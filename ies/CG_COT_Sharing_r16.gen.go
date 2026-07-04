// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-COT-Sharing-r16 (line 6675).

var cGCOTSharingR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "noCOT-Sharing-r16"},
		{Name: "cot-Sharing-r16"},
	},
}

const (
	CG_COT_Sharing_r16_NoCOT_Sharing_r16 = 0
	CG_COT_Sharing_r16_Cot_Sharing_r16   = 1
)

type CG_COT_Sharing_r16 struct {
	Choice            int
	NoCOT_Sharing_r16 *struct{}
	Cot_Sharing_r16   *struct {
		Duration_r16              int64
		Offset_r16                int64
		ChannelAccessPriority_r16 int64
	}
}

func (ie *CG_COT_Sharing_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(cGCOTSharingR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r16_NoCOT_Sharing_r16:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case CG_COT_Sharing_r16_Cot_Sharing_r16:
		c := (*ie.Cot_Sharing_r16)
		if err := e.EncodeInteger(c.Duration_r16, per.Constrained(1, 39)); err != nil {
			return err
		}
		if err := e.EncodeInteger(c.Offset_r16, per.Constrained(1, 39)); err != nil {
			return err
		}
		if err := e.EncodeInteger(c.ChannelAccessPriority_r16, per.Constrained(1, 4)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CG-COT-Sharing-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CG_COT_Sharing_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(cGCOTSharingR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CG-COT-Sharing-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CG_COT_Sharing_r16_NoCOT_Sharing_r16:
		ie.NoCOT_Sharing_r16 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case CG_COT_Sharing_r16_Cot_Sharing_r16:
		ie.Cot_Sharing_r16 = &struct {
			Duration_r16              int64
			Offset_r16                int64
			ChannelAccessPriority_r16 int64
		}{}
		c := (*ie.Cot_Sharing_r16)
		{
			v, err := d.DecodeInteger(per.Constrained(1, 39))
			if err != nil {
				return err
			}
			c.Duration_r16 = v
		}
		{
			v, err := d.DecodeInteger(per.Constrained(1, 39))
			if err != nil {
				return err
			}
			c.Offset_r16 = v
		}
		{
			v, err := d.DecodeInteger(per.Constrained(1, 4))
			if err != nil {
				return err
			}
			c.ChannelAccessPriority_r16 = v
		}
	default:
		return &per.DecodeError{TypeName: "CG-COT-Sharing-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
