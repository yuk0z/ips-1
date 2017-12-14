package irc

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"time"
)

type Account struct {
	Username string
	FullName string
	Nick     string
}

type Client struct {
	Account Account
	Channel string
	Data    string
	Server  Server
}

type Command struct {
	name     string
	shortcut string
	handler  func(c *Client, msg []string, output chan string) bool
}

var (
	clientCommands []Command
)

func (c *Client) Login(nick string) bool {
	return Reg(c.Server.Conn, nick, c.Account.Username, c.Account.FullName)
}

func (c *Client) JoinChannel(channel string) bool {
	c.Channel = channel
	return Join(c.Server.Conn, channel)
}

func (c *Client) LeaveChannel() bool {
	return Part(c.Server.Conn, c.Channel)
}

func (c *Client) HandleData() bool {
	message, _ := bufio.NewReader(c.Server.Conn).ReadString('\r')
	if len(message) > 0 {
		fmt.Printf("Message:%s\n", message)
	}
	return true
}

func (c *Client) LogMessage(nick string, msg string) bool {
	fmt.Printf(
		"%s - [%s] <%s> %s\n",
		time.Now().String(), c.Channel, c.Account.Nick, msg)
	return true
}

func (c *Client) Close() bool {
	c.Server.Conn.Close()
	return true
}

func (c *Client) Auth(passwd string) bool {
	return Msg(c.Server.Conn, "NickServ", "IDENTIFY "+passwd)
}

func AddListener(wg sync.WaitGroup, reader io.Reader) {

}

func AddUserCommandListener(wg sync.WaitGroup, reader io.Reader) {

}

// Syntax: connect server-name port-number
func Connect(c *Client, msg []string, output chan string) bool {
	if len(msg) != 2 {
		output <- "Error: Wrong numbers of arguments"
		return false
	}

	c.Server = Server{
		Hostname: msg[0],
		Port:     msg[1],
	}
	if !c.Server.Connect() {
		output <- "Error: Cannot connect to " + c.Server.String()
		return false
	}

	output <- "Connected to " + c.Server.String()

	return true
}

func ListAvaiableServers(c *Client, msg string) {

}

func (c *Client) Run() bool {
	var wg sync.WaitGroup

	//if !client.Login(settings.UserData.Nickname) {
	//	os.Exit(1)
	//}

	//client.Auth(string(settings.UserData.Password))

	// Should wait NOTIFY message
	//time.Sleep(10 * time.Second)

	//client.JoinChannel(settings.UserData.Chat)

	//for client.HandleData() {
	//}

	//client.Close()

	//AddListener(wg, )

	return true
}

func (c *Client) Init() bool {
	clientCommands = []Command{
		{"CONNECT", "C", Connect},
	}
	return true
}
