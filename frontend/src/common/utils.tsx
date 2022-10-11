import {
  uniqueNamesGenerator,
  Config,
  adjectives,
  colors,
} from "unique-names-generator"

const customConfig: Config = {
  dictionaries: [adjectives, colors],
  separator: "-",
  length: 2,
}

export const randomName = () => {
  return uniqueNamesGenerator(customConfig)
}

export const saveRoomToLocalStorage = (roomName, nickName) => {
  var rooms: {}[] = JSON.parse(localStorage.getItem("rooms") || "[]")
  rooms = rooms.filter((room: any) => room.roomName !== roomName)
  rooms.push({ roomName: roomName, nickName: nickName })
  localStorage.setItem("rooms", JSON.stringify(rooms))
}

export const nameToColor = (name: string) => {
  let r = 0
  let g = 0
  let b = 0
  for (let i = 0; i < name.length / 3; i++) {
    let code = name.charCodeAt(i)
    g = g + code
    code = name.charCodeAt(i * 2)
    b = b + code
    code = name.charCodeAt(i * 3)
    r = r + code
  }
  return [r % 256, g % 256, b % 256]
}
