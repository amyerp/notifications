package matrix

import (
	"encoding/json"
	"fmt"
	. "notifications/global"
	. "notifications/grpc_requests"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"

	"github.com/spf13/viper"
)

func CreateChat(t *pb.Request, user string) (string, error) {

	matrixhost := viper.GetString(fmt.Sprintf("%s.matrix.host", MicroServiceName))
	token := viper.GetString(fmt.Sprintf("%s.matrix.token", MicroServiceName))

	URL := fmt.Sprintf("%s/_matrix/client/r0/createRoom", matrixhost)

	msg := make(map[string]interface{})
	//	intstate := make(map[string]interface{})
	//	contnt := make(map[string]string)
	//	contnt["guest_access"] = "can_join"
	//	intstate["type"] = "m.room.guest_access"
	//	intstate["state_key"] = ""
	//	intstate["content"] = contnt
	msg["name"] = "AMY Bot"
	msg["preset"] = "private_chat"
	msg["visibility"] = "private"
	//	msg["initial_state"] = intstate

	resp, err := JsonReq(URL, msg, token, "POST", "Bearer", "Authorization")

	if err != nil {
		return "", err
	}

	//Get Chat ID
	/*
	   {
	     "room_id": "!sefiuhWgwghwWgh:example.com"
	   }
	*/

	type MatrixResponse struct {
		RoomID string `json:"room_id"`
	}

	var result MatrixResponse

	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		//  fmt.Println("Can not unmarshal JSON")
		return "", err
	}

	go UpdateSocialNetwork(t, user, "matrix", result.RoomID)

	return result.RoomID, nil

}
