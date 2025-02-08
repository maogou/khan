package message

import "encoding/xml"

type NameCardMsgXml struct {
	XMLName                 xml.Name `xml:"msg"`
	Username                string   `xml:"username,attr"`
	Nickname                string   `xml:"nickname,attr"`
	Fullpy                  string   `xml:"fullpy,attr"`
	Shortpy                 string   `xml:"shortpy,attr"`
	Alias                   string   `xml:"alias,attr"`
	ImageStatus             int      `xml:"imagestatus,attr"`
	Scene                   int      `xml:"scene,attr"`
	Province                string   `xml:"province,attr"`
	City                    string   `xml:"city,attr"`
	Sign                    string   `xml:"sign,attr"`
	Sex                     int      `xml:"sex,attr"`
	CertFlag                int      `xml:"certflag,attr"`
	CertInfo                string   `xml:"certinfo,attr"`
	BrandIconUrl            string   `xml:"brandIconUrl,attr"`
	BrandHomeUrl            string   `xml:"brandHomeUrl,attr"`
	BrandSubscriptConfigUrl string   `xml:"brandSubscriptConfigUrl,attr"`
	BrandFlags              int      `xml:"brandFlags,attr"`
	RegionCode              string   `xml:"regionCode,attr"`
}

type PostNameCardConfig struct {
	Content string `json:"content"`
	MsgType int    `json:"msg_type"`
	ToWxid  string `json:"to_wxid"`
}

/**
{
  "appid": "wx_WijJeN12eAvakvBXkh7m3",
  "to_wxid_list": [
    {
      "content": "<msg username=\"wxid_zngcn5z2gw2i12\" nickname=\"不能说的秘密\" fullpy=\"\" shortpy=\"\" alias=\"wxid_zngcn5z2gw2i12\"\r\n     imagestatus=\"3\" scene=\"17\" province=\"\" city=\"\" sign=\"\" sex=\"1\"\r\n     certflag=\"0\" certinfo=\"\" brandIconUrl=\"\" brandHomeUrl=\"\" brandSubscriptConfigUrl=\"\"\r\n     brandFlags=\"0\" regionCode=\"CN\" ></msg>\r\n",
      "msg_type": 42,
      "to_wxid": "xingmaogou"
    }
  ]
}
*/

type PostNameCardRequest struct {
	Appid      string               `json:"appid"`
	ToWxidList []PostNameCardConfig `json:"to_wxid_list"`
}

/**
{
  "ret": 0,
  "msg": "success",
  "data": {
    "BaseResponse": {
      "ret": 0,
      "errMsg": {}
    },
    "Count": 1,
    "List": [
      {
        "Ret": 0,
        "ToUsetName": {
          "string": "xingmaogou"
        },
        "MsgId": 0,
        "NewClientMsgid": 3857738358,
        "Createtime": 1737704983,
        "servertime": 1737704983,
        "Type": 42,
        "NewMsgId": 4064532974335847000
      }
    ],
    "NoKnow": 0
  }
}
*/

type PostNameCardResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Count int `json:"Count"`
		List  []struct {
			Ret        int `json:"Ret"`
			ToUsetName struct {
				String string `json:"string"`
			} `json:"ToUsetName"`
			MsgId          int   `json:"MsgId"`
			NewClientMsgid int64 `json:"NewClientMsgid"`
			Createtime     int   `json:"Createtime"`
			Servertime     int   `json:"servertime"`
			Type           int   `json:"Type"`
			NewMsgId       int64 `json:"NewMsgId"`
		} `json:"List"`
		NoKnow int `json:"NoKnow"`
	} `json:"data"`
}
