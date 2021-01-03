package module

import "net/http"

type Data interface {
	Valid() bool
}

// 希望类型值是不可变的，将字段名称的首地址全为小写
type Request struct {
	httpReq *http.Request
	depth uint32
}

//通过这个函数对内部字段进行初始化
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

// 获取内部字段的值
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

//获取内部字段的值
func (req *Request) Depth() uint32 {
	return req.depth
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

type Response struct {
	httpResp *http.Response
	depth uint32
}

func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

func (resp *Response) Depth() uint32 {
	return resp.depth
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

type Item map[string]interface{}

func (item Item) Valid() bool{
	return item != nil
}