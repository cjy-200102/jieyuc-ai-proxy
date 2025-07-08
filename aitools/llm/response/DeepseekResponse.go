package response

import (
	"encoding/json"
	"jieyuc-ai-proxy/aitools/llm/vo"
)

type DeepseekResponse struct {
	Id      string                `json:"id"`
	Object  string                `json:"object"`
	Created int                   `json:"created"`
	Model   string                `json:"model"`
	Choices []*DeepseekChoiceInfo `json:"choices"`
	Usage   *DeepseekUsage        `json:"usage"`
}

type DeepseekChoiceInfo struct {
	Index        int                             `json:"index"`
	Message      *vo.LargeLanguageModelMessageVo `json:"message"`
	FinishReason string                          `json:"finish_reason"`
}

type DeepseekUsage struct {
	PromptTokens          int `json:"prompt_tokens"`
	CompletionTokens      int `json:"completion_tokens"`
	TotalTokens           int `json:"total_tokens"`
	PromptCacheHitTokens  int `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"`
}

func ParseDeepseekResponse(respJson string) (resp *DeepseekResponse) {
	resp = &DeepseekResponse{}
	bytes := []byte(respJson)
	err := json.Unmarshal(bytes, resp)
	if err != nil {
		return nil
	}
	return
}

func (resp *DeepseekResponse) String() string {
	respJson, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(respJson)
}
