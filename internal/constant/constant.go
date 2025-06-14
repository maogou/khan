package constant

const (
	AppName     = "khan"
	VERSION     = "v1.0.6"
	QID         = "qid"
	ZEROLOG     = "zerolog"
	XID         = "X-Request-Id"
	GithubRepo  = "github.com/maogou/vRobot@latest"
	GO111MODULE = "on"
	GOPROXY     = "https://goproxy.cn/,direct"
	Windows     = "windows"
	DebugPpro   = "debug/pprof"

	TypeNameSystemMessages = "SystemMessages"
	TypeNameSystemSuccess  = "Long_Serve_START_SUCCESS"
	TypeNameSystemClose    = "Long_Serve_Close"
	TypeNameSystemFailure  = "Long_Serve_START_FAILURE"
	TypeNameSystemTimeout  = "Long_Serve_START_Timeout"
	TypeNameMsgAddMsg      = "AddMsg"
	TypeNameMsgDelContacts = "DelContacts"
	TypeNameMsgModContacts = "ModContacts"
	TypeNameMsgFinderMsg   = "FinderMsg"
	TypeNameMsgOffline     = "Offline"

	License    = "khan_license"
	LicenseKey = "khan_license_key"
	License37  = "37"
	License73  = "73"
	License919 = "919"

	AesWXidKey  = "7d4bb489fc2026181e73483b2224a539"
	AesRedisKey = "e23c03daed4736f9c3da32daff4cd00f"
	AesBackup   = "c9cff2759b28cb51e73c72a48d9fa6e3"

	KhanFree  = "khan_free"
	KhanToken = "Khan-Token"

	Country      = "中国"
	DeviceName   = "iPad"
	Model        = "iPad"
	ProtoVersion = "8.0.48"

	WXQrCodeExpire       = 5 * 60
	WXLoginCache         = "wx_login_cache_"
	WXQrCodeCache        = "wx_qr_code_"
	WXCallbackCache      = "wx_callback_"
	WXLongCache          = "long_connect_cache_"
	WXLongOK             = "longOk"
	WXLongAlreadyConnect = "已开启长链,无需再度开启"

	WXMsgTypePayment    = "2000"
	WXPaymentCacheKey   = "wx_2000_"
	WXMsgTypeRedPacket  = "2001"
	WXRedPacketCacheKey = "wx_20001_"
	WXSnsCacheKey       = "wx_2003_"

	SnsFun          = "sns"
	RedPacketFun    = "redPacket"
	TransferFun     = "transfer"
	CollectMoneyFun = "collectMoney"

	FavIds = "fav_id_"
)
