# Routing
flannel에서 가상 네트워크의 라우팅 기능을 수행하는 데 필요한 코드


### router.go
```go
// Route present a specific route
type Route struct {
	InterfaceIndex    int
	DestinationSubnet *net.IPNet
	GatewayAddress    net.IP
}
```
- `InterfaceIndex` : 라우팅 대상인 인터페이스의 인덱스
  - 인덱스는 시스템에서 인터페이스를 식별하는 데 사용
- `DestinationSubnet` : 목적지 서브넷의 IP 주소와 서브넷 마스크를 포함(`net.IPNet)
  - 이를 통해 라우터는 패킷의 목적지가 어떤 서브넷에 속하는지를 확인할 수 있다.
- `GatewayAddress` : 패킷이 목적지 서브넷 밖에 있을 때, 다음 라우터의 IP 주소
  - 이 정보를 통해 패킷이 다음 라우터로 전달되도록 한다.
  - [게이트웨이](https://github.com/royroyee/gonet/blob/main/03-layer/03-network-layer/README.md#%EA%B2%8C%EC%9D%B4%ED%8A%B8%EC%9B%A8%EC%9D%B4gateway)

```go
// Router manages network routes
type Router interface {
	// GetAllRoutes returns all existing routes
	GetAllRoutes() ([]Route, error)

	// GetRoutesFromInterfaceToSubnet returns all routes from the given Interface to the given subnet
	GetRoutesFromInterfaceToSubnet(interfaceIndex int, destinationSubnet *net.IPNet) ([]Route, error)

	// CreateRoute creates a new route
	CreateRoute(interfaceIndex int, destinationSubnet *net.IPNet, gatewayAddress net.IP) error

	// DeleteRoute removes an existing route
	DeleteRoute(interfaceIndex int, destinationSubnet *net.IPNet, gatewayAddress net.IP) error
}
```
- `GetAllRoutes()` : 사용 가능한 모든 라우팅 테이블 항목 반환
- `GetRoutesFromInterfaceToSubnet()` : 특정 인터페이스에서 목적지 서브넷으로 가는 모든 라우팅 테이블 항목 반환
- `CreateRoute()` : 새로운 라우팅 테이블 항목을 추가
- `DeleteRoute()` : 기존 라우팅 테이블 항목을 삭`