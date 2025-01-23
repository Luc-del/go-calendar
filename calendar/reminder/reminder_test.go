package reminder_test

import (
	"go-calendar/calendar/reminder"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	t.Run("empty description", func(t *testing.T) {
		got := reminder.Render(reminder.RelativeTrigger{
			Timing:   reminder.Before,
			Duration: 3*24*time.Hour + time.Second,
		})

		expected := `
BEGIN:VALARM
ACTION:DISPLAY
TRIGGER:-P3DT1S
END:VALARM`

		assert.Equal(t, expected, got)
	})

	t.Run("with custom description", func(t *testing.T) {
		got := reminder.Render(reminder.RelativeTrigger{
			Timing:   reminder.After,
			Duration: 90 * time.Minute,
		},
			reminder.WithDescription("my desc"),
		)

		expected := `
BEGIN:VALARM
ACTION:DISPLAY
DESCRIPTION:my desc
TRIGGER:PT1H30M
END:VALARM`

		assert.Equal(t, expected, got)
	})
}
