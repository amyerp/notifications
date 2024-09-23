package model

import "gorm.io/gorm"

type Notifications struct {
	gorm.Model
	UUID    string `gorm:"column:uuid;type:varchar(254);UNIQUE;NOT NULL;" json:"uuid"`
	UID     string `gorm:"column:uid;type:varchar(60);DEFAULT '';" json:"uid"`
	Message string `gorm:"column:message;type:longtext;DEFAULT '';" json:"message"`
	Seen    bool   `gorm:"column:seen;type:bool;DEFAULT false;" json:"seen"`
	Fav     bool   `gorm:"column:fav;type:bool;DEFAULT false;" json:"favorite"`
}

type MatrixChat struct {
	gorm.Model
	ChatID      string `gorm:"column:chatid;type:varchar(60);DEFAULT '';" json:"chat_id"`
	Username    string `gorm:"column:username;type:varchar(60);DEFAULT '';" json:"username"`
	DeviceID    string `gorm:"column:deviceid;type:varchar(60);DEFAULT '';" json:"device_id"`
	PublicKey   string `gorm:"column:pubkey;type:varchar(60);DEFAULT '';" json:"public_key"`
	PrivateKey  string `gorm:"column:privatekey;type:varchar(60);DEFAULT '';" json:"private_key"`
	LastMessage string `gorm:"column:lastmessage;type:longtext;DEFAULT '';" json:"last_message"`
	Msgtype     string `gorm:"column:msgtype;type:varchar(60);DEFAULT '';" json:"msg_type"`
	Sender      string `gorm:"column:sender;type:varchar(60);DEFAULT '';" json:"sender"`
	EventID     string `gorm:"column:eventid;type:varchar(60);DEFAULT '';" json:"event_id"`
	FileID      string `gorm:"column:fileid;type:varchar(60);DEFAULT '';" json:"file_id"`
}

type MatrixSync struct {
	Rooms []MatrixRooms `json:"rooms"`
	//	AccountData []string      `json:"account_data"`
	//	Receipts    []string      `json:"receipts"`
	End string `json:"end"`
}

type MatrixRooms struct {
	RoomID     string         `json:"room_id"`
	Membership string         `json:"membership"`
	Visavility string         `json:"visability"`
	Messages   MatrixMessages `json:"messages"`
	// State       []string `json:"state"`
	// AccountData []string `json:"account_data"`
}

type MatrixMessages struct {
	Chunk []MatrixMessageChunk `json:"chunk"`
	Start string               `json:"start"`
	End   string               `json:"end"`
}

type MatrixMessageChunk struct {
	ChunkType string               `json:"type"`
	RoomID    string               `json:"room_id"`
	Sender    string               `json:"sender"`
	Content   MatrixMessageContent `json:"content"`
	ServerTs  int                  `json:"origin_server_ts"`
	EventID   string               `json:"event_id"`
	UserID    string               `json:"user_id"`
	Age       int                  `json:"age"`
}

type MatrixMessageContent struct {
	Body       string `json:"body"`
	MsgType    string `json:"msgtype"`
	Algorithm  string `json:"algorithm"`
	DeviceID   string `json:"device_id"`
	Ciphertext string `json:"ciphertext"`
	SenderKey  string `json:"sender_key"`
	SessionID  string `json:"session_id"`
}

type RoomEventResponse struct {
	Start string        `json:"start"`
	End   string        `json:"end"`
	Chunk []ChunkStruct `json:"chunk"`
}

type ChunkStruct struct {
	ServerTime int           `json:"origin_server_ts"`
	UserID     string        `json:"user_id"`
	EventID    string        `json:"event_id"`
	Content    ContentStruct `json:"content"`
	RoomID     string        `json:"room_id"`
	EventType  string        `json:"type"`
	Age        int           `json:"age"`
}

type ContentStruct struct {
	Body    string `json:"body"`
	MsgType string `json:"msgtype"`
	Name    string `json:"name"`
}

type UserData struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
