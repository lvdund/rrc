// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DataCollectionPreference-r19 (line 2822).

var dataCollectionPreferenceR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataCollectionRequest-r19", Optional: true},
		{Name: "dataCollectionPreferredConfigurationList-r19", Optional: true},
		{Name: "dataCollectionStopConfigurationList-r19", Optional: true},
	},
}

const (
	DataCollectionPreference_r19_DataCollectionRequest_r19_True = 0
)

var dataCollectionPreferenceR19DataCollectionRequestR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DataCollectionPreference_r19_DataCollectionRequest_r19_True},
}

var dataCollectionPreferenceR19DataCollectionPreferredConfigurationListR19Constraints = per.SizeRange(1, common.MaxNrofServingCells)

var dataCollectionPreferenceR19DataCollectionStopConfigurationListR19Constraints = per.SizeRange(1, common.MaxNrofServingCells)

type DataCollectionPreference_r19 struct {
	DataCollectionRequest_r19                    *int64
	DataCollectionPreferredConfigurationList_r19 []DataCollectionCandidateList_r19
	DataCollectionStopConfigurationList_r19      []DataCollectionList_r19
}

func (ie *DataCollectionPreference_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionPreferenceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCollectionRequest_r19 != nil, ie.DataCollectionPreferredConfigurationList_r19 != nil, ie.DataCollectionStopConfigurationList_r19 != nil}); err != nil {
		return err
	}

	// 3. dataCollectionRequest-r19: enumerated
	{
		if ie.DataCollectionRequest_r19 != nil {
			if err := e.EncodeEnumerated(*ie.DataCollectionRequest_r19, dataCollectionPreferenceR19DataCollectionRequestR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. dataCollectionPreferredConfigurationList-r19: sequence-of
	{
		if ie.DataCollectionPreferredConfigurationList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionPreferenceR19DataCollectionPreferredConfigurationListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionPreferredConfigurationList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionPreferredConfigurationList_r19 {
				if err := ie.DataCollectionPreferredConfigurationList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. dataCollectionStopConfigurationList-r19: sequence-of
	{
		if ie.DataCollectionStopConfigurationList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionPreferenceR19DataCollectionStopConfigurationListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionStopConfigurationList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionStopConfigurationList_r19 {
				if err := ie.DataCollectionStopConfigurationList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *DataCollectionPreference_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionPreferenceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dataCollectionRequest-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(dataCollectionPreferenceR19DataCollectionRequestR19Constraints)
			if err != nil {
				return err
			}
			ie.DataCollectionRequest_r19 = &idx
		}
	}

	// 4. dataCollectionPreferredConfigurationList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionPreferenceR19DataCollectionPreferredConfigurationListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionPreferredConfigurationList_r19 = make([]DataCollectionCandidateList_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionPreferredConfigurationList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. dataCollectionStopConfigurationList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionPreferenceR19DataCollectionStopConfigurationListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionStopConfigurationList_r19 = make([]DataCollectionList_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionStopConfigurationList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
