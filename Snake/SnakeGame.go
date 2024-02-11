package snake

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Coords struct {
	x int
	y int
}

type Snake struct {
	head Coords
	body []Coords
}

const LEFT int = 0
const RIGHT int = 1
const UP int = 2
const DOWN int = 3

const GRID_LENGTH = 10

type SnakeGame struct {
	Snake Snake
	Food  Coords
}

func (snake Snake) contains(coords Coords) bool {
	for _, bodyCoords := range snake.body {
		if bodyCoords.x == coords.x && bodyCoords.y == coords.y {
			return true
		}
	}
	return snake.head.x == coords.x && snake.head.y == coords.y
}

// Return true if snake's head is inside its body or its head is out of bounds
func (snake Snake) ShouldDie() bool {
	for _, bodyCoords := range snake.body {
		if bodyCoords.x == snake.head.x && bodyCoords.y == snake.head.y {
			fmt.Println("Head in body", snake)
			return true
		}
	}
	return snake.head.outOfBounds()
}

func (coords Coords) outOfBounds() bool {
	return !(0 <= coords.x && coords.x < GRID_LENGTH && coords.y < GRID_LENGTH && 0 <= coords.y)
}

func (coords Coords) GetNeighbour(direction int) Coords {
	switch direction {
	case LEFT:
		return Coords{x: coords.x - 1, y: coords.y}
	case RIGHT:
		return Coords{x: coords.x + 1, y: coords.y}
	case UP:
		return Coords{x: coords.x, y: coords.y - 1}
	case DOWN:
		return Coords{x: coords.x, y: coords.y + 1}
	default:
		panic("Invalid direction")
	}
}

func (gameState *SnakeGame) getVision() []float64 {
	var output [10]float64
	output[0] = math.Min(math.Abs(float64(gameState.Snake.head.x-gameState.Food.x)), 8.0) / 8.0
	output[1] = math.Min(math.Abs(float64(gameState.Snake.head.y-gameState.Food.y)), 8.0) / 8.0

	directionsAsCoordChanges := []Coords{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
	for idx, coordChanges := range directionsAsCoordChanges {
		distance := gameState.getDistanceInDirection(coordChanges.x, coordChanges.y)
		normalisedDistance := math.Min(8.0, float64(distance)) / 8.0
		output[2+idx] = normalisedDistance
	}
	return output[:]
}

func (gameState SnakeGame) getDistanceInDirection(xModifier int, yModifier int) int {
	//Increment at least once else the snake is dead.
	canaryCoords := Coords{gameState.Snake.head.x + xModifier, gameState.Snake.head.y + yModifier}
	distance := 0

	// Until going in that distance causes death by out of bounds or snake collision
	for !gameState.Snake.contains(canaryCoords) && !canaryCoords.outOfBounds() {
		canaryCoords = Coords{canaryCoords.x + xModifier, canaryCoords.y + yModifier}
		distance++
	}
	return distance
}

func (gameState *SnakeGame) SpawnFood() {
	rand.Seed(time.Now().Unix())
	gameState.Food.x = rand.Intn(GRID_LENGTH)
	gameState.Food.y = rand.Intn(GRID_LENGTH)

	if gameState.Snake.contains(gameState.Food) {
		gameState.SpawnFood()
	}
}

func (gameState *SnakeGame) Move(direction int) {
	oldPos := gameState.Snake.head
	gameState.Snake.head = gameState.Snake.head.GetNeighbour(direction)

	foodCollision := gameState.Snake.head.x == gameState.Food.x && gameState.Snake.head.y == gameState.Food.y

	for idx, _ := range gameState.Snake.body {
		tmp := gameState.Snake.body[idx]
		gameState.Snake.body[idx] = oldPos
		oldPos = tmp
	}
	if foodCollision {
		gameState.Snake.body = append(gameState.Snake.body, oldPos)
		gameState.SpawnFood()
	}
}

func NewGame() *SnakeGame {
	rand.Seed(time.Now().UnixNano())
	snakeHead := Coords{x: GRID_LENGTH / 2, y: GRID_LENGTH / 2}

	newGame := SnakeGame{
		Snake: Snake{
			head: snakeHead,
			body: []Coords{snakeHead.GetNeighbour(rand.Intn(4))},
		},
	}
	newGame.SpawnFood()
	return &newGame
}
