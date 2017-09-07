package providers

type ProviderInterface interface {
	Register() (err error)
}
