// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DataCollectionCandidateList-r19 (line 2829).

var dataCollectionCandidateListR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dataCollectionServCellIndex-r19"},
		{Name: "dataCollectionCandidateIdList-r19", Optional: true},
	},
}

var dataCollectionCandidateListR19DataCollectionCandidateIdListR19Constraints = per.SizeRange(1, common.MaxNrofDataCollectionCandidateConfigs_r19)

type DataCollectionCandidateList_r19 struct {
	DataCollectionServCellIndex_r19   ServCellIndex
	DataCollectionCandidateIdList_r19 []DataCollectionCandidateConfigId_r19
}

func (ie *DataCollectionCandidateList_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionCandidateListR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCollectionCandidateIdList_r19 != nil}); err != nil {
		return err
	}

	// 2. dataCollectionServCellIndex-r19: ref
	{
		if err := ie.DataCollectionServCellIndex_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. dataCollectionCandidateIdList-r19: sequence-of
	{
		if ie.DataCollectionCandidateIdList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionCandidateListR19DataCollectionCandidateIdListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionCandidateIdList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionCandidateIdList_r19 {
				if err := ie.DataCollectionCandidateIdList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *DataCollectionCandidateList_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionCandidateListR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dataCollectionServCellIndex-r19: ref
	{
		if err := ie.DataCollectionServCellIndex_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. dataCollectionCandidateIdList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionCandidateListR19DataCollectionCandidateIdListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionCandidateIdList_r19 = make([]DataCollectionCandidateConfigId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionCandidateIdList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
