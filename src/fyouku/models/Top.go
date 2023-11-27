package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// 获取动漫排行榜
func GetChannelTop(channelId int) []Video {
	var info []Video
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/video/channel/ranking?channelId=" + strconv.Itoa(channelId))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Data
	}

	return info
}

// 获取少女排行榜
func GetTypeTop(typeId int) []Video {
	var info []Video
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/video/type/ranking?typeId=" + strconv.Itoa(typeId) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Data
	}

	return info
}
