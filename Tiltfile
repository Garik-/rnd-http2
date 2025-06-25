# -*- mode: Python -*-

docker_build('example-http2', './http2', dockerfile='deployments/Dockerfile')
docker_build('example-http1', './http1', dockerfile='deployments/Dockerfile')

yaml = helm('deployments')
k8s_yaml(yaml)

k8s_resource('example-http1', port_forwards=8080)
k8s_resource('example-http2', port_forwards='8090:8090')