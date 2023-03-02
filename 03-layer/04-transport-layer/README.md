# Transport Layer
4계층, 전송 계층(Transport Layer)
이전에 배운 1,2,3 Layer 들은 목적지로 데이터를 보낼 순 있지만 데이터가 손상되거나 유실되더라도 아무 처리를 해주지 않는다.

> Transport Layer 는 목적지에 신뢰할 수 있는 데이터를 전달하기 위해 필요하다.

### 기능
1. **오류를 점검**하는 기능
   - 데이터가 제대로 도착했는 지 확인 (네트워크 계층(Layer 3)은 목적지까지 데이터를 전달)
2. **전송된 데이터의 목적지가 어떤 애플리케이션인지 식별**하는 기능

### 연결형 통신 & 비연결형 통신
> Transport Layer 는 **신뢰성, 정확성**과 **효율성** 으로 구분할 수 있다.
- `연결형 통신` : **신뢰**할 수 있고 **정확**한 데이터를 전달하는 통신
  1. 송신자 : 데이터를 보내도 되나?
  2. 수신자 : 보내도 된다.
  3. 송신자 : 보냄
  4. 수신자 : 받음
  5. 송신자 : 확인
-  대표적으로 TCP 가 연결형 통신에 해당한다.

- `비연결형 통신` : **효율**적으로 데이터를 전달하는 통신
  1. 송신자 : 데이터를 보냄
  2. 수신자 : 확인
- 대표적으로 UDP 가 비연결형 통신에 해당한다.

## TCP
> 신뢰성과 정확성을 우선으로 하는 연결형 통신 프로토콜
- `TCP 헤더` : TCP 로 전송할 때 붙이는 헤더
  - IP 헤더와 마찬가지로 많은 정보를 가지고 있다.
- `세그먼트(segement)` : TCP 헤더가 붙은 데이터를 세그먼트 라고 한다.

### TCP 헤더 구성 요소 (주요 부분만)
- 출발지 포트 번호(16비트)
- 목적지 포트 번호(16비트)
- 확인 응답 번호(32비트)
- 코드 비트(6비트)
  >데이터를 전송하려면 먼저 연결이라는 가상의 독점 통신로를 확보해야 한다. 코드 비트에 연결의 제어정보가 기록된다.
  - 주요 코드 비트
    - `SYN`: 연결 요청
    - `ACK` : 확인 응답
    - `FIN` : 연결 종료

### TCP 3-way HandShake
- 위에서 언급한 `SYN`, `ACK` 를 사용하여 연결을 확립할 수 있는데, 신뢰할 수 있는 연결을 하려면 데이터를 전송하기 전에 패킷을 교환한다. 이 때 패킷 요청을 3번 교환하는 것을 3-way Handshake 라고 한다.

다음은 컴퓨터 1 과 컴퓨터 2 의 연결 예시이다.

#### 연결 확립

1. 컴퓨터1 은 컴퓨터2에게 **연결 확립 요청(SYN)**
   - 통신을 하려면 컴퓨터2 에게 허가를 받아야 한다. 때문에 컴퓨터1은 컴퓨터2로 연결 확립 허가를 받기 위한 요청(SYN)을 보낸다.
   

2. 컴퓨터2는 **연결 확립 응답(SYN)**을 하고, **연결 확립 요청(ACK)** 한다.
   - 컴퓨터2는 컴퓨터1이 보낸 요청을 받은 후에 허가한다는 응답을 회신하기 위해 연결 확립(ACK)을 보낸다. 동시에 컴퓨터2도 컴퓨터1에게 데이터 전송을 허가 받기 위해 연결확립 요청(SYN)을 보낸다.


3. 컴퓨터1은 **연결 확립 응답(SYN)**을 한다.
   - 컴퓨터2의 요청을 받은 컴퓨터1은 컴퓨터2로 허가한다는 응답으로 연결 확립 응답(ACK)을 보낸다.

#### 연결 종료
1. 컴퓨터1 에서 컴퓨터2로 연결 종료 요청(FIN)을 보낸다.
2. 컴퓨터2 에서 컴퓨터1로 연결 종료 응답(ACK)을 반환한다.
3. 동시에 컴퓨터2에서도 컴퓨터1로 연결 종료 요청(FIN)을 보낸다.
4. 컴퓨터1에서 컴퓨터2로 연결 종료 응답(ACK)을 반환한다.
> TCP는 신뢰성과 정확성을 중요하게 여기는 프로토콜이다. 연결을 확립할 때 뿐만 아니라, 연결을 종료할 때도 FIN 을 이용하여 확인한다.