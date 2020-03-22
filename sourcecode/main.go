package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/icodeface/grdp"
	"github.com/icodeface/grdp/glog"
	flag "github.com/spf13/pflag"
	"golang.org/x/crypto/ssh"
)

type Conn struct {
	user string
	auth []ssh.AuthMethod
	addr string
}

var (
	client   *ssh.Client
	session  *ssh.Session
	password string
	host     string
	port     int
	choice   int
	user     string
	cmd      string
	help     bool
)

func init() {
	flag.StringVarP(&password, "password", "w", " ", "if use user & password to login; the password is for user.")
	flag.StringVarP(&host, "host", "s", " ", "remote host addr ip")
	flag.IntVarP(&port, "port", "p", 22, "remote host port")
	flag.StringVarP(&user, "user", "u", "root", "remote host user")
	flag.StringVarP(&cmd, "cmd", "c", "uname -a", "execute cmd in server")
	flag.BoolVarP(&help, "help", "h", false, "this help")
	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `Usage [gordpssh] [-s host] [-p port] [-u user] [-w password] [-c cmd]

 - ssh example:  
	1) gordpssh -u root -p 9527 -s 1.1.1.1 -w 123456 -c "uname -a"
	2) gordpssh --user root --port 9527 --host 1.1.1.1 --password 123456 --cmd "uname -a"
 - if use user & pasword; the password is for user.
 - rdp example is the same as ssh example.

Options:
`)
	flag.PrintDefaults()
}

func (c *Conn) SetConf() (err error) {
	c.addr = fmt.Sprintf("%s:%d", host, port)
	c.user = user
	c.auth = make([]ssh.AuthMethod, 0)
	var method ssh.AuthMethod
	method = ssh.Password(password)
	c.auth = append(c.auth, method)
	return nil
}

func (c *Conn) SetSession() (session *ssh.Session, err error) {
	client, err = ssh.Dial("tcp", c.addr, &ssh.ClientConfig{
		User: c.user,
		Auth: c.auth,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Millisecond * 500,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func rdpScan() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	client := grdp.NewClient(fmt.Sprintf("%s:%d", host, port), glog.DEBUG)
	err := client.Login(user, password)
	if err != nil {
		color.Red("[-] login failed,", err)
	} else {
		color.Green("[+] login success")
	}
} //3389端口指纹识别及登录

func sshScan() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	c := Conn{}
	err := c.SetConf()
	if err != nil {
		log.Fatalln(err)
	}
	session, err = c.SetSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	fmt.Println("[+] ssh command result is:")
	err = session.Run(cmd)
	if err == nil {
		fmt.Println("[+] host and ssh port is:", host, port)
	}
} //ssh端口识别及登录shell执行命令

func main() {
	flagtext, err := ioutil.ReadFile("flag.txt")
	if err != nil {
		panic(err.Error())
	}
	color.Magenta("----------------------------------------------------------")
	color.Red("%s\n", flagtext)
	color.Blue("{a light ssh and rdp management tool}")
	color.Magenta("----------------------------------------------------------")
	color.Yellow("[+] list of choices:")
	color.Yellow("1. ssh operation")
	color.Yellow("2. rdp operation")
	color.Green("[+] Use -h option then type any of [1] or [2] choice to get help messages.")
	color.Green("[+] when you finish type string, then choose choice to ssh or rdp.")
	color.Green("[+] now choose your choice to ssh or rdp")
	fmt.Scanf("%d", &choice)
	color.Magenta("- {your choice [%d]}\n", choice)
	time.Sleep(time.Second * 1)
	switch choice {
	case 1:
		sshScan()
	case 2:
		rdpScan()
	default:
		color.HiRed("...):you choice is not support...")
	}
}
