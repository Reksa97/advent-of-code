import { readFileSync } from 'fs'

let max = 0
let maxThree = [0, 0, 0]
let calories = 0
const input = readFileSync('./1/input.txt')
input.toString().split(/\n/).forEach((line) => {
  if (line.length === 0) {
    if (calories > max) max = calories
    if (maxThree[0] < calories) {
      maxThree[0] = calories
      maxThree.sort()
    }
    calories = 0
    return
  }

  calories += parseInt(line)
});

console.log('part 1', max)
console.log(
  'part 2',
  maxThree,
  maxThree.reduce((accumulator, currentValue) => currentValue+accumulator, 0)
)