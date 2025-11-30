package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_InformationCommon_r16_msg1_FDM_r16_Enum_one   aper.Enumerated = 0
	RA_InformationCommon_r16_msg1_FDM_r16_Enum_two   aper.Enumerated = 1
	RA_InformationCommon_r16_msg1_FDM_r16_Enum_four  aper.Enumerated = 2
	RA_InformationCommon_r16_msg1_FDM_r16_Enum_eight aper.Enumerated = 3
)

type RA_InformationCommon_r16_msg1_FDM_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RA_InformationCommon_r16_msg1_FDM_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RA_InformationCommon_r16_msg1_FDM_r16", err)
	}
	return nil
}

func (ie *RA_InformationCommon_r16_msg1_FDM_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RA_InformationCommon_r16_msg1_FDM_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
