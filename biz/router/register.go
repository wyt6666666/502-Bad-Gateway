// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"douyin/biz/router/messageapi"
	"douyin/biz/router/relationapi"
	userapi "douyin/biz/router/userapi"
	videoapi "douyin/biz/router/videoapi"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	userapi.Register(r)
	videoapi.Register(r)
	messageapi.Register(r)
	relationapi.Register(r)
}
