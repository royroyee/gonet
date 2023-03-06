# Layer 5, Application Layer
> 애플리케이션이 동작하는 계층
>
여기에선 Session, Presentation Layer 을 포함해서 서술

- Application Layer 에서는 클라이언트의 요청을 전달하기 위해 통신의 대상(서버 등)이 이해할 수 있는 데이터로 변환하고 Transport Layer 로 전달하는 역할을 한다.
  - 이때도 역시 Application Layer 의 프로토콜을 사용해야 한다.

### 대표적인 Application Layer 의 프로토콜
- `HTTP` : 웹 사이트 접속
- `DNS` : 이름 해석
- `FTP` : 파일 전송
- `SMTP` : 메일 송신
- `POP3` : 메일 수신

### 각 계층 별로 기술,프로토콜 정리(요약)
> Application Layer 에서 보내려는 데이터는 하위 계층으로 (순서대로) 전달되어 처리된다.

- `Application Layer` : HTTP, DNS, FTP, SMTP, POP3 , 기타
- `Transport Layer` : TCP , UDP
- `Network Layer` : IP 등
- `Data Link Layer` : Ethernet
- `Physical Layer` : 전기 신호 변환

---

## 웹 서버 Web Server
인터넷에서 핵심역할을 하고 있는 `WWW` 웹서버
- `HTML`, `URL`, `HTTP` 세가지 기술 사용
- html 파일 , 이미지 파일이 웹 서버에 전송되는 방식
  - 문서와 이미지는 각각 별도로 요쳥하므로, 사용자에게 보이는 순서가 제각각일 때가 있다.

## HTTP
클라이언트(웹 브라우저)는 웹 사이트를 보기 위해 서버의 80번 포트를 사용하여 HTTP 통신을 한다.
클라이언트에서 HTTP 요청, 서버에서 HTTP 응답을 반환한다.
(구체적인 내용은 추후 서술)

### Keep-Alive
기존 HTTP/1.0 버전에서는 요청을 보낼 때마다 연결 수립 과 연결 끊기를 반복했지만, HTTP/1.1 버전에서부터 연결을 한 번 수립하면 데이터 교환을 마칠 때까지 유지하고,
데이터 교환을 모두 끝내면 연결을 끊는 구조 가  추가되었다. 이 기능을 `keep alive` 라고 한다.
- 매번 연결 수립 및 끊기를 하지 않아도 되어 성능이 향상된다.

### HTTP/2
HTTP/1.1 버전에선 요청을 보낸 순서대로 요청을 반환한다. 만약 이전 요청 처리 시간이 길어지면 다음 요청에 대한 처리도 계속 늦어지게 된다.
> HTTP/2 버전은 이러한 문제를 해결하기 위해 요청을 보낸 순서대로 응답을 반환하지 않아도 된다.

---

## DNS
URL을 IP 주소로 변환하는 서비스(시스템)
- `이름 해석(name resolution)` : IP 주소가 아닌 https://github.com/royroyee 과 같은 주소로 사용하여 접속하여 돕는 것
  - `DNS 서버` 가 https://github.com/royroyee 에 접속하면 이 웹 사이트 서버의 IP 주소를 알려준다.
- `도메인 이름` : https://github.com/royroyee 처럼 컴퓨터, 네트워크를 식별하기 위해 붙여진 이름
- `호스트 이름(서버 이름)` : 도메인 이름 앞에 있는 `WWW` 등

### 예시
https://github.com/royroyee 웹 브라우저가 이 URL 을 입력하여 접속 시도 시 
1. 클라이언트는 DNS 서버에 https://github.com/royroyee 의 IP 주소를 요청
2. DNS 서버는 해당 요청에 해당하는 도메인 이름의 IP 주소를 반환
3. 클라이언트는 받은 IP 주소로 서버에 접속
> DNS 서버는 전 세계에 흝어져 있으며 요청받은 DNS 서버가 해당 도메인 이름의 IP 주소를 모르는 경우 다른 DNS 서버에 요청하기도 한다.


## SMTP, POP3 (메일 서버)
- `SMTP` : 메일을 보내는 데 사용되는 프로토콜 , 포트 번호 25
- `POP3` : 메일을 받는 데 사용되는 프로토콜, 포트 번호 110