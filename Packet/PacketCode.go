// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Packet

import "strconv"

type PacketCode byte

const (
	PacketCodeNONE                       PacketCode = 0
	PacketCodeUserLoginClient            PacketCode = 1
	PacketCodeUserLoginServer            PacketCode = 2
	PacketCodeUserConnectVerifyClient    PacketCode = 3
	PacketCodeUserConnectVerifyServer    PacketCode = 4
	PacketCodeUserInfoClient             PacketCode = 5
	PacketCodeUserInfoServer             PacketCode = 6
	PacketCodeUserConnectDoneClient      PacketCode = 7
	PacketCodeUserConnectDoneServer      PacketCode = 8
	PacketCodeUserErrorCloseServerNotify PacketCode = 9
	PacketCodeUserHeartBeatClient        PacketCode = 10
	PacketCodeUserHeartBeatServer        PacketCode = 11
)

var EnumNamesPacketCode = map[PacketCode]string{
	PacketCodeNONE:                       "NONE",
	PacketCodeUserLoginClient:            "UserLoginClient",
	PacketCodeUserLoginServer:            "UserLoginServer",
	PacketCodeUserConnectVerifyClient:    "UserConnectVerifyClient",
	PacketCodeUserConnectVerifyServer:    "UserConnectVerifyServer",
	PacketCodeUserInfoClient:             "UserInfoClient",
	PacketCodeUserInfoServer:             "UserInfoServer",
	PacketCodeUserConnectDoneClient:      "UserConnectDoneClient",
	PacketCodeUserConnectDoneServer:      "UserConnectDoneServer",
	PacketCodeUserErrorCloseServerNotify: "UserErrorCloseServerNotify",
	PacketCodeUserHeartBeatClient:        "UserHeartBeatClient",
	PacketCodeUserHeartBeatServer:        "UserHeartBeatServer",
}

var EnumValuesPacketCode = map[string]PacketCode{
	"NONE":                       PacketCodeNONE,
	"UserLoginClient":            PacketCodeUserLoginClient,
	"UserLoginServer":            PacketCodeUserLoginServer,
	"UserConnectVerifyClient":    PacketCodeUserConnectVerifyClient,
	"UserConnectVerifyServer":    PacketCodeUserConnectVerifyServer,
	"UserInfoClient":             PacketCodeUserInfoClient,
	"UserInfoServer":             PacketCodeUserInfoServer,
	"UserConnectDoneClient":      PacketCodeUserConnectDoneClient,
	"UserConnectDoneServer":      PacketCodeUserConnectDoneServer,
	"UserErrorCloseServerNotify": PacketCodeUserErrorCloseServerNotify,
	"UserHeartBeatClient":        PacketCodeUserHeartBeatClient,
	"UserHeartBeatServer":        PacketCodeUserHeartBeatServer,
}

func (v PacketCode) String() string {
	if s, ok := EnumNamesPacketCode[v]; ok {
		return s
	}
	return "PacketCode(" + strconv.FormatInt(int64(v), 10) + ")"
}