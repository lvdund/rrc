package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n3   aper.Enumerated = 0
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n4   aper.Enumerated = 1
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n5   aper.Enumerated = 2
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n6   aper.Enumerated = 3
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n7   aper.Enumerated = 4
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n8   aper.Enumerated = 5
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n10  aper.Enumerated = 6
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n20  aper.Enumerated = 7
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n50  aper.Enumerated = 8
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n100 aper.Enumerated = 9
	RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16_Enum_n200 aper.Enumerated = 10
)

type RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16", err)
	}
	return nil
}

func (ie *RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
