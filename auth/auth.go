package auth

import (
	"github.com/formeo/sshManyRunner/config"
	"golang.org/x/crypto/ssh"
)

type Auth struct {
	conf  *config.ConfStruct
	Hosts []string
}

func NewAuth(conf *config.ConfStruct) *Auth {
	return &Auth{
		conf:  conf,
		Hosts: nil,
	}
}

func (a *Auth) FillHosts() {
	for _, host := range a.conf.CmdConf.Aliases {
		if host.Enabled {
			a.Hosts = append(a.Hosts, host.Name+":"+host.Port)
		}
	}
}

func (a *Auth) connectToHost(host string) (*ssh.Client, *ssh.Session, error) {

	sshConfig := &ssh.ClientConfig{
		User: a.conf.CmdConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(a.conf.CmdConf.Password)},
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

func (a *Auth) RunCmd(nodeName string, command string, c chan string) {

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
