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
