package webpush

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	enum2 "notify-api/app/manager/push/enum"
	pushTypes "notify-api/app/manager/push/interfaces"
	"notify-api/app/manager/push/item"
	"notify-api/app/models"
	"notify-api/db/dao"

	"github.com/SherClockHolmes/webpush-go"
)

var client = &http.Client{}

type Provider struct {
	VAPIDPublicKey  string
	VAPIDPrivateKey string
	Mailto          string
}

func (p *Provider) Init(cfg pushTypes.Config) error {
	p.VAPIDPublicKey = cfg[VAPIDPublicKey]
	p.VAPIDPrivateKey = cfg[VAPIDPrivateKey]
	p.Mailto = cfg[Mailto]
	return nil
}

func (p *Provider) getOption() *webpush.Options {
	return &webpush.Options{
		HTTPClient:      client,
		TTL:             60 * 60 * 24,
		Subscriber:      p.Mailto,
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
	}
}

func (p *Provider) Send(ctx context.Context, msg *item.PushMessage) error {
	tokens, ok := dao.Device.GetUserDeviceChannelTokens(ctx, msg.User, p.Name())
	if !ok {
		return errors.New("webpush get user device channel tokens failed")
	}
	if len(tokens) == 0 {
		return nil
	}

	data, err := json.Marshal(models.FromPushMessage(*msg))
	if err != nil {
		return err
	}

	option := p.getOption()
	switch msg.Priority {
	case enum2.PriorityHigh:
		option.Urgency = webpush.UrgencyHigh
	case enum2.PriorityNormal:
		option.Urgency = webpush.UrgencyNormal
	case enum2.PriorityLow:
		option.Urgency = webpush.UrgencyLow
	}

	size := len(tokens)
	c := make(chan error)
	for _, v := range tokens {
		s := &webpush.Subscription{}
		err = json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}

		go func() {
			resp, err := webpush.SendNotificationWithContext(ctx, data, s, option)
			_ = resp.Body.Close()
			c <- err
		}()
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-c:
			if err != nil {
				return err
			} else {
				size--
			}
			if size == 0 {
				return nil
			}
		}
	}
}

type Config struct {
	VAPIDPublicKey  string
	VAPIDPrivateKey string
	Mailto          string
}

const VAPIDPublicKey = "VAPIDPublicKey"
const VAPIDPrivateKey = "VAPIDPrivateKey"
const Mailto = "Mailto"

func (p *Provider) Config() []string {
	return []string{VAPIDPublicKey, VAPIDPrivateKey, Mailto}
}

func (p *Provider) Name() enum2.Sender {
	return enum2.SenderWebPush
}