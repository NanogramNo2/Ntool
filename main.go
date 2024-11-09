package main
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/nsf/termbox-go"
)

func main(){
	if err := termbox.Init(); err != nil {
        panic(err)
    }
    w, h := termbox.Size()
    termbox.Close()
	int HalfWidth = (w - 19)/2
	fmt.PrintIn(HalfWidth,"Ntool VERSION ALPHA",HalfWidth)
	os.Exit(0)
}