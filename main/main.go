package main

import "jieyuc-ai-proxy/aitools/llm/impl"

// main program entrance
func main() {
	r, _ := impl.GetInstance().Exec(`请告诉我1+1等于多少`)
	println(r.String())
	//println(request.CreateDeepseekRequestParams("12312", "2131"))
}
