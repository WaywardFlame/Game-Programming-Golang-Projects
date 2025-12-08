package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Block struct {
	position rl.Vector2
	visible  bool
} // 62 x 62 size when drawing

func newBlock(x float32, y float32) Block {
	return Block{position: rl.NewVector2(x, y), visible: true}
}

func createBlocks() [3][10]Block {
	var blocks [3][10]Block
	var size float32 = 65
	var y float32
	for i := 0; i < 3; i++ {
		if i == 0 {
			y = 75
		} else if i == 1 {
			y = 140
		} else if i == 2 {
			y = 205
		}
		for k := 0; k < 10; k++ {
			blocks[i][k] = newBlock(75+(float32(k)*size), y)
		}
	}
	return blocks
}
