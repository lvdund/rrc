// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PagingUE-Identity (line 878).

var pagingUEIdentityConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ng-5G-S-TMSI"},
		{Name: "fullI-RNTI"},
	},
}

const (
	PagingUE_Identity_Ng_5G_S_TMSI = 0
	PagingUE_Identity_FullI_RNTI   = 1
)

type PagingUE_Identity struct {
	Choice       int
	Ng_5G_S_TMSI *NG_5G_S_TMSI
	FullI_RNTI   *I_RNTI_Value
}

func (ie *PagingUE_Identity) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(pagingUEIdentityConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PagingUE_Identity_Ng_5G_S_TMSI:
		if err := ie.Ng_5G_S_TMSI.Encode(e); err != nil {
			return err
		}
	case PagingUE_Identity_FullI_RNTI:
		if err := ie.FullI_RNTI.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "PagingUE-Identity",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PagingUE_Identity) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(pagingUEIdentityConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PagingUE-Identity", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PagingUE_Identity_Ng_5G_S_TMSI:
		ie.Ng_5G_S_TMSI = new(NG_5G_S_TMSI)
		if err := ie.Ng_5G_S_TMSI.Decode(d); err != nil {
			return err
		}
	case PagingUE_Identity_FullI_RNTI:
		ie.FullI_RNTI = new(I_RNTI_Value)
		if err := ie.FullI_RNTI.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "PagingUE-Identity", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
