package main

import (
	"testing"
	"fmt"
)



func TestSearchHash(t *testing.T) {
	hash := make([]int,100000)
	InsertHash(hash,10000,100)
	InsertHash(hash,10000,101)
	InsertHash(hash,10000,102)
	InsertHash(hash,10000,103)
	InsertHash(hash,10000,105)
	fmt.Printf("hash search:    %d",SearchHash(hash,100000,105))
	fmt.Printf("hash search:    %d",SearchHash(hash,100000,10000))
}
