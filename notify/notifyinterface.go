package notify

type Notifier interface {
	Notify(intensity string, timeLeft string)
}
