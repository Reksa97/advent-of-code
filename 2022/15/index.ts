import { readFileSync } from 'fs'
import { Position, positionToString } from '../utils'

interface Sensor extends Position {
  distanceToBeacon?: number
  closestBeacon?: Position
}

const input = readFileSync('./15/input.txt').toString().split(/\n/)

const distanceBetween = (a: Position, b: Position): number => {
  return Math.abs(a.x - b.x) + Math.abs(a.y - b.y)
}

const sensors: Sensor[] = []
const beacons = new Set<string>()
let minX = Infinity
let maxX = -Infinity

input.forEach((line) => {
  const [sensor, beacon] = line.split(': ')
  const sensorPos: Sensor = {
    x: parseInt(sensor.substring(sensor.indexOf('x=') + 2, sensor.indexOf(','))),
    y: parseInt(sensor.substring(sensor.indexOf('y=') + 2))
  }
  const beaconPos: Position = {
    x: parseInt(beacon.substring(beacon.indexOf('x=') + 2, beacon.indexOf(','))),
    y: parseInt(beacon.substring(beacon.indexOf('y=') + 2))
  }
  sensorPos.distanceToBeacon = distanceBetween(sensorPos, beaconPos)
  sensorPos.closestBeacon = beaconPos

  if (minX > sensorPos.x - sensorPos.distanceToBeacon) minX = sensorPos.x - sensorPos.distanceToBeacon
  if (maxX < sensorPos.x + sensorPos.distanceToBeacon) maxX = sensorPos.x + sensorPos.distanceToBeacon

  sensors.push(sensorPos)
  beacons.add(positionToString(beaconPos))
})

const getSensorIfCannotHaveBeacon = (pos: Position, countExistingBeacons = true) => {
  if (countExistingBeacons && beacons.has(positionToString(pos))) {
    return undefined
  }

  for (const sensor of sensors) {
    if (distanceBetween(sensor, pos) <= sensor.distanceToBeacon) {
      return sensor
    }
  }
  return undefined
}

const findDistressBeacon = (maxXandY: number): Position => {
  for (let x = 0; x <= maxXandY; x++) {
    for (let y = 0; y <= maxXandY; y++) {
      const sensor = getSensorIfCannotHaveBeacon({ x, y }, false)
      if (!sensor) {
        return { x, y }
      }
      const skipToY = (sensor.y + sensor.distanceToBeacon) - Math.abs(x - sensor.x)
      y = skipToY
    }
  }
}

const y = 2000000

let cannotContainBeacon = 0
for (let x = minX; x <= maxX; x++) {
  if (getSensorIfCannotHaveBeacon({ x, y }) !== undefined) {
    cannotContainBeacon++
  }
}

console.time('part 1')
console.log(`part 1: ${cannotContainBeacon}`)
console.timeEnd('part 1')

console.time('part 2')
const distressBeacon = findDistressBeacon(4000000)
const tuningFrequency = distressBeacon.x*4000000+distressBeacon.y
console.log(`part 2: ${tuningFrequency}`)
console.timeEnd('part 2')
