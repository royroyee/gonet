# ICMP 

`IP 프로토콜` 은 오류 보고 및 수정 기능이 없다. 만약 무엇인가 잘못이 일어나면 어떻게 해결해야 할까?
이러한 부분을 해결해주는 것이 바로 `ICMP` 프로토콜이다.
> IP 는 신뢰성이 없는 프로토콜이다.

### ICMP 의 위치
- 네트워크 계층
- IP 프로토콜의 동반 프로토콜
- 직접 데이터링크(Layer2) 계층으로 전달되지 않는다.
> IP 데이터그램 내에 캡슐화 된다.
![icmp 캡슐화.png](img%2Ficmp%20%EC%BA%A1%EC%8A%90%ED%99%94.png)
> IP 데이터가 ICMP 메시지 임을 알리기 위해 IP 데이터그램의 프로토콜 필드의 값은 1이 된다.


### ICMP 메시지
크게 두 종류로 나눌 수 있다.

1.  **Error-reporting messages**
    - Type
      - 3 : Destination unreachable
      - 4 : Source quench
      - 11 : Time exceeded
      - 12 : Parameter problem
      - 5 : Redirection

2. **Query messages**
   - Type
     - 8 or 0 : Echo request or reply
     - 13 or 14 : Timestamp request or reply
   

### ICMP 메시지 형식(구조)
![ICMP 형식.png](img%2FICMP%20%ED%98%95%EC%8B%9D.png)

- `Type` : 메시지의 유형
- `Code` : 각 타입에 대한 코드 값(이유)
- `Checksum` : ICMP 메시지 자체(헤더+데이터)의 오류를 검사하는 필드
- `Rest of the header, Data Section`: 이 필드는 메시지 유형 마다 다르므로, 아래에서 다시 다룬다.

### Error - reporting messages
각 유형에 따라 여러가지 코드가 있다.
- 구조는 위의 그림에서 Rest of the Header는 unused 필드가 된다.(항상 0)
- 나머지 구조는 동일하다.
</br>

5개의 타입 중 두 가지만 여기서 서술(Destination unreachable, Time exeeded)

### Destination unreachable (목적지 도달 불가)
> 라우터가 데이터그램을 전달할 수 없거나, 호스트가 데이터그램을 배달할 수 없을 때 데이터그램은 폐기되고 라우터나 호스트는 데이터그램을 시작했던 발신지 호스트에게 목적지 도달 불가 메시지를 보낸다.
- Code (0~15) 로 이유를 나타낸다.
  - 0 : 하드웨어 고장 등의 이유로 네트워크 도달 불가능
  - 1 : 호스트에 도달할 수 없다.(하드웨어 고장 등)
  - 이하 생략

### Time exceeded(시간 경과)
- `TTL` 파트에서 언급했듯이, 패킷은 다음 라우터를 찾기 위해 라우터는 라우팅 테이블을 사용하는데, 만약 오류 등의 문제로 루프나 사이클을 지날 수 있다.
  - 무한 루프 등을 방지 하기 위해서 각 데이터그램은 `TTL` 필드를 가지고 있다.
  - 데이터그램이 라우터를 방문할 때 이 필드의 값은 1씩 감소된다.
  - 0이 되면 라우터는 이 데이터그램을 폐기한다.
> 이렇게 폐기가 될 때, 라우터는 시간 경과 메시지를 원 발신지에 송신해야 한다.

### Query messages

### 1. 에코 요청과 응답
고장 진단의 목적으로 설계되어, 이 메시지를 사용하여 네트워크 문제를 발견할 수 있다.
</br>

에코 요청과 에코 응답 메시지의 조합으로 두 시스템(호스트 or 라우터)이 서로 통신할 수 있는 지 결정할 수 있다.
</br>

호스트나 라우터는 `에코 요청 메시지`를 다른 호스트,라우터에게 보낼 수 있고 `에코 요청 메시지`를 받은 호스트,라우터는 `에코 응답 메시지를 생성`하여 원래의 송신자에게 보낸다.

### 구조
![echo 구조.png](img%2Fecho%20%EA%B5%AC%EC%A1%B0.png)
- Type : 8 -> Echo Request
- Type : 0 -> Echo Reply
- `Identifier` : 프로세스 ID, 여러 개의 ping이 동일 호스트에서 실행되는 경우 식별하기 위해 사용
- `Sequence number` : 요구 패킷의 순서, 0부터 시작
- `Code` : 0


### 2. 타임스탬프 요청과 응답
두 호스트나 라우터는 타임스탬프 요청과 응답 메시지를 사용하여 IP 데이터그램이 이 둘 사이를 지나가는 데 필요한 왕복시간(round-trip time)을 결정할 수 있다.

- `발신지` : 타임스탬프 요청 메시지를 생성
    - 발신지는 출발시각에서의 자신의 시계 값을 세계 표준시로 표현한 값을 원래 타임스탬프(original timestamp) 필드에 삽입한다.
    - 나머지 Recevie, Transmit 타임스탬프 필드는 0으로 채워져 있다.


- `목적지` : 타임스탬프 응답 메시지를 생성
  - 요청 메시지에 있는 original 타임스탬프 값을 응답 메시지의 같은 필드에 복사한다.
  - receive 타임스탬프 필드에 요청이 수신된 시점에서의 자신의 시계 값을 세계 표준시로 표현하여 삽입
  - transmit 타임스탬프 필드에는 응답 메시지가 출발하는 시점의 시계값을 세계 표준시로 표현하여 삽입

### 구조
![timestamp 구조.png](img%2Ftimestamp%20%EA%B5%AC%EC%A1%B0.png)
- Type : 13 -> Timestamp Request
- Type : 14 -> Timestamp Reply
- Code : 0
- Identifier :  프로세스 ID
- Sequence number : 요구 패킷의 순서

3개의 timestamp 필드 들은 32비트 길이를 가진다.
- 자세한 내용은 위의 타임스탬프 요청과 응답 발신지,목적지 내용 참고

### Ping
ping 프로그램을 사용하여 호스트가 정상적으로 작동하고, 응답하는 지를 점검할 수 있다.


## 실제 ICMP 패킷 분석
> 와이어샤크, 로컬 환경, ping 8.8.8.8(구글)이용

### 1. Echo Request
![echo request_type.png](img%2Fecho%20request_type.png)
- 먼저 타입이다. Type : 8 이므로, Echo Request 임을 알 수 있다.
- Timestamp 필드도 같이 보내진 것을 알 수 있다.

### 2. Echo Reply
![echo_reply.png](img%2Fecho_reply.png)
- Type : 0 이므로 Echo Reply 임을 알 수 있다.

TTL 만료 패킷 예제