package parser

import (
	"encoding/json"
	"time"
)

type Events struct {
	Events []Event `json:"events"`
}

type Event struct {
	Title       string       `json:"title"`
	StartDate   time.Time    `json:"start_date"`
	Duration    JSONDuration `json:"duration"`
	Description string       `json:"description,omitempty"`
	Location    string       `json:"location,omitempty"`
	Reminders   []Reminder   `json:"reminders,omitempty"`
}

type Reminder struct {
	Type        string       `json:"type"`
	Description string       `json:"description,omitempty"`
	AfterEvent  bool         `json:"after_event,omitempty"`
	Duration    JSONDuration `json:"duration,omitempty"`
}

// JSONDuration is a custom duration type to handle JSON unmarshalling
type JSONDuration time.Duration

// UnmarshalJSON converts string duration (e.g., "1h30m0s") to time.Duration
func (d *JSONDuration) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		return err
	}
	*d = JSONDuration(duration)
	return nil
}
