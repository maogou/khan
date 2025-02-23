package v1

import "encoding/xml"

type WeChatMessageType49 struct {
	XMLName    xml.Name `xml:"msg"`
	FromUser   string   `xml:"fromusername"`
	AppMessage struct {
		Type        string `xml:"type"`
		Title       string `xml:"title"`
		Description string `xml:"des"`
	} `xml:"appmsg"`
}

type XmlContent struct {
	String string `json:"string"`
}

type CallbackXmlContent struct {
	MsgType int        `json:"MsgType"`
	Content XmlContent `json:"Content"`
}

type FavItem struct {
	XMLName xml.Name `xml:"favitem"`
	Source  struct {
		SourceID string `xml:"sourceid,attr"`
	} `xml:"source"`
}
