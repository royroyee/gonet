# Flannel Code Analysis
**Flannel 코드를 분석하며 네트워크 프로그래밍, Kubernetes에 대해 더 깊게 학습하기**

#### [원본 코드](https://github.com/flannel-io/flannel/tree/master/pkg)

> powershell pkg, _windows.go, amd64.go 등의 운영체제에 따라 다른 코드들은 분석을 제외합니다. (Linux 전용 코드만 분석)
## Package
- powershell, version pkg 를 제외한 총 10개의 pkg 분석
- IPv6 도 현재는 제외함(추후 필요 시 작성 예정)

### backend
- Flannel의 백엔드 구현에 사용되는 코드 
- Flannel이 호스팅하는 네트워크를 설정하는 데 필요한 인프라를 제공
- 논리적인 네트워크를 만드는 데 사용

### ip
- IP 주소와 관련된 코드를 제공
- IP 주소를 만들고 변환하고, 비교하고 쿼리하는 데 사용된다.
- Flannel 네트워크의 설정 및 관리를 도와준다.

### ipmath
- IP 주소의 계산을 담당
- IP 주소와 CIDR 마스크 간에 변환을 수행

### iptables
- iptabels 규칙을 설정 및 관리
- Flannel 은 이 패키지를 사용해 호스트 시스템에서 패킷 전달을 가능하게 하고, NAT 패키지를 구성한다.

### mac
- MAC 주소를 생성 및 변환, 비교하는 데 사용
- Flannel 에서 MAC 주소는 호스트 간의 고유 식별자로 사용된다.

### ns
- Linux 네임스페이스를 구성 및 관리하는 데 사용
- Flannel은 호스트의 네임스페이스를 구성하고 이를 통해 논리적인 네트워크를 호스트 간에 공유한다.

### retry
- 네트워크 연결을 다시 시도하는 코드를 제공
- Flannel은 이 패키지를 사용하여 네트워크 연결 문제를 해결하고, 더 나은 호스트 간 통신을 가능하게 한다.

### routing
- Flannel의 논리적인 네트워크에 대한 라우팅을 구성 및 관리
- 호스트 간의 패킷 전달을 가능하게 하며 서브넷의 구성과 호스트 간의 라우팅 테이블을 설정

### subnet 
- 서브넷을 구성 및 관리
- IP 대역을 관리하여 모든 노드에서 일관성 있는 IP 대역 사용을 가능하게 한다.


