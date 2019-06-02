package main

import (
	"flag"
	"fmt"
	"github.com/formeo/sshManyRunner/config"
	"golang.org/x/crypto/ssh"
)

type Auth struct {
	User  string
	Pass  string
	Hosts []string
}

func NewAuth(filename string) (a *Auth, err error) {
	conf, err := config.New(filename)
	if err != nil {
		return nil, err
	}
	a = new(Auth)
	a.User = conf.CmdConf.Username
	a.Pass = conf.CmdConf.Password

	for _, host := range conf.CmdConf.Aliases {
		if host.Enabled {
			a.Hosts = append(a.Hosts, host.Name+":"+host.Port)
		}
	}
	return a, nil
}

func (a *Auth) runCmd(nodeName string, command string, c chan string) {

	client, session, err := a.connectToHost(nodeName)
	if err != nil {
		c <- err.Error()
		return
	}

	out, err := session.Output(command)
	if err != nil {

		c <- err.Error()
		return

	}

	err = client.Close()
	if err != nil {

		c <- err.Error()
		return

	}
	c <- string(out)
	return

}

func main() {
	a, err := NewAuth("config.json")
	if err != nil {
		fmt.Println("Внимание! Файл настройки не найден")
		panic(err)
	}
	commandPtr := flag.String("command", "ifconfig", "please write a command")
	for _, host := range a.Hosts {
		c := make(chan string)
		fmt.Println("for node: ", host)
		go a.runCmd(host, *commandPtr, c)
		fmt.Println("main function message: ", <-c)
	}
}

func (a *Auth) connectToHost(host string) (*ssh.Client, *ssh.Session, error) {

	sshConfig := &ssh.ClientConfig{
		User: a.User,
		Auth: []ssh.AuthMethod{ssh.Password(a.Pass)},
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		err := client.Close()
		if err != nil {
			return nil, nil, err
		}

		return nil, nil, err
	}

	return client, session, nil
}
