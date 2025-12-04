package main

import (
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AIState int

const (
	Seeking   = 0
	Gathering = 1
	Returning = 2
	Patrol    = 3
	Rest      = 4
)

type AppleThiefAI struct {
	Creature     *Creature
	State        AIState
	SightRange   float32
	TargetPos    rl.Vector2
	ScoreZone    *ScoreZone
	worldApples  *[]*Apple
	ChangedState bool
	Timer        float32
	TickCount    int
}

func NewAppleThiefAI(creature *Creature, scoreZone *ScoreZone, worldApples *[]*Apple) *AppleThiefAI {
	return &AppleThiefAI{
		Creature:     creature,
		State:        Seeking,
		SightRange:   1000,
		ScoreZone:    scoreZone,
		worldApples:  worldApples,
		ChangedState: false,
		Timer:        0,
		TickCount:    0,
	}
}

func (ai *AppleThiefAI) SetState(newState AIState) {
	ai.ChangedState = true
	ai.State = newState
}

func (ai *AppleThiefAI) Tick() {
	if ai.ChangedState {
		ai.Timer = 0
		ai.TickCount = 0
		ai.ChangedState = false
	}
	switch ai.State {
	case Seeking:
		ai.TickSeek()
	case Gathering:
		ai.TickGather()
	case Returning:
		ai.TickReturn()
	case Patrol:
		ai.TickPatrol()
	case Rest:
		ai.TickRest()
	}
	ai.Timer += rl.GetFrameTime()
	ai.TickCount += 1
}

func (ai *AppleThiefAI) FindNearestApple() (*Apple, bool) {
	var nearestApple *Apple = nil
	minDist := float32(ai.SightRange)

	for _, apple := range *ai.worldApples {
		if apple.Carried {
			continue
		}
		dist := rl.Vector2Distance(ai.Creature.Pos, apple.Pos)
		if dist > ai.SightRange {
			continue
		}
		if dist < minDist {
			minDist = dist
			nearestApple = apple
		}
	}
	return nearestApple, nearestApple != nil
}

func (ai *AppleThiefAI) TickSeek() {
	if len(ai.Creature.Apples) >= CREATURE_MAX_APPLES {
		ai.SetState(Returning)
		return
	}

	apple, found := ai.FindNearestApple()
	if found {
		ai.TargetPos = apple.Pos
		ai.SetState(Gathering)
	} else if !found && len(ai.Creature.Apples) == 0 {
		ai.SetState(Patrol)
	} else if len(ai.Creature.Apples) > 0 {
		ai.SetState(Returning)
	}
}

func (ai *AppleThiefAI) TickGather() {
	dist := rl.Vector2Distance(ai.Creature.Pos, ai.TargetPos)

	if dist < APPLE_SIZE+CREATURE_SIZE {
		ai.Creature.GatherApples(ai.worldApples)
		ai.SetState(Seeking)
		return
	}

	ai.Creature.MoveCreatureTowards(ai.TargetPos)
}

func (ai *AppleThiefAI) TickReturn() {
	if len(ai.Creature.Apples) == 0 {
		ai.SetState(Seeking)
		return
	}

	dist := rl.Vector2Distance(ai.Creature.Pos, ai.ScoreZone.Pos)

	if dist < SCORE_ZONE_SIZE {
		ai.Creature.DepositApple(ai.ScoreZone)
		if len(ai.Creature.Apples) == 0 {
			ai.SetState(Seeking)
		}
		return
	}

	ai.Creature.MoveCreatureTowards(ai.ScoreZone.Pos)
}

func (ai *AppleThiefAI) TickRest() {
	if ai.TickCount == 0 {
		fmt.Println("I'm resting :3")
	}

	if ai.Timer < 5 { //do nothing for 5 seconds
		return
	}
	ai.SetState(Patrol)
}

func (ai *AppleThiefAI) TickPatrol() {
	if ai.TickCount == 0 {
		x := float32(rand.IntN(100) + SCORE_ZONE_SIZE + 20)
		y := float32(rand.IntN(100) + SCORE_ZONE_SIZE + 20)
		if rand.IntN(2) == 1 {
			x *= -1
		}
		if rand.IntN(2) == 1 {
			y *= -1
		}
		ai.TargetPos = rl.NewVector2(x, y)
		fmt.Println("I'm patrolling")
	}
	dist := rl.Vector2Distance(ai.Creature.Pos, ai.TargetPos)
	if dist < 5 {
		ai.SetState(Rest)
	}
	ai.Creature.MoveCreatureTowards(ai.TargetPos)
}
