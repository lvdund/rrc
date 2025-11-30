package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n1  aper.Enumerated = 0
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n2  aper.Enumerated = 1
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n4  aper.Enumerated = 2
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n8  aper.Enumerated = 3
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n10 aper.Enumerated = 4
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n12 aper.Enumerated = 5
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n14 aper.Enumerated = 6
)

type SRS_Resource_resourceMapping_r17_nrofSymbols_r17 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SRS_Resource_resourceMapping_r17_nrofSymbols_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SRS_Resource_resourceMapping_r17_nrofSymbols_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping_r17_nrofSymbols_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SRS_Resource_resourceMapping_r17_nrofSymbols_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
