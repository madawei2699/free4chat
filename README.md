# free4chat

[free4.chat](https://free4.chat/) is an instant audio conferencing service.

It is designed by the [local first](https://www.inkandswitch.com/local-first/) and `privacy first` principle, and is very easy to use.

## Features

- **Common**
  - [ ] Use [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) to replace http protocol of JSON-RPC
- **Room**
  - [ ] Text chat, can sent text or emoji
  - [ ] Room permission setting, like public/private type setting
    - private room can't been seen on room discovery, and it needs password to enter. The password is [End-to-End Encryption](https://blog.excalidraw.com/end-to-end-encryption/), server only need check the answer which given by the client like the `PoW in blockchain`<sup>*</sup>
    - [ ] Public rooms discovery, like hot room list or filter rooms by type/tag
- **User**
  - [ ] Robot user, like game robot who can play or facilitate game
    - robot use [Web Speech API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Speech_API) to play with user in room
    - robot can play some voice games like language learning, technical interview, etc.
      - [Gartic Phone - The Telephone Game](https://garticphone.com/lobby)
      - [ESL Game - Not only practicing English speaking](https://esl.bmpi.dev/)
      - [Gartic.io - Draw, Guess, WIN](https://gartic.io/)
  - [ ] User real-time collaboration, like whiteboard, you draw I guess, etc.
- **Infra**
  - [ ] Backend service cluster, auto scaling, load balancing, etc.
  - [ ] Security enhancement, like coturn TLS setup, end-to-end encryption, etc.
  - [ ] Privacy enhancement.
  - [ ] IPV6 support.

__NOTE__: 
- `*` means it can be considered a VIP feature.

## Contribution

If you are interested in `webRTC`, `peer-to-peer(P2P)`, `real-time collaboration(CRDT)`, `distributed system` or `robot design`, you can join this project and contact with me by [twitter](https://twitter.com/madawei2699).

## Thanks

free4.chat is build on the top of [Kraken](https://github.com/bmpi-dev/kraken), [Mornin](https://github.com/lyricat/mornin.fm), [coturn](https://github.com/coturn/coturn) and [Pion](https://github.com/pion), thanks for their heart of open source.

These websites also inspired me:

- [Random voice and text chat rooms that you’ll love. | Speakrandom](https://www.speakrandom.com/)
- [Practice Speaking English Online Free - Language Practice Community](https://www.free4talk.com/)
- [Agora Real-Time Voice and Video Engagement](https://www.agora.io/en/)
