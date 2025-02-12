package dto

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   []string    `json:"error"`
}

func (r *Response) DefaultOK() {
	r.Code = 200
	r.Message = "success"
}

func (r *Response) DefaultCreated() {
	r.Code = 201
	r.Message = "created"
}

func (r *Response) DefaultBadRequest() {
	r.Code = 400
	r.Message = "Request body not match"
}

func (r *Response) DefaultUnauthorized() {
	r.Code = 401
	r.Message = "Unauthorized access"
}

func (r *Response) DefaultForbidden() {
	r.Code = 403
	r.Message = "Forbidden access"
}

func (r *Response) DefaultNotFound() {
	r.Code = 404
	r.Message = "Record not found"
}

func (r *Response) DefaultNotAllowed() {
	r.Code = 405
	r.Message = "Method not allowed"
}

func (r *Response) DefaultConflict() {
	r.Code = 409
	r.Message = "New data already exists"
}

func (r *Response) DefaultInternalError() {
	r.Code = 500
	r.Message = "Request failed, server error"
}
