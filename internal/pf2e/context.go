package pf2e

type Context struct {
	Config Configuration
}

func NewContext() Context {
	return Context{Config: NewConfiguration()}
}
