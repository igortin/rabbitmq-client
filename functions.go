package rabbit

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func checkArgs() {
	if len(os.Args) < 3 {
		log.Printf("Usage: %s [binding_key] [msg]", os.Args[0]) }
}





func GetCorelId() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa(rand.Intn(10000))
}

func GetRoutingKey(args []string) []string {
	return args[1:len(args)-1]
}

func GetEvent(args []string) string {
	return args[len(args)-1]
}
