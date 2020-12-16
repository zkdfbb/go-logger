package main

import (
	"github.com/zkdfbb/go-logger"
)

func main() {

	log := logger.NewLogger()

	apiConfig := &logger.ApiConfig{
		Url:        "http://127.0.0.1:8081/index.php",
		Method:     "GET",
		Headers:    map[string]string{},
		IsVerify:   false,
		VerifyCode: 0,
	}
	logger.Attach("api", logger.LOGGER_LEVEL_DEBUG, apiConfig)
	logger.SetAsync()

	logger.Emergency("this is a emergency log!")
	logger.Alert("this is a alert log!")

	logger.Flush()
}
