package calendar

import (
	"fmt"
	"go-calendar/model"
)

const calendarTemplate = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:go-calendar%s
END:VCALENDAR`

func Render(events ...model.Renderer) string {
	return fmt.Sprintf(calendarTemplate, model.AggregateRenderers(events...)())
}
