# Github User

To run the github.py script in a container, do the following:

Build the container image
```
podman build -t github-user .
```

Run the container image with the `-it` flag
```
podman run -it github-user
```

The `-it` flag runs the container in interactive mode so that you can enter a github username that the python script can use as its input.
