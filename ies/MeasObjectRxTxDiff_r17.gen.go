// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasObjectRxTxDiff-r17 (line 9633).

var measObjectRxTxDiffR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-Ref-r17", Optional: true},
	},
}

var measObjectRxTxDiff_r17DlRefR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "prs-Ref-r17"},
		{Name: "csi-RS-Ref-r17"},
	},
}

const (
	MeasObjectRxTxDiff_r17_Dl_Ref_r17_Prs_Ref_r17    = 0
	MeasObjectRxTxDiff_r17_Dl_Ref_r17_Csi_RS_Ref_r17 = 1
)

type MeasObjectRxTxDiff_r17 struct {
	Dl_Ref_r17 *struct {
		Choice         int
		Prs_Ref_r17    *struct{}
		Csi_RS_Ref_r17 *struct{}
	}
}

func (ie *MeasObjectRxTxDiff_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectRxTxDiffR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_Ref_r17 != nil}); err != nil {
		return err
	}

	// 3. dl-Ref-r17: choice
	{
		if ie.Dl_Ref_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(measObjectRxTxDiff_r17DlRefR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Dl_Ref_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Dl_Ref_r17).Choice {
			case MeasObjectRxTxDiff_r17_Dl_Ref_r17_Prs_Ref_r17:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MeasObjectRxTxDiff_r17_Dl_Ref_r17_Csi_RS_Ref_r17:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Dl_Ref_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *MeasObjectRxTxDiff_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectRxTxDiffR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dl-Ref-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Dl_Ref_r17 = &struct {
				Choice         int
				Prs_Ref_r17    *struct{}
				Csi_RS_Ref_r17 *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(measObjectRxTxDiff_r17DlRefR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_Ref_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MeasObjectRxTxDiff_r17_Dl_Ref_r17_Prs_Ref_r17:
				(*ie.Dl_Ref_r17).Prs_Ref_r17 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MeasObjectRxTxDiff_r17_Dl_Ref_r17_Csi_RS_Ref_r17:
				(*ie.Dl_Ref_r17).Csi_RS_Ref_r17 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
