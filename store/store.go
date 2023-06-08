package store

type Store struct {
	videoStatus string
}

func New() Store {
	return Store{
		videoStatus: "213",
	}
}

func (s *Store) addVideo() {

}
