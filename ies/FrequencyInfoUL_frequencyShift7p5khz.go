package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FrequencyInfoUL_frequencyShift7p5khz_Enum_true aper.Enumerated = 0
)

type FrequencyInfoUL_frequencyShift7p5khz struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FrequencyInfoUL_frequencyShift7p5khz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FrequencyInfoUL_frequencyShift7p5khz", err)
	}
	return nil
}

func (ie *FrequencyInfoUL_frequencyShift7p5khz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FrequencyInfoUL_frequencyShift7p5khz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
