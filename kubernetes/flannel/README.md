# Flannel
**Flannel 이란 쿠버네티스 클러스터 내에서 네트워크 추상화 계층을 제공하는 오픈소스 소프트웨어이다.**
> 간단한 구조와 설정 방법으로 네트워크 환경에 상관없이 쉽게 사용할 수 있어, 쿠버네티스에서 가장 많이 사용되는 CNI 플러그인 중 하나

### 사전지식
#### [Kubernetes CNI(Container Network Interface)](https://github.com/royroyee/gonet/tree/main/kubernetes/cni)

### [Flannel 코드 분석](https://github.com/royroyee/gonet/tree/main/kubernetes/flannel/code-analysis)


### 기능
- 가상 네트워크를 구성하여 모든 노드 간에 통신이 가능하도록 한다.
  - 이를 통해 클러스터 내의 모든 노드 간에 컨테이너와 통신할 수 있다.
- Flannel 은 IP 주소를 할당받고 관리하며, 이를 통해 클러스터 내에서 런타임 파드 간에 통신을 할 수 있도록 한다.

### 동작 방식(Basic)
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


### 동일한 Node 에 있는 Pod 간의 통신 과정

1. **Pod 생성**


2. **CNI Plugin(Flannel) 호출** : Pod 이 생성되면 Kubernetes는 해당 Pod 에 대한 CNI Plugin을 호출
   - CNI Plugin은 Pod이 사용할 가상 네트워크 인터페이스를 생성하고, 해당 인터페이스에 IP 주소를 할당한다.
   - Flannel 은 가상 IP 주소를 할당하는 동시에, 해당 노드의 물리적인 IP 주소를 가상 IP 주소와 매핑시킨다.


3. **컨테이너 실행** : 각 컨테이너는 Pod 내에서 고유한 가상 네트워크 인터페이스를 가지게 된다.


4. **네트워크 연결** : 동일한 worker node에 있는 Pod 간의 통신이 필요할 때, 각 컨테이너는 가상 네트워크 인터페이스(veth)를 통해 `cni0`인터페이스로 패킷을 전송한다.
   - `cni0` : Flannel 의 가상 브릿지 인터페이스. Pod 간의 통신을 위한 가상 네트워크 인터페이스를 생성 및 이를 물리적인 네트워크 인터페이스와 매핑하는 역할을 수행


5. **패킷 전송** 동일한 노드에 있는 Pod 간의 통신이 필요할 때, 각 컨테이너는 가상 네트워크 인터페이스(veth)를 통해 cni0 인터페이스로 패킷을 전송한다. 
   - `cni0`는 ARP 프로토콜을 사용하여 컨테이너 간의 MAC 주소를 확인하고, 패킷을 전송한다.


6. **ARP Proxy** : Flannel은 ARP Proxy 기능을 제공하여 동일한 노드에 있는 Pod 간의 통신을 최적화 한다.
   - ARP Proxy 는 `cni0` 인터페이스에 대한 ARP 요청을 가로채서 해당 요청에 대한 응답을 바로 제공하는 역할을 한다.
   - 이를 통해 동일한 노드 내에 있는 Pod 간의 통신이 더욱 빠르고 효율적으로 이루어질 수 있다.

### 서로 다른 Node 에 있는 Pod 간의 통신 과정

1. **물리적인 네트워크**(eth) : 다른 노드 간의 Pod 간 통신은 물리적인 네트워크를 통해 이루어진다.
   - 노드 A의 네트워크 인터페이스와, 노드 B의 네트워크 인터페이스는 물리적인 네트워크 케이블로 연결되어 있다.
   > 주의. Flannel 은 Overlay Network 를 사용하지만, 이것도 물리적인 네트워크 위에서 작동되는 것이다.


2. **가상 네트워크 인터페이스**(veth) 생성 : 다른 노드 간의 Pod 간 통신을 위해 Flannel 은 먼저 가상 네트워크 인터페이스를 생성


3. **라우팅 테이블 설정** : 라우팅 테이블을 설정한다. 라우팅 테이블은 각 노드의 가상 IP 주소와 물리적인 IP 주소를 매핑하는 역할을 한다.
    - 이를 통해 Flannel은 각 노드 간의 가상 IP 주소를 사용하여 Pod 간의 통신을 가능하게 한다.


4. **패킷 전송** : 다른 노드 간의 Pod 들의 통신이 필요한 경우 각 Pod의 가상 네트워크 인터페이스(veth)를 통해 패킷이 전송된다.


5. **ARP 프로토콜** : 각 노드는 ARP 프로토콜을 사용하여 상대방의 MAC 주소를 확인한다.


6. **IPsec** : Flannel 은 다른 노드 간의 Pod 들의 통신을 보안하기 위해 `IPsec` 기술을 사용
   - 패킷을 암호화하고 안전하게 전송
   - 이를 통해 네트워크 상에서 전송되는 데이터의 안전성을 보장