// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fyoukuapi/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {

	//用户注册
	web.Router("/user/register", &controllers.UserController{}, "put:UserRegister")
	//用户登录
	web.Router("/user/login", &controllers.UserController{}, "get:UserLogin")
	//发送消息
	web.Router("/user/send/message", &controllers.UserController{}, "post:UserSendMessage")

	//获取频道顶部广告
	web.Router("/channel/advert", &controllers.VideoController{}, "get:ChannelAdvert")
	//获取热播列表
	web.Router("/channel/hot", &controllers.VideoController{}, "get:ChannelHotList")
	//根据频道下的地区ID获取推荐视频
	web.Router("/channel/recommend/region", &controllers.VideoController{}, "get:ChannelRegionRecommendList")
	//按照类型获取推荐
	web.Router("/channel/recommend/type", &controllers.VideoController{}, "get:ChannelTypeRecommendList")
	//视频列表
	web.Router("/channel/video", &controllers.VideoController{}, "get:ChannelVideo")
	//获取频道下地区信息
	web.Router("/channel/region", &controllers.BaseController{}, "get:ChannelRegion")
	//获取频道下类型信息
	web.Router("/channel/type", &controllers.BaseController{}, "get:ChannelType")

	//视频详情
	web.Router("/video/info", &controllers.VideoController{}, "get:VideoInfo")
	//视频剧集列表
	web.Router("/video/episodes/list", &controllers.VideoController{}, "get:VideoEpisodesList")
	//频道排行榜
	web.Router("/video/channel/ranking", &controllers.VideoController{}, "get:ChannelRanking")
	//类型排行榜
	web.Router("/video/type/ranking", &controllers.VideoController{}, "get:TypeRanking")

	//评论列表
	web.Router("/comment/list", &controllers.CommentController{}, "get:CommentList")
	//保存评论
	web.Router("/comment/save", &controllers.CommentController{}, "post:CommentSave")

}
