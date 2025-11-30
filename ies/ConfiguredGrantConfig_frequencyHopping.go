package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_frequencyHopping_Enum_intraSlot aper.Enumerated = 0
	ConfiguredGrantConfig_frequencyHopping_Enum_interSlot aper.Enumerated = 1
)

type ConfiguredGrantConfig_frequencyHopping struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ConfiguredGrantConfig_frequencyHopping) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_frequencyHopping", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_frequencyHopping) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_frequencyHopping", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
