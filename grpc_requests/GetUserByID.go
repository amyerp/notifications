package grpc_requests

import (
	"encoding/json"
	"fmt"
	. "notifications/model"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func GetUserByID(t *pb.Request, id string) (user UserData) {
	host := viper.GetString("server.internal_host")
	port := viper.GetString("server.grpc_port")

	s := &pb.Request{}
	module := "auth"
	param := "getuserbyid"
	method := "GET"
	s.Module = &module
	s.Param = &param
	s.Sign = t.Sign
	s.Method = &method
	s.UID = t.UID
	args := make(map[string]interface{})
	args["uid"] = id
	argst := ToMapStringAny(args)
	s.Args = argst

	resp := GRPCConnect(host, port, s)

	SetErrorLog(fmt.Sprintf("%v", resp))

	byte, err := json.Marshal(resp["user"])
	if err != nil {
		SetErrorLog(err.Error())
		return user
	}

	err = json.Unmarshal(byte, &user)
	if err != nil {
		SetErrorLog(err.Error())
		return user
	}

	return user

}
