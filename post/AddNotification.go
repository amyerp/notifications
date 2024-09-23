package post

import (
	. "notifications/model"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
)

func AddNotification(uid string, message string) {

	nid := Hashgen(64)

	newevent := &Notifications{}
	newevent.UUID = nid
	newevent.UID = uid
	newevent.Message = message

	db, err := ConnectDBv2()
	if err != nil {
		return
	}

	db.Conn.Debug().Create(&newevent)

}
