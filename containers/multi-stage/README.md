## Multi-stage OCI Image Builds

This directory contains a simple Go program that serves an HTML page with ASCII art.
This example demonstrates the benefits of using multi-stage OCI builds to create small and efficient container images.

## What are Multi-Stage Builds?

Multi-stage builds are a powerful feature that allows you to use multiple `FROM` statements in your `Containerfile`.
Each `FROM` statement starts a new stage, and you can copy artifacts from one stage to another.
This is useful for building applications, as it allows you to separate the build environment from the runtime environment,
resulting in smaller and more secure images.

## Steps to Build and Run the Image

Follow these steps to build and run the OCI image using the provided Go code and `Containerfile`.
Here `podman` command is used but `Docker` CLI also can build with multiple stages (FROM directives)

### Step 1: Build the OCI Image

In your terminal, navigate to the directory containing the `serve-it.go` file and the `Containerfile`.
Then, run the following command to build the image:

```sh
podman build -t ascii-art-server .
```

### Step 2: Run the Container

Once the image is built, run the container with the following command:

```sh
podman run --rm -p 8080:8080 ascii-art-server
```

This command will start the container and map port 8080 on your host to port 8080 in the container.

### Step 3: Access the Server
Open a web browser and navigate to http://localhost:8080. You should see a webpage with the ASCII art served by the Go program.

### Summary
In this exercise, you learned about the benefits of multi-stage builds in OCI images. By using a multi-stage build,
you can create small and efficient images that contain only the necessary runtime dependencies, while
keeping the build dependencies separate. This approach enhances security and reduces the size of the final image.
