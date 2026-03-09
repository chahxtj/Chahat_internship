## CHAHAT'S ROOK CHEATSHEET
## Rook Setup Cheatsheet (Minikube)
This cheatsheet explains how to install Rook with Ceph on Minikube, create a StorageClass, provision a PersistentVolumeClaim (PVC), and experiment with VolumeSnapshots

## Prerequisites
Make sure the following tools are installed:
Rook
Ceph
Minikube
kubectl
Git

## Verify installation:
minikube version
kubectl version --client
git --version

## Start Minikube Cluster

Start a Kubernetes cluster using the QEMU driver
minikube start --driver=qemu --cpus=2 --memory=6g

Verify the cluster:
kubectl get nodes

Expected output:
NAME       STATUS   ROLES           AGE
minikube   Ready    control-plane

## Clone Rook Repository

Clone the official Rook repository

git clone https://github.com/rook/rook.git
cd rook/deploy/examples

## Install Rook Operator
Deploy required Kubernetes resources

# Install CRDs
kubectl create -f crds.yaml
# Install common resources
kubectl create -f common.yaml
# Deploy Rook operator
kubectl create -f operator.yaml

Verify operator:
kubectl -n rook-ceph get pods

Expected:
rook-ceph-operator-xxxxx   Running

## Configure Cluster for Minikube
Since Minikube is a single node cluster, modify cluster.yaml

Open the file:
nano cluster.yaml

Update the configuration:
mon:
  count: 1
  allowMultiplePerNode: true

mgr:
  count: 1
Deploy Ceph Cluster

-> Create the Ceph cluster
kubectl create -f cluster.yaml

Check cluster status:
kubectl -n rook-ceph get cephcluster

Example output:
NAME        DATADIRHOSTPATH   MONCOUNT   AGE   PHASE   MESSAGE
rook-ceph   /var/lib/rook     1          12m   Ready   Cluster created successfully

## Verify Running Pods

Check all pods created by Rook
kubectl -n rook-ceph get pods

Example:
rook-ceph-mon-a
rook-ceph-mgr-a
rook-ceph-crashcollector
csi-rbdplugin
csi-cephfsplugin
rook-ceph-operator

## Deploy Toolbox

The toolbox pod allows running Ceph commands

Create toolbox:
kubectl create -f toolbox.yaml

Enter toolbox:
kubectl -n rook-ceph exec -it deploy/rook-ceph-tools -- bash

Check cluster status:
ceph status

Exit toolbox:
exit

## Create StorageClass

Enable dynamic provisioning using RBD

kubectl create -f csi/rbd/storageclass.yaml

Verify:
kubectl get storageclass

Example:
rook-ceph-block

## Create PersistentVolumeClaim

Create pvc.yaml.

apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: rook-pvc
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: rook-ceph-block
  resources:
    requests:
      storage: 1Gi

Apply the PVC:
kubectl apply -f pvc.yaml

Verify:
kubectl get pvc

## Test PVC with Pod

- > Create pod.yaml.

apiVersion: v1
kind: Pod
metadata:
  name: pvc-test
spec:
  containers:
  - name: app
    image: busybox
    command: ["sleep","3600"]
    volumeMounts:
    - mountPath: "/data"
      name: storage
  volumes:
  - name: storage
    persistentVolumeClaim:
      claimName: rook-pvc

- > Apply:

kubectl apply -f pod.yaml

## Create VolumeSnapshotClass

Create snapshotclass.yaml

apiVersion: snapshot.storage.k8s.io/v1
kind: VolumeSnapshotClass
metadata:
  name: rook-ceph-snapclass
driver: rook-ceph.rbd.csi.ceph.com
deletionPolicy: Delete

Apply:

kubectl apply -f snapshotclass.yaml

## Create Volume Snapshot

-> Create snapshot.yaml

apiVersion: snapshot.storage.k8s.io/v1
kind: VolumeSnapshot
metadata:
  name: rook-pvc-snapshot
spec:
  volumeSnapshotClassName: rook-ceph-snapclass
  source:
    persistentVolumeClaimName: rook-pvc

Apply:

kubectl apply -f snapshot.yaml

Verify snapshot:

kubectl get volumesnapshot

## Cleanup

Delete resources:
kubectl delete -f pod.yaml
kubectl delete -f pvc.yaml
kubectl delete -f snapshot.yaml
kubectl -n rook-ceph delete cephcluster rook-ceph

Stop Minikube:
minikube stop