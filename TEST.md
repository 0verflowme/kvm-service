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

## Testing
1. The VM policy Used for tests can be found in ``examples`` folder with name ``kvmpolicy1.yaml``
```yaml
apiVersion: security.kubearmor.com/v1
kind: KubeArmorVirtualMachine
metadata:
  name: testvm1
  labels:
    name: vm1
    vm: true
```
2. The policy to apply on VM i'm using here to block the user to see the contents of /etc/passwd.
```yaml
apiVersion: security.kubearmor.com/v1
kind: KubeArmorHostPolicy
metadata:
  name: khp-02
spec:
  nodeSelector:
    matchLabels:
      name: vm1
  severity: 5
  file:
    matchPaths:
    - path: /etc/passwd
  action:
    Block
```
3. Apply the above VM policy using ``./karmor vm policy add policy.yaml`` also you can delete policy using ``./karmor vm policy delete policy.yaml``
4. Run ``cat /etc/passwd`` and see if you can see the contents.
5. To view the logs, you can use ``karmor log`` to see logs.
```
== Alert / 2022-01-10 17:11:03.904713 ==
Cluster Name: default
Host Name: 7b0a06b40e20
Policy Name: ew-khp-01
Severity: 5
Type: MatchedHostPolicy
Source: cat
Operation: File
Resource: /etc/passwd
Data: syscall=SYS_OPENAT fd=-100 flags=/etc/passwd
Action: Block
Result: Passed
```
Alternatively, you can exec into kubearmor docker and view ``/tmp/kubearmor.log``
```json
vagrant@kvmservice-dev:~$ docker exec -it kubearmor tail -n 3 /tmp/kubearmor.log
{"timestamp":1641834663,"updatedTime":"2022-01-10T17:11:03.903782Z","hostName":"7b0a06b40e20","hostPid":7011,"ppid":6248,"pid":7011,"uid":1000,"type":"HostLog","source":"cat","operation":"Fi
le","resource":"/usr/lib/x86_64-linux-gnu/gconv/gconv-modules.cache","data":"syscall=SYS_OPENAT fd=-100 flags=/usr/lib/x86_64-linux-gnu/gconv/gconv-modules.cache","result":"Passed"}
{"timestamp":1641834663,"updatedTime":"2022-01-10T17:11:03.903706Z","hostName":"7b0a06b40e20","hostPid":7011,"ppid":6248,"pid":7011,"uid":1000,"type":"HostLog","source":"cat","operation":"Fi
le","resource":"/etc/locale.alias","data":"syscall=SYS_OPENAT fd=-100 flags=/usr/share/locale/locale.alias","result":"Passed"}
{"timestamp":1641834663,"updatedTime":"2022-01-10T17:11:03.903734Z","hostName":"7b0a06b40e20","hostPid":7011,"ppid":6248,"pid":7011,"uid":1000,"type":"HostLog","source":"cat","operation":"Fi
le","resource":"/usr/lib/locale/C.UTF-8/LC_IDENTIFICATION","data":"syscall=SYS_OPENAT fd=-100 flags=/usr/lib/locale/C.UTF-8/LC_IDENTIFICATION","result":"Passed"}
```
**This Test was failed While Testing**
