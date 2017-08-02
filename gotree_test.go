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
	flag.IntVar(&nodes, "nodes", 100000, "The number of nodes in the tree")
}

func TestGotree(t *testing.T) {
	vals := []float32{}
	for i := 0; i < nodes; i++ {
		vals = append(vals, float32(rand.Int()))
	}

	var buildSum time.Duration
	var traverseSum time.Duration
	for i := 0; i < iters; i++ {
		start := time.Now()
		tree := Build(vals)
		buildSum += time.Since(start)

		start = time.Now()
		InOrder(tree)
		traverseSum += time.Since(start)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
	fmt.Fprintf(writer, "Args \t Iterations: \t %v \t\n", iters)
	fmt.Fprintf(writer, "\t Nodes: \t %v \t\n", nodes)
	fmt.Fprintf(writer, "Mean Times \t Build Tree: \t %v \t\n", buildSum/time.Duration(iters))
	fmt.Fprintf(writer, "\t Traverse Tree: \t %v \t\n", traverseSum/time.Duration(iters))
	writer.Flush()
}

func BenchmarkTreeBuild(b *testing.B) {
	vals := make([]float32, nodes)
	for i := 0; i < nodes; i++ {
		vals[i] = float32(rand.Int())
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Build(vals)
	}
}

func BenchmarkTreeTraverse(b *testing.B) {
	vals := []float32{}
	for i := 0; i < nodes; i++ {
		vals = append(vals, float32(rand.Int()))
	}

	tree := Build(vals)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		InOrder(tree)
	}
}
