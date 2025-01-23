package event

import (
	"fmt"
	"go-calendar/model"
	"time"

	"github.com/google/uuid"
)

// intentional start with \n
const eventTemplate = `
BEGIN:VEVENT
UID:%s
DTSTAMP:%s
DTSTART:%s
DTEND:%s
SUMMARY:%s
DESCRIPTION:%s%s%s
END:VEVENT`

const dateFormat = "20060102T150405Z"

func NewRenderer(title string, start time.Time, d time.Duration, opts ...Option) func() string {
	return func() string {
		return Render(title, start, d, opts...)
	}
}

func Render(title string, start time.Time, d time.Duration, opts ...Option) string {
	a := &event{
		id:          uuid.NewString(),
		start:       start,
		duration:    d,
		title:       title,
		description: "Event generated using go-calendar.",
	}

	for _, opt := range opts {
		opt(a)
	}

	return a.render()
}

type event struct {
	id                 string
	start              time.Time
	duration           time.Duration
	title, description string
	location           string
	reminders          []model.Renderer
}

func (c event) render() string {
	return fmt.Sprintf(eventTemplate,
		c.id,
		time.Now().Format(dateFormat),
		c.start.Format(dateFormat),
		c.start.Add(c.duration).Format(dateFormat),
		c.title,
		c.description,
		c.location,
		model.AggregateRenderers(c.reminders...)(),
	)
}

// TODO location

type Option func(*event) *event

func WithDescription(d string) Option {
	return func(c *event) *event {
		c.description = d
		return c
	}
}

func WithLocation(l string) Option {
	return func(c *event) *event {
		c.location = "\nLOCATION:" + l
		return c
	}
}

func WithReminder(r model.Renderer) Option {
	return func(c *event) *event {
		c.reminders = append(c.reminders, r)
		return c
	}
}
