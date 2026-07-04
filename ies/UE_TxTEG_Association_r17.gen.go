// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-TxTEG-Association-r17 (line 3594).

var uETxTEGAssociationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-TxTEG-ID-r17"},
		{Name: "nr-TimeStamp-r17"},
		{Name: "associatedSRS-PosResourceIdList-r17"},
		{Name: "servCellId-r17", Optional: true},
	},
}

var uETxTEGAssociationR17UeTxTEGIDR17Constraints = per.Constrained(0, common.MaxNrOfTxTEG_ID_1_r17)

var uETxTEGAssociationR17AssociatedSRSPosResourceIdListR17Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResources_r16)

type UE_TxTEG_Association_r17 struct {
	Ue_TxTEG_ID_r17                     int64
	Nr_TimeStamp_r17                    NR_TimeStamp_r17
	AssociatedSRS_PosResourceIdList_r17 []SRS_PosResourceId_r16
	ServCellId_r17                      *ServCellIndex
}

func (ie *UE_TxTEG_Association_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uETxTEGAssociationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServCellId_r17 != nil}); err != nil {
		return err
	}

	// 2. ue-TxTEG-ID-r17: integer
	{
		if err := e.EncodeInteger(ie.Ue_TxTEG_ID_r17, uETxTEGAssociationR17UeTxTEGIDR17Constraints); err != nil {
			return err
		}
	}

	// 3. nr-TimeStamp-r17: ref
	{
		if err := ie.Nr_TimeStamp_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. associatedSRS-PosResourceIdList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uETxTEGAssociationR17AssociatedSRSPosResourceIdListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.AssociatedSRS_PosResourceIdList_r17))); err != nil {
			return err
		}
		for i := range ie.AssociatedSRS_PosResourceIdList_r17 {
			if err := ie.AssociatedSRS_PosResourceIdList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. servCellId-r17: ref
	{
		if ie.ServCellId_r17 != nil {
			if err := ie.ServCellId_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_TxTEG_Association_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uETxTEGAssociationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ue-TxTEG-ID-r17: integer
	{
		v0, err := d.DecodeInteger(uETxTEGAssociationR17UeTxTEGIDR17Constraints)
		if err != nil {
			return err
		}
		ie.Ue_TxTEG_ID_r17 = v0
	}

	// 3. nr-TimeStamp-r17: ref
	{
		if err := ie.Nr_TimeStamp_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. associatedSRS-PosResourceIdList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uETxTEGAssociationR17AssociatedSRSPosResourceIdListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.AssociatedSRS_PosResourceIdList_r17 = make([]SRS_PosResourceId_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.AssociatedSRS_PosResourceIdList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. servCellId-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.ServCellId_r17 = new(ServCellIndex)
			if err := ie.ServCellId_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
