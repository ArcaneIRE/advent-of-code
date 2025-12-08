package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))

	fmt.Println(solve_part_one(input))

	fmt.Println(solve_part_two(input))
}

type Coordinate struct {
	x, y, z int
}

type Connection struct {
	distance    float64
	source      Coordinate
	destination Coordinate
}

type ConnectionHeap []Connection

func (h ConnectionHeap) Len() int           { return len(h) }
func (h ConnectionHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h ConnectionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ConnectionHeap) Push(x any)        { *h = append(*h, x.(Connection)) }
func (h *ConnectionHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func create_coordinate_list(input string) []Coordinate {
	var junction_boxes []Coordinate
	for line := range strings.SplitSeq(input[:len(input)-2], "\n") {
		values := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])

		junction_boxes = append(junction_boxes, Coordinate{
			x: x,
			y: y,
			z: z,
		})
	}
	return junction_boxes
}

func (source Coordinate) calculate_euclidean_distance(destination Coordinate) float64 {
	dx := source.x - destination.x
	dy := source.y - destination.y
	dz := source.z - destination.z

	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func calculate_connection_distances(junction_boxes []Coordinate) ConnectionHeap {
	var connection_heap ConnectionHeap
	for src_idx := range junction_boxes { // For every junction box
		for dest_idx := src_idx + 1; dest_idx < len(junction_boxes); dest_idx++ { // Create connection to every other, skipping existing ones
			connection := Connection{
				source:      junction_boxes[src_idx],
				destination: junction_boxes[dest_idx],
				distance:    junction_boxes[src_idx].calculate_euclidean_distance(junction_boxes[dest_idx]),
			}
			heap.Push(&connection_heap, connection)
		}
	}
	return connection_heap
}

func get_circuit_sizes(adjacency_list map[Coordinate][]Coordinate) []int {
	var circuit_sizes []int
	visited := make(map[Coordinate]bool)
	for key := range adjacency_list {
		_, is_visited := visited[key]
		if is_visited {
			continue
		}
		size := 0
		to_visit := []Coordinate{key}
		for i := 0; i < len(to_visit); i++ { // While adjacency list is not empty
			next := to_visit[i]
			_, is_visited := visited[next]
			if !is_visited {
				size++
				visited[next] = true
				to_visit = append(to_visit, adjacency_list[next]...)
			}
		}
		circuit_sizes = append(circuit_sizes, size)
	}
	return circuit_sizes
}

func solve_part_one(input string) string {
	const NUM_CONNECTIONS int = 1000

	junction_boxes := create_coordinate_list(input)
	connection_heap := calculate_connection_distances(junction_boxes)

	// Created graph with edges for 1000 lowest connections
	adjacency_list := make(map[Coordinate][]Coordinate, 1000)
	for range NUM_CONNECTIONS {
		connection := heap.Pop(&connection_heap).(Connection)
		adjacency_list[connection.source] = append(adjacency_list[connection.source], connection.destination)
		adjacency_list[connection.destination] = append(adjacency_list[connection.destination], connection.source)
	}

	circuit_sizes := get_circuit_sizes(adjacency_list)

	// Return the product of the three largest circuits
	sort.Sort(sort.Reverse(sort.IntSlice(circuit_sizes)))
	return strconv.Itoa(circuit_sizes[0] * circuit_sizes[1] * circuit_sizes[2])
}

func is_completed_circuit(adjacency_list map[Coordinate][]Coordinate, max int) bool {
	circuit_sizes := get_circuit_sizes(adjacency_list)
	if len(circuit_sizes) > 0 && circuit_sizes[0] == max {
		return true
	}
	return false
}

func solve_part_two(input string) string {
	junction_boxes := create_coordinate_list(input)
	connection_heap := calculate_connection_distances(junction_boxes)

	// Add connections until graph is complete
	adjacency_list := make(map[Coordinate][]Coordinate, 1000)
	var last_added Connection
	for !is_completed_circuit(adjacency_list, len(junction_boxes)) {
		connection := heap.Pop(&connection_heap).(Connection)
		last_added = connection
		adjacency_list[connection.source] = append(adjacency_list[connection.source], connection.destination)
		adjacency_list[connection.destination] = append(adjacency_list[connection.destination], connection.source)
	}

	return strconv.Itoa(int(last_added.source.x) * int(last_added.destination.x))
}
