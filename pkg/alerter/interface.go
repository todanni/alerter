package alerter

type Service interface {
	SendLoginAlert() error
	SendRegisterAlert() error
	SendActivationAlert() error
}
