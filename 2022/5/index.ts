import { readFileSync } from 'fs'

const [startPositions, input] = readFileSync('./5/input.txt').toString().split(/\n\n/)

const getStackStartPositions = () => {
  let stacks: string[][] = []
  for (let i = 0; i <= 9; i++) stacks.push([])

  startPositions.split(/\n/).slice(0,-1).forEach(line => {
    const stackItems = line.match(/.{1,4}/g)
    stackItems.forEach((item, i) => {
      const indexOfStartingBracket = item.indexOf('[')
      const char = (indexOfStartingBracket >= 0) ? item[indexOfStartingBracket+1] : undefined
      if (char !== undefined) {
        stacks[i+1].unshift(char)
      }
    })
  })
  return stacks
}

let stacks = getStackStartPositions()

input.split(/\n/).forEach((line) => {
  const [amount, from, to] = line.split(' ').map(s => parseInt(s)).filter(n => !isNaN(n))
  for (let i = 0; i < amount; i++) {
    stacks[to].push(stacks[from].pop())
  }
})

let answer = ''
stacks.slice(1).forEach(s => answer += s.at(-1))
console.log('part 1', answer)

stacks = getStackStartPositions()

input.split(/\n/).forEach((line) => {
  const [amount, from, to] = line.split(' ').map(s => parseInt(s)).filter(n => !isNaN(n))
  const cratesToMove = stacks[from].splice(-amount)
  stacks[to].push(...cratesToMove)
})

answer = ''
stacks.slice(1).forEach(s => answer += s.at(-1))
console.log('part 2', answer)