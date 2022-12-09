import { readFileSync } from 'fs'

const input = readFileSync('./8/input.txt').toString().split(/\n/)
let grid = []
input.forEach((line) => {
  const newLine = []
  for (const tree of line) {
    newLine.push(parseInt(tree))
  }
  grid.push(newLine)
})

const maxX = grid[0].length - 1
const maxY = grid.length - 1

let visibleTrees = 0

const getTreeHeight = (x: number, y: number): number => {
  return grid[y][x]
}

const isVisible = (x: number, y: number): boolean => {
  const treeHeight = getTreeHeight(x, y)

  let visibleFromLeft = true
  for (let xx = x - 1; xx >= 0; xx--) {
    const possiblyBlockingTreeHeight = getTreeHeight(xx, y)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      visibleFromLeft = false
      break
    }
  }

  let visibleFromRight = true
  for (let xx = x + 1; xx <= maxX; xx++) {
    const possiblyBlockingTreeHeight = getTreeHeight(xx, y)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      visibleFromRight = false
      break
    }
  }

  let visibleFromBottom = true
  for (let yy = y + 1; yy <= maxY; yy++) {
    const possiblyBlockingTreeHeight = getTreeHeight(x, yy)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      visibleFromBottom = false
      break
    }
  }

  let visibleFromTop = true
  for (let yy = y - 1; yy >= 0; yy--) {
    const possiblyBlockingTreeHeight = getTreeHeight(x, yy)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      visibleFromTop = false
      break
    }
  }

  return visibleFromLeft || visibleFromRight || visibleFromBottom || visibleFromTop
}

const calculateScenicScore = (x: number, y: number): number => {
  const treeHeight = getTreeHeight(x, y)

  let visibleToLeft = 0
  for (let xx = x - 1; xx >= 0; xx--) {
    visibleToLeft++
    const possiblyBlockingTreeHeight = getTreeHeight(xx, y)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      break
    }
  }

  let visibleToRight = 0
  for (let xx = x + 1; xx <= maxX; xx++) {
    visibleToRight++
    const possiblyBlockingTreeHeight = getTreeHeight(xx, y)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      break
    }
  }

  let visibleFromBottom = 0
  for (let yy = y + 1; yy <= maxY; yy++) {
    visibleFromBottom++
    const possiblyBlockingTreeHeight = getTreeHeight(x, yy)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      break
    }
  }

  let visibleFromTop = 0
  for (let yy = y - 1; yy >= 0; yy--) {
    visibleFromTop++
    const possiblyBlockingTreeHeight = getTreeHeight(x, yy)
    if (possiblyBlockingTreeHeight >= treeHeight) {
      break
    }
  }

  return visibleToRight * visibleToLeft * visibleFromBottom * visibleFromTop
}

let treeCount = 0
let bestTreeHouseCandidate = {
  x: -1,
  y: -1,
  scenicScore: -1
}
for (let x = 0; x <= maxX; x++) {
  for (let y = 0; y <= maxY; y++) {
    treeCount++
    if (isVisible(x, y)) visibleTrees++
    const scenicScore = calculateScenicScore(x, y)
    if (scenicScore > bestTreeHouseCandidate.scenicScore)
      bestTreeHouseCandidate = {
        x,
        y,
        scenicScore: scenicScore
      }
  }
}

console.log('visible trees', visibleTrees, 'out of', treeCount)
console.log(`best treehouse candidate is x:${bestTreeHouseCandidate.x} y:${bestTreeHouseCandidate.y} and its scenic score is ${bestTreeHouseCandidate.scenicScore}`)
