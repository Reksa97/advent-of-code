import { readFileSync } from 'fs'

interface File { size: number, name: string }
interface Folder { name: string, children: Folder[], files: File[], parent: Folder }

const input = readFileSync('./7/input.txt').toString().split(/\n/)
const directoryTree: Folder = {
  name: '/', children: [], files: [], parent: null
}
let currentDirectory = directoryTree
let readingLs = false
input.forEach((line) => {
  if (line.at(0) === '$') {
    const cdCommandIndex = line.indexOf('cd ')
    if (cdCommandIndex >= 0) {
      readingLs = false
      const cdDir = line.substring(cdCommandIndex+3)
      if (cdDir == '..') {
        currentDirectory = currentDirectory.parent
      } else if (cdDir === '/') {
        currentDirectory = directoryTree
      } else {
        currentDirectory = currentDirectory.children.find(c => c.name === cdDir)
      }
      return
    }

    if (line.substring(2) == 'ls') {
      readingLs = true
      return
    }
  }
  if (readingLs) {
    if (line.startsWith('dir ')) {
      const directoryName = line.substring(4)
      currentDirectory.children.push({ name: directoryName, children: [], files: [], parent: currentDirectory })
    } else {
      const [size, fileName] = line.split(' ')
      const sizeInt = parseInt(size)
      currentDirectory.files.push({ size: sizeInt, name: fileName })
    }
  }
})

let sumOfTotalSizes = 0

const calculateFolderSize = (directoryTree: Folder) => {
  let size = 0
  size += directoryTree.children.map(c => calculateFolderSize(c)).reduce((accumulator, current) => current + accumulator, 0) 
  directoryTree.files.forEach(file => {
    size += file.size
  })

  if (size <= 100000) sumOfTotalSizes += size
  return size
}

const rootFolderSize = calculateFolderSize(directoryTree)
console.log('sum of total sizes (max 100000 per directory)', sumOfTotalSizes)


const TOTAL_DISK_SPACE = 70000000
const MIN_UNUSED_SPACE = 30000000

const unusedSpace = TOTAL_DISK_SPACE - rootFolderSize
const needToDeleteSize = MIN_UNUSED_SPACE - unusedSpace

let smallestDeletableFolderSize = Infinity
const findSmallestFolderToDelete = (directoryTree: Folder, minSize: number) => {
  let size = 0
  size += directoryTree.children.map(
    c => findSmallestFolderToDelete(c, minSize)
  ).reduce(
    (accumulator, current) => current + accumulator, 0
  ) 
  directoryTree.files.forEach(file => {
    size += file.size
  })

  if (size >= minSize && size < smallestDeletableFolderSize) smallestDeletableFolderSize = size
  return size
}

findSmallestFolderToDelete(directoryTree, needToDeleteSize)

console.log('smallest folder that frees up enough space has size', smallestDeletableFolderSize)
