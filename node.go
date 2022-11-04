package main

import (
	_"fmt"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}