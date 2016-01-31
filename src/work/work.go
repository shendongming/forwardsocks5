package work

//工作节点
import (
	"log"
	"net"
	"os"
	"os/user"
	"fmt"
	"runtime"
	"io/ioutil"
	"bytes"
	"bufio"
	"crypto/sha1"
	socks5 "go-socks5"
	"encoding/hex"
	"strings"
	uuid "github.com/satori/go.uuid"
)

//读取文件的工具方法
func read_file(filename string) string {
	if _, err := os.Stat(filename); err != nil {
		return ""
	}
	buf, err2 := ioutil.ReadFile(filename)
	if err2 != nil {
		return ""
	}
	return string(buf)
}

//获取本地的会话id
func get_session_id() string {

	session_file := os.TempDir() + "/work-session-id"
	log.Println("read file:", session_file)
	session_id := read_file(session_file)
	if len(session_id) == 0 {
		log.Println("write session id:", session_id)
		session_id = uuid.NewV4().String()
		ioutil.WriteFile(session_file, []byte(session_id), 0644)
	}
	return session_id
}

func getInfo() []byte {

	var buf bytes.Buffer
	host, err := os.Hostname()
	if err != nil {
		log.Println("get HostName error:", err)

	}
	user_info, _ := user.Current()
	session_id := get_session_id()

	buf.WriteString(fmt.Sprintf("OS: %s\n", runtime.GOOS))
	buf.WriteString(fmt.Sprintf("go: %s\n", runtime.Version()))
	user_id := fmt.Sprintf("<%s>%s@%s", user_info.Name, user_info.Username, host)
	buf.WriteString(fmt.Sprintf("session: %s\n", session_id))
	buf.WriteString(fmt.Sprintf("user: %s\n", user_id))
	return buf.Bytes()
}
//连接服务器
func Connect(addr string, signKey string) {


	//todo
	//info := getInfo()

	log.Println("connect:", addr)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)

	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}


	for {


		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			log.Println("connect error", err.Error())
			return
		}
		log.Println("connect ok")
		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)

		hi_str, err := reader.ReadString('\n')
		if err != nil {
			log.Println("get error:", err.Error())
			return
		}
		log.Println("hi_str", hi_str)

		log.Printf("get hi:[%v]\n", hi_str)

		hi_sha1 := sha1.Sum([]byte(hi_str + signKey))
		hi_sha1_hex := hex.EncodeToString(hi_sha1[:])

		session_id := get_session_id()
		writer.WriteString(hi_sha1_hex + "\n")
		writer.WriteString(session_id + "\n")
		writer.Flush()

		//有连接就发送2个字节 go
		go_ret, err := reader.ReadString('\n')
		go_ret = strings.TrimRight(go_ret, "\n")
		log.Println("go ret:", go_ret)
		log.Printf("get go ret:[%v],[%v]\n", go_ret, go_ret[0:3])
		if go_ret[0:3] != "go:" {
			return
		}
		log.Printf("ok start go!\n")
		//协商完毕 ,进行 socks5 服务
		go run_socks5(conn, server)

	}
}
func run_socks5(conn net.Conn, socks *socks5.Server) {
	socks.ServeConn(conn)
}