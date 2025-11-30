package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_ResourceSet_srs_PowerControlAdjustmentStates_Enum_sameAsFci2         aper.Enumerated = 0
	SRS_ResourceSet_srs_PowerControlAdjustmentStates_Enum_separateClosedLoop aper.Enumerated = 1
)

type SRS_ResourceSet_srs_PowerControlAdjustmentStates struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SRS_ResourceSet_srs_PowerControlAdjustmentStates) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SRS_ResourceSet_srs_PowerControlAdjustmentStates", err)
	}
	return nil
}

func (ie *SRS_ResourceSet_srs_PowerControlAdjustmentStates) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SRS_ResourceSet_srs_PowerControlAdjustmentStates", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
