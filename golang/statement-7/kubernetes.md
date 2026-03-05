# STATEMENT-7 NOTES ON KUBERNETES
<!-- refer comments for notes and specific details -->
# KUBERNETES
Kubernetes is a platform used to manage containers automatically
When we run only one container it is simple. But real applications run **many containers across many machines**. Managing them manually becomes difficult

Kubernetes helps in:
* starting containers
* restarting containers if they crash
* scaling applications when traffic increases
* managing updates without stopping the app

Container = runs the application
Kubernetes = manages those containers
<!-- containers run the app , kubernetes manages multiple containers automatically -->

# NEED FOR CONTAINERS 
Problems before containers:
application depended on specific OS version
required certain libraries installed
runtime versions had to match
application worked on one machine but failed on another
<!-- environment mismatch problem dependency issues across machines -->
Containers solve these problems by packaging the application together with its dependencies.
Containers:run in isolated environments,share the host OS kernel & include application + libraries + runtime

NOTE -> Containers are much lighter than virtual machines
Eg:Laptop/Server
Container A (Java App)
Container B (Node App)
Container C (Python App)
All containers share the same operating system but run independently
<!-- containers isolate applications n each container has its own environment -->

# NEED CONTAINER ORCHESTRATION 
Running a few containers manually is manageable but large systems may run hundreds or thousands of containers
So we use container orchestration tools like Kubernetes
<!-- orchestration means automatic container managementhandles large scale container systems -->
Kubernetes is a container orchestration platform that manages containers automatically

**It provides high availability ,scalability & disaster recovery**

- High Availability ->If a container or node fails, Kubernetes automatically creates a new one so the application keeps running.
<!-- self healing system  n ensures application is always available -->
- Scaling -> Kubernetes can increase or decrease the number of containers depending on system load.
Scaling can be manual & automatic
It can scale based on: CPU usage ,memory usage & traffic load
<!-- automatic scaling helps handle heavy traffic -->
- Disaster Recovery -> Kubernetes can restore applications if servers fail , Cluster configuration and application data are stored so workloads can be recreated.
<!-- improves system reliability n helps recover system after failures -->

# KUBERNETES BASIC ARCHITECTURE
A Kubernetes cluster has two main parts - Master Node & worker Nodes
Example structure:
Master node manages the cluster.
Worker nodes run the applications.
<!-- cluster = master + worker nodes , master controls system -->

# MASTER NODE (CONTROL PLANE)
The master node manages and controls the Kubernetes cluster
It mainly handles scheduling, cluster state and management tasks
<!-- main control center , manages cluster state -->
Main components in the master node:
- API Server
API server is the main entry point of the Kubernetes cluster
All requests like kubectl commands or dashboard actions first go to the API server
<!-- gateway to cluster , processes requests -->
- Scheduler
Scheduler decides which worker node should run a pod
It checks available resources like CPU and memory and then assigns the pod to the best node.
<!-- pod placement resource balancing -->
- Controller Manager
Controller manager continuously checks the cluster state
If the actual state does not match the desired state, it fixes it automatically
Controller creates one more pod.
<!-- keeps cluster stable , auto correction -->
- etcd
etcd is a key value database used by Kubernetes , It stores cluster configuration, pod information and other important data
<!-- stores cluster data , source of truth -->

# WORKER NODES
Worker nodes run the actual application containers
Main components inside worker node:
- container runtime
- kubelet
- kube-proxy <!-- application runs here -->

**Container Runtime**
Container runtime is responsible for running containers inside pods
Common examples include containerd and CRI-O
<!-- runs containers , executes container images -->
**kubelet**
kubelet runs on every worker node and communicates with the master node
It ensures that the pods assigned to the node are running properly
<!-- node agent , manages pods -->
**kube-proxy**
kube-proxy manages networking inside the cluster
It helps route traffic between pods and services
<!-- networking component, traffic routing -->

# VIRTUAL NETWORK (CNI)
Kubernetes uses a virtual network so all nodes work like a single system , Each pod gets its own IP address and can communicate with other pods
<!-- pod to pod networking , cluster communication -->

# POD AND CONTAINER
Pod is the smallest deployable unit in Kubernetes
Pods contain one or more containers
Containers inside the same pod share:network, storage and IP address
But sometimes multiple containers are used together in one pod
<!-- containers live inside pods ,pod acts as wrapper for containers -->
Each pod gets its own IP address from the cluster network, pods communicate with each other using these IP's
If a pod restarts it gets a new IP address
Because of this, we use services to provide stable access
<!-- pod ip keeps changing, service solves this issue -->

# POD LIFECYCLE
Pods are temporary objects
They may stop because of:application crash , node failure and scaling events
When a pod dies, Kubernetes automatically creates another one
The new pod gets:new IP address, new identity
<!-- pods are ephemeral , kubernetes recreates them automatically -->

# SERVICES
Services provide a stable way to access pods
Because pod IPs change frequently, services act as a permanent entry point

A service provides: stable IP & load balancing

Traffic flow eg : User → Service → Pod A / Pod B / Pod C
If one pod fails, service sends traffic to remaining pods

<!-- service hides pod changes, distributes traffic between pods -->
# REPLICASET
ReplicaSet ensures a certain number of pods are always running
eg: replicas = 3 ,Kubernetes will keep 3 pods running, If one pod crashes, another pod is created automatically
<!-- maintains desired number of pods, improves application availability -->
# DEPLOYMENT
Deployment is used to manage ReplicaSets and pods
Instead of creating pods directly, we create deployments
Deployment allows:rolling updates , rollback to previous versions & easy scaling
Working flow: Deployment → ReplicaSet → Pods
<!-- commonly used for application deployment, manages updates safely -->

# DAEMONSET
DaemonSet ensures that a specific pod runs on every node
Eg: If cluster has 4 nodes → 4 pods will run

Used for system services like:logging agents & monitoring tools
<!-- one pod per node, useful for background services -->

# NODES
Node is a machine in the Kubernetes cluster
It can be:physical server or virtual machine
Nodes run pods and containers
Master node manages worker nodes
<!-- node = machine in cluster, pods run on worker nodes -->

# KUBECTL COMMANDS
kubectl is the command line tool used to interact with a Kubernetes cluster.
Common commands:
kubectl get nodes
kubectl get pods
kubectl get deployments
kubectl get services

To create resources from YAML file:
* kubectl apply -f file.yaml

To delete resources:
* kubectl delete -f file.yaml

To see detailed information:
* kubectl describe pod pod-name

To view logs:
* kubectl logs pod-name

To scale deployment:
* kubectl scale deployment name --replicas=5
<!-- kubectl communicates with api server, main CLI tool for kubernetes -->

# Summing up
Containers run applications
Pods wrap containers
Nodes run pods
ReplicaSets maintain number of pods
Deployments manage application updates
Services provide stable networking
DaemonSets run pods on every node
Kubectl allows us to manage everything from the command line
Kubernetes combines all these components to run containerized applications in a reliable and automated way
<!-- main goal = automation and scalability , ensures applications run smoothly in production -->