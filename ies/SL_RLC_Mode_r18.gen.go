// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RLC-Mode-r18 (line 2333).

var sLRLCModeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-AM-Mode-r18"},
		{Name: "sl-UM-Mode-r18"},
	},
}

const (
	SL_RLC_Mode_r18_Sl_AM_Mode_r18 = 0
	SL_RLC_Mode_r18_Sl_UM_Mode_r18 = 1
)

type SL_RLC_Mode_r18 struct {
	Choice         int
	Sl_AM_Mode_r18 *struct{}
	Sl_UM_Mode_r18 *struct{}
}

func (ie *SL_RLC_Mode_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLRLCModeR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RLC_Mode_r18_Sl_AM_Mode_r18:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case SL_RLC_Mode_r18_Sl_UM_Mode_r18:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-RLC-Mode-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_RLC_Mode_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLRLCModeR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-RLC-Mode-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_RLC_Mode_r18_Sl_AM_Mode_r18:
		ie.Sl_AM_Mode_r18 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case SL_RLC_Mode_r18_Sl_UM_Mode_r18:
		ie.Sl_UM_Mode_r18 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "SL-RLC-Mode-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
