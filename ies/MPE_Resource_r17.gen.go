// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MPE-Resource-r17 (line 12499).

var mPEResourceR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mpe-ResourceId-r17"},
		{Name: "cell-r17", Optional: true},
		{Name: "additionalPCI-r17", Optional: true},
		{Name: "mpe-ReferenceSignal-r17"},
	},
}

var mPE_Resource_r17MpeReferenceSignalR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "csi-RS-Resource-r17"},
		{Name: "ssb-Resource-r17"},
	},
}

const (
	MPE_Resource_r17_Mpe_ReferenceSignal_r17_Csi_RS_Resource_r17 = 0
	MPE_Resource_r17_Mpe_ReferenceSignal_r17_Ssb_Resource_r17    = 1
)

type MPE_Resource_r17 struct {
	Mpe_ResourceId_r17      MPE_ResourceId_r17
	Cell_r17                *ServCellIndex
	AdditionalPCI_r17       *AdditionalPCIIndex_r17
	Mpe_ReferenceSignal_r17 struct {
		Choice              int
		Csi_RS_Resource_r17 *NZP_CSI_RS_ResourceId
		Ssb_Resource_r17    *SSB_Index
	}
}

func (ie *MPE_Resource_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mPEResourceR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cell_r17 != nil, ie.AdditionalPCI_r17 != nil}); err != nil {
		return err
	}

	// 2. mpe-ResourceId-r17: ref
	{
		if err := ie.Mpe_ResourceId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. cell-r17: ref
	{
		if ie.Cell_r17 != nil {
			if err := ie.Cell_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. additionalPCI-r17: ref
	{
		if ie.AdditionalPCI_r17 != nil {
			if err := ie.AdditionalPCI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. mpe-ReferenceSignal-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(mPE_Resource_r17MpeReferenceSignalR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Mpe_ReferenceSignal_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Mpe_ReferenceSignal_r17.Choice {
		case MPE_Resource_r17_Mpe_ReferenceSignal_r17_Csi_RS_Resource_r17:
			if err := ie.Mpe_ReferenceSignal_r17.Csi_RS_Resource_r17.Encode(e); err != nil {
				return err
			}
		case MPE_Resource_r17_Mpe_ReferenceSignal_r17_Ssb_Resource_r17:
			if err := ie.Mpe_ReferenceSignal_r17.Ssb_Resource_r17.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Mpe_ReferenceSignal_r17.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MPE_Resource_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mPEResourceR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mpe-ResourceId-r17: ref
	{
		if err := ie.Mpe_ResourceId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. cell-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Cell_r17 = new(ServCellIndex)
			if err := ie.Cell_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. additionalPCI-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
			if err := ie.AdditionalPCI_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. mpe-ReferenceSignal-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(mPE_Resource_r17MpeReferenceSignalR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Mpe_ReferenceSignal_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MPE_Resource_r17_Mpe_ReferenceSignal_r17_Csi_RS_Resource_r17:
			ie.Mpe_ReferenceSignal_r17.Csi_RS_Resource_r17 = new(NZP_CSI_RS_ResourceId)
			if err := ie.Mpe_ReferenceSignal_r17.Csi_RS_Resource_r17.Decode(d); err != nil {
				return err
			}
		case MPE_Resource_r17_Mpe_ReferenceSignal_r17_Ssb_Resource_r17:
			ie.Mpe_ReferenceSignal_r17.Ssb_Resource_r17 = new(SSB_Index)
			if err := ie.Mpe_ReferenceSignal_r17.Ssb_Resource_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
