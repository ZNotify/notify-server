package push

import (
	"crypto/rand"
	"fmt"
	"github.com/ZNotify/server/config"
	"github.com/ZNotify/server/db/entity"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const APIURL = "https://api.xmpush.xiaomi.com/v2/message/user_account"

var MiPushClient *http.Client

func InitMiPushClient() {
	MiPushClient = &http.Client{}
}

func SendViaMiPush(msg *entity.Message) error {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	notifyID := n.Int64()

	// Build MIPush request
	title := msg.Title
	content := msg.Content
	long := msg.Long

	msgID := msg.ID

	intentUriFormat := "intent:#Intent;launchFlags=0x14000000;component=top.learningman.push/.TranslucentActivity;S.userID=%s;S.long=%s;S.msgID=%s;S.title=%s;S.createdAt=%s;S.content=%s;end"
	intentUri := fmt.Sprintf(intentUriFormat,
		url.PathEscape(msg.UserID),
		url.PathEscape(long),
		url.PathEscape(msgID),
		url.PathEscape(title),
		url.PathEscape(msg.CreatedAt.Format(time.RFC3339)),
		url.PathEscape(content))

	postData := url.Values{
		"user_account":            {msg.UserID},
		"payload":                 {long},
		"restricted_package_name": {"top.learningman.push"},
		"pass_through":            {"0"},
		"title":                   {title},
		"description":             {content},
		"notify_id":               {strconv.Itoa(int(notifyID))},
		"extra.id":                {msgID},
		"extra.notify_effect":     {"2"}, // https://dev.mi.com/console/doc/detail?pId=1278#_3_2
		"extra.intent_uri":        {intentUri},
	}.Encode()

	req, err := http.NewRequest(
		http.MethodPost,
		APIURL,
		strings.NewReader(postData))

	req.Header.Set("Authorization", fmt.Sprintf("key=%s", config.MiPushSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := MiPushClient.Do(req)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	return nil
}
