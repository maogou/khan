package download

type DownloadEmojiRequest struct {
	Appid   string   `json:"appid"`
	Md5List []string `json:"md5_list"`
}

type DownloadEmojiResponse struct {
	Data struct {
		BaseResponse struct {
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
			Ret int `json:"ret"`
		} `json:"BaseResponse"`
		EmojiList []struct {
			AesKey     string `json:"aesKey"`
			EncryptUrl string `json:"encryptUrl"`
			ExternMd5  string `json:"externMd5"`
			ExternUrl  string `json:"externUrl"`
			Md5        string `json:"md5"`
			ProductId  string `json:"productId"`
			ThumbUrl   string `json:"thumbUrl"`
			Url        string `json:"url"`
		} `json:"emojiList"`
	} `json:"data"`
	Msg string `json:"msg"`
	Ret int    `json:"ret"`
}
