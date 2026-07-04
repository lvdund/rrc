// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RadioLinkMonitoringRS (line 13214).

var radioLinkMonitoringRSConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "radioLinkMonitoringRS-Id"},
		{Name: "purpose"},
		{Name: "detectionResource"},
	},
}

const (
	RadioLinkMonitoringRS_Purpose_BeamFailure = 0
	RadioLinkMonitoringRS_Purpose_Rlf         = 1
	RadioLinkMonitoringRS_Purpose_Both        = 2
)

var radioLinkMonitoringRSPurposeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioLinkMonitoringRS_Purpose_BeamFailure, RadioLinkMonitoringRS_Purpose_Rlf, RadioLinkMonitoringRS_Purpose_Both},
}

var radioLinkMonitoringRSDetectionResourceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	RadioLinkMonitoringRS_DetectionResource_Ssb_Index    = 0
	RadioLinkMonitoringRS_DetectionResource_Csi_RS_Index = 1
)

type RadioLinkMonitoringRS struct {
	RadioLinkMonitoringRS_Id RadioLinkMonitoringRS_Id
	Purpose                  int64
	DetectionResource        struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
	}
}

func (ie *RadioLinkMonitoringRS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(radioLinkMonitoringRSConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. radioLinkMonitoringRS-Id: ref
	{
		if err := ie.RadioLinkMonitoringRS_Id.Encode(e); err != nil {
			return err
		}
	}

	// 3. purpose: enumerated
	{
		if err := e.EncodeEnumerated(ie.Purpose, radioLinkMonitoringRSPurposeConstraints); err != nil {
			return err
		}
	}

	// 4. detectionResource: choice
	{
		choiceEnc := e.NewChoiceEncoder(radioLinkMonitoringRSDetectionResourceConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.DetectionResource.Choice), false, nil); err != nil {
			return err
		}
		switch ie.DetectionResource.Choice {
		case RadioLinkMonitoringRS_DetectionResource_Ssb_Index:
			if err := ie.DetectionResource.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case RadioLinkMonitoringRS_DetectionResource_Csi_RS_Index:
			if err := ie.DetectionResource.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.DetectionResource.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RadioLinkMonitoringRS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(radioLinkMonitoringRSConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. radioLinkMonitoringRS-Id: ref
	{
		if err := ie.RadioLinkMonitoringRS_Id.Decode(d); err != nil {
			return err
		}
	}

	// 3. purpose: enumerated
	{
		v1, err := d.DecodeEnumerated(radioLinkMonitoringRSPurposeConstraints)
		if err != nil {
			return err
		}
		ie.Purpose = v1
	}

	// 4. detectionResource: choice
	{
		choiceDec := d.NewChoiceDecoder(radioLinkMonitoringRSDetectionResourceConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.DetectionResource.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RadioLinkMonitoringRS_DetectionResource_Ssb_Index:
			ie.DetectionResource.Ssb_Index = new(SSB_Index)
			if err := ie.DetectionResource.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case RadioLinkMonitoringRS_DetectionResource_Csi_RS_Index:
			ie.DetectionResource.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.DetectionResource.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
