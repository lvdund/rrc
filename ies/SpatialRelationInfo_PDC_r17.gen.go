// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SpatialRelationInfo-PDC-r17 (line 15707).

var spatialRelationInfoPDCR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "referenceSignal"},
	},
}

var spatialRelationInfo_PDC_r17ReferenceSignalConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
		{Name: "dl-PRS-PDC"},
		{Name: "srs"},
	},
}

const (
	SpatialRelationInfo_PDC_r17_ReferenceSignal_Ssb_Index    = 0
	SpatialRelationInfo_PDC_r17_ReferenceSignal_Csi_RS_Index = 1
	SpatialRelationInfo_PDC_r17_ReferenceSignal_Dl_PRS_PDC   = 2
	SpatialRelationInfo_PDC_r17_ReferenceSignal_Srs          = 3
)

type SpatialRelationInfo_PDC_r17 struct {
	ReferenceSignal struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
		Dl_PRS_PDC   *NR_DL_PRS_ResourceID_r17
		Srs          *struct {
			ResourceId SRS_ResourceId
			UplinkBWP  BWP_Id
		}
	}
}

func (ie *SpatialRelationInfo_PDC_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(spatialRelationInfoPDCR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. referenceSignal: choice
	{
		choiceEnc := e.NewChoiceEncoder(spatialRelationInfo_PDC_r17ReferenceSignalConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal.Choice {
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Ssb_Index:
			if err := ie.ReferenceSignal.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Csi_RS_Index:
			if err := ie.ReferenceSignal.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Dl_PRS_PDC:
			if err := ie.ReferenceSignal.Dl_PRS_PDC.Encode(e); err != nil {
				return err
			}
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Srs:
			c := (*ie.ReferenceSignal.Srs)
			if err := c.ResourceId.Encode(e); err != nil {
				return err
			}
			if err := c.UplinkBWP.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SpatialRelationInfo_PDC_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(spatialRelationInfoPDCR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. referenceSignal: choice
	{
		choiceDec := d.NewChoiceDecoder(spatialRelationInfo_PDC_r17ReferenceSignalConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Ssb_Index:
			ie.ReferenceSignal.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Csi_RS_Index:
			ie.ReferenceSignal.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Dl_PRS_PDC:
			ie.ReferenceSignal.Dl_PRS_PDC = new(NR_DL_PRS_ResourceID_r17)
			if err := ie.ReferenceSignal.Dl_PRS_PDC.Decode(d); err != nil {
				return err
			}
		case SpatialRelationInfo_PDC_r17_ReferenceSignal_Srs:
			ie.ReferenceSignal.Srs = &struct {
				ResourceId SRS_ResourceId
				UplinkBWP  BWP_Id
			}{}
			c := (*ie.ReferenceSignal.Srs)
			{
				if err := c.ResourceId.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.UplinkBWP.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
