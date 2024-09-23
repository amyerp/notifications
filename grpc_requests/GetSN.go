// Copyright 2020 - 2024 Alexey Yanchenko <mail@yanchenko.me>
//
// This file is part of the Gufo library.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package grpc_requests

import (
	"fmt"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func GetChatID(t *pb.Request, uid string, snid string) (chatid string, username string) {

	host := viper.GetString("server.internal_host")
	port := viper.GetString("server.grpc_port")

	s := &pb.Request{}
	module := "catalog"
	inf := "person"
	infb := "social_network_by_id"
	method := "GET"
	s.Module = &module
	s.Param = &inf
	s.ParamID = &uid
	s.ParamIDD = &infb
	s.Sign = t.Sign
	s.Method = &method
	args := make(map[string]interface{})
	args["snid"] = snid
	argst := ToMapStringAny(args)
	s.Args = argst

	ans := GRPCConnect(host, port, s)
	SetErrorLog(fmt.Sprintf("%v", ans))
	chatid = fmt.Sprintf("%v", ans["chatid"])
	username = fmt.Sprintf("%v", ans["username"])
	return chatid, username

}
