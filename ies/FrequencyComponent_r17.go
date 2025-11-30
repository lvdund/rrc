package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FrequencyComponent_r17_Enum_activeCarrier     aper.Enumerated = 0
	FrequencyComponent_r17_Enum_configuredCarrier aper.Enumerated = 1
	FrequencyComponent_r17_Enum_activeBWP         aper.Enumerated = 2
	FrequencyComponent_r17_Enum_configuredBWP     aper.Enumerated = 3
)

type FrequencyComponent_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FrequencyComponent_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FrequencyComponent_r17", err)
	}
	return nil
}

func (ie *FrequencyComponent_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FrequencyComponent_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
