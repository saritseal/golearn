package graphs

import (
	"algorithms"
	"fmt"
	"log"
	"strings"
)

type graphErrors struct {
	errorCode int
	context   map[string]interface{}
	message   string
}

func (err graphErrors) Error() string {
	return ""
}

type Key string
type PropertyValue interface{}
type GraphProperties map[Key]PropertyValue

type vertex struct {
	id         algorithms.UUID
	properties GraphProperties
	connectons []*vertex
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
	edges  []*edge
	vertex []*vertex
}

func CreateGraph() *Graph {

	edges := make([]*edge, 0, 100)
	vertices := make([]*vertex, 0, 100)

	return &Graph{
		edges:  edges,
		vertex: vertices,
	}
}

func (graph *Graph) getVertexList() []*vertex {
	return graph.vertex
}

func (graph *Graph) getEdgeList() []*edge {
	return graph.edges
}

func (graph *Graph) FindVertex(id algorithms.UUID) (*vertex, error) {

	for _, vtx := range graph.getVertexList() {
		if vtx.id == id {
			return vtx, nil
		}
	}
	return nil, graphErrors{100, map[string]interface{}{"vertexId": id}, "vertex not found with the id"}
}

func (graph *Graph) FindAllEdges(id algorithms.UUID) ([]*edge, error) {
	edges := make([]*edge, 0)

	for _, edge := range graph.getEdgeList() {
		if edge.fromVertex.id == id {
			edges = append(edges, edge)
		}
	}
	if len(edges) != 0 {
		return edges, nil
	} else {
		return nil, graphErrors{100, map[string]interface{}{"vertexId": id}, "there are no edges for the vertices id"}
	}
}

func (graph *Graph) AddVertex(vertex *vertex) *Graph {
	graph.vertex = append(graph.getVertexList(), vertex)
	return graph
}

func (graph *Graph) AddEdge(edge *edge) *Graph {
	graph.edges = append(graph.getEdgeList(), edge)
	return graph
}

func (graph Graph) String() string {

	vertexString := make([]string, 0)
	vertices := graph.getVertexList()
	for i := 0; i < len(vertices); i++ {
		vertexString = append(vertexString, vertices[i].String()+"\n")
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
	GetVertex(id algorithms.UUID) *vertex
}

func (vertex *vertex) GetVertexConnections() []algorithms.UUID {
	vertices := make([]algorithms.UUID, 0)

	for _, vtx := range vertex.connectons {
		vertices = append(vertices, vtx.id)
	}

	return vertices
}

func (vertex *vertex) String() string {
	return fmt.Sprintf("Vertex( id : %s, properties : %v, connections: [%s])", vertex.id, vertex.properties,
		fmt.Sprintf("%s", vertex.GetVertexConnections()))
}

func (vertex *vertex) AddConnection(toVertex *vertex) {
	vertex.connectons = append(vertex.connectons, toVertex)
}

func CreateVertex(vertexProp GraphProperties) *vertex {
	return &vertex{
		id:         algorithms.GenerateUUID(),
		properties: vertexProp,
		connectons: make([]*vertex, 0),
	}
}

func (edge edge) String() string {
	return fmt.Sprintf("Edge( fromVertex : %s, toVertex: %s, properties : %v)", edge.fromVertex, edge.toVertex, edge.properties)
}

func createEdge(edgeProp GraphProperties, fromVertex *vertex, toVertex *vertex, directed bool) *edge {

	fromVertex.AddConnection(toVertex)
	log.Println(fromVertex.id, fromVertex)

	toVertex.AddConnection(fromVertex)
	log.Println(toVertex.id, toVertex)

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
