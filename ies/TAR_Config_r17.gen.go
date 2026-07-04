// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TAR-Config-r17 (line 15966).

var tARConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetThresholdTA-r17", Optional: true},
		{Name: "timingAdvanceSR-r17", Optional: true},
	},
}

const (
	TAR_Config_r17_OffsetThresholdTA_r17_Ms0dot5 = 0
	TAR_Config_r17_OffsetThresholdTA_r17_Ms1     = 1
	TAR_Config_r17_OffsetThresholdTA_r17_Ms2     = 2
	TAR_Config_r17_OffsetThresholdTA_r17_Ms3     = 3
	TAR_Config_r17_OffsetThresholdTA_r17_Ms4     = 4
	TAR_Config_r17_OffsetThresholdTA_r17_Ms5     = 5
	TAR_Config_r17_OffsetThresholdTA_r17_Ms6     = 6
	TAR_Config_r17_OffsetThresholdTA_r17_Ms7     = 7
	TAR_Config_r17_OffsetThresholdTA_r17_Ms8     = 8
	TAR_Config_r17_OffsetThresholdTA_r17_Ms9     = 9
	TAR_Config_r17_OffsetThresholdTA_r17_Ms10    = 10
	TAR_Config_r17_OffsetThresholdTA_r17_Ms11    = 11
	TAR_Config_r17_OffsetThresholdTA_r17_Ms12    = 12
	TAR_Config_r17_OffsetThresholdTA_r17_Ms13    = 13
	TAR_Config_r17_OffsetThresholdTA_r17_Ms14    = 14
	TAR_Config_r17_OffsetThresholdTA_r17_Ms15    = 15
	TAR_Config_r17_OffsetThresholdTA_r17_Spare13 = 16
	TAR_Config_r17_OffsetThresholdTA_r17_Spare12 = 17
	TAR_Config_r17_OffsetThresholdTA_r17_Spare11 = 18
	TAR_Config_r17_OffsetThresholdTA_r17_Spare10 = 19
	TAR_Config_r17_OffsetThresholdTA_r17_Spare9  = 20
	TAR_Config_r17_OffsetThresholdTA_r17_Spare8  = 21
	TAR_Config_r17_OffsetThresholdTA_r17_Spare7  = 22
	TAR_Config_r17_OffsetThresholdTA_r17_Spare6  = 23
	TAR_Config_r17_OffsetThresholdTA_r17_Spare5  = 24
	TAR_Config_r17_OffsetThresholdTA_r17_Spare4  = 25
	TAR_Config_r17_OffsetThresholdTA_r17_Spare3  = 26
	TAR_Config_r17_OffsetThresholdTA_r17_Spare2  = 27
	TAR_Config_r17_OffsetThresholdTA_r17_Spare1  = 28
)

var tARConfigR17OffsetThresholdTAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TAR_Config_r17_OffsetThresholdTA_r17_Ms0dot5, TAR_Config_r17_OffsetThresholdTA_r17_Ms1, TAR_Config_r17_OffsetThresholdTA_r17_Ms2, TAR_Config_r17_OffsetThresholdTA_r17_Ms3, TAR_Config_r17_OffsetThresholdTA_r17_Ms4, TAR_Config_r17_OffsetThresholdTA_r17_Ms5, TAR_Config_r17_OffsetThresholdTA_r17_Ms6, TAR_Config_r17_OffsetThresholdTA_r17_Ms7, TAR_Config_r17_OffsetThresholdTA_r17_Ms8, TAR_Config_r17_OffsetThresholdTA_r17_Ms9, TAR_Config_r17_OffsetThresholdTA_r17_Ms10, TAR_Config_r17_OffsetThresholdTA_r17_Ms11, TAR_Config_r17_OffsetThresholdTA_r17_Ms12, TAR_Config_r17_OffsetThresholdTA_r17_Ms13, TAR_Config_r17_OffsetThresholdTA_r17_Ms14, TAR_Config_r17_OffsetThresholdTA_r17_Ms15, TAR_Config_r17_OffsetThresholdTA_r17_Spare13, TAR_Config_r17_OffsetThresholdTA_r17_Spare12, TAR_Config_r17_OffsetThresholdTA_r17_Spare11, TAR_Config_r17_OffsetThresholdTA_r17_Spare10, TAR_Config_r17_OffsetThresholdTA_r17_Spare9, TAR_Config_r17_OffsetThresholdTA_r17_Spare8, TAR_Config_r17_OffsetThresholdTA_r17_Spare7, TAR_Config_r17_OffsetThresholdTA_r17_Spare6, TAR_Config_r17_OffsetThresholdTA_r17_Spare5, TAR_Config_r17_OffsetThresholdTA_r17_Spare4, TAR_Config_r17_OffsetThresholdTA_r17_Spare3, TAR_Config_r17_OffsetThresholdTA_r17_Spare2, TAR_Config_r17_OffsetThresholdTA_r17_Spare1},
}

const (
	TAR_Config_r17_TimingAdvanceSR_r17_Enabled = 0
)

var tARConfigR17TimingAdvanceSRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TAR_Config_r17_TimingAdvanceSR_r17_Enabled},
}

type TAR_Config_r17 struct {
	OffsetThresholdTA_r17 *int64
	TimingAdvanceSR_r17   *int64
}

func (ie *TAR_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tARConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OffsetThresholdTA_r17 != nil, ie.TimingAdvanceSR_r17 != nil}); err != nil {
		return err
	}

	// 3. offsetThresholdTA-r17: enumerated
	{
		if ie.OffsetThresholdTA_r17 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetThresholdTA_r17, tARConfigR17OffsetThresholdTAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. timingAdvanceSR-r17: enumerated
	{
		if ie.TimingAdvanceSR_r17 != nil {
			if err := e.EncodeEnumerated(*ie.TimingAdvanceSR_r17, tARConfigR17TimingAdvanceSRR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TAR_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tARConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. offsetThresholdTA-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(tARConfigR17OffsetThresholdTAR17Constraints)
			if err != nil {
				return err
			}
			ie.OffsetThresholdTA_r17 = &idx
		}
	}

	// 4. timingAdvanceSR-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(tARConfigR17TimingAdvanceSRR17Constraints)
			if err != nil {
				return err
			}
			ie.TimingAdvanceSR_r17 = &idx
		}
	}

	return nil
}
