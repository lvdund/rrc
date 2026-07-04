// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BH-LogicalChannelIdentity-r16 (line 5188).

var bHLogicalChannelIdentityR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "bh-LogicalChannelIdentity-r16"},
		{Name: "bh-LogicalChannelIdentityExt-r16"},
	},
}

const (
	BH_LogicalChannelIdentity_r16_Bh_LogicalChannelIdentity_r16    = 0
	BH_LogicalChannelIdentity_r16_Bh_LogicalChannelIdentityExt_r16 = 1
)

type BH_LogicalChannelIdentity_r16 struct {
	Choice                           int
	Bh_LogicalChannelIdentity_r16    *LogicalChannelIdentity
	Bh_LogicalChannelIdentityExt_r16 *BH_LogicalChannelIdentity_Ext_r16
}

func (ie *BH_LogicalChannelIdentity_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bHLogicalChannelIdentityR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BH_LogicalChannelIdentity_r16_Bh_LogicalChannelIdentity_r16:
		if err := ie.Bh_LogicalChannelIdentity_r16.Encode(e); err != nil {
			return err
		}
	case BH_LogicalChannelIdentity_r16_Bh_LogicalChannelIdentityExt_r16:
		if err := ie.Bh_LogicalChannelIdentityExt_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BH-LogicalChannelIdentity-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BH_LogicalChannelIdentity_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bHLogicalChannelIdentityR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BH-LogicalChannelIdentity-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BH_LogicalChannelIdentity_r16_Bh_LogicalChannelIdentity_r16:
		ie.Bh_LogicalChannelIdentity_r16 = new(LogicalChannelIdentity)
		if err := ie.Bh_LogicalChannelIdentity_r16.Decode(d); err != nil {
			return err
		}
	case BH_LogicalChannelIdentity_r16_Bh_LogicalChannelIdentityExt_r16:
		ie.Bh_LogicalChannelIdentityExt_r16 = new(BH_LogicalChannelIdentity_Ext_r16)
		if err := ie.Bh_LogicalChannelIdentityExt_r16.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "BH-LogicalChannelIdentity-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
