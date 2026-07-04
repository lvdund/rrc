// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BeamLinkMonitoringRS-r17 (line 13234).

var beamLinkMonitoringRSR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "beamLinkMonitoringRS-Id-r17"},
		{Name: "detectionResource-r17"},
	},
}

var beamLinkMonitoringRS_r17DetectionResourceR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	BeamLinkMonitoringRS_r17_DetectionResource_r17_Ssb_Index    = 0
	BeamLinkMonitoringRS_r17_DetectionResource_r17_Csi_RS_Index = 1
)

type BeamLinkMonitoringRS_r17 struct {
	BeamLinkMonitoringRS_Id_r17 BeamLinkMonitoringRS_Id_r17
	DetectionResource_r17       struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
	}
}

func (ie *BeamLinkMonitoringRS_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamLinkMonitoringRSR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. beamLinkMonitoringRS-Id-r17: ref
	{
		if err := ie.BeamLinkMonitoringRS_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. detectionResource-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(beamLinkMonitoringRS_r17DetectionResourceR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.DetectionResource_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.DetectionResource_r17.Choice {
		case BeamLinkMonitoringRS_r17_DetectionResource_r17_Ssb_Index:
			if err := ie.DetectionResource_r17.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case BeamLinkMonitoringRS_r17_DetectionResource_r17_Csi_RS_Index:
			if err := ie.DetectionResource_r17.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.DetectionResource_r17.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *BeamLinkMonitoringRS_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamLinkMonitoringRSR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. beamLinkMonitoringRS-Id-r17: ref
	{
		if err := ie.BeamLinkMonitoringRS_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. detectionResource-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(beamLinkMonitoringRS_r17DetectionResourceR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.DetectionResource_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case BeamLinkMonitoringRS_r17_DetectionResource_r17_Ssb_Index:
			ie.DetectionResource_r17.Ssb_Index = new(SSB_Index)
			if err := ie.DetectionResource_r17.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case BeamLinkMonitoringRS_r17_DetectionResource_r17_Csi_RS_Index:
			ie.DetectionResource_r17.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.DetectionResource_r17.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
