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

### Setup

1. Run kvmservice with sudo privilages as ``sudo ./kvmservice --non-k8s``
2. Add new VM policy using karmor as ``./karmor vm add policyname.yaml``
3. Now get the script using ``./karmor vm --non-k8s getscript VMNAME``
4. Prepare a new VM and run the script on the VM , this will setup the VM to run Kubearmor using karmor , you can also Verify is VM is connected to kvmservice on the window running kvmservice binary will produce logs that it got a connection.
```
2022-01-10 22:40:09.014331	INFO	ETCD: Getting values key:/kvm-svc-identity-podip-maps/1090
2022-01-10 22:40:09.015815	INFO	ETCD: err: No data
2022-01-10 22:40:09.015865	INFO	ETCD: Getting values key:/kvm-opr-identity-to-label-maps/
2022-01-10 22:40:09.016610	INFO	ETCD: Key:/kvm-opr-identity-to-label-maps/1090 Value:def=pqr
2022-01-10 22:40:09.016670	INFO	Validated the identity from the etcd DB identity:1090 is unique for label:def=pqr
2022-01-10 22:40:09.016714	INFO	New connection received RegisterAgentIdentity: 1090 podIp: 192.168.1.8
2022-01-10 22:40:09.016750	INFO	ETCD: putting values with TTL key:/kvm-svc-identity-podip-maps/1090 value:192.168.1.8
2022-01-10 22:40:09.051421	INFO	Started Policy Streamer
```
5. Also transfer Karmor binary on VM if you don't want to install go to compile on VM
