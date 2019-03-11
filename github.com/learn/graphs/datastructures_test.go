package graphs

import (
	"log"
	"testing"
)

func TestCreateVertex(t *testing.T) {
	prop := make(GraphProperties)
	key := Key("name")
	prop[key] = "testvertex"
	vertex := CreateVertex(prop)
	log.Println(vertex)
}

func TestCreateDirectedEdge(t *testing.T) {
	propFrom := make(GraphProperties)
	keyFrom := Key("name")
	propFrom[keyFrom] = "testvertexFrom"
	vertexFrom := CreateVertex(propFrom)

	propTo := make(GraphProperties)
	keyTo := Key("name")
	propTo[keyTo] = "testvertexTo"
	vertexTo := CreateVertex(propTo)

	prop := make(GraphProperties)
	key := Key("name")
	prop[key] = "testedge"

	edge := CreateDirectedEdge(prop, vertexFrom, vertexTo)
	log.Println(edge)
}

func BenchmarkCreateVertex(t *testing.B) {
	for i := 0; i < t.N; i++ {
		prop := make(GraphProperties)
		key := Key("name")
		prop[key] = "testvertex"

		CreateVertex(prop)
		//log.Println(vertex)
	}
}

func BenchmarkCreateDirectedEdge(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		propFrom := make(GraphProperties)
		keyFrom := Key("name")
		propFrom[keyFrom] = "testvertexFrom"
		vertexFrom := CreateVertex(propFrom)

		propTo := make(GraphProperties)
		keyTo := Key("name")
		propTo[keyTo] = "testvertexTo"
		vertexTo := CreateVertex(propTo)

		prop := make(GraphProperties)
		key := Key("name")
		prop[key] = "testedge"

		CreateDirectedEdge(prop, vertexFrom, vertexTo)
		//log.Println(edge)
	}
}

func BenchmarkParallelEdgeCreate(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			propFrom := make(GraphProperties)
			keyFrom := Key("name")
			propFrom[keyFrom] = "testvertexFrom"
			vertexFrom := CreateVertex(propFrom)

			propTo := make(GraphProperties)
			keyTo := Key("name")
			propTo[keyTo] = "testvertexTo"
			vertexTo := CreateVertex(propTo)

			prop := make(GraphProperties)
			key := Key("name")
			prop[key] = "testedge"

			CreateDirectedEdge(prop, vertexFrom, vertexTo)
		}
	})
}

func TestCreateGraph(t *testing.T) {
	graph := CreateGraph()

	propFromEd := make(GraphProperties)
	keyFromEd := Key("name")
	propFromEd[keyFromEd] = "testvertexFromEd"
	vertexFromEd := CreateVertex(propFromEd)

	propFrom := make(GraphProperties)
	keyFrom := Key("name")
	propFrom[keyFrom] = "testvertexFrom"
	vertexFrom := CreateVertex(propFrom)

	propTo := make(GraphProperties)
	keyTo := Key("name")
	propTo[keyTo] = "testvertexTo"
	vertexTo := CreateVertex(propTo)

	prop := make(GraphProperties)
	key := Key("name")
	prop[key] = "testedge"

	edge := CreateDirectedEdge(prop, vertexFrom, vertexTo)

	graph.AddVertex(*vertexFromEd).AddVertex(*vertexFrom).AddVertex(*vertexTo).AddEdge(*edge)

	log.Println(*graph)
}
