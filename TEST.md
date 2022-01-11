## Setup to Run Kvmservice in non-k8s envoirnment

### Install
```
 sudo apt-get install etcd
 git clone https://github.com/0verflowme/kvm-service
 git checkout non-k8s
 ```

you'll also need Karmor in order to use kvm-service
```
git clone https://github.com/seswarrajan/kubearmor-client
```
Compile karmor using make
Navigate to src/service in Kvmservice and run make
```
cd src/service
make
```
It'll produce a binary ``kvmservice``
