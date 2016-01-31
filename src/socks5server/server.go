package socks5server


import (
	"net"
	"log"
	"time"
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

const (
// 尝试创建处理器时的conn.read 的timeout
// 实际是一次性读取数据，所以这个超时指的是客户端必须10秒内发出第一和数据
	handlerNewTimeout = 10 * time.Second
	socks5Version = uint8(5)
// 默认一个连接的总处理时间，一般都会被实际的处理器修改掉。
	handlerBaseTimeout = 10 * time.Minute
)


type BridgeServer struct {

	Socks5Addr       string // TCP 监听地址 :1086
	WorkAddress      string // TCP 监听地址 :1088
	SignKey          string // 签名的密钥
	workConn         net.Listener
	socks5Conn       net.Listener

	read_socks5_chan chan net.Conn
							//s := make(chan string, 3)

}

//s := make(chan string, 3)

func NewServer(socks5Addr string, workAddress string, signKey string) *BridgeServer {

	server := &BridgeServer{}
	server.Socks5Addr = socks5Addr
	server.WorkAddress = workAddress
	server.SignKey = signKey
	server.read_socks5_chan = make(chan net.Conn, 100)

	return server
}

func (server *BridgeServer)ListenAndServe() error {


	log.Printf("log listen:%s", server.Socks5Addr)

	socks5Conn, err := net.Listen("tcp", server.Socks5Addr)
	if err != nil {

		return err
	}

	workConn, err := net.Listen("tcp", server.WorkAddress)
	if err != nil {

		return err
	}

	server.socks5Conn = socks5Conn
	server.workConn = workConn

	go server.ServerSocks5()
	server.ServerWork()
	return nil
}


//work 工作节点的守护
func (server *BridgeServer) ServerWork() error {

	conn := server.workConn
	defer conn.Close()
	for {
		rw, e := conn.Accept()

		if e != nil {
			log.Printf("Accept error: %v", e)
		}
		go server.handlerWorkConn(rw)
	}
	return nil
}


func (server *BridgeServer)handlerWorkConn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("work failed:", err)
		}
	}()
	defer conn.Close()

	//bufio.re

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	//检查签名
	hi_str := "hi:" + uuid.NewV4().String() + "\n"
	writer.WriteString(hi_str)
	writer.Flush()

	ok_ask := sha1.Sum([]byte(hi_str + server.SignKey))
	ok_ask_hex := hex.EncodeToString(ok_ask[:])
	ask, err := reader.ReadString('\n')
	if (err != nil) {
		fmt.Println("GET ERROR:", err.Error())
		return
	}
	ask = strings.TrimRight(ask, "\n")
	if ask != ok_ask_hex {
		fmt.Println("sign error!", conn.RemoteAddr())
		return
	}
	fmt.Println("sign ok!", conn.RemoteAddr())

	//获取会话id,客户端唯一标识
	session_id, error := reader.ReadString('\n')
	if error != nil {
		fmt.Println("get session id error")
		return
	}
	session_id = strings.TrimRight(session_id, "\n")

	fmt.Println("get session id:", session_id,conn.RemoteAddr())

	//等客户端链接
	socks5_client_conn := <-server.read_socks5_chan
	log.Printf("get a socks5 CONN:%v", socks5_client_conn.RemoteAddr())
	writer.WriteString(fmt.Sprintf("go:%v\n",socks5_client_conn.RemoteAddr()))
	writer.Flush()


	ch := make(chan int, 2)

	go func() {
		pipe_conn(socks5_client_conn, conn);
		ch <- 1
		log.Println("pipe socks5_client_conn->conn end")
	}()


	go func() {
		pipe_conn(conn, socks5_client_conn);
		ch <- 1
		log.Println("pipe conn->socks5_client_conn end")
	}()

	log.Println("wait copy all  end")
	//某个client客户端的数据全部给复制给 某个work的连接
	<-ch
	<-ch
	log.Println("copy all  end")

}

func pipe_conn(reader net.Conn, writer net.Conn) {
	for {
		log.Println("reader RemoteAddr addr", reader.RemoteAddr())
		log.Println("reader LocalAddr addr", reader.LocalAddr())
		log.Println("write RemoteAddr:", writer.RemoteAddr())
		log.Println("write LocalAddr:", writer.LocalAddr())
		buf := make([]byte, 1024)
		read_len, read_error := reader.Read(buf)
		if read_error != nil {
			log.Println("read socks5_client_conn data error:", read_error.Error())
			break
		}
		log.Println("read date is len:", read_len)
		if read_len == 0 {
			log.Println("read socks5_client_conn data len =0")
			break
		}
		write_len, write_error := writer.Write(buf[0:read_len])
		if write_error != nil {
			log.Println("write data error")
			break
		}
		log.Println("write date is len:", write_len)
		if write_len == 0 {
			log.Println("write date is len")
			break
		}
	}
}
//socks5 工作
func (server *BridgeServer) ServerSocks5() error {
	conn := server.socks5Conn
	defer conn.Close()

	for {
		log.Println("wait socks5 client")
		rw, e := conn.Accept()

		if e != nil {
			log.Printf("Accept error: %v", e)
		}
		go server.handlerSocks5Conn(rw)
		log.Println("read socks5 client next")
	}
}


//socks5进程
func (server *BridgeServer)handlerSocks5Conn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("work failed:", err)
		}
	}()
	log.Println("get socks5 client:", conn.RemoteAddr())
	server.read_socks5_chan <- conn

	log.Println("GET CONN OK send chan")

}