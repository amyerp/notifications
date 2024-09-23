package matrix

import (
	"fmt"
	. "notifications/global"

	. "github.com/gogufo/gufo-api-gateway/gufodao"

	"github.com/spf13/viper"
)

func SetAvatar(chatid string) error {

	matrixhost := viper.GetString(fmt.Sprintf("%s.matrix.host", MicroServiceName))
	token := viper.GetString(fmt.Sprintf("%s.matrix.token", MicroServiceName))
	avatar := viper.GetString(fmt.Sprintf("%s.matrix.avatar", MicroServiceName))

	URL := fmt.Sprintf("%s/_matrix/client/r0/rooms/%s/state/m.room.avatar/", matrixhost, chatid)

	msg := make(map[string]interface{})
	msg["url"] = avatar

	_, err := JsonReq(URL, msg, token, "PUT", "Bearer", "Authorization")

	if err != nil {
		return err
	}

	return nil

}
