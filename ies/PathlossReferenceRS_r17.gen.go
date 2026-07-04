// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PathlossReferenceRS-r17 (line 10923).

var pathlossReferenceRSR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pathlossReferenceRS-Id-r17"},
		{Name: "referenceSignal-r17"},
		{Name: "additionalPCI-r17", Optional: true},
	},
}

var pathlossReferenceRS_r17ReferenceSignalR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	PathlossReferenceRS_r17_ReferenceSignal_r17_Ssb_Index    = 0
	PathlossReferenceRS_r17_ReferenceSignal_r17_Csi_RS_Index = 1
)

type PathlossReferenceRS_r17 struct {
	PathlossReferenceRS_Id_r17 PathlossReferenceRS_Id_r17
	ReferenceSignal_r17        struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
	}
	AdditionalPCI_r17 *AdditionalPCIIndex_r17
}

func (ie *PathlossReferenceRS_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pathlossReferenceRSR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalPCI_r17 != nil}); err != nil {
		return err
	}

	// 2. pathlossReferenceRS-Id-r17: ref
	{
		if err := ie.PathlossReferenceRS_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. referenceSignal-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(pathlossReferenceRS_r17ReferenceSignalR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal_r17.Choice {
		case PathlossReferenceRS_r17_ReferenceSignal_r17_Ssb_Index:
			if err := ie.ReferenceSignal_r17.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case PathlossReferenceRS_r17_ReferenceSignal_r17_Csi_RS_Index:
			if err := ie.ReferenceSignal_r17.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal_r17.Choice), Constraint: "index out of range"}
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

	return nil
}

func (ie *PathlossReferenceRS_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pathlossReferenceRSR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pathlossReferenceRS-Id-r17: ref
	{
		if err := ie.PathlossReferenceRS_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. referenceSignal-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(pathlossReferenceRS_r17ReferenceSignalR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PathlossReferenceRS_r17_ReferenceSignal_r17_Ssb_Index:
			ie.ReferenceSignal_r17.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal_r17.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case PathlossReferenceRS_r17_ReferenceSignal_r17_Csi_RS_Index:
			ie.ReferenceSignal_r17.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal_r17.Csi_RS_Index.Decode(d); err != nil {
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

	return nil
}
