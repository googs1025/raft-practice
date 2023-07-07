package cache

type Cache interface {
	SetItem(key string, value string) error
	GetItem(key string) (string, error)
	DelItem(key string) error
}
