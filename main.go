package main

import (
	"fmt"
	"leetCode/alg"
)

func main() {
	alg.InOrderRecursion(alg.GetTreeNode())
	fmt.Println("##################")
	alg.InOrder(alg.GetTreeNode())
}
