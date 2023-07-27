package domain

type RegistrationHandler interface {
	Handle() (chan bool, error)
}
