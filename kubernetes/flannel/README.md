# Flannel
**Flannel 이란 쿠버네티스 클러스터 내에서 네트워크 추상화 계층을 제공하는 오픈소스 소프트웨어이다.**
> 간단한 구조와 설정 방법으로 네트워크 환경에 상관없이 쉽게 사용할 수 있어, 쿠버네티스에서 가장 많이 사용되는 CNI 플러그인 중 하나

### 사전지식
#### [Kubernetes CNI(Container Network Interface)](https://github.com/royroyee/gonet/tree/main/kubernetes/cni)




### 기능
- 가상 네트워크를 구성하여 모든 노드 간에 통신이 가능하도록 한다.
  - 이를 통해 클러스터 내의 모든 노드 간에 컨테이너와 통신할 수 있다.
- Flannel 은 IP 주소를 할당받고 관리하며, 이를 통해 클러스터 내에서 런타임 파드 간에 통신을 할 수 있도록 한다.

### 동작 방식
1. Flannel 은 클러스터 내의 모든 노드에 설치된다. (**Daemon set**)
  - 노드의 개수만큼 flannel pod 이 실행된다.
  - 이 때, 각 노드에는 `flannel0` 이라는 가상 인터페이스가 생성된다.
  ```
  ubuntu@master:~$ kubectl get pods -n=kube-flannel
  
  NAME                    READY   STATUS    RESTARTS   AGE
  kube-flannel-ds-5ctcp   1/1     Running   0          44d
  kube-flannel-ds-97vct   1/1     Running   0          44d
  kube-flannel-ds-hs6rc   1/1     Running   0          44d
  ```
  
2. 각 호스트에서는 flannelId 라는 프로세스가 실행되며, etcd와 통신하여 IP 주소를 할당 받는다.
  - etcd는 flannel의 중앙 집중식 스토리지 역할을 수행
3. IP 주소를 할당 받은 노드는 이를 사용하여 클러스터 내 다른 노드 또는 pod 과 통신한다.
  - 이 때 각 노드 간에 네트워크 알고리즘들이 사용된다. (VXLAN, GRE, Host-GW 등)

### 인터페이스 확인
- flannel 은 flannel.1 과 cni0 인터페이스를 노드에 추가한다.
  - cni0인터페이스는 flannel pod 이외에 pod이 1개 이상 존재해야 생성되므로 워커 노드에 없을 수 있다.
```
ip -c -br addr(리눅스 운영체제에서 네트워크 인터페이스 정보를 출력하는 명령어)

flannel.1        UNKNOWN        10.244.0.0/32 fe80::d4f0:90ff:fe5a:271f/64 
cni0             UP             10.244.0.1/24 fe80::acf1:1dff:fe4e:b385/64 
```