package controllers

import (
	"encoding/json"
	"fyoukuapi/models"
	"github.com/siddontang/go/websocket"
)

type WsData struct {
	CurrentTime int
	EpisodesId  int
}

type BarrageController struct {
	BaseController
}

// 列表弹幕
func (this *BarrageController) WsFunc() {
	var (
		conn     *websocket.Conn
		err      error
		data     []byte
		barrages []models.BarrageData
	)
	if conn, err = websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil); err != nil {
		goto ERR
	}

	////心跳
	//go func() {
	//	for {
	//		if err = conn.WriteMessage(websocket.TextMessage, []byte("心跳")); err != nil {
	//			return
	//		}
	//		time.Sleep(5 * time.Second)
	//	}
	//}()

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal(data, &wsData)
		endTime := wsData.CurrentTime + 60
		//获取弹幕数据
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err != nil {
			goto ERR
		}
		msgJSON, err := json.Marshal(barrages)
		if err != nil {
			goto ERR
		}
		err = conn.WriteMessage(websocket.TextMessage, msgJSON)
		if err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

// 保存弹幕
func (this *BarrageController) SaveBarrage() {
	var (
		episodesId  int
		videoId     int
		uid         int
		content     string
		currentTime int
	)

	episodesId, _ = this.GetInt("episodesId", 0)
	if episodesId == 0 {
		this.JsonResult(1, "必须指定评论视频集数")
	}
	videoId, _ = this.GetInt("videoId", 0)
	if videoId == 0 {
		this.JsonResult(1, "必须指定评论视频")
	}
	uid, _ = this.GetInt("uid", 0)
	if uid == 0 {
		this.JsonResult(1, "请先登录")
	}
	content = this.GetString("content", "")
	if content == "" {
		this.JsonResult(1, "内容不能为空")
	}
	currentTime, _ = this.GetInt("currentTime", 1)
	err := models.Savebarrage(episodesId, videoId, uid, content, currentTime)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "保存成功")
}
