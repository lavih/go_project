sudo podman rm -f go || true
sudo podman run -d --name go -v $(pwd):/tmp docker.io/golang:1.22.3-alpine sh -c 'sleep infinity'