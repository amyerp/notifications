package get

import (
	"fmt"
	"notifications/model"
	"strconv"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func GetNotifications(t *pb.Request) (response *pb.Response) {

	args := ToMapStringInterface(t.Args)
	ans := make(map[string]interface{})

	allevents := []*model.Notifications{}
	limit := 50

	if args["limit"] != nil {
		limit, _ = strconv.Atoi(fmt.Sprintf("%v", args["limit"]))
	}

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}
		return ErrorReturn(t, 500, "000027", err.Error())
	}

	db.Conn.Debug().Model(model.Notifications{}).Where("uid = ?", *t.UID).Order("created_at desc").Limit(limit).Find(&allevents)

	ans["notifications"] = allevents
	response = Interfacetoresponse(t, ans)

	return response
}
