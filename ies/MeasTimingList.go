package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasTimingList struct {
	Value []MeasTiming `lb:1,ub:maxMeasFreqsMN,madatory`
}

func (ie *MeasTimingList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasTiming]([]*MeasTiming{}, aper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasTimingList", err)
	}
	return nil
}

func (ie *MeasTimingList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasTiming]([]*MeasTiming{}, aper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
	fn := func() *MeasTiming {
		return new(MeasTiming)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasTimingList", err)
	}
	ie.Value = []MeasTiming{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
