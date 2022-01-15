export function uuidv4 () {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    const r = Math.random() * 16 | 0
    const v = c === 'x' ? r : (r & 0x3 | 0x8)
    return v.toString(16)
  })
}

export function uuidToColor (id) {
  let g = 0
  let b = 0
  for (let i = 0; i < id.length / 2; i++) {
    let code = id.charCodeAt(i)
    g = g + code
    code = id.charCodeAt(i * 2)
    b = b + code
  }
  return [ g % 256, b % 256 ]
}
