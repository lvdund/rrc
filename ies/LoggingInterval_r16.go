package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LoggingInterval_r16_Enum_ms320    aper.Enumerated = 0
	LoggingInterval_r16_Enum_ms640    aper.Enumerated = 1
	LoggingInterval_r16_Enum_ms1280   aper.Enumerated = 2
	LoggingInterval_r16_Enum_ms2560   aper.Enumerated = 3
	LoggingInterval_r16_Enum_ms5120   aper.Enumerated = 4
	LoggingInterval_r16_Enum_ms10240  aper.Enumerated = 5
	LoggingInterval_r16_Enum_ms20480  aper.Enumerated = 6
	LoggingInterval_r16_Enum_ms30720  aper.Enumerated = 7
	LoggingInterval_r16_Enum_ms40960  aper.Enumerated = 8
	LoggingInterval_r16_Enum_ms61440  aper.Enumerated = 9
	LoggingInterval_r16_Enum_infinity aper.Enumerated = 10
)

type LoggingInterval_r16 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *LoggingInterval_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode LoggingInterval_r16", err)
	}
	return nil
}

func (ie *LoggingInterval_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode LoggingInterval_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
