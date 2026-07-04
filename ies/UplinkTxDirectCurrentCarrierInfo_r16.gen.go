// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxDirectCurrentCarrierInfo-r16 (line 16428).

var uplinkTxDirectCurrentCarrierInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellIndex-r16"},
		{Name: "servCellInfo-r16"},
	},
}

var uplinkTxDirectCurrentCarrierInfo_r16ServCellInfoR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "bwp-Id-r16"},
		{Name: "deactivatedCarrier-r16"},
	},
}

const (
	UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_Bwp_Id_r16             = 0
	UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_DeactivatedCarrier_r16 = 1
)

const (
	UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_DeactivatedCarrier_r16_Deactivated = 0
)

var uplinkTxDirectCurrentCarrierInfoR16ServCellInfoR16DeactivatedCarrierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_DeactivatedCarrier_r16_Deactivated},
}

type UplinkTxDirectCurrentCarrierInfo_r16 struct {
	ServCellIndex_r16 ServCellIndex
	ServCellInfo_r16  struct {
		Choice                 int
		Bwp_Id_r16             *BWP_Id
		DeactivatedCarrier_r16 *int64
	}
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxDirectCurrentCarrierInfoR16Constraints)
	_ = seq

	// 1. servCellIndex-r16: ref
	{
		if err := ie.ServCellIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. servCellInfo-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(uplinkTxDirectCurrentCarrierInfo_r16ServCellInfoR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ServCellInfo_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ServCellInfo_r16.Choice {
		case UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_Bwp_Id_r16:
			if err := ie.ServCellInfo_r16.Bwp_Id_r16.Encode(e); err != nil {
				return err
			}
		case UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_DeactivatedCarrier_r16:
			if err := e.EncodeEnumerated((*ie.ServCellInfo_r16.DeactivatedCarrier_r16), uplinkTxDirectCurrentCarrierInfoR16ServCellInfoR16DeactivatedCarrierR16Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ServCellInfo_r16.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxDirectCurrentCarrierInfoR16Constraints)
	_ = seq

	// 1. servCellIndex-r16: ref
	{
		if err := ie.ServCellIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. servCellInfo-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(uplinkTxDirectCurrentCarrierInfo_r16ServCellInfoR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ServCellInfo_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_Bwp_Id_r16:
			ie.ServCellInfo_r16.Bwp_Id_r16 = new(BWP_Id)
			if err := ie.ServCellInfo_r16.Bwp_Id_r16.Decode(d); err != nil {
				return err
			}
		case UplinkTxDirectCurrentCarrierInfo_r16_ServCellInfo_r16_DeactivatedCarrier_r16:
			ie.ServCellInfo_r16.DeactivatedCarrier_r16 = new(int64)
			v, err := d.DecodeEnumerated(uplinkTxDirectCurrentCarrierInfoR16ServCellInfoR16DeactivatedCarrierR16Constraints)
			if err != nil {
				return err
			}
			(*ie.ServCellInfo_r16.DeactivatedCarrier_r16) = v
		}
	}

	return nil
}
