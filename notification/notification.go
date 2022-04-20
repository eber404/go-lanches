package notification

import "strings"

type NotificationType string

const (
	Error = "error"
	Warn  = "warn"
	Info  = "info"
)

type Notification struct {
	Message string
	Type    NotificationType
}

type NotificationMethods interface {
	hasErrors() bool
	getErrorMessages() []string
}

type Notifications struct {
	Notifications []Notification
}

var instance *Notification = nil
var notifications []Notification = nil

func init() {
	if instance == nil {
		instance = &Notification{}
	}
}

func HasErrors() bool {
	hasError := false

	for _, notification := range notifications {
		if notification.Type == NotificationType(Error) {
			hasError = true
		}
	}

	return hasError
}

func GetInstance() *Notification {
	return instance
}

func Add(Message string, Type NotificationType) {
	notification := Notification{Message: Message, Type: Type}
	notifications = append(notifications, notification)
}

func GetErrorMessages() string {
	var messages []string

	for _, notification := range notifications {
		if notification.Type == NotificationType(Error) {
			messages = append(messages, notification.Message)
		}
	}

	return strings.Join(messages, " | ")
}
