package post

import (
	"fmt"
	. "notifications/matrix"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
)

/*
important data: userid - can be array, message, notification type
*/
func InviteMe(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["users"] == nil {
		fmt.Printf("Missing important data")
		return ErrorReturn(t, 404, "000012", "Missing important data")

	}

	users := p.Sanitize(fmt.Sprintf("%v", args["users"]))
	chatid := ""
	if args["chatid"] != nil {
		chatid = p.Sanitize(fmt.Sprintf("%v", args["chatid"]))
	}

	//1. Check Notification Settings by notification type and user to understand which chanels we use for notifications

	res, err := InviteUser(t, users, chatid)
	if err != nil {
		ans["error"] = err.Error()
	}

	ans["answer"] = res

	response = Interfacetoresponse(t, ans)
	return response
}
