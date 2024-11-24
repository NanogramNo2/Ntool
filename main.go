package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("Ntool ALPHA 1.0")

	//サーバー機能は後付け
	fmt.Println("CONNECTING SERVER...")
	ws, dialErr := websocket.Dial("ws://localhost:5103/auth", "", "http://localhost:5103/auth")
	if dialErr != nil {
		log.Fatal(dialErr)
	}

	defer ws.Close()
authagain:
	fmt.Print("TYPE PASSKEY:")
	var key string
	fmt.Scan(&key)
	sendRestMsg(ws, key)
	fmt.Println(" ")
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
			panic(recvErr)
		}
		fmt.Println("Receive : " + recvMsg)
		switch recvMsg {
		case "authorized":
			ws.Close()
			home()
		case "failed":
			fmt.Println("Auth : FAILED , TYPE PASSKEY AGAIN")
			goto authagain

		default:
			fmt.Println("Error : UNKNOWN ERROR , EXITING SYSTEM")
			panic("unknown error")
		}
	}
}
func home() {
home:
	cmd := exec.Command("clear") 
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(" ____________ ____________")
	fmt.Println("|1. Discord  |6. SearchUpd|")
	fmt.Println("|2. WebSocket|7.          |")
	fmt.Println("|3. Settings |            |")
	fmt.Println("|4. Help     |            |")
	fmt.Println("|5. Exit     |            |")
	fmt.Println(" -------------------------")

	fmt.Print("num>")
	var selNum string
	fmt.Scan(&selNum)
	switch selNum {
	case "1":
		cmd := exec.Command("clear") 
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(" _______________ ____________")
		fmt.Println("|1. Check User  |6. Exit     |")
		fmt.Println("|2. Send Message|            |")
		fmt.Println("|3. Typer       |            |")
		fmt.Println("|4. Login as Bot|            |")
		fmt.Println("|5. Return home |            |")
		fmt.Println(" ----------------------------")
	case "2":
		cmd := exec.Command("clear") 
		cmd.Stdout = os.Stdout
		cmd.Run()
		var dialurl string
		fmt.Print("Dial URL > ")
		fmt.Scan(&dialurl)
		var newdialurl string = "ws://" + dialurl
		var httpdialurl string = "http://" + dialurl
		ws, dialErr := websocket.Dial(newdialurl, "", httpdialurl)
		if dialErr != nil {
			fmt.Println("WebSocket : ERROR NOT CONNECTED")
			scanner := bufio.NewScanner(os.Stdin)
			for {
				scanner.Scan()
				var in string
				in = scanner.Text()
				fmt.Println("in: ", in)
				switch in {
				case "":
					goto home
				default:
					goto home
				}
			}

		}
		defer ws.Close()
		fmt.Println(" ")
		fmt.Println("WebSocket : CONNECTED")
		fmt.Println(" __________________")
		fmt.Println("|1. Send Request   |")
		fmt.Println("|2. Receive Request|")
		fmt.Println("|3. Disconnect and |")
		fmt.Println("|   Return home    |")
		fmt.Println(" ------------------")

	case "3":

	case "4":

	case "5":
		exitingFlag:
		fmt.Print("Do you really want to Exit?(y/n) > ")
		var exitFlag string
		fmt.Scan(&exitFlag)
		if(exitFlag == "y"){
			os.Exit(0)
		}else if (exitFlag == "n"){
			goto home
		}else {
			goto exitingFlag
		}
	case "6":

	default:

	}

}
func sendRestMsg(ws *websocket.Conn, msg string) {
	sendErr := websocket.Message.Send(ws, msg)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
}