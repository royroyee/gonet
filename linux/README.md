# Linux Network


## iptables
**리눅스에서 네트워크 패킷을 필터링하고, 포트포워딩, NAT, 패킷 로깅 등 다양한 기능을 제공하는 프로그램**
- 네트워크 패킷의 전반적인 흐름을 `체인` 이라는 단위로 관리
  - `체인` : 패킷의 흐름을 처리하기 위한 규칙을 가지고 있으며, 패킷이 체인을 따라 흐르면서 각각의 규칙에 따라 처리가 이루어진다.

### iptables 체인
1. INPUT 체인 : 호스트로 들어오는 패킷을 처리
2. OUTPUT 체인 : 호스트에서 나가는 패킷을 처리
3. FORWARD 체인 : 호스트를 통과하는 패킷을 처리
4. PREROUTING 체인 : 호스트로 들어오는 패킷을 라우팅하기 전에 처리
- 각 체인에 대해 규칙을 추가하거나 삭제할 수 있다.


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