import { readFileSync } from 'fs'

const input = readFileSync('./6/input.txt').toString()

for (let i = 4; i <= input.length; i++) {
  const lastFour = input.substring(i-4, i)
  const charsInLastFour = {}
  for (const c of lastFour) {
    charsInLastFour[c] = true
  }
  if (Object.keys(charsInLastFour).length === 4) {
    console.log('found packet marker', lastFour, 'at index', i)
    break
  }
}

for (let i = 14; i <= input.length; i++) {
  const lastFourteen = input.substring(i-14, i)
  const charsInLastFourteen = {}
  for (const c of lastFourteen) {
    charsInLastFourteen[c] = true
  }
  if (Object.keys(charsInLastFourteen).length === 14) {
    console.log('found message marker', lastFourteen, 'at index', i)
    break
  }
}
