package biz

import (
	"context"
	"myapp/internal/data"
)

type HelloRequest interface {
	data.Hello
	SetName(string)
	GetName() string
	BizHello(ctx context.Context, request HelloReq) HelloResponse
}

func NewHelloRequest(hello data.Hello) HelloRequest {
	return HelloRequest(&HelloReq{hello, ""})
}

type HelloReq struct {
	data.Hello
	name string
}

func (h *HelloReq) SetName(s string) {
	h.name = s
}

func (h *HelloReq) GetName() string {
	return h.name
}

func NewHelloReq(name string) HelloReq {
	return HelloReq{name: name}
}

type HelloResponse interface {
	GetMessage() string
}

type helloResponse struct {
	message string
}

func (h *helloResponse) GetMessage() string {
	return h.message
}

func (h *HelloReq) BizHello(ctx context.Context, request HelloReq) HelloResponse {
	message := request.ReSendMessage(ctx, request.GetName())
	res := helloResponse{message}
	return HelloResponse(&res)
}
