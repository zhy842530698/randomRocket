package randomRocket

import (
	"encoding/binary"
	"net"
)

var CONNECT_SUCCESS = []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

type LsServer struct {
	Cipher     *Cipher
	ListenAddr *net.TCPAddr
}

// 新建一个服务端
// 服务端的职责是:
// 1. 监听来自本地代理客户端的请求
// 2. 解密本地代理客户端请求的数据，解析 SOCKS5 协议，连接用户浏览器真正想要连接的远程服务器
// 3. 转发用户浏览器真正想要连接的远程服务器返回的数据的加密后的内容到本地代理客户端
func NewLsServer(passwd string, listenAddr string) (*LsServer, error) {
	bsPasswd, err := ParsePassword(passwd)
	if err != nil {
		return nil, err
	}
	lis, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	return &LsServer{
		NewCipher(bsPasswd),
		lis,
	}, nil
}
func (l *LsServer) Listen(didListen func(lis *net.TCPAddr)) error {
	return ListenEncryptedTCP(l.ListenAddr, l.Cipher, nil, didListen)
}
func (l *LsServer) handleConn(lconn *SecureTCPConn) {
	defer lconn.Close()
	buf := make([]byte, 256)
	_, err := lconn.DecodeRead(buf)
	//非法协议
	if err != nil || buf[0] != 0x05 {
		return
	}
	lconn.EncodeWrite([]byte{0x05, 0x00})
	n, err := lconn.DecodeRead(buf)
	if err != nil || n < 7 {
		return
	}
	//cmd代表客户端想要的链接类型，值长度为1，
	if buf[1] != 0x01 {
		return
	}
	var dIP []byte
	//按照dtype区分，代表请求远程服务器的类型，值长度为一个字节
	switch buf[3] {
	case 0x01:
		//	IPV4
		dIP = buf[4 : 4+net.IPv4len]
	case 0x03:
		//域名转化
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[5:n-2]))
		if err != nil {
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		dIP = buf[4 : 4+net.IPv6len]
	default:
		return
	}
	dPort := buf[n-2:]
	dstAddr := &net.TCPAddr{
		IP:   dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}
	//	作为客户端去访问远程服务
	dstServer, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		return
	}
	defer dstServer.Close()
	//清除所有数据
	dstServer.SetLinger(0)
	//告诉客户端链接成功了
	lconn.EncodeWrite(CONNECT_SUCCESS)
	//	进行转发
	go func() {
		err := lconn.DecodeCopy(dstServer)
		if err != nil {
			lconn.Close()
			dstServer.Close()
		}
	}()
	//这里会出现网络错误
	(&SecureTCPConn{
		Cipher:          lconn.Cipher,
		ReadWriteCloser: dstServer,
	}).EncodeCopy(lconn)
}
