package domain

type RegistrationReq struct {
	Labels []Label
}

type RegistrationResp struct {
	NodeId string
}
