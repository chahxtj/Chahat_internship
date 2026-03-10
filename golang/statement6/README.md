## Statement 6: gRPC Trie Service
This implements a Trie service using gRPC in Go.
Clients can add, remove, check, and list words over the network. The project includes Kubernetes manifests and works with Podman for building and running container images.

Features:
Add and remove words in Trie
Check if a word exists
List all stored words
gRPC API for networked access
Kubernetes deployment with separate server and client pods

Folder Structure
statement6/
├── trie/       # Trie logic
├── proto/      # gRPC service definition
├── server/     # gRPC server
├── client/     # gRPC client
├── k8s/        # Kubernetes manifests
└── README.md
Setup
1. Initialize Go module
cd statement6
go mod tidy
2. Generate gRPC code
protoc --go_out=. --go-grpc_out=. proto/string_service.proto
3. Build images with Podman
podman build -t trie-server ./server
podman build -t trie-client ./client
Running Locally

Start the server:
podman run -it --rm trie-server

In a new terminal, start the client:
podman run -it --rm trie-client

Use client commands:
Add : add a word
Remove : delete a word
Check : check if a word exists
List : list all words
Quit : exit client

Kubernetes Deployment
Apply server and service:
kubectl apply -f k8s/server.yaml

Apply client pod:
kubectl apply -f k8s/client.yaml

Check pods:
kubectl get pods
kubectl logs <client-pod-name>

