package matrix

import (
	"fmt"
	. "notifications/global"
	. "notifications/grpc_requests"

	pb "github.com/gogufo/gufo-api-gateway/proto/go"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	"github.com/spf13/viper"
)

func NotifyByMatrix(t *pb.Request, user string, message string) {

	//1. Check does user has matrix Account

	chatid, _ := GetChatID(t, user, "matrix")

	//2. Check does user has chatid. If not, Create chatid

	if chatid == "" {
		//Create chat with user
		newchatid, err := CreateChat(t, user)
		if err != nil {
			return
		}

		//Set avatar
		SetAvatar(newchatid)

		//Invite user to chat
		_, err = InviteUser(t, user, newchatid)

		if err != nil {
			return
		}

		chatid = newchatid
	}

	//3. Send message
	msg := make(map[string]interface{})
	msg["msgtype"] = "m.text"
	msg["body"] = message

	matrixhost := viper.GetString(fmt.Sprintf("%s.matrix.host", MicroServiceName))
	token := viper.GetString(fmt.Sprintf("%s.matrix.token", MicroServiceName))

	URL := fmt.Sprintf("%s/_matrix/client/r0/rooms/%s/send/m.room.message", matrixhost, chatid)

	/*
		type Events struct {
			gorm.Model
			EvenetID   string `gorm:"column:eventid;type:varchar(254);UNIQUE;NOT NULL;"`
			UID        string `gorm:"column:uid;type:varchar(60);DEFAULT '';"`
			ActivityID string `gorm:"column:activityid;type:varchar(60);DEFAULT '';"`
			AccountID  string `gorm:"column:accountid;type:varchar(60);DEFAULT '';"`
			Event      string `gorm:"column:event;type:longtext;DEFAULT '';"`
		}

		newevent := &Events{}
		newevent.EvenetID = "eventid"
		newevent.UID = ""
		newevent.ActivityID = ""
		newevent.AccountID = "accountid"
		newevent.Event = URL


		db.Conn.Debug().Create(&newevent)
	*/

	_, err := JsonReq(URL, msg, token, "POST", "Bearer", "Authorization")

	if err != nil {
		fmt.Printf(err.Error())
	}

	//	fmt.Printf("DONE")

}
