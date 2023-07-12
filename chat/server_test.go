package chat

import (
	"testing"

	"github.com/vogo/logger"
)

func TestServer(t *testing.T) {
	appID := EnvString("APP_ID")
	apiKey := EnvString("API_KEY")
	apiSecret := EnvString("API_SECRET")

	if appID == "" || apiKey == "" || apiSecret == "" {
		t.Skip("APP_ID, API_KEY, API_SECRET are required")
	}

	s := NewServer(appID, apiKey, apiSecret)
	session, sessionErr := s.GetSession("123456789")

	if sessionErr != nil {
		panic(sessionErr)
	}

	questions := []string{
		"我想做一个AI相关的产品，请你以苏格拉底的方式对我进行提问，一次一个问题",
		"你的建议是什么？",
		"接下来我该怎么做？",
	}

	answer := ""

	var err error

	for _, question := range questions {
		logger.Info("question:", question)
		answer, err = session.Send(question)

		if err != nil {
			panic(err)
		}

		logger.Info("answer:", answer)
	}
}
