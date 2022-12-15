export interface Position {
  x: number,
  y: number,
  parent?: Position
}

export const positionToString = (position: Position): string => {
  return `${position.x}:${position.y}`
}