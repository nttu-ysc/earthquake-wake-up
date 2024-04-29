package notify

type Notifier interface {
	Notify(message string)
}
