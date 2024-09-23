package post

import (
	"fmt"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
)

func ForgotEmail(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["password"] == nil || args["email"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Important Data")
	}

	//Get user's email - request to Users

	email := p.Sanitize(fmt.Sprintf("%v", args["email"]))
	otp := p.Sanitize(fmt.Sprintf("%v", args["password"]))
	//	lang := p.Sanitize(fmt.Sprintf("%v", args["lang"]))
	msg := []string{}
	msga := "New Passowrd for your AMY ERP/CRM account:"
	msg = append(msg, msga)
	msga = p.Sanitize(fmt.Sprintf("<code>%s</code>", otp))
	msg = append(msg, msga)

	ms := &MailSettings{}
	ms.Custom = false
	go SendHTMLEmail(email, "Dear User,", msg, "New AMY Password", "email.html", nil, ms)

	ans["answer"] = "Email Sent"
	response = Interfacetoresponse(t, ans)
	return response

}
