package utils

type Setuprelease[T any] struct {
	Choice uint64
	Release *struct{}
	Setup   *T
}