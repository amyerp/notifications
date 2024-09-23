package post

import (
	"encoding/json"
	. "notifications/grpc_requests"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

type MaiRequest struct {
	Title    string   `json:"title"`
	Message  []string `json:"message"`
	Users    string   `json:"users"`
	Template string   `json:"template"`
}

func NotifyUserByEmail(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)

	if args["title"] == nil || args["message"] == nil || args["users"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Important Data")
	}

	data := MaiRequest{}

	JsonArgs, err := json.Marshal(args)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = json.Unmarshal(JsonArgs, &data)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	if data.Template == "" {
		data.Template = "email.html"
	}

	//Get user's email - request to Users
	user := GetUserByID(t, data.Users)
	email := user.Email

	ms := &MailSettings{}
	ms.Custom = false
	go SendHTMLEmail(email, user.Name, data.Message, data.Title, data.Template, nil, ms)

	ans["answer"] = "Email Sent"
	response = Interfacetoresponse(t, ans)
	return response
}
