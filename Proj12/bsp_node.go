package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BSPNode struct {
	Left, Right *BSPNode
	Rect        Rectangle
	Room        Rectangle
	Color       rl.Color
}

type Rectangle struct {
	X, Y, Width, Height int32
}

func (n *BSPNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *BSPNode) FindRoom() Rectangle {
	if n.Room.Width > 0 {
		return n.Room
	}

	var lRoom, rRoom Rectangle
	if n.Left != nil {
		lRoom = n.Left.FindRoom()
	}
	if n.Right != nil {
		rRoom = n.Right.FindRoom()
	}

	if lRoom.Width > 0 && rRoom.Width > 0 {
		if rand.Intn(2) == 0 {
			return lRoom
		}
		return rRoom
	} else if lRoom.Width > 0 {
		return lRoom
	}
	return rRoom
}
