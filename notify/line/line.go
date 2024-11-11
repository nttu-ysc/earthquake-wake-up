package line

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/nttu-ysc/earthquake-wake-up/configs"
	"github.com/nttu-ysc/earthquake-wake-up/notify"
)

var _ notify.Notifier = (*line)(nil)

const url = "https://notify-api.line.me/api/notify"

type line struct {
	ctx    context.Context
	tokens string
}

func NewLine(ctx context.Context, c *configs.Config) notify.Notifier {
	return &line{
		ctx:    ctx,
		tokens: c.Line.Token,
	}
}

func (l *line) Notify(intensity string, timeLeft string) {
	message := fmt.Sprintf(notify.DefaultNotifyMsg, intensity, timeLeft)
	wg := new(sync.WaitGroup)
	for _, token := range strings.Split(l.tokens, ",") {
		wg.Add(1)
		go func(token string) {
			defer wg.Done()
			client := new(http.Client)
			client.Timeout = 30 * time.Second
			req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(fmt.Sprintf("message=\n%s", message))))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			if resp, err := client.Do(req); err != nil {
				fmt.Println(err)
			} else {
				defer resp.Body.Close()
				byt, _ := io.ReadAll(resp.Body)
				fmt.Println(string(byt))
			}
		}(token)
	}
	wg.Wait()
}
