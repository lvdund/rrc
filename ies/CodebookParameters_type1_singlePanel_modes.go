package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookParameters_type1_singlePanel_modes_Enum_mode1         aper.Enumerated = 0
	CodebookParameters_type1_singlePanel_modes_Enum_mode1andMode2 aper.Enumerated = 1
)

type CodebookParameters_type1_singlePanel_modes struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CodebookParameters_type1_singlePanel_modes) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CodebookParameters_type1_singlePanel_modes", err)
	}
	return nil
}

func (ie *CodebookParameters_type1_singlePanel_modes) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CodebookParameters_type1_singlePanel_modes", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
