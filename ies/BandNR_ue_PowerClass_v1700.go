package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_ue_PowerClass_v1700_Enum_pc5 aper.Enumerated = 0
	BandNR_ue_PowerClass_v1700_Enum_pc6 aper.Enumerated = 1
	BandNR_ue_PowerClass_v1700_Enum_pc7 aper.Enumerated = 2
)

type BandNR_ue_PowerClass_v1700 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandNR_ue_PowerClass_v1700) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandNR_ue_PowerClass_v1700", err)
	}
	return nil
}

func (ie *BandNR_ue_PowerClass_v1700) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandNR_ue_PowerClass_v1700", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
