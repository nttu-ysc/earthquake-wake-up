package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/nttu-ysc/earthquake-wake-up/configs"
	"github.com/nttu-ysc/earthquake-wake-up/notify"
)

var _ notify.Notifier = (*discord)(nil)

type DiscordMessage struct {
	Content string   `json:"content"`
	Embeds  []Embeds `json:"embeds"`
}

type Embeds struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	Color       int           `json:"color"`
	Fields      []EmbedsField `json:"fields"`
	Image       Image         `json:"image"`
	Footer      Footer        `json:"footer"`
}

type Image struct {
	URL string `json:"url"`
}

type EmbedsField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Footer struct {
	Text    string `json:"text"`
	IconUrl string `json:"icon_url"`
}

var DefaultEmbed Embeds = Embeds{
	Title:       "好晃~~~",
	Description: "",
	Url:         "",
	Color:       0,
	Fields:      []EmbedsField{},
	Image: Image{
		URL: "https://i.imgur.com/bKBb4QJ.gif",
	},
	Footer: Footer{},
}

type discord struct {
	ctx     context.Context
	webhook string
}

func NewDiscord(ctx context.Context, c *configs.Config) notify.Notifier {
	return &discord{
		ctx:     ctx,
		webhook: c.Discord.Webhook,
	}
}

// Notify implements notify.Notifier.
func (d *discord) Notify(intensity string, timeLeft string) {
	message := fmt.Sprintf(notify.DefaultNotifyMsg, intensity, timeLeft)

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(DiscordMessage{
		Content: message,
		Embeds: []Embeds{
			{
				Title:       "好晃~~~",
				Description: "",
				Url:         "",
				Color:       5814783,
				Fields: []EmbedsField{
					{
						Name:   "地震震度",
						Value:  intensity,
						Inline: true,
					},
					{
						Name:   "剩餘可逃跑秒數",
						Value:  timeLeft,
						Inline: true,
					},
				},
				Image: Image{
					URL: "https://i.imgur.com/bKBb4QJ.gif",
				},
			},
		},
	}); err != nil {
		fmt.Println(err)
		return
	}

	client := new(http.Client)
	client.Timeout = time.Second * 10

	req, err := http.NewRequest(http.MethodPost, d.webhook, body)
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("status code: %d\n body: %s", resp.StatusCode, respBody)
}
