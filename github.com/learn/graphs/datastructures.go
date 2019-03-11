package graphs

import (
	"algorithms"
	"fmt"
	"strings"
)

type Key string
type PropertyValue interface{}
type GraphProperties map[Key]PropertyValue

type vertex struct {
	id         algorithms.UUID
	properties GraphProperties
}

//Edge this is the structure for the edge
type edge struct {
	id         algorithms.UUID
	fromVertex *vertex
	toVertex   *vertex
	directed   bool
	properties GraphProperties
}

type Graph struct {
	edges  []edge
	vertex []vertex
}

func CreateGraph() *Graph {

	edges := make([]edge, 0, 100)
	vertices := make([]vertex, 0, 100)

	return &Graph{
		edges:  edges,
		vertex: vertices,
	}
}

func (graph *Graph) getVertexList() []vertex {
	return graph.vertex
}

func (graph *Graph) getEdgeList() []edge {
	return graph.edges
}

func (graph *Graph) AddVertex(vertex vertex) *Graph {
	graph.vertex = append(graph.getVertexList(), vertex)
	return graph
}

func (graph *Graph) AddEdge(edge edge) *Graph {
	graph.edges = append(graph.getEdgeList(), edge)
	return graph
}

func (graph Graph) String() string {

	vertexString := make([]string, 0)
	vertices := graph.getVertexList()
	for i := 0; i < len(vertices); i++ {
		vertexString = append(vertexString, vertices[i].String())
	}

	edgeString := make([]string, 0)
	edges := graph.getEdgeList()
	for i := 0; i < len(edges); i++ {
		edgeString = append(edgeString, edges[i].String())
	}

	return fmt.Sprintf("Graph( \n vertices : [%s], \n edges : [%v] \n)", strings.Join(vertexString, ","), strings.Join(edgeString, ","))
}

type GraphMethods interface {
	String() string
}

func (vertex vertex) String() string {
	return fmt.Sprintf("Vertex( id : %s, properties : %v)", vertex.id, vertex.properties)
}

func CreateVertex(vertexProp GraphProperties) *vertex {
	return &vertex{
		id:         algorithms.GenerateUUID(),
		properties: vertexProp}
}

func (edge edge) String() string {
	return fmt.Sprintf("Edge( fromVertex : %s, toVertex: %s, properties : %v)", edge.fromVertex, edge.toVertex, edge.properties)
}

func createEdge(edgeProp GraphProperties, fromVertex *vertex, toVertex *vertex, directed bool) *edge {
	return &edge{
		id:         algorithms.GenerateUUID(),
		fromVertex: fromVertex,
		toVertex:   toVertex,
		directed:   true,
		properties: edgeProp}
}

//CreateDirectedEdge this is a function to created a directed graph
func CreateDirectedEdge(edgeProp GraphProperties, fromVertex *vertex, toVertex *vertex) *edge {
	return createEdge(edgeProp, fromVertex, toVertex, true)
}
