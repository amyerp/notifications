package get

import (
	"fmt"
	. "notifications/global"
	. "notifications/grpc_requests"
	. "notifications/matrix"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func GetMessage(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})

	userid := *t.UID
	setingskey := fmt.Sprintf("%s.matrix.domain", MicroServiceName)
	domain := viper.GetString(setingskey)

	//1. Get ChatID and username
	chatid, username := GetChatID(t, userid, "matrix")
	if chatid == "" {
		return ErrorReturn(t, 400, "00003", "User not found")
	}

	user := fmt.Sprintf("@%s:%s", username, domain)

	message, err := GetSync(user, chatid)
	//	message, err := GetKey(user, chatid)

	if err != nil {
		return ErrorReturn(t, 400, "00006", err.Error())
	}

	ans["notifications"] = message
	response = Interfacetoresponse(t, ans)

	return response
}
