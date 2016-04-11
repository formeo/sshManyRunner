package main

import (
	"fmt"	
	"golang.org/x/crypto/ssh"
	"github.com/formeo/sshManyRunner/config"
)
var cConf config.MyJsonName

var user string
var pass string

func init(){
	
	cConfs, err := config.New("config.json")
	
	if err != nil {
		panic(err)
	}
	cConf = cConfs		
	user = cConf.Cmdconf.Username
	pass = cConf.Cmdconf.Password
	
}



func runCmd(nodename string,c chan string){
	
	client, session, err := connectToHost(nodename)
	        if err != nil {
		    c <-err.Error()
			return
	    }	
	
	//out, err := session.CombinedOutput("top")
	out, err := session.Output("ifconfig")
	if err != nil {
	
		c <-err.Error()
		return
		
		
	}

	client.Close()
    c <- string(out)
	return
	

	
}


func main() {
	fmt.Println("Start")
	
     for _, ms := range cConf.Cmdconf.Aliases{
        c := make(chan string)    
		fmt.Println("for node: ",ms.Name)    
        go runCmd(ms.Name+":"+ms.Port,c)
		fmt.Println("main function message: ",<-c)
    
	
  
		 
}


	
}

func connectToHost(host string) (*ssh.Client, *ssh.Session, error) {
	
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}