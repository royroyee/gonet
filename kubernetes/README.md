# Kubernetes

Kubernetes Network

쿠버네티스의 네트워크에 대해서

## Section
- Kubernetes Network
- [CNI](https://github.com/royroyee/gonet/tree/main/kubernetes/cni)
- [Flannel](https://github.com/royroyee/gonet/blob/main/kubernetes/flannel/README.md)


## Kubernetes Network 구조
Kubernetes 에서 기본적으로 사용하는 네트워크 구조는 크게 2가지가 있다.

### 1. Docker Bridge Network(Kubernetes Default Network)
[Docker Network](https://github.com/royroyee/gonet/tree/main/docker)
- 쿠버네티스에서는 Pod을 생성할 때, 각 Pod 마다 고유한 IP 주소를 할당한다.
- 이때, **각 노드에서 실행 중인 컨테이너는 동일한 IP 주소를 가지게 된다.**
- 따라서 **다른 노드에서 실행 중인 컨테이너와 IP 주소가 충돌**하게 되어 **통신에 문제**가 발생할 수 있다. (다른 노드의 Pod 끼리의 통신이 불가하다.)
> 때문에 쿠버네티스는 보통 CNI Plugin 을 설치하여 이러한 IP 주소 충돌 문제, 다른 노드 간의 통신을 해결한다.


#### kube-proxy
> 쿠버네티스의 네트워크 프록시 서비스
- **각 노드에서 실행**되며, iptables 를 사용하여 각 Pod 에 대한 **로드밸런싱**을 수행하며 이를 통해 **Pod 간 통신을 가능**하게 한다.
- 쿠버네티스의 서비스 구성 중 하나인 `Cluster IP` 서비스와 `Node Port` 서비스를 지원한다. 
  - 즉, 쿠버네티스에서 **같은 노드 안에 있는 Pod 간의 통신을 가능**하게 하고 **Pod을 외부로 노출**하는 다양한 서비스를 지원한다.

### 2. CNI Plugin(Flannel, Calico..)
> 클러스터 내에서 컨테이너 네트워크를 구성하는 데 사용
- CNI Plugin 을 설치하여 컨테이너 네트워크를 구성하면, **각 노드에서 실행 중인 pod 이 동일한 가상 네트워크에 속하게 된다.**
- 따라서 **멀티 노드 클러스터에서 실행 중인 Pod 간에도 통신이 가능**해진다.
- Flannel, Calico 같은 다양한 CNI Plugin 들이 있다.
- 보통 [Overlay Network](https://github.com/royroyee/gonet/blob/main/03-layer/03-network-layer/Overlay.md) 를 사용한다.

> 주의! 
> 
> kube-proxy : 서비스 및 Pod 간의 통신 관리에 초점
>
>  Flannel(CNI Plugin) : 호스트(노드) 간 통신에 초점