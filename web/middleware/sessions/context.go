package sessions

type contextKey string

func (c contextKey) String() string {
	return "github.com/tomadele/go-common/web/middleware/" + string(c)
}
