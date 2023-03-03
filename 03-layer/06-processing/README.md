# Processing
> OSI 모델의 각 계층 간에 데이터가 전달되고 처리되는 전체 과정 알아보기

## 네트워크 구성 (요약)
- `Application Layer(Session,Presentation)` : 애플리케이션 등에서 사용하는 데이터를 송수신 하는 데 필요 (HTTP 등)
- `Transport Layer` : 목적지에 데이터를 정확하게 전달하는 데 필요 (TCP, UDP)
- `Network Layer` : 다른 네트워크에 있는 목적지에 데이터를 전달하는 데 필요(IP, 라우터)
- `Data Link Layer` : LAN에서 데이터를 송수신 하는 데 필요(이더넷, MAC 주소)
- `Physical Layer` : 데이터를 전기 신호로 변환하는 데 필요

### 예시 (랜 카드, 캡슐화 예제)
> 웹 브라우저에 URL 을 입력할 때부터, 웹 서버에 도착할 때 까지의 캡슐화 예제
1. 웹 사이트에 접속해야 하므로, application layer 에서부터 시작(URL 입력 후 Enter 누르면, 캡슐화가 시작)
   - tcp 3-way Handshake 는 이미 완료되어 연결이 확립되었다고 가정
   

2. application layer 에서 HTTP 프로토콜을 이용하여 HTTP 메세지를 보낸다. (GET /index.html HTTP/1.1 ~)


3. 해당 데이터가 transport layer 에 전달된다.


4. transport layer 에서는 `TCP 헤더` 가 붙어 어느 애플리케이션에 데이터를 보낼 지 식별한다.
   - 출발지 포트 번호는 잘 알려진 포트가 아닌(1025 이상) 포트 중에서 랜덤으로 할당 (여기선 3100번이 할당되었다고 가정)
   - 목적지 포트 번호는 HTTP 이므로 80번 포트 할당
   - 웹 브라우저의 3100번 포트에서 웹 서버의 80번 포트로 데이터를 전송이 가능해진다.


5. **세그먼트** 가 되어 network layer 에 전달된다.
   - 세그먼트 : TCP 헤더가 붙은 데이터


6. network layer 에서는 세그먼트에 **IP 헤더**를 붙인다.
   - IP 헤더에는 출발지 IP 주소와 목적지 IP 주소 등이 추가된다.
   - IP 헤더가 붙은 데이터를 IP 패킷이라고 한다.


7. 그 다음 data link layer 로 데이터가 전달된다.


8. data link layer 에서는 **이더넷 헤더**가 추가된다.
   - 이데넛 헤더가 추가된 데이터를 **이더넷 프레임** 이라고 한다.


9. physical layer 로 전달되어 전기신호로 변환되어 네트워크로 전송된다.

### 스위치, 라우터에서의 데이터 전달 및 처리
추후 작성 예정