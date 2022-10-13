import { weightedRand } from "./utils"

export const AUDIO_TRACK_CONSTRAINTS: MediaTrackConstraints = {
  advanced: [
    { autoGainControl: true },
    { noiseSuppression: true },
    { echoCancellation: true },
  ],
}

export const LOCAL_PEER_ID = "local-peer"

const MBps = 1024 * 8

export const BANDWIDTH_LIMITS = {
  audio: 1 * MBps,
}

export const GET_API_SERVER_URL = () => {
  // consider TURN server currently can not scale in cluster, so we use wight random to select the server
  const choice: number = weightedRand({ 0: 0.5, 1: 0.5 })()
  const serverAddrs = [
    "wss://rtc1.free4.chat/socket",
    "wss://rtc2.free4.chat/socket",
  ]
  return serverAddrs[choice]
}
