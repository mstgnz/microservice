package config

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error"`
	Data    any    `json:"data"`
}

func (r *Response) SetStatus(status bool) *Response {
	r.Status = status
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) SetError(error any) *Response {
	r.Error = error
	return r
}

func (r *Response) SetData(data any) *Response {
	r.Data = data
	return r
}
