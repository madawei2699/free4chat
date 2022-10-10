export interface UserInfo {
  name: string
  room: string
  className?: string
  audioStream?: MediaStream | null
  peerId: string
  muteState?: boolean | false
}
