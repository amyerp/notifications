package matrix

import (
	"fmt"
	. "notifications/global"
	"strings"
	"time"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	"github.com/spf13/viper"
)

/*
1. Определяем язык запроса. Если язык не соответсвует нашему выдаем ошибку
2. Определяем вопросительное ли преложение или утвердительное (по знаку вопроса и по вопросительным словам)
2.2. Если вопросительное значит есть запрос это или GET или TRACE
2.3. Определяем существителные (предмет запроса)
2.4. Определяем прилагательные (уочнение запроса)
2.5. Определяем местоимения (принадлежность запроса)

3. Если предложение утвердительное определяем глаголы (действие)
3.1. Определяем существителные (предмет запроса)
3.2. Определяем прилагательные (уочнение запроса)
3.3. Определяем местоимения (принадлежность запроса)

4. Если запрос требует диалога - помечаем разговор как диалог
*/

func AmyAnswer(message string, sender string, chatid string) {
	lowerStr := strings.ToLower(message)
	keys := []string{"hi", "day", "wife"}
	keysinmsg := []string{}

	for i := 0; i < len(keys); i++ {

		if strings.Contains(lowerStr, keys[i]) {
			keysinmsg = append(keysinmsg, keys[i])
		}

	}

	if len(keysinmsg) > 0 {

		for n := 0; n < len(keysinmsg); n++ {

			SendMSQbyKey(keysinmsg[n], sender, chatid)

		}

	}

}

func SendMSQbyKey(key string, sender string, chatid string) {

	message := ""

	switch key {
	case "hi":
		message = "Hi, Alex"
	case "day":
		currentTime := time.Now()
		strtime := currentTime.Format("02 January 2006")
		message = fmt.Sprintf("Today is %s", strtime)
	case "wife":
		message = "Alex, your wife's name is Anna"
	}

	if message != "" {
		msg := make(map[string]interface{})
		msg["msgtype"] = "m.text"
		msg["body"] = message

		matrixhost := viper.GetString(fmt.Sprintf("%s.matrix.host", MicroServiceName))
		token := viper.GetString(fmt.Sprintf("%s.matrix.token", MicroServiceName))

		URL := fmt.Sprintf("%s/_matrix/client/r0/rooms/%s/send/m.room.message", matrixhost, chatid)

		JsonReq(URL, msg, token, "POST", "Bearer", "Authorization")

	}

}
