package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nttu-ysc/earthquake-wake-up/configs"
	"github.com/nttu-ysc/earthquake-wake-up/notify"
	"github.com/nttu-ysc/earthquake-wake-up/notify/discord"
	"github.com/nttu-ysc/earthquake-wake-up/notify/file"
	"github.com/nttu-ysc/earthquake-wake-up/notify/line"
	"github.com/spf13/viper"
)

var notifiers = map[string]func(ctx context.Context, c *configs.Config) notify.Notifier{
	"line":    line.NewLine,
	"file":    file.NewFile,
	"discord": discord.NewDiscord,
}

func main() {
	c := getConfig()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.AppTimeout)*time.Second)
	defer cancel()

	notifyManager := prepareNotificationManager(ctx, c)

	// parse args
	notifyManager.Notify(os.Args[1], os.Args[2])
}

func prepareNotificationManager(ctx context.Context, c *configs.Config) *notify.NotificationManager {
	notifyManager := notify.NewNotificationManager()
	for _, n := range c.NotifyGroups {
		if _, ok := notifiers[n]; !ok {
			log.Printf("invalid notify group: %s", n)
			continue
		}
		notifyManager.AddNotifier(notifiers[n](ctx, c))
	}

	return notifyManager
}

func getConfig() *configs.Config {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	c := &configs.Config{}
	if err := viper.Unmarshal(c); err != nil {
		panic(fmt.Errorf("unable to decode into config struct, %v", err))
	}
	return c
}
