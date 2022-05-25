package entities

type SMSRes struct {
	DataResponse struct {
		SendSMSResponse struct {
			SendSMSResult struct {
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				Trans_id   string `xml:"trans_id"`
			} `xml:"sendSMSResult"`
		} `xml:"sendSMSResponse"`
	} `xml:"Body"`
}

type SMSReq struct {
	PrivateKey string
	UserID     string
	Trans_ID   string
	MsisDN     string
	HeaderSMS  string
	Message    string
}

// var (
// 	privateKey string = os.Getenv("privateKey")
// 	userid     string = os.Getenv("userid")
// 	trans_id   string
// 	msisdn     string
// 	headerSMS  string = os.Getenv("headerSMS")
// 	message    string
// )