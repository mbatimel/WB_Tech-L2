package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct{
	host string
	port string
	connectionTime time.Duration
}

func NewClientConnection(host string, port string, cTime string) *Client{
	t, err := strconv.Atoi(cTime)
	if err != nil{
		log.Fatal(err)
	}
	return &Client{
		host: host,
		port: port,
		connectionTime: time.Duration(t) * time.Second,
	}
}

func call(client *Client, ctx context.Context, out context.CancelFunc){
	conn, err := net.DialTimeout("tcp", client.host+":"+client.port, client.connectionTime)
	if err != nil {
		log.Fatal(err)
	}
	defer out()
	for {
		select{
		case <-ctx.Done():
			_, err := fmt.Fprintf(conn, "time is up")
			if err != nil {
				log.Print(err)
			}
			log.Printf("time is up...")
			err = conn.Close()
			if err != nil {
				log.Print(err)
			}
			return
		default:
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("message: ")
			text, err := rd.ReadString('\n')
			if err != nil {
				log.Print("read error: ")
			}
			_, err = fmt.Fprintf(conn, text+"\n")
			if err != nil {
				log.Print(err)
			}
			fb, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Print(err)
			}
			fmt.Println("from server: ",fb)
			
		}
	}
	
}
func main(){
	ln, _ := net.Listen("tcp", "localhost:3000")
	conn, _ := ln.Accept()
	//временный сервер для подключения
	MyHandle(conn)

	err := ln.Close()
	if err != nil {
		log.Print(err)
	}
	timeOut := flag.String("timeout", "10", "time to work with the server")
	flag.Parse()
	args :=flag.Args()
	if len(args) <2 {
		log.Fatal("should be host and port")
	}
	port := args[0]
	host := args[1]
	client := NewClientConnection(host, port, *timeOut)
	ctx := context.Background()
	ctx, out := context.WithTimeout(ctx, client.connectionTime)
	call(client, ctx, out)
	fmt.Println("out connection")
	
}
func MyHandle(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		switch message {
		case "time is up":
			return
		default:
			fmt.Print("Message Received:", message)
			newMessage := strings.ToUpper(message)
			_, err := conn.Write([]byte(time.Now().String() + " " + newMessage + "\n"))
			if err != nil {
				log.Print(err)
			}
		}
	}
}