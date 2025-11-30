package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasResultListWLAN_r16 struct {
	Value []LogMeasResultWLAN_r16 `lb:1,ub:maxWLAN_Id_Report_r16,madatory`
}

func (ie *LogMeasResultListWLAN_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*LogMeasResultWLAN_r16]([]*LogMeasResultWLAN_r16{}, aper.Constraint{Lb: 1, Ub: maxWLAN_Id_Report_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode LogMeasResultListWLAN_r16", err)
	}
	return nil
}

func (ie *LogMeasResultListWLAN_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*LogMeasResultWLAN_r16]([]*LogMeasResultWLAN_r16{}, aper.Constraint{Lb: 1, Ub: maxWLAN_Id_Report_r16}, false)
	fn := func() *LogMeasResultWLAN_r16 {
		return new(LogMeasResultWLAN_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode LogMeasResultListWLAN_r16", err)
	}
	ie.Value = []LogMeasResultWLAN_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
