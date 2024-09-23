```
https://spec.matrix.org/legacy/r0.0.0/client_server.html#put-matrix-client-r0-rooms-roomid-state-eventtype

https://github.com/vidister/matrix-message-decrypter
https://github.com/chrisport/go-lang-detector
```

Authorization: Bearer TheTokenHere
 !apDgNYhUcmKciSERVL:upload.fiatbay.com - ALex
1. Login to Riot
```
curl -XPOST -d '{"type":"m.login.password", "user":"amy", "password":"oA88WDCdsxpL"}' "https://im.mymates.gmbh/_matrix/client/r0/login"
```
Response:
```
{"user_id":"@amy:upload.fiatbay.com","access_token":"syt_YW15_GcllUmkPJmzOwKLusrVI_08rzzo","home_server":"upload.fiatbay.com","device_id":"UVNXRCCXYU"}

```
2. Create room
```
curl -XPOST -d '{"room_alias_name":"alex@yanchenko.me"}' "https://im.mymates.gmbh/_matrix/client/r0/createRoom?access_token=syt_YW15_GcllUmkPJmzOwKLusrVI_08rzzo"

```
Response:
```
{"room_id":"!oNUFtpolKAShMXkuzU:upload.fiatbay.com","room_alias":"#alex@yanchenko.me:upload.fiatbay.com"}
```

```
{
  "type": "m.room.member",
  "sender": "@amy:upload.fiatbay.com",
  "content": {
    "membership": "join",
    "displayname": "AMY ERP/CRM",
    "avatar_url": "mxc://upload.fiatbay.com/VNZfvPEBkOMuMfjlOttEBXwy"
  },
  "state_key": "@amy:upload.fiatbay.com",
  "origin_server_ts": 1711963741166,
  "unsigned": {
    "age": 122
  },
  "event_id": "$k9LlLdOfL9QWVZPuBNQG4QNhZq3TD1NVIROGtqbfwuY",
  "room_id": "!apDgNYhUcmKciSERVL:upload.fiatbay.com"
}
```

3. Invite user
```
curl -XPOST -d '{"user_id":"@alex:upload.fiatbay.com"}' "https://im.mymates.gmbh/_matrix/client/r0/rooms/%21oNUFtpolKAShMXkuzU:upload.fiatbay.com/invite?access_token=syt_YW15_GcllUmkPJmzOwKLusrVI_08rzzo"
```

```
{
  "type": "m.room.member",
  "sender": "@amy:upload.fiatbay.com",
  "content": {
    "is_direct": true,
    "membership": "invite",
    "displayname": "alex",
    "avatar_url": "mxc://upload.fiatbay.com/MRLMDtefsNUCgJCNUUuYqLPl"
  },
  "state_key": "@alex:upload.fiatbay.com",
  "origin_server_ts": 1711963741501,
  "unsigned": {
    "age": 230
  },
  "event_id": "$gW0kTDBHWVYZCyTp84zwANydzL8DTT53sVd_U4aOgFE",
  "room_id": "!apDgNYhUcmKciSERVL:upload.fiatbay.com"
}

```

4. Send Message
```
 curl -XPOST -d '{"msgtype":"m.text", "body":"Hello, Alex"}' "https://im.mymates.gmbh/_matrix/client/r0/rooms/%21oNUFtpolKAShMXkuzU:upload.fiatbay.com/send/m.room.message?access_token=syt_YW15_GcllUmkPJmzOwKLusrVI_08rzzo"
```
Response:
```
{"event_id":"$LVXcxLFlsd6pBVcsZ1w0PzWUVUdHKqB_72yff2J3-gM"}
```
