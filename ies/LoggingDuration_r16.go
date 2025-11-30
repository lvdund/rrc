package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LoggingDuration_r16_Enum_min10  aper.Enumerated = 0
	LoggingDuration_r16_Enum_min20  aper.Enumerated = 1
	LoggingDuration_r16_Enum_min40  aper.Enumerated = 2
	LoggingDuration_r16_Enum_min60  aper.Enumerated = 3
	LoggingDuration_r16_Enum_min90  aper.Enumerated = 4
	LoggingDuration_r16_Enum_min120 aper.Enumerated = 5
	LoggingDuration_r16_Enum_spare2 aper.Enumerated = 6
	LoggingDuration_r16_Enum_spare1 aper.Enumerated = 7
)

type LoggingDuration_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *LoggingDuration_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode LoggingDuration_r16", err)
	}
	return nil
}

func (ie *LoggingDuration_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode LoggingDuration_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
