// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-Identity-EUTRA-5GC (line 5533).

var pLMNIdentityEUTRA5GCConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "plmn-Identity-EUTRA-5GC"},
		{Name: "plmn-index"},
	},
}

const (
	PLMN_Identity_EUTRA_5GC_Plmn_Identity_EUTRA_5GC = 0
	PLMN_Identity_EUTRA_5GC_Plmn_Index              = 1
)

type PLMN_Identity_EUTRA_5GC struct {
	Choice                  int
	Plmn_Identity_EUTRA_5GC *PLMN_Identity
	Plmn_Index              *int64
}

func (ie *PLMN_Identity_EUTRA_5GC) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(pLMNIdentityEUTRA5GCConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PLMN_Identity_EUTRA_5GC_Plmn_Identity_EUTRA_5GC:
		if err := ie.Plmn_Identity_EUTRA_5GC.Encode(e); err != nil {
			return err
		}
	case PLMN_Identity_EUTRA_5GC_Plmn_Index:
		if err := e.EncodeInteger((*ie.Plmn_Index), per.Constrained(1, common.MaxPLMN)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PLMN-Identity-EUTRA-5GC",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PLMN_Identity_EUTRA_5GC) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(pLMNIdentityEUTRA5GCConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PLMN-Identity-EUTRA-5GC", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PLMN_Identity_EUTRA_5GC_Plmn_Identity_EUTRA_5GC:
		ie.Plmn_Identity_EUTRA_5GC = new(PLMN_Identity)
		if err := ie.Plmn_Identity_EUTRA_5GC.Decode(d); err != nil {
			return err
		}
	case PLMN_Identity_EUTRA_5GC_Plmn_Index:
		ie.Plmn_Index = new(int64)
		v, err := d.DecodeInteger(per.Constrained(1, common.MaxPLMN))
		if err != nil {
			return err
		}
		(*ie.Plmn_Index) = v
	default:
		return &per.DecodeError{TypeName: "PLMN-Identity-EUTRA-5GC", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
