package service

import "github.com/stretchr/testify/mock"

type MockAlerterService struct {
	mock.Mock
}

func (a *MockAlerterService) SendLoginAlert() error {
	args := a.Called()
	return args.Error(0)
}

func (a *MockAlerterService) SendRegisterAlert() error {
	args := a.Called()
	return args.Error(0)
}

func (a *MockAlerterService) SendActivationAlert() error {
	args := a.Called()
	return args.Error(0)
}
