package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_restrictedSetConfig_Enum_unrestrictedSet    aper.Enumerated = 0
	RACH_ConfigCommon_restrictedSetConfig_Enum_restrictedSetTypeA aper.Enumerated = 1
	RACH_ConfigCommon_restrictedSetConfig_Enum_restrictedSetTypeB aper.Enumerated = 2
)

type RACH_ConfigCommon_restrictedSetConfig struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RACH_ConfigCommon_restrictedSetConfig) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_restrictedSetConfig", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_restrictedSetConfig) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_restrictedSetConfig", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
