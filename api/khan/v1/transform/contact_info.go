package transform

/*
*
{"appid":"wx_WijJeN12eAvakvBXkh7m3","to_wxid_list":["zdgougou"]}
*/
type ContactInfoRequest struct {
	Appid      string   `json:"appid"`
	ToWxidList []string `json:"to_wxid_list"`
}

/*
*
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{"string":""}},"contactList":[{"username":"zdgougou","ret":0,"contact":{"UserName":{"string":"zdgougou"},"NickName":{"string":"???"},"PyInitial":{"string":"J"},"QuanPin":{"string":"jian"},"Sex":2,"ImgBuf":{"iLen":0},"BitMask":4294967295,"BitVal":3,"ImgFlag":2,"Remark":{"string":""},"RemarkPyinitial":{"string":""},"RemarkQuanPin":{"string":""},"Province":"Shaanxi","City":"Xi'an","HasWeiXinHdHeadImg":1,"VerifyFlag":0,"Alias":"zdgougou","SnsUserInfo":{"SnsFlag":1},"Country":"CN","BigHeadImgUrl":"https://wx.qlogo.cn/mmhead/ver_1/L1z08xST8nlpNdM5KhFuNMKqWG5qDFOs3iauwibwEead5jDP1S3wCNyzuUMtqDiajmTfIfy2mkvqyo3qvqUnWDowMdAzoJNwRyd00iaYJzvhNOY/0","SmallHeadImgUrl":"https://wx.qlogo.cn/mmhead/ver_1/L1z08xST8nlpNdM5KhFuNMKqWG5qDFOs3iauwibwEead5jDP1S3wCNyzuUMtqDiajmTfIfy2mkvqyo3qvqUnWDowMdAzoJNwRyd00iaYJzvhNOY/132"}}]}}
*/
type ContactInfoResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		ContactList []struct {
			Username string `json:"username"`
			Ret      int    `json:"ret"`
			Contact  struct {
				UserName struct {
					String string `json:"string"`
				} `json:"UserName"`
				NickName struct {
					String string `json:"string"`
				} `json:"NickName"`
				PyInitial struct {
					String string `json:"string"`
				} `json:"PyInitial"`
				QuanPin struct {
					String string `json:"string"`
				} `json:"QuanPin"`
				Sex    int `json:"Sex"`
				ImgBuf struct {
					ILen int `json:"iLen"`
				} `json:"ImgBuf"`
				BitMask int64 `json:"BitMask"`
				BitVal  int   `json:"BitVal"`
				ImgFlag int   `json:"ImgFlag"`
				Remark  struct {
					String string `json:"string"`
				} `json:"Remark"`
				RemarkPyinitial struct {
					String string `json:"string"`
				} `json:"RemarkPyinitial"`
				RemarkQuanPin struct {
					String string `json:"string"`
				} `json:"RemarkQuanPin"`
				Province           string `json:"Province"`
				City               string `json:"City"`
				HasWeiXinHdHeadImg int    `json:"HasWeiXinHdHeadImg"`
				VerifyFlag         int    `json:"VerifyFlag"`
				Alias              string `json:"Alias"`
				SnsUserInfo        struct {
					SnsFlag int `json:"SnsFlag"`
				} `json:"SnsUserInfo"`
				Country         string `json:"Country"`
				BigHeadImgUrl   string `json:"BigHeadImgUrl"`
				SmallHeadImgUrl string `json:"SmallHeadImgUrl"`
			} `json:"contact"`
		} `json:"contactList"`
	} `json:"data"`
}
