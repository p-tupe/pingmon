package internal

import (
	"context"
	"fmt"
	"log"
	"os"
)

var storeFile *os.File

func InitStore(ctx context.Context) {
	var err error
	storeFile, err = os.OpenFile(cfg.Store, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error while opening file store:", err.Error())
	}
	defer storeFile.Close()
	<-ctx.Done()
}

func WriteToStore(ping *Ping) {
	entry := fmt.Appendf(make([]byte, 0, 50), "%v,%v,%v\n", ping.URL, ping.LastPing, ping.OK)
	_, err := storeFile.Write(entry)
	if err != nil {
		log.Println("Error while writing to store:", err.Error())
	}
}
