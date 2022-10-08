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
