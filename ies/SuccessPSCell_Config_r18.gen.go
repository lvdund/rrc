// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SuccessPSCell-Config-r18 (line 26414).

var successPSCellConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdPercentageT304-SCG-r18", Optional: true},
		{Name: "thresholdPercentageT310-SCG-r18", Optional: true},
		{Name: "thresholdPercentageT312-SCG-r18", Optional: true},
	},
}

const (
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_P40    = 0
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_P60    = 1
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_P80    = 2
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare5 = 3
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare4 = 4
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare3 = 5
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare2 = 6
	SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare1 = 7
)

var successPSCellConfigR18ThresholdPercentageT304SCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_P40, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_P60, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_P80, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare5, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare4, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare3, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare2, SuccessPSCell_Config_r18_ThresholdPercentageT304_SCG_r18_Spare1},
}

const (
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_P40    = 0
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_P60    = 1
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_P80    = 2
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare5 = 3
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare4 = 4
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare3 = 5
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare2 = 6
	SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare1 = 7
)

var successPSCellConfigR18ThresholdPercentageT310SCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_P40, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_P60, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_P80, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare5, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare4, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare3, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare2, SuccessPSCell_Config_r18_ThresholdPercentageT310_SCG_r18_Spare1},
}

const (
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P20    = 0
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P40    = 1
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P60    = 2
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P80    = 3
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare4 = 4
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare3 = 5
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare2 = 6
	SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare1 = 7
)

var successPSCellConfigR18ThresholdPercentageT312SCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P20, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P40, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P60, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_P80, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare4, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare3, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare2, SuccessPSCell_Config_r18_ThresholdPercentageT312_SCG_r18_Spare1},
}

type SuccessPSCell_Config_r18 struct {
	ThresholdPercentageT304_SCG_r18 *int64
	ThresholdPercentageT310_SCG_r18 *int64
	ThresholdPercentageT312_SCG_r18 *int64
}

func (ie *SuccessPSCell_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(successPSCellConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThresholdPercentageT304_SCG_r18 != nil, ie.ThresholdPercentageT310_SCG_r18 != nil, ie.ThresholdPercentageT312_SCG_r18 != nil}); err != nil {
		return err
	}

	// 3. thresholdPercentageT304-SCG-r18: enumerated
	{
		if ie.ThresholdPercentageT304_SCG_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ThresholdPercentageT304_SCG_r18, successPSCellConfigR18ThresholdPercentageT304SCGR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. thresholdPercentageT310-SCG-r18: enumerated
	{
		if ie.ThresholdPercentageT310_SCG_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ThresholdPercentageT310_SCG_r18, successPSCellConfigR18ThresholdPercentageT310SCGR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. thresholdPercentageT312-SCG-r18: enumerated
	{
		if ie.ThresholdPercentageT312_SCG_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ThresholdPercentageT312_SCG_r18, successPSCellConfigR18ThresholdPercentageT312SCGR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SuccessPSCell_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(successPSCellConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. thresholdPercentageT304-SCG-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(successPSCellConfigR18ThresholdPercentageT304SCGR18Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdPercentageT304_SCG_r18 = &idx
		}
	}

	// 4. thresholdPercentageT310-SCG-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(successPSCellConfigR18ThresholdPercentageT310SCGR18Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdPercentageT310_SCG_r18 = &idx
		}
	}

	// 5. thresholdPercentageT312-SCG-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(successPSCellConfigR18ThresholdPercentageT312SCGR18Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdPercentageT312_SCG_r18 = &idx
		}
	}

	return nil
}
