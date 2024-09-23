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
func RoomAvatar(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["chatid"] == nil {
		fmt.Printf("Missing important data")
		return ErrorReturn(t, 404, "000012", "Missing important data")

	}

	chatid := p.Sanitize(fmt.Sprintf("%v", args["chatid"]))

	//1. Check Notification Settings by notification type and user to understand which chanels we use for notifications

	err := SetAvatar(chatid)
	if err != nil {
		ans["error"] = err.Error()
	}

	ans["answer"] = "Done"

	response = Interfacetoresponse(t, ans)
	return response
}
