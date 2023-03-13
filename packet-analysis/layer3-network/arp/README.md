# ARP Packet
[ARP에 대해서](https://github.com/royroyee/gonet/tree/main/03-layer/03-network-layer#arpaddress-resolution-protocol-%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C)

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
![캡슐화(arp).png](img%2F%EC%BA%A1%EC%8A%90%ED%99%94%28arp%29.png)
`ARP 패킷`은 데이터 링크(Layer2) 프레임으로 **캡슐화** 된다.

즉, `ARP 패킷` 은 이더넷 프레임에 의해 캡슐화 되어 있다.

### ARP 프로세스
1. 송신자는 타깃의 `IP 주소`를 알고 있다고 가정


2. `IP` 가 `ARP` 에게 `ARP 요청 메시지`를 생성하도록 요청한다.
   - 이 요청 메시지에서 송신자의 물리 주소와 IP 주소 그리고 타깃의 IP 주소는 채워지지만 타깃의 물리 주소 필드는 0으로 채워진다.
   - 물리 주소는 MAC 주소로 생각하면 된다.


3. 이 메시지는 데이터링크 계층(Layer2)에 전달되고 여기서 송신자의 물리 주소를 프레임의 발신지 주소로, 물리 브로드캐스트 주소를 프레임의 목적지 주소로 하는 프레임에 의해 `캡슐화` 된다.


4. 모든 호스트나 라우터는 이 프레임을 수신한다.
   - 프레임은 브로드 캐스트 목적지 주소를 가지고 있으므로 모든 스테이션은 이 메시지를 자신의 `ARP` 에게 전달한다.
   - 타깃을 제외하고는 모든 라우터,호스트는 이 패킷을 폐기한다. 
   - 타깃은 `IP 주소`를 인식한다.


5. 타깃 호스트 또는 라우터는 자신의 물리 주소를 포함하는 `ARP 응답 메시지`를 보낸다. 이 메시지는 `유니캐스트` 된다.


6. 송신자는 응답 메시지를 받고 타깃의 물리 주소를 알게 된다.


7. 타깃에게 보내질 데이터를 포함하고 있는 `IP 데이터그램`은 이 물리 주소를 가지는 프레임으로 `캡슐화` 되어 목적지에 `유니캐스트` 된다.
    - 쉽게 말하면, 목적지 호스트의 `IP 주소` 를 이제 알게 되었으니 이전에 생성한 `IP 패킷`에 MAC 주소를 추가하여 이더넷 프레임을 생성하고, 생성된 이더넷 프레임을 목적지 호스트 또는 라우터로 전송한다.(유니캐스트)

 
## 실제 ARP 패킷 분석

### Basic
> 와이어샤크, 로컬환경 이용

### 1. ARP Request

### Destination
![arp_destination.png](img%2Farp_destination.png)
- 위에서 언급했듯이 ARP 패킷은 이더넷 프레임에 캡슐화 되어있고, 따라서 프레임의 목적지 주소를 확인해야한다.
- ff:ff:ff:ff:ff:ff 이므로, 브로드캐스트 임을 알 수 있다.

### ARP Request 패킷
![arp request.png](img%2Farp%20request.png)
- 주의 깊게 봐야할 부분
- Opcode : 1 -> request
- `ARP Request ` 이므로 목적지의 IP 주소만 알 뿐 MAC 주소는 알지 못하는 상태다.
- 즉, `Target MAC address` 가 0으로 표시되었다.


### 2. ARP Reply

### Destination
![arp_reply_destination.png](img%2Farp_reply_destination.png)
- `ARP Reply` 패킷이므로 당연히 나의 컴퓨터 mac 주소를 가리키는 모습이다.

### ARP Reply 패킷
![arp_reply_packet.png](img%2Farp_reply_packet.png)
- 여기서 주의깊게 봐야할 부분
- Opcode : 2 -> Reply
- Target MAC address 가 나의 컴퓨터 mac 주소를 반환함으로써 목적지의 mac 주소를 반환하는 것을 알 수 있다.