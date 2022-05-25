package ltc

import (
	"WayhaSMS/pkg/entities"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SendSMS(*entities.SMSReq) (string) {
	r := entities.SMSReq{}
	encrypted := Encrypt(&r)
	agent := fiber.AcquireAgent()
	req := agent.Request()
	req.SetRequestURI(os.Getenv("LTC_URL"))
	req.Header.SetMethod(fiber.MethodPost)
	agent.ContentType("text/xml; charset=utf-8")
	agent.BodyString(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://Services.laotel.com/">
	<soapenv:Header/>
	<soapenv:Body>
	   <ser:sendSMS>
		  <ser:msg>
			 <ser:header>
				<ser:userid>` + r.UserID + `</ser:userid>
				<ser:key>` + encrypted + `</ser:key>
				<ser:trans_id> ` + r.Trans_ID + `</ser:trans_id>
				<ser:verson></ser:verson>
			 </ser:header>
			 <ser:msisdn> ` + r.MsisDN + ` </ser:msisdn>
			 <ser:headerSMS> ` + r.HeaderSMS + `</ser:headerSMS>
			 <ser:message> ` + r.Message + `</ser:message>
			 <ser:sms_type></ser:sms_type>
		  </ser:msg>
	   </ser:sendSMS>
	</soapenv:Body>
 </soapenv:Envelope>`)
	if err := agent.Parse(); err != nil {
		panic(err)
	}
	_, body, errs := agent.Bytes() // ...
	if errs != nil {
		panic(errs)
	}
	var data entities.SMSRes
	err := xml.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", data.DataResponse.SendSMSResponse.SendSMSResult.ResultCode)
	return data.DataResponse.SendSMSResponse.SendSMSResult.ResultCode
}

func Encrypt(r *entities.SMSReq) string {
	cmd := exec.Command("java", "-jar", os.Getenv("encrypt_path"), r.UserID+r.Trans_ID+r.MsisDN, r.PrivateKey)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	encrypted := strings.Trim(string(out), "OK\n")
	return encrypted
}
