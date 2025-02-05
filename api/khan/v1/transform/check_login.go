package transform

type CheckLoginRequest struct {
	Appid string `json:"appid"`
	Nkey  string `json:"nkey"`
	Uuid  string `json:"uuid"`
}

type CheckLoginResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		StatusInfo struct {
			Uuid                    string `json:"uuid"`
			Status                  int    `json:"status"`
			HeadImgUrl              string `json:"headImgUrl"`
			PushLoginUrlexpiredTime int    `json:"pushLoginUrlexpiredTime"`
			NickName                string `json:"nickName"`
			ExpiredTime             int    `json:"expiredTime"`
		} `json:"status_info"`
		LoginInfo struct {
			BaseResponse struct {
				Ret    int `json:"ret"`
				ErrMsg struct {
					String string `json:"string"`
				} `json:"errMsg"`
			} `json:"baseResponse"`
			UnifyAuthSectFlag int `json:"unifyAuthSectFlag"`
			AuthSectResp      struct {
				Uin           int `json:"uin"`
				SvrPubEcdhkey struct {
					Nid int `json:"nid"`
					Key struct {
						ILen   int    `json:"iLen"`
						Buffer string `json:"buffer"`
					} `json:"key"`
				} `json:"svrPubEcdhkey"`
				SessionKey struct {
					ILen   int    `json:"iLen"`
					Buffer string `json:"buffer"`
				} `json:"sessionKey"`
				AutoAuthKey struct {
					ILen   int    `json:"iLen"`
					Buffer string `json:"buffer"`
				} `json:"autoAuthKey"`
				WtloginRspBuffFlag int `json:"wtloginRspBuffFlag"`
				WtloginRspBuff     struct {
					ILen int `json:"iLen"`
				} `json:"wtloginRspBuff"`
				WtloginImgRespInfo struct {
					Ksid struct {
						ILen int `json:"iLen"`
					} `json:"ksid"`
					ImgBuf struct {
						ILen int `json:"iLen"`
					} `json:"imgBuf"`
				} `json:"wtloginImgRespInfo"`
				WxVerifyCodeRespInfo struct {
					VerifyBuff struct {
						ILen int `json:"iLen"`
					} `json:"verifyBuff"`
				} `json:"wxVerifyCodeRespInfo"`
				CliDbencryptKey struct {
					ILen int `json:"iLen"`
				} `json:"cliDbencryptKey"`
				CliDbencryptInfo struct {
					ILen int `json:"iLen"`
				} `json:"cliDbencryptInfo"`
				A2Key struct {
					ILen int `json:"iLen"`
				} `json:"a2Key"`
				ShowStyle struct {
					KeyCount int      `json:"keyCount"`
					Key      []string `json:"key"`
				} `json:"showStyle"`
				NewVersion          int    `json:"newVersion"`
				UpdateFlag          int    `json:"updateFlag"`
				AuthResultFlag      int    `json:"authResultFlag"`
				Fsurl               string `json:"fsurl"`
				MmtlsControlBitFlag int    `json:"mmtlsControlBitFlag"`
				ServerTime          int    `json:"serverTime"`
				ClientSessionKey    struct {
					ILen   int    `json:"iLen"`
					Buffer string `json:"buffer"`
				} `json:"clientSessionKey"`
				ServerSessionKey struct {
					ILen   int    `json:"iLen"`
					Buffer string `json:"buffer"`
				} `json:"serverSessionKey"`
				EcdhControlFlag int `json:"ecdhControlFlag"`
			} `json:"authSectResp"`
			AcctSectResp struct {
				UserName         string `json:"userName"`
				NickName         string `json:"nickName"`
				BindUin          int    `json:"bindUin"`
				Alias            string `json:"alias"`
				Status           int    `json:"status"`
				PluginFlag       int    `json:"pluginFlag"`
				RegType          int    `json:"regType"`
				SafeDevice       int    `json:"safeDevice"`
				OfficialUserName string `json:"officialUserName"`
				OfficialNickName string `json:"officialNickName"`
				PushMailStatus   int    `json:"pushMailStatus"`
				Fsurl            string `json:"fsurl"`
			} `json:"acctSectResp"`
			NetworkSectResp struct {
				NewHostList struct {
					Count int `json:"count"`
					List  []struct {
						Host     string `json:"host"`
						Redirect string `json:"redirect"`
					} `json:"list"`
				} `json:"newHostList"`
				NetworkControl struct {
					PortList         string `json:"portList"`
					TimeoutList      string `json:"timeoutList"`
					MinNoopInterval  int    `json:"minNoopInterval"`
					MaxNoopInterval  int    `json:"maxNoopInterval"`
					TypingInterval   int    `json:"typingInterval"`
					NoopIntervalTime int    `json:"noopIntervalTime"`
				} `json:"networkControl"`
				BuiltinIplist struct {
					LongConnectIpcount  int `json:"longConnectIpcount"`
					ShortConnectIpcount int `json:"shortConnectIpcount"`
					LongConnectIplist   []struct {
						Ip   string `json:"ip"`
						Host string `json:"host"`
					} `json:"longConnectIplist"`
					ShortConnectIplist []struct {
						Ip   string `json:"ip"`
						Host string `json:"host"`
					} `json:"shortConnectIplist"`
					Seq int `json:"seq"`
				} `json:"builtinIplist"`
			} `json:"networkSectResp"`
			AxAuthSecRespList struct {
				Count int `json:"count"`
			} `json:"axAuthSecRespList"`
			Id string `json:"id"`
		} `json:"login_info"`
	} `json:"data"`
}
