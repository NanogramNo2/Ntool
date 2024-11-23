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
	
	//サーバー機能は後付け
	fmt.PrintIn("CONNECTING SERVER...")
	ws, dialErr := websocket.Dial("ws://localhost:8080", "", "http://localhost:8080")
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
	/*if key == "TestingKey"　{
		fmt.PrintIn("Auth : Authorized Passkey " + key)
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
	else {
		error = error + 1
		if	error >= 3 {
			fmt.PrintIn("Error : TOO MANY PASSKEY MISTAKE , EXITING SYSTEM")
		}
		else{
			fmt.PrintIn("Auth : FAILED , TYPE PASSKEY AGAIN")
			goto authagain
		}
	}*/
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
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()

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
	
	fmt.PrintIn{
		"
		 ____________ ____________
		|1. Discord  |6. SearchUpd|
		|2. WebSocket|7.          |
		|3. Settings |            |
		|4. Help     |            |
		|5. Exit     |            |
		 -------------------------
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