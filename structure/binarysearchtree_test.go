package main

import (
	"testing"
	"fmt"
)

var bst ItemBinarySearchTree

func fillTree(bst *ItemBinarySearchTree) {
	bst.Insert(8, 8)
	bst.Insert(4, 4)
	bst.Insert(10, 10)
	bst.Insert(2, 2)
	bst.Insert(6, 6)
	bst.Insert(1, 1)
	bst.Insert(3, 3)
	bst.Insert(5, 5)
	bst.Insert(7, 7)
	bst.Insert(9, 9)
}

func TestInsert(t *testing.T) {
	fillTree(&bst)
	bst.String()

	bst.Insert(11, 11)
	bst.String()
}


func TestSearch(t *testing.T) {
	if !bst.Search(1) || !bst.Search(8) || !bst.Search(11) {
		t.Errorf("search not working")
	}
}


func TestMin(t *testing.T) {
	fmt.Printf("min :%d", bst.Min())
	if  bst.Min() != 1 {
		t.Errorf("min should be 1")
	}
}

func TestMax(t *testing.T) {
	if bst.Max() != 11 {
		t.Errorf("max should be 11")
	}
}