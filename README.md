# GoNet
**Golang + Network**

## Introduction
이 레파지토리는 컴퓨터 네트워크에 대한 개인적인 학습을 목표로 만들어졌습니다. 네트워크에 대해 학습한 내용을 작성합니다. 

네트워크에 대한 가장 기본적인 이론부터 심화 개념 까지 다루고, 네트워크에 관련된 다양한 개념들을 작성합니다.


배운 개념들을 토대로 동시성과 병렬성을 지원하고, 네트워크 프로그래밍에 다양한 패키지와 라이브러리를 제공하는 Go 언어로 직접 코드를 작성하여 올립니다.


미흡한 내용들이 있을 수 있으며, 지속적으로 업데이트 할 예정입니다.

## Sections

1. [First Steps Network](https://github.com/royroyee/gonet/tree/main/01-first-steps-network)
   - 네트워크에 대한 가장 기본적인 용어들, 시작점
2. [Rules](https://github.com/royroyee/gonet/tree/main/02-rules)
   - 네트워크 규칙. 프로토콜과 OSI 7계층
3. [Layer](https://github.com/royroyee/gonet/tree/main/03-layer)
   - OSI 7계층에 대해 자세하게. IP,라우터,TCP 등 핵심 개념들을 계층으로 나누어 학습
   - Layer 1 Physical Layer
   - [Layer 2 Data Link Layer](https://github.com/royroyee/gonet/tree/main/03-layer/02-data-link-layer)
   - [Layer 3 Network Layer](https://github.com/royroyee/gonet/tree/main/03-layer/03-network-layer)
   - [Layer 4 Transport Layer](https://github.com/royroyee/gonet/tree/main/03-layer/04-transport-layer)
      - [Example Code](https://github.com/royroyee/gonet/tree/main/03-layer/04-transport-layer/example)
         - TCP Server / Client : tcp_server.go / tcp_client.go

### 예정된 사항들, 관련 레파지토리
1. Kubernetes , Docker Networking 
2. Kubernetes Flannel 분석 및 나만의 클론 코딩
3. [GoHTTPerf](https://github.com/royroyee/gohttperf) : HTTP Benchmarking tool
4. Kubernetes CNI + SmartNIC
5. 네트워크 심화 개념