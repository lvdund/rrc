package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_powerControlLoopToUse_Enum_n0 aper.Enumerated = 0
	ConfiguredGrantConfig_powerControlLoopToUse_Enum_n1 aper.Enumerated = 1
)

type ConfiguredGrantConfig_powerControlLoopToUse struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ConfiguredGrantConfig_powerControlLoopToUse) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_powerControlLoopToUse", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_powerControlLoopToUse) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_powerControlLoopToUse", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
