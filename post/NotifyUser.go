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
func NotifyUser(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["users"] == nil || args["message"] == nil || args["ntype"] == nil {
		fmt.Printf("Missing important data")
		return ErrorReturn(t, 404, "000012", "Missing important data")

	}

	users := p.Sanitize(fmt.Sprintf("%v", args["users"]))
	message := p.Sanitize(fmt.Sprintf("%v", args["message"]))
	ntype := p.Sanitize(fmt.Sprintf("%v", args["ntype"]))

	//Send general notification
	go AddNotification(*t.UID, message)

	//1. Check Notification Settings by notification type and user to understand which chanels we use for notifications

	if ntype == "matrix" {

		go NotifyByMatrix(t, users, message)

	}

	ans["answer"] = "Done"

	response = Interfacetoresponse(t, ans)
	return response
}
