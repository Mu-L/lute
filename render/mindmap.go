// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package render

import "github.com/88250/lute/parse"

func RenderMindmap(listContent string) string {
	tree := parse.Parse("", []byte(listContent), nil)


	_ = tree
	return "{\n \"name\": \"flare\",\n \"children\": [\n  {\n   \"name\": \"analytics\",\n   \"children\": [\n    {\n     \"name\": \"cluster\",\n     \"children\": [\n      {\"name\": \"AgglomerativeCluster\", \"value\": 3938},\n      {\"name\": \"CommunityStructure\", \"value\": 3812},\n      {\"name\": \"HierarchicalCluster\", \"value\": 6714},\n      {\"name\": \"MergeEdge\", \"value\": 743}\n     ]\n    },\n    {\n     \"name\": \"graph\",\n     \"children\": [\n      {\"name\": \"BetweennessCentrality\", \"value\": 3534},\n      {\"name\": \"LinkDistance\", \"value\": 5731},\n      {\"name\": \"MaxFlowMinCut\", \"value\": 7840},\n      {\"name\": \"ShortestPaths\", \"value\": 5914},\n      {\"name\": \"SpanningTree\", \"value\": 3416}\n     ]\n    },\n    {\n     \"name\": \"optimization\",\n     \"children\": [\n      {\"name\": \"AspectRatioBanker\", \"value\": 7074}\n     ]\n    }\n   ]\n  },\n  {\n   \"name\": \"animate\",\n   \"children\": [\n    {\"name\": \"Easing\", \"value\": 17010},\n    {\"name\": \"FunctionSequence\", \"value\": 5842},\n    {\n     \"name\": \"interpolate\",\n     \"children\": [\n      {\"name\": \"ArrayInterpolator\", \"value\": 1983},\n      {\"name\": \"ColorInterpolator\", \"value\": 2047},\n      {\"name\": \"DateInterpolator\", \"value\": 1375},\n      {\"name\": \"Interpolator\", \"value\": 8746},\n      {\"name\": \"MatrixInterpolator\", \"value\": 2202},\n      {\"name\": \"NumberInterpolator\", \"value\": 1382},\n      {\"name\": \"ObjectInterpolator\", \"value\": 1629},\n      {\"name\": \"PointInterpolator\", \"value\": 1675},\n      {\"name\": \"RectangleInterpolator\", \"value\": 2042}\n     ]\n    },\n    {\"name\": \"ISchedulable\", \"value\": 1041},\n    {\"name\": \"Parallel\", \"value\": 5176},\n    {\"name\": \"Pause\", \"value\": 449},\n    {\"name\": \"Scheduler\", \"value\": 5593},\n    {\"name\": \"Sequence\", \"value\": 5534},\n    {\"name\": \"Transition\", \"value\": 9201},\n    {\"name\": \"Transitioner\", \"value\": 19975},\n    {\"name\": \"TransitionEvent\", \"value\": 1116},\n    {\"name\": \"Tween\", \"value\": 6006}\n   ]\n  },\n  {\n   \"name\": \"data\",\n   \"children\": [\n    {\n     \"name\": \"converters\",\n     \"children\": [\n      {\"name\": \"Converters\", \"value\": 721},\n      {\"name\": \"DelimitedTextConverter\", \"value\": 4294},\n      {\"name\": \"GraphMLConverter\", \"value\": 9800},\n      {\"name\": \"IDataConverter\", \"value\": 1314},\n      {\"name\": \"JSONConverter\", \"value\": 2220}\n     ]\n    },\n    {\"name\": \"DataField\", \"value\": 1759},\n    {\"name\": \"DataSchema\", \"value\": 2165},\n    {\"name\": \"DataSet\", \"value\": 586},\n    {\"name\": \"DataSource\", \"value\": 3331},\n    {\"name\": \"DataTable\", \"value\": 772},\n    {\"name\": \"DataUtil\", \"value\": 3322}\n   ]\n  },\n  {\n   \"name\": \"display\",\n   \"children\": [\n    {\"name\": \"DirtySprite\", \"value\": 8833},\n    {\"name\": \"LineSprite\", \"value\": 1732},\n    {\"name\": \"RectSprite\", \"value\": 3623},\n    {\"name\": \"TextSprite\", \"value\": 10066}\n   ]\n  },\n  {\n   \"name\": \"flex\",\n   \"children\": [\n    {\"name\": \"FlareVis\", \"value\": 4116}\n   ]\n  },\n  {\n   \"name\": \"physics\",\n   \"children\": [\n    {\"name\": \"DragForce\", \"value\": 1082},\n    {\"name\": \"GravityForce\", \"value\": 1336},\n    {\"name\": \"IForce\", \"value\": 319},\n    {\"name\": \"NBodyForce\", \"value\": 10498},\n    {\"name\": \"Particle\", \"value\": 2822},\n    {\"name\": \"Simulation\", \"value\": 9983},\n    {\"name\": \"Spring\", \"value\": 2213},\n    {\"name\": \"SpringForce\", \"value\": 1681}\n   ]\n  },\n  {\n   \"name\": \"query\",\n   \"children\": [\n    {\"name\": \"AggregateExpression\", \"value\": 1616},\n    {\"name\": \"And\", \"value\": 1027},\n    {\"name\": \"Arithmetic\", \"value\": 3891},\n    {\"name\": \"Average\", \"value\": 891},\n    {\"name\": \"BinaryExpression\", \"value\": 2893},\n    {\"name\": \"Comparison\", \"value\": 5103},\n    {\"name\": \"CompositeExpression\", \"value\": 3677},\n    {\"name\": \"Count\", \"value\": 781},\n    {\"name\": \"DateUtil\", \"value\": 4141},\n    {\"name\": \"Distinct\", \"value\": 933},\n    {\"name\": \"Expression\", \"value\": 5130},\n    {\"name\": \"ExpressionIterator\", \"value\": 3617},\n    {\"name\": \"Fn\", \"value\": 3240},\n    {\"name\": \"If\", \"value\": 2732},\n    {\"name\": \"IsA\", \"value\": 2039},\n    {\"name\": \"Literal\", \"value\": 1214},\n    {\"name\": \"Match\", \"value\": 3748},\n    {\"name\": \"Maximum\", \"value\": 843},\n    {\n     \"name\": \"methods\",\n     \"children\": [\n      {\"name\": \"add\", \"value\": 593},\n      {\"name\": \"and\", \"value\": 330},\n      {\"name\": \"average\", \"value\": 287},\n      {\"name\": \"count\", \"value\": 277},\n      {\"name\": \"distinct\", \"value\": 292},\n      {\"name\": \"div\", \"value\": 595},\n      {\"name\": \"eq\", \"value\": 594},\n      {\"name\": \"fn\", \"value\": 460},\n      {\"name\": \"gt\", \"value\": 603},\n      {\"name\": \"gte\", \"value\": 625},\n      {\"name\": \"iff\", \"value\": 748},\n      {\"name\": \"isa\", \"value\": 461},\n      {\"name\": \"lt\", \"value\": 597},\n      {\"name\": \"lte\", \"value\": 619},\n      {\"name\": \"max\", \"value\": 283},\n      {\"name\": \"min\", \"value\": 283},\n      {\"name\": \"mod\", \"value\": 591},\n      {\"name\": \"mul\", \"value\": 603},\n      {\"name\": \"neq\", \"value\": 599},\n      {\"name\": \"not\", \"value\": 386},\n      {\"name\": \"or\", \"value\": 323},\n      {\"name\": \"orderby\", \"value\": 307},\n      {\"name\": \"range\", \"value\": 772},\n      {\"name\": \"select\", \"value\": 296},\n      {\"name\": \"stddev\", \"value\": 363},\n      {\"name\": \"sub\", \"value\": 600},\n      {\"name\": \"sum\", \"value\": 280},\n      {\"name\": \"update\", \"value\": 307},\n      {\"name\": \"variance\", \"value\": 335},\n      {\"name\": \"where\", \"value\": 299},\n      {\"name\": \"xor\", \"value\": 354},\n      {\"name\": \"-\", \"value\": 264}\n     ]\n    },\n    {\"name\": \"Minimum\", \"value\": 843},\n    {\"name\": \"Not\", \"value\": 1554},\n    {\"name\": \"Or\", \"value\": 970},\n    {\"name\": \"Query\", \"value\": 13896},\n    {\"name\": \"Range\", \"value\": 1594},\n    {\"name\": \"StringUtil\", \"value\": 4130},\n    {\"name\": \"Sum\", \"value\": 791},\n    {\"name\": \"Variable\", \"value\": 1124},\n    {\"name\": \"Variance\", \"value\": 1876},\n    {\"name\": \"Xor\", \"value\": 1101}\n   ]\n  },\n  {\n   \"name\": \"scale\",\n   \"children\": [\n    {\"name\": \"IScaleMap\", \"value\": 2105},\n    {\"name\": \"LinearScale\", \"value\": 1316},\n    {\"name\": \"LogScale\", \"value\": 3151},\n    {\"name\": \"OrdinalScale\", \"value\": 3770},\n    {\"name\": \"QuantileScale\", \"value\": 2435},\n    {\"name\": \"QuantitativeScale\", \"value\": 4839},\n    {\"name\": \"RootScale\", \"value\": 1756},\n    {\"name\": \"Scale\", \"value\": 4268},\n    {\"name\": \"ScaleType\", \"value\": 1821},\n    {\"name\": \"TimeScale\", \"value\": 5833}\n   ]\n  },\n  {\n   \"name\": \"util\",\n   \"children\": [\n    {\"name\": \"Arrays\", \"value\": 8258},\n    {\"name\": \"Colors\", \"value\": 10001},\n    {\"name\": \"Dates\", \"value\": 8217},\n    {\"name\": \"Displays\", \"value\": 12555},\n    {\"name\": \"Filter\", \"value\": 2324},\n    {\"name\": \"Geometry\", \"value\": 10993},\n    {\n     \"name\": \"heap\",\n     \"children\": [\n      {\"name\": \"FibonacciHeap\", \"value\": 9354},\n      {\"name\": \"HeapNode\", \"value\": 1233}\n     ]\n    },\n    {\"name\": \"IEvaluable\", \"value\": 335},\n    {\"name\": \"IPredicate\", \"value\": 383},\n    {\"name\": \"IValueProxy\", \"value\": 874},\n    {\n     \"name\": \"math\",\n     \"children\": [\n      {\"name\": \"DenseMatrix\", \"value\": 3165},\n      {\"name\": \"IMatrix\", \"value\": 2815},\n      {\"name\": \"SparseMatrix\", \"value\": 3366}\n     ]\n    },\n    {\"name\": \"Maths\", \"value\": 17705},\n    {\"name\": \"Orientation\", \"value\": 1486},\n    {\n     \"name\": \"palette\",\n     \"children\": [\n      {\"name\": \"ColorPalette\", \"value\": 6367},\n      {\"name\": \"Palette\", \"value\": 1229},\n      {\"name\": \"ShapePalette\", \"value\": 2059},\n      {\"name\": \"SizePalette\", \"value\": 2291}\n     ]\n    },\n    {\"name\": \"Property\", \"value\": 5559},\n    {\"name\": \"Shapes\", \"value\": 19118},\n    {\"name\": \"Sort\", \"value\": 6887},\n    {\"name\": \"Stats\", \"value\": 6557},\n    {\"name\": \"Strings\", \"value\": 22026}\n   ]\n  },\n  {\n   \"name\": \"vis\",\n   \"children\": [\n    {\n     \"name\": \"axis\",\n     \"children\": [\n      {\"name\": \"Axes\", \"value\": 1302},\n      {\"name\": \"Axis\", \"value\": 24593},\n      {\"name\": \"AxisGridLine\", \"value\": 652},\n      {\"name\": \"AxisLabel\", \"value\": 636},\n      {\"name\": \"CartesianAxes\", \"value\": 6703}\n     ]\n    },\n    {\n     \"name\": \"controls\",\n     \"children\": [\n      {\"name\": \"AnchorControl\", \"value\": 2138},\n      {\"name\": \"ClickControl\", \"value\": 3824},\n      {\"name\": \"Control\", \"value\": 1353},\n      {\"name\": \"ControlList\", \"value\": 4665},\n      {\"name\": \"DragControl\", \"value\": 2649},\n      {\"name\": \"ExpandControl\", \"value\": 2832},\n      {\"name\": \"HoverControl\", \"value\": 4896},\n      {\"name\": \"IControl\", \"value\": 763},\n      {\"name\": \"PanZoomControl\", \"value\": 5222},\n      {\"name\": \"SelectionControl\", \"value\": 7862},\n      {\"name\": \"TooltipControl\", \"value\": 8435}\n     ]\n    },\n    {\n     \"name\": \"data\",\n     \"children\": [\n      {\"name\": \"Data\", \"value\": 20544},\n      {\"name\": \"DataList\", \"value\": 19788},\n      {\"name\": \"DataSprite\", \"value\": 10349},\n      {\"name\": \"EdgeSprite\", \"value\": 3301},\n      {\"name\": \"NodeSprite\", \"value\": 19382},\n      {\n       \"name\": \"render\",\n       \"children\": [\n        {\"name\": \"ArrowType\", \"value\": 698},\n        {\"name\": \"EdgeRenderer\", \"value\": 5569},\n        {\"name\": \"IRenderer\", \"value\": 353},\n        {\"name\": \"ShapeRenderer\", \"value\": 2247}\n       ]\n      },\n      {\"name\": \"ScaleBinding\", \"value\": 11275},\n      {\"name\": \"Tree\", \"value\": 7147},\n      {\"name\": \"TreeBuilder\", \"value\": 9930}\n     ]\n    },\n    {\n     \"name\": \"events\",\n     \"children\": [\n      {\"name\": \"DataEvent\", \"value\": 2313},\n      {\"name\": \"SelectionEvent\", \"value\": 1880},\n      {\"name\": \"TooltipEvent\", \"value\": 1701},\n      {\"name\": \"VisualizationEvent\", \"value\": 1117}\n     ]\n    },\n    {\n     \"name\": \"legend\",\n     \"children\": [\n      {\"name\": \"Legend\", \"value\": 20859},\n      {\"name\": \"LegendItem\", \"value\": 4614},\n      {\"name\": \"LegendRange\", \"value\": 10530}\n     ]\n    },\n    {\n     \"name\": \"operator\",\n     \"children\": [\n      {\n       \"name\": \"distortion\",\n       \"children\": [\n        {\"name\": \"BifocalDistortion\", \"value\": 4461},\n        {\"name\": \"Distortion\", \"value\": 6314},\n        {\"name\": \"FisheyeDistortion\", \"value\": 3444}\n       ]\n      },\n      {\n       \"name\": \"encoder\",\n       \"children\": [\n        {\"name\": \"ColorEncoder\", \"value\": 3179},\n        {\"name\": \"Encoder\", \"value\": 4060},\n        {\"name\": \"PropertyEncoder\", \"value\": 4138},\n        {\"name\": \"ShapeEncoder\", \"value\": 1690},\n        {\"name\": \"SizeEncoder\", \"value\": 1830}\n       ]\n      },\n      {\n       \"name\": \"filter\",\n       \"children\": [\n        {\"name\": \"FisheyeTreeFilter\", \"value\": 5219},\n        {\"name\": \"GraphDistanceFilter\", \"value\": 3165},\n        {\"name\": \"VisibilityFilter\", \"value\": 3509}\n       ]\n      },\n      {\"name\": \"IOperator\", \"value\": 1286},\n      {\n       \"name\": \"label\",\n       \"children\": [\n        {\"name\": \"Labeler\", \"value\": 9956},\n        {\"name\": \"RadialLabeler\", \"value\": 3899},\n        {\"name\": \"StackedAreaLabeler\", \"value\": 3202}\n       ]\n      },\n      {\n       \"name\": \"layout\",\n       \"children\": [\n        {\"name\": \"AxisLayout\", \"value\": 6725},\n        {\"name\": \"BundledEdgeRouter\", \"value\": 3727},\n        {\"name\": \"CircleLayout\", \"value\": 9317},\n        {\"name\": \"CirclePackingLayout\", \"value\": 12003},\n        {\"name\": \"DendrogramLayout\", \"value\": 4853},\n        {\"name\": \"ForceDirectedLayout\", \"value\": 8411},\n        {\"name\": \"IcicleTreeLayout\", \"value\": 4864},\n        {\"name\": \"IndentedTreeLayout\", \"value\": 3174},\n        {\"name\": \"Layout\", \"value\": 7881},\n        {\"name\": \"NodeLinkTreeLayout\", \"value\": 12870},\n        {\"name\": \"PieLayout\", \"value\": 2728},\n        {\"name\": \"RadialTreeLayout\", \"value\": 12348},\n        {\"name\": \"RandomLayout\", \"value\": 870},\n        {\"name\": \"StackedAreaLayout\", \"value\": 9121},\n        {\"name\": \"TreeMapLayout\", \"value\": 9191}\n       ]\n      },\n      {\"name\": \"Operator\", \"value\": 2490},\n      {\"name\": \"OperatorList\", \"value\": 5248},\n      {\"name\": \"OperatorSequence\", \"value\": 4190},\n      {\"name\": \"OperatorSwitch\", \"value\": 2581},\n      {\"name\": \"SortOperator\", \"value\": 2023}\n     ]\n    },\n    {\"name\": \"Visualization\", \"value\": 16540}\n   ]\n  }\n ]\n}\n"
}

