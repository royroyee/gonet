# iptables
[iptables](https://github.com/royroyee/gonet/blob/main/linux/README.md#iptables) 를 이용하여 서브넷에 대한 라우팅 규칙을 생성

## iptables.go
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

### func `MasqRules()`
> Flannel 에서 사용하는 iptables 규칙들을 생성하는 함수

`cluster_cidrs` : 클러스터 내에서 사용되는 CIDR 블록 목록(IPv4)

```go
rules[0] = IPTablesRule{"nat", "-A", "POSTROUTING", []string{"-m", "comment", "--comment", "flanneld masq", "-j", "FLANNEL-POSTRTG"}}

rules[1] = IPTablesRule{"nat", "-A", "FLANNEL-POSTRTG", []string{"-m", "mark", "--mark", kubeProxyMark, "-m", "comment", "--comment", "flanneld masq", "-j", "RETURN"}}
```
- rules[0] : `POSTROUTING` 체인에 flannel의 masquerade 정책을 추가
- rules[1] : `FLANNEL-POSTRTG`체인에 kube-proxy에서 마크한 패킷을 무시하도록 RETURN 하는 정책을 추가
