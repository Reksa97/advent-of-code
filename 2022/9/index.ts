import { readFileSync } from 'fs'

const input = readFileSync('./9/input.txt').toString().split(/\n/)

interface Position {
  x: number,
  y: number
}

const calculateTailPositions = (knotAmount: number): number => {
  const tailPositions = new Set<string>()
  const knotPositions: Position[] = []
  for (let i = 0; i < knotAmount; i++) knotPositions.push({ x: 0, y: 0 })
  input.forEach((line) => {
    const [direction, amount] = line.split(' ')
    const amountInt = parseInt(amount)
    for (let i = 0; i < amountInt; i++) {
      switch (direction) {
        case "U":
          knotPositions[0].y++
          break
        case "D":
          knotPositions[0].y--
          break
        case "L":
          knotPositions[0].x--
          break
        case "R":
          knotPositions[0].x++
          break
      }
      for (let i = 1; i < knotAmount; i++) {
        const xDiff = knotPositions[i - 1].x - knotPositions[i].x
        const yDiff = knotPositions[i - 1].y - knotPositions[i].y
        if (Math.abs(xDiff) > 1 || Math.abs(yDiff) > 1) {
          const xTranslation = xDiff > 0 ? 1 : -1;
          const yTranslation = yDiff > 0 ? 1 : -1;
          if (xDiff !== 0) {
            knotPositions[i].x += xTranslation
          }
          if (yDiff !== 0) {
            knotPositions[i].y += yTranslation
          }
        }
        if (i === knotAmount - 1) tailPositions.add(`${knotPositions[i].x}:${knotPositions[i].y}`)
      }
    }
  })
  return tailPositions.size
}

console.log(`part 1: tail visited ${calculateTailPositions(2)} places`)
console.log(`part 2: tail visited ${calculateTailPositions(10)} places`)
