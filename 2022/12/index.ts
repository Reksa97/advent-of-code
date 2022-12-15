import { readFileSync } from 'fs'
import { Position, positionToString } from '../utils'

interface VisitedPositions {
  [position: string]: boolean
}

const ELEVATIONS = 'abcdefghijklmnopqrstuvwxyz'

const input = readFileSync('./12/input.txt').toString().split(/\n/)
const grid = []
const endPosition: Position = { x: -1, y: -1 }
input.forEach((line, y) => {
  const ePosition = line.indexOf('E')
  if (ePosition >= 0) {
    endPosition.x = ePosition
    endPosition.y = y
  }
  grid.push(line)
})

const gridWidth = grid[0].length
const gridHeight = grid.length

const getStartingPositions = (part: 1 | 2): Position[] => {
  const startPositions: Position[] = []
  input.forEach((line, y) => {
    const sPosition = line.indexOf('S')
    if (sPosition >= 0) {
      startPositions.push({ x: sPosition, y })
    }
    if (part === 2) {
      for (let x = 0; x < line.length; x++) {
        if (line.charAt(x) === 'a') startPositions.push({ x, y })
      }
    }
  })

  return startPositions
}

const getElevationAt = (position: Position): string => {
  const elevation = grid[position.y].charAt(position.x)
  if (elevation === 'S') return 'a'
  if (elevation === 'E') return 'z'
  return elevation
}

const canMove = (from: Position, to: Position): boolean => {
  if (to.x < 0 || to.x >= gridWidth) return false
  if (to.y < 0 || to.y >= gridHeight) return false

  const fromElevation = getElevationAt(from)
  const toElevation = getElevationAt(to)

  return ELEVATIONS.indexOf(toElevation) <= ELEVATIONS.indexOf(fromElevation) + 1
}

const getAdjacent = (pos: Position) => {
  const adjacent: Position[] = []
  // Up
  const positionUp: Position = {
    x: pos.x,
    y: pos.y - 1
  }
  if (canMove(pos, positionUp)) adjacent.push(positionUp)

  // Down
  const positionDown: Position = {
    x: pos.x,
    y: pos.y + 1
  }
  if (canMove(pos, positionDown)) adjacent.push(positionDown)

  // Right
  const positionRight: Position = {
    x: pos.x + 1,
    y: pos.y
  }
  if (canMove(pos, positionRight)) adjacent.push(positionRight)

  // Left
  const positionLeft: Position = {
    x: pos.x - 1,
    y: pos.y
  }
  if (canMove(pos, positionLeft)) adjacent.push(positionLeft)

  return adjacent
}

const BFS = (startPositions: Position[], goal: Position): number => {
  let shortestPath = Infinity
  for (const position of startPositions) {
    const visited: VisitedPositions = {}
    const queue = []
    visited[positionToString(position)] = true
    queue.push(position)
    while (queue.length > 0) {
      const pos = queue.shift()

      if (positionToString(pos) === positionToString(goal)) {
        let parent = pos
        let count = 0
        while (parent = parent.parent) {
          count++
        }
        if (count < shortestPath) {
          shortestPath = count
        }
        break
      }
      getAdjacent(pos).forEach((adjacent, i) => {
        if (!visited[positionToString(adjacent)]) {
          adjacent.parent = pos
          visited[positionToString(adjacent)] = true;
          queue.push(adjacent);
        }
      });
    }
  }
  return shortestPath
}

let startPositions = getStartingPositions(1)
console.log(`part 1: shortest route has ${BFS(startPositions, endPosition)} positions`);

startPositions = getStartingPositions(2)
console.log(`part 2: shortest route has ${BFS(startPositions, endPosition)} positions`)
