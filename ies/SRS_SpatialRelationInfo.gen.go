// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-SpatialRelationInfo (line 15596).

var sRSSpatialRelationInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId", Optional: true},
		{Name: "referenceSignal"},
	},
}

var sRS_SpatialRelationInfoReferenceSignalConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
		{Name: "srs"},
	},
}

const (
	SRS_SpatialRelationInfo_ReferenceSignal_Ssb_Index    = 0
	SRS_SpatialRelationInfo_ReferenceSignal_Csi_RS_Index = 1
	SRS_SpatialRelationInfo_ReferenceSignal_Srs          = 2
)

type SRS_SpatialRelationInfo struct {
	ServingCellId   *ServCellIndex
	ReferenceSignal struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
		Srs          *struct {
			ResourceId SRS_ResourceId
			UplinkBWP  BWP_Id
		}
	}
}

func (ie *SRS_SpatialRelationInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSSpatialRelationInfoConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServingCellId != nil}); err != nil {
		return err
	}

	// 2. servingCellId: ref
	{
		if ie.ServingCellId != nil {
			if err := ie.ServingCellId.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. referenceSignal: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_SpatialRelationInfoReferenceSignalConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal.Choice {
		case SRS_SpatialRelationInfo_ReferenceSignal_Ssb_Index:
			if err := ie.ReferenceSignal.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case SRS_SpatialRelationInfo_ReferenceSignal_Csi_RS_Index:
			if err := ie.ReferenceSignal.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		case SRS_SpatialRelationInfo_ReferenceSignal_Srs:
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

func (ie *SRS_SpatialRelationInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSSpatialRelationInfoConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. servingCellId: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ServingCellId = new(ServCellIndex)
			if err := ie.ServingCellId.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. referenceSignal: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_SpatialRelationInfoReferenceSignalConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_SpatialRelationInfo_ReferenceSignal_Ssb_Index:
			ie.ReferenceSignal.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case SRS_SpatialRelationInfo_ReferenceSignal_Csi_RS_Index:
			ie.ReferenceSignal.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		case SRS_SpatialRelationInfo_ReferenceSignal_Srs:
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
