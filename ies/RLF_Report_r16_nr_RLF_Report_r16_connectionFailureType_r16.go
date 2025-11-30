package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16_Enum_rlf aper.Enumerated = 0
	RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16_Enum_hof aper.Enumerated = 1
)

type RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16", err)
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
