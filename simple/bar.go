package simple

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(fooRepository *BarRepository) *BarService {
	return &BarService{BarRepository: fooRepository}
}
