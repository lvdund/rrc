// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedDataCollectionAssistanceConfig-r19 (line 26626).

var loggedDataCollectionAssistanceConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "loggedDataCollectionMemoryThreshold-r19", Optional: true},
	},
}

const (
	LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_KB16   = 0
	LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_KB32   = 1
	LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_KB48   = 2
	LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_Spare1 = 3
)

var loggedDataCollectionAssistanceConfigR19LoggedDataCollectionMemoryThresholdR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_KB16, LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_KB32, LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_KB48, LoggedDataCollectionAssistanceConfig_r19_LoggedDataCollectionMemoryThreshold_r19_Spare1},
}

type LoggedDataCollectionAssistanceConfig_r19 struct {
	LoggedDataCollectionMemoryThreshold_r19 *int64
}

func (ie *LoggedDataCollectionAssistanceConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedDataCollectionAssistanceConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LoggedDataCollectionMemoryThreshold_r19 != nil}); err != nil {
		return err
	}

	// 3. loggedDataCollectionMemoryThreshold-r19: enumerated
	{
		if ie.LoggedDataCollectionMemoryThreshold_r19 != nil {
			if err := e.EncodeEnumerated(*ie.LoggedDataCollectionMemoryThreshold_r19, loggedDataCollectionAssistanceConfigR19LoggedDataCollectionMemoryThresholdR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LoggedDataCollectionAssistanceConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedDataCollectionAssistanceConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. loggedDataCollectionMemoryThreshold-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(loggedDataCollectionAssistanceConfigR19LoggedDataCollectionMemoryThresholdR19Constraints)
			if err != nil {
				return err
			}
			ie.LoggedDataCollectionMemoryThreshold_r19 = &idx
		}
	}

	return nil
}
