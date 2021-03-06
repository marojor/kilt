package kiltapi

import (
	"github.com/falcosecurity/kilt/pkg/hocon"
	"github.com/falcosecurity/kilt/pkg/kilt"
)

func NewKiltFromHocon(definition string) *kilt.Kilt {
	impl := hocon.NewKiltHocon(definition)
	return kilt.NewKilt(impl)
}
