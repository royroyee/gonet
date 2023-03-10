# Linux Network


## iptables
**리눅스에서 네트워크 패킷을 필터링하고, 포트포워딩, NAT, 패킷 로깅 등 다양한 기능을 제공하는 프로그램**
- 네트워크 패킷의 전반적인 흐름을 `체인` 이라는 단위로 관리
  - `체인` : 패킷의 흐름을 처리하기 위한 규칙을 가지고 있으며, **패킷이 체인을 따라 흐르면서 각각의 규칙에 따라 처리가 이루어진다.**
  - `타겟` : 패킷이 규칙과 일치할 때 동작을 취하는 타겟을 지원한다.
> 주의! :  레이어3 NAT 기능은 IP 주소 변환만 하는 것이 맞지만, 리눅스 iptables 에서는 NAT 기능 뿐만 아니라 다양한 규칙들을 포함해서 처리한다.


### iptables Table
**정책을 적용할 네트워크 패킷의 유형을 결정**

- `Filter` : 기본 테이블로 패킷을 거부하거나 허용하는 등의 기본적인 패킷 필터링 작업을 수행
- `NAT` : 네트워크 주소 변환 작업에 사용되는 테이블. 포트포워딩, IP 마스커레이드, 로드밸런싱 등의 작업을 처리
- `Mangle` : 라우팅, TOS(데이터그램의 우선순위), TTL(생존 시간) 등의 네트워크 패킷 처리 작업을 수행(자주 사용되진 않음)
  - TOS (Type of Service) : 패킷의 TOS(패킷 전송의 우선순위)를 변경하고자 하는 경우에 사용
  - TTL(Time to Live) : 패킷의 TTL을 변경하고자 하는 경우에 사용
  - MARK: 패킷에 특별한 표시 값을 설정하고자 하는 경우에 사용
- `Raw` : NAT 테이블 전에 동작, RAW 체인과 PREROUTING, OUTPUT 체인에 대해 룰을 적용. 일반적으로 인바운드 패킷의 제어와 관련된 룰을 구성하는 데 사용
- `Security` : SELinux 보안 기능과 같은 리눅스 보안 기술을 적용하는 데 사용

### iptables Chain
패킷이나 커널 데이터를 처리하는 일련의 규칙 
> 체인은 패킷의 경로를 결정하고 이러한 패킷 경로를 따라 규칙이 적용되어 패킷이 수락되거나 거부되는 데 사용
- 패킷이 이동하는 경로를 나타낸다.
  
(그림)
- `INPUT` 체인 : 호스트(시스템)로 들어오는 패킷을 처리, 패킷이 로컬 시스템에서 처리되는 지 여부를 결정한다.
  - 방화벽으로 들어오는 패킷, 목적지 주소가 방화벽 IP인 모든 패킷을 의미한다.

- `OUTPUT` 체인 : 호스트(시스템)에서 나가는 패킷을 처리, 패킷이 시스템에서 정상적으로 처리되는지 여부를 결정
- `FORWARD` 체인 : 패킷이 라우터를 통과하는 경우에 사용되는 체인, 패킷이 라우터에서 전달되는지 여부를 결정
- `PREROUTING` 체인 : 패킷이 라우터로 들어오기 전에 사용되는 체인, 패킷의 **도착지 IP 주소**를 변경하는 NAT 처리를 수행(DNAT)
- `POSTROUTING` 체인 : 패킷이 라우터를 통과한 후에 사용되는 체인, 패킷의 **출발지 IP 주소**를 변경하는 NAT 처리를 수행(SNAT)
> `INPUT, OUTPUT, FORWARD` : 패킷 필터링을 위해 사용되는 체인

> `PREROUTING`, `POSTROUTING` NAT 처리를 위해 사용되는 체인 

- 각 체인에 대해 규칙을 추가하거나 삭제할 수 있다.

### iptables Target(Action)
> 패킷 규칙이 일치할 때 취하는 동작을 지정
- `ACCEPT` : 패킷을 받아들인다.
- `DROP` : 패킷을 버린다. (패킷이 전송된 적이 없던 것처럼), 아무런 메세지도 전송하지 않는다.
- `REJECT` : 패킷을 버리고 이와 동시에 사용자에게 오류 메세지를 보낸다(ICMP 타입 형태의 메시지).
- `LOG` : 패킷을 로그에 남긴다.
- `SNAT` : 출발지 IP 주소를 변경
- `DNAT` : 목적지 IP 주소를 변경
- `MASQUERADE` : 로컬 네트워크에서 외부 네트워크로 나가는 패킷에 대해 NAT를 수행
- `REDIRECT` : 패킷을 로컬 포트로 리다이렉트
- `MARK` : 패킷에 특정한 마크를 설정
- `TEE` : 패킷을 복사하여 다른 호스트로 전송
- `SNTP` : 패킷의 시간 정보를 변경
- `RETURN` : 패킷을 처리하지 않고 즉시 호출자로 돌려 보내는 역할. 일종의 조건문으로 사용되며, 패킷이 매칭되었을 때 특정 규칙을 즉시 종료시키고 다음 규칙으로 넘어가지 않도록 한다.


### iptables Option
- `-A` : [APPEND]  규칙을 추가하는 옵션. 새로운 규칙을 마지막에 삽입
- `-I` : [INSERT]  규칙을 삽입하는 옵션. 새로운 규칙을 명시된 위치에 삽입
- `-D` : [DELETE]  규칙을 삭제하는 옵션. 삭제하고자 하는 규칙 번호를 지정하여 사용합.
- `-R` : [REPLACE]  기존의 규칙을 교체하는 옵션. 교체하고자 하는 규칙 번호와 새로운 규칙을 지정하여 사용.
- `-F` : [FLUSH]  특정 테이블의 모든 규칙을 삭제하는 옵션입니다.
- `-P` : [POLICY]  체인에 대한 기본 규칙을 설정하는 옵션입니다. 보통 `INPUT`, `FORWARD`, `OUTPUT` 체인에 대한 기본 규칙을 설정할 때 사용.
- `-L` : [LIST]  정책을 나열하는 옵션. 현재 설정된 규칙들을 보여준다.

### iptables Match
- `-s/--source` : 출발지 주소 매칭 옵션. 해당 옵션 뒤에 IP 주소나 도메인 이름, 서브넷 등을 입력하여 출발지 주소가 일치하는 패킷만 매칭시킬 수 있다.

- `-d/--destination` : 목적지 주소 매칭 옵션. 해당 옵션 뒤에 IP 주소나 도메인 이름, 서브넷 등을 입력하여 목적지 주소가 일치하는 패킷만 매칭시킬 수 있다.

- `-p/--protocol` : 프로토콜 매칭 옵션. 해당 옵션 뒤에 TCP, UDP, ICMP 등 프로토콜 이름을 입력하여 해당 프로토콜을 사용하는 패킷만 매칭시킬 수 있다.

- `-i/--in-interface` : 입력 인터페이스 매칭 옵션. 해당 옵션 뒤에 네트워크 인터페이스 이름을 입력하여 해당 인터페이스로 들어오는 패킷만 매칭시킬 수 있다.

- `-o/--out-interface` : 출력 인터페이스 매칭 옵션. 해당 옵션 뒤에 네트워크 인터페이스 이름을 입력하여 해당 인터페이스로 나가는 패킷만 매칭시킬 수 있다.

- `-j/--jump` : 매치되는 패킷을 처리할 액션을 지정하는 옵션. 해당 옵션 뒤에 ACCEPT, DROP, REJECT 등의 액션을 입력하여 매칭된 패킷을 어떻게 처리할지 결정할 수 있다.

### iptables Match Extension
> 트래픽 필터링을 위해 패킷의 특정 속성을 매치시켜 필터링 할 수 있게 해주는 확장 기술

- `-m state` : 패킷의 커넥션 상태를 매치시켜 필터링. --state 옵션과 함께 사용

- `-m tcp` : TCP 패킷의 속성을 매치시켜 필터링. 예를 들어, --dport 옵션과 함께 사용하면 목적지 포트번호로 필터링할 수 있다.

- `-m udp` : UDP 패킷의 속성을 매치시켜 필터링. TCP와 마찬가지로 --dport 옵션과 함께 사용하면 목적지 포트번호로 필터링할 수 있다.

- `-m limit` : 패킷의 속도를 제한. 초당 허용되는 패킷 수를 지정할 수 있다.

- `-m string` : 패킷의 데이터(페이로드)에서 특정 문자열을 찾아 필터링.

- `-m mac` : MAC 주소를 기반으로 패킷을 필터링.

- `-m iprange` : 특정 IP 주소 범위를 필터링.

- `-m comment` : 패킷에 주석을 추가하거나 패킷 주석을 기반으로 필터링.

- `-m mark` : 특정 마크값과 일치하는 IP 패킷을 선택하는 데 사용


### Example
```
iptables -A INPUT -p tcp --dport 80 -j ACCEPT
```
- `INPUT` 체인에 tcp 프로토콜을 사용하여 80번 포트로 들어오는 모든 패킷을 `ACCEPT` 하겠다는 것을 설정하는 명령어
  - `--dport 80` 대상 포트가 80번인 패킷을 선택(destination port)
  
```
iptables -A PREROUTING -t nat -i eth0 -p tcp --dport 80 -j DNAT --to-destination 192.168.1.100:8080
iptables -A FORWARD -p tcp -d 192.168.1.100 --dport 8080 -j ACCEPT
```
1. eth0 인터페이스를 통해 들어오는 tcp 80포트 패킷을 192.168.1.100 서버의 8080 포트로 전달 (DNAT)
2. 192.168.1.100 서버로 전달되는 패킷 중에 8080 포트로 전달되는 tcp 패킷을 ACCEPT

```go
// flannel iptables.go
iptables -t nat -A FLANNEL-POSTRTG -m mark --mark 0x4000/0x4000 -m comment --comment "flanneld masq" -j RETURN
```
- nat 테이블에서 FLANNEL-POSTRTG 체인에 새로운 규칙을 추가하는 명령어
- 특정 마크 값(0x4000/0x4000)을 패킷을 탐지하여 flanneld masq 라는 주석을 붙인 후 RETURN 옵션으로 반환한다.
  - 즉, 패킷이 특정 마크와 일치하는 지 확인하고 만약 패킷의 마크가 일치하면 주석을 달고 RETURN 액션을 수행하여 다음 규칙으로 이동하지 않도록 한다.

> Plus : iptables는 정책을 적용할 때, 먼저 적용된 정책부터 차례대로 처리하므로 규칙의 순서는 매우 중요하다.
## IP Masquerade
> Masquerade : 가면
- 리눅스의 [NAT](https://github.com/royroyee/gonet/tree/main/03-layer/03-network-layer#natnetwork-address-translation) 기능


- IP 마스커레이드도 NAT 처럼 사설 IP 주소를 공인 IP 주소로 변환하여 사설 IP 주소만 배정받은 호스트가 외부 통신이 가능하도록 한다.


- NAT 과 다른 점의 핵심은 임의의 포트 번호를 배정한다는 것이다.


- 사설 IP 주소를 가지고 있는 호스트가 인터넷에 접속할 때 
  1. 패킷의 출발지 IP 주소는 호스트의 사설 IP 주소로 설정
  2. 패킷은 라우터로 전송되고, 라우터는 패킷의 출발지 IP 주소를 **마스커레이드 규칙**에 따라 공인 IP 주소와 **임의의 포트 번호를 가진** 새로운 패킷으로 변경 


- 외부에서 해당 호스트에 접속할 때 : 접속할 때 설정한 임의로 배정 받은 포트 번호를 통해 패킷을 보낸 사설 IP 로 패킷을 받게 한다.

### NAT vs IP Masquerade
둘 다 사설 IP 주소를 사용하는 호스트가 공인 IP 주소를 사용하는 인터넷에 접속할 때 호스트의 사설 IP 주소를 공인 IP 주소로 변환하여 인터넷에 접속할 수 있게 해주는 역할을 한다.
- `NAT` : 호스트의 사설 IP 주소를 변환할 때, **공인 IP 주소와 포트 번호를 매핑하여 변환**한다.
  - 때문에 **NAT를 사용하는 호스트는 인터넷에 접속할 때 고정된 포트번호를 사용**

- `IP 마스커레이드 ` : 호스트의 사설 IP 주소를 변환할 때, 공인 IP 주소만 사용하고 **포트 번호는 무작위로 할당**
  - 때문에 IP 마스커레이드를 사용하는 **호스트는 인터넷에 접속할 때 매번 다른 포트를 사용**

> IP 마스커레이드는 호스트의 사설 IP 주소를 인터넷에 노출시키지 않고, 매번 다른 포트 번호를 사용하여 보안성을 높이는 기술
> 

### 실제 사용 예제 : Kubernetes Network(Flannel)
> Kubernetes 의 Network Plugin 종류 중 하나인 Flannel은 Pod의 IP 주소를 알 수 없게 만들기 위해 IP 마스커레이드를 사용한다. 
> 
> 이를 통해 overlay network의 IP 주소 공간이 호스트 네트워크의 IP 주소 공간과 겹쳐도 Pod IP 주소의 충돌을 방지할 수 있다.


---

## Netfilter
> Linux 커널에서 제공하는 네트워크 패킷 필터링 프레임워크
- 패킷의 흐름을 분석하고, 특정 규칙에 따라 패킷을 처리하거나 블록한다.
- `iptables`와 같은 다양한 도구와 함께 사용
  - `ìptables` 는 위에서 언급한 것처럼, `netfliter` 의 규칙을 설정 및 관리하기 위한 명령어 도구이다.
- 다양한 기능을 제공한다.
  - NAT 
  - Packet filtering : 특정 패킷을 차단 또는 허용하는 기능. 서버의 접근제어 또는 방화벽기능 구현 가능
  - Packet mangling : 필요시 패킷 헤더의 값을 변경

### Netfilter Hook

- 기본적으로 5개의 훅 존재



## Bridge
**두 개 이상의 네트워크를 연결하기 위한 장치**
- 보통 물리적으로 분리된 두 개 이상의 네트워크를 하나로 연결하는 역할을 한다.
  - 예시: 서로 다른 두 개의 LAN 카드(NIC)가 있다고 가정 시 브릿지로 연결하면 하나의 LAN 카드로 연결한 것과 같은 효과를 낼 수 있다.
>브릿지를 통해 연결된 두 개 이상의 네트워크는 마치 하나의 네트워크처럼 동작하며, 서로 통신할 수 있다.
- 마치 브릿지에 연결된 모든 장치들이 같은 IP 대역에 속하는 것처럼 동작한다.
- 브릿지는 리눅스에서 나온 가상적인 개념이다.
- 도커에서도 이 브릿지를 생성하여 컨테이너를 연결한다.
- 브릿지에도 마찬가지로 MAC 주소, IP 주소가 할당되며 **가상의 스위치 역할**을 한다.
- 브릿지에 연결된 네트워크 인터페이스 또한 각자의 고유한 MAC 주소와 IP 주소를 가진다. 

## veth
**가상 이더넷 네트워크 인터페이스**
- 일반적으로 가상 머신, 컨테이너 내에서 사용
- 도커 컨테이너 내부에서는 `eth0` 인터페이스가 존재하지만, 호스트 컴퓨터에는 `eth0` 인터페이스가 없다.
  - 따라서, `veth` 인터페이스를 사용하여 도커 컨테이너의 eth0 인터페이스를 호스트 컴퓨터와 연결한다.

## eth0
**물리적인 네트워크 인터페이스**
- 컴퓨터의 물리적인 네트워크 카드와 연결되어 있다.
- `eth` : Ethernet 인터페이스를 의미
- `0` : 첫 번재 Ethernet 인터페이스를 의미



## 리눅스 네트워크 구성 방법

### 1. 네트워크 인터페이스 구성 확인
```
ifconfig -a (-a: 비활성화된 인터페이스도 보이기 위해 all 옵션 추가)
ens3: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.25.235.132  netmask 255.255.255.0  broadcast 172.25.235.255
        inet6 fe80::f816:3eff:fefd:81b4  prefixlen 64  scopeid 0x20<link>
        ether fa:16:3e:fd:81:b4  txqueuelen 1000  (Ethernet)
        RX packets 175386  bytes 191689004 (191.6 MB)
        RX errors 0  dropped 291  overruns 0  frame 0
        TX packets 68233  bytes 6685069 (6.6 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 9616  bytes 2823080 (2.8 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 9616  bytes 2823080 (2.8 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```
#### 구성요소
- `ens3` : 네트워크 인터페이스 명
- `flags` : 인터페이스 상태 (UP, BROADCAST, RUNNING, MULTICAST)
- `mtu` : 인터페이스 최대 전송 단위
- `inet` : 인터페이스 IPv4 주소
- `netmask` : 인터페이스 IPv4 netmask. (IPv6 는 넷마스크 사용X)
- `broadcast` : 인터페이스 broadcast 주소
- `inet6` : 인터페이스 IPv6 주소
- `prefixlen` : IP 주소에서 netmask bit
- `scopeid` : IPv6 범위
- `ether`: 인터페이스 MAC 주소
- `RX packets` : 받은 패킷 정보(bytes, errors, dropped, overruns, fram)
- `TX packetss` : 보낸 패킷 정보(bytes, erros, dropped, overruns, carrier, collisions)

### 2. Routing

#### Routing Table 확인하기
```
ubuntu@ubuntu:~$ route -n

Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.25.235.254  0.0.0.0         UG    100    0        0 ens3
169.254.169.254 172.25.235.91   255.255.255.255 UGH   100    0        0 ens3
172.25.235.0    0.0.0.0         255.255.255.0   U     0      0        0 ens3

```
- `Destination` : 목적지 네트워크
- `Gateway` : 목적지로 가기위한 게이트웨이 주소
- `Genmask` : 목적지 네트워크의 넷마스크 주소 (목적지 IP 와 Genmask를 AND 연산한 결과가 목적지 네트워크)
- `Flags` : 해당 경로에 대한 정보를 알려주는 기호
  - U : up
  - H : 목적지 Host
  - G : 게이트웨이 사용
- `Metric` : 목적지 네트워크 까지의 거리
- `Ref` : 경로를 참조한 횟수
- `Use` : 경로를 탐색한 횟수
- `IFace` : 네트워크 인터페이스