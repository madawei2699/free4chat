---
title: Server Process State
order: 2
---

These data from RTC1 server.

## Peers

### zeroth-peach

```text
<https://www.free4.chat/room?id=1>, { iceServers: [turn:18.182.39.157:55637?transport=udp, turn:54.238.157.32:49999?transport=tcp], iceTransportPolicy: relay, bundlePolicy: balanced, rtcpMuxPolicy: require, iceCandidatePoolSize: 0 }
```

### depressed-emerald

```text
<https://www.free4.chat/room?id=1>, { iceServers: [turn:18.182.39.157:54046?transport=udp, turn:18.182.39.157:49999?transport=tcp], iceTransportPolicy: relay, bundlePolicy: balanced, rtcpMuxPolicy: require, iceCandidatePoolSize: 0 }
```

## Get Room PID

```elixir
:global.whereis_name("1") # #PID<42721.14101.2>
```

## Room PID

```elixir
%{
  network_options: [
    integrated_turn_options: [
      ip: {0, 0, 0, 0},
      mock_ip: {18, 182, 39, 157},
      ports_range: {50000, 65355},
      cert_file: nil
    ],
    integrated_turn_domain: "rtc2.free4.chat",
    dtls_pkey: <<48, 130, 4, 162, 2, 1, 0, 2, 130, 1, 1, 0, 175, 207, 62, 97,
      128, 225, 107, 191, 3, 120, 202, 150, 59, 209, 121, 110, 133, 253, 217,
      59, 160, 78, 123, 11, 53, 222, 165, 174, 145, 246, 153, 150, 30, 197,
      ...>>,
    dtls_cert: <<48, 130, 2, 210, 48, 130, 1, 186, 2, 1, 1, 48, 13, 6, 9, 42,
      134, 72, 134, 247, 13, 1, 1, 5, 5, 0, 48, 47, 49, 11, 48, 9, 6, 3, 85, 4,
      6, 19, 2, 80, 76, 49, 15, 48, 13, ...>>
  ],
  peer_channels: %{
    "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a" => #PID<42721.14475.2>,
    "7e5ba85d-ae0d-4132-b76a-d2c1ae965869" => #PID<0.4839.2>
  },
  room_id: "1",
  rtc_engine: #PID<42721.14103.2>,
  simulcast?: false,
  trace_ctx: %{}
}
```

## Channel Peer PID

```elixir
%Phoenix.Socket{
  assigns: %{
    peer_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
    room_id: "1",
    room_pid: #PID<42721.14101.2>
  },
  channel: Free4chatWeb.PeerChannel,
  channel_pid: #PID<0.4839.2>,
  endpoint: Free4chatWeb.Endpoint,
  handler: Free4chatWeb.UserSocket,
  id: nil,
  joined: true,
  join_ref: "5",
  private: %{log_handle_in: :debug, log_join: :info},
  pubsub_server: Free4chat.PubSub,
  ref: nil,
  serializer: Phoenix.Socket.V2.JSONSerializer,
  topic: "room:1",
  transport: :websocket,
  transport_pid: #PID<0.4836.2>
}
```

## RTC Engine PID

```elixir
%Membrane.Core.Pipeline.State{
  module: Membrane.RTC.Engine,
  synchronization: %{
    clock_provider: %{choice: :auto, clock: nil, provider: nil},
    clock_proxy: #PID<42721.14104.2>,
    timers: %{}
  },
  internal_state: %{
    component_path: "pipeline@<0.14103.2>",
    display_manager: nil,
    endpoints: %{
      "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a" => %Membrane.RTC.Engine.Endpoint{
        id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
        inbound_tracks: %{
          "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8" => %Membrane.RTC.Engine.Track{
            type: :audio,
            stream_id: "4df6c1a4-9184-40de-af17-c3e62afb1d1f",
            id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
            origin: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
            fmtp: %ExSDP.Attribute.FMTP{
              pt: 111,
              profile_level_id: nil,
              level_asymmetry_allowed: nil,
              packetization_mode: nil,
              max_mbps: nil,
              max_smbps: nil,
              max_fs: nil,
              max_dpb: nil,
              max_br: nil,
              maxaveragebitrate: nil,
              maxplaybackrate: nil,
              minptime: 10,
              stereo: nil,
              cbr: nil,
              useinbandfec: true,
              usedtx: nil,
              profile_id: nil,
              max_fr: nil,
              apt: nil,
              rtx_time: nil,
              repair_window: nil,
              dtmf_tones: nil,
              redundant_payloads: nil
            },
            encoding: :OPUS,
            simulcast_encodings: [],
            clock_rate: 48000,
            format: [:RTP, :raw],
            active?: true,
            metadata: %{"active" => true},
            ctx: %{
              Membrane.WebRTC.Extension => [
                %ExSDP.Attribute.Extmap{
                  id: 3,
                  uri: "http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
                  direction: nil,
                  attributes: []
                }
              ]
            }
          }
        }
      },
      "7e5ba85d-ae0d-4132-b76a-d2c1ae965869" => %Membrane.RTC.Engine.Endpoint{
        id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
        inbound_tracks: %{
          "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca" => %Membrane.RTC.Engine.Track{
            type: :audio,
            stream_id: "3743bf60-0462-402a-8c9c-371f6b68009a",
            id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca",
            origin: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
            fmtp: %ExSDP.Attribute.FMTP{
              pt: 111,
              profile_level_id: nil,
              level_asymmetry_allowed: nil,
              packetization_mode: nil,
              max_mbps: nil,
              max_smbps: nil,
              max_fs: nil,
              max_dpb: nil,
              max_br: nil,
              maxaveragebitrate: nil,
              maxplaybackrate: nil,
              minptime: 10,
              stereo: nil,
              cbr: nil,
              useinbandfec: true,
              usedtx: nil,
              profile_id: nil,
              max_fr: nil,
              apt: nil,
              rtx_time: nil,
              repair_window: nil,
              dtmf_tones: nil,
              redundant_payloads: nil
            },
            encoding: :OPUS,
            simulcast_encodings: [],
            clock_rate: 48000,
            format: [:RTP, :raw],
            active?: true,
            metadata: %{"active" => true},
            ctx: %{
              Membrane.WebRTC.Extension => [
                %ExSDP.Attribute.Extmap{
                  id: 3,
                  uri: "http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
                  direction: nil,
                  attributes: []
                }
              ]
            }
          }
        }
      }
    },
    filters: %{
      "3e615a3b-d2af-4349-b2a1-3991ca081ce1:4571a8a6-7a27-48c0-a0ef-415380a62ef9" => %Membrane.RTP.DepayloaderBin{
        depayloader: Membrane.RTP.Opus.Depayloader,
        clock_rate: 48000
      },
      "4246f1fd-0ea3-4bb6-ae1f-ac2236077af8:c76bcf95-ddd9-4cd5-8046-720d3ba70ddd" => %Membrane.RTP.DepayloaderBin{
        depayloader: Membrane.RTP.Opus.Depayloader,
        clock_rate: 48000
      },
      "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8" => %Membrane.RTP.DepayloaderBin{
        depayloader: Membrane.RTP.Opus.Depayloader,
        clock_rate: 48000
      },
      "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca" => %Membrane.RTP.DepayloaderBin{
        depayloader: Membrane.RTP.Opus.Depayloader,
        clock_rate: 48000
      },
      "ce48a73f-27c6-4a5f-8cfc-605721c09bc7:a04c5574-b7b5-42ec-91f9-c008c29f5cc7" => %Membrane.RTP.DepayloaderBin{
        depayloader: Membrane.RTP.Opus.Depayloader,
        clock_rate: 48000
      },
      "e95e0e61-1086-4ad2-b541-d06dd4b13bf6:3e54916b-0768-46c2-846e-f3dff7e78899" => %Membrane.RTP.DepayloaderBin{
        depayloader: Membrane.RTP.Opus.Depayloader,
        clock_rate: 48000
      }
    },
    id: "1",
    peers: %{
      "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a" => %Membrane.RTC.Engine.Peer{
        id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
        metadata: %{"displayName" => "depressed-emerald"}
      },
      "7e5ba85d-ae0d-4132-b76a-d2c1ae965869" => %Membrane.RTC.Engine.Peer{
        id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
        metadata: %{"displayName" => "zeroth-peach"}
      }
    },
    pending_subscriptions: [],
    subscriptions: %{
      "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a" => %{
        "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca" => %Membrane.RTC.Engine.Subscription{
          endpoint_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
          track_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca",
          format: :RTP,
          status: :active,
          opts: [default_simulcast_encoding: "m"]
        }
      },
      "7e5ba85d-ae0d-4132-b76a-d2c1ae965869" => %{
        "4246f1fd-0ea3-4bb6-ae1f-ac2236077af8:c76bcf95-ddd9-4cd5-8046-720d3ba70ddd" => %Membrane.RTC.Engine.Subscription{
          endpoint_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
          track_id: "4246f1fd-0ea3-4bb6-ae1f-ac2236077af8:c76bcf95-ddd9-4cd5-8046-720d3ba70ddd",
          format: :RTP,
          status: :active,
          opts: [default_simulcast_encoding: "m"]
        },
        "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8" => %Membrane.RTC.Engine.Subscription{
          endpoint_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
          track_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
          format: :RTP,
          status: :active,
          opts: [default_simulcast_encoding: "m"]
        },
        "ce48a73f-27c6-4a5f-8cfc-605721c09bc7:a04c5574-b7b5-42ec-91f9-c008c29f5cc7" => %Membrane.RTC.Engine.Subscription{
          endpoint_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
          track_id: "ce48a73f-27c6-4a5f-8cfc-605721c09bc7:a04c5574-b7b5-42ec-91f9-c008c29f5cc7",
          format: :RTP,
          status: :active,
          opts: [default_simulcast_encoding: "m"]
        },
        "e95e0e61-1086-4ad2-b541-d06dd4b13bf6:3e54916b-0768-46c2-846e-f3dff7e78899" => %Membrane.RTC.Engine.Subscription{
          endpoint_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
          track_id: "e95e0e61-1086-4ad2-b541-d06dd4b13bf6:3e54916b-0768-46c2-846e-f3dff7e78899",
          format: :RTP,
          status: :active,
          opts: [default_simulcast_encoding: "m"]
        }
      }
    },
    telemetry_label: [room_id: "1"],
    trace_context: %{}
  },
  children: %{
    {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"} => %Membrane.ChildEntry{
      name: {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"},
      module: Membrane.RTC.Engine.Endpoint.WebRTC,
      options: %Membrane.RTC.Engine.Endpoint.WebRTC{
        rtc_engine: #PID<42721.14103.2>,
        direction: :sendrecv,
        ice_name: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
        handshake_opts: [
          client_mode: false,
          dtls_srtp: true,
          pkey: <<48, 130, 4, 162, 2, 1, 0, 2, 130, 1, 1, 0, 175, 207, 62, 97,
            128, 225, 107, 191, 3, 120, 202, 150, 59, 209, 121, 110, 133, 253,
            217, 59, 160, 78, 123, ...>>,
          cert: <<48, 130, 2, 210, 48, 130, 1, 186, 2, 1, 1, 48, 13, 6, 9, 42,
            134, 72, 134, 247, 13, 1, 1, 5, 5, 0, 48, 47, 49, 11, 48, 9, 6, 3,
            ...>>
        ],
        filter_codecs: &Membrane.WebRTC.SDP.filter_mappings/1,
        log_metadata: [peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"],
        webrtc_extensions: [Membrane.WebRTC.Extension.TWCC],
        extensions: %{},
        integrated_turn_domain: "rtc2.free4.chat",
        integrated_turn_options: [
          ip: {0, 0, 0, 0},
          mock_ip: {18, 182, 39, 157},
          ports_range: {50000, 65355},
          cert_file: nil
        ],
        owner: #PID<42721.14101.2>,
        trace_context: %{},
        parent_span: {:span_ctx, 61616268124636252307583551697024517257,
         7077553910427508515, 1, [], true, false, true,
         {:otel_span_ets,
          #Function<2.11568973/1 in :otel_tracer_server.on_end/1>}},
        video_tracks_limit: nil,
        rtcp_receiver_report_interval: nil,
        rtcp_sender_report_interval: nil,
        simulcast_config: %Membrane.RTC.Engine.Endpoint.WebRTC.SimulcastConfig{
          enabled: false,
          default_encoding: #Function<2.55342776/1 in Free4chat.Room.handle_info/2>
        },
        peer_metadata: %{"displayName" => "depressed-emerald"},
        telemetry_label: [
          room_id: "1",
          peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"
        ]
      },
      component_type: :bin,
      pid: #PID<42721.14476.2>,
      clock: nil,
      sync: :membrane_no_sync,
      spec_ref: #Reference<42721.2281437309.2531000321.230833>,
      playback_sync: :synced,
      terminating?: false
    },
    {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"} => %Membrane.ChildEntry{
      name: {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"},
      module: Membrane.RTC.Engine.Endpoint.WebRTC,
      options: %Membrane.RTC.Engine.Endpoint.WebRTC{
        rtc_engine: #PID<42721.14103.2>,
        direction: :sendrecv,
        ice_name: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
        handshake_opts: [
          client_mode: false,
          dtls_srtp: true,
          pkey: <<48, 130, 4, 162, 2, 1, 0, 2, 130, 1, 1, 0, 175, 207, 62, 97,
            128, 225, 107, 191, 3, 120, 202, 150, 59, 209, 121, 110, 133, 253,
            217, 59, 160, 78, ...>>,
          cert: <<48, 130, 2, 210, 48, 130, 1, 186, 2, 1, 1, 48, 13, 6, 9, 42,
            134, 72, 134, 247, 13, 1, 1, 5, 5, 0, 48, 47, 49, 11, 48, 9, 6,
            ...>>
        ],
        filter_codecs: &Membrane.WebRTC.SDP.filter_mappings/1,
        log_metadata: [peer_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"],
        webrtc_extensions: [Membrane.WebRTC.Extension.TWCC],
        extensions: %{},
        integrated_turn_domain: "rtc2.free4.chat",
        integrated_turn_options: [
          ip: {0, 0, 0, 0},
          mock_ip: {18, 182, 39, 157},
          ports_range: {50000, 65355},
          cert_file: nil
        ],
        owner: #PID<42721.14101.2>,
        trace_context: %{},
        parent_span: {:span_ctx, 61616268124636252307583551697024517257,
         7077553910427508515, 1, [], true, false, true,
         {:otel_span_ets,
          #Function<2.11568973/1 in :otel_tracer_server.on_end/1>}},
        video_tracks_limit: nil,
        rtcp_receiver_report_interval: nil,
        rtcp_sender_report_interval: nil,
        simulcast_config: %Membrane.RTC.Engine.Endpoint.WebRTC.SimulcastConfig{
          enabled: false,
          default_encoding: #Function<2.55342776/1 in Free4chat.Room.handle_info/2>
        },
        peer_metadata: %{"displayName" => "zeroth-peach"},
        telemetry_label: [
          room_id: "1",
          peer_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"
        ]
      },
      component_type: :bin,
      pid: #PID<0.4843.2>,
      clock: nil,
      sync: :membrane_no_sync,
      spec_ref: #Reference<42721.2281437309.2531000321.197949>,
      playback_sync: :synced,
      terminating?: false
    },
    {:tee,
     "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"} => %Membrane.ChildEntry{
      name: {:tee,
       "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"},
      module: Membrane.RTC.Engine.PushOutputTee,
      options: %Membrane.RTC.Engine.PushOutputTee{
        codec: :OPUS,
        telemetry_label: [
          room_id: "1",
          peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
          track_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8:"
        ]
      },
      component_type: :element,
      pid: #PID<42721.14533.2>,
      clock: nil,
      sync: :membrane_no_sync,
      spec_ref: #Reference<42721.2281437309.2531000321.231291>,
      playback_sync: :synced,
      terminating?: false
    },
    {:tee,
     "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"} => %Membrane.ChildEntry{
      name: {:tee,
       "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
      module: Membrane.RTC.Engine.PushOutputTee,
      options: %Membrane.RTC.Engine.PushOutputTee{
        codec: :OPUS,
        telemetry_label: [
          room_id: "1",
          peer_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
          track_id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca:"
        ]
      },
      component_type: :element,
      pid: #PID<42721.14302.2>,
      clock: nil,
      sync: :membrane_no_sync,
      spec_ref: #Reference<42721.2281437309.2531000321.198284>,
      playback_sync: :synced,
      terminating?: false
    }
  },
  crash_groups: %{
    "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a" => %Membrane.Core.Parent.CrashGroup{
      name: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
      mode: :temporary,
      members: [
        endpoint: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
        tee: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"
      ],
      alive_members_pids: [#PID<42721.14476.2>, #PID<42721.14533.2>],
      triggered?: false,
      crash_initiator: nil
    },
    "7e5ba85d-ae0d-4132-b76a-d2c1ae965869" => %Membrane.Core.Parent.CrashGroup{
      name: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
      mode: :temporary,
      members: [
        endpoint: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
        tee: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"
      ],
      alive_members_pids: [#PID<0.4843.2>, #PID<42721.14302.2>],
      triggered?: false,
      crash_initiator: nil
    }
  },
  delayed_playback_change: nil,
  links: [
    %Membrane.Core.Parent.Link{
      from: %Membrane.Core.Parent.Link.Endpoint{
        child: {:tee,
         "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"},
        pad_spec: {Membrane.Pad, :output,
         {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"}},
        pad_ref: {Membrane.Pad, :output,
         {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"}},
        pid: #PID<42721.14533.2>,
        pad_props: %{options: []},
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_mode: :manual,
          demand_unit: :buffers,
          direction: :output,
          mode: :push,
          name: :output,
          options: nil
        }
      },
      to: %Membrane.Core.Parent.Link.Endpoint{
        child: {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"},
        pad_spec: {Membrane.Pad, :input,
         "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"},
        pad_ref: {Membrane.Pad, :input,
         "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"},
        pid: #PID<0.4843.2>,
        pad_props: %{
          auto_demand_size: nil,
          min_demand_factor: nil,
          options: [],
          target_queue_size: nil,
          throttling_factor: 1,
          toilet_capacity: nil
        },
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_unit: :buffers,
          direction: :input,
          mode: :pull,
          name: :input,
          options: nil
        }
      }
    },
    %Membrane.Core.Parent.Link{
      from: %Membrane.Core.Parent.Link.Endpoint{
        child: {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"},
        pad_spec: {Membrane.Pad, :output,
         {"4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
          nil}},
        pad_ref: {Membrane.Pad, :output,
         {"4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
          nil}},
        pid: #PID<42721.14476.2>,
        pad_props: %{options: []},
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_unit: :buffers,
          direction: :output,
          mode: :pull,
          name: :output,
          options: nil
        }
      },
      to: %Membrane.Core.Parent.Link.Endpoint{
        child: {:tee,
         "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"},
        pad_spec: :input,
        pad_ref: :input,
        pid: #PID<42721.14533.2>,
        pad_props: %{
          auto_demand_size: nil,
          min_demand_factor: nil,
          options: [],
          target_queue_size: nil,
          throttling_factor: 1,
          toilet_capacity: nil
        },
        pad_info: %{
          availability: :always,
          caps: :any,
          demand_mode: :auto,
          demand_unit: :buffers,
          direction: :input,
          mode: :pull,
          name: :input,
          options: nil
        }
      }
    },
    %Membrane.Core.Parent.Link{
      from: %Membrane.Core.Parent.Link.Endpoint{
        child: {:tee,
         "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
        pad_spec: {Membrane.Pad, :output,
         {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"}},
        pad_ref: {Membrane.Pad, :output,
         {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"}},
        pid: #PID<42721.14302.2>,
        pad_props: %{options: []},
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_mode: :manual,
          demand_unit: :buffers,
          direction: :output,
          mode: :push,
          name: :output,
          options: nil
        }
      },
      to: %Membrane.Core.Parent.Link.Endpoint{
        child: {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"},
        pad_spec: {Membrane.Pad, :input,
         "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
        pad_ref: {Membrane.Pad, :input,
         "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
        pid: #PID<42721.14476.2>,
        pad_props: %{
          auto_demand_size: nil,
          min_demand_factor: nil,
          options: [],
          target_queue_size: nil,
          throttling_factor: 1,
          toilet_capacity: nil
        },
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_unit: :buffers,
          direction: :input,
          mode: :pull,
          name: :input,
          options: nil
        }
      }
    },
    %Membrane.Core.Parent.Link{
      from: %Membrane.Core.Parent.Link.Endpoint{
        child: {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"},
        pad_spec: {Membrane.Pad, :output,
         {"7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca",
          nil}},
        pad_ref: {Membrane.Pad, :output,
         {"7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca",
          nil}},
        pid: #PID<0.4843.2>,
        pad_props: %{options: []},
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_unit: :buffers,
          direction: :output,
          mode: :pull,
          name: :output,
          options: nil
        }
      },
      to: %Membrane.Core.Parent.Link.Endpoint{
        child: {:tee,
         "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
        pad_spec: :input,
        pad_ref: :input,
        pid: #PID<42721.14302.2>,
        pad_props: %{
          auto_demand_size: nil,
          min_demand_factor: nil,
          options: [],
          target_queue_size: nil,
          throttling_factor: 1,
          toilet_capacity: nil
        },
        pad_info: %{
          availability: :always,
          caps: :any,
          demand_mode: :auto,
          demand_unit: :buffers,
          direction: :input,
          mode: :pull,
          name: :input,
          options: nil
        }
      }
    }
  ],
  pending_specs: %{},
  playback: %Membrane.Core.Playback{
    state: :playing,
    pending_state: nil,
    target_state: :playing,
    async_state_change: false
  }
}
```

## WebRTC Endipoint for depressed-emerald

```elixir
%Membrane.Core.Bin.State{
  module: Membrane.RTC.Engine.Endpoint.WebRTC,
  synchronization: %{
    clock: nil,
    clock_provider: %{
      choice: :auto,
      clock: #PID<42721.14104.2>,
      provider: Membrane.Parent
    },
    clock_proxy: #PID<42721.14477.2>,
    latency: 0,
    parent_clock: #PID<42721.14104.2>,
    stream_sync: :membrane_no_sync,
    timers: %{}
  },
  internal_state: %{
    direction: :sendrecv,
    display_manager: nil,
    extensions: %{},
    ice_name: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
    inbound_tracks: %{
      "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8" => %Membrane.RTC.Engine.Track{
        type: :audio,
        stream_id: "4df6c1a4-9184-40de-af17-c3e62afb1d1f",
        id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
        origin: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
        fmtp: %ExSDP.Attribute.FMTP{
          pt: 111,
          profile_level_id: nil,
          level_asymmetry_allowed: nil,
          packetization_mode: nil,
          max_mbps: nil,
          max_smbps: nil,
          max_fs: nil,
          max_dpb: nil,
          max_br: nil,
          maxaveragebitrate: nil,
          maxplaybackrate: nil,
          minptime: 10,
          stereo: nil,
          cbr: nil,
          useinbandfec: true,
          usedtx: nil,
          profile_id: nil,
          max_fr: nil,
          apt: nil,
          rtx_time: nil,
          repair_window: nil,
          dtmf_tones: nil,
          redundant_payloads: nil
        },
        encoding: :OPUS,
        simulcast_encodings: [],
        clock_rate: 48000,
        format: [:RTP, :raw],
        active?: true,
        metadata: %{"active" => true},
        ctx: %{
          Membrane.WebRTC.Extension => [
            %ExSDP.Attribute.Extmap{
              id: 3,
              uri: "http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
              direction: nil,
              attributes: []
            }
          ]
        }
      }
    },
    integrated_turn_domain: "rtc2.free4.chat",
    integrated_turn_options: [
      ip: {0, 0, 0, 0},
      mock_ip: {18, 182, 39, 157},
      ports_range: {50000, 65355},
      cert_file: nil
    ],
    outbound_tracks: %{
      "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca" => %Membrane.RTC.Engine.Track{
        type: :audio,
        stream_id: "3743bf60-0462-402a-8c9c-371f6b68009a",
        id: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca",
        origin: "7e5ba85d-ae0d-4132-b76a-d2c1ae965869",
        fmtp: %ExSDP.Attribute.FMTP{
          pt: 111,
          profile_level_id: nil,
          level_asymmetry_allowed: nil,
          packetization_mode: nil,
          max_mbps: nil,
          max_smbps: nil,
          max_fs: nil,
          max_dpb: nil,
          max_br: nil,
          maxaveragebitrate: nil,
          maxplaybackrate: nil,
          minptime: 10,
          stereo: nil,
          cbr: nil,
          useinbandfec: true,
          usedtx: nil,
          profile_id: nil,
          max_fr: nil,
          apt: nil,
          rtx_time: nil,
          repair_window: nil,
          dtmf_tones: nil,
          redundant_payloads: nil
        },
        encoding: :OPUS,
        simulcast_encodings: [],
        clock_rate: 48000,
        format: [:RTP, :raw],
        active?: true,
        metadata: %{"active" => true},
        ctx: %{
          Membrane.WebRTC.Extension => [
            %ExSDP.Attribute.Extmap{
              id: 3,
              uri: "http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
              direction: nil,
              attributes: []
            }
          ]
        }
      }
    },
    owner: #PID<42721.14101.2>,
    rtc_engine: #PID<42721.14103.2>,
    simulcast_config: %Membrane.RTC.Engine.Endpoint.WebRTC.SimulcastConfig{
      enabled: false,
      default_encoding: #Function<2.55342776/1 in Free4chat.Room.handle_info/2>
    },
    telemetry_label: [
      room_id: "1",
      peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"
    ],
    track_id_to_metadata: %{
      "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8" => %{
        "active" => true
      }
    },
    video_tracks_limit: nil
  },
  playback: %Membrane.Core.Playback{
    state: :playing,
    pending_state: nil,
    target_state: :playing,
    async_state_change: false
  },
  children: %{
    endpoint_bin: %Membrane.ChildEntry{
      name: :endpoint_bin,
      module: Membrane.WebRTC.EndpointBin,
      options: %Membrane.WebRTC.EndpointBin{
        inbound_tracks: [],
        outbound_tracks: [],
        direction: :sendrecv,
        handshake_opts: [
          client_mode: false,
          dtls_srtp: true,
          pkey: <<48, 130, 4, 162, 2, 1, 0, 2, 130, 1, 1, 0, 175, 207, 62, 97,
            128, 225, 107, 191, 3, 120, 202, 150, 59, 209, 121, 110, 133, 253,
            217, 59, 160, 78, ...>>,
          cert: <<48, 130, 2, 210, 48, 130, 1, 186, 2, 1, 1, 48, 13, 6, 9, 42,
            134, 72, 134, 247, 13, 1, 1, 5, 5, 0, 48, 47, 49, 11, 48, 9, 6,
            ...>>
        ],
        rtcp_receiver_report_interval: nil,
        rtcp_sender_report_interval: nil,
        filter_codecs: &Membrane.WebRTC.SDP.filter_mappings/1,
        extensions: [Membrane.WebRTC.Extension.TWCC],
        log_metadata: [peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"],
        integrated_turn_options: [
          ip: {0, 0, 0, 0},
          mock_ip: {18, 182, 39, 157},
          ports_range: {50000, 65355},
          cert_file: nil
        ],
        simulcast?: false,
        trace_metadata: [ice_name: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"],
        trace_context: %{},
        parent_span: {:span_ctx, 61616268124636252307583551697024517257,
         13997271554974898790, 1, [], true, false, true,
         {:otel_span_ets,
          #Function<2.11568973/1 in :otel_tracer_server.on_end/1>}},
        telemetry_label: [
          room_id: "1",
          peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"
        ]
      },
      component_type: :bin,
      pid: #PID<42721.14480.2>,
      clock: nil,
      sync: :membrane_no_sync,
      spec_ref: #Reference<42721.2281437309.2531000321.230837>,
      playback_sync: :synced,
      terminating?: false
    }
  },
  delayed_playback_change: nil,
  name: {:endpoint, "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a"},
  pads_info: %{
    input: %{
      accepted_caps: :any,
      availability: :on_request,
      demand_unit: :buffers,
      direction: :input,
      mode: :pull,
      name: :input,
      options: nil
    },
    output: %{
      accepted_caps: :any,
      availability: :on_request,
      demand_unit: :buffers,
      direction: :output,
      mode: :pull,
      name: :output,
      options: nil
    }
  },
  pads_data: %{
    {Membrane.Pad, :input,
     "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"} => %Membrane.Bin.PadData{
      ref: {Membrane.Pad, :input,
       "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
      options: %{},
      accepted_caps: :any,
      availability: :on_request,
      direction: :input,
      mode: :pull,
      name: :input,
      link_id: {#Reference<42721.2281437309.2531000321.231188>,
       #Reference<42721.2281437309.2531000321.231190>},
      endpoint: %Membrane.Core.Parent.Link.Endpoint{
        child: :endpoint_bin,
        pad_spec: {Membrane.Pad, :input,
         "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
        pad_ref: {Membrane.Pad, :input,
         "7e5ba85d-ae0d-4132-b76a-d2c1ae965869:3ea36e7f-2abc-401b-9e39-10a9f0b5e8ca"},
        pid: #PID<42721.14480.2>,
        pad_props: %{
          auto_demand_size: nil,
          min_demand_factor: nil,
          options: [use_payloader?: false],
          target_queue_size: nil,
          throttling_factor: 1,
          toilet_capacity: nil
        },
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_unit: :buffers,
          direction: :input,
          mode: :pull,
          name: :input,
          options: [
            use_payloader?: [
              spec: "boolean()",
              default: true,
              description: "Defines if incoming stream should be payloaded based on given encoding.\nOtherwise the stream is assumed  be in RTP format.\n"
            ]
          ]
        }
      },
      linked?: true,
      response_received?: false,
      spec_ref: #Reference<42721.2281437309.2531000321.231192>,
      demand_unit: :buffers
    },
    {Membrane.Pad, :output,
     {"4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
      nil}} => %Membrane.Bin.PadData{
      ref: {Membrane.Pad, :output,
       {"4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
        nil}},
      options: %{},
      accepted_caps: :any,
      availability: :on_request,
      direction: :output,
      mode: :pull,
      name: :output,
      link_id: {#Reference<42721.2281437309.2531000321.231291>,
       #Reference<42721.2281437309.2531000321.231299>},
      endpoint: %Membrane.Core.Parent.Link.Endpoint{
        child: :endpoint_bin,
        pad_spec: {Membrane.Pad, :output,
         {"4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
          nil}},
        pad_ref: {Membrane.Pad, :output,
         {"4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8",
          nil}},
        pid: #PID<42721.14480.2>,
        pad_props: %{options: [extensions: [], use_depayloader?: false]},
        pad_info: %{
          availability: :on_request,
          caps: :any,
          demand_unit: :buffers,
          direction: :output,
          mode: :pull,
          name: :output,
          options: [
            extensions: [
              spec: "[Membrane.RTP.SessionBin.extension_t()]",
              default: [],
              description: "List of general extensions that will be applied to the SessionBin's output pad"
            ],
            use_depayloader?: [
              spec: "boolean()",
              default: true,
              description: "Defines if the outgoing stream should get depayloaded.\n\nThis option should be used as a convenience, it is not necessary as the new track notification\nreturns a depayloading filter's definition that can be attached to the output pad\nto work the same way as with the option set to true.\n"
            ]
          ]
        }
      },
      linked?: true,
      response_received?: false,
      spec_ref: #Reference<42721.2281437309.2531000321.231300>,
      demand_unit: :buffers
    }
  },
  parent_pid: #PID<42721.14103.2>,
  crash_groups: %{},
  children_log_metadata: [
    parent_path: ["pipeline@<0.14103.2>"],
    rtc_engine_id: "1"
  ],
  links: [],
  pending_specs: %{}
}
```

## Tee #PID<42721.14533.2>

```elixir
%Membrane.Core.Element.State{
  module: Membrane.RTC.Engine.PushOutputTee,
  type: :filter,
  name: {:tee,
   "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8"},
  internal_state: %{
    caps: %Membrane.RTP{},
    codec: :OPUS,
    telemetry_label: [
      room_id: "1",
      peer_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a",
      track_id: "4ccc918f-fc26-48a2-bd3b-16bda46e2a6a:b5b38167-8644-4028-b3dd-dfeccf3b1bf8:"
    ]
  },
  pads_info: %{
    input: %{
      accepted_caps: :any,
      availability: :always,
      demand_mode: :auto,
      demand_unit: :buffers,
      direction: :input,
      mode: :pull,
      name: :input,
      options: nil
    },
    output: %{
      accepted_caps: :any,
      availability: :on_request,
      demand_mode: :manual,
      demand_unit: :buffers,
      direction: :output,
      mode: :push,
      name: :output,
      options: nil
    }
  },
  pads_data: %{
    :input => %Membrane.Element.PadData{
      accepted_caps: :any,
      availability: :always,
      caps: %Membrane.RTP{},
      start_of_stream?: true,
      end_of_stream?: false,
      direction: :input,
      mode: :pull,
      name: :input,
      ref: :input,
      options: %{},
      pid: #PID<42721.14537.2>,
      other_ref: :output,
      input_queue: nil,
      demand: 2846,
      demand_mode: :auto,
      demand_unit: :buffers,
      other_demand_unit: :buffers,
      auto_demand_size: 4000,
      sticky_messages: [],
      toilet: nil,
      associated_pads: []
    },
    {Membrane.Pad, :output, {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"}} => %Membrane.Element.PadData{
      accepted_caps: :any,
      availability: :on_request,
      caps: %Membrane.RTP{},
      start_of_stream?: true,
      end_of_stream?: false,
      direction: :output,
      mode: :push,
      name: :output,
      ref: {Membrane.Pad, :output,
       {:endpoint, "7e5ba85d-ae0d-4132-b76a-d2c1ae965869"}},
      options: %{},
      pid: #PID<0.4989.2>,
      other_ref: {Membrane.Pad, :input, 2237834518},
      input_queue: nil,
      demand: nil,
      demand_mode: :manual,
      demand_unit: :buffers,
      other_demand_unit: :buffers,
      auto_demand_size: nil,
      sticky_messages: [],
      toilet: {Membrane.Core.Element.Toilet,
       {#PID<0.5235.2>, #Reference<0.381373121.3066953729.115270>}, 200,
       #PID<0.4989.2>, 1, 0},
      associated_pads: []
    }
  },
  parent_pid: #PID<42721.14103.2>,
  playback: %Membrane.Core.Playback{
    state: :playing,
    pending_state: nil,
    target_state: :playing,
    async_state_change: false
  },
  playback_buffer: %Membrane.Core.Element.PlaybackBuffer{q: #Qex<[]>},
  supplying_demand?: false,
  delayed_demands: MapSet.new([]),
  synchronization: %{
    clock: nil,
    latency: 0,
    parent_clock: #PID<42721.14104.2>,
    stream_sync: :membrane_no_sync,
    timers: %{}
  },
  demand_size: nil
}
```
