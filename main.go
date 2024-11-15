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
	
	/* server auth
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
	*/
	fmt.PrintIn{
		"
		 ___________ ____________
		|1.         |6. SearchUpd|
		|2.         |			 |
		|3. Settings|            |
		|4. Help    |            |
		|5. Exit    |            |
		 ----------- ------------
		"
	}
	fmt.Print("num>")
	fmt.scan(&selNum)
	switch selNum {
	case 1:

	case 2:

	case 3:

	case 4:
		
	case 5:
		
	case 6:

	default:
		
	
	} 
	
}
/* 
func sendRestMsg(ws *websocket.Conn, msg string) {
	sendErr := websocket.Message.Send(ws, msg)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
}
*/