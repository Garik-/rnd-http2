# -*- mode: Python -*-

docker_build('example-http2', './http2', dockerfile='deployments/Dockerfile')
docker_build('example-http1', './http1', dockerfile='deployments/Dockerfile')

yaml = helm('deployments')
k8s_yaml(yaml)