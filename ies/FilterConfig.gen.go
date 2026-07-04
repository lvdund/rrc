// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FilterConfig (line 12797).

var filterConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "filterCoefficientRSRP", Optional: true},
		{Name: "filterCoefficientRSRQ", Optional: true},
		{Name: "filterCoefficientRS-SINR", Optional: true},
	},
}

type FilterConfig struct {
	FilterCoefficientRSRP    *FilterCoefficient
	FilterCoefficientRSRQ    *FilterCoefficient
	FilterCoefficientRS_SINR *FilterCoefficient
}

func (ie *FilterConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(filterConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FilterCoefficientRSRP != nil && ie.FilterCoefficientRSRP.Value != FilterCoefficient_Fc4, ie.FilterCoefficientRSRQ != nil && ie.FilterCoefficientRSRQ.Value != FilterCoefficient_Fc4, ie.FilterCoefficientRS_SINR != nil && ie.FilterCoefficientRS_SINR.Value != FilterCoefficient_Fc4}); err != nil {
		return err
	}

	// 2. filterCoefficientRSRP: ref
	{
		if ie.FilterCoefficientRSRP != nil && ie.FilterCoefficientRSRP.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientRSRP.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. filterCoefficientRSRQ: ref
	{
		if ie.FilterCoefficientRSRQ != nil && ie.FilterCoefficientRSRQ.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientRSRQ.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. filterCoefficientRS-SINR: ref
	{
		if ie.FilterCoefficientRS_SINR != nil && ie.FilterCoefficientRS_SINR.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientRS_SINR.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FilterConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(filterConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. filterCoefficientRSRP: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FilterCoefficientRSRP = new(FilterCoefficient)
			if err := ie.FilterCoefficientRSRP.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientRSRP = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	// 3. filterCoefficientRSRQ: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FilterCoefficientRSRQ = new(FilterCoefficient)
			if err := ie.FilterCoefficientRSRQ.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientRSRQ = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	// 4. filterCoefficientRS-SINR: ref
	{
		if seq.IsComponentPresent(2) {
			ie.FilterCoefficientRS_SINR = new(FilterCoefficient)
			if err := ie.FilterCoefficientRS_SINR.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientRS_SINR = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	return nil
}
