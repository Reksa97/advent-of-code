import { readFileSync } from 'fs'

const input = readFileSync('./10/input.txt').toString().split(/\n/)

let cycle = 0
let x = 1
let sumOfSignalStrength = 0
const interestingCycles = [20, 60, 100, 140, 180, 220]
const pixels: string[] = []

const incrementCycle = (amount = 1) => {
  for (let i = 0; i < amount; i++) {
    const indexInRow = (cycle % 40)
    const pixel = (x - 1 <= indexInRow && indexInRow <= x + 1) ? '#' : '.'
    pixels.push(pixel)

    cycle++

    if (interestingCycles.includes(cycle)) {
      const signalStrength = cycle * x
      sumOfSignalStrength += signalStrength
    }
  }
}

input.forEach((line) => {
  if (line === 'noop') {
    incrementCycle()
    return
  }
  const [_, addX] = line.split(' ')
  incrementCycle(2)
  x += parseInt(addX)
})

console.log(`sum of interesting signal strengths is ${sumOfSignalStrength}\n`)

let row = ''
let i = 0
for (const pixel of pixels) {
  row += pixel
  i++
  if (i % 40 === 0) {
    console.log(row)
    row = ''
  }
}
