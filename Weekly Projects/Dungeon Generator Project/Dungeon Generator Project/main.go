package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "Brian Statom - Dungeon Gen Assignment")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	dungeon := MakeDungeon()
	BuildBSP(1, &dungeon.Cells, rl.NewRectangle(0, 0, 1600, 900))
	GenerateHallways(&dungeon.Cells, 5, 1)
	GenerateHallways(&dungeon.Cells, 4, 1)
	GenerateHallways(&dungeon.Cells, 3, 1)
	GenerateHallways(&dungeon.Cells, 2, 1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		if rl.IsKeyPressed(rl.KeyR) {
			dungeon = MakeDungeon()
			BuildBSP(1, &dungeon.Cells, rl.NewRectangle(0, 0, 1600, 900))
			GenerateHallways(&dungeon.Cells, 5, 1)
			GenerateHallways(&dungeon.Cells, 4, 1)
			GenerateHallways(&dungeon.Cells, 3, 1)
			GenerateHallways(&dungeon.Cells, 2, 1)
		}

		DrawBSP(dungeon.Cells)

		rl.EndDrawing()
	}
}

func DrawBSP(cell Cell) {
	if cell.hallway != rl.NewRectangle(0, 0, 0, 0) {
		rl.DrawRectangle(int32(cell.hallway.X), int32(cell.hallway.Y), int32(cell.hallway.Width), int32(cell.hallway.Height), rl.Gray)
	}
	if cell.left == nil && cell.right == nil {
		rl.DrawRectangleLines(int32(cell.rec.X), int32(cell.rec.Y), int32(cell.rec.Width), int32(cell.rec.Height), rl.Green)
		rl.DrawRectangle(int32(cell.room.X), int32(cell.room.Y), int32(cell.room.Width), int32(cell.room.Height), rl.Gray)
		return
	}
	DrawBSP(*cell.left)
	DrawBSP(*cell.right)
}
