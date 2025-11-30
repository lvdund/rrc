package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC3_r16_duration_r16_Enum_sf1 aper.Enumerated = 0
	SSB_MTC3_r16_duration_r16_Enum_sf2 aper.Enumerated = 1
	SSB_MTC3_r16_duration_r16_Enum_sf3 aper.Enumerated = 2
	SSB_MTC3_r16_duration_r16_Enum_sf4 aper.Enumerated = 3
	SSB_MTC3_r16_duration_r16_Enum_sf5 aper.Enumerated = 4
)

type SSB_MTC3_r16_duration_r16 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SSB_MTC3_r16_duration_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SSB_MTC3_r16_duration_r16", err)
	}
	return nil
}

func (ie *SSB_MTC3_r16_duration_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SSB_MTC3_r16_duration_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
