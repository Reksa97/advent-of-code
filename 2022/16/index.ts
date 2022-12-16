import { readFileSync } from 'fs'

type Valve = {
  name: string
  flowRate: number
  isOpen: boolean
  leadsTo: string[]
  timeToValves: {
    [key: string]: number
  }
}

const TUNNELS_START = 'tunnels lead to valves '
const TUNNEL_START = 'tunnel leads to valve '
const input = readFileSync('./16/input.txt').toString().split(/\n/)
const valves: {
  [name: string]: Valve
} = {}
input.forEach((line) => {
  const [valveRow, tunnelsRow] = line.split('; ')
  const valveParts = valveRow.split(' ')
  const valveName = valveParts.at(1)
  const valveRate = parseInt(valveParts.at(-1).split('=').at(-1))

  const hasMultipleTunnels = tunnelsRow.includes(TUNNELS_START)
  const tunnels = tunnelsRow.substring(hasMultipleTunnels ? TUNNELS_START.length : TUNNEL_START.length).split(', ')
  valves[valveName] = ({
    name: valveName,
    flowRate: valveRate,
    leadsTo: tunnels,
    isOpen: false,
    timeToValves: {}
  })
})

type ValveBTS = {
  name: string
  parent?: ValveBTS
}

const minutesToTravel = (from: Valve, to: Valve): number => {
  const queue: ValveBTS[] = []
  queue.push({ name: from.name, parent: undefined })
  const visited = new Set<string>()
  visited.add(from.name)
  while (queue.length > 0) {
    const current = queue.shift()

    if (current.name === to.name) {
      let parent = current
      let count = 0
      while (parent = parent.parent) {
        count++
      }
      return count
    }
    valves[current.name].leadsTo.forEach((adjacent, i) => {
      if (!visited[adjacent]) {
        visited[adjacent] = true
        queue.push({ name: adjacent, parent: current })
      }
    });
  }
  return -1
}

console.log('calculate times')
// Calculate time to travel between valves
for (const from of Object.keys(valves)) {
  for (const to of Object.keys(valves)) {
    if (from === to) continue
    valves[from].timeToValves[valves[to].name] = minutesToTravel(valves[from], valves[to])
    // Can go back?
  }
}

const keys = Object.keys(valves).filter(k => k !== 'AA' && valves[k].flowRate > 0)

const getPossibleRoutes = (): string[][] => {
  const possibleRoutes = []
  const route = ['AA']
  const possiblePoints = keys.slice()
  const addRest = (current: string[], points: string[]) => {

    for (let i = 0; i < points.length; i++) {
      const nextPoints = points.slice()
      const nextPoint = nextPoints.splice(i, 1)
      const next = [...current, ...nextPoint]
      const t = next.reduce((acc, cur, i, arr) => acc + (i > 0 ? valves[arr[i - 1]].timeToValves[cur] + 1 : 0), 0)

      possibleRoutes.push(next.slice())
      if (t > 32) {
        return
      }
      addRest(next, nextPoints)
    }
  }
  addRest(route, possiblePoints)
  return possibleRoutes
}

const possibleRoutes = getPossibleRoutes()

let maxReleasedPressure = 0
let routeI = 0
//const BEST_ROUTE = -1 //718
//const BEST_ROUTE = 718 //718
const BEST_ROUTE = 85635 //718
for (const route of possibleRoutes) {
  //if (routeI % 50000 === 0) console.log('left', possibleRoutes.length - routeI)
  routeI++
  let releasedPressure = 0
  let releasePerMinute = 0
  let minute = 0

  const nextMinute = (minuteDelta = 1) => {
    if (routeI === BEST_ROUTE && minute === 29) console.log('asd', minuteDelta, minute + minuteDelta > 30, 30 - minute)
    if (minute + minuteDelta > 30) {
      minuteDelta = 30 - minute
    }
    releasedPressure += releasePerMinute * minuteDelta

    if (routeI === BEST_ROUTE) {
      for (let i = 1; i < minuteDelta; i++) {
        console.log(`Minute ${minute + i} released ${releasePerMinute}`)
      }
    }

    minute += minuteDelta
    if (routeI === BEST_ROUTE) console.log(`Minute ${minute} released ${releasePerMinute}`)
  }

  if (routeI === BEST_ROUTE) console.log(route)
  for (let i = 0; i < route.length; i++) {
    const current = route[i]

    if (routeI === BEST_ROUTE) console.log(current)
    if (current !== 'AA') {
      nextMinute()
      releasePerMinute += valves[current].flowRate
      if (routeI === BEST_ROUTE) console.log('opened', current, valves[current].flowRate)
    }

    if (i < route.length - 1) {
      nextMinute(valves[current].timeToValves[route[i + 1]])
      /* for (let travelTime = valves[current].timeToValves[route[i + 1]]; travelTime > 0; travelTime--) {
        //if (minute >= 30) break routeLoop
      } */
    }
  }

  while (minute < 30) nextMinute()

  if (maxReleasedPressure < releasedPressure) {
    console.log('best route', routeI)
    maxReleasedPressure = releasedPressure
  }
}


console.log(`released pressure ${maxReleasedPressure}`) // 1796

