package notify

import (
	"fmt"
	"os"
	"time"

	"github.com/toorop/pushover"
)

// SendToPushover sends a message via pushover API. It uses env vars PUSHOVER_APP_TOKEN and PUSHOVER_USER to auth.
func SendToPushover(title, message string, priority int) error {

	po, err := pushover.NewPushover(os.Getenv("PUSHOVER_APP_TOKEN"), os.Getenv("PUSHOVER_USER"))
	if err != nil {
		return fmt.Errorf("Error calling pushover API: %s", err)
	}

	_, _, err = po.Push(&pushover.Message{
		Title:     title,
		Message:   message,
		Priority:  priority,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		return fmt.Errorf("Error sending pushover message: %s", err)
	}

	return nil
}
