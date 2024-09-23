package post

import (
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {
	switch *t.Param {
	case "notifyuser":
		response = NotifyUser(t)
	case "notifyuserbyemail":
		response = NotifyUserByEmail(t)
	case "otp_email":
		response = OTPByEmail(t)
	case "confirm_email":
		response = ConfirmEmail(t)
	case "forgot_email":
		response = ForgotEmail(t)
	case "createchat":
		response = MakeChat(t)
	case "inviteme":
		response = InviteMe(t)
	case "roomavatar":
		response = RoomAvatar(t)
	default:
		response = ErrorReturn(t, 404, "000012", "Missing argument")
	}

	return response
}
