package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"os"

	"github.com/robfig/cron"
)

var (
	address string
	port string
	logpath string
	status string
	every string
)

func init()  {
	flag.StringVar(&address, "address", "www.growingio.com", "check's IP or server name")
	flag.StringVar(&port, "port", "80", "check's port")
	flag.StringVar(&every, "every", "30s", "check time intervals")
	flag.StringVar(&logpath, "log", "check.log", "check's log file path")

	log.SetFlags(log.Ldate|log.Lmicroseconds|log.LUTC)
}

func checkConn()  {
	conn, err := net.DialTimeout("tcp", address + ":" + port, 5 * time.Second)
	if err != nil {
		status = "Unreachable"
		log.Println("Connection error:", err.Error())
	} else {
		status = "Online"
		log.Println("Connection Success:", conn.RemoteAddr())
		defer conn.Close()
	}
}

func main()  {
	flag.Parse()
	logfile, err := os.OpenFile(logpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.SetOutput(logfile)

	c := cron.New()
	spec := "@every " + every
	err = c.AddFunc(spec, func() {
		checkConn()
	})
	if err != nil {
		fmt.Println("cron error:", err.Error())
	}
	c.Start()
	defer c.Stop()
	select{}
}