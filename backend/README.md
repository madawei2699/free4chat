# Free4chat

## Environment variables

Available runtime environment variables:

```shell
VIRTUAL_HOST={host passed to the endpoint config, defaults to "localhost" on non-production environments}

USE_TLS={"true" or "false", if set to "true" then https will be used and certificate paths will be required}
KEY_FILE_PATH={path to certificate key file, used when "USE_TLS" is set to true}
CERT_FILE_PATH={path to certificate file, used when "USE_TLS" is set to true}

INTEGRATED_TURN_IP={TURN server IP address}
INTEGRATED_TURN_PORT_RANGE={port range, where UDP TURN will try to open ports. By default set to 50000-59999}
INTEGRATED_TCP_TURN_PORT={port number of TCP TURN}
INTEGRATED_TLS_TURN_PORT={port number of TLS TURN, used when "INTEGRATED_TURN_PKEY" and "INTEGRATED_TURN_CERT" are provided}
INTEGRATED_TURN_PKEY={SSL certificate for TLS TURN}
INTEGRATED_TURN_PKEY={SSL private key for TLS TURN}

METRICS_SCRAPE_INTERVAL={number of seconds between `Membrane.RTC.Engine` metrics reports scrapes}
```

## Run manually

### Dependencies

#### Mac OS X

```
brew install srtp clang-format ffmpeg
```

#### Ubuntu

```
sudo apt-get install libsrtp2-dev libavcodec-dev libavformat-dev libavutil-dev
```

### To run
First install all dependencies:
```
mix deps.get
```

In order to run, type:

```
mix phx.server 
```

Then go to <http://localhost:4000/>.

## Run with docker

### To run:

**IMPORTANT** If you intend to use TLS remember that setting paths in the `.env` file is not enough. Those paths will be used inside the docker container therefore besides setting env variables you will need to mount those paths to the docker container on your own. You can do it by adding `-v` flag with proper paths to `docker` command.

Default environmental variables are available in `.env` file. To run with docker, firstly you have to set up integrated TURNs environments. You can add these 2 lines to `.env` file and add `--env-file .env` flag to `docker` command or set envionment variables omitting `.env` file with `-e` flag, as in examples below
```bash
INTEGRATED_TURN_PORT_RANGE=50000-50050
INTEGRATED_TURN_IP={IPv4 address of one of your network interfaces}
```
`INTEGRATED_TURN_PORT_RANGE` describes the range, where TURN servers will try to open ports. The bigger the range is, the more users server will be able to handle. Useful when not using the `-network=host` option to limit the UDP ports used only to ones exposed from a Docker container.
`INTEGRATED_TURN_IP` is the IP address, on which TURN servers will listen. You can get it from the output of `ifconfig` command. To make the server available from your local network, you can set it to an address like `192.168.*.*`.

To start a container with image from Docker Hub, run
```bash
$ docker run --network=host -e INTEGRATED_TURN_IP=<IPv4 address> free4chat:latest
```
or
```bash
$ docker run -p 50000-50050:50000-50050/udp -p 4000:4000/tcp -e INTEGRATED_TURN_PORT_RANGE=50000-50050 -e INTEGRATED_TURN_IP=<IPv4 address> free4chat:latest
```

***NOTE 1*** There might be a problem with running `--network=host` on `macOS`, so the latter command is recommended on that operating system.

***NOTE 2*** Remember to update the range of published ports if setting the `INTEGRATED_TURN_PORT_RANGE` to a value different than in this guide.

Alternatively, you can build docker image from source
```bash
$ docker build  -t free4chat .
```

And start a container with this image
```bash
$ docker run --network=host -e INTEGRATED_TURN_IP=<IPv4 address> free4chat
```
```bash
$ docker run -p 50000-50050:50000-50050/udp -p 4000:4000/tcp -e INTEGRATED_TURN_PORT_RANGE=50000-50050 -e INTEGRATED_TURN_IP=<IPv4 address> free4chat
```

Then go to <http://localhost:4000/>.

## OpenTelemetry
By default OpenTelemetry is turned off. You can turn it on by going to `config/runtime.exs` and changing `otel_state` to one of four possible values:
* `:local` - OpenTelemetry traces will be printed on stdout
* `:zipkin` - OpenTelemetry traces are sent to Zipkin. You can change the url traces are sent to in `config/runtime.exs`. To setup zipkin you can run this command `docker run -d -p 9411:9411 openzipkin/zipkin`.
* `:honeycomb` - OpenTelemetry traces are sent to Honeycomb. You have to specify "x-honeycomb-team", which is API KEY for this service.

## Thanks

Free4chat backend is build on top of [Membrane Framework](https://github.com/membraneframework), and the origin version is modified from [Membrane Videoroom](https://github.com/membraneframework/membrane_videoroom).
