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
func MakeChat(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["users"] == nil {
		fmt.Printf("Missing important data")
		return ErrorReturn(t, 404, "000012", "Missing important data")

	}

	users := p.Sanitize(fmt.Sprintf("%v", args["users"]))

	//1. Check Notification Settings by notification type and user to understand which chanels we use for notifications

	answer, err := CreateChat(t, users)
	if err != nil {
		ans["error"] = err.Error()
	}

	ans["answer"] = answer

	response = Interfacetoresponse(t, ans)
	return response
}
