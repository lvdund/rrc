// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DataCollectionCandidateConfig-r19 (line 26597).

var dataCollectionCandidateConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataCollectionServCellIndex-r19"},
		{Name: "dataCollectionCandidateConfigParameterToAddModList-r19", Optional: true},
		{Name: "dataCollectionCandidateConfigParameterToReleaseList-r19", Optional: true},
	},
}

var dataCollectionCandidateConfigR19DataCollectionCandidateConfigParameterToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofDataCollectionCandidateConfigs_r19)

var dataCollectionCandidateConfigR19DataCollectionCandidateConfigParameterToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofDataCollectionCandidateConfigs_r19)

type DataCollectionCandidateConfig_r19 struct {
	DataCollectionServCellIndex_r19                         ServCellIndex
	DataCollectionCandidateConfigParameterToAddModList_r19  []DataCollectionCandidateConfigParameters_r19
	DataCollectionCandidateConfigParameterToReleaseList_r19 []DataCollectionCandidateConfigId_r19
}

func (ie *DataCollectionCandidateConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionCandidateConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCollectionCandidateConfigParameterToAddModList_r19 != nil, ie.DataCollectionCandidateConfigParameterToReleaseList_r19 != nil}); err != nil {
		return err
	}

	// 3. dataCollectionServCellIndex-r19: ref
	{
		if err := ie.DataCollectionServCellIndex_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. dataCollectionCandidateConfigParameterToAddModList-r19: sequence-of
	{
		if ie.DataCollectionCandidateConfigParameterToAddModList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionCandidateConfigR19DataCollectionCandidateConfigParameterToAddModListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionCandidateConfigParameterToAddModList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionCandidateConfigParameterToAddModList_r19 {
				if err := ie.DataCollectionCandidateConfigParameterToAddModList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. dataCollectionCandidateConfigParameterToReleaseList-r19: sequence-of
	{
		if ie.DataCollectionCandidateConfigParameterToReleaseList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionCandidateConfigR19DataCollectionCandidateConfigParameterToReleaseListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionCandidateConfigParameterToReleaseList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionCandidateConfigParameterToReleaseList_r19 {
				if err := ie.DataCollectionCandidateConfigParameterToReleaseList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *DataCollectionCandidateConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionCandidateConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dataCollectionServCellIndex-r19: ref
	{
		if err := ie.DataCollectionServCellIndex_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. dataCollectionCandidateConfigParameterToAddModList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionCandidateConfigR19DataCollectionCandidateConfigParameterToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionCandidateConfigParameterToAddModList_r19 = make([]DataCollectionCandidateConfigParameters_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionCandidateConfigParameterToAddModList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. dataCollectionCandidateConfigParameterToReleaseList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionCandidateConfigR19DataCollectionCandidateConfigParameterToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionCandidateConfigParameterToReleaseList_r19 = make([]DataCollectionCandidateConfigId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionCandidateConfigParameterToReleaseList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
