package nearearth

import "fmt"

//go:generate ffakes -i ImpactService
type ImpactService interface {
	GetAll() (Impacts, error)
}

func (e Impact) String() string {
	return fmt.Sprintf("%s: %s", e.Designation, e.DiscoveryDate)
}
func (e Impacts) String() string {
	out := ""
	for _, event := range e {
		out += event.String() + "\n"
	}
	return out
}

type Impacts []Impact

type Impact struct {
	Designation   string `json:"designation"`
	DiscoveryDate string `json:"discovery_date"`
	HMag          string `json:"h_mag,omitempty"`
	MoidAu        string `json:"moid_au"`
	QAu1          string `json:"q_au_1"`
	QAu2          string `json:"q_au_2,omitempty"`
	PeriodYr      string `json:"period_yr,omitempty"`
	IDeg          string `json:"i_deg"`
	Pha           string `json:"pha"`
	OrbitClass    string `json:"orbit_class"`
}
type NoOpImpactService struct{}

func (n NoOpImpactService) GetAll() (Impacts, error) {
	return nil, fmt.Errorf("no op service")
}
func Count(s ImpactService) (int, error) {
	c, err := NewImpactCounter(s)
	if err != nil {
		return 0, err
	}
	count, err := c.count()
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrFailedCount, err)
	}
	return count, err
}

func GetAllImpacts(s ImpactService) (Impacts, error) {
	return s.GetAll()
}

var ErrFailedCount = fmt.Errorf("failed to count near earth impacts")
