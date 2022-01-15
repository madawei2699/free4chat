import { API_BASE } from '@/constants'
import { uuidv4 } from '@/utils/uuid'
import Base64 from '@/utils/base64'

const constraints = {
  audio: true,
  video: false
}

const configuration:any = {
  iceServers: [],
  iceTransportPolicy: 'relay',
  bundlePolicy: 'max-bundle',
  rtcpMuxPolicy: 'require',
  sdpSemantics: 'unified-plan'
}

const _AudioContext:any = (window as any).AudioContext || (window as any).webkitAudioContext
const audioCtx = new _AudioContext()
let ucid = ''
let uid = ''
let nickname = ''
let rnameRPC = ''
let unameRPC = ''
let onConnect = (pc:any, stream:any, analyser:any, trackId:string, uid:string, nickname:string) => {}
let onDisconnect = (trackId:string) => {}
let onError = (err:any) => {}
let onResume = (err:any) => {}

let pc:any = null
let running:boolean = false

export function launch (room, _nickname, _uid, _onConnect, _onDisconnect, _onResume, _onError) {
  const uname = _uid + ':' + Base64.encode(_nickname)
  rnameRPC = encodeURIComponent(room)
  unameRPC = encodeURIComponent(uname)
  uid = _uid
  onConnect = _onConnect
  onDisconnect = _onDisconnect
  onResume = _onResume
  onError = _onError
  nickname = _nickname
  running = true
  start()
}

export function stop () {
  if (pc) {
    pc.close()
  }
  running = false
  console.log('stop: ' + running)
}

async function rpc (method, params) {
  const rpcUrl = API_BASE
  try {
    const response = await fetch(rpcUrl, {
      method: 'POST', // *GET, POST, PUT, DELETE, etc.
      mode: 'cors', // no-cors, *cors, same-origin
      cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      credentials: 'omit', // include, *same-origin, omit
      headers: {
        'Content-Type': 'application/json'
      },
      redirect: 'follow', // manual, *follow, error
      referrerPolicy: 'no-referrer', // no-referrer, *client
      body: JSON.stringify({ id: uuidv4(), method, params }) // body data type must match "Content-Type" header
    })
    return response.json() // parses JSON response into native JavaScript objects
  } catch (err) {
    console.log('fetch error', method, params, err)
    return await rpc(method, params)
  } finally {
  }
}

async function subscribe (pc:any) {
  const res = await rpc('subscribe', [rnameRPC, unameRPC, ucid])
  // console.log('subscribe', res)
  if (res.hasOwnProperty('error')) {
    console.log('try to reconnect', res.error.description)
    if (onResume) {
      const prom:any = onResume(res.error)
      prom.then(() => {
        if (running) {
          console.log('reconnect in 0.5s')
          setTimeout(async () => {
            pc.close()
            await start()
          }, 500)
        }
      })
    }
    return
  }
  if (res.data && res.data.type === 'offer') {
    console.log('subscribe offer', res.data)
    await pc.setRemoteDescription(res.data)
    const sdp = await pc.createAnswer()
    await pc.setLocalDescription(sdp)
    await rpc('answer', [rnameRPC, unameRPC, ucid, JSON.stringify(sdp)])
  }
  setTimeout(function () {
    subscribe(pc)
  }, 3000)
}

export async function start () {
  if (!running) {
    return
  }
  try {
    const servers:any = await rpc('turn', [uid])
    configuration.iceServers = servers.data
  } catch (err) {
    console.log('failed to get server', err)
    return
  }
  try {
    document.querySelectorAll('.peer').forEach((el:any) => el.remove())

    pc = new RTCPeerConnection(configuration)
    pc.createDataChannel('chat') // FIXME remove this line

    // chatChannel.onopen = (event) => {
    //   chatChannel.send('Hi you!')
    // }

    // chatChannel.onmessage = (event) => {
    //   chatChannel.send(event.data)
    // }

    pc.onicecandidate = ({ candidate }) => {
      rpc('trickle', [rnameRPC, unameRPC, ucid, JSON.stringify(candidate)])
    }

    pc.ontrack = (event) => {
      console.log('ontrack', event)

      const stream = event.streams[0]
      const sid = decodeURIComponent(stream.id)
      const id = sid.split(':')[0]
      // console.log('sid', sid)
      let name = sid.split(':')[1]
      try {
        name = Base64.decode(name)
      } catch (err) {
        console.log('failed to decode name', name, err)
      }
      // console.log(stream, id, name)

      if (id === uid) {
        return
      }

      event.track.onmute = (event) => {
        if (onDisconnect) {
          const trackId:string = (event as any).target.id as string
          console.log('onmute', trackId, event)
          onDisconnect(trackId)
        }
      }

      const analyser = audioCtx.createAnalyser();
      analyser.fftSize = 256;
      analyser.minDecibels = -80;
      analyser.maxDecibels = -10;
      analyser.smoothingTimeConstant = 0.85;
      const source = audioCtx.createMediaStreamSource(stream)
      source.connect(analyser)

      if (onConnect) {
        const trackId:string = (event as any).track.id as string
        onConnect(pc, stream, analyser, trackId, id, name)
      }
    }

    let stream
    try {
      stream = await navigator.mediaDevices.getUserMedia(constraints)
    } catch (err) {
      // document.getElementById('microphone').style.display = 'block'
      console.error(err)
      if (onError) {
        onError(err)
      }
      return
    }
    const analyser = audioCtx.createAnalyser()
    analyser.fftSize = 256
    analyser.minDecibels = -80
    analyser.maxDecibels = -10
    analyser.smoothingTimeConstant = 0.85
    const source = audioCtx.createMediaStreamSource(stream)
    const gainNode = audioCtx.createGain()
    // Reduce micphone's volume to 0.01 (not 0)
    // or safari will give non-sense data in getFloatFrequencyData() and getByteTimeDomainData()
    // after switch mornin's tab to background
    gainNode.gain.value = 0.01
    source.connect(analyser)
    analyser.connect(gainNode)
    gainNode.connect(audioCtx.destination)

    if (onConnect) {
      onConnect(pc, stream, analyser, 'me', uid, nickname)
    }

    audioCtx.resume()

    stream.getTracks().forEach((track) => {
      pc.addTrack(track, stream)
    })
    await pc.setLocalDescription(await pc.createOffer())

    const res = await rpc('publish', [rnameRPC, unameRPC, JSON.stringify(pc.localDescription)])
    console.log(res)
    if (res.data && res.data.sdp.type === 'answer') {
      await pc.setRemoteDescription(res.data.sdp)
      ucid = res.data.track
      subscribe(pc)
    }
  } catch (err) {
    if (onError) {
      onError(err)
    }
  }
}
