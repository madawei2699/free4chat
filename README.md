# free4chat

[free4.chat](https://free4.chat/) is a __real-time__ audio chat service.

It is designed by the [local first](https://www.inkandswitch.com/local-first/) and `privacy first` principle, and is very easy to use.

> :warning: **This project is just using for technical test purpose, use at all your risk!**
>
> :warning: **There is freedom of speech, but I cannot guarantee freedom after speech.** (- Idi Amin)

## Features

- **Real-time Communicating**
  - [x] Voice chat in room
  - [x] Text chat in room, can sent text or emoji
    - [ ] Persist text messages to browser localStorage
  - [ ] Can send arbitrary data by WebRTC datachannel 🚩
  - [ ] Room permission setting, like public/private type setting
    - private room can't been seen on room discovery, and it needs password to enter. The password is [End-to-End Encryption](https://blog.excalidraw.com/end-to-end-encryption/), server only need check the answer which given by the client like the `PoW` in blockchain.
    - [ ] Public rooms discovery, like hot room list or filter rooms by type/tag
- **Real-time Collaborating**  
  - [ ] User real-time collaboration, like whiteboard, you draw I guess, etc.
    - [ ] Use [CRDT](https://crdt.tech/) to impelement real-time collaboration
      - https://github.com/liveblocks/liveblocks
      - https://github.com/derekkraan/delta_crdt_ex
    - [ ] Whiteboard
      - https://github.com/tldraw/tldraw
- **Real-time Contesting**
  - [ ] Robot user, like game robot who can play or facilitate game
    - robot use [Web Speech API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Speech_API) to play with user in room
    - robot can play some voice games like language learning, technical interview, etc.
      - [Gartic Phone - The Telephone Game](https://garticphone.com/lobby)
      - [ESL Game - Not only practicing English speaking](https://esl.bmpi.dev/)
      - [Gartic.io - Draw, Guess, WIN](https://gartic.io/)

## Architecture

- **Common**
  - [x] Use [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) to replace http protocol of JSON-RPC
  - Compatibility
    - [ ] Make Safari(WebKit) compatibility better

- **Tech Stack**
  - [x] Use Elixir/Phoenix to rewrite the backend code
  - [x] Use Recat/Next.js to rewrite the frontend code

- **Infra**
  - [x] ~~Use docker to deploy to PaaS platform like [Railway](https://railway.app/) or [Fly](https://fly.io/)~~ 
  - [x] Use GitHub Actions + AWS Lightsail to deploy backend server
  - [x] Backend service cluster, auto scaling, load balancing, etc.
    - Backend service use Elixir libcluster to build cluster
    - Frontend app use the client load balance strategy
    - [x] [TURN cluster scale](https://github.com/membraneframework/membrane_ice_plugin/issues/20)
      - Fix by start turn before the libcluster, still wait the upstream library to fix it normally
  - [ ] Security enhancement, like coturn TLS setup, end-to-end encryption, etc.
    - [ ] TURN enable TLS
  - [ ] Privacy enhancement.
  - [ ] IPV6 support.

## Contribution

If you are interested in `webRTC`, `peer-to-peer(P2P)`, `real-time collaboration(CRDT)`, `distributed system` or `robot design`, you can join this project and contact with me by [twitter](https://twitter.com/madawei2699).

## Thanks

- free4.chat Elixir version is build on the top of [Membrane Framework](https://github.com/membraneframework), thanks for their heart of open source.
- [free4.chat Golang version](https://github.com/madawei2699/free4chat/tree/golang) is build on the top of [Kraken](https://github.com/bmpi-dev/kraken), [Mornin](https://github.com/lyricat/mornin.fm), [coturn](https://github.com/coturn/coturn) and [Pion](https://github.com/pion), thanks for their heart of open source.
- These websites also inspired me:
  - [Random voice and text chat rooms that you’ll love. | Speakrandom](https://www.speakrandom.com/)
  - [Practice Speaking English Online Free - Language Practice Community](https://www.free4talk.com/)
  - [Agora Real-Time Voice and Video Engagement](https://www.agora.io/en/)
  - [An open network for secure, decentralized communication - Matrix](https://matrix.org/)