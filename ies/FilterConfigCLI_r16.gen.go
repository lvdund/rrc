// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FilterConfigCLI-r16 (line 12803).

var filterConfigCLIR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "filterCoefficientSRS-RSRP-r16", Optional: true},
		{Name: "filterCoefficientCLI-RSSI-r16", Optional: true},
	},
}

type FilterConfigCLI_r16 struct {
	FilterCoefficientSRS_RSRP_r16 *FilterCoefficient
	FilterCoefficientCLI_RSSI_r16 *FilterCoefficient
}

func (ie *FilterConfigCLI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(filterConfigCLIR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FilterCoefficientSRS_RSRP_r16 != nil && ie.FilterCoefficientSRS_RSRP_r16.Value != FilterCoefficient_Fc4, ie.FilterCoefficientCLI_RSSI_r16 != nil && ie.FilterCoefficientCLI_RSSI_r16.Value != FilterCoefficient_Fc4}); err != nil {
		return err
	}

	// 2. filterCoefficientSRS-RSRP-r16: ref
	{
		if ie.FilterCoefficientSRS_RSRP_r16 != nil && ie.FilterCoefficientSRS_RSRP_r16.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientSRS_RSRP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. filterCoefficientCLI-RSSI-r16: ref
	{
		if ie.FilterCoefficientCLI_RSSI_r16 != nil && ie.FilterCoefficientCLI_RSSI_r16.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientCLI_RSSI_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FilterConfigCLI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(filterConfigCLIR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. filterCoefficientSRS-RSRP-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FilterCoefficientSRS_RSRP_r16 = new(FilterCoefficient)
			if err := ie.FilterCoefficientSRS_RSRP_r16.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientSRS_RSRP_r16 = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	// 3. filterCoefficientCLI-RSSI-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FilterCoefficientCLI_RSSI_r16 = new(FilterCoefficient)
			if err := ie.FilterCoefficientCLI_RSSI_r16.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientCLI_RSSI_r16 = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	return nil
}
