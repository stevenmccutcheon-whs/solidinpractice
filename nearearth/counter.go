package nearearth

import "fmt"

type Counter struct {
	nieService ImpactService
}

func NewImpactCounter(service ImpactService) (Counter, error) {
	if service == nil {
		return Counter{nieService: NoOpImpactService{}}, fmt.Errorf("provided service is nil; must provide a valid service")
	}
	return Counter{nieService: service}, nil
}

func (n Counter) count() (int, error) {
	nei, err := n.nieService.GetAll()
	if err != nil {
		return 0, err
	}
	return len(nei), nil
}
