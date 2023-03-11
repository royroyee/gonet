# ARP Packet
[ARP에 대해서]()

## ARP 패킷 형식

![arp packet.png](img%2Farp%20packet.png)

- `Hardware Type`: 16비트, ARP 가 수행되고 있는 네트워크 유형을 정의하는 필드
  - 각 LAN은 유형에 따라 번호를 할당 받는다.
  - ex) 이더넷 : 1


- `Protocol Type` : 16비트, 프로토콜을 정의하는 필드
  - `ARP`는 어떤 상위 프로토콜과도 사용될 수 있다.
   - ex) IPv4 : 0800(16진수)


- `Hardware Length` : 8비트, 물리 주소의 길이를 바이트 단위로 정의
  - ex) 이더넷 : 6


- `Protocol Length` : 8비트, 논리 주소의 길이를 바이트 단위로 정의
  - ex) IPv4 : 4


- `Operation` : 16비트, 패킷의 유형을 정의(요청/응답)
  - `ARP 요청` : 1
  - `ARP 응답` : 2


- `Sender hardware address`(발신지 하드웨어 주소) : 가변 길이 필드, 송신자의 물리 주소를 나타낸다.
  - ex) 이더넷 : 6바이트


- `Sender protocol address`(발신지 프로토콜 주소) : 가변 길이 필드, IP 주소와 같은 송신자의 논리 주소를 나타낸다.
  - ex) IP 프로토콜 : 4바이트


- `Taget hardware address`(타깃 하드웨어 주소) : 가변 길이 필드, 타깃의 물리 주소를 정의
  - ex) 이더넷 : 6바이트
  - ARP 요청 메시지 : 송신자는 타깃의 물리 주소를 모르므로 0 이다.


- `Target protocol address`(타깃 프로토콜 주소) : 가변 길이 필드, IP 주소와 같은 타깃의 논리 주소
  - ex) IPv4: 4 바이트

## 캡슐화

 
## 실제 ARP 패킷 분석

### Basic

### 1. ARP Request

### 2. ARP Reply