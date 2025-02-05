package transform

type ChatroomAnnouncementRequest struct {
	Appid   string `json:"appid"`
	GroupId string `json:"group_id"`
}

type ChatroomAnnouncementResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Announcement            string `json:"Announcement"`
		ChatRoomInfoVersion     int    `json:"ChatRoomInfoVersion"`
		AnnouncementEditor      string `json:"AnnouncementEditor"`
		AnnouncementPublishTime int    `json:"AnnouncementPublishTime"`
		ChatRoomStatus          int    `json:"ChatRoomStatus"`
		ChatRoomBusinessType    int    `json:"ChatRoomBusinessType"`
		RoomTools               struct {
			RoomToolsWxAppCount int `json:"RoomToolsWxAppCount"`
		} `json:"RoomTools"`
	} `json:"data"`
}
