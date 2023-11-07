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
	//ns := web.NewNamespace("/v1",
	//	web.NSNamespace("/object",
	//		web.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	web.NSNamespace("/user",
	//		web.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//)
	//web.AddNamespace(ns)

	web.Router("/user/register", &controllers.UserController{}, "put:UserRegister")
	web.Router("/user/login", &controllers.UserController{}, "get:UserLogin")

	web.Router("/channel/advert", &controllers.VideoController{}, "get:ChannelAdvert")
	web.Router("/channel/hot", &controllers.VideoController{}, "get:ChannelHotList")
	web.Router("/channel/recommend/region", &controllers.VideoController{}, "get:ChannelRegionRecommendList")
	web.Router("/channel/recommend/type", &controllers.VideoController{}, "get:ChannelTypeRecommendList")
	web.Router("/channel/video", &controllers.VideoController{}, "get:ChannelVideo")
	web.Router("/channel/region", &controllers.VideoController{}, "get:ChannelRegion")
	web.Router("/channel/type", &controllers.VideoController{}, "get:ChannelType")

	web.Router("/video/info", &controllers.VideoController{}, "get:VideoInfo")
	web.Router("/video/episodes/list", &controllers.VideoController{}, "get:VideoEpisodesList")

	web.Router("/comment/list", &controllers.CommentController{}, "get:CommentList")
	web.Router("/comment/save", &controllers.CommentController{}, "get:CommentSave")
}
