import { readFileSync } from 'fs'

let myPoints = 0
const input = readFileSync('./2/input.txt').toString().split(/\n/)
input.forEach((line) => {
  const [they, me] = line.split(' ')
  // A = rock, B = paper, C = scissors
  // X = I lose, Y = draw, Z = I win

  if (me === 'X') myPoints += 1
  if (me === 'Y') myPoints += 2
  if (me === 'Z') myPoints += 3

  // If they win
  if ((they === 'A' && me === 'Z') || (they === 'B' && me === 'X') || (they === 'C' && me === 'Y')) {
    // No points for me
  } // If I win
  else if ((they === 'A' && me === 'Y') || (they === 'B' && me === 'Z') || (they === 'C' && me === 'X')) {
    myPoints += 6
  } // If draw
  else {
    myPoints += 3
  }
})
console.log('part 1', myPoints)

myPoints = 0
input.forEach((line) => {
  const [they, me] = line.split(' ')
  // A = rock, B = paper, C = scissors
  // X = I lose, Y = draw, Z = I win

  if (me === 'X') {
    if (they === 'A') myPoints += 3 // I choose scissors
    if (they === 'B') myPoints += 1 // I choose rock
    if (they === 'C') myPoints += 2 // I choose paper
  } else if (me === 'Y') {
    myPoints += 3
    if (they === 'A') myPoints += 1 // I choose rock
    if (they === 'B') myPoints += 2 // I choose paper
    if (they === 'C') myPoints += 3 // I choose scissors
  } else if (me === 'Z') {
    myPoints += 6
    if (they === 'A') myPoints += 2 // I choose paper
    if (they === 'B') myPoints += 3 // I choose scissors
    if (they === 'C') myPoints += 1 // I choose rock
  }
})

console.log('part 2', myPoints)