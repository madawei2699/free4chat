export interface UserInfo {
  name: string
  room: string
  className?: string
  audioStream?: MediaStream | null
  peerId: string
  muteState?: boolean | false
}

export interface Message {
  peerId: string
  name: string
  text: string
}

export interface Color {
  r: string
  g: string
  b: string
}
