package notify

import (
	"sync"
)

type NotificationManager struct {
	notifiers []Notifier
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{}
}

func (n *NotificationManager) AddNotifier(notifier Notifier) {
	n.notifiers = append(n.notifiers, notifier)
}

func (n *NotificationManager) Notify(intensity string, timeLeft string) {
	if len(n.notifiers) == 0 {
		return
	}
	wg := new(sync.WaitGroup)
	for _, notifier := range n.notifiers {
		wg.Add(1)
		go func(n Notifier, intensity string, timeLeft string) {
			defer wg.Done()
			n.Notify(intensity, timeLeft)
		}(notifier, intensity, timeLeft)
	}
	wg.Wait()
}
