// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1610 (line 18387).

var cAParametersNRDCV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFR-NR-DC-PwrSharingMode1-r16", Optional: true},
		{Name: "intraFR-NR-DC-PwrSharingMode2-r16", Optional: true},
		{Name: "intraFR-NR-DC-DynamicPwrSharing-r16", Optional: true},
		{Name: "asyncNRDC-r16", Optional: true},
	},
}

const (
	CA_ParametersNRDC_v1610_IntraFR_NR_DC_PwrSharingMode1_r16_Supported = 0
)

var cAParametersNRDCV1610IntraFRNRDCPwrSharingMode1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1610_IntraFR_NR_DC_PwrSharingMode1_r16_Supported},
}

const (
	CA_ParametersNRDC_v1610_IntraFR_NR_DC_PwrSharingMode2_r16_Supported = 0
)

var cAParametersNRDCV1610IntraFRNRDCPwrSharingMode2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1610_IntraFR_NR_DC_PwrSharingMode2_r16_Supported},
}

const (
	CA_ParametersNRDC_v1610_IntraFR_NR_DC_DynamicPwrSharing_r16_Short = 0
	CA_ParametersNRDC_v1610_IntraFR_NR_DC_DynamicPwrSharing_r16_Long  = 1
)

var cAParametersNRDCV1610IntraFRNRDCDynamicPwrSharingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1610_IntraFR_NR_DC_DynamicPwrSharing_r16_Short, CA_ParametersNRDC_v1610_IntraFR_NR_DC_DynamicPwrSharing_r16_Long},
}

const (
	CA_ParametersNRDC_v1610_AsyncNRDC_r16_Supported = 0
)

var cAParametersNRDCV1610AsyncNRDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1610_AsyncNRDC_r16_Supported},
}

type CA_ParametersNRDC_v1610 struct {
	IntraFR_NR_DC_PwrSharingMode1_r16   *int64
	IntraFR_NR_DC_PwrSharingMode2_r16   *int64
	IntraFR_NR_DC_DynamicPwrSharing_r16 *int64
	AsyncNRDC_r16                       *int64
}

func (ie *CA_ParametersNRDC_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraFR_NR_DC_PwrSharingMode1_r16 != nil, ie.IntraFR_NR_DC_PwrSharingMode2_r16 != nil, ie.IntraFR_NR_DC_DynamicPwrSharing_r16 != nil, ie.AsyncNRDC_r16 != nil}); err != nil {
		return err
	}

	// 2. intraFR-NR-DC-PwrSharingMode1-r16: enumerated
	{
		if ie.IntraFR_NR_DC_PwrSharingMode1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IntraFR_NR_DC_PwrSharingMode1_r16, cAParametersNRDCV1610IntraFRNRDCPwrSharingMode1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. intraFR-NR-DC-PwrSharingMode2-r16: enumerated
	{
		if ie.IntraFR_NR_DC_PwrSharingMode2_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IntraFR_NR_DC_PwrSharingMode2_r16, cAParametersNRDCV1610IntraFRNRDCPwrSharingMode2R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. intraFR-NR-DC-DynamicPwrSharing-r16: enumerated
	{
		if ie.IntraFR_NR_DC_DynamicPwrSharing_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IntraFR_NR_DC_DynamicPwrSharing_r16, cAParametersNRDCV1610IntraFRNRDCDynamicPwrSharingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. asyncNRDC-r16: enumerated
	{
		if ie.AsyncNRDC_r16 != nil {
			if err := e.EncodeEnumerated(*ie.AsyncNRDC_r16, cAParametersNRDCV1610AsyncNRDCR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraFR-NR-DC-PwrSharingMode1-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1610IntraFRNRDCPwrSharingMode1R16Constraints)
			if err != nil {
				return err
			}
			ie.IntraFR_NR_DC_PwrSharingMode1_r16 = &idx
		}
	}

	// 3. intraFR-NR-DC-PwrSharingMode2-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1610IntraFRNRDCPwrSharingMode2R16Constraints)
			if err != nil {
				return err
			}
			ie.IntraFR_NR_DC_PwrSharingMode2_r16 = &idx
		}
	}

	// 4. intraFR-NR-DC-DynamicPwrSharing-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1610IntraFRNRDCDynamicPwrSharingR16Constraints)
			if err != nil {
				return err
			}
			ie.IntraFR_NR_DC_DynamicPwrSharing_r16 = &idx
		}
	}

	// 5. asyncNRDC-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1610AsyncNRDCR16Constraints)
			if err != nil {
				return err
			}
			ie.AsyncNRDC_r16 = &idx
		}
	}

	return nil
}
