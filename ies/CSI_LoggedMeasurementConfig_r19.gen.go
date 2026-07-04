// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-LoggedMeasurementConfig-r19 (line 6978).

var cSILoggedMeasurementConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-LoggedMeasurementConfigId-r19"},
		{Name: "csi-LoggedResourceConfig-r19"},
		{Name: "loggingPeriodicity-r19", Optional: true},
		{Name: "csi-LoggedMeasurementEventTriggerConfig-r19", Optional: true},
	},
}

const (
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N2     = 0
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N3     = 1
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N4     = 2
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N5     = 3
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare4 = 4
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare3 = 5
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare2 = 6
	CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare1 = 7
)

var cSILoggedMeasurementConfigR19LoggingPeriodicityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N2, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N3, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N4, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_N5, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare4, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare3, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare2, CSI_LoggedMeasurementConfig_r19_LoggingPeriodicity_r19_Spare1},
}

type CSI_LoggedMeasurementConfig_r19 struct {
	Csi_LoggedMeasurementConfigId_r19           CSI_LoggedMeasurementConfigId_r19
	Csi_LoggedResourceConfig_r19                CSI_ResourceConfigId
	LoggingPeriodicity_r19                      *int64
	Csi_LoggedMeasurementEventTriggerConfig_r19 *CSI_LoggedMeasurementEventTriggerConfig_r19
}

func (ie *CSI_LoggedMeasurementConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSILoggedMeasurementConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LoggingPeriodicity_r19 != nil, ie.Csi_LoggedMeasurementEventTriggerConfig_r19 != nil}); err != nil {
		return err
	}

	// 3. csi-LoggedMeasurementConfigId-r19: ref
	{
		if err := ie.Csi_LoggedMeasurementConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. csi-LoggedResourceConfig-r19: ref
	{
		if err := ie.Csi_LoggedResourceConfig_r19.Encode(e); err != nil {
			return err
		}
	}

	// 5. loggingPeriodicity-r19: enumerated
	{
		if ie.LoggingPeriodicity_r19 != nil {
			if err := e.EncodeEnumerated(*ie.LoggingPeriodicity_r19, cSILoggedMeasurementConfigR19LoggingPeriodicityR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. csi-LoggedMeasurementEventTriggerConfig-r19: ref
	{
		if ie.Csi_LoggedMeasurementEventTriggerConfig_r19 != nil {
			if err := ie.Csi_LoggedMeasurementEventTriggerConfig_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_LoggedMeasurementConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSILoggedMeasurementConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-LoggedMeasurementConfigId-r19: ref
	{
		if err := ie.Csi_LoggedMeasurementConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. csi-LoggedResourceConfig-r19: ref
	{
		if err := ie.Csi_LoggedResourceConfig_r19.Decode(d); err != nil {
			return err
		}
	}

	// 5. loggingPeriodicity-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cSILoggedMeasurementConfigR19LoggingPeriodicityR19Constraints)
			if err != nil {
				return err
			}
			ie.LoggingPeriodicity_r19 = &idx
		}
	}

	// 6. csi-LoggedMeasurementEventTriggerConfig-r19: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Csi_LoggedMeasurementEventTriggerConfig_r19 = new(CSI_LoggedMeasurementEventTriggerConfig_r19)
			if err := ie.Csi_LoggedMeasurementEventTriggerConfig_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
