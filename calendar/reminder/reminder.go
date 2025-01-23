package reminder

import (
	"fmt"
)

// intentional start with \n
const reminderTemplate = `
BEGIN:VALARM
ACTION:DISPLAY%s
%s
END:VALARM`

func Render(t trigger, opts ...Option) string {
	a := &reminder{
		trigger: t,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a.render()
}

func NewRenderer(t trigger, opts ...Option) func() string {
	return func() string {
		return Render(t, opts...)
	}
}

type trigger interface {
	render() string
}

type reminder struct {
	description string
	trigger     trigger
}

func (a reminder) render() string {
	return fmt.Sprintf(reminderTemplate, a.description, a.trigger.render())
}

type Option func(*reminder) *reminder

func WithDescription(d string) Option {
	return func(a *reminder) *reminder {
		a.description = "\nDESCRIPTION:" + d
		return a
	}
}
