package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n2  aper.Enumerated = 0
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n3  aper.Enumerated = 1
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n4  aper.Enumerated = 2
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n6  aper.Enumerated = 3
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n8  aper.Enumerated = 4
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n9  aper.Enumerated = 5
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n10 aper.Enumerated = 6
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n12 aper.Enumerated = 7
)

type PUCCH_ConfigCommon_additionalPRBOffset_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PUCCH_ConfigCommon_additionalPRBOffset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ConfigCommon_additionalPRBOffset_r17", err)
	}
	return nil
}

func (ie *PUCCH_ConfigCommon_additionalPRBOffset_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ConfigCommon_additionalPRBOffset_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
