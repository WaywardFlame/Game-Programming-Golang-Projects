package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dungeon struct {
	Cells Cell
}

type Cell struct {
	rec       rl.Rectangle
	left      *Cell
	right     *Cell
	sister    *Cell
	room      rl.Rectangle
	hallway   rl.Rectangle
	connected bool
}

func MakeDungeon() Dungeon {
	d := Dungeon{}
	d.Cells = Cell{rec: rl.NewRectangle(0, 0, 0, 0), left: nil, right: nil, sister: nil, room: rl.NewRectangle(0, 0, 0, 0), hallway: rl.NewRectangle(0, 0, 0, 0), connected: false}
	return d
}

func BuildBSP(iteration int, cell *Cell, r rl.Rectangle) {
	var direction int = rand.IntN(2)
	partition := float32(rand.IntN(21)+40) / 100 // between 0.40 and 0.60

	var rec1 rl.Rectangle
	var rec2 rl.Rectangle
	var position float32
	if direction == 0 { // split vertically
		position = r.Width * partition
		rec1 = rl.NewRectangle(r.X, r.Y, position, r.Height)
		rec2 = rl.NewRectangle(r.X+position, r.Y, r.Width-position, r.Height)
	} else if direction == 1 { // split horizontally
		position = r.Height * partition
		rec1 = rl.NewRectangle(r.X, r.Y, r.Width, position)
		rec2 = rl.NewRectangle(r.X, r.Y+position, r.Width, r.Height-position)
	}

	c1 := Cell{rec: rec1, left: nil, right: nil, sister: nil, room: rl.NewRectangle(0, 0, 0, 0), connected: false}
	c2 := Cell{rec: rec2, left: nil, right: nil, sister: nil, room: rl.NewRectangle(0, 0, 0, 0), connected: false}
	if iteration == 4 {
		c1.room = GenerateRoom(rec1)
		c2.room = GenerateRoom(rec2)
	}
	c1.sister = &c2
	c2.sister = &c1

	cell.left = &c1
	cell.right = &c2

	if iteration != 4 {
		BuildBSP(iteration+1, &c1, rec1)
		BuildBSP(iteration+1, &c2, rec2)
	}
}

func GenerateRoom(cell rl.Rectangle) rl.Rectangle {
	maxWidth := cell.Width * 0.8
	maxHeight := cell.Height * 0.8
	minWidth := cell.Width * 0.4
	minHeight := cell.Height * 0.4

	newWidth := float32(rand.IntN(int(maxWidth-minWidth))) + minWidth
	newHeight := float32(rand.IntN(int(maxHeight-minHeight))) + minHeight

	minX := (cell.Width * 0.1)
	minY := (cell.Height * 0.1)
	seedX := (maxWidth - newWidth)
	seedY := (maxHeight - newHeight)

	newX := float32(rand.IntN(int(seedX))) + minX + cell.X
	newY := float32(rand.IntN(int(seedY))) + minY + cell.Y

	return rl.NewRectangle(newX, newY, newWidth, newHeight)
}

// For everything below, I'm sure there is a much better way to do,
// but I don't know what that way would be.

func GenerateHallways(cell *Cell, level int, depth int) {
	if cell.connected {
		return
	}

	if depth == 5 { // if bottom most level
		room1 := cell.room
		room2 := cell.sister.room

		newHall := CompareRooms(room1, room2)
		cell.hallway = newHall
		cell.sister.hallway = newHall

		cell.connected = true
		cell.sister.connected = true
		return
	} else if depth == level { // if other level, assumes lower level hallways already created
		leftsideRooms := make([]rl.Rectangle, 0, 20)
		rightsideRooms := make([]rl.Rectangle, 0, 20)
		GetRooms(*cell, &leftsideRooms)
		GetRooms(*cell.sister, &rightsideRooms)

		room1, room2 := CompareDistance(leftsideRooms, rightsideRooms)

		newHall := CompareRooms(room1, room2)
		cell.hallway = newHall
		cell.sister.hallway = newHall

		cell.connected = true
		cell.sister.connected = true
		return
	}
	GenerateHallways(cell.left, level, depth+1)
	GenerateHallways(cell.right, level, depth+1)
}

func GetRooms(cell Cell, leftsideRooms *[]rl.Rectangle) {
	if cell.left == nil && cell.right == nil {
		(*leftsideRooms) = append((*leftsideRooms), cell.room)
		return
	}
	GetRooms(*cell.left, leftsideRooms)
	GetRooms(*cell.right, leftsideRooms)
}

func CompareDistance(leftsideRooms []rl.Rectangle, rightsideRoom []rl.Rectangle) (rl.Rectangle, rl.Rectangle) {
	var room1 rl.Rectangle
	var room2 rl.Rectangle
	var minDistance float32
	for i := 0; i < len(leftsideRooms); i++ {
		for k := 0; k < len(rightsideRoom); k++ {
			if i == 0 && k == 0 {
				room1 = leftsideRooms[0]
				room2 = rightsideRoom[0]
				minDistance = GetDistance(room1, room2)
				continue
			}
			temp := GetDistance(leftsideRooms[i], rightsideRoom[k])
			if temp < minDistance {
				minDistance = temp
				room1 = leftsideRooms[i]
				room2 = rightsideRoom[k]
			}
		}
	}
	return room1, room2
}

func GetDistance(room1 rl.Rectangle, room2 rl.Rectangle) float32 {
	left := room2.X+room2.Width < room1.X
	right := room1.X+room1.Width < room2.X
	bottom := room2.Y+room2.Height < room1.Y
	top := room1.Y+room1.Height < room2.Y
	if top && left {
		return rl.Vector2Distance(rl.NewVector2(room1.X, room1.Y+room1.Height), rl.NewVector2(room2.X+room2.Width, room2.Y))
	} else if left && bottom {
		return rl.Vector2Distance(rl.NewVector2(room1.X, room1.Y), rl.NewVector2(room2.X+room2.Width, room2.Y+room2.Height))
	} else if bottom && right {
		return rl.Vector2Distance(rl.NewVector2(room1.X+room1.Width, room1.Y), rl.NewVector2(room2.X, room2.Y+room2.Height))
	} else if right && top {
		return rl.Vector2Distance(rl.NewVector2(room1.X+room1.Width, room1.Y+room1.Height), rl.NewVector2(room2.X, room2.Y))
	} else if left {
		return room1.X - (room2.X + room2.Width)
	} else if right {
		return room2.X - (room1.X + room1.Width)
	} else if bottom {
		return room1.Y - (room2.Y + room2.Height)
	} else if top {
		return room2.Y - (room1.Y + room1.Height)
	}
	return 0
}

func CompareRooms(room1 rl.Rectangle, room2 rl.Rectangle) rl.Rectangle {
	if room1.Y+room1.Height < room2.Y { // FIXME: does not account for diagonals
		var newHeight float32
		var newY float32
		if room1.Y+room1.Height < room2.Y {
			newHeight = room2.Y - (room1.Y + room1.Height)
			newHeight += 10
			newY = room1.Y + room1.Height
			newY -= 5
		} else {
			newHeight = room1.Y - (room2.Y + room2.Height)
			newHeight += 10
			newY = room2.Y + room2.Height
			newY -= 5
		}
		var newWidth float32 = 15
		var newX float32
		noPos := true
		if room1.X <= room2.X {
			var i float32
			for i = room1.X; i+15 <= room1.Width+room1.X && noPos; i++ {
				if i >= room2.X && i+15 <= room2.Width+room2.X {
					noPos = false
				}
			}
			newX = i
		} else {
			var i float32
			for i = room2.X; i+15 <= room2.Width+room2.X && noPos; i++ {
				if i >= room1.X && i+15 <= room1.Width+room1.X {
					noPos = false
				}
			}
			newX = i
		}
		newHall := rl.NewRectangle(newX, newY, newWidth, newHeight)
		return newHall
	} else if room1.X+room1.Width < room2.X { // FIXME: does not account for diagonals
		var newWidth float32
		var newX float32
		if room1.X+room1.Width < room2.X {
			newWidth = room2.X - (room1.X + room1.Width)
			newWidth += 10
			newX = room1.X + room1.Width
			newX -= 5
		} else {
			newWidth = room1.X - (room2.X + room2.Width)
			newWidth += 10
			newX = room2.X + room2.Width
			newX -= 5
		}
		var newHeight float32 = 15
		var newY float32
		noPos := true
		if room1.Y <= room2.Y {
			var i float32
			for i = room1.Y; i+15 <= room1.Height+room1.Y && noPos; i++ {
				if i >= room2.Y && i+15 <= room2.Height+room2.Y {
					noPos = false
				}
			}
			newY = i
		} else {
			var i float32
			for i = room2.Y; i+15 <= room2.Height+room2.Y && noPos; i++ {
				if i >= room1.Y && i+15 <= room1.Height+room1.Y {
					noPos = false
				}
			}
			newY = i
		}
		newHall := rl.NewRectangle(newX, newY, newWidth, newHeight)
		return newHall
	}
	return rl.NewRectangle(0, 0, 0, 0)
}
