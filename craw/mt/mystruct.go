package mt

type Data struct {
	BffDatas interface{}
	BffData  []string
}

type BffData struct {
	IsOK    string `json:"isOK"`
	Context struct {
		JsContext struct {
		} `json:"jsContext"`
		WarmUpSeq             interface{} `json:"warmUpSeq"`
		UseAsWarmUpData       bool        `json:"useAsWarmUpData"`
		TimelyWarmUp          bool        `json:"timelyWarmUp"`
		AllowCrossEnvCalls    interface{} `json:"allowCrossEnvCalls"`
		FuncScriptExecTimeout interface{} `json:"funcScriptExecTimeout"`
		ScriptID              int         `json:"scriptId"`
		HeadersMap            struct {
		} `json:"headersMap"`
		ParamsMap struct {
		} `json:"paramsMap"`
		Headers struct {
			MTSIRemoteIP     string `json:"MTSI-remote-ip"`
			UpstreamName     string `json:"$upstream_name"`
			MTSIFlowStrategy string `json:"MTSI-flow-strategy"`
			MTSIScore        string `json:"MTSI-score"`
			Accept           string `json:"Accept"`
			SafaInternal     string `json:"safa-internal"`
			MTSIFlag         string `json:"MTSI-flag"`
			IP               string `json:"ip"`
			XForwardedProto  string `json:"X-Forwarded-Proto"`
			UserAgent        string `json:"User-Agent"`
			Host             string `json:"Host"`
			AcceptEncoding   string `json:"Accept-Encoding"`
			MTSIChecked      string `json:"MTSI-checked"`
			XForwardedFor    string `json:"X-Forwarded-For"`
			PostmanToken     string `json:"Postman-Token"`
			Href             string `json:"href"`
			XRealIP          string `json:"X-Real-IP"`
			MTSIRequestCode  string `json:"MTSI-request-code"`
		} `json:"headers"`
		Params struct {
			PageIdentifier string `json:"page_identifier"`
			Env            struct {
				ClientVersion interface{} `json:"clientVersion"`
			} `json:"env"`
			TechnicianID string `json:"technicianId"`
		} `json:"params"`
		ReqSeqInFuncScript interface{} `json:"reqSeqInFuncScript"`
		ExecParamTypeEnum  struct {
		} `json:"execParamTypeEnum"`
		ExecTotalConsume int  `json:"execTotalConsume"`
		StressTestReq    bool `json:"stressTestReq"`
	} `json:"context"`
	Store struct {
		Client bool `json:"client"`
		Env    struct {
			IsBrowser bool   `json:"isBrowser"`
			IsNode    bool   `json:"isNode"`
			Host      string `json:"host"`
			Href      string `json:"href"`
			Query     struct {
				TechnicianID string `json:"technicianId"`
			} `json:"query"`
			Ua              string `json:"ua"`
			IsBeta          bool   `json:"isBeta"`
			IsLocal         bool   `json:"isLocal"`
			IsX             bool   `json:"isX"`
			IsIOS           bool   `json:"isIOS"`
			IsAndroid       bool   `json:"isAndroid"`
			Os              string `json:"os"`
			IsWX            bool   `json:"isWX"`
			IsWXApp         bool   `json:"isWXApp"`
			IsMtWXApp       bool   `json:"isMtWXApp"`
			IsDpWXApp       bool   `json:"isDpWXApp"`
			IsGroupWXApp    bool   `json:"isGroupWXApp"`
			Version         string `json:"version"`
			Type            string `json:"type"`
			IsDpmerchant    bool   `json:"isDpmerchant"`
			IsMerchantApp   bool   `json:"isMerchantApp"`
			IsDPApp         bool   `json:"isDPApp"`
			IsMTApp         bool   `json:"isMTApp"`
			IsWMApp         bool   `json:"isWMApp"`
			IsKLApp         bool   `json:"isKLApp"`
			IsErpbossproApp bool   `json:"isErpbossproApp"`
			IsWMBApp        bool   `json:"isWMBApp"`
			IsApp           bool   `json:"isApp"`
			IsMT            bool   `json:"isMT"`
			IsMTURL         bool   `json:"isMTUrl"`
			IsRainbow       bool   `json:"isRainbow"`
			UUID            string `json:"uuid"`
			Dpid            string `json:"dpid"`
			Cookie          struct {
			} `json:"cookie"`
			UseVConsole bool `json:"useVConsole"`
			Geo         struct {
			} `json:"geo"`
		} `json:"env"`
	} `json:"store"`
	Geo          bool `json:"geo"`
	ResponseData []struct {
		Code int `json:"code"`
		Data struct {
			TraceID string `json:"traceId"`
			Code    int    `json:"code"`
			Data    struct {
				StatisticValues struct {
				} `json:"statisticValues"`
				HeadActionPoints interface{} `json:"headActionPoints"`
				ExtValues        struct {
					CardInfo struct {
					} `json:"cardInfo"`
				} `json:"extValues"`
				TechCategoryID int `json:"techCategoryId"`
				Share          struct {
					Image string `json:"image"`
					Title string `json:"title"`
					URL   string `json:"url"`
					Desc  string `json:"desc"`
				} `json:"share"`
				ShopIDForFe string `json:"shopIdForFe"`
				AttrValues  struct {
					WorkYearsStr string      `json:"workYearsStr"`
					Skills       []string    `json:"skills"`
					Name         string      `json:"name"`
					JobNumber    int         `json:"jobNumber"`
					PhotoURL     string      `json:"photoUrl"`
					Summary      string      `json:"summary"`
					WorkYears    int         `json:"workYears"`
					JobNumberStr interface{} `json:"jobNumberStr"`
				} `json:"attrValues"`
				TechnicianID      int    `json:"technicianId"`
				ShopCategoryForFe string `json:"shopCategoryForFe"`
			} `json:"data"`
		} `json:"data"`
		Msg string `json:"msg"`
	} `json:"responseData"`
	RequestRepairs []struct {
		Type   string `json:"type"`
		Params []struct {
			TechnicianID string `json:"technicianId"`
		} `json:"params"`
	} `json:"requestRepairs"`
}
