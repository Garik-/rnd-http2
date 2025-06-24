# R&D HTTP/2 + linkerd benchmark

## Requirements  
- Docker Desktop + Kubernetes enabled
- Helm
- Tilt.dev
- k6

## Usage

`tilt up`

В общем хочу сделать sidecar linked и проверить че лучше будет http1 или http2 h2c в плане RTT