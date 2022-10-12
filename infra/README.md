# Infrastructure for free4.chat

## Deploy backend server

### AWS Lightsail

### Cluster

### ~~Fly.io~~

__NOTE__: There are some stumbling blocks that make it difficult to deploy to Fly.io:

- IPV6: Fly.io internal is IPV6 private network, but [ei_connect does not handle IPv6](https://github.com/erlang/otp/issues/5068) which make it hard to start our backend server, because it need use CNode to start [ex_dtls](https://github.com/membraneframework/ex_dtls). Althoght we can sue `NIF` mode to start ex_dtls, but it seems [libcluster](https://github.com/bitwalker/libcluster) still can not connect the node in Fly.io IPV6.
- TURN: WebRTC need TURN server to relay the package behind on the NAT network, and the TURN need open lots of UDP ports, but Fly.io needs our server to listen on [fly-global-services](https://fly.io/docs/app-guides/udp-and-tcp/#udp-must-listen-on-the-same-port-externally-and-internally) which is conflict with the TURN server address.

<details>
  <summary>fly.io config backup</summary>

  ```
fly secrets set SECRET_KEY_BASE=my_secret_value # `mix phx.gen.secret` to generate secret value for production
```

```bash
flyctl deploy --build-arg DASHBOARD_AUTH_USERNAME=admin --build-arg DASHBOARD_AUTH_PASSWORD=xxx
```
</details>

## Deploy frontend server
