package sessions

type contextKey string

func (c contextKey) String() string {
	return "go-common/web/middleware/" + string(c)
}
