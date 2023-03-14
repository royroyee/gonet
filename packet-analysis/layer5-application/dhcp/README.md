# DHCP

[DHCP란?](https://github.com/royroyee/gonet/tree/main/03-layer/05-application-layer#dhcp)

## 패킷 구조
![DHCP 패킷.png](img%2FDHCP%20%ED%8C%A8%ED%82%B7.png)
- `Operation code`(동작 코드) : 8비트, DHCP 패킷의 종류를 나타내는 필드
  - 요청 : 1 / 응답 2
- `Hardware type` :  8비트, 물리 네트워크의 유형
  - 이더넷 : 1
- `Hardware length` : 8비트, 물리주소의 길이를 바이트 단위로 나타낸다.
  - 이더넷 : 6
- `hop count`(홉 개수) : 8비투, 패킷이 전달 될 수 있는 최대 홉 개수를 나타낸다.
- `transaction ID` : 4바이트, 요청에 대한 응답을 확인하기 위하여 사용
  - 서버는 요청메시지 내에 있는 것과 동일한 값을 응답메시지에 기록
- `Number of seconds`(초 단위 경과시간) : 16비트, 클라이언트가 부팅 된 후 경과된 시간을 초 단위로 나타낸다. 
- `Flag` : 16비트, 맨 왼쪽 비트만 사용하고 나머지 비트들은 0으로 설정
  - 맨 왼쪽 비트 : 서버로 부터의 응답이 유니캐스트가 아니라 강제적인 보르대크스트 응답이어야 함을 명시
    - 0 : 유니캐스트
    - 1 : 브로드캐스트
- `Client IP address` : 4바이트, 클라이언트 IP 주소를 포함한다.
  - 만약 클라이언트가 이 정보를 가지고 있지 않다면 모두 0으로 지정
- `Your IP address` : 4바이트, 클라이언트 IP 주소를 포함, 요청을 받은 서버에 의해 응답 메시지에 기록
- `Server IP address` : 4바이트, 서버의 IP 주소를 포함, 서버는 이 값을 응답 메시지에 기록
- `Gateway IP address` : 4바이트, 라우터에 IP 주소를 나타내고 서버에 의해 응답 메시지에 기록
- `Client hardware address` : 클라이언트의 물리 주소
- `Server name` : 64바이트, 서버가 응답 메시지에 기록하는 선택 항목, null 문자열로 끝나는 서버의 도메인 이름을 포함
- `Boot filename` : 128바이트, 서버가 응답 메시지에 기록하는 선택 항목
- `Option` : 64바이트 목적은 두 가지
  - 1. 네트워크 마스크 또는 기본 라우터 주소와 같은 추가적 정보 전달하기 위해
  - 2. 특정 제조업체의 해당 정보를 전다하기 위해

### [DHCP 동작 방식](https://github.com/royroyee/gonet/tree/main/03-layer/05-application-layer#dhcp-%EB%8F%99%EC%9E%91-%EB%B0%A9%EC%8B%9D)