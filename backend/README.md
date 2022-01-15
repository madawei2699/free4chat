# Kraken

🐙 High performance WebRTC SFU implemented with pure Go.

## Architecture

Kraken SFU only supports simple group audio conferencing, more features may be added easily.

Both Unified Plan and RTCP-MUX supported, so that only one UDP port per participant despite the number of participants in a room.

### monitor [WIP]

This is the daemon that load balance all engine instances according to their system load, and it will direct all peers in a room to the same engine instance.

### engine

The engine handles rooms, all peers in a room should connect to the same engine instance. No need to create rooms, a room is just an ID to distribute streams.

Access the engine with HTTP JSON-RPC, some pseudocode to demonstrate the full procedure.

```javascript
var roomId = getUrlQueryParameter('room');
var userId = uuidv4();
var trackId;

var pc = new RTCPeerConnection(configuration);

// send ICE candidate to engine
pc.onicecandidate = ({candidate}) => {
  rpc('trickle', [roomId, userId, trackId, JSON.stringify(candidate)]);
};

// play the audio stream when available
pc.ontrack = (event) => {
  el = document.createElement(event.track.kind)
  el.id = aid;
  el.srcObject = stream;
  el.autoplay = true;
  document.getElementById('peers').appendChild(el)
};

// setup local audio stream from microphone
const stream = await navigator.mediaDevices.getUserMedia(constraints);
stream.getTracks().forEach((track) => {
  pc.addTrack(track, stream);
});
await pc.setLocalDescription(await pc.createOffer());

// RPC publish to roomId, with SDP offer
var res = await rpc('publish', [roomId, userId, JSON.stringify(pc.localDescription)]);
// publish should respond an SDP answer
var jsep = JSON.parse(res.data.jsep);
if (jsep.type == 'answer') {
  await pc.setRemoteDescription(jsep);
  trackId = res.data.track;
  subscribe(pc);
}

// RPC subscribe to roomId periodically
async function subscribe(pc) {
  var res = await rpc('subscribe', [roomId, userId, trackId]);
  var jsep = JSON.parse(res.data.jsep);
  if (jsep.type == 'offer') {
    await pc.setRemoteDescription(jsep);
    var sdp = await pc.createAnswer();
    await pc.setLocalDescription(sdp);
    // RPC anwser the subscribe offer
    await rpc('answer', [roomId, userId, trackId, JSON.stringify(sdp)]);
  }
  setTimeout(function () {
    subscribe(pc);
  }, 3000);
}

async function rpc(method, params = []) {
  const response = await fetch('http://localhost:7000', {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({id: uuidv4(), method: method, params: params})
  });
  return response.json();
}
```

## Quick Start

Setup Golang development environment at first.

```
git clone https://github.com/MixinNetwork/kraken
cd kraken && go build

cp config/engine.example.toml config/engine.toml
ip address # get your network interface name, edit config/engine.toml

./kraken -c config/engine.toml -s engine
```

Get the source code of either [kraken.fm](https://github.com/MixinNetwork/kraken.fm) or [Mornin](https://github.com/fox-one/mornin.fm), follow their guides to use your local kraken API.

## Community

Kraken is built with [Pion](https://github.com/pion/webrtc), we have discussions over their Slack.
