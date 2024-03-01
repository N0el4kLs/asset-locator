package providers

type IP2DomainEngine interface {
	Name() string
	SearchIP2Domain(ip string) (string, error)
}
