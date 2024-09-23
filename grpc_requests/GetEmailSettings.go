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
	"encoding/json"
	"fmt"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

type EmailSettings struct {
	Address  string `json:"address"`
	Host     string `json:"host"`
	Password string `json:"password"`
	IsSecure bool   `json:"issecure"`
	Port     string `json:"port"`
	Reply    string `json:"reply"`
	Title    string `json:"title"`
	User     string `json:"user"`
}

// api/v3/catalog/activity/{activityid}/getemailsettings
func GetEmailSettings(t *pb.Request, activityid string) (settings EmailSettings) {

	host := viper.GetString("server.internal_host")
	port := viper.GetString("server.grpc_port")

	s := &pb.Request{}
	module := "catalog"
	inf := "activity"
	infb := "getemailsettings"
	method := "GET"
	s.Module = &module
	s.Param = &inf
	s.ParamID = &activityid
	s.ParamIDD = &infb
	s.Sign = t.Sign
	s.Method = &method

	resp := GRPCConnect(host, port, s)

	SetErrorLog(fmt.Sprintf("%v", resp))

	byte, err := json.Marshal(resp)
	if err != nil {
		SetErrorLog(err.Error())
		return settings
	}

	err = json.Unmarshal(byte, &settings)
	if err != nil {
		SetErrorLog(err.Error())
		return settings
	}

	return settings

}
