package main

//https://codetree.dev/golang-graph-traversal/

func main() {
	g := NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 1)
}

type Vertex struct {
	// Key is the unique identifier of the vertex
	Key int
	// Vertices will describe vertices connected to this one
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
}

// We then create a constructor function for the Vertex

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
	// This will decide if it's a directed or undirected graph
	directed bool
}

// We defined constructor functions that create
// new directed or undirected graphs respectively

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

// AddVertex creates a new vertex with the given
// key and adds it to the graph
func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

// here, we import the graph we defined in the previous section as the `graph` package
func DFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	// we maintain a map of visited nodes to prevent visiting the same
	// node more than once
	visited := map[int]bool{}

	if startVertex == nil {
		return
	}
	visited[startVertex.Key] = true
	visitCb(startVertex.Key)

	// for each of the adjacent vertices, call the function recursively
	// if it hasn't yet been visited
	for _, v := range startVertex.Vertices {
		if visited[v.Key] {
			continue
		}
		DFS(g, v, visitCb)
	}
}

// create a node that holds the graphs vertex as data
type node struct {
	v    *Vertex
	next *node
}

// create a queue data structure
type queue struct {
	head *node
	tail *node
}

// enqueue adds a new node to the tail of the queue
func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}

	// if the queue is empty, set the head and tail as the node value
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

// dequeue removes the head from the queue and returns it
func (q *queue) dequeue() *Vertex {
	n := q.head
	// return nil, if head is empty
	if n == nil {
		return nil
	}

	q.head = q.head.next

	// if there wasn't any next node, that
	// means the queue is empty, and the tail
	// should be set to nil
	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

func BFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	// initialize queue and visited vertices map
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}

	currentVertex := startVertex
	// start a continuous loop
	for {
		// visit the current node
		visitCb(currentVertex.Key)
		visitedVertices[currentVertex.Key] = true

		// for each neighboring vertex, push it to the queue
		// if it isn't already visited
		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}

		// change the current vertex to the next one
		// in the queue
		currentVertex = vertexQueue.dequeue()
		// if the queue is empty, break out of the loop
		if currentVertex == nil {
			break
		}
	}
}
