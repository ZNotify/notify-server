package push

import (
	"notify-api/push/host/telegram"
	"notify-api/push/host/websocket"
	"notify-api/push/provider/fcm"
	"notify-api/push/provider/webpush"
	"notify-api/push/provider/wns"
	pushTypes "notify-api/push/types"
)

type senders = []pushTypes.Sender

var availableSenders = senders{
	new(fcm.Provider),
	new(webpush.Provider),
	new(wns.Provider),
	new(websocket.Host),
	new(telegram.Host),
}

var activeSenders = senders{}

func IsValid(id string) bool {
	for _, v := range activeSenders {
		if v.Name() == id {
			return true
		}
	}
	return false
}

func TryGetInitialTokenMeta(channel string) (string, bool) {
	for _, v := range activeSenders {
		if v.Name() == channel {
			if host, ok := v.(pushTypes.SenderWithInitialTokenMeta); ok {
				return host.GetInitialTokenMeta(), true
			}
		}
	}
	return "", false
}
