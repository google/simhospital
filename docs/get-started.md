# Get started

-   [Run your own instance of Simulated Hospital](#run-your-own-instance-of-simulated-hospital)
    *   [Download](#download)
    *   [Run with Bazel](#run-with-bazel)
    *   [Create Docker image](#create-docker-image)
    *   [Publish Docker image](#publish-docker-image)
-   [Run in Docker](#run-in-docker)
-   [Troubleshooting](#troubleshooting)
    *   [Error: cannot parse locations file: no such file or directory](#error-cannot-parse-locations-file-no-such-file-or-directory)

## Run the latest version of Simulated Hospital

Prerequisites: install [docker](https://www.docker.com/).

The Simulated Hospital team publishes new images of Simulated Hospital
periodically. To run the latest release:

```shell
docker run --rm -it -p 8000:8000 eu.gcr.io/simhospital-images/simhospital:latest health/simulator
```

## Run your own instance of Simulated Hospital

You can download and run the code of Simulated Hospital. The advantage of
running Simulated Hospital this way is that you can modify the source code and
customize the tool.

Prerequisites: install [bazel](https://bazel.build/) and
[git](https://git-scm.com/downloads).

### Download

Download the code into a `simhospital` local folder.

```shell
git clone https://github.com/google/simhospital.git
```

`cd` into the new folder and create a local variable for the path. This will
make the next steps easier.

```shell
cd simhospital
LOCAL_DIR=$(pwd)
```

Make sure that the code builds and that the tests run:

```shell
bazel build //...
bazel test //...
```

### Run with Bazel

Run Simulated Hospital with the default arguments. From `${LOCAL_DIR}`:

```shell
bazel run //cmd/simulator:simulator -- \
  --local_path=${LOCAL_DIR}
```

Stop the simulator with Ctrl-C.

You need to set the `local_path` argument when running with `bazel run` so that
Simulated Hospital knows where to find the default configuration files.
Alternatively, you can set the value of all of the
[data configuration arguments](./arguments.md#data-configuration). The
`local_path` argument is not needed if you
[run Simulated Hospital in Docker](#run-in-docker).

Add other command line arguments to customize the behaviour. For instance, store
the generated HL7v2 messages in a file instead of printing them on screen:

```shell
bazel run //cmd/simulator:simulator -- \
  --local_path=${LOCAL_DIR} \
  --output=file \
  --output_file=${LOCAL_DIR}/hl7_messages.out
```

See the full list of [command line arguments](./arguments.md).

### Create Docker image

You can create a Simulated Hospital image to run in Docker. The Docker image
already has all the paths configured, so you do not need any extra flags.

Prerequisites: install [docker](https://www.docker.com/).

From `${LOCAL_DIR}`:

1.  Build the image. You will see the name of the repository and tag
    *(bazel:simhospital_container_image)* printed at the end:

    ```shell
    $ bazel run //:simhospital_container_image
    INFO: Analyzed target //:simhospital_container_image (0 packages loaded, 0 targets configured).
    INFO: Found 1 target...
    Target //:simhospital_container_image up-to-date:
      bazel-bin/simhospital_container_image-layer.tar
    INFO: Elapsed time: 0.655s, Critical Path: 0.48s
    INFO: 13 processes: 13 linux-sandbox.
    INFO: Build completed successfully, 18 total actions
    INFO: Build completed successfully, 18 total actions
    Loaded image ID: sha256:b5693797e9104264d01d9cd853e97f3cf5468c6107ac7f85d0db095e6c26690b
    Tagging b5693797e9104264d01d9cd853e97f3cf5468c6107ac7f85d0db095e6c26690b as bazel:simhospital_container_image
    ```

1.  Check that the image has been loaded in your Docker images:

    ```shell
    $ docker images
    REPOSITORY                          TAG                           IMAGE ID            CREATED             SIZE
    bazel                               simhospital_container_image   b5693797e910        50 years ago        174MB
    ```

See [Run in Docker](#run-in-docker) for how to run the image in Docker.

### Publish Docker image

You can create your own Docker image of Simulated Hospital and upload it to your
container registry.

From `${LOCAL_DIR}`:

1.  In the file `BUILD.bazel`, modify the `repository` field in the
    `simhospital_image_push` rule to point to your repository, e.g.
    `repository = "my-cool-repo/simhospital"`.

1.  Push the image to the repository:

    ```shell
    bazel run //:simhospital_image_push
    ```

1.  You can then download it with:

    ```shell
    docker pull gcr.io/my-cool-repo/simhospital
    ```

## Run in Docker

Follow these steps to run Simulated Hospital in a Docker container:

1.  Get a Simulated Hospital Docker image. You can either
    [create a Docker image from your local code](#create-docker-image), or
    download an image from a registry (see
    [Publish Docker image](#publish-docker-image) for how to push and pull your
    image to/from a registry).

1.  Create a variable to make the next steps easier.

    *   If you downloaded the Docker image from a registry (replace
        <my-cool-repo> with your registry):

        ```shell
        IMAGE=gcr.io/my-cool-repo/simhospital
        ```

    *   If you created the image with bazel:

        ```shell
        IMAGE=bazel:simhospital_container_image
        ```

1.  Run the image:

    ```shell
    docker run --rm -it -p 8000:8000 $IMAGE health/simulator
    ```

If you want to load your own configuration files instead of the default ones,
you need to mount the files in the default locations, for instance:

```shell
docker run --rm -it -p 8000:8000 -v ABSOLUTE_PATH_TO_LOCAL_ALLERGIES_FILE:/configs/hl7_messages/allergies.csv $IMAGE health/simulator
```

Alternatively, you can copy the file somewhere else, and use the command line
arguments to point to it:

```shell
docker run --rm -it -p 8000:8000 -v ABSOLUTE_PATH_TO_LOCAL_ALLERGIES_FILE:/configs/allergies.csv $IMAGE health/simulator --allergies_file=configs/allergies.csv
```

See [the command line arguments](./arguments.md) for more configuration files
that you can use.

## Troubleshooting

### Error: cannot parse locations file: no such file or directory

The file with the locations is the first file that Simulated Hospital tries to
open. If it cannot be found, it is very likely that Simulated Hospital does not
know where to find the configuration files. When running locally with `bazel
run`, set the `-local_path` argument to tell Simulated Hospital where to find
your files.

Full error:

```shell
Cannot create Hospital Runner
error="cannot create default hospital config: cannot create Location Manager: cannot parse locations file configs/hl7_messages/locations.yml:: open configs/hl7_messages/locations.yml: no such file or directory
```
