// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DataCollectionList-r19 (line 2834).

var dataCollectionListR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dataCollectionStopServCellIndex-r19"},
		{Name: "dataCollectionIdList-r19", Optional: true},
	},
}

var dataCollectionListR19DataCollectionIdListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_ReportConfigurations)

type DataCollectionList_r19 struct {
	DataCollectionStopServCellIndex_r19 ServCellIndex
	DataCollectionIdList_r19            []CSI_ReportConfigId
}

func (ie *DataCollectionList_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionListR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCollectionIdList_r19 != nil}); err != nil {
		return err
	}

	// 2. dataCollectionStopServCellIndex-r19: ref
	{
		if err := ie.DataCollectionStopServCellIndex_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. dataCollectionIdList-r19: sequence-of
	{
		if ie.DataCollectionIdList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(dataCollectionListR19DataCollectionIdListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DataCollectionIdList_r19))); err != nil {
				return err
			}
			for i := range ie.DataCollectionIdList_r19 {
				if err := ie.DataCollectionIdList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *DataCollectionList_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionListR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dataCollectionStopServCellIndex-r19: ref
	{
		if err := ie.DataCollectionStopServCellIndex_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. dataCollectionIdList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(dataCollectionListR19DataCollectionIdListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DataCollectionIdList_r19 = make([]CSI_ReportConfigId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DataCollectionIdList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
