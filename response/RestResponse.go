package response

const codeOfSuccess = 0
const codeOfFailed = -1000

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BuildSuccess(date interface{}) *RestResponse {
	resp := &RestResponse{
		Code: codeOfSuccess,
		Data: date,
		Msg:  "success",
	}
	return resp
}

func BuildFailed(msg string) *RestResponse {
	resp := &RestResponse{
		Code: codeOfFailed,
		Msg:  msg,
	}
	return resp
}
