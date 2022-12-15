import { readFileSync } from 'fs'
import { positionToString } from '../utils'

interface Position {
  x: number,
  y: number
}

const ARROW = ' -> '

const sandStartPosition: Position = {
  x: 500,
  y: 0
}

const rockPositions = new Set<string>()
const input = readFileSync('./14/input.txt').toString().split(/\n/)

let minX = sandStartPosition.x
let maxX = sandStartPosition.x
let minY = sandStartPosition.y
let maxY = sandStartPosition.y

input.forEach((line) => {
  const rockEdges: Position[] = line.split(ARROW).map(p => {
    const [x, y] = p.split(',').map(c => parseInt(c))
    if (y > maxY) maxY = y
    if (y < minY) minY = y
    if (x > maxX) maxX = x
    if (x < minX) minX = x
    return { x, y }
  })

  for (let i = 1; i < rockEdges.length; i++) {
    const from = rockEdges[i - 1]
    const to = rockEdges[i]
    rockPositions.add(positionToString(from))

    if (from.x === to.x) {
      const yStep = (to.y - from.y > 0) ? 1 : -1
      const x = from.x
      let y = from.y
      while (y !== to.y) {
        y += yStep
        rockPositions.add(positionToString({ x, y }))
      }
      continue
    }

    if (from.y === to.y) {
      const xStep = (to.x - from.x > 0) ? 1 : -1
      const y = from.y
      let x = from.x
      while (x !== to.x) {
        x += xStep
        rockPositions.add(positionToString({ x, y }))
      }
      continue
    }

    // One of two above if cases should be always true, otherwise input is malformed
  }
})
const floorY = maxY + 2
maxY++
maxX++
minY--
minX--

const simulateSand = (startPosition: Position, rocks: Set<string>, floorY: number, part: 1 | 2) => {
  const sand = new Set<string>()

  const isPositionFree = (pos: Position) => {
    const posString = positionToString(pos)
    const positionIsFree = !rocks.has(posString) && !sand.has(posString)
    return positionIsFree
  }

  if (part === 2) {
    sand.add(positionToString(startPosition))
    let leftX = startPosition.x
    let rightX = startPosition.x
    for (let y = startPosition.y + 1; y < floorY; y++) {
      leftX--
      rightX++
      for (let x = leftX; x <= rightX; x++) {
        const positionString = positionToString({ x, y })
        const positionIsFree = !rocks.has(positionString)
        const sandComingAbove = [{ x: x - 1, y: y - 1 }, { x: x, y: y - 1 }, { x: x + 1, y: y - 1 }].map(p => positionToString(p)).some(p => sand.has(p))
        if (positionIsFree && sandComingAbove) {
          sand.add(positionString)
        }
      }
    }
    minX = leftX
    maxX = rightX
    return { sand }
  }

  const lastSandPath = new Set<string>()

  sandLoop: while (true) {
    lastSandPath.clear()
    const newSand: Position = { x: startPosition.x, y: startPosition.y }
    let atRest = false
    while (!atRest) {
      lastSandPath.add(positionToString(newSand))
      if (isPositionFree({ x: newSand.x, y: newSand.y + 1 })) {
        newSand.y++
        if (newSand.y > maxY) {
          lastSandPath.add(positionToString(newSand))
          break sandLoop
        }
        continue
      }
      if (isPositionFree({ y: newSand.y + 1, x: newSand.x - 1 })) {
        newSand.y++
        newSand.x--
        continue
      }
      if (isPositionFree({ y: newSand.y + 1, x: newSand.x + 1 })) {
        newSand.y++
        newSand.x++
        continue
      }
      atRest = true
    }
    sand.add(positionToString(newSand))
  }
  return { sand, lastSandPath }
}

const printTiles = (part: 1 | 2, sand: Set<string>, lastSandPath?: Set<string>) => {
  for (let y = minY; y <= floorY; y++) {
    let row = ''
    for (let x = minX; x <= maxX; x++) {
      const positionString = positionToString({ x, y })
      if (positionToString(sandStartPosition) === positionString) {
        row += sand.has(positionString) ? '🟡' : '🔻'
        continue
      }
      if (sand.has(positionString)) {
        row += '🟨'
        continue
      }
      if (rockPositions.has(positionString)) {
        row += '🟫'
        continue
      }
      if (lastSandPath && lastSandPath.has(positionString)) {
        row += '🔻'
        continue
      }
      if (y === floorY && part === 2) {
        row += '⬛️'
        continue
      }
      row += '⬜️'
    }
    console.log(row)
  }
}

let { sand, lastSandPath } = simulateSand(sandStartPosition, rockPositions, floorY, 1)
printTiles(1, sand, lastSandPath)

console.log(`part 1: there are ${sand.size} units of settled sand`)

const secondResult = simulateSand(sandStartPosition, rockPositions, floorY, 2)
printTiles(2, secondResult.sand)

console.log(`part 2: there are ${secondResult.sand.size} units of settled sand`)
