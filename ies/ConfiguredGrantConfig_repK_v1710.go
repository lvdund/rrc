package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_repK_v1710_Enum_n12 aper.Enumerated = 0
	ConfiguredGrantConfig_repK_v1710_Enum_n16 aper.Enumerated = 1
	ConfiguredGrantConfig_repK_v1710_Enum_n24 aper.Enumerated = 2
	ConfiguredGrantConfig_repK_v1710_Enum_n32 aper.Enumerated = 3
)

type ConfiguredGrantConfig_repK_v1710 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ConfiguredGrantConfig_repK_v1710) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_repK_v1710", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_repK_v1710) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_repK_v1710", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
