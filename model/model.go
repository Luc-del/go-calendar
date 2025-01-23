package model

type Renderer func() string

func AggregateRenderers(renderers ...Renderer) Renderer {
	return func() string {
		var res string
		for _, r := range renderers {
			res += r()
		}

		return res
	}
}
