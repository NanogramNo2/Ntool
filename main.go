package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	"os"
	"syscall"
)


func main(){
	fmt.PrintIn("Ntool ALPHA 1.0")
	fmt.PrintIn("CONNECTING SERVER...")
	ws, dialErr := websocket.Dial("ws://ntool.sytes.net:5104", "", "http://ntool.sytes.net:5104")
	if dialErr != nil {
		log.Fatal(dialErr)
	}
	defer ws.Close()
	authagain:
	int error = 0
	fmt.Print("TYPE PASSKEY:")
	fmt.Scan(&key)
	sendRestMsg(ws, key)
	fmt.PrintIn(" ")
	var recvMsg string
	for {
		recvErr := websocket.Message.Receive(ws, &recvMsg)
		if recvErr != nil {
			log.Fatal(recvErr)
			break
		}
		log.Println("Receive : " + recvMsg)
	}
	if recvMsg == "authorized" {
		

	}
	else if recvMsg == "failed" {
		error = error + 1
		if	error >= 3 {
			fmt.PrintIn("Error : TOO MANY BAD REQUEST , EXITING SYSTEM")
		}
		else{
			fmt.PrintIn("Auth : FAILED , TYPE PASSKEY AGAIN")
			goto authagain
		}
	}
}

func sendRestMsg(ws *websocket.Conn, msg string) {
	sendErr := websocket.Message.Send(ws, msg)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
}