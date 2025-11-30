package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n1   aper.Enumerated = 0
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n2   aper.Enumerated = 1
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n4   aper.Enumerated = 2
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n6   aper.Enumerated = 3
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n8   aper.Enumerated = 4
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n10  aper.Enumerated = 5
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n20  aper.Enumerated = 6
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n50  aper.Enumerated = 7
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n100 aper.Enumerated = 8
	RA_InformationCommon_r16_msgA_TransMax_r17_Enum_n200 aper.Enumerated = 9
)

type RA_InformationCommon_r16_msgA_TransMax_r17 struct {
	Value aper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *RA_InformationCommon_r16_msgA_TransMax_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode RA_InformationCommon_r16_msgA_TransMax_r17", err)
	}
	return nil
}

func (ie *RA_InformationCommon_r16_msgA_TransMax_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode RA_InformationCommon_r16_msgA_TransMax_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
