package mock

import "github.com/ruprict/wtf"

//GaugeService is a mock for wtf.

type GaugeService struct {
	GaugeFn      func(id string) (*wtf.Gauge, error)
	GaugeInvoked bool

	CreateGaugeFn      func(gauge *wtf.Gauge) error
	CreateGaugeInvoked bool

	DeleteGaugeFn      func(id string) error
	DeleteGaugeInvoked bool

	SaveDataPointFn      func(input *wtf.DataPoint) error
	SaveDataPointInvoked bool
}

func (s *GaugeService) Gauge(id string) (*wtf.Gauge, error) {
	s.GaugeInvoked = true
	return s.GuageFn(id)
}

func (s *GaugeService) CraeteGauge(gauge *wtf.Gauge) error {
	s.CreateGaugeInvoked = true
	return s.CreateGuageFn(gauge)
}

func (s *GaugeService) DeleteCraeteGauge(id string) error {
	s.DeleteGaugeInvoked = true
	return s.DeleteGuageFn(id)
}

func (s *GaugeService) SaveDataPoint(input *wtf.DataPoint) error {
	s.SaveDataPointInvoked = true
	return s.SaveDataPointFn(input)
}
