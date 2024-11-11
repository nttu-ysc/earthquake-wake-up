package discord

import (
	"context"
	"testing"

	"github.com/nttu-ysc/earthquake-wake-up/configs"
)

func TestDiscord(t *testing.T) {
	conf := &configs.Config{
		AppTimeout:   10,
		NotifyGroups: []string{"discord"},
		Discord: configs.Discord{
			Webhook: "https://discord.com/api/webhooks/1302894221472301117/FNz3O1bctkuh_2l-R3rx2jPUofhk5m0hsdm2uWQK7fboZFvdJkMkGT4mAimh_RzNTSwu",
		},
	}
	discord := NewDiscord(context.Background(), conf)
	discord.Notify("2", "10")
}
