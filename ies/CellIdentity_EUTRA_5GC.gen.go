// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellIdentity-EUTRA-5GC (line 5538).

var cellIdentityEUTRA5GCConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellIdentity-EUTRA"},
		{Name: "cellId-index"},
	},
}

const (
	CellIdentity_EUTRA_5GC_CellIdentity_EUTRA = 0
	CellIdentity_EUTRA_5GC_CellId_Index       = 1
)

type CellIdentity_EUTRA_5GC struct {
	Choice             int
	CellIdentity_EUTRA *per.BitString
	CellId_Index       *int64
}

func (ie *CellIdentity_EUTRA_5GC) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(cellIdentityEUTRA5GCConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CellIdentity_EUTRA_5GC_CellIdentity_EUTRA:
		if err := e.EncodeBitString((*ie.CellIdentity_EUTRA), per.FixedSize(28)); err != nil {
			return err
		}
	case CellIdentity_EUTRA_5GC_CellId_Index:
		if err := e.EncodeInteger((*ie.CellId_Index), per.Constrained(1, common.MaxPLMN)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CellIdentity-EUTRA-5GC",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CellIdentity_EUTRA_5GC) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(cellIdentityEUTRA5GCConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CellIdentity-EUTRA-5GC", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CellIdentity_EUTRA_5GC_CellIdentity_EUTRA:
		ie.CellIdentity_EUTRA = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(28))
		if err != nil {
			return err
		}
		(*ie.CellIdentity_EUTRA) = v
	case CellIdentity_EUTRA_5GC_CellId_Index:
		ie.CellId_Index = new(int64)
		v, err := d.DecodeInteger(per.Constrained(1, common.MaxPLMN))
		if err != nil {
			return err
		}
		(*ie.CellId_Index) = v
	default:
		return &per.DecodeError{TypeName: "CellIdentity-EUTRA-5GC", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
