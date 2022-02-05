package services

var (
	ItemService itemServiceIterface = &itemService{}
)

type itemServiceIterface interface {
	GetItem()
	SaveItem()
}

type itemService struct{}

func (s *itemService) GetItem() {

}

func (s *itemService) SaveItem() {

}
