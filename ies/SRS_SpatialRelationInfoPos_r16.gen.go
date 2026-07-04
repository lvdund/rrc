// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-SpatialRelationInfoPos-r16 (line 15608).

var sRSSpatialRelationInfoPosR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "servingRS-r16"},
		{Name: "ssb-Ncell-r16"},
		{Name: "dl-PRS-r16"},
	},
}

const (
	SRS_SpatialRelationInfoPos_r16_ServingRS_r16 = 0
	SRS_SpatialRelationInfoPos_r16_Ssb_Ncell_r16 = 1
	SRS_SpatialRelationInfoPos_r16_Dl_PRS_r16    = 2
)

var sRSSpatialRelationInfoPosR16ServingRSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId", Optional: true},
		{Name: "referenceSignal-r16"},
	},
}

var sRSSpatialRelationInfoPosR16ServingRSR16ReferenceSignalR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-IndexServing-r16"},
		{Name: "csi-RS-IndexServing-r16"},
		{Name: "srs-SpatialRelation-r16"},
	},
}

const (
	SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Ssb_IndexServing_r16    = 0
	SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Csi_RS_IndexServing_r16 = 1
	SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16 = 2
)

var sRSSpatialRelationInfoPosR16ServingRSR16ReferenceSignalR16SrsSpatialRelationR16ResourceSelectionR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "srs-ResourceId-r16"},
		{Name: "srs-PosResourceId-r16"},
	},
}

const (
	SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16_ResourceSelection_r16_Srs_ResourceId_r16    = 0
	SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16_ResourceSelection_r16_Srs_PosResourceId_r16 = 1
)

type SRS_SpatialRelationInfoPos_r16 struct {
	Choice        int
	ServingRS_r16 *struct {
		ServingCellId       *ServCellIndex
		ReferenceSignal_r16 struct {
			Choice                  int
			Ssb_IndexServing_r16    *SSB_Index
			Csi_RS_IndexServing_r16 *NZP_CSI_RS_ResourceId
			Srs_SpatialRelation_r16 *struct {
				ResourceSelection_r16 struct {
					Choice                int
					Srs_ResourceId_r16    *SRS_ResourceId
					Srs_PosResourceId_r16 *SRS_PosResourceId_r16
				}
				UplinkBWP_r16 BWP_Id
			}
		}
	}
	Ssb_Ncell_r16 *SSB_InfoNcell_r16
	Dl_PRS_r16    *DL_PRS_Info_r16
}

func (ie *SRS_SpatialRelationInfoPos_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sRSSpatialRelationInfoPosR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_ServingRS_r16:
		c := (*ie.ServingRS_r16)
		sRSSpatialRelationInfoPosR16ServingRSR16Seq := e.NewSequenceEncoder(sRSSpatialRelationInfoPosR16ServingRSR16Constraints)
		if err := sRSSpatialRelationInfoPosR16ServingRSR16Seq.EncodePreamble([]bool{c.ServingCellId != nil}); err != nil {
			return err
		}
		if c.ServingCellId != nil {
			if err := c.ServingCellId.Encode(e); err != nil {
				return err
			}
		}
		{
			choiceEnc := e.NewChoiceEncoder(sRSSpatialRelationInfoPosR16ServingRSR16ReferenceSignalR16Constraints)
			if err := choiceEnc.EncodeChoice(int64(c.ReferenceSignal_r16.Choice), false, nil); err != nil {
				return err
			}
			switch c.ReferenceSignal_r16.Choice {
			case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Ssb_IndexServing_r16:
				if err := c.ReferenceSignal_r16.Ssb_IndexServing_r16.Encode(e); err != nil {
					return err
				}
			case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Csi_RS_IndexServing_r16:
				if err := c.ReferenceSignal_r16.Csi_RS_IndexServing_r16.Encode(e); err != nil {
					return err
				}
			case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16:
				c := (*c.ReferenceSignal_r16.Srs_SpatialRelation_r16)
				{
					choiceEnc := e.NewChoiceEncoder(sRSSpatialRelationInfoPosR16ServingRSR16ReferenceSignalR16SrsSpatialRelationR16ResourceSelectionR16Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.ResourceSelection_r16.Choice), false, nil); err != nil {
						return err
					}
					switch c.ResourceSelection_r16.Choice {
					case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16_ResourceSelection_r16_Srs_ResourceId_r16:
						if err := c.ResourceSelection_r16.Srs_ResourceId_r16.Encode(e); err != nil {
							return err
						}
					case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16_ResourceSelection_r16_Srs_PosResourceId_r16:
						if err := c.ResourceSelection_r16.Srs_PosResourceId_r16.Encode(e); err != nil {
							return err
						}
					}
				}
				if err := c.UplinkBWP_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	case SRS_SpatialRelationInfoPos_r16_Ssb_Ncell_r16:
		if err := ie.Ssb_Ncell_r16.Encode(e); err != nil {
			return err
		}
	case SRS_SpatialRelationInfoPos_r16_Dl_PRS_r16:
		if err := ie.Dl_PRS_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SRS-SpatialRelationInfoPos-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SRS_SpatialRelationInfoPos_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sRSSpatialRelationInfoPosR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SRS-SpatialRelationInfoPos-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SRS_SpatialRelationInfoPos_r16_ServingRS_r16:
		ie.ServingRS_r16 = &struct {
			ServingCellId       *ServCellIndex
			ReferenceSignal_r16 struct {
				Choice                  int
				Ssb_IndexServing_r16    *SSB_Index
				Csi_RS_IndexServing_r16 *NZP_CSI_RS_ResourceId
				Srs_SpatialRelation_r16 *struct {
					ResourceSelection_r16 struct {
						Choice                int
						Srs_ResourceId_r16    *SRS_ResourceId
						Srs_PosResourceId_r16 *SRS_PosResourceId_r16
					}
					UplinkBWP_r16 BWP_Id
				}
			}
		}{}
		c := (*ie.ServingRS_r16)
		sRSSpatialRelationInfoPosR16ServingRSR16Seq := d.NewSequenceDecoder(sRSSpatialRelationInfoPosR16ServingRSR16Constraints)
		if err := sRSSpatialRelationInfoPosR16ServingRSR16Seq.DecodePreamble(); err != nil {
			return err
		}
		if sRSSpatialRelationInfoPosR16ServingRSR16Seq.IsComponentPresent(0) {
			c.ServingCellId = new(ServCellIndex)
			if err := (*c.ServingCellId).Decode(d); err != nil {
				return err
			}
		}
		{
			choiceDec := d.NewChoiceDecoder(sRSSpatialRelationInfoPosR16ServingRSR16ReferenceSignalR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			c.ReferenceSignal_r16.Choice = int(index)
			switch index {
			case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Ssb_IndexServing_r16:
				c.ReferenceSignal_r16.Ssb_IndexServing_r16 = new(SSB_Index)
				if err := c.ReferenceSignal_r16.Ssb_IndexServing_r16.Decode(d); err != nil {
					return err
				}
			case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Csi_RS_IndexServing_r16:
				c.ReferenceSignal_r16.Csi_RS_IndexServing_r16 = new(NZP_CSI_RS_ResourceId)
				if err := c.ReferenceSignal_r16.Csi_RS_IndexServing_r16.Decode(d); err != nil {
					return err
				}
			case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16:
				c.ReferenceSignal_r16.Srs_SpatialRelation_r16 = &struct {
					ResourceSelection_r16 struct {
						Choice                int
						Srs_ResourceId_r16    *SRS_ResourceId
						Srs_PosResourceId_r16 *SRS_PosResourceId_r16
					}
					UplinkBWP_r16 BWP_Id
				}{}
				c := (*c.ReferenceSignal_r16.Srs_SpatialRelation_r16)
				{
					choiceDec := d.NewChoiceDecoder(sRSSpatialRelationInfoPosR16ServingRSR16ReferenceSignalR16SrsSpatialRelationR16ResourceSelectionR16Constraints)
					index, _, _, err := choiceDec.DecodeChoice()
					if err != nil {
						return err
					}
					c.ResourceSelection_r16.Choice = int(index)
					switch index {
					case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16_ResourceSelection_r16_Srs_ResourceId_r16:
						c.ResourceSelection_r16.Srs_ResourceId_r16 = new(SRS_ResourceId)
						if err := c.ResourceSelection_r16.Srs_ResourceId_r16.Decode(d); err != nil {
							return err
						}
					case SRS_SpatialRelationInfoPos_r16_ServingRS_r16_ReferenceSignal_r16_Srs_SpatialRelation_r16_ResourceSelection_r16_Srs_PosResourceId_r16:
						c.ResourceSelection_r16.Srs_PosResourceId_r16 = new(SRS_PosResourceId_r16)
						if err := c.ResourceSelection_r16.Srs_PosResourceId_r16.Decode(d); err != nil {
							return err
						}
					}
				}
				{
					if err := c.UplinkBWP_r16.Decode(d); err != nil {
						return err
					}
				}
			}
		}
	case SRS_SpatialRelationInfoPos_r16_Ssb_Ncell_r16:
		ie.Ssb_Ncell_r16 = new(SSB_InfoNcell_r16)
		if err := ie.Ssb_Ncell_r16.Decode(d); err != nil {
			return err
		}
	case SRS_SpatialRelationInfoPos_r16_Dl_PRS_r16:
		ie.Dl_PRS_r16 = new(DL_PRS_Info_r16)
		if err := ie.Dl_PRS_r16.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "SRS-SpatialRelationInfoPos-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
