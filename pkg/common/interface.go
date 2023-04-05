package common

type IService interface {
	Start() error
	Stop() error
	Name() string
}
