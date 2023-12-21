package helpers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
)

func PrettyStruct(data interface{}) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	log.Info(string(val))
}
