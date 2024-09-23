package matrix

import (
	"fmt"
	. "notifications/global"
	. "notifications/grpc_requests"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"

	"github.com/spf13/viper"
)

func InviteUser(t *pb.Request, user string, chatid string) (string, error) {

	matrixhost := viper.GetString(fmt.Sprintf("%s.matrix.host", MicroServiceName))
	token := viper.GetString(fmt.Sprintf("%s.matrix.token", MicroServiceName))
	domain := viper.GetString(fmt.Sprintf("%s.matrix.domain", MicroServiceName))

	if chatid == "" {
		//Get chatid from DB

		chatid, _ := GetChatID(t, user, "matrix")

		if chatid == "" {
			return "No Chat ID", nil
		}
	}

	URL := fmt.Sprintf("%s/_matrix/client/r0/rooms/%s/invite", matrixhost, chatid)
	fullname := fmt.Sprintf("@%s:%s", user, domain)

	msg := make(map[string]interface{})
	msg["user_id"] = fullname

	_, err := JsonReq(URL, msg, token, "POST", "Bearer", "Authorization")

	if err != nil {
		return "", err
	}

	return "", nil

}
