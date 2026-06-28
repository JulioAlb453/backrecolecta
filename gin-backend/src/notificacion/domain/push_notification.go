//go:build fcm

package domain

type PushNotification struct {
	Title string
	Body  string
	Type  string
	Data  map[string]string
}

type SendResult struct {
	Success bool
	Error   string
}
