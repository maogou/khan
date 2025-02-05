package transform

type GroupDetailRequest struct {
	Appid       string   `json:"appid"`
	GroupIdList []string `json:"group_id_list"`
}

type GroupDetailResponse struct {
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
			} `json:"NickName"`
			Pyinitial struct {
			} `json:"Pyinitial"`
			QuanPin struct {
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
			PersonalCard       int    `json:"PersonalCard"`
			HasWeiXinHdHeadImg int    `json:"HasWeiXinHdHeadImg"`
			VerifyFlag         int    `json:"VerifyFlag"`
			Level              int    `json:"Level"`
			Source             int    `json:"Source"`
			ChatRoomOwner      string `json:"ChatRoomOwner"`
			WeiboFlag          int    `json:"WeiboFlag"`
			AlbumStyle         int    `json:"AlbumStyle"`
			AlbumFlag          int    `json:"AlbumFlag"`
			SnsUserInfo        struct {
				SnsFlag       int `json:"SnsFlag"`
				SnsBgobjectId int `json:"SnsBgobjectId"`
				SnsFlagEx     int `json:"SnsFlagEx"`
			} `json:"SnsUserInfo"`
			SmallHeadImgUrl string `json:"SmallHeadImgUrl"`
			CustomizedInfo  struct {
				BrandFlag int `json:"BrandFlag"`
			} `json:"CustomizedInfo"`
			EncryptUserName       string `json:"EncryptUserName"`
			AdditionalContactList struct {
				LinkedinContactItem struct {
				} `json:"LinkedinContactItem"`
			} `json:"AdditionalContactList"`
			ChatroomVersion    int `json:"ChatroomVersion"`
			ChatroomMaxCount   int `json:"ChatroomMaxCount"`
			ChatroomAccessType int `json:"ChatroomAccessType"`
			NewChatroomData    struct {
				MemberCount    int `json:"MemberCount"`
				ChatRoomMember []struct {
					UserName           string `json:"UserName"`
					NickName           string `json:"NickName"`
					ChatroomMemberFlag int    `json:"ChatroomMemberFlag"`
					InviterUserName    string `json:"InviterUserName,omitempty"`
				} `json:"ChatRoomMember"`
				InfoMask int `json:"InfoMask"`
			} `json:"NewChatroomData"`
			DeleteFlag       int `json:"DeleteFlag"`
			PhoneNumListInfo struct {
				Count int `json:"Count"`
			} `json:"PhoneNumListInfo"`
			ChatroomInfoVersion  int `json:"ChatroomInfoVersion"`
			DeleteContactScene   int `json:"DeleteContactScene"`
			ChatroomStatus       int `json:"ChatroomStatus"`
			ExtFlag              int `json:"ExtFlag"`
			ChatRoomBusinessType int `json:"chatRoomBusinessType"`
			TextStatusFlag       int `json:"textStatusFlag"`
			RingBackSetting      struct {
				FinderObjectId int `json:"finderObjectId"`
				StartTs        int `json:"startTs"`
				EndTs          int `json:"endTs"`
			} `json:"ringBackSetting"`
			BitMask2  int64 `json:"bitMask2"`
			BitValue2 int   `json:"bitValue2"`
		} `json:"ContactList"`
		Ret    []int `json:"Ret"`
		Ticket []struct {
		} `json:"Ticket"`
	} `json:"data"`
}
