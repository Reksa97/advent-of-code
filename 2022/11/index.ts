import { readFileSync } from 'fs'

interface Monkey {
  index?: number
  items?: number[]
  itemsInspected?: number
  operation?: (old: number) => number
  testDivisibleBy?: number
  ifTrueThrowTo?: number
  ifFalseThrowTo?: number
}

const MONKEY_LINE = 'Monkey '
const STARTING_ITEMS_LINE = '  Starting items: '
const OPERATION_LINE = '  Operation: new = old '
const TEST_LINE = '  Test: divisible by '
const TEST_IF_TRUE_LINE = '    If true: throw to monkey '
const TEST_IF_FALSE_LINE = '    If false: throw to monkey '
const getInitialMonkeys = (input: string[]) => {
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
        const operator = operationParts[0]
        const variable = operationParts[1]
        switch (operator) {
          case '*':
            newValue = old * (variable === 'old' ? old : parseInt(variable))
            break
          case '+':
            newValue = old + (variable === 'old' ? old : parseInt(variable))
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
  return monkeys
}

const calculateMonkeyBusiness = (input: string[], rounds: number, part: 1 | 2) => {
  let monkeys = getInitialMonkeys(input)
  const productOfDivisors = monkeys.map(m => m.testDivisibleBy).reduce((acc, cur) => acc * cur, 1)
  for (let round = 1; round <= rounds; round++) {
    for (let i = 0; i < monkeys.length; i++) {
      while (monkeys[i].items.length > 0) {
        monkeys[i].itemsInspected++
        // Inspect item
        const item = monkeys[i].items.shift()
        let worryLevel = monkeys[i].operation(item)

        if (part === 1) {
          // Divide worry level
          worryLevel = ~~(worryLevel / 3)
        } else {
          // Calculate modulo to keep worry level manageable with a computer
          worryLevel = worryLevel % productOfDivisors
        }

        // Test worry level and throw to another monkey
        if (worryLevel % monkeys[i].testDivisibleBy === 0) {
          monkeys[monkeys[i].ifTrueThrowTo].items.push(worryLevel)
        } else {
          monkeys[monkeys[i].ifFalseThrowTo].items.push(worryLevel)
        }
      }
    }
  }
  return monkeys.sort(
    (a, b) => b.itemsInspected - a.itemsInspected
  ).slice(0, 2).reduce((acc, cur) => acc * cur.itemsInspected, 1)
}

const input = readFileSync('./11/input.txt').toString().split(/\n/)
console.log(`part 1: busiest two monkies had monkey business level of ${calculateMonkeyBusiness(input, 20, 1)}`)
console.log(`part 2: busiest two monkies had monkey business level of ${calculateMonkeyBusiness(input, 10_000, 2)}`)
