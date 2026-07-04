// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TMGI-r17 (line 28644).

var tMGIR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Id-r17"},
		{Name: "serviceId-r17"},
	},
}

var tMGI_r17PlmnIdR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "plmn-Index"},
		{Name: "explicitValue"},
	},
}

const (
	TMGI_r17_Plmn_Id_r17_Plmn_Index    = 0
	TMGI_r17_Plmn_Id_r17_ExplicitValue = 1
)

var tMGIR17ServiceIdR17Constraints = per.FixedSize(3)

type TMGI_r17 struct {
	Plmn_Id_r17 struct {
		Choice        int
		Plmn_Index    *int64
		ExplicitValue *PLMN_Identity
	}
	ServiceId_r17 []byte
}

func (ie *TMGI_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tMGIR17Constraints)
	_ = seq

	// 1. plmn-Id-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(tMGI_r17PlmnIdR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Plmn_Id_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Plmn_Id_r17.Choice {
		case TMGI_r17_Plmn_Id_r17_Plmn_Index:
			if err := e.EncodeInteger((*ie.Plmn_Id_r17.Plmn_Index), per.Constrained(1, common.MaxPLMN)); err != nil {
				return err
			}
		case TMGI_r17_Plmn_Id_r17_ExplicitValue:
			if err := ie.Plmn_Id_r17.ExplicitValue.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Plmn_Id_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 2. serviceId-r17: octet-string
	{
		if err := e.EncodeOctetString(ie.ServiceId_r17, tMGIR17ServiceIdR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TMGI_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tMGIR17Constraints)
	_ = seq

	// 1. plmn-Id-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(tMGI_r17PlmnIdR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Plmn_Id_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TMGI_r17_Plmn_Id_r17_Plmn_Index:
			ie.Plmn_Id_r17.Plmn_Index = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, common.MaxPLMN))
			if err != nil {
				return err
			}
			(*ie.Plmn_Id_r17.Plmn_Index) = v
		case TMGI_r17_Plmn_Id_r17_ExplicitValue:
			ie.Plmn_Id_r17.ExplicitValue = new(PLMN_Identity)
			if err := ie.Plmn_Id_r17.ExplicitValue.Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. serviceId-r17: octet-string
	{
		v1, err := d.DecodeOctetString(tMGIR17ServiceIdR17Constraints)
		if err != nil {
			return err
		}
		ie.ServiceId_r17 = v1
	}

	return nil
}
