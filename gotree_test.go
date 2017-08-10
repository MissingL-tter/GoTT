package gotree

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"text/tabwriter"
	"time"
)

var iters int
var nodes int

func init() {
	flag.IntVar(&iters, "iters", 10, "The number of iterations to average execution time over")
	flag.IntVar(&nodes, "nodes", 1000000, "The number of nodes in the tree")
}

// TestTreeStruct verifies that the tree is being structured correctly during build
//
// It explicitly checks each node and if a check fails it will print the error and fail.
func TestTreeStruct(t *testing.T) {

	vals := []float32{5, 3, 2, 4, 7, 6, 8}

	tree := Build(vals)

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
	if tree.Value != vals[0] && tree.Level == 0 {
		fmt.Fprintf(writer, "Root was incorrect \t Wanted: %v \t Got: %v \t\n", vals[0], tree.Value)
		t.Fail()
	}
	if tree.Left.Value != vals[1] && tree.Level == 1 {
		fmt.Fprintf(writer, "Left subtree incorrect \t Wanted: %v \t Got: %v \t\n", vals[1], tree.Value)
		t.Fail()
	}
	if tree.Left.Left.Value != vals[2] && tree.Level == 2 {
		fmt.Fprintf(writer, "Left subtree left leaf incorrect \t Wanted: %v \t Got: %v \t\n", vals[2], tree.Value)
		t.Fail()
	}
	if tree.Left.Right.Value != vals[3] && tree.Level == 2 {
		fmt.Fprintf(writer, "Left subtree right leaf incorrect \t Wanted: %v \t Got: %v \t\n", vals[3], tree.Value)
		t.Fail()
	}
	if tree.Right.Value != vals[4] && tree.Level == 1 {
		fmt.Fprintf(writer, "Right subtree incorrect \t Wanted: %v \t Got: %v \t\n", vals[4], tree.Value)
		t.Fail()
	}
	if tree.Right.Left.Value != vals[5] && tree.Level == 2 {
		fmt.Fprintf(writer, "Right subtree left leaf incorrect \t Wanted: %v \t Got: %v \t\n", vals[5], tree.Value)
		t.Fail()
	}
	if tree.Right.Right.Value != vals[6] && tree.Level == 2 {
		fmt.Fprintf(writer, "Right subtree right leaf incorrect \t Wanted: %v \t Got: %v \t\n", vals[6], tree.Value)
		t.Fail()
	}
	writer.Flush()
}

// TestGotree builds and traverses a tree of the same values and prints the results
//
// The number of nodes is defined with the flag "nodes" and the number of iterations is defined with the flag "iters".
// This test will not fail and simply provides controlled way to benchmark the tree without using Go's benchmark utility.
func TestGotree(t *testing.T) {

	vals := []float32{}
	for i := 0; i < nodes; i++ {
		vals = append(vals, float32(rand.Intn(nodes)))
	}

	var buildSum time.Duration
	var traverseSum time.Duration
	var traverseParallelSum time.Duration
	for i := 0; i < iters; i++ {
		start := time.Now()
		tree := Build(vals)
		buildSum += time.Since(start)

		start = time.Now()
		tree.InOrder()
		traverseSum += time.Since(start)

		start = time.Now()
		tree.InOrderParallel()
		traverseParallelSum += time.Since(start)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
	fmt.Fprintf(writer, "Args \t Iterations: \t %v \t\n", iters)
	fmt.Fprintf(writer, "\t Nodes: \t %v \t\n", nodes)
	fmt.Fprintf(writer, "Mean Times \t Build Tree: \t %v \t\n", buildSum/time.Duration(iters))
	fmt.Fprintf(writer, "\t Traverse Tree: \t %v \t\n", traverseSum/time.Duration(iters))
	fmt.Fprintf(writer, "\t Traverse Parallel: \t %v \t\n", traverseParallelSum/time.Duration(iters))
	writer.Flush()
}

// BenchmarkBuild provides a benchmark for the building of a tree
func BenchmarkBuild(b *testing.B) {
	vals := make([]float32, nodes)
	for i := 0; i < nodes; i++ {
		vals[i] = float32(rand.Intn(nodes))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Build(vals)
	}
}

// BenchmarkInOrder provides a benchmark for the in order traversal of a tree
func BenchmarkInOrder(b *testing.B) {
	vals := []float32{}
	for i := 0; i < nodes; i++ {
		vals = append(vals, float32(rand.Intn(nodes)))
	}

	tree := Build(vals)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree.InOrder()
	}
}

// BenchmarkInOrder provides a benchmark for the in order traversal of a tree making use of multiple threads
func BenchmarkInOrderParallel(b *testing.B) {
	vals := []float32{}
	for i := 0; i < nodes; i++ {
		vals = append(vals, float32(rand.Intn(nodes)))
	}

	tree := Build(vals)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree.InOrderParallel()
	}
}
