// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DataCollectionPreferenceConfig-r19 (line 26589).

var dataCollectionPreferenceConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataCollectionCandidateConfigToAddModList-r19", Optional: true},
		{Name: "dataCollectionCandidateConfigToReleaseList-r19", Optional: true},
	},
}

var dataCollectionPreferenceConfigR19DataCollectionCandidateConfigToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofServingCells)

var dataCollectionPreferenceConfigR19DataCollectionCandidateConfigToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofServingCells)

type DataCollectionPreferenceConfig_r19 struct {
	DataCollectionCandidateConfigToAddModList_r19  []DataCollectionCandidateConfig_r19
	DataCollectionCandidateConfigToReleaseList_r19 []ServCellIndex
}

func (ie *DataCollectionPreferenceConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionPreferenceConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCollectionCandidateConfigToAddModList_r19 != nil, ie.DataCollectionCandidateConfigToReleaseList_r19 != nil}); err != nil {
		return err
	}

	// 3. dataCollectionCandidateConfigToAddModList-r19: sequence-of
	{
		if ie.DataCollectionCandidateConfigToAddModList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionPreferenceConfigR19DataCollectionCandidateConfigToAddModListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionCandidateConfigToAddModList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionCandidateConfigToAddModList_r19 {
				if err := ie.DataCollectionCandidateConfigToAddModList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. dataCollectionCandidateConfigToReleaseList-r19: sequence-of
	{
		if ie.DataCollectionCandidateConfigToReleaseList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionPreferenceConfigR19DataCollectionCandidateConfigToReleaseListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionCandidateConfigToReleaseList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionCandidateConfigToReleaseList_r19 {
				if err := ie.DataCollectionCandidateConfigToReleaseList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *DataCollectionPreferenceConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionPreferenceConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dataCollectionCandidateConfigToAddModList-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionPreferenceConfigR19DataCollectionCandidateConfigToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionCandidateConfigToAddModList_r19 = make([]DataCollectionCandidateConfig_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionCandidateConfigToAddModList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. dataCollectionCandidateConfigToReleaseList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionPreferenceConfigR19DataCollectionCandidateConfigToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionCandidateConfigToReleaseList_r19 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionCandidateConfigToReleaseList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
