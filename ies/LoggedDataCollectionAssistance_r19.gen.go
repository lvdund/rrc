// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedDataCollectionAssistance-r19 (line 2839).

var loggedDataCollectionAssistanceR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lowPowerState-r19", Optional: true},
		{Name: "memoryStatus-r19", Optional: true},
	},
}

const (
	LoggedDataCollectionAssistance_r19_LowPowerState_r19_True = 0
)

var loggedDataCollectionAssistanceR19LowPowerStateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggedDataCollectionAssistance_r19_LowPowerState_r19_True},
}

const (
	LoggedDataCollectionAssistance_r19_MemoryStatus_r19_Full           = 0
	LoggedDataCollectionAssistance_r19_MemoryStatus_r19_AboveThreshold = 1
)

var loggedDataCollectionAssistanceR19MemoryStatusR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggedDataCollectionAssistance_r19_MemoryStatus_r19_Full, LoggedDataCollectionAssistance_r19_MemoryStatus_r19_AboveThreshold},
}

type LoggedDataCollectionAssistance_r19 struct {
	LowPowerState_r19 *int64
	MemoryStatus_r19  *int64
}

func (ie *LoggedDataCollectionAssistance_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedDataCollectionAssistanceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LowPowerState_r19 != nil, ie.MemoryStatus_r19 != nil}); err != nil {
		return err
	}

	// 3. lowPowerState-r19: enumerated
	{
		if ie.LowPowerState_r19 != nil {
			if err := e.EncodeEnumerated(*ie.LowPowerState_r19, loggedDataCollectionAssistanceR19LowPowerStateR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. memoryStatus-r19: enumerated
	{
		if ie.MemoryStatus_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MemoryStatus_r19, loggedDataCollectionAssistanceR19MemoryStatusR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LoggedDataCollectionAssistance_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedDataCollectionAssistanceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. lowPowerState-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(loggedDataCollectionAssistanceR19LowPowerStateR19Constraints)
			if err != nil {
				return err
			}
			ie.LowPowerState_r19 = &idx
		}
	}

	// 4. memoryStatus-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(loggedDataCollectionAssistanceR19MemoryStatusR19Constraints)
			if err != nil {
				return err
			}
			ie.MemoryStatus_r19 = &idx
		}
	}

	return nil
}
