package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_prs_ProcessingWindowType1A_r17_Enum_option1 aper.Enumerated = 0
	BandNR_prs_ProcessingWindowType1A_r17_Enum_option2 aper.Enumerated = 1
	BandNR_prs_ProcessingWindowType1A_r17_Enum_option3 aper.Enumerated = 2
)

type BandNR_prs_ProcessingWindowType1A_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandNR_prs_ProcessingWindowType1A_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandNR_prs_ProcessingWindowType1A_r17", err)
	}
	return nil
}

func (ie *BandNR_prs_ProcessingWindowType1A_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandNR_prs_ProcessingWindowType1A_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
