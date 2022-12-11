import { readFileSync } from 'fs'

const input = readFileSync('./11/input.txt').toString().split(/\n/)

const MONKEY_LINE = 'Monkey '
const STARTING_ITEMS_LINE = '  Starting items: '
const OPERATION_LINE = '  Operation: new = old '
const TEST_LINE = '  Test: divisible by '
const TEST_IF_TRUE_LINE = '    If true: throw to monkey '
const TEST_IF_FALSE_LINE = '    If false: throw to monkey '

interface Monkey {
  index?: number
  items?: number[]
  itemsInspected?: number
  operation?: (old: number) => number
  testDivisibleBy?: number
  ifTrueThrowTo?: number
  ifFalseThrowTo?: number
}

const monkeys: Monkey[] = []
let monkey: Monkey = {}
input.forEach((line) => {
  if (line.startsWith(MONKEY_LINE)) {
    monkey.index = monkeys.length
    monkey.itemsInspected = 0
    return
  }

  if (line.startsWith(STARTING_ITEMS_LINE)) {
    monkey.items = line.substring(STARTING_ITEMS_LINE.length).split(', ').map(i => parseInt(i))
    return
  }

  if (line.startsWith(OPERATION_LINE)) {
    const operationParts = line.substring(OPERATION_LINE.length).split(' ')
    monkey.operation = (old: number) => {
      let newValue = old
      switch (operationParts[0]) {
        case '*':
          newValue = old * (operationParts[1] === 'old' ? old : parseInt(operationParts[1]))
          break
        case '+':
          newValue = old + (operationParts[1] === 'old' ? old : parseInt(operationParts[1]))
          break
      }
      return newValue
    }
    return
  }

  if (line.startsWith(TEST_LINE)) {
    monkey.testDivisibleBy = parseInt(line.substring(TEST_LINE.length))
    return
  }

  if (line.startsWith(TEST_IF_TRUE_LINE)) {
    monkey.ifTrueThrowTo = parseInt(line.substring(TEST_IF_TRUE_LINE.length))
    return
  }

  if (line.startsWith(TEST_IF_FALSE_LINE)) {
    monkey.ifFalseThrowTo = parseInt(line.substring(TEST_IF_FALSE_LINE.length))
    return
  }

  if (line === '') {
    monkeys.push({ ...monkey })
    monkey = {}
  }
})
monkeys.push({ ...monkey })
monkey = {}

for (let round = 1; round <= 20; round++) {
  for (let i = 0; i < monkeys.length; i++) {
    while (monkeys[i].items.length > 0) {
      monkeys[i].itemsInspected++
      // Inspect item
      const item = monkeys[i].items.shift()
      let worryLevel = monkeys[i].operation(item)
      // Divide worry level
      worryLevel = ~~(worryLevel / 3)

      // Test worry level and throw to another monkey
      if (worryLevel % monkeys[i].testDivisibleBy === 0) {
        monkeys[monkeys[i].ifTrueThrowTo].items.push(worryLevel)
      } else {
        monkeys[monkeys[i].ifFalseThrowTo].items.push(worryLevel)
      }
    }
  }
}

const topTwoMonkeyBusinessLevel = monkeys.sort((a, b) => b.itemsInspected - a.itemsInspected).slice(0, 2).reduce((acc, cur) => acc * cur.itemsInspected, 1)

console.log(`busiest two monkies had monkey business level of ${topTwoMonkeyBusinessLevel}`)

