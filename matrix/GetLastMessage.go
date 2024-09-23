package matrix

import (
	"encoding/json"
	"fmt"
	. "notifications/global"
	"notifications/model"

	. "github.com/gogufo/gufo-api-gateway/gufodao"

	"github.com/spf13/viper"
)

func GetSync(user string, chatid string) (string, error) {

	matrixhost := viper.GetString(fmt.Sprintf("%s.matrix.host", MicroServiceName))
	token := viper.GetString(fmt.Sprintf("%s.matrix.token", MicroServiceName))

	URL := fmt.Sprintf("%s/_matrix/client/r0/initialSync", matrixhost)
	// /_matrix/client/v3/sync

	resp, err := JsonGet(URL, nil, token, "Bearer", "Authorization")

	if err != nil {
		return "", err
	}

	var result model.MatrixSync

	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer

		return "", err
	}

	db, _ := ConnectDBv2()

	sn := model.MatrixChat{}

	if result.Rooms[0].Messages.Chunk[9].ChunkType == "m.room.message" {
		sn.LastMessage = result.Rooms[0].Messages.Chunk[9].Content.Body

		sn.Msgtype = result.Rooms[0].Messages.Chunk[9].Content.MsgType
		sn.Sender = result.Rooms[0].Messages.Chunk[9].Sender
		sn.EventID = result.Rooms[0].Messages.Chunk[9].EventID

		chatid := result.Rooms[0].Messages.Chunk[9].RoomID
		sn.ChatID = chatid

		//Check Does such message is exist
		curdata := model.MatrixChat{}
		msgdata := model.MatrixChat{}

		needanswer := false
		isamy := false
		if sn.Sender == "@amy:upload.fiatbay.com" {
			isamy = true
		}

		chatrows := db.Conn.Debug().Model(curdata).Where("chatid = ?", chatid).First(&curdata)
		if chatrows.RowsAffected == 0 {
			//Create first line
			db.Conn.Debug().Model(sn).Save(&sn)
			if !isamy {
				needanswer = true
			}

		} else {
			//Check Does message is same
			msgrows := db.Conn.Debug().Model(msgdata).Where("chatid = ? AND eventid = ?", chatid, sn.EventID).First(&msgdata)
			if msgrows.RowsAffected == 0 {
				//Update
				db.Conn.Debug().Model(sn).Where("chatid = ?", chatid).Updates(map[string]interface{}{"lastmessage": sn.LastMessage, "msgtype": sn.Msgtype, "sender": sn.Sender, "eventid": sn.EventID})
				if !isamy {
					needanswer = true
				}
			}

			//No action need

		}

		if needanswer {
			go AmyAnswer(sn.LastMessage, sn.Sender, sn.ChatID)
		}

		return sn.LastMessage, nil
	}

	return "Message is encrypted or it is  not a message", nil

}
