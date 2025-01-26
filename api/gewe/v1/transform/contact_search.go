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
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{"string":"Everything is OK"}},"UserName":{"string":"xingmaogou"},"NickName":{"string":"气盖世の籍"},"Pyinitial":{"string":"wxid_6gc1cmd7y1fn22"},"QuanPin":{"string":"wxid_6gc1cmd7y1fn22"},"Sex":1,"ImgBuf":{"iLen":0},"Province":"Shaanxi","City":"Xi'an","Signature":"因为感恩，所以我一直铭记着那些帮助过我的人！","PersonalCard":1,"VerifyFlag":0,"WeiboFlag":0,"AlbumStyle":0,"AlbumFlag":0,"SnsUserInfo":{"SnsFlag":0,"SnsBgobjectId":0,"SnsFlagEx":0},"Country":"CN","CustomizedInfo":{"BrandFlag":0},"ContactCount":0,"BigHeadImgUrl":"http://wx.qlogo.cn/mmhead/ver_1/GZ4vqt7CSEfPuaj9Me3ibYSAqnpvIzA1iaiaFlw3iblElpd1G4Slib2s3uJdibicteic3D7wqVgyosMBxFLCqI8XMtrj3Ek4NTy9UACeYEaPWpUUQkPdzUptmaQLHOe4p4yFCibSE/0","SmallHeadImgUrl":"http://wx.qlogo.cn/mmhead/ver_1/GZ4vqt7CSEfPuaj9Me3ibYSAqnpvIzA1iaiaFlw3iblElpd1G4Slib2s3uJdibicteic3D7wqVgyosMBxFLCqI8XMtrj3Ek4NTy9UACeYEaPWpUUQkPdzUptmaQLHOe4p4yFCibSE/132","ResBuf":{"iLen":0},"MatchType":0}}
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
		Province     string `json:"Province"`
		City         string `json:"City"`
		Signature    string `json:"Signature"`
		PersonalCard int    `json:"PersonalCard"`
		VerifyFlag   int    `json:"VerifyFlag"`
		WeiboFlag    int    `json:"WeiboFlag"`
		AlbumStyle   int    `json:"AlbumStyle"`
		AlbumFlag    int    `json:"AlbumFlag"`
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
		MatchType int `json:"MatchType"`
	} `json:"data"`
}
