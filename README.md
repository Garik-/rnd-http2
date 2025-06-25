# R&D HTTP/2 + linkerd benchmark

## Requirements  
- Docker Desktop + Kubernetes enabled
- Helm
- Tilt.dev
- k6

## Usage

`tilt up`

## Idea

В общем хочу сделать sidecar linked и проверить че лучше будет http1 или http2 h2c в плане RTT

### Итоги
Как оказалось linkerd-proxy просто передает трафик как есть, пошлешь ему h2c будет h2c, пошлешь http/1.1 будет http/1.1 он не занимается "конвертацией трафика" это просто transparent proxy.

nginx-ingress-controller не умеет h2c, если попробовать через него передать h2c он передает http/1.1

k6 не умеет делать http запросы в h2c

Можно общаться внутри кластера по h2c если клиент и сервер его поддерживают, а преобразовывать трафик поставив envoy или sidecar proxy на go, не имеет никакого практического смысла

## TODO
Можно конечно сделать не ingress контроллер, а использовать traefik-contoller который заставить преобразовывать трафик в h2c, 

Протестировать можно нагрузку через hey
```
hey -n 1000 -c 50 http://<traefik-ip>/
hey -n 1000 -c 50 -proto h2c http://<traefik-ip>/
```