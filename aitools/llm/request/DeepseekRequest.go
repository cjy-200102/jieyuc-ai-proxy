package request

import (
	"encoding/json"
	"jieyuc-ai-proxy/aitools/llm/vo"
)

type DeepseekRequest struct {
	Model    string                            `json:"model"`
	Messages []*vo.LargeLanguageModelMessageVo `json:"messages"`
}

func CreateDeepseekRequestParams(prompt, modelName string) (reqJson string, err error) {
	message := &vo.LargeLanguageModelMessageVo{
		Role:    "user",
		Content: prompt,
	}
	messages := []*vo.LargeLanguageModelMessageVo{message}
	request := DeepseekRequest{
		Model:    modelName,
		Messages: messages,
	}
	jsonByte, err := json.Marshal(request)
	reqJson = string(jsonByte)
	return
}
