---
title: Arch Draws
order: 1
---

## Membrane Framework

### Adding a track (Client)

```sequence
actor client
participant server
actor other_client
client->>server: renegotiateTracks
server->>client: custom(offerData)
client->>server: custom(sdpOffer)
par
    client->>server: custom(candidate)
and
    server->>client: custom(sdpAnswer)
    server->>client: custom(candidate)
    server->>other_client: tracksAdded
    server->>client: tracksAdded 
    rect rgb(135, 204, 232)
        note right of server: renegotiation
        server->>other_client: custom(offerData)
        other_client->>server: custom(sdpOffer)
        par
            other_client->>server: custom(candidate)
        and
            server->>other_client: custom(sdpAnswer)
            server->>other_client: custom(candidate)
        end
    end
end
```

### Adding a track (Server)

```sequence
autonumber
participant Free4chatWeb.PeerChannel
participant Free4chat.Room
participant Membrane.RTC.Engine
participant Membrane.RTC.Engine.Endpoint.WebRTC
participant Membrane.WebRTC.EndpointBin

Note right of Free4chatWeb.PeerChannel: Phoenix.Channel
Note left of Free4chat.Room: GenServer

Free4chatWeb.PeerChannel ->> Free4chat.Room: :media_event<custom<renegotiateTracks>>
Free4chat.Room ->> Membrane.RTC.Engine: :media_event<custom<renegotiateTracks>>
rect rgb(135, 204, 232)
    Note right of Membrane.RTC.Engine: Membrane.Pipeline
    Note right of Membrane.RTC.Engine.Endpoint.WebRTC: Membrane.Bin<br>{children: [Membrane.WebRTC.EndpointBin]}
    Note left of Membrane.WebRTC.EndpointBin: Membrane.Bin<br>{children: [Membrane.ICE.Endpoint, <br>Membrane.RTP.SessionBin, <br>Membrane.Funnel]}
    Membrane.RTC.Engine ->> Membrane.RTC.Engine.Endpoint.WebRTC: :custom_media_event<renegotiateTracks>
    Membrane.RTC.Engine.Endpoint.WebRTC ->> Membrane.WebRTC.EndpointBin: :signal<:renegotiate_tracks>
    Membrane.WebRTC.EndpointBin ->> Membrane.RTC.Engine.Endpoint.WebRTC: :signal<:offer_data>
    Membrane.RTC.Engine.Endpoint.WebRTC  ->> Membrane.RTC.Engine: :custom_media_event<offerData>
end
Membrane.RTC.Engine ->> Free4chatWeb.PeerChannel: custom<offerData>
```
