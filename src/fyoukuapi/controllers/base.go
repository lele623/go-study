package controllers

import (
	"compress/gzip"
	"encoding/json"
	"fyoukuapi/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"io"
	"strings"
)

type BaseController struct {
	web.Controller
}

// JsonResult 响应 json 结果
func (this *BaseController) JsonResult(code int, msg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		logs.Error(err)
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//this.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")//解决回退出现json的问题
	//使用gzip原始，json数据会只有原本数据的10分之一左右
	if strings.Contains(strings.ToLower(this.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		this.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		//gzip压缩
		w := gzip.NewWriter(this.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJSON)
		w.Flush()
	} else {
		io.WriteString(this.Ctx.ResponseWriter, string(returnJSON))
	}
	this.StopRun()
}

// 获取频道下的地区
func (this *BaseController) ChannelRegion() {
	var (
		channelId int
	)

	channelId, _ = this.GetInt("channelId", 0)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	channelType, res := models.GetChannelRegion(channelId)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", channelType)

}

// 获取频道类型
func (this *BaseController) ChannelType() {

}
