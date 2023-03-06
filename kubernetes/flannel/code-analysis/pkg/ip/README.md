# ip package
****ip package 의 주요 코드들을 분석****


## 사전지식
### golang `net` package 의 Interface 이해하기
- 네트워크 인터페이스를 나타내는 구조체
```go
type Interface struct {
	Index        int          // positive integer that starts at one, zero is never used
	MTU          int          // maximum transmission unit
	Name         string       // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
}
```
- `Index` : 인터페이스의 고유한 숫자 인덱스, 일반적으로 1부터 시작하며 0은 사용하지 않음
- `MTU` : 인터페이스의 최대 전송 단위, **네트워크 패킷의 최대 크기**를 결정
- `Name` : 인터페이스의 이름 (ex eth0.100)
- `HardwareAddr` : 인터페이스의 MAC(하드웨어)주소
- `Flags` : 인터페이스의 특성을 설명하는 플래그


### Flannel 에서의 링크 로컬 주소와 글로벌 주소
- 링크 로컬 주소 : 같은 노드에 있는 컨테이너끼리의 통신을 위한 주소
- 글로벌 주소 : 다른 노드나 외부 인터넷과 통신하기 위한 주소
  - 다른 노드와 통신하기 위해서는 노드 간에 IP 주소가 할당되어야 하고, 이 때 각 노드에 할당되는 IP 주소는 서로 다른 네트워크 대역에 속해 있어야 한다.
  - Flannel은 각 노드에게 고유한 IP 주소를 할당하여 다른 노드와 통신이 가능하도록 한다.


---

### endianess.go

> 바이트 순서를 일관되게 맞추어주기 위해 엔디안 처리를 하는 코드
- IP 패킷은 여러 바이트로 이루어진 값을 전달하는데, 이 값들은 순서가 중요하다. 이를 전송 순서 라고 한다(Byte Order)

## iface.go
> 네트워크 인터페이스와 관련된 정보를 가져오는 함수들을 제공하는 것이 주된 기능
- 주어진 네트워크 인터페이스에 대한 IPv4 또는 IPv6 주소 목록을 가져오는 함수 (`GetInterfaceIP4Addrs()` 등)
- 인터페이스에 대한 기본 게이트웨이 인터페이스를 가져오는 함수 (` GetDefaultGatewayInterface()` 등)
- 주어진 IP 주소를 가진 네트워크 인터페이스를 가져오는 함수 (`GetInterfaceByIP()` 등)
- 기타 등등

### 주요 함수

### 1. GetInterfaceIP4Addrs()
```go
func GetInterfaceIP4Addrs(iface *net.Interface) ([]net.IP, error) {
	addrs, err := getIfaceAddrs(iface)
	if err != nil {
		return nil, err
	}

	ipAddrs := make([]net.IP, 0)

	// prefer non link-local addr
	ll := make([]net.IP, 0)

	for _, addr := range addrs {
		if addr.IP.To4() == nil {
			continue
		}

		if addr.IP.IsGlobalUnicast() {
			ipAddrs = append(ipAddrs, addr.IP)
			continue
		}

		if addr.IP.IsLinkLocalUnicast() {
			ll = append(ll, addr.IP)
		}
	}

	if len(ll) > 0 {
		// didn't find global but found link-local. it'll do.
		ipAddrs = append(ipAddrs, ll...)
	}

	if len(ipAddrs) > 0 {
		return ipAddrs, nil
	}

	return nil, errors.New("No IPv4 address found for given interface")
}
```
1. `getIfaceAddrs()` 함수를 이용하여 매개변수로 받은 `net.Interface 에 할당된 모든 IPv4 주소 목록을 가져온다.
2. `IsGlobalUnicast()` 함수를 이용하여 글로벌 유니캐스트 주소인지 확인한다.
   - 글로벌 유니캐스트 주소면 목록에 추가하고 다음 주소로(continue), 아니면 링크 로컬 유니캐스트 주소인지 확인한다.

### 2. GetDefaultGatewayInterface()
```go
func GetDefaultGatewayInterface() (*net.Interface, error) {
	routes, err := netlink.RouteList(nil, syscall.AF_INET)
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		if route.Dst == nil || route.Dst.String() == "0.0.0.0/0" {
			if route.LinkIndex <= 0 {
				return nil, errors.New("Found default route but could not determine interface")
			}
			return net.InterfaceByIndex(route.LinkIndex)
		}
	}

	return nil, errors.New("Unable to find default route")
}
```
1. `netlink.RouteList()` 함수를 이용해 라우팅 테이블 목록을 가져온다. 
   - `AF_INET` : IPv4 주소체계
2. 가져온 라우팅 정보를 for loop 로 돌리면서, 목적지가 `0.0.0.0/0` (기본 경로) 인 경로를 찾는다.
   - 시스템에서 사용하는 기본 게이트웨이를 의미한다.
3. 해당 인터페이스를 찾으면, 그 인터페이스가 flannel 이 패킷을 보내는 데 사용할 인터페이스 이다.

#### 왜 기본 게이트웨이를 찾을까?
- 노드에서 외부와 통신하기 위해 필요한 게이트웨이 정보를 알아내기 위해서
  - 외부와 통신하기 위해서는 통신을 중계해주는 게이트웨이가 필요한데, 이 게이트웨이의 주소는 일반적으로 `라우팅 테이블` 에 설정된다.
  - 보통 가장 기본적인 경로인 기본 게이트웨이(0.0.0.0/0) 으로 설정된다.

### 3. GetInterfaceByIP() 
```go
func GetInterfaceByIP(ip net.IP) (*net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		err := GetInterfaceIP4AddrMatch(&iface, ip)
		if err == nil {
			return &iface, nil
		}
	}

	return nil, errors.New("No interface with given IP found")
}
```
1. `net.Interfaces()` 함수를 통해 모든 인터페이스를 가져온다.
2. `GetInterfaceIPAddrMatch()` 함수를 통해 해당 인터페이스에 주어진 IPv4 주소와 일치하는 IP 가 있는 지 확인한다. 

## ipnet.go
> IP를 생성하고 파싱하는 코드 (IPv4 주소 처리에 대한 함수들)
```go
type IP4 uint32

func FromBytes(ip []byte) IP4 {
	return IP4(uint32(ip[3]) |
		(uint32(ip[2]) << 8) |
		(uint32(ip[1]) << 16) |
		(uint32(ip[0]) << 24))
}

func FromIP(ip net.IP) IP4 {
	ipv4 := ip.To4()

	if ipv4 == nil {
		panic("Address is not an IPv4 address")
	}

	return FromBytes(ipv4)
}
```
- 이런 식으로 IPv4 를 정의하고, 생성한다.
- 이외 코드 생략