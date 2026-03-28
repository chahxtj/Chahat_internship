## STATEMENT-5 NOTES ON CONTAINERS AND PODMAN
<!-- refer comments for notes and specific details-->
## Containers

A container is a lightweight and portable software unit that packages an application along with all its required dependencies, libraries, and runtime environment. It ensures that the application runs consistently across different computing environments. Containers share the host operating system kernel but operate in isolated environments 
Containers are widely used in modern application development and deployment because they provide portability, efficiency, and scalability
<!--it works same in dev ,testing, production everywhere -->
<!-- no diff os same kernel but isolated proces-->

## Need for containers 
The main reason containers are used is to avoid environment-related problems. An application that works perfectly on a developer’s system may fail in production due to missing packages or version mismatches.Containers ensure consistency because the same image is used everywhere. This approach makes deployment easier and more reliable
<!-- fixed blueprint where all dependencies are already included-->

Containers operate on top of a container engine such as Podman. The architecture consists of:
-> Hardware
Host Operating System
Container Engine
Container Images
Running Containers

Containers use Linux features such as:
Namespaces (for process, network, and file system isolation)<!--- provides isolation for process, n/w etc -->
Control Groups (cgroups) for resource management <!-- controls resources limits -->
This allows multiple containers to run independently on the same host system.

## Continer tools
Container tools are software used to create, run, and manage containers
Some common tools are ->
Podman – A daemonless container engine used to build and run containers <!-- no bg central service is run as its daemonless , in docker - dockered runs here nothing like that -->
Docker – The most popular container platform with a daemon-based architecture
Buildah – A tool used to build container images
CRI-O – A container runtime specifically designed to work with Kubernetes and follows OCI standards.
<!--Kubernetes – A platform used to manage and orchestrate containers It is not a container engine, but a container orchestration platform , orchestration means merging many containers hence helps in load balancing ,scaling n/w, containers start/stop etc -->

These tools help in:
Creating container images
Running containers
Managing container lifecycle
Deploying and scaling applications

# Difference Between Container and Virtual Machine

| Basis of Comparison | Container | Virtual Machine (VM) |
|---------------------|------------|-----------------------|
| *Definition* | Lightweight environment that packages application with its dependencies | Software-based virtual computer that runs a full operating system
| *Operating System* | Shares host OS kernel | Has its own guest operating system
| *Size* | Small (MBs) | Large (GBs)
| *Startup Time* | Starts in seconds | Takes minutes to boot
| *Resource Usage* | Uses fewer resources | Uses more CPU, memory and storage
| *Isolation Level* | Process-level isolation | Hardware-level isolation 
| *Performance* | Near-native performance | Slightly slower due to hypervisor overhead
| *Portability*| Highly portable | Less portable 
| *Management Tool* | Podman, Docker | VMware, VirtualBox 
| *Use Case* | Microservices, cloud applications | Running multiple operating systems

<!-- vm virtualizes h/w and container does on os level -->

## Container Image 
A container image is a read-only layered file system used to create containers .It serves as a blueprint for running containers.
Images are built from a Containerfile and consist of multiple layers. Each instruction in the Containerfile creates a new layer. These layers are cached and reused to optimize build performance.
The lifecycle of a container image follows this flow:
Containerfile → Build Process → Image → Container Runtime → Running Container

<!--Images are immutable, meaning once created they cannot be modified any change creates a new image layer-->

## Containerfile
Containerfile is a text file that contains instructions for building a container image. The instructions are executed sequentially from top to bottom during the build process.
A typical Containerfile defines the base image, copies application files, installs dependencies, configures environment variables, exposes ports, and specifies the default command to run when the container starts.
<!--order is imp because caching and buildup speed is affected otherwise -->

## CONTAINER FLOW AND COMMANDS 
**The typical flow of a Containerfile is as follows:**

-> FROM – It specifies the base image
Eg: FROM golang:1.22

-> WORKDIR – Sets the working directory inside the container
Eg: WORKDIR /app

-> COPY – Copies files from host system to container
Eg: COPY . .

-> RUN – Executes commands during image build process (BUILD TIME)
Eg: RUN apt update

-> EXPOSE – Defines the network port used by the application
Eg: EXPOSE 8080

-> CMD – Specifies the default command to run when the container starts (RUN TIME)
Eg: CMD ["go", "run", "main.go"] ( RUNS MAIN.GO)
CMD ["echo", "Container started successfully"]

<!--RUN is executed while building the image. It is mainly used for installing packages or setting up the environment.
CMD runs when the container actually starts. It defines what process should run inside the container.
So in simple terms:
RUN prepares the environment
CMD starts the application -->

Each instruction creates a separate image layer, which improves build efficiency and caching.

## PODMAN
Podman is a container engine used to build, run, and manage containers. It is **daemonless** meaning it does not require a background service to operate. It supports rootless containers and follows OCI standards.

<!--Podman is compatible with most Docker commands-->
<!-- because it follows oci standards and hence compatible with kuberenets also as any container image following oci can run on the other if both follws oci like docker image can run on podman if it also follows oci standards (Open Container Initiative , industry standard)-->

## Common Podman Commands
-> Build Image
podman build -t chahat_img .

-> List Images
podman images

-> Run Container
podman run -it chahat_img <!--interactive mode-->

-> Run in Background
podman run -d chahat_img

-> Port Mapping
podman run -p 8080:8080 chahat_img

-> List Running Containers
podman ps

-> Stop Container
podman stop <containerid>

->Remove Container
podman rm <containerid>

-> Remove Image
podman rmi chahat_img

-> List all containers 
podman ps -a <!-- including stopped-->

->  Check logs
podman logs <containerid>

## Container Image Build Process
When the command "podman build" is executed:
The container engine reads the Containerfile
Instructions are executed sequentially
Layers are created for each step
The final image is stored locally with a tag

## Running a Container
When the command "podman run" is executed:
A container is created from the image
An isolated environment is set up
Resource limits are applied
The default CMD instruction is executed
The application starts inside the container

## Advantages of Containers
→ Lightweight and efficient
Because containers share the host OS kernel and do not require a full guest OS
→ Faster startup time
Since no OS boot process is required.
→ Portable
Same image runs across environments
→ Better resource utilization
Multiple containers can run on the same host while using minimal CPU and memory compared to virtual machines
→ Simplifies application deployment
All dependencies are packaged inside the image, so deployment issues due to missing libraries are reduced
→ Supports microservices architecture
Each service can run in a separate container, allowing independent development, deployment, and scaling

