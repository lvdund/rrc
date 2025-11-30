package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_repK_RV_Enum_s1_0231 aper.Enumerated = 0
	ConfiguredGrantConfig_repK_RV_Enum_s2_0303 aper.Enumerated = 1
	ConfiguredGrantConfig_repK_RV_Enum_s3_0000 aper.Enumerated = 2
)

type ConfiguredGrantConfig_repK_RV struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ConfiguredGrantConfig_repK_RV) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_repK_RV", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_repK_RV) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_repK_RV", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
