package main

import (
	"os"
	"fmt"
	"time"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	for ; ; {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
	}

}
