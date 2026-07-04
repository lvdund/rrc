// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-SpatialRelationInfo (line 12287).

var pUCCHSpatialRelationInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-SpatialRelationInfoId"},
		{Name: "servingCellId", Optional: true},
		{Name: "referenceSignal"},
		{Name: "pucch-PathlossReferenceRS-Id"},
		{Name: "p0-PUCCH-Id"},
		{Name: "closedLoopIndex"},
	},
}

var pUCCH_SpatialRelationInfoReferenceSignalConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
		{Name: "srs"},
	},
}

const (
	PUCCH_SpatialRelationInfo_ReferenceSignal_Ssb_Index    = 0
	PUCCH_SpatialRelationInfo_ReferenceSignal_Csi_RS_Index = 1
	PUCCH_SpatialRelationInfo_ReferenceSignal_Srs          = 2
)

const (
	PUCCH_SpatialRelationInfo_ClosedLoopIndex_I0 = 0
	PUCCH_SpatialRelationInfo_ClosedLoopIndex_I1 = 1
)

var pUCCHSpatialRelationInfoClosedLoopIndexConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_SpatialRelationInfo_ClosedLoopIndex_I0, PUCCH_SpatialRelationInfo_ClosedLoopIndex_I1},
}

type PUCCH_SpatialRelationInfo struct {
	Pucch_SpatialRelationInfoId PUCCH_SpatialRelationInfoId
	ServingCellId               *ServCellIndex
	ReferenceSignal             struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
		Srs          *PUCCH_SRS
	}
	Pucch_PathlossReferenceRS_Id PUCCH_PathlossReferenceRS_Id
	P0_PUCCH_Id                  P0_PUCCH_Id
	ClosedLoopIndex              int64
}

func (ie *PUCCH_SpatialRelationInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHSpatialRelationInfoConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServingCellId != nil}); err != nil {
		return err
	}

	// 2. pucch-SpatialRelationInfoId: ref
	{
		if err := ie.Pucch_SpatialRelationInfoId.Encode(e); err != nil {
			return err
		}
	}

	// 3. servingCellId: ref
	{
		if ie.ServingCellId != nil {
			if err := ie.ServingCellId.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. referenceSignal: choice
	{
		choiceEnc := e.NewChoiceEncoder(pUCCH_SpatialRelationInfoReferenceSignalConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal.Choice {
		case PUCCH_SpatialRelationInfo_ReferenceSignal_Ssb_Index:
			if err := ie.ReferenceSignal.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case PUCCH_SpatialRelationInfo_ReferenceSignal_Csi_RS_Index:
			if err := ie.ReferenceSignal.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		case PUCCH_SpatialRelationInfo_ReferenceSignal_Srs:
			if err := ie.ReferenceSignal.Srs.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal.Choice), Constraint: "index out of range"}
		}
	}

	// 5. pucch-PathlossReferenceRS-Id: ref
	{
		if err := ie.Pucch_PathlossReferenceRS_Id.Encode(e); err != nil {
			return err
		}
	}

	// 6. p0-PUCCH-Id: ref
	{
		if err := ie.P0_PUCCH_Id.Encode(e); err != nil {
			return err
		}
	}

	// 7. closedLoopIndex: enumerated
	{
		if err := e.EncodeEnumerated(ie.ClosedLoopIndex, pUCCHSpatialRelationInfoClosedLoopIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_SpatialRelationInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHSpatialRelationInfoConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pucch-SpatialRelationInfoId: ref
	{
		if err := ie.Pucch_SpatialRelationInfoId.Decode(d); err != nil {
			return err
		}
	}

	// 3. servingCellId: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ServingCellId = new(ServCellIndex)
			if err := ie.ServingCellId.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. referenceSignal: choice
	{
		choiceDec := d.NewChoiceDecoder(pUCCH_SpatialRelationInfoReferenceSignalConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PUCCH_SpatialRelationInfo_ReferenceSignal_Ssb_Index:
			ie.ReferenceSignal.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case PUCCH_SpatialRelationInfo_ReferenceSignal_Csi_RS_Index:
			ie.ReferenceSignal.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		case PUCCH_SpatialRelationInfo_ReferenceSignal_Srs:
			ie.ReferenceSignal.Srs = new(PUCCH_SRS)
			if err := ie.ReferenceSignal.Srs.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. pucch-PathlossReferenceRS-Id: ref
	{
		if err := ie.Pucch_PathlossReferenceRS_Id.Decode(d); err != nil {
			return err
		}
	}

	// 6. p0-PUCCH-Id: ref
	{
		if err := ie.P0_PUCCH_Id.Decode(d); err != nil {
			return err
		}
	}

	// 7. closedLoopIndex: enumerated
	{
		v5, err := d.DecodeEnumerated(pUCCHSpatialRelationInfoClosedLoopIndexConstraints)
		if err != nil {
			return err
		}
		ie.ClosedLoopIndex = v5
	}

	return nil
}
