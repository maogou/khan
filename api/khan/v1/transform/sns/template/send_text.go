package template

type CDATA struct {
	Content string `xml:",cdata"`
}

type TimelineObject struct {
	ID                  int64         `xml:"id"`
	Username            string        `xml:"username"`
	CreateTime          string        `xml:"createTime"`
	ContentDesc         CDATA         `xml:"contentDesc"`
	ContentDescShowType int           `xml:"contentDescShowType"`
	ContentDescScene    int           `xml:"contentDescScene"`
	Private             int           `xml:"private"`
	SightFolded         int           `xml:"sightFolded"`
	ShowFlag            int           `xml:"showFlag"`
	AppInfo             AppInfo       `xml:"appInfo"`
	SourceUserName      string        `xml:"sourceUserName"`
	SourceNickName      string        `xml:"sourceNickName"`
	StatisticsData      string        `xml:"statisticsData"`
	StatExtStr          string        `xml:"statExtStr"`
	ContentObject       ContentObject `xml:"ContentObject"`
	ActionInfo          ActionInfo    `xml:"actionInfo"`
	Location            Location      `xml:"location"`
	PublicUserName      string        `xml:"publicUserName"`
	Streamvideo         Streamvideo   `xml:"streamvideo"`
}

type AppInfo struct {
	ID            string `xml:"id"`
	Version       string `xml:"version"`
	AppName       string `xml:"appName"`
	InstallUrl    string `xml:"installUrl"`
	FromUrl       string `xml:"fromUrl"`
	IsForceUpdate int    `xml:"isForceUpdate"`
	IsHidden      int    `xml:"isHidden"`
}

type ContentObject struct {
	ContentStyle int    `xml:"contentStyle"`
	Title        string `xml:"title"`
	Description  string `xml:"description"`
	MediaList    string `xml:"mediaList"`
	ContentUrl   string `xml:"contentUrl"`
}

type ActionInfo struct {
	AppMsg AppMsg `xml:"appMsg"`
}

type AppMsg struct {
	MessageAction string `xml:"messageAction"`
}

type Location struct {
	POIClassifyID   string `xml:"poiClassifyId,attr"`
	POIName         string `xml:"poiName,attr"`
	POIAddress      string `xml:"poiAddress,attr"`
	POIClassifyType string `xml:"poiClassifyType,attr"`
	City            string `xml:"city,attr"`
}

type Streamvideo struct {
	Streamvideourl      string `xml:"streamvideourl"`
	Streamvideothumburl string `xml:"streamvideothumburl"`
	Streamvideoweburl   string `xml:"streamvideoweburl"`
}
