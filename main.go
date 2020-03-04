package main

import (
	"flag"
	"fmt"
	"github.com/formeo/sshManyRunner/auth"
	"github.com/formeo/sshManyRunner/config"
)

func main() {
	conf, err := config.New("config.json")
	if err != nil {
		fmt.Println("Внимание! Файл настройки не найден")
		panic(err)
	}
	application := auth.NewAuth(conf)
	application.FillHosts()

	commandPtr := flag.String("command", "ifconfig", "please write a command")
	for _, host := range application.Hosts {
		c := make(chan string)
		fmt.Println("for node: ", host)
		go application.RunCmd(host, *commandPtr, c)
		fmt.Println("main function message: ", <-c)
	}
}
