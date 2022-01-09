package fetcher

type Fetcher interface {
	Fetch(store Store) interface{}
}

type Store struct {
	storeType StoreType
}

type StoreType string

const (
	gitHub   StoreType = "github"
	postgres StoreType = "postgres"
)

func Fetch(fetcher Fetcher) {

}
