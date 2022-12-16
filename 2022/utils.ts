export interface Position {
  x: number,
  y: number,
  parent?: Position
}

export const positionToString = (position: Position): string => {
  return `${position.x}:${position.y}`
}

export const permutator = <T>(inputArr: T[]): T[][] => {
  let result = [];

  const permute = (arr: T[], m = []) => {
    if (arr.length === 0) {
      result.push(m)
    } else {
      for (let i = 0; i < arr.length; i++) {
        let curr = arr.slice();
        let next = curr.splice(i, 1);
        permute(curr.slice(), m.concat(next))
      }
    }
  }
  permute(inputArr)
  return result
}