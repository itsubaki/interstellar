package launcher

type Launcher interface {
	Register(in *RegisterInput) *RegisterOutput
	List(in *ListInput) *ListOutput
}

type RegisterInput struct {
}

type RegisterOutput struct {
}

type ListInput struct {
}

type ListOutput struct {
}
