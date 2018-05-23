package main

import (
	"math"
)

// adjacent tells us whether two points are adjacent to each other (have a distance of 1)
func adjacent(a, b Point) bool {
	return manhattanDistance(a, b) == 1
}

// calculate the manhattan distance between two points
func manhattanDistance(a, b Point) int {
	return absInt(a.X-b.X) + absInt(a.Y-b.Y)
}

// determines whether a point is within the board
func isInBounds(p Point, width, height int) bool {
	// Check if x out of bounds
	if p.X < 0 || p.X >= width {
		return false
	}

	// Check if y out of bounds
	if p.Y > 0 || (-1*height) <= p.Y {
		return false
	}

	return true
}

// helper function for calculating absolute values of int values
func absInt(v int) int {
	return int(math.Abs(float64(v)))
}

// checks whether the point would be a deadly collision
func isDeadlyCollision(p Point, sl *SnakeList, width, height int) bool {
	// check if this point goes out of bounds
	if !isInBounds(p, width, height) {
		// walls are deadly
		return true
	}

	// check if this point overlaps a snake body
	for _, snek := range *sl {
		for _, s := range snek.Body {
			isSame(p, s)
		}
	}

	return false
}

// returns whether two points have the same coordinates
func isSame(a, b Point) bool {
	if a.X == b.X && a.Y == b.Y {
		// collides with a snake body
		return true
	}
	return false
}

// gets the moves that a snake can make
func getAvailableMoves(you Snake, all SnakeList, width, height int) []Point {
	snakeHead := you.Body[0]
	available := getSurroundingSquares(snakeHead)

	// remove any points that are the same as where a snake body part already is
	for _, snek := range all {
		for _, sb := range snek.Body {
			available = removePoint(available, sb)
		}
	}

	// remove any points that are outside of the board
	available = removeOutOfBounds(available, width, height)

	// what's left should be only points that are safe
	return available
}

// removes any points from a list that have the same coordinates
func removePoint(from []Point, p Point) []Point {
	var remainingPoints []Point
	for i, _ := range from {
		// Skip any points with the same x,y
		if isSame(from[i], p) {
			continue
		}
		remainingPoints = append(remainingPoints, from[i])
	}

	return remainingPoints
}

// remove any points that are out of the board
func removeOutOfBounds(from []Point, width, height int) []Point {
	var inBounds []Point
	for _, p := range from {
		if isInBounds(p, width, height) {
			inBounds = append(inBounds, p)
		}
	}

	return inBounds
}

// gets surrounding squares for a point
func getSurroundingSquares(p Point) []Point {
	return []Point{
		Point{X: p.X + 1, Y: p.Y},
		Point{X: p.X, Y: p.Y + 1},
		Point{X: p.X + 1, Y: p.Y - 1},
		Point{X: p.X - 1, Y: p.Y},
	}
}

// finds the closest food to a point
func findClosestFood(p Point, food PointList) *Point {
	// no food
	if len(food) == 0 {
		return nil
	}

	// use manhattanDistance to find the closest food
	var closestFood Point
	closestDistance := 99999999999
	for _, f := range food {
		// calculate how close this food is
		dist := manhattanDistance(p, f)
		// check if it's closer
		if dist < closestDistance {
			closestDistance = dist
			closestFood = f
		}
	}

	return &closestFood
}
