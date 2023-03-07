# iptables
[iptables](https://github.com/royroyee/gonet/blob/main/linux/README.md#iptables) 를 이용하여 서브넷에 대한 라우팅 규칙을 생성

### iptables.go
> iptables를 사용하여 네트워크 패킷을 필터링 및 변환
```go
type IPTables interface {
	AppendUnique(table string, chain string, rulespec ...string) error
	ChainExists(table, chain string) (bool, error)
	ClearChain(table, chain string) error
	Delete(table string, chain string, rulespec ...string) error
	Exists(table string, chain string, rulespec ...string) (bool, error)
}
```
go언어의 `iptables` Package 를 이용.

자세한 내용 : https://pkg.go.dev/github.com/coreos/go-iptables@v0.5.0/iptables
- `AppendUnique()` : 특정 iptables 테이블의 특정 체인에 규칙을 추가
  - 이미 규칙이 존재하는 경우 중복으로 추가하지 않음
- `ChainExists()` : 특정 iptables 테이블의 특정 체인이 존재하는 지 확인
- `ClearChain()` : 특정 iptables 테이블의 특정 체인에서 모든 규칙을 제거
- `Delete()`: 특정 iptables 테이블의 특정 체인에서 규칙을 제거
- `Exists()` : 특정 iptables 테이블의 특정 체인에서 특정 규칙이 존재하는 지 확인