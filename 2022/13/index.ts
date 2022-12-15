import { readFileSync } from 'fs'

const compare = (leftPacket: any[], rightPacket: any[]) => {
  const maxLength = Math.max(leftPacket.length, rightPacket.length)
  const zipped = Array.from({ length: maxLength }).map((_, i) => [leftPacket[i], rightPacket[i]])
  for (let [left, right] of zipped) {
    if (right === undefined) return 1 // Right is greater
    if (left === undefined) return -1 // Left is greater

    if (Number.isInteger(left) && Number.isInteger(right)) {
      if (left > right) return 1 // Right is greater
      if (left < right) return -1 // Left is greater
      continue
    }

    const result = compare(
      Array.isArray(left) ? left :
        [left], Array.isArray(right) ? right : [right]
    )

    if (result !== 0) {
      return result;
    }
  }
  return 0
}

const input = readFileSync('./13/input.txt').toString().split(/\n\n/)

let sum = 0
const allPackets = [[[2]], [[6]]]
input.forEach((line, i) => {
  const [leftPacket, rightPacket] = line.split(/\n/).map(r => JSON.parse(r))
  allPackets.push(leftPacket, rightPacket)
  sum += compare(leftPacket, rightPacket) < 0 ? i + 1 : 0
})

console.log(`part 1: sum of indicises in correct order is ${sum}`)

const separatorPacketsAsStrings = ['[[2]]', '[[6]]']
const decoderKey = allPackets.sort(compare).reduce((acc, curr, i) => {
  const isSeparator = separatorPacketsAsStrings.includes(JSON.stringify(curr))
  return acc * (isSeparator ? i + 1 : 1)
}, 1)

console.log(`part 2: decoder key is ${decoderKey}`)
