package domain

type RegistrationHandler interface {
	Handle() (chan bool, error)
}

type RegistrationReq struct {
}

type RegistrationResp struct {
	NodeId string
}
