package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var PetApi = new(Pet)

type Pet struct {
	ResApi
}

type PetCheckParam struct {
	Name string `p:"username" v:"required|参数为必填项"`
}

func (a *Pet) PetArchives(r *ghttp.Request) {
	var name = r.Get("name")
	var check *PetCheckParam
	//参数验证
	if err := r.Parse(&check); err != nil {
		r.Response.WriteJsonExit(ResApi{
			Code: 1,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	g.Log("name glog")
	r.Response.WriteJsonExit(ResApi{
		Code: 0,
		Msg:  "",
		Data: name,
	})

	//r.Response.WriteExit("petArchives")

}

func (a *Pet) FindPetInfo(r *ghttp.Request) {
	var name = r.Get("name")

	r.Response.Writeln(name)
}
