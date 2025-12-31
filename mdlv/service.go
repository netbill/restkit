package mdlv

type Service struct {
	skUser string
	ctxKey any
}

func New(skUser string, ctxKey any) Service {
	return Service{
		skUser: skUser,
		ctxKey: ctxKey,
	}
}
