# R&D HTTP/2 + linkerd benchmark

## Requirements  
- Docker Desktop + Kubernetes enabled
- Helm
- Tilt.dev
- k6

## Usage

`tilt up`

В общем хочу сделать sidecar linked и проверить че лучше будет http1 или http2 h2c в плане RTT

итоги на сегодня
для того что бы linked работал контейнеры которые участвуют в общении должны содержать linkerd-proxy

сам linkerd передает трафик просто как есть даж с ingress nginx conroller
можно поставить traefik и настроить у него транспорт так что бы он делал h2c но тогда зачем линкерд не понятно

В общем если поставить просто ingress controller то он передает трафик как ест
