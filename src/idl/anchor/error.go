package anchor

type ErrorCode struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Msg  string `json:"msg,omitempty"`
}
