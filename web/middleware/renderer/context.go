package renderer

type contextKey string

func (c contextKey) String() string {
	return "go-common/web/middleware/" + string(c)
}
