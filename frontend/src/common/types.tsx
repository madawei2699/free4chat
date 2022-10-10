export interface UserInfo {
  name: string
  className?: string
  audioStream?: MediaStream | null
  peerId: string
  muteState?: boolean | false
}
