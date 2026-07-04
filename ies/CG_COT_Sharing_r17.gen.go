// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-COT-Sharing-r17 (line 6684).

var cGCOTSharingR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "noCOT-Sharing-r17"},
		{Name: "cot-Sharing-r17"},
	},
}

const (
	CG_COT_Sharing_r17_NoCOT_Sharing_r17 = 0
	CG_COT_Sharing_r17_Cot_Sharing_r17   = 1
)

type CG_COT_Sharing_r17 struct {
	Choice            int
	NoCOT_Sharing_r17 *struct{}
	Cot_Sharing_r17   *struct {
		Duration_r17 int64
		Offset_r17   int64
	}
}

func (ie *CG_COT_Sharing_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(cGCOTSharingR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r17_NoCOT_Sharing_r17:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case CG_COT_Sharing_r17_Cot_Sharing_r17:
		c := (*ie.Cot_Sharing_r17)
		if err := e.EncodeInteger(c.Duration_r17, per.Constrained(1, 319)); err != nil {
			return err
		}
		if err := e.EncodeInteger(c.Offset_r17, per.Constrained(1, 319)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CG-COT-Sharing-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CG_COT_Sharing_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(cGCOTSharingR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CG-COT-Sharing-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CG_COT_Sharing_r17_NoCOT_Sharing_r17:
		ie.NoCOT_Sharing_r17 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case CG_COT_Sharing_r17_Cot_Sharing_r17:
		ie.Cot_Sharing_r17 = &struct {
			Duration_r17 int64
			Offset_r17   int64
		}{}
		c := (*ie.Cot_Sharing_r17)
		{
			v, err := d.DecodeInteger(per.Constrained(1, 319))
			if err != nil {
				return err
			}
			c.Duration_r17 = v
		}
		{
			v, err := d.DecodeInteger(per.Constrained(1, 319))
			if err != nil {
				return err
			}
			c.Offset_r17 = v
		}
	default:
		return &per.DecodeError{TypeName: "CG-COT-Sharing-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
