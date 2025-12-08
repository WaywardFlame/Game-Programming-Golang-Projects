package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Planet struct {
	PhysicsBody
	health int
}

// creates the planet
func createPlanet(planetTexture rl.Texture2D) Planet {
	pb := newPhysicsBody(rl.NewVector2(640, 360), rl.NewVector2(0, 0), 16, planetTexture, true)
	return Planet{health: 100, PhysicsBody: pb}
}
