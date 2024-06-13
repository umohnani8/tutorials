# Cowsay

This tutorial walks you through building a container image that runs the cowsay application.

The Containerfile installs the cowsay package and the entrypoint and cmd lines tells the container to run said application when the container is started.

To run the container, run the following command
```
podman run cowsay
```
