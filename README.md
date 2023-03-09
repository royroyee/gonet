# GoNet
**Golang + Network**

## Introduction
이 레파지토리는 컴퓨터 네트워크에 대한 개인적인 학습을 목표로 만들어졌습니다. 네트워크에 대해 학습한 내용을 작성합니다. 

네트워크에 대한 가장 기본적인 이론부터 심화 개념 까지 다루고, 네트워크에 관련된 다양한 개념들을 작성합니다.

배운 개념들을 토대로 동시성과 병렬성을 지원하고, 네트워크 프로그래밍에 다양한 패키지와 라이브러리를 제공하는 Go 언어로 직접 코드를 작성하여 올립니다.

또한 서버, 클라우드에 대한 내용과 Kubernetes, Docker 에 관련된 Network 요소들도 다루어 배운 네트워크 지식들을 응용하고 더 깊게 배웁니다.

미흡한 내용들이 있을 수 있으며, 지속적으로 업데이트 할 예정입니다.

## Sections

1. [**First Steps Network**](https://github.com/royroyee/gonet/tree/main/01-first-steps-network)
   - 네트워크에 대한 가장 기본적인 용어들, 시작점
2. [**Rules**](https://github.com/royroyee/gonet/tree/main/02-rules)
   - 네트워크 규칙. 프로토콜과 OSI 7계층
3. [**Layer**](https://github.com/royroyee/gonet/tree/main/03-layer)
   - OSI 7계층에 대해 자세하게. IP,라우터,TCP 등 핵심 개념들을 계층으로 나누어 학습
   
   - [**Layer 1 Physical Layer**](https://github.com/royroyee/gonet/tree/main/03-layer/01-physical-layer)
      - LAN 카드(NIC), 허브
     
   - [**Layer 2 Data Link Layer**](https://github.com/royroyee/gonet/tree/main/03-layer/02-data-link-layer)
     - 이더넷, MAC 주소, ARP 프로토콜
     
   - [**Layer 3 Network Layer**](https://github.com/royroyee/gonet/tree/main/03-layer/03-network-layer)
     - 라우터, IP, 게이트웨이, DHCP
     - [Overlay Network](https://github.com/royroyee/gonet/blob/main/03-layer/03-network-layer/Overlay.md)
     
   - [**Layer 4 Transport Layer**](https://github.com/royroyee/gonet/tree/main/03-layer/04-transport-layer)
     - TCP/UDP
     - [Example Code](https://github.com/royroyee/gonet/tree/main/03-layer/04-transport-layer/example)
         
   - [**Layer 5 Application Layer**](https://github.com/royroyee/gonet/tree/main/03-layer/05-application-layer)
     - Session, Presentation Layer 포함
     - 웹 서버, HTTP, DNS
     
   - [**Processing**](https://github.com/royroyee/gonet/tree/main/03-layer/06-processing)
     - 랜 카드, 스위치/라우터, 웹 서버에서의 데이터 전달 및 처리

4. [**Cloud**](https://github.com/royroyee/gonet/tree/main/cloud)
    - Cloud computing 

5. [**Docker**](https://github.com/royroyee/gonet/tree/main/docker)
   - Docker Network

6. [**Kubernetes**](https://github.com/royroyee/gonet/tree/main/kubernetes)
    - Kubernetes Network
      - [CNI](https://github.com/royroyee/gonet/blob/main/kubernetes/cni/README.md)
      - [Flannel](https://github.com/royroyee/gonet/tree/main/kubernetes/flannel)
        - [Flannel Code Analysis](https://github.com/royroyee/gonet/tree/main/kubernetes/flannel/code-analysis)

7. [**Linux**](https://github.com/royroyee/gonet/tree/main/linux)
   - iptables
   - Bridge
   - IP Masquerade


### 예정된 사항들, 네트워크 관련 프로젝트 레파지토리
1. 네트워크 심화 개념
2. [rb](https://github.com/boanlab/rb) : REST API Benchmarking tool 
3. Kubernetes CNI + SmartNIC
4. HTTP
5. gRPC
