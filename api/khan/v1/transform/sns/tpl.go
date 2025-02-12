package sns

type TextTimelineObject struct {
	Content string
}

type UrlTimelineObject struct {
	Title       string
	Description string
	Content     string
	ContentUrl  string
	ThumbUrl    string
}

type ImageTimelineItem struct {
	FileMd5  string
	FileUrl  string
	ThumbUrl string
}

type ImageTimelineObject struct {
	Content string
	Items   []ImageTimelineItem
}

type SnsVideoTimelineItem struct {
	Content       string
	FileUrl       string
	ThumbUrl      string
	VideoDuration string
}

type SnsVideoTimelineObject struct {
	Content string
	Video   SnsVideoTimelineItem
}
