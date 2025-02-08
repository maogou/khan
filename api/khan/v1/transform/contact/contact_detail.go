package contact

/*
*
{"appid":"wx_WijJeN12eAvakvBXkh7m3","to_wxid_list":["xingmaogou","zdgougou"]}
*/
type ContactDetailRequest struct {
	Appid      string   `json:"appid"`
	ToWxidList []string `json:"to_wxid_list"`
}

/*
*
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{}},"ContactCount":2,"ContactList":[{"UserName":{"string":"xingmaogou"},"NickName":{"string":"???????????????"},"Pyinitial":{"string":"QGSJ"},"QuanPin":{"string":"qigaishiji"},"Sex":1,"ImgBuf":{"iLen":0},"BitMask":4294967295,"BitVal":3,"ImgFlag":3,"Remark":{},"RemarkPyinitial":{},"RemarkQuanPin":{},"ContactType":0,"RoomInfoCount":0,"DomainList":{},"ChatRoomNotify":0,"AddContactScene":0,"Province":"Shaanxi","City":"Xi'an","Signature":"??????????????????????????????????????????????????????????????????","PersonalCard":1,"HasWeiXinHdHeadImg":1,"VerifyFlag":0,"Level":0,"Source":9,"WeiboFlag":0,"AlbumStyle":0,"AlbumFlag":0,"SnsUserInfo":{"SnsFlag":1,"SnsBgimgId":"http://shmmsns.qpic.cn/mmsns/hAWvo7oLMKDUn7jJLlXcaXCb2SapTBrT28HXkdqiaAQQIeiatcW0ku6tiaMKRmKInfmSw9xvQHKmno/0","SnsBgobjectId":14179628696879706783,"SnsFlagEx":7297},"Country":"CN","BigHeadImgUrl":"https://wx.qlogo.cn/mmhead/ver_1/32mfqJXnYQbZaic4FwQu4M2FoGkDPy7P2o6zQPeerb0ia2HpRsOibB8YBxVRYzJsyy2WsQaw2Fd9OvHibCadskcZhicAPeZPFtdztrPZwLA9owgGPkH9Teic2mOyzxWYibylrof/0","SmallHeadImgUrl":"https://wx.qlogo.cn/mmhead/ver_1/32mfqJXnYQbZaic4FwQu4M2FoGkDPy7P2o6zQPeerb0ia2HpRsOibB8YBxVRYzJsyy2WsQaw2Fd9OvHibCadskcZhicAPeZPFtdztrPZwLA9owgGPkH9Teic2mOyzxWYibylrof/132","MyBrandList":"\u003cbrandlist count=\"0\" ver=\"826430565\"\u003e\u003c/brandlist\u003e","CustomizedInfo":{"BrandFlag":0},"HeadImgMd5":"41fee7249e7fca8daf994ba135ba7e06","EncryptUserName":"v3_020b3826fd0301000000000091a5214e167666000000501ea9a3dba12f95f6b60a0536a1adb6a45ca29c79e8efeb399e19ac72cf5b9e6455a3bc98294e597e2903a8a64e469114559167ee9ea7ee5f27937a41@stranger","AdditionalContactList":{"LinkedinContactItem":{}},"ChatroomVersion":0,"ChatroomMaxCount":0,"ChatroomAccessType":0,"NewChatroomData":{"MemberCount":0,"InfoMask":0},"DeleteFlag":0,"PhoneNumListInfo":{"Count":0},"ChatroomInfoVersion":0,"DeleteContactScene":0,"ChatroomStatus":0,"ExtFlag":0,"chatRoomBusinessType":0,"friendUserName":"xingmaogou","textStatusFlag":1,"ringBackSetting":{"finderObjectId":0,"startTs":0,"endTs":0},"bitMask2":4294967295,"bitValue2":0},{"UserName":{"string":"zdgougou"},"NickName":{"string":"???"},"Pyinitial":{"string":"J"},"QuanPin":{"string":"jian"},"Sex":2,"ImgBuf":{"iLen":0},"BitMask":4294967295,"BitVal":3,"ImgFlag":3,"Remark":{},"RemarkPyinitial":{},"RemarkQuanPin":{},"ContactType":0,"RoomInfoCount":0,"DomainList":{},"ChatRoomNotify":0,"AddContactScene":0,"Province":"Shaanxi","City":"Xi'an","Signature":"?????????????????????????????????","PersonalCard":1,"HasWeiXinHdHeadImg":1,"VerifyFlag":0,"Level":0,"Source":1000017,"Alias":"zdgougou","WeiboFlag":0,"AlbumStyle":0,"AlbumFlag":0,"SnsUserInfo":{"SnsFlag":1,"SnsBgimgId":"http://shmmsns.qpic.cn/mmsns/nGGS2R26WoDTejhrsLLpoBlqGLt8DibHRXoY5xlLDwVBytGAhfBod76j06mniaibqdibHodQqUQzbNM/0","SnsBgobjectId":14250985526642553423,"SnsFlagEx":6785},"Country":"CN","BigHeadImgUrl":"https://wx.qlogo.cn/mmhead/ver_1/L1z08xST8nlpNdM5KhFuNMKqWG5qDFOs3iauwibwEead5jDP1S3wCNyzuUMtqDiajmTfIfy2mkvqyo3qvqUnWDowMdAzoJNwRyd00iaYJzvhNOY/0","SmallHeadImgUrl":"https://wx.qlogo.cn/mmhead/ver_1/L1z08xST8nlpNdM5KhFuNMKqWG5qDFOs3iauwibwEead5jDP1S3wCNyzuUMtqDiajmTfIfy2mkvqyo3qvqUnWDowMdAzoJNwRyd00iaYJzvhNOY/132","MyBrandList":"\u003cbrandlist count=\"0\" ver=\"830369402\"\u003e\u003c/brandlist\u003e","CustomizedInfo":{"BrandFlag":0},"HeadImgMd5":"c5aed7bae5ca972f80cf8d2df4f38047","EncryptUserName":"v3_020b3826fd0301000000000016123032d08d1d000000501ea9a3dba12f95f6b60a0536a1adb6a45ca29c79e8efeb399e19ac72d0a6fa0a7e6ecbf20a49caee030d93eaef4397769ce44ebe3965f909e9ba9919@stranger","AdditionalContactList":{"LinkedinContactItem":{}},"ChatroomVersion":0,"ChatroomMaxCount":0,"ChatroomAccessType":0,"NewChatroomData":{"MemberCount":0,"InfoMask":0},"DeleteFlag":0,"PhoneNumListInfo":{"Count":0},"ChatroomInfoVersion":0,"DeleteContactScene":0,"ChatroomStatus":0,"ExtFlag":0,"chatRoomBusinessType":0,"friendUserName":"qq609654215","textStatusFlag":1,"ringBackSetting":{"finderObjectId":0,"startTs":0,"endTs":0},"bitMask2":4294967295,"bitValue2":0}],"Ret":[0,0],"Ticket":[{},{}]}}
*/
type ContactDetailResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		ContactCount int `json:"ContactCount"`
		ContactList  []struct {
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
			BitMask int64 `json:"BitMask"`
			BitVal  int   `json:"BitVal"`
			ImgFlag int   `json:"ImgFlag"`
			Remark  struct {
			} `json:"Remark"`
			RemarkPyinitial struct {
			} `json:"RemarkPyinitial"`
			RemarkQuanPin struct {
			} `json:"RemarkQuanPin"`
			ContactType   int `json:"ContactType"`
			RoomInfoCount int `json:"RoomInfoCount"`
			DomainList    struct {
			} `json:"DomainList"`
			ChatRoomNotify     int    `json:"ChatRoomNotify"`
			AddContactScene    int    `json:"AddContactScene"`
			Province           string `json:"Province"`
			City               string `json:"City"`
			Signature          string `json:"Signature"`
			PersonalCard       int    `json:"PersonalCard"`
			HasWeiXinHdHeadImg int    `json:"HasWeiXinHdHeadImg"`
			VerifyFlag         int    `json:"VerifyFlag"`
			Level              int    `json:"Level"`
			Source             int    `json:"Source"`
			WeiboFlag          int    `json:"WeiboFlag"`
			AlbumStyle         int    `json:"AlbumStyle"`
			AlbumFlag          int    `json:"AlbumFlag"`
			SnsUserInfo        struct {
				SnsFlag       int     `json:"SnsFlag"`
				SnsBgimgId    string  `json:"SnsBgimgId"`
				SnsBgobjectId float64 `json:"SnsBgobjectId"`
				SnsFlagEx     int     `json:"SnsFlagEx"`
			} `json:"SnsUserInfo"`
			Country         string `json:"Country"`
			BigHeadImgUrl   string `json:"BigHeadImgUrl"`
			SmallHeadImgUrl string `json:"SmallHeadImgUrl"`
			MyBrandList     string `json:"MyBrandList"`
			CustomizedInfo  struct {
				BrandFlag int `json:"BrandFlag"`
			} `json:"CustomizedInfo"`
			HeadImgMd5            string `json:"HeadImgMd5"`
			EncryptUserName       string `json:"EncryptUserName"`
			AdditionalContactList struct {
				LinkedinContactItem struct {
				} `json:"LinkedinContactItem"`
			} `json:"AdditionalContactList"`
			ChatroomVersion    int `json:"ChatroomVersion"`
			ChatroomMaxCount   int `json:"ChatroomMaxCount"`
			ChatroomAccessType int `json:"ChatroomAccessType"`
			NewChatroomData    struct {
				MemberCount int `json:"MemberCount"`
				InfoMask    int `json:"InfoMask"`
			} `json:"NewChatroomData"`
			DeleteFlag       int `json:"DeleteFlag"`
			PhoneNumListInfo struct {
				Count int `json:"Count"`
			} `json:"PhoneNumListInfo"`
			ChatroomInfoVersion  int    `json:"ChatroomInfoVersion"`
			DeleteContactScene   int    `json:"DeleteContactScene"`
			ChatroomStatus       int    `json:"ChatroomStatus"`
			ExtFlag              int    `json:"ExtFlag"`
			ChatRoomBusinessType int    `json:"chatRoomBusinessType"`
			FriendUserName       string `json:"friendUserName"`
			TextStatusFlag       int    `json:"textStatusFlag"`
			RingBackSetting      struct {
				FinderObjectId int `json:"finderObjectId"`
				StartTs        int `json:"startTs"`
				EndTs          int `json:"endTs"`
			} `json:"ringBackSetting"`
			BitMask2  int64  `json:"bitMask2"`
			BitValue2 int    `json:"bitValue2"`
			Alias     string `json:"Alias,omitempty"`
		} `json:"ContactList"`
		Ret    []int `json:"Ret"`
		Ticket []struct {
		} `json:"Ticket"`
	} `json:"data"`
}
