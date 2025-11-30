package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_CodeBlockGroupTransmissionList_r16 struct {
	Value []PDSCH_CodeBlockGroupTransmission `lb:1,ub:2,madatory`
}

func (ie *PDSCH_CodeBlockGroupTransmissionList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PDSCH_CodeBlockGroupTransmission]([]*PDSCH_CodeBlockGroupTransmission{}, aper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PDSCH_CodeBlockGroupTransmissionList_r16", err)
	}
	return nil
}

func (ie *PDSCH_CodeBlockGroupTransmissionList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PDSCH_CodeBlockGroupTransmission]([]*PDSCH_CodeBlockGroupTransmission{}, aper.Constraint{Lb: 1, Ub: 2}, false)
	fn := func() *PDSCH_CodeBlockGroupTransmission {
		return new(PDSCH_CodeBlockGroupTransmission)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PDSCH_CodeBlockGroupTransmissionList_r16", err)
	}
	ie.Value = []PDSCH_CodeBlockGroupTransmission{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
