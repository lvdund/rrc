package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_BarringInfoSet_uac_BarringTime_Enum_s4   aper.Enumerated = 0
	UAC_BarringInfoSet_uac_BarringTime_Enum_s8   aper.Enumerated = 1
	UAC_BarringInfoSet_uac_BarringTime_Enum_s16  aper.Enumerated = 2
	UAC_BarringInfoSet_uac_BarringTime_Enum_s32  aper.Enumerated = 3
	UAC_BarringInfoSet_uac_BarringTime_Enum_s64  aper.Enumerated = 4
	UAC_BarringInfoSet_uac_BarringTime_Enum_s128 aper.Enumerated = 5
	UAC_BarringInfoSet_uac_BarringTime_Enum_s256 aper.Enumerated = 6
	UAC_BarringInfoSet_uac_BarringTime_Enum_s512 aper.Enumerated = 7
)

type UAC_BarringInfoSet_uac_BarringTime struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UAC_BarringInfoSet_uac_BarringTime) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSet_uac_BarringTime", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSet_uac_BarringTime) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSet_uac_BarringTime", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
