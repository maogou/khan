package transform

/*
*
{"appid":"wx_WijJeN12eAvakvBXkh7m3","user_name":"xingmaogou"}
*/
type ContactSearchRequest struct {
	Appid    string `json:"appid"`
	UserName string `json:"user_name"`
}

/*
*
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{"string":"Everything is OK"}},"UserName":{"string":"v3_020b3826fd03010000000000520c10334fe6b3000000501ea9a3dba12f95f6b60a0536a1adb6a45ca29c79e8efeb399e19ac72c3a7e26f7dd4407b1c3246692f4e5998ca85e69272f1e1a3138e18465b28cc3c@stranger"},"NickName":{"string":"???"},"Pyinitial":{"string":"wxid_7yr31zjouhv212"},"QuanPin":{"string":"wxid_7yr31zjouhv212"},"Sex":1,"ImgBuf":{"iLen":0},"Signature":"????????????????????????????????????","PersonalCard":0,"VerifyFlag":0,"WeiboFlag":0,"AlbumStyle":0,"AlbumFlag":0,"AlbumBgimgId":"919535318204416_919535318204416","SnsUserInfo":{"SnsFlag":0,"SnsBgobjectId":0,"SnsFlagEx":0},"Country":"CN","CustomizedInfo":{"BrandFlag":0},"ContactCount":0,"BigHeadImgUrl":"http://wx.qlogo.cn/mmhead/ver_1/QunmnGMbsqlBIp0fRAazuUrxjOSgGwtadhTv2CibFBAHqk6jkFjM1wqqPvQd3eK9IuaCyPy9ic0WQXicOduVXNxSicX3WPUyiaDN7c18ZpOv2edl10hVCPEuT2EWTrSyIOGfhibRU9Ay9P8sTGIiagJQ1uW4A/0","SmallHeadImgUrl":"http://wx.qlogo.cn/mmhead/ver_1/QunmnGMbsqlBIp0fRAazuUrxjOSgGwtadhTv2CibFBAHqk6jkFjM1wqqPvQd3eK9IuaCyPy9ic0WQXicOduVXNxSicX3WPUyiaDN7c18ZpOv2edl10hVCPEuT2EWTrSyIOGfhibRU9Ay9P8sTGIiagJQ1uW4A/132","ResBuf":{"iLen":0},"AntispamTicket":"v4_000b708f0b040000010000000000e8a963c43e1b21ec31f429d995671000000050ded0b020927e3c97896a09d47e6e9ecbd68f743b4377543f8a336c4dc06a39d3fddea4e84da8d75311b8eb2aa80e0c9168727afad98857fdda60227655b8bf49c6b5eefc2242fc4120e47299c5a80e1e6f93bbc7c6444d51@stranger","MatchType":0}}
*/

type ContactSearchResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		UserName struct {
			String string `json:"string"`
		} `json:"UserName"`
		NickName struct {
			String string `json:"string"`
		} `json:"NickName"`
		Pyinitial struct {
			String string `json:"string"`
		} `json:"Pyinitial"`
		QuanPin struct {
			String string `json:"string"`
		} `json:"QuanPin"`
		Sex    int `json:"Sex"`
		ImgBuf struct {
			ILen int `json:"iLen"`
		} `json:"ImgBuf"`
		Signature    string `json:"Signature"`
		PersonalCard int    `json:"PersonalCard"`
		VerifyFlag   int    `json:"VerifyFlag"`
		WeiboFlag    int    `json:"WeiboFlag"`
		AlbumStyle   int    `json:"AlbumStyle"`
		AlbumFlag    int    `json:"AlbumFlag"`
		AlbumBgimgId string `json:"AlbumBgimgId"`
		SnsUserInfo  struct {
			SnsFlag       int `json:"SnsFlag"`
			SnsBgobjectId int `json:"SnsBgobjectId"`
			SnsFlagEx     int `json:"SnsFlagEx"`
		} `json:"SnsUserInfo"`
		Country        string `json:"Country"`
		CustomizedInfo struct {
			BrandFlag int `json:"BrandFlag"`
		} `json:"CustomizedInfo"`
		ContactCount    int    `json:"ContactCount"`
		BigHeadImgUrl   string `json:"BigHeadImgUrl"`
		SmallHeadImgUrl string `json:"SmallHeadImgUrl"`
		ResBuf          struct {
			ILen int `json:"iLen"`
		} `json:"ResBuf"`
		AntispamTicket string `json:"AntispamTicket"`
		MatchType      int    `json:"MatchType"`
	} `json:"data"`
}
