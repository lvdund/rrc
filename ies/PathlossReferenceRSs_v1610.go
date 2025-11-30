package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PathlossReferenceRSs_v1610 struct {
	Value []PUCCH_PathlossReferenceRS_r16 `lb:1,ub:maxNrofPUCCH_PathlossReferenceRSsDiff_r16,madatory`
}

func (ie *PathlossReferenceRSs_v1610) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PUCCH_PathlossReferenceRS_r16]([]*PUCCH_PathlossReferenceRS_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofPUCCH_PathlossReferenceRSsDiff_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PathlossReferenceRSs_v1610", err)
	}
	return nil
}

func (ie *PathlossReferenceRSs_v1610) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PUCCH_PathlossReferenceRS_r16]([]*PUCCH_PathlossReferenceRS_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofPUCCH_PathlossReferenceRSsDiff_r16}, false)
	fn := func() *PUCCH_PathlossReferenceRS_r16 {
		return new(PUCCH_PathlossReferenceRS_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PathlossReferenceRSs_v1610", err)
	}
	ie.Value = []PUCCH_PathlossReferenceRS_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
