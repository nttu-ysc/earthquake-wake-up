package file

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nttu-ysc/earthquake-wake-up/configs"
	"github.com/nttu-ysc/earthquake-wake-up/notify"
)

var _ notify.Notifier = (*file)(nil)

type file struct {
	ctx  context.Context
	path string
}

func NewFile(ctx context.Context, c *configs.Config) notify.Notifier {
	return &file{
		ctx:  ctx,
		path: c.File.Path,
	}
}

func (f *file) Notify(intensity string, timeLeft string) {
	message := fmt.Sprintf(notify.DefaultNotifyMsg, intensity, timeLeft)
	var file *os.File
	if _, err := os.Stat(f.path); err != nil && os.IsNotExist(err) {
		file, _ = os.Create(f.path)
	} else {
		file, _ = os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY, 0644)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Printf("\n%s", message)
}
