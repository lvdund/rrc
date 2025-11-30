package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRB_CountMSB_InfoList struct {
	Value []DRB_CountMSB_Info `lb:1,ub:maxDRB,madatory`
}

func (ie *DRB_CountMSB_InfoList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*DRB_CountMSB_Info]([]*DRB_CountMSB_Info{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode DRB_CountMSB_InfoList", err)
	}
	return nil
}

func (ie *DRB_CountMSB_InfoList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*DRB_CountMSB_Info]([]*DRB_CountMSB_Info{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn := func() *DRB_CountMSB_Info {
		return new(DRB_CountMSB_Info)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode DRB_CountMSB_InfoList", err)
	}
	ie.Value = []DRB_CountMSB_Info{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
