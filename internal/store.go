package internal

import (
	"context"
	"fmt"
	"log"
	"os"
)

var writeChan chan<- (*Ping)

func InitWriteChan(ctx context.Context) {
	ch := make(chan (*Ping), 10)
	writeChan = ch

	storeFile, err := os.OpenFile(cfg.Store, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error while opening file store:", err.Error())
	}
	defer storeFile.Close()

	for {
		select {
		case <-ctx.Done():
			return

		case ping := <-ch:
			data := fmt.Appendf(make([]byte, 0, 50), "%v,%v,%v\n", ping.URL, ping.LastPing, ping.OK)
			_, err := storeFile.Write(data)
			if err != nil {
				log.Println("Error while writing to store:", err.Error())
			}
		}
	}
}
