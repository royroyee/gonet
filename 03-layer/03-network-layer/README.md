# Layer 3,  Network Layer
**Layer 3 - 네트워크 계층**

03-data-link-layer 에서 배운 데이터 링크 계층으로만은 인터넷이나 다른 네트워크로는 데이터를 전송할 수 없다. (같은 네트워크에 있는 컴퓨터로만 데이터 전송이 가능)
> 네트워크 간의 통신을 가능하게 하는 것이 네트워크 계층의 역할

## NAT(Network Address Translation)
- **사설 IP 주소를 공인 IP 주소로 변환하여 인터넷에 연결하는 기술**
- IP 패킷의 TCP/UDP 포트 숫자와 소스 및 목적지의 IP 주소 등을 재기록 하면서 라우터를 통해 네트워크 트래픽을 주고 받는 기술
- 라우터나 방화벽과 같은 네트워크 장비에서 수행

### NAT 작동 방식
- **사설 IP 주소를 가지고 있는 호스트가 인터넷에 접속할 때** : 해당 호스트의 IP 패킷의 출발지 주소를 공인 IP 주소로 바꿔준다. 
- **외부에서 해당 호스트에 접속할 때** : 공인 IP 주소를 이용하여 라우터나 방화벽이 패킷을 수신한 후에 해당 호스트의 사설 IP 주소로 패킷을 전달해준다.

> 보통 사설네트워크에 속한 여러개의 호스트가 하나의 공인 IP 주소를 사용하여 인터넷을 접속하기 위해서 사용

### 예시
- [리눅스 IP Masquerade](https://github.com/royroyee/gonet/blob/main/linux/README.md#ip-%EB%A7%88%EC%8A%A4%EC%BB%A4%EB%A0%88%EC%9D%B4%EB%93%9Cmasquerade)

---

## 라우터(Router)
- NAT기능을 가진 외부 네트워크와 내부 네트워크를 연결해주는 하나의 장치
- 서로 다른 네트워크 간에 통신을 하려면 라우터가 필요 (L3 스위치가 라우터의 역할을 하기도 한다.)
> 라우터를 통해 네트워크를 분리할 수 있다.

## 라우팅(Routing)
>IP 주소로 목적지를 지정하는 것 뿐만 아니라 **어떤 경로**로 보낼지도 결정해야 한다.
- 목적지 IP 주소까지 어떤 경로로 데이터를 보낼지 결정하는 것을 라우팅 이라고 한다.
- 최단 거리를 찾는 알고리즘을 사용한다.
- 현재 네트워크에서 다른 네트워크로 **최적의 경로를 통해** 데이터를 전송한다.
  - 이 경로 정보가 등록되어 있는 테이블을 **라우팅 테이블** 이라고 한다.

### 라우팅 테이블(Routing Table)
- 라우팅 알고리즘을 통해 만들어지는 테이
- 라우팅 테이블은 라우터가 데이터를 전송할 때, 수신 측 호스트의 IP 주소를 기반으로 경로를 선택하는 데 사용. 
- 라우터는 수신 측 호스트의 IP 주소와 일치하는 목적지 네트워크 주소를 라우팅 테이블에서 찾는다. 
- 라우팅 테이블에는 여러 경로가 있을 수 있으며, 라우터는 이 중에서 최적 경로를 선택한다. 
  - 최적 경로를 선택하는 데에는 여러 가지 요인이 영향을 미칠 수 있으며, 이러한 요인에는 경로의 대역폭, 지연 시간, 비용 등이 포함될 수 있다.
  
## 포워딩(Forwarding)
> 포워드는 데이터 패킷을 수신하여 적절한 출력 포트로 전달하는 과정

### 포워딩 테이블(Forward Table)
- 포워딩 테이블은 라우터가 수신한 패킷을 적절한 출력 인터페이스로 전달하기 위한 정보를 담고 있는 테이블
    - 라우팅 테이블은 라우터가 수신한 데이터 패킷의 최종 목적지를 결정하는 데 사용되는 정보를 담고 있는 테이블

### 레이어2 스위치 vs 라우터
> 스위치의 목적 : 여러 장치를 연결하여 네트워크 생성
> 
> 라우터의 목적 : 서로 다른 네트워크를 연결
1. **계층** : 스위치는 데이터링크 계층에서 MAC 주소를 기반으로, 라우터는 네트워크 계층에서 IP 주소를 기반으로 동작
2. **브로드캐스트 도메인 구분** : 스위치는 불가능, 라우터는 가능
   - `브로드캐스트 도메인` : 브로드캐스트 패킷이 도달할 수 있는 영역
   - 레이어2 스위치 : 스위치 내에 연결된 모든 포트에 데이터를 전송한다. 이러한 동작 방식은 대규모 네트워크에서 브로드캐스트 패킷이 전체 네트워크 대역폭을 차지하고, 네트워크 성능을 떨어뜨리는 원인이 된다.
   - 라우터 : 라우터는 라우팅 테이블을 사용하여 패킷을 전송하는데, 라우팅 테이블은 목적지 주소를 기반으로 전송 유무를 결정하므로 브로드캐스트 도메인을 구분할 수 있다.
3. **불명확한 목적지를 가진 데이터 처리** : 스위치는 모든 포트로 데이터를 전송하는 브로드캐스트를 동작, 라우터는 목적지가 불명확한 데이터는 버린다.
4. **관리자 설정** : 스위치는 필요 없으나, 라우터는 관리자의 설정으로 라우팅 테이블 생성 및 통신을 해야한다.
5. **데이터 전송범위** 
   - 레이어2 스위치 : 동일 네트워크
   - 라우터 : 다른 네트워크
---

## 게이트웨이(Gateway)
- 서로 다른 두 개 이상의 네트워크를 연결하여 통신이 가능하게 해주는 역할
- **다른 네트워크 영역으로 가기 위해 필수적으로 거쳐야 하는 곳** 이며 외부는 다른 프로토콜을 이용하여 통신한다.
  - 다른 네트워크 영역으로 가기 위해 프로토콜 변환 등 다양한 기능이 숨어져있다.
  - 프록시 서버 등과 역할은 유사하나, 프록시 서버는 게이트웨이와는 조금 개념이 다르며, 5~7계층에서 작동한다.
> 소프트웨어 측을 강조할 땐 게이트웨이, 하드웨어적 측면을 강조할 땐 라우터라고도 한다. 즉, 같은 개념이라고 볼 수 있다. 

> 게이트웨이는 패킷이 목적지 서브넷과 다른 서브넷으로 전달될 때 패킷이 거쳐가는 경로라고도 할 수 있다.
---

## IP
> LAN 에서는 MAC 주소만으로도 통신이 가능하지만, 다른 네트워크에는 데이터를 보낼 수 없다. MAC 주소가 아닌 네트워크를 식별할 수 있는 다른 주소
- IP 주소로 어떤 컴퓨터로 데이터를 보낼 지 목적지를 설정한다.
- 네트워크 계층의 대표적인 프로토콜
- 네트워크 계층에서 캡슐화를 할 때, 데이터에 `IP 헤더` 를 붙인다.
  - 이를 `IP 패킷` 이라고 한다.
  - IP 헤더의 주요 요소 : 출발지/목적지 IP 주소


## IP 주소
- IP 주소는 인터넷 서비스 제공자(ISP)에게 받을 수 있다.
- 32비트, 구분하기 쉽도록 10진수로 표시
- IP 버전 : IPv4 / IPv6
- 유형 : 공인 IP 주소 / 사설 IP 주소
- 네트워크 ID / 호스트 ID 두 가지 정보가 합쳐저서 IP 주소를 이룸
  - `공인 IP 주소` : 인터넷에 직접 연결되는 컴퓨터나 라우터에 할당
  - `사설 IP 주소` : 회사나 가정의 LAN에 있는 컴퓨터에 할당
  > 공인 IP 주소를 절약하기 위해 라우터에만 할당하고 LAN 안에 있는 컴퓨터에는 LAN의 네트워크 관리자가 자유롭게 사설 IP 주소를 할당하거나 라우터의 DHCP 기능을 사용하여 주소를 자동으로 할당한다.
  - `네트워크 ID` : 어떤 네트워크 인지를 나타낸다.
  - `호스트 ID` : 해당 네트워크의 어느 컴퓨터인지 나타낸다.
---

## DHCP
- IP 주소와 같은 TCP/IP 네트워크 구성 정보를 자동으로 할당하는 프로토콜
- **IP 주소를 자동으로 할당하기 위한 서비스**
> IP 주소, 서브넷 마스크, 기본 게이트웨이, DNS 서버 등을 자동으로 할당하여 네트워크 설정을 간편하게 한다.
- 단점으로는, DHCP 서버에 IP 할당을 의존하므로, DHCP 서버가 다운되면 IP 할당이 제대로 이루어지지 않을 수 있다.

### IPv4
- 32비트 , 약 43억개의 IP 주소 생성 가능 (초기에 43억개 정도로도 충분하다고 생각했었으나, 현재는 부족하여 IPv6 등의 대안을 마련)

### IPv6 
- 128비트, 약 340간 개의 IP 주소 생성 가능
- 급속도로 고갈된 IP 주소의 대책으로 생겨난 프로토콜

## IP 주소 클래스
- 네트워크 ID 와 호스트 ID 의 크기는 클래스마다 다르다.
- A~E 클래스가 있지만 C 까지만 서술한다.
- 네트워크 ID 비트가 적을 수록 큰 네트워크

### A 클래스
- 대규모 네트워크 주소
- `네트워크 ID` : 앞 8비트
- `호스트 ID` : 뒤 24비트
- 사설 IP 주소 범위 : 10.0.0.0 ~ 10.255.255.255

### B 클래스
- 중형 네트워크 주소
- `네트워크 ID` : 앞 16비트
- `호스트 ID` : 뒤 16비트
- 사설 IP 주소 범위 : 172.16.0.0 ~ 172.31.255.255

### C 클래스
- 소규모 네트워크 주소
- `네트워크 ID` : 앞 24비트
- `소규모 ID` : 뒤 24비트
- 사설 IP 주소 범위 : 192.168.0.0 ~ 192.168.255.255

## 데이터를 전송하는 주소

### 유니캐스트(Unicast)
- 단 한 곳으로 데이터를 전송하는 것
- 유니캐스트에 이용하는 IP 주소가 유니캐스트 IP 주소이다.
- 목적지 호스트의 유니캐스트 IP 주소를 IP 헤더의 목적지 IP 주소로 지정
- IP 헤더
  - 목적지 IP : 유니캐스트
  - 출발지 IP : 유니캐스트
> 단, 완전히 같은 데이터를 복수의 주소로 보낼 땐 효율적으로 전송하기 위해 브로드캐스트 또는 멀티캐스트를 사용

### 멀티캐스트(Multicast)
- 같은 애플리케이션이 동작하는 등 특정 그룹에 포함되는 호스트에 완전히 똑같은 데이터를 전송하는 것
  - 즉, 멀티캐스트 그룹에 포함되는 호스트는 반드시 같은 네트워크라고 할 수는 없음
- IP 헤더의 목적지 IP 주소에 멀티캐스트 IP 주소를 지정
- IP 헤더
  - 목적지 IP : 멀티캐스트
  - 출발지 IP : 유니캐스트

### 브로드캐스트(Broadcast)
- 같은 네트워크 상의 모든 호스트에 완전히 똑같은 데이터를 전송하는 것
- IP 헤더의 목적지 IP 주소에 브로드캐스트 IP 주소를 지정
- IP 헤더
  - 목적지 IP : 브로드캐스트
  - 출발지 IP : 유니캐스트

## 네트워크 & 브로드캐스트 주소
- 컴퓨터나 라우터가 자신의 IP 주소로 사용하면 안 되는 주소

### 네트워크 주소
- 호스트 ID가 10진수로 0인 주소
- 전체 네트워크에서 **작은 네트워크** 를 식별하는 데 사용되고 그 네트워크 전체를 대표하는 주소
  - 한 마디로 **전체 네트워크의 대표 주소**
  - ex) `192.168.2.0` : 192.168.2.0 ~ 192.168.2.255 를 대표하는 주소

---

### 브로드캐스트 주소
- 네트워크에 있는 컴퓨터나 장비 모두에게 한 번에 데이터를 전송하는 데 사용되는 전용 IP 
> 브로드캐스팅은 송신 호스트가 전송한 데이터가 네트워크에 연결된 모든 호스트에 전송되는 방식을 의미한다.
  - 호스트 ID 가 10진수로 255인 주소
    - ex) `192.168.1.255` : 이 주소로 데이터를 전송하면 네트워크 안에 있는 모든 컴퓨터가 데이터를 받게 된다. (브로드캐스팅)

### 브로드캐스트 도메인
브로드캐스트 도메인이란 LAN 상에서 어떤 단말이 브로드캐스트 패킷을 송출할 때 이 패킷에 대해 네트워크에서 영향을 받는 영역 또는 그 패킷을 수신할 수 있는 단말들의 집합

### 브로드캐스트 도메인 분할 단위
#### VLAN
VLAN 에 대한 자세한 정보 : [링크]() 
- 하나의 스위치에서 여러개의 LAN을 나누어서 사용이 가능
- 이렇게 가상으로 나누어진 VLAN은 하나의 논리적인 브로드캐스트 도메인이 된다.
#### 라우터
라우터 또는 L3 스위치
- 브로드캐스트 도메인은 LAN 상의 모든 단말에 패킷을 전송하게 되므로 L3 를 넘어갈 수 없다. 때문에 라우터가 브로드캐스트 도메인을 나누어줄 수 있다.

### 브로드캐스트 도메인을 나누는 이유
- 브로드캐스트는 연결된 단말 모두에 패킷을 보내게 되는데, 때문에 대역폭을 낭비할 수 있다.
- 브로드캐스트 패킷을 받지 않아도 되는 단말에서 패킷을 받을 수 있는 낭비 때문에,네트워크 설계시 적절히 나누어주어야 한다.
---

## 서브넷
- 네트워크를 `클래스` 로만 나누는 것은 매우 비효율적이다.
  - 예를들어, A 클래스에서 호스트 ID 는 24비트에 달한다. 이 많은 네트워크에 브로드캐스트 패킷을 전송하면 매우 비효율적 
  - 또는 기업 마다 네트워크 ID, 호스트 ID 의 수요가 크게 다를 텐데, 클래스만으로 나누는 것은 비효율적
> 때문에, 네트워크를 작은 네트워크로 분할하는 것을 서브넷팅 이라고 한다. 이렇게 분할된 네트워크를 서브넷이라고 한다.

- `네트워크 ID`, `호스트 ID` 을 구분하는 방법이다.
### 서브넷팅
![subnet mask.png](img%2Fsubnet%20mask.png)
- 네트워크 ID 길이를 증가 시키고 호스트 ID 길이를 감소시킨다.
- 때문에, 서브네트워크로 나누어진다.

## 서브넷 마스크
> 네트워크 ID와 호스트 ID 를 식별하기 위한 값
- 위에서 `서브넷 ID` 가 추가되었는데, 이렇게 서브넷팅을 해버리면 어디까지가 네트워크 ID 인지 등의 판별이 어려워진다.
- 이런 문제를 해결하기 위해 **서브넷 마스크** 를 이용한다.
- 정확한 구분 방법들은 추후 작성 예정
> 프리픽스 표기법 : IP 주소에서 네트워크 부분을 나타내는 방법 중 하나. IP 주소와 함께 슬래시(/)로 구분되며 네트워크 ID의 비트 수를 나타낸다
> 
> 예를 들어, 192.168.1.0/24라는 IP 주소는 24비트가 네트워크 ID, 나머지 8비트는 호스트ID 를 나타낸다.

### Example 서브네팅 적용 전
![서브네팅 적용 전.png](img%2F%EC%84%9C%EB%B8%8C%EB%84%A4%ED%8C%85%20%EC%A0%81%EC%9A%A9%20%EC%A0%84.png)
- B 클래스(network id : 16비트)
- 2^16 개의 호스트를 가진 한 개의 네트워크
- 한 개의 연결을 통하여 전체 네트워크가 인터넷 내의 한 개의 라우터에 연결
- `/16` 으로 표기

### 서브네팅 적용 후
![서브네팅 적용 후.png](img%2F%EC%84%9C%EB%B8%8C%EB%84%A4%ED%8C%85%20%EC%A0%81%EC%9A%A9%20%ED%9B%84.png)
- 위의 예제의 같은 네트워크를 `서브네팅`을 이용한 사례
- 전체 네트워크는 여전히 같은 라우터를 통해 인터넷에 연결
- 단, 네트워크는 사설 라우터를 사용하여 네트워크를 네 개의 서브네트워크로 나누고 있음
- 각 서브네트워크는 거의 2^14 개의 호스트를 가질 수 있음


## TTL
> "Time to Live" 의 약자로 네트워크 프로토콜에서 패킷이 유효한 시간을 나타내는 값
- 패킷의 출발지와 목적지 사이의 경로를 따라 이동하는 동안 사용
- 패킷은 출발지에서 TTL 값을 가지고 출발하며, 이 값은 경로상의 각 라우터에서 1씩 감소
- 만약 어떤 라우터에서 TTL 값이 0이 되면 라우터가 감지하여 해당 패킷을 버리고 `ICMP 메시지`를 발신지 host에 보내게 된다.
    - TTL 값을 사용하지 않았다면 **패킷이 무한히 라우터 사이를 전달**하면서 네트워크 전체가 다운될 수도 있다.
- 보통 1~255 정수값으로 설정. 
> 값이 낮을수록 패킷이 목적지에 도달하기 전에 폐기할 가능성이 높아지므로 TTL 값은 데이터의 중요도와 전송 거리등을 고려하여 결정해야 한다.

---

## 주소 변환
인터넷은 라우터와 같은 네트워크 간 연결장치에 의해 상호 연결된 네트워크로 구성된다.

발신지 호스트가 보낸 패킷은 목적지 호스트에 도착하기 전에, 여러 개의 다른 네트워크를 지나갈 수 있다.

호스트와 라우터는 Network 레벨에서 자신의 논리 주소(Logical Address)로 인식된다.
- 논리주소 : 소프트웨어 상에서 구현되므로 논리주소라 하며, 전 세계에서 유일해야 한다.
- TCP/IP 프로토콜에선 IP 주소가 `논리 주소`에 해당한다.

패킷이 논리적인 네트워크만 통과할까? 패킷들은 호스트, 라우터에 도달하기 위해 **물리적인 네트워크도 통과해야 한다.**
- 물리적인 레벨에서 호스트와 라우터들은 `물리 주소` 에 의해 인식된다. (MAC 주소, 호스트나 라우터 내부에 설치된 NIC 에 들어있다.)

> 즉, 물리 주소와 논리 주소는 서로 다른 식별자이고 둘 다 필요하다. 

## ARP(Address Resolution Protocol) 프로토콜 
목적지 컴퓨터의 IP 주소를 이용하여 MAC 주소를 찾기 위한 프로토콜 (IP 주소를 MAC 주소와 매칭 시키기 위한 프로토콜)

예를 들어 어떤 호스트나 라우터가 다른 호스트나 라우터에 보낼 `IP 데이터그램`을 가지고 있다면?
- 송신자는 수신자의 논리 주소인 `IP 주소`를 가지고 있다는 의미와 동일하다.

그러나, `IP 데이터그램`은 **물리적인 네트워크를 통과하기 위해 프레임(이더넷 프레임) 내에 캡슐화** 되어야 한다.
- 즉 송신자는 수신자의 물리 주소(MAC) 도 알아야 한다. 그러나 MAC 주소를 모른다면 어떡해야 할까? 이때 사용되는 것이 `ARP 프로토콜`이다.

### 과정
호스트나 라우터가 같은 네트워크 상에 있는 다른 호스트나 라우터의 물리 주소를 필요로 할 때(즉, 송신자가 수신자의 물리주소(MAC)를 모를 때)
1. ARP 질의 패킷을 보낸다. 
    - 이 질의 패킷에는 송신자의 MAC 주소, IP 주소 , 수신자의 IP 주소를 포함한다.
2. 송신자는 수신자의 MAC 주소를 모르므로 네트워크 상에서 브로드캐스트 한다.
3. 라우터 상의 모든 호스트와 라우터는 이 ARP 질의 패킷을 수신하고 처리한다.
   - 단, 해당 **IP 주소를 가지고 있는 수신자만**이 IP 주소를 인식하고 **ARP 응답 패킷을 돌려보낸다.**


> ARP 요청 : 출발지 컴퓨터가 이더넷 프레임을 전송하려면 목적지 컴퓨터의 MAC 주소를 지정해야 하는데, 이를 모를 때 MAC 주소를 알아내기 위해 네트워크에 브로드캐스트를 하는 것
- 즉, APR Request 는 연결된 네트워크 장비와 컴퓨터에 모두 전달된다.
> APR 응답 : 위의 요청에 대해 지정된 IP 주소를 갖고 있는 컴퓨터가 MAC 주소를 응답으로 보내는 것
- IP 주소와 MAC 주소를 일대일 매칭하여 LAN(Layer 2)에서 목적지를 제대로 찾아갈 수 있도록 돕는다.
- **LAN 의 범위를 ARP 가 닿는 모든 범위**라고 칭하기도 한다.

### ARP Table
- IP 주소와 MAC 주소를 일대일 매칭시킨 정보를 정리해둔 Table
- L2 네트워크에선, ARP 테이블을 수집하기 위해 Broadcast 를 이용한다.(ARP Request)





## ICMP(Internet Control Message Protocol)
> 인터넷 프로토콜(IP)의 일부로, 네트워크 상황에 대한 메시지를 전송하는 역할을 한다.

- 네트워크에서 데이터를 전송하면서 오류나 문제가 발생한 경우 ICMP는 오류 메시지를 생성하여 이를 송신자에게 전달하고, 이에 따라 해당 문제를 해결할 수 있도록 돕는다.
- IP 패킷을 전달하는 과정에서 발생하는 여러가지 문제들을 알리는데 사용


- `Destination Unreachable` : 목적지 호스트에 도달할 수 없음을 알리는 메시지
- `Time Exceeded` : IP 패킷이 (TTL) 값에 도달하지 못하여 제거됨을 알리는 메시지
- `Redirect` : 패킷을 다른 라우터로 보내야 함을 알리는 메시지
- `Echo Request` : 네트워크 상태를 확인하는 메시지
- `Echo Reply` : Echo Request에 대한 응답 메시지
> 따라서 ICMP는 인터넷 프로토콜의 핵심 요소 중 하나이며, 네트워크 통신을 유지하고 문제를 해결하는 데 매우 중요한 역할을 한다.

자세한 내용 : [ICMP 패킷 분석]()