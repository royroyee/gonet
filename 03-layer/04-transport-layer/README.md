# 4 Layer, Transport Layer
4계층, 전송 계층(Transport Layer)
이전에 배운 1,2,3 Layer 들은 목적지로 데이터를 보낼 순 있지만 데이터가 손상되거나 유실되더라도 아무 처리를 해주지 않는다.

> Transport Layer 는 목적지에 신뢰할 수 있는 데이터를 전달하기 위해 필요하다.

### 기능
1. **오류를 점검**하는 기능
   - 데이터가 제대로 도착했는 지 확인 (네트워크 계층(Layer 3)은 목적지까지 데이터를 전달)
2. **전송된 데이터의 목적지가 어떤 애플리케이션인지 식별**하는 기능

### 연결형 통신 & 비연결형 통신
> Transport Layer 는 **신뢰성, 정확성**과 **효율성** 으로 구분할 수 있다.

### 연결형 통신 서비스
**신뢰**할 수 있고 **정확**한 데이터를 전달하는 통신

1. 가장 먼저 클라이언트와 서버의 연결을 설정한다.
2. 데이터 교환은 연결이 설정된 이후에 가능하다.
3. 데이터 교환이 완료된 후에는 연결이 해지된다.

간단 요약
  1. 송신자 : 데이터를 보내도 되나?
  2. 수신자 : 보내도 된다.
  3. 송신자 : 보냄
  4. 수신자 : 받음
  5. 송신자 : 확인
-  대표적으로 TCP 가 연결형 통신에 해당한다.

### 비연결형 통신 
**효율**적으로 데이터를 전달하는 통신

1. 발신지 프로세스(응용프로그램)는 메세지를 전송 계층에서 수신 가능한 크기의 **여러 개의 데이터 조각으로 나눈** 후, 데이터 조각들을 **하나씩 전송 계층으로 전달**한다.
2. **전송 계층은 이 데이터들의 관계는 고려하지 않고**, 각각의 데이터 조각을 **독립적인 하나의 단위로 간주**한다.
3. 응용 계층으로 부터 하나의 조각이 들어오면, **전송 계층은 조각을 패킷으로 캡슐화 한 후에 전송**한다.
4. 이 조각들은 **순서에 맞게 비연결형 전송 계층에 전달**되는 데,**전송 계층에선 아까 언급했듯이 패킷간의 연관성이 없으므로**, **목적지에는 순서가 어긋나게 도달**할 수 있다.
   - 특히 하나라도 손실되면 더 악화.

정말 간단하게 요약하면   
  1. 송신자 : 데이터를 보냄
  2. 수신자 : 확인
- 대표적으로 UDP 가 비연결형 통신에 해당한다.

---

## TCP
> 신뢰성과 정확성을 우선으로 하는 연결형 통신 프로토콜
- `TCP 헤더` : TCP 로 전송할 때 붙이는 헤더
  - IP 헤더와 마찬가지로 많은 정보를 가지고 있다.
- `세그먼트(segement)` : TCP 헤더가 붙은 데이터를 세그먼트 라고 한다.
> TCP는 데이터 전송 시 세그먼트를 여러 개로 분할하여 전송하고, 이를 통해 데이터 전송의 신뢰성과 속도를 높일 수 있다.

### TCP 헤더 구성 요소 (주요 부분만)
- **출발지 포트 번호(Source port number)** : 16비트
- **목적지 포트 번호(Destination port number)** : 16비트
- **일련번호(Sequence number)** : 32비트, 송신 측에서 수신 측에 이 데이터가 몇 번째 데이터인지 알려주는 역할
  - 패킷은 나누어서 수신되기 때문에, 수신측에서 순서에 맞게 조립해야 데이터를 해석할 수 있다. 때문에 sequence number 가 필요
- **확인 응답 번호(Acknowledgment number)** : 32비트, 다음에 받을 것으로 예상되는 데이터 옥텟의 순서번호를 나타낸다.
  - 몇 번째 데이터를 수신했는지 송신 측에 알려준다. -> 다음 번호의 데이터를 요청하는 데 사용된다.
- **코드 비트(6비트)**
  >데이터를 전송하려면 먼저 연결이라는 가상의 독점 통신로를 확보해야 한다. 코드 비트에 연결의 제어정보가 기록된다.
  - 주요 코드 비트
    - `SYN`: 연결 요청
    - `ACK` : 확인 응답
    - `FIN` : 연결 종료
- **체크섬(Checksum)** : 16비트, 중간에 훼손이 되었는지 또는 변조되었는 지 체크
- **윈도우 크기(Window size)** : 16비트, 버퍼의 한계 크기(얼마나 많은 용량의 데이터를 저장해 둘 수 있는지)
  - 효율성 때문에 송신 측은 수신 측에 세그먼트를 연속해서 보내는 데, 이때 수신 측에서 오버플로우가 발생하지 않도록 하기 위해 버퍼의 한계 크기를 헤더에 넣는다.

### TCP 3-way HandShake
- 위에서 언급한 `SYN`, `ACK` 를 사용하여 연결을 확립할 수 있는데, 신뢰할 수 있는 연결을 하려면 데이터를 전송하기 전에 패킷을 교환한다. 이 때 패킷 요청을 3번 교환하는 것을 3-way Handshake 라고 한다.
> TCP 가 통신을 맺는 과정을 3-way Handshake 라고 한다.

다음은 컴퓨터 1 과 컴퓨터 2 의 연결 예시이다.

#### 연결 확립

1. 컴퓨터1 은 컴퓨터2에게 **연결 확립 요청(SYN)**
   - 통신을 하려면 컴퓨터2 에게 허가를 받아야 한다. 때문에 컴퓨터1은 컴퓨터2로 연결 확립 허가를 받기 위한 요청(SYN)을 보낸다.
   

2. 컴퓨터2는 **연결 확립 응답(ACK)** 과, **연결 확립 요청(SYN)** 한다.(같은 패킷에 ACK 플래그를 1, SYN 플래그를 1로 설정)
   - 컴퓨터2는 컴퓨터1이 보낸 요청을 받은 후에 허가한다는 응답을 회신하기 위해 연결 확립(ACK)과. 동시에 컴퓨터2도 컴퓨터1에게 데이터 전송을 허가 받기 위해 연결확립 요청(SYN)을 보낸다.


3. 컴퓨터1은 **연결 확립 응답(SYN)** 을 한다.
   - 컴퓨터2의 요청을 받은 컴퓨터1은 컴퓨터2로 허가한다는 응답으로 연결 확립 응답(ACK)을 보낸다.

### TCP 4-way HandShake
> TCP 소켓 연결을 종료하는 과정을 4-way handshake 라고 한다.
#### 연결 종료
1. 컴퓨터1 에서 컴퓨터2로 연결 종료 요청(FIN)을 보낸다.
2. 컴퓨터2 에서 컴퓨터1로 연결 종료 응답(ACK)을 반환한다.
3. 동시에 컴퓨터2에서도 컴퓨터1로 연결 종료 요청(FIN)을 보낸다.
4. 컴퓨터1에서 컴퓨터2로 연결 종료 응답(ACK)을 반환한다.
> TCP는 신뢰성과 정확성을 중요하게 여기는 프로토콜이다. 연결을 확립할 때 뿐만 아니라, 연결을 종료할 때도 FIN 을 이용하여 확인한다.


## 포트 번호(Port number)
- Transport Layer 의 기능 중 하나인 전송된 데이터의 목적지가 **어떤 애플리케이션인지 구분하는 역할**을 하기 위해 필요한 것
- TCP 헤더의 출발지 포트 번호 , 목적지 포트 번호가 이에 해당
- 대표적인 포트 번호 
  - SSH : 22
  - DNS : 53
  - HTTP : 80
  - HTTPS : 443

---

## UDP
> TCP 와 달리 효율성을 중요하게 여기는 프로토콜
- 비연결형 통신, TCP 처럼 확인 작업을 일일이 하지 않고 데이터를 효율적으로 빠르게 보낸다.
  - 스트리밍 방식으로 전송하는 동영상 서비스 등에 사용 (TCP 데이터 통신으로 동영상을 전송하면 느리다.)
- **브로드캐스트** 통신에 적합하다.
  - TCP 는 3-way handshake 등의 과정이 필요하기 때문에 브로드캐스트와 같이 불특정 다수에게 보내는 통신에게는 적합하지 않음
### UDP 헤더
1. 출발지 포트 번호 (16비트)
2. 목적지 포트 번호 (16비트)
3. 길이(16비트)
4. 체크섬(16비트)
- TCP 헤더에 비해 매우 적은 정보들이 담겨져 있다.