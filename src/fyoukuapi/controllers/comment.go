package controllers

import (
	"fyoukuapi/models"
	"time"
)

type CommentController struct {
	BaseController
}

// 获取评论列表
func (this *CommentController) CommentList() {
	var (
		episodesId int
		page       int
		limit      int
	)

	episodesId, _ = this.GetInt("episodesId", 0)
	if episodesId == 0 {
		this.JsonResult(1, "episodesId不能为空")
	}
	page, _ = this.GetInt("page", 0)
	limit, _ = this.GetInt("limit", 12)

	commentList, count, err := models.GetCommentList(episodesId, page, limit)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}

	for k, _ := range commentList {
		commentList[k].AddTimeTitle = time.Unix(commentList[k].AddTime, 0).Format("2006-01-02 15:04:05")
	}

	data := map[string]interface{}{
		"list":  commentList,
		"count": count,
	}
	this.JsonResult(0, "操作成功", data)
}

// 保存评论
func (this *CommentController) CommentSave() {
	var (
		uid        int
		videoId    int
		episodesId int
		content    string
	)

	uid, _ = this.GetInt("uid", 0)
	if uid == 0 {
		this.JsonResult(1, "请先登录")
	}
	videoId, _ = this.GetInt("videoId", 0)
	if videoId == 0 {
		this.JsonResult(1, "必须指定评论视频")
	}
	episodesId, _ = this.GetInt("episodesId", 0)
	if episodesId == 0 {
		this.JsonResult(1, "必须指定评论视频集数")
	}
	content = this.GetString("content", "")
	if content == "" {
		this.JsonResult(1, "必须指定评论视频集数")
	}

	comment, err := models.CommentSave(uid, videoId, episodesId, content)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "操作成功", comment)
}
