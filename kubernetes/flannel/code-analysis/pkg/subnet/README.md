# Subnet
**subnet package의 주요 코드들을 분석**

## 사전지식
[서브넷이란?](https://github.com/royroyee/gonet/tree/main/03-layer/03-network-layer#%EC%84%9C%EB%B8%8C%EB%84%B7)
> Flannel은 가상의 네트워크 상에서 각 노드의 IP 대역을 분해하기 위해 서브넷을 사용한다.
- 서브넷은 IP 대역와 CIDR 프리픽스로 구성되며, 각 노드에서 사용할 수 있는 IP 대역을 지정한다.



---

## 주요 코드


### 1. config.go
> Flannel의 구성 파일을 검증하고, IPv4 및 IPv6 네트워크를 관리하는 데 사용
- `Config` struct : Flannel 구성을 저장, IPv4, IPv6 둘 다 가능, 네트워크와 백엔드 등의 값이 할당
- `CheckNetworkConfig()` : `Config` 구조체의 유효성을 검사하는 함수
- `GetFlannelNetwork()` : IPv4 네트워크에 대한 정보를 가져오는 함수
- `parseBackendType()` : 구성 파일에서 백엔드 유형을 검색하고 반환하는 함수 
   
### 2. subnet.go
> Flannel의 서브넷 관리와 관련된 기능들을 구현

- `MakeSubnetKey()` : 서브넷 키 생성
- `WriteSubnetFile()` : 서브넷 파일 작성
