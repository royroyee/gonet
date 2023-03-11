# Ethernet Packet

[이더넷 Basic](https://github.com/royroyee/gonet/tree/main/03-layer/02-data-link-layer#%EC%9D%B4%EB%8D%94%EB%84%B7-ethernet)

### 이더넷(Ethernet)
> Layer 2의 프로토콜. 표준화된 유선 LAN(Local Area Network) 기술. 
> 
- 이더넷은 일반적으로 케이블을 사용하여 컴퓨터, 스위치, 라우터, 모뎀 등의 장치들을 서로 연결하여 통신.
- CSMA/CD(Carrier Sense Multiple Access with Collision Detection) 방식을 사용하여 충돌을 감지하고 회피하는 프로토콜.
  - 따라서 여러 대의 컴퓨터가 동시에 네트워크를 사용하더라도 충돌을 최소화하고 안정적인 통신을 유지할 수 있다.

### 프레임(Frame)
> 이더넷 LAN에서 보내지는 패킷

![이더넷 프레임.png](img%2F%EC%9D%B4%EB%8D%94%EB%84%B7%20%ED%94%84%EB%A0%88%EC%9E%84.png)
### 프레임 형식
- 총 8개 필드로 구분
- 1byte = 8bit


- `Preamble` : 7bytes
  - 물리 계층에서 추가.(실제로는 프레임의 일부라고 볼 수 없는 이유)


- `SFD`(시작 프레임 식별자) : 다음 필드가 목적지 주소임을 알린다.(마지막 두 비트 11)
  - 물리 계층에서 추가.


- `DA`(목적지 주소) : 6bytes, 패킷을 수신하는 목적지 MAC 주소 


- `SA`(발신지 주소) : 6bytes, 패킷 송신자의 MAC(물리) 주소


- `Length or Type` : 길이 또는 유형 필드
  - IEEE 표준에선 데이터 필드에 있는 바이트의 수를 나타내는 길이로 사용


- `Data` : 상위 계층 프로토콜로부터 캡슐화된 데이터를 전달
  - 46~1500 bytes


- `CRC` : 마지막 필드에는 오류 검출 정보가 포함

### 이더넷 MAC 주소
- 6bytes(48bits) 로 16진수로 표현
  - 총 12자리의 16진수로 보통 표현한다. (16진수 1자리 : 4비트)

### Example
> 4A:30:10:21:10:1A


### 유니캐스트, 멀티캐스트, 브로드캐스트 주소
> 발신지 주소는 항상 유니캐스트

**목적지 주소**
- 유니캐스트 : 목적지 주소의 첫 번째 바이트의 최하위 비트가 0
  - ex) 4A:30:10:21:10:1A -> A -> 1010(짝수) 
- 멀티캐스트 : 유니캐스트 조건을 만족하지 않는 경우 
  - ex) 47:20:1B:2E:08:EE -> 7 -> 0111(홀수)
- 브로드캐스트 : 48비트 모두 1인 경우 
  - FF:FF:FF:FF:FF

--- 

## 실제 이더넷 패킷 캡처

### Basic
> 와이어샤크, 로컬환경, 구글 8.8.8.8 로 테스트

![ping 테스트.png](img%2Fping%20%ED%85%8C%EC%8A%A4%ED%8A%B8.png)
- 8.8.8.8 과의 패킷만 보기위해 필터(Capture -> Options -> host 8.8.8.8 입력)


![Destination.png](img%2FDestination.png)
- 목적지 주소 00 1a 1e 03 90 50

![source.png](img%2Fsource.png)
- 송신지 주소 (내 컴퓨터 mac 주소와 일치 확인)