package request

type LLMProxyRequest struct {
	Model        string `json:"model"`
	Prompt       string `json:"prompt"`
	OutputFormat string `json:"output_format"`
}
