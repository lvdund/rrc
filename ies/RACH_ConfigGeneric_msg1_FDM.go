package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_msg1_FDM_Enum_one   aper.Enumerated = 0
	RACH_ConfigGeneric_msg1_FDM_Enum_two   aper.Enumerated = 1
	RACH_ConfigGeneric_msg1_FDM_Enum_four  aper.Enumerated = 2
	RACH_ConfigGeneric_msg1_FDM_Enum_eight aper.Enumerated = 3
)

type RACH_ConfigGeneric_msg1_FDM struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RACH_ConfigGeneric_msg1_FDM) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_msg1_FDM", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_msg1_FDM) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_msg1_FDM", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
