package etagproviders

type ETagProviderName string

type ETagProvider interface {
	GetETag() string
	GetName() string
	IsSupported() bool
}
