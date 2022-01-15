# INFRA

## AWS lightsail

Server type: 1 GB RAM, 1 vCPU, 40 GB SSD

## TLS/SSL

### cloudflare

- create DNS A/CNAME record for the domain `free4.chat` and point it to `vercel` DNS server.
  - Frontend static content is served from `vercel` CDN.

### ~~certbot(replaced by cloudflare proxy)~~

- Refer doc: [certbot instructions](https://certbot.eff.org/instructions?ws=other&os=ubuntufocal)

Note:

- Must disable `cloudflare proxy` for certbot to work.
- After setting up, it will give the certificate path.
  - Certificate Path: /etc/letsencrypt/live/free4.chat-0001/fullchain.pem
  - Private Key Path: /etc/letsencrypt/live/free4.chat-0001/privkey.pem

## coturn

First install coturn package:

```bash
sudo apt-get -y update
sudo apt-get install coturn
systemctl stop coturn
```

then enable turn server by changing `/etc/default/coturn` file:

```text
TURNSERVER_ENABLED=1
```

use `./configs/turnserver.conf` file to configure turn server, move it to `/etc/turnserver.conf` file.

because we want to use tcp port 443 for turn server, and this will lead a [issue](https://github.com/coturn/coturn/issues/421#issuecomment-597552224), so it need to update the `/lib/systemd/system/coturn.service` file:

- Add `AmbientCapabilities=CAP_NET_BIND_SERVICE` line to `[Service]` section.

finally do these commands to start turn server:

```bash
sudo touch /var/log/turn.log
sudo chown turnserver:turnserver /var/log/turn.log
sudo systemctl start coturn
```

After starting turn server, it needs to test whether it is working, so we can do this ICE test.

### ICE test

```python3
import hashlib
import hmac
import base64
from time import time

user = 'your-arbitrary-username'
secret = 'replace-it-with-your-secret'

ttl = 24 * 3600 # Time to live
timestamp = int(time()) + ttl
username = str(timestamp) + ':' + user
dig = hmac.new(bytes(secret, 'utf-8'), bytes(username, 'utf-8'), hashlib.sha1).digest()
password = base64.b64encode(dig).decode()

print('username: %s' % username)
print('password: %s' % password)
```

Execute this python script to get the temporary username and password, and then go to this [icetest](https://icetest.info/) website to test the ICE features, and if you can see a result like the following.

So the input parameters are:

- URL: turn:13.214.210.176:443?transport=tcp `use tcp protocol`
- Username: 1642315510:test `use python script to get the username and password`
- Credential: gMdSa3Ajy+ZQxVcpsky0moHr3n4= `use python script to get the username and password`

```text
IceGatheringState: complete
host    udp undefined:58765 N/A
relay   udp undefined:57429 1.2.3.4:12345
```

Note that there a `relay` line, that means it works!

## Nginx

**After setup the backend service**.

Firstly, install the `nginx` service:

```bash
sudo apt update
sudo apt install nginx
```

then copy the config file to `/etc/nginx/sites-enabled/`:

```bash
sudo cp free4chat.nginx /etc/nginx/sites-enabled/
```

Reload nginx:

```bash
sudo nginx -t
sudo systemctl reload nginx
```
