package handler

import (
	"github.com/gin-gonic/gin"
	llm "jieyuc-ai-proxy/aitools/llm/impl"
	"jieyuc-ai-proxy/request"
	"jieyuc-ai-proxy/response"
	"net/http"
)

func DeepseekExec(ctx *gin.Context) {
	req := &request.LLMProxyRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, response.BuildFailed("invalid params"))
	}

	llmInstance := llm.GetInstance()
	resp, err := llmInstance.Exec(req.Prompt)

	if err != nil {
		ctx.JSON(http.StatusOK, response.BuildFailed("llm execute error"))
		return
	}
	choices := resp.Choices
	ctx.JSON(http.StatusOK, response.BuildSuccess(choices))
}
