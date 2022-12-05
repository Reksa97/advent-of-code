import { readFileSync } from 'fs'

const input = readFileSync('./4/input.txt').toString().split(/\n/)
let count = 0
input.forEach((line) => {
  const [first, second] = line.split(',')
  const [firstStart, firstEnd] = first.split('-').map(c => parseInt(c))
  const [secondStart, secondEnd] = second.split('-').map(c => parseInt(c))

  if (
    (firstStart >= secondStart && firstEnd <= secondEnd)
    || (secondStart >= firstStart && secondEnd <= firstEnd)
  ) {
    count++
  }
})
console.log('part1', count)

count = 0
input.forEach((line) => {
  const [first, second] = line.split(',')
  const [firstStart, firstEnd] = first.split('-').map(c => parseInt(c))
  const [secondStart, secondEnd] = second.split('-').map(c => parseInt(c))

  if (
    (firstStart >= secondStart && firstStart <= secondEnd)
    || (secondStart >= firstStart && secondStart <= firstEnd)
    || (firstEnd <= secondEnd && firstEnd >= secondStart)
    || (secondEnd <= firstEnd && secondEnd >= firstStart)
  ) {
    count++
  }
})

console.log('part2', count)