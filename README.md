### gordpssh

``` bash
----------------------------------------------------------
 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/
{a light ssh and rdp management tool}
----------------------------------------------------------
```
 - By Lucifer11
 - My QQ:1185151867
 - [My Blog](https://fdlucifer.github.io/)

### help

``` bash
D:\Go\testgofiles\oldboygo\hackingwithgolang\gordpssh>gordpssh.exe -h
...
1
- {your choice [1]}
Usage [gordpssh] [-s host] [-p port] [-u user] [-w password] [-c cmd]

 - ssh example:
        1) gordpssh -u root -p 9527 -s 1.1.1.1 -w 123456 -c "uname -a"
        2) gordpssh --user root --port 9527 --host 1.1.1.1 --password 123456 --cmd "uname -a"
 - if use user & pasword; the password is for user.
 - rdp example is the same as ssh example.

Options:
  -c, --cmd string        execute cmd in server (default "uname -a")
  -h, --help              this help
  -s, --host string       remote host addr ip (default " ")
  -w, --password string   if use user & password to login; the password is for user. (default " ")
  -p, --port int          remote host port (default 22)
  -u, --user string       remote host user (default "root")
```

 - ![](/pics/help.jpg)

### ssh management

``` bash
D:\Go\testgofiles\oldboygo\hackingwithgolang\gordpssh>gordpssh.exe -s <ip addr> -p <port> -u <username> -w <password>
...
1
- {your choice [1]}
[+] ssh command result is:
Linux yisu-5de4c53d34764 3.10.0-327.el7.x86_64 #1 SMP Thu Nov 19 22:10:57 UTC 2015 x86_64 x86_64 x86_64 GNU/Linux
[+] host and ssh port is: 110.42.1.217 22
```

 - ![](/pics/ssh.jpg)

### rdp management

``` bash
D:\Go\testgofiles\oldboygo\hackingwithgolang\gordpssh>gordpssh.exe -s 39.104.21.7 -p 3389 -u adminstrator -w 12345678
...
2
- {your choice [2]}
39.104.21.7:3389
[DEBUG][x224 sendConnectionRequest 0ee000000000000100080003000000]
[DEBUG][tpkt Write 030000130ee000000000000100080003000000]
[DEBUG][GetBytes:  0300]
[DEBUG][tpkt recvHeader 0300 <nil>]
[DEBUG][tptk recvHeader FASTPATH_ACTION_X224, wait for recvExtendedHeader]
[DEBUG][GetBytes:  0013]
[DEBUG][tpkt recvExtendedHeader 0013 <nil>]
[DEBUG][tpkt wait recvData]
[DEBUG][GetBytes:  0ed00000123400020f080002000000]
[DEBUG][tpkt recvData 0ed00000123400020f080002000000 <nil>]
[DEBUG][x224 recvConnectionConfirm 0ed00000123400020f080002000000]
[+] Os Fingerprint is: [0ed00000123400020f080002000000]
[+] os version is 2012R2 OR 8
[INFO][TYPE_RDP_NEG_RSP]
[INFO][*** NLA Security selected ***]
[INFO][StartNLA]
[INFO][StartTLS]
[DEBUG][recvChallenge 30820102a003020106a181fa3081f73081f4a081f10481ee4e544c4d53535000020000001e001e003800000035828a623c69b461abc3866d00000000000000009800980056000000060380250000000f69005a003000330064006f00720066006f006d0079003700340072005a0002001e0069005a003000330064006f00720066006f006d0079003700340072005a0001001e0069005a003000330064006f00720066006f006d0079003700340072005a0004001e0069005a003000330064006f00720066006f006d0079003700340072005a0003001e0069005a003000330064006f00720066006f006d0079003700340072005a0007000800d71b5d074800d60100000000]
```

 - ![](/pics/rdp.jpg)

### 优点特色

 - 轻量的ssh，rdp管理客户端
 - 连接ssh端口并能使用-c执行linux命令
 - 连接rdp端口
 - 识别rdp端口所在主机的指纹判断主机操作系统类型
 - 持续更新

### QQ OR 微信

 - 1185151867
 - By Lucifer11 :)