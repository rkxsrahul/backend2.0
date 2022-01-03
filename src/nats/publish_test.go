package nats

import (
	"flag"
	"log"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
)

// func TestPublish(t *testing.T) {
// 	data := RequestData{
// 		Method:  "",
// 		URL:     "",
// 		UUID:    "",
// 		Status:  "",
// 		Subject: "",
// 	}
// 	body, _ := json.Marshal(data)
// 	Publish(body, "")
// }

// func FirstCheck(t *testing.T) {
// 	log.Println(config.Conf.Service.Build, "eiwubd")
// }

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// setup for reading flags for deciding whether to do configuration with toml or env variables
	conf := flag.String("conf", "environment", "set configuration from toml file or environment variables")

	flag.Parse()

	log.Println(conf)
}

func TestInitConnection(t *testing.T) {
	log.Println(config.Conf.Service.Build, "eiwubd")
}
