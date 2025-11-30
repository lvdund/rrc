package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms5    aper.Enumerated = 0
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms10   aper.Enumerated = 1
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms20   aper.Enumerated = 2
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms40   aper.Enumerated = 3
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms80   aper.Enumerated = 4
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms160  aper.Enumerated = 5
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_spare2 aper.Enumerated = 6
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_spare1 aper.Enumerated = 7
)

type SSB_Configuration_r16_ssb_Periodicity_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SSB_Configuration_r16_ssb_Periodicity_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SSB_Configuration_r16_ssb_Periodicity_r16", err)
	}
	return nil
}

func (ie *SSB_Configuration_r16_ssb_Periodicity_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SSB_Configuration_r16_ssb_Periodicity_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
