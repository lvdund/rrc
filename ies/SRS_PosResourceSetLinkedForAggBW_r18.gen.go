// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosResourceSetLinkedForAggBW-r18 (line 15776).

var sRSPosResourceSetLinkedForAggBWR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosResourceSetLinked-r18"},
		{Name: "srs-LinkedConfig-r18"},
	},
}

var sRS_PosResourceSetLinkedForAggBW_r18SrsLinkedConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrc-connectedState-r18"},
		{Name: "rrc-inactiveState-r18"},
	},
}

const (
	SRS_PosResourceSetLinkedForAggBW_r18_Srs_LinkedConfig_r18_Rrc_ConnectedState_r18 = 0
	SRS_PosResourceSetLinkedForAggBW_r18_Srs_LinkedConfig_r18_Rrc_InactiveState_r18  = 1
)

var sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellAndBWP-Id-r18"},
	},
}

var sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-SpecificCarrier-r18"},
		{Name: "freqInfo-r18"},
	},
}

type SRS_PosResourceSetLinkedForAggBW_r18 struct {
	Srs_PosResourceSetLinked_r18 SRS_PosResourceSetId_r16
	Srs_LinkedConfig_r18         struct {
		Choice                 int
		Rrc_ConnectedState_r18 *struct{ ServingCellAndBWP_Id_r18 ServingCellAndBWP_Id_r17 }
		Rrc_InactiveState_r18  *struct {
			Scs_SpecificCarrier_r18 SCS_SpecificCarrier
			FreqInfo_r18            ARFCN_ValueNR
		}
	}
}

func (ie *SRS_PosResourceSetLinkedForAggBW_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosResourceSetLinkedForAggBWR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. srs-PosResourceSetLinked-r18: ref
	{
		if err := ie.Srs_PosResourceSetLinked_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. srs-LinkedConfig-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_PosResourceSetLinkedForAggBW_r18SrsLinkedConfigR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Srs_LinkedConfig_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Srs_LinkedConfig_r18.Choice {
		case SRS_PosResourceSetLinkedForAggBW_r18_Srs_LinkedConfig_r18_Rrc_ConnectedState_r18:
			c := (*ie.Srs_LinkedConfig_r18.Rrc_ConnectedState_r18)
			sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Seq := e.NewSequenceEncoder(sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Constraints)
			if err := sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.ServingCellAndBWP_Id_r18.Encode(e); err != nil {
				return err
			}
		case SRS_PosResourceSetLinkedForAggBW_r18_Srs_LinkedConfig_r18_Rrc_InactiveState_r18:
			c := (*ie.Srs_LinkedConfig_r18.Rrc_InactiveState_r18)
			sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Seq := e.NewSequenceEncoder(sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Constraints)
			if err := sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.Scs_SpecificCarrier_r18.Encode(e); err != nil {
				return err
			}
			if err := c.FreqInfo_r18.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Srs_LinkedConfig_r18.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SRS_PosResourceSetLinkedForAggBW_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosResourceSetLinkedForAggBWR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. srs-PosResourceSetLinked-r18: ref
	{
		if err := ie.Srs_PosResourceSetLinked_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. srs-LinkedConfig-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_PosResourceSetLinkedForAggBW_r18SrsLinkedConfigR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Srs_LinkedConfig_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_PosResourceSetLinkedForAggBW_r18_Srs_LinkedConfig_r18_Rrc_ConnectedState_r18:
			ie.Srs_LinkedConfig_r18.Rrc_ConnectedState_r18 = &struct{ ServingCellAndBWP_Id_r18 ServingCellAndBWP_Id_r17 }{}
			c := (*ie.Srs_LinkedConfig_r18.Rrc_ConnectedState_r18)
			sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Seq := d.NewSequenceDecoder(sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Constraints)
			if err := sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcConnectedStateR18Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.ServingCellAndBWP_Id_r18.Decode(d); err != nil {
					return err
				}
			}
		case SRS_PosResourceSetLinkedForAggBW_r18_Srs_LinkedConfig_r18_Rrc_InactiveState_r18:
			ie.Srs_LinkedConfig_r18.Rrc_InactiveState_r18 = &struct {
				Scs_SpecificCarrier_r18 SCS_SpecificCarrier
				FreqInfo_r18            ARFCN_ValueNR
			}{}
			c := (*ie.Srs_LinkedConfig_r18.Rrc_InactiveState_r18)
			sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Seq := d.NewSequenceDecoder(sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Constraints)
			if err := sRSPosResourceSetLinkedForAggBWR18SrsLinkedConfigR18RrcInactiveStateR18Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.Scs_SpecificCarrier_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.FreqInfo_r18.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
