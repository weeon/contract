package contract

type Config interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
