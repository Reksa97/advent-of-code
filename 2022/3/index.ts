import { readFileSync } from 'fs'

const priorities = {
  'a': 1,
  'b': 2,
  'c': 3,
  'd': 4,
  'e': 5,
  'f': 6,
  'g': 7,
  'h': 8,
  'i': 9,
  'j': 10,
  'k': 11,
  'l': 12,
  'm': 13,
  'n': 14,
  'o': 15,
  'p': 16,
  'q': 17,
  'r': 18,
  's': 19,
  't': 20,
  'u': 21,
  'v': 22,
  'w': 23,
  'x': 24,
  'y': 25,
  'z': 26,
  'A': 27,
  'B': 28,
  'C': 29,
  'D': 30,
  'E': 31,
  'F': 32,
  'G': 33,
  'H': 34,
  'I': 35,
  'J': 36,
  'K': 37,
  'L': 38,
  'M': 39,
  'N': 40,
  'O': 41,
  'P': 42,
  'Q': 43,
  'R': 44,
  'S': 45,
  'T': 46,
  'U': 47,
  'V': 48,
  'W': 49,
  'X': 50,
  'Y': 51,
  'Z': 52,
}

let sum = 0
const input = readFileSync('./3/input.txt').toString().split(/\n/)
input.forEach((line) => {
  const midPoint = line.length / 2
  const map = {}
  const first = line.slice(0, midPoint)
  const second = line.slice(midPoint, line.length)

  for (const c of first) {
    map[c] = true
  }
  for (const c of second) {
    if (map[c]) {
      sum += priorities[c]
      break
    }
  }
})

console.log('part1', sum)

sum = 0
let i = 0
let groupItems = {}
let groupCount = 0
let lineCount = 0
input.forEach((line) => {

  const elfItems = {}

  for (const c of line) {
    elfItems[c] = true
  }

  Object.keys(elfItems).forEach(c => {
    groupItems[c] = (groupItems[c] ?? 0) + 1
  })

  i++
  lineCount++
  if (i > 2) {
    groupCount++
    Object.keys(groupItems).forEach(k => {
      if (groupItems[k] === 3) {
        sum += priorities[k]
      }
    })
    i = 0
    groupItems = {}
  }
})


console.log('part2', sum)