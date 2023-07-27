package domain

type RegistrationHandler interface {
	Handle() (chan bool, error)
}

type RegistrationReq struct {
	Labels []Label
}

type RegistrationResp struct {
	NodeId string
}
