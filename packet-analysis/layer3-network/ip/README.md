# ip


## ip 데이터그램
![IP 헤더.png](img%2FIP%20%ED%97%A4%EB%8D%94.png)
> ip 계층 가변길이 패킷
- IP 헤더 : 20~40 bytes
- 라우팅 및 전달에 필요한 정보를 포함
- 4바이트 단위로 표시

### 구성 요소
- `VER`(버전) : 4비트, IP 프로토콜의 버전을 나타냄(현재 IPv4)


- `HLEN`(헤더 길이) : 4비트, 헤더 길이가 20~60바이트 사이에서 변할 수 있으므로 필요한 필드
  - 해당 값에 4를 곱하여 구한다.
  - 선택사항이 없다면 필드의 값은 5 (5*4 = 20)
  - 선택사항 필드가 최대 길이라면 값은 15 (15 * 4 = 60)


- `Service Type`(서비스 유형) : 서비스 유형에서의 우선순위 등을 표현 (`TOS` 라고 불리기도 했었다.)
  - `DSCP` (Differentiated Service Code Point) : 요구하는 서비스의 우선순위에 대한 유형, 6비트
    - 즉, IP 데이터그램이 라우터에서 어떻게 처리해야 할지 정하는 역할
    - 0~7 의 우선순위가 존재하며 만약 라우터에 혼잡이 발생하여 데이터그램 일부를 폐기해야 한다면, 가장 낮은 우선순위 값을 가진 데이터그램이 가장 먼저 폐기 된다.
    - CS0 : 0 (16진수) / CS1 : 8 / CS2 : 16 .. CS7: 56

  - `ECN` (Explicit Congestion Notification) : 라우터가 패킷을 바로 폐기 하지 않고 최종 노드에 혼잡을 알리는 용도로 사용
    - 00 : 패킷이 ECN 기능을 사용하지 않음
    - 01 or 10 : 보내는 측이 종단점에서 ECN 기능을 사용
    - 11 : 라우터에서 혼잡이 발생할 때 


- `Total length`(전체 길이) : 헤더와 데이터를 포함하는 전체 길이를 바이트 단위로 나타낸다.
  - 상위 계층으로부터 받은 데이터 길이를 알아내기 위해 전체 길이에서 헤더 길이를 빼면 된다.
  - 데이터 길이 = 전체 길이 - 헤더 길이(HLEN*4)

> Identification, Flags, offset 은 모두 단편화와 관련된 필드이다.

- `Identification`(식별) : : 생성되는 각각의 패킷마다 부여된느 고유의 번호. 
  - 패킷은 제 2계층 프로토콜의 최대 전송 단위(MTU) 값에 따라 여러 개의 프로그먼트(Fragment)로 분할되어 처리되는데, 분할되어 온 fragment들을 원래의 패킷으로 재조립할 때 이 식별자 값을 기준으로 한다.
  - 단편화 된 임의의 패킷이 어떠한 데이터인지 구분하기 위해서 고유번호를 할당한다.(Fragment Identification)
  - 단편화 되지 않은 패킷 : 오직 하나의 고유 ID
  - 만약 3개로 단편화 된 패킷이라면? : 3개의 같은 아이디를 가진 패킷이 존재


- `Flags` : IP 패킷(데이터그램)의 분할(fragmentation) 가능 여부와 마지막 fragment인지 아닌지를 알리기 위해 사용되는 필드.
  - 3비트
    - Reserved bit(사용하지 않으므로 항상 0) / DF(Don't fragment bit) / MF(More fragment bit)
  - DF bit(두 번째 비트)  
    - 1 : 이 데이터그램을 단편화 하면 안된다는 의미(단편화 되지 않았다)
  - MF bit 
    - 1(0x2) : 분할된 패킷이 더 존재한다는 의미
    - 0(0x0) : 분할된 패킷이 더 이상 없다는 의미
  
- `Fragmentation offset` : 13비트, 분할 된 패킷을 수신 측에서 재조합 할 때 패킷들의 순서들을 파악하기 위한 필드
  - 4000바이트 패킷이 3개의 패킷으로 단편화 되었을 때(분할)
    - 첫 번째 패킷 : Flags : 1 / offset : 0 (0~1399번까지의 바이트를 가지고 있으며 0/8 = 0 이므로)
    - 두 번째 패킷 : Flags : 1 / offset : 175 (1400~2799번 까지의 바이트를 가지고 있으며 1400/8 = 175 이므로)
    - 세 번째 패킷 : Flags : 0(더 없으므로) / offset : 350 (2800~3999번 까지의 바이트를 가지고 있으며 2800/8 = 350 이므로)
    > 3 개의 패킷의 Identification 은 모두 동일
  
    

- `TTL(Time to Live)`(수명) : 제한된 수명, 값이 0이 되면 데이터그램은 폐기 된다.
  - 데이터그램에 의해 방문되는 최대 홉(라우터) 수를 제어하기 위해 사용
  - 각 라우터는 이 필드의 값을 1씩 감소시키며, 값이 0이 되면 라우터는 데이터그램을 폐기
  - 예시: 발신지가 패킷이 지역 네트워크 내에서만 전달되기를 원하여 값을 1로 설정 -> 첫 번째 라우터에서 0으로 감소되어 데이터그램이 폐기


- `Protocol` : 8비트, IP 계층의 서비스를 사용하는 상위 계층 프로토콜을 정의
  - IP 데이터그램은 TCP, UDP, ICMP, IGMP 등과 같은 여러 종류의 상위 계층 프로토콜을 캡슐화 할 수 있다.


- `Header Checksum` : IP 헤더의 오류 검사를 위한 필드. 
  - TCP와 UDP를 포함하여 IP 데이터그램으로 캡슐화되는 프로토콜은 대부분 헤더 및 데이터를 포함하는 체크섬 필드를 가지고 있기 때문에 IP 데이터그램의 체크섬 필드는 단순히 IP 헤더에 대한 오류 검사만을 수행.


- `Source IP Address` : 32비트, 발신지의 IP 주소를 정의
  - IP 데이터그램이 발신지에서 목적지까지 전달되는 동안 이 값은 변해서는 안 된다.


- `Destination Address` : 32비트, 목적지 IP 주소를 정의
  - 마찬가지로 전달되는 동안 변하지 않아야 할 값

## 단편화(Fragmentation)
데이터그램은 다른 네트워크를 지나갈 수 있다. 각 라우터는 수신한 프레임에서 IP 데이터그램을 `역 캡슐화`한 후 처리하고 다른 프레임 속에 `캡슐화` 한다.
</br>
- 수신한 프레임의 형식과 크기는 프레임이 방금 지나온 네트워크가 사용하는 프로토콜에 의존한다.
- 송신할 프레임의 형식과 크기는 프레임이 바로 다음에 지나갈 네트워크가 사용하는 프로토콜에 의존한다.

### 최대 전달 단위(MTU)
대부분의 프로토콜에서 각 데이터링크 계층(Layer 2)은 **자신의 프레임 형식을 가지고 있다**.
</br>
프레임의 형식에 정의된 필드 중의 하나는 데이터 필드의 최대 크기인 최대 전달 단위(MTU : Maximum Transfer Unit) 이다.

- 데이터그램이 프레임 속에 캡슐화가 될 때 데이터그램의 크기는 이 최대 크기(MTU)보다는 작아야 한다.
  - 보통 이 MTU는 네트워크 내에서 사용되는 하드웨어, 소프트웨어의해 주어지는 제한 조건에 의해 결정된다.

- MTU 의 값은 물리 네트워크 프로토콜마다 다르다.
  - `Ethernet LAN` : **1500 바이트** (가장 자주, 기본으로 쓰이는)
  - FDDI LAN : 4352 바이트

- IP 프로토콜을 물리적 네트워크에 독립적으로 만들기 위해 IP 데이터그램의 최대 길이는 65.535 바이트 이다.
  - 이렇게 MTU 가 큰 프로토콜을 사용한다면 패킷의 전달은 효율적일 수 있을 것이다.

> 그러나 MTU 가 작은 다른 네트워크에서는 데이터그램을 나누어서 보내야 하는데, 이것을 단편화(Fragmentation) 라고 한다.

---

## 실제 IP 패킷 캡처

### Basic
> 와이어샤크, 로컬 환경, yes24 접속 했을 때의 응답 패킷

### 1. VER, HLEN 
![ip version.png](img%2Fip%20version.png)
- `45`
- 45 에서 4는 VER 이므로 IPv4를 말하고, 5는 HLEN(5*4 = 20) 즉 헤더 길이는 20바이트를 의미

### 2. Service Type
![service type.png](img%2Fservice%20type.png)
- `00`
- `DSCP` : CS0 (가장 낮은 우선순위(default)) 
- `ECN` : 00 패킷이 ECN 기능을 사용하지 않음(대부분)

### 3. Total Length
![total length.png](img%2Ftotal%20length.png)
- `171` bytes

### 4. Identification
![identification.png](img%2Fidentification.png)
- flags 를 보고 단편화 되지 않을 것을 확인 -> 고유 ID 인 걸 알 수 있음

### 5. Flags
![flags.png](img%2Fflags.png)
- `0x2` : 위에서 언급했듯이 첫 비트는 사용하지 않는 비트(Reserved bit)
- DF 비트가 1이므로, 단편화 되지 않았다는 것을 알 수 있다.

### 6. Fragmentation offset
![offset.png](img%2Foffset.png)
- 앞 40 은 flags 값
- `00` 이므로 첫 번째 패킷. 단편화 되지 않았으므로 자동으로 처음이자 마지막 패킷 


### 7. TTL
![ttl.png](img%2Fttl.png)
- 242 : 총 242번 이동 가능

### 8. Protocol
![protocol.png](img%2Fprotocol.png)
- `6` : TCP

### 9. Checksum
![checksum.png](img%2Fchecksum.png)
- Correct
- 계산은 생략

### 10. Source / Destination IP Address
![ip address.png](img%2Fip%20address.png)
- 해당 패킷은 yes24 로 부터 접속했을 때 reply 패킷이다.
- Source Address 는 송신자의(yes24) ip 주소
- Destination Address : 내 컴퓨터 의 ip 주소

### Fragment Example
- 작성 예정