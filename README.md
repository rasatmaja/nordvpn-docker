# GO NordVPN
Go app that interact with the NordVPN Linux client to make it easier to connect to servers and manage settings using environment variables.

### Why this project exist
I currently have multiple apps running in Docker that I want to connect to the NordVPN network. Instead of installing NordVPN in each container, I'd prefer to create a single container and route the network for the other containers through this NordVPN container.

However, I encountered a problem while creating a Docker image with the NordVPN client. Frequently, after starting the container, the app displays an error message saying it can't reach the daemon. This application is designed to address that issue by checking and attempting to start the daemon whenever a container is launched, before logging in or connecting to the NordVPN server.

Moreover, I aim to set up a system where features of NordVPN can be configured through environment variables, so I don’t have to manually enter the container console to adjust the settings.

---

## Feature
* Easy Connection Management: Connect to NordVPN servers with simple commands.
* Environment Variable Support: Manage settings and configurations through environment variables.

## Technology Stack
* Programming Language: Go
* Dependencies: Utilizes the [NordVPN](https://support.nordvpn.com/hc/en-us/articles/20196094470929-Installing-NordVPN-on-Linux-distributions) command-line interface for backend operations.

## Usage 

```sh
docker pull ghcr.io/rasatmaja/nordvpn-docker:latest
```

### Environment Variable

| ENV                               | Description                                               |
| ---                               | ---                                                       |
| NORDVPN_TOKEN                     | Token used to authenticate with NordVPN                   |
| NORDVPN_DEFAULT_CONNECT_COUNTRY   | Default country to connect to                             |
| NORDVPN_DEFAULT_TECHNOLOGY        | Connection technology, options are: NORDLYNX or OPENVPN   |
| NORDVPN_ENABLE_LAN_DISCOVERY      | Enable or disable LAN discovery                           |
| NORDVPN_ENABLE_KILL_SWITCH        | Enable or disable the kill switch feature                 |
| NORDVPN_ENABLE_IPV6               | Enable or disable IPv6                                    |
| NORDVPN_ENABLE_FIREWALL           | Enable or disable the firewall                            |
| NORDVPN_ENABLE_AUTO_CONNECT       | Enable or disable auto-connect                            |

### Docker Compose
```yml
services:
  gonordvpn:
    image: ghcr.io/rasatmaja/nordvpn-docker:latest
    container_name: gonordvpn
    environment:
      NORDVPN_TOKEN: xxxx
      NORDVPN_DEFAULT_CONNECT_COUNTRY: singapore
      NORDVPN_DEFAULT_TECHNOLOGY: NORDLYNX
      NORDVPN_ENABLE_IPV6: false
      NORDVPN_ENABLE_KILL_SWITCH: true
      NORDVPN_ENABLE_AUTO_CONNECT: true
      NORDVPN_ENABLE_LAN_DISCOVERY: true
    cap_add: 
      - NET_ADMIN
    sysctls:
      - net.ipv6.conf.all.disable_ipv6=1
    healthcheck:
      test: ["CMD", "gonordvpn", "healthcheck"]
      interval: 1m
      timeout: 10s
      retries: 5
    ports:
      - 3000:3000 # expose port for firefox
      - 3001:3001 # expose port for firefox
    restart: unless-stopped  
  firefox:
    image: lscr.io/linuxserver/firefox:latest
    container_name: firefox
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Etc/UTC
    shm_size: "1gb"
    restart: unless-stopped
    network_mode: service:gonordvpn
    depends_on:
      gonordvpn:
        condition: service_healthy
```

## What I Learn from this Project
Here are a few things I learned while building this project:

1. Building a Go app to create a script that interacts with the command-line interface of the NordVPN application client.
2. Build and deploy docker image to GitHube Package Registy
3. Set up a VSCode `devcontainer` configuration with custom settings and commands
