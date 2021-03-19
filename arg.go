package main

import (
	"flag"
	"fmt"
)

func main() {
	// CLI
	textPtr := flag.String("text", "", "Text to Parse")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric")
	flag.parse()

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)

}
