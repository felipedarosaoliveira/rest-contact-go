package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/felipedarosaoliveira/rest-contact-go/domain/contact"
	"github.com/felipedarosaoliveira/rest-contact-go/internal/server/http"
)

func main() {
	message := "Servidor Rest Go"
	fmt.Println(message)
	contactService := contact.NewService()
	handler := http.NewHandler(contactService)
	server := http.New("8080", handler)
	server.ListenAndServe()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shudown()
}
