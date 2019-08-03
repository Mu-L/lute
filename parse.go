// Lute - A structured markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lute

import (
	"regexp"
)

func Parse(name, text string) (*Tree, error) {
	// TODO: 移到 lex 阶段
	text = sanitize(text)

	t := &Tree{Name: name, text: text, context: &Context{}}
	err := t.parse()

	return t, err
}

var newlinesRegexp = regexp.MustCompile("\r[\n\u0085]?|[\u2424\u2028\u0085]")
var nullRegexp = regexp.MustCompile("\u0000")

func sanitize(text string) (ret string) {
	ret = newlinesRegexp.ReplaceAllString(text, "\n")
	nullRegexp.ReplaceAllString(ret, "\uFFFD")

	return
}

type BlockContainer struct {
	nodes []Node
}

func (bc *BlockContainer) push(node Node) {
	bc.nodes = append(bc.nodes, node)
}

func (bc *BlockContainer) pop() (ret Node) {
	length := len(bc.nodes)
	if 1 > length {
		return nil
	}

	ret = bc.nodes[length-1]
	bc.nodes = bc.nodes[:length-1]

	return
}

func (bc *BlockContainer) peek() Node {
	length := len(bc.nodes)
	if 1 > length {
		return nil
	}

	return bc.nodes[length-1]
}

// Context use to store common data in parsing.
type Context struct {
	linkRefDef map[string]*Link
	curLines   []items

	// Blocks parsing

	tip                                                      Node
	oldtip                                                   Node
	currentLine                                              items
	currentLineLen                                           int
	offset, column, nextNonspace, nextNonspaceColumn, indent int
	indented, blank, partiallyConsumedTab, allClosed         bool
	lastMatchedContainer                                     Node

	// Inlines parsing

	pos               int
	delimiters        *delimiter
	brackets          *delimiter
	previousDelimiter *delimiter
}

func (context *Context) advanceOffset(count int, columns bool) {
	var currentLine = context.currentLine
	var charsToTab, charsToAdvance int
	var c *item
	for ; count > 0; {
		c = currentLine[context.offset]
		if c.isTab() {
			charsToTab = 4 - (context.column % 4)
			if columns {
				context.partiallyConsumedTab = charsToTab > count
				if charsToTab > count {
					charsToAdvance = count
				} else {
					charsToAdvance = charsToTab
				}
				context.column += charsToAdvance
				if !context.partiallyConsumedTab {
					context.offset += 1
				}
				count -= charsToAdvance
			} else {
				context.partiallyConsumedTab = false
				context.column += charsToTab
				context.offset += 1
				count -= 1
			}
		} else {
			context.partiallyConsumedTab = false
			context.offset += 1
			context.column += 1 // assume ascii; block starts are ascii
			count -= 1
		}
	}
}

func (context *Context) advanceNextNonspace() {
	context.offset = context.nextNonspace
	context.column = context.nextNonspaceColumn
	context.partiallyConsumedTab = false
}

func (context *Context) findNextNonspace() {
	i := context.offset
	cols := context.column

	var c *item
	for {
		c = context.currentLine[i]
		if c.isSpace() {
			i++
			cols++
		} else if c.isTab() {
			i++
			cols += 4 - (cols % 4)
		} else {
			break
		}
	}
	context.blank = c.val == "\n" || c.val == "\r" || "" == c.val
	context.nextNonspace = i
	context.nextNonspaceColumn = cols
	context.indent = context.nextNonspaceColumn - context.column
	context.indented = context.indent >= 4
}

// Finalize and close any unmatched blocks.
func (context *Context) closeUnmatchedBlocks() {
	if !context.allClosed {
		// finalize any blocks not matched
		for context.oldtip != context.lastMatchedContainer {
			parent := context.oldtip.Parent()
			context.finalize(context.oldtip)
			context.oldtip = parent
		}
		context.allClosed = true
	}
}

// Finalize a block.  Close it and do any necessary postprocessing,
// e.g. creating string_content from strings, setting the 'tight'
// or 'loose' status of a list, and parsing the beginnings
// of paragraphs for reference definitions.  Reset the tip to the
// parent of the closed block.
func (context *Context) finalize(block Node) {
	var parent = block.Parent()
	block.Close()
	block.Finalize(context)
	context.tip = parent
}

// Add block of type tag as a child of the tip.  If the tip can't
// accept children, close and finalize it and try its parent,
// and so on til we find a block that can accept children.
func (context *Context) addChild(child Node) {
	for !context.tip.CanContain(child.Type()) {
		context.finalize(context.tip)
	}

	context.tip.AppendChild(context.tip, child)
	context.tip = child
}

// Returns true if the two list items are of the same type,
// with the same delimiter and bullet character.  This is used
// in agglomerating list items into lists.
func (context *Context) listsMatch(list_data, item_data *ListData) bool {
	return list_data.typ == item_data.typ &&
		list_data.delimiter == item_data.delimiter &&
		list_data.bulletChar == item_data.bulletChar
}

// Tree is the representation of the markdown ast.
type Tree struct {
	Root      *Root
	Name      string
	text      string
	lex       *lexer
	peekCount int
	context   *Context
}

func (t *Tree) Render(renderer *Renderer) (output string, err error) {
	err = renderer.Render(t.Root)
	if nil != err {
		return "", err
	}
	output = renderer.writer.String()

	return
}

func (t *Tree) nonSpaceTab(tokens items) (spaces, tabs int, remains items) {
	i := 0
Loop:
	for ; i < len(tokens); i++ {
		token := tokens[i]
		switch token.typ {
		case itemTab:
			tabs++
		case itemSpace:
			spaces++
		default:
			break Loop
		}
	}

	remains = tokens[i:]

	return
}

func (t *Tree) skipBlankLines() (blankLines []items) {
	for {
		line := t.nextLine()
		if line.isEOF() {
			return
		}

		//tokens := t.removeStartBlockquoteMarker(line, t.context.BlockquoteLevel)
		if !line.isBlankLine() {
			t.backupLine(line)
			return
		}

		blankLines = append(blankLines, line)
	}
}

func (t *Tree) indentOffset(tokens items, indentSpaces int) (ret items) {
	var nonWhitespaces items
	compSpaces := 0
	i := 0
	for ; i < len(tokens); i++ {
		typ := tokens[i].typ
		if itemSpace == typ {
			compSpaces++
		} else if itemTab == typ {
			compSpaces += 4
		} else {
			nonWhitespaces = append(nonWhitespaces, tokens[i:]...)
			break
		}
	}

	remains := compSpaces - indentSpaces
	if 0 > remains {
		return nonWhitespaces
	}

	for j := 0; j < remains; j++ {
		ret = append(ret, &item{itemSpace, 0, " ", 0})
	}
	ret = append(ret, nonWhitespaces...)

	return
}

func (t *Tree) nextLine() (line items) {
	length := len(t.context.curLines)
	if 0 < length {
		line = t.context.curLines[0]
		t.context.curLines = t.context.curLines[1:]

		return
	}

	for {
		token := t.lex.nextItem()
		line = append(line, token)
		if token.isNewline() || token.isEOF() {
			return
		}
	}
}

func (t *Tree) backupLine(line items) {
	if 0 < len(t.context.curLines) {
		t.context.curLines = append([]items{line}, t.context.curLines...)
	} else {
		t.context.curLines = append(t.context.curLines, line)
	}
}

// Parsing.

// recover is the handler that turns panics into returns from the top level of Parse.
func (t *Tree) recover(err *error) {
	e := recover()
	if e != nil {
		*err = e.(error)
	}
}

func (t *Tree) parse() (err error) {
	defer t.recover(&err)

	t.lex = lex(t.Name, t.text)
	t.Root = &Root{&BaseNode{typ: NodeRoot}}
	t.parseBlocks()
	t.parseInlines()
	t.lex = nil

	return nil
}
