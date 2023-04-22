package api

type UserAPI struct {
}

func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

func (*UserAPI) Nil() bool { return true }
