package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1579966190)
	ginrpc.AddGenOne("Hello.Hello", "/block", []string{"post", "get"})
	ginrpc.AddGenOne("Hello.Hello2", "hello.hello2", []string{"post"})
	ginrpc.AddGenOne("Hello.Hello3", "hello.hello3", []string{"post"})
}
