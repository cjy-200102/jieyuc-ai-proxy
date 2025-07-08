package impl

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"jieyuc-ai-proxy/aitools/llm/request"
	"jieyuc-ai-proxy/aitools/llm/response"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	instance = &DeepseekExecutor{}
)

const apiUrl = "https://api.deepseek.com/chat/completions"

type DeepseekExecutor struct {
	apiKey string
}

func init() {
	llmConfigPath := os.Getenv("LLM_CONFIG_LOCAL_PATH")
	viper.SetConfigFile(llmConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	instance.apiKey = viper.GetString("deepseek.api_key")
}

func GetInstance() *DeepseekExecutor {
	return instance
}

func (exec *DeepseekExecutor) buildHeader() (headers map[string][]string) {
	headers = make(map[string][]string)
	headers["Content-Type"] = []string{"application/json"}
	headers["Authorization"] = []string{fmt.Sprintf("Bearer %s", exec.apiKey)}
	return
}

// Exec deepseek执行
func (exec *DeepseekExecutor) Exec(prompt string) (*response.DeepseekResponse, error) {
	client := http.Client{Timeout: 60 * time.Second}
	params, err := request.CreateDeepseekRequestParams(prompt, "deepseek-chat")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(params))
	if err != nil {
		log.Println(err)
	}

	req.Header = exec.buildHeader()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response.ParseDeepseekResponse(string(body)), nil
}
