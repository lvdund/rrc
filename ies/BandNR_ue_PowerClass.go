package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_ue_PowerClass_Enum_pc1 aper.Enumerated = 0
	BandNR_ue_PowerClass_Enum_pc2 aper.Enumerated = 1
	BandNR_ue_PowerClass_Enum_pc3 aper.Enumerated = 2
	BandNR_ue_PowerClass_Enum_pc4 aper.Enumerated = 3
)

type BandNR_ue_PowerClass struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BandNR_ue_PowerClass) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BandNR_ue_PowerClass", err)
	}
	return nil
}

func (ie *BandNR_ue_PowerClass) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BandNR_ue_PowerClass", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
