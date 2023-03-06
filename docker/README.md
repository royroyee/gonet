# Docker
Docker 네트워크에 대해서

### 사전 지식
 [Linux Network](https://github.com/royroyee/gonet/tree/main/linux)
 / [NIC](https://github.com/royroyee/gonet/tree/main/03-layer/01-physical-layer#lan-cardnic)
 /  [iptables](https://github.com/royroyee/gonet/blob/main/linux/README.md#iptables)
 /  [Bridge](https://github.com/royroyee/gonet/blob/main/linux/README.md#bridge)
 / [eth0](https://github.com/royroyee/gonet/blob/main/linux/README.md#eth0)
 / [veth]()


## Docker Network 구조

![docker network.png](..%2F..%2F..%2FPictures%2Fdocker%20network.png)


### docker0
호스트의 eth0 네트워크 인터페이스와 직접적으로 연결되는 **가상 네트워크 인터페이스**
- 컨테이너끼리 통신하려면 가상의 네트워크가 필요하다. 때문에 도커는 가상의 네트워크 인터페이스인 `docker0`를 제공한다.
- 도커 컨테이너 간 통신 뿐만 아니라, 호스트와 컨테이너 사이의 네트워크 통신도 이 인터페이를 통해 이루어진다.

### docker0 작동 방식
1. 도커 엔진이 실행될 때 `docker0` 라는 가상 브릿지를 생성
   - 이 가상 브릿지는 물리적인 네트워크 카드(NIC)와 연결되어 있으며 가상 브릿지 내부에서 컨테이너들이 IP 주소를 할당받고 통신한다.


2. 컨테이너가 실행될 때, 도커 엔진은 해당 컨테이너를 위한 가상 네트워크 인터페이스(veth)를 생성한다.
   - 이 `veth` 는 호스트의 네트워크 인터페이스(eth0)와 짝을 이루며, `veth`한 쪽은 컨테이너 내부에 할당된다. (위 그림 참고)


3. 컨테이너가 네트워크 통신을 하면, veth의 컨테이너 내부 쪽에서 전송된 패킷은 호스트의 가상 네트워크 인터페이스를 통해 docker0 가상 브리지로 전송된다. 이후 docker0 가상 브리지 내부에서 패킷이 라우팅되어 다른 컨테이너로 전달된다.


4. 컨테이너가 호스트와 네트워크 통신을 할 경우, docker0 가상 브리지를 통해 호스트의 물리적 네트워크 카드와 연결되어 있는 라우터나 스위치 등을 거쳐 외부와 통신한다.

### docker0 확인하기
```
$ ifconfig

docker0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        inet6 fe80::42:bbff:fed3:5cd9  prefixlen 64  scopeid 0x20<link>
        ether 02:42:bb:d3:5c:d9  txqueuelen 0  (Ethernet)
        RX packets 213809  bytes 13175722 (13.1 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 394838  bytes 522512946 (522.5 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```