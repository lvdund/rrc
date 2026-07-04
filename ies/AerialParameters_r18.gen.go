// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AerialParameters-r18 (line 16477).

var aerialParametersR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aerialUE-Capability-r18", Optional: true},
		{Name: "altitudeMeas-r18", Optional: true},
		{Name: "altitudeBasedSSB-ToMeasure-r18", Optional: true},
		{Name: "eventAxHy-r18", Optional: true},
		{Name: "flightPathReporting-r18", Optional: true},
		{Name: "flightPathAvailabilityIndicationUAI-r18", Optional: true},
		{Name: "multipleCellsMeasExtension-r18", Optional: true},
		{Name: "nr-NS-PmaxListAerial-r18", Optional: true},
		{Name: "simulMultiTriggerSingleMeasReport-r18", Optional: true},
		{Name: "sl-A2X-Service-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	AerialParameters_r18_AerialUE_Capability_r18_Supported = 0
)

var aerialParametersR18AerialUECapabilityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_AerialUE_Capability_r18_Supported},
}

const (
	AerialParameters_r18_AltitudeMeas_r18_Supported = 0
)

var aerialParametersR18AltitudeMeasR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_AltitudeMeas_r18_Supported},
}

const (
	AerialParameters_r18_AltitudeBasedSSB_ToMeasure_r18_Supported = 0
)

var aerialParametersR18AltitudeBasedSSBToMeasureR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_AltitudeBasedSSB_ToMeasure_r18_Supported},
}

const (
	AerialParameters_r18_EventAxHy_r18_Supported = 0
)

var aerialParametersR18EventAxHyR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_EventAxHy_r18_Supported},
}

const (
	AerialParameters_r18_FlightPathReporting_r18_Supported = 0
)

var aerialParametersR18FlightPathReportingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_FlightPathReporting_r18_Supported},
}

const (
	AerialParameters_r18_FlightPathAvailabilityIndicationUAI_r18_Supported = 0
)

var aerialParametersR18FlightPathAvailabilityIndicationUAIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_FlightPathAvailabilityIndicationUAI_r18_Supported},
}

const (
	AerialParameters_r18_MultipleCellsMeasExtension_r18_Supported = 0
)

var aerialParametersR18MultipleCellsMeasExtensionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_MultipleCellsMeasExtension_r18_Supported},
}

const (
	AerialParameters_r18_Nr_NS_PmaxListAerial_r18_Supported = 0
)

var aerialParametersR18NrNSPmaxListAerialR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_Nr_NS_PmaxListAerial_r18_Supported},
}

const (
	AerialParameters_r18_SimulMultiTriggerSingleMeasReport_r18_Supported = 0
)

var aerialParametersR18SimulMultiTriggerSingleMeasReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_SimulMultiTriggerSingleMeasReport_r18_Supported},
}

const (
	AerialParameters_r18_Sl_A2X_Service_r18_Brid       = 0
	AerialParameters_r18_Sl_A2X_Service_r18_Daa        = 1
	AerialParameters_r18_Sl_A2X_Service_r18_BridAndDAA = 2
)

var aerialParametersR18SlA2XServiceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_Sl_A2X_Service_r18_Brid, AerialParameters_r18_Sl_A2X_Service_r18_Daa, AerialParameters_r18_Sl_A2X_Service_r18_BridAndDAA},
}

const (
	AerialParameters_r18_Ext_CondEventAxHy_r19_Supported = 0
)

var aerialParametersR18ExtCondEventAxHyR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AerialParameters_r18_Ext_CondEventAxHy_r19_Supported},
}

type AerialParameters_r18 struct {
	AerialUE_Capability_r18                 *int64
	AltitudeMeas_r18                        *int64
	AltitudeBasedSSB_ToMeasure_r18          *int64
	EventAxHy_r18                           *int64
	FlightPathReporting_r18                 *int64
	FlightPathAvailabilityIndicationUAI_r18 *int64
	MultipleCellsMeasExtension_r18          *int64
	Nr_NS_PmaxListAerial_r18                *int64
	SimulMultiTriggerSingleMeasReport_r18   *int64
	Sl_A2X_Service_r18                      *int64
	CondEventAxHy_r19                       *int64
}

func (ie *AerialParameters_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(aerialParametersR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CondEventAxHy_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AerialUE_Capability_r18 != nil, ie.AltitudeMeas_r18 != nil, ie.AltitudeBasedSSB_ToMeasure_r18 != nil, ie.EventAxHy_r18 != nil, ie.FlightPathReporting_r18 != nil, ie.FlightPathAvailabilityIndicationUAI_r18 != nil, ie.MultipleCellsMeasExtension_r18 != nil, ie.Nr_NS_PmaxListAerial_r18 != nil, ie.SimulMultiTriggerSingleMeasReport_r18 != nil, ie.Sl_A2X_Service_r18 != nil}); err != nil {
		return err
	}

	// 3. aerialUE-Capability-r18: enumerated
	{
		if ie.AerialUE_Capability_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AerialUE_Capability_r18, aerialParametersR18AerialUECapabilityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. altitudeMeas-r18: enumerated
	{
		if ie.AltitudeMeas_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AltitudeMeas_r18, aerialParametersR18AltitudeMeasR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. altitudeBasedSSB-ToMeasure-r18: enumerated
	{
		if ie.AltitudeBasedSSB_ToMeasure_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AltitudeBasedSSB_ToMeasure_r18, aerialParametersR18AltitudeBasedSSBToMeasureR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. eventAxHy-r18: enumerated
	{
		if ie.EventAxHy_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EventAxHy_r18, aerialParametersR18EventAxHyR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. flightPathReporting-r18: enumerated
	{
		if ie.FlightPathReporting_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathReporting_r18, aerialParametersR18FlightPathReportingR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. flightPathAvailabilityIndicationUAI-r18: enumerated
	{
		if ie.FlightPathAvailabilityIndicationUAI_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathAvailabilityIndicationUAI_r18, aerialParametersR18FlightPathAvailabilityIndicationUAIR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. multipleCellsMeasExtension-r18: enumerated
	{
		if ie.MultipleCellsMeasExtension_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultipleCellsMeasExtension_r18, aerialParametersR18MultipleCellsMeasExtensionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. nr-NS-PmaxListAerial-r18: enumerated
	{
		if ie.Nr_NS_PmaxListAerial_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Nr_NS_PmaxListAerial_r18, aerialParametersR18NrNSPmaxListAerialR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. simulMultiTriggerSingleMeasReport-r18: enumerated
	{
		if ie.SimulMultiTriggerSingleMeasReport_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SimulMultiTriggerSingleMeasReport_r18, aerialParametersR18SimulMultiTriggerSingleMeasReportR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-A2X-Service-r18: enumerated
	{
		if ie.Sl_A2X_Service_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_A2X_Service_r18, aerialParametersR18SlA2XServiceR18Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "condEventAxHy-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CondEventAxHy_r19 != nil}); err != nil {
				return err
			}

			if ie.CondEventAxHy_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.CondEventAxHy_r19, aerialParametersR18ExtCondEventAxHyR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AerialParameters_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(aerialParametersR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. aerialUE-Capability-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(aerialParametersR18AerialUECapabilityR18Constraints)
			if err != nil {
				return err
			}
			ie.AerialUE_Capability_r18 = &idx
		}
	}

	// 4. altitudeMeas-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(aerialParametersR18AltitudeMeasR18Constraints)
			if err != nil {
				return err
			}
			ie.AltitudeMeas_r18 = &idx
		}
	}

	// 5. altitudeBasedSSB-ToMeasure-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(aerialParametersR18AltitudeBasedSSBToMeasureR18Constraints)
			if err != nil {
				return err
			}
			ie.AltitudeBasedSSB_ToMeasure_r18 = &idx
		}
	}

	// 6. eventAxHy-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(aerialParametersR18EventAxHyR18Constraints)
			if err != nil {
				return err
			}
			ie.EventAxHy_r18 = &idx
		}
	}

	// 7. flightPathReporting-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(aerialParametersR18FlightPathReportingR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathReporting_r18 = &idx
		}
	}

	// 8. flightPathAvailabilityIndicationUAI-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(aerialParametersR18FlightPathAvailabilityIndicationUAIR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathAvailabilityIndicationUAI_r18 = &idx
		}
	}

	// 9. multipleCellsMeasExtension-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(aerialParametersR18MultipleCellsMeasExtensionR18Constraints)
			if err != nil {
				return err
			}
			ie.MultipleCellsMeasExtension_r18 = &idx
		}
	}

	// 10. nr-NS-PmaxListAerial-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(aerialParametersR18NrNSPmaxListAerialR18Constraints)
			if err != nil {
				return err
			}
			ie.Nr_NS_PmaxListAerial_r18 = &idx
		}
	}

	// 11. simulMultiTriggerSingleMeasReport-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(aerialParametersR18SimulMultiTriggerSingleMeasReportR18Constraints)
			if err != nil {
				return err
			}
			ie.SimulMultiTriggerSingleMeasReport_r18 = &idx
		}
	}

	// 12. sl-A2X-Service-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(aerialParametersR18SlA2XServiceR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_A2X_Service_r18 = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "condEventAxHy-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(aerialParametersR18ExtCondEventAxHyR19Constraints)
			if err != nil {
				return err
			}
			ie.CondEventAxHy_r19 = &v
		}
	}

	return nil
}
