// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LPWUS-SupportedBandInfo-r19 (line 25962).

var lPWUSSupportedBandInfoR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandIndicator-r19"},
		{Name: "lpwus-OOK-r19", Optional: true},
		{Name: "lpwus-OFDM-r19", Optional: true},
		{Name: "lpwus-LP-SS-r19", Optional: true},
		{Name: "minimumTimeGap-r19"},
	},
}

const (
	LPWUS_SupportedBandInfo_r19_Lpwus_OOK_r19_Supported = 0
)

var lPWUSSupportedBandInfoR19LpwusOOKR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_SupportedBandInfo_r19_Lpwus_OOK_r19_Supported},
}

const (
	LPWUS_SupportedBandInfo_r19_Lpwus_OFDM_r19_Supported = 0
)

var lPWUSSupportedBandInfoR19LpwusOFDMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_SupportedBandInfo_r19_Lpwus_OFDM_r19_Supported},
}

const (
	LPWUS_SupportedBandInfo_r19_Lpwus_LP_SS_r19_Supported = 0
)

var lPWUSSupportedBandInfoR19LpwusLPSSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_SupportedBandInfo_r19_Lpwus_LP_SS_r19_Supported},
}

const (
	LPWUS_SupportedBandInfo_r19_MinimumTimeGap_r19_Cap1 = 0
	LPWUS_SupportedBandInfo_r19_MinimumTimeGap_r19_Cap2 = 1
	LPWUS_SupportedBandInfo_r19_MinimumTimeGap_r19_Cap3 = 2
)

var lPWUSSupportedBandInfoR19MinimumTimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_SupportedBandInfo_r19_MinimumTimeGap_r19_Cap1, LPWUS_SupportedBandInfo_r19_MinimumTimeGap_r19_Cap2, LPWUS_SupportedBandInfo_r19_MinimumTimeGap_r19_Cap3},
}

type LPWUS_SupportedBandInfo_r19 struct {
	SupportedBandIndicator_r19 FreqBandIndicatorNR
	Lpwus_OOK_r19              *int64
	Lpwus_OFDM_r19             *int64
	Lpwus_LP_SS_r19            *int64
	MinimumTimeGap_r19         int64
}

func (ie *LPWUS_SupportedBandInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lPWUSSupportedBandInfoR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lpwus_OOK_r19 != nil, ie.Lpwus_OFDM_r19 != nil, ie.Lpwus_LP_SS_r19 != nil}); err != nil {
		return err
	}

	// 2. supportedBandIndicator-r19: ref
	{
		if err := ie.SupportedBandIndicator_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. lpwus-OOK-r19: enumerated
	{
		if ie.Lpwus_OOK_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_OOK_r19, lPWUSSupportedBandInfoR19LpwusOOKR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. lpwus-OFDM-r19: enumerated
	{
		if ie.Lpwus_OFDM_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_OFDM_r19, lPWUSSupportedBandInfoR19LpwusOFDMR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. lpwus-LP-SS-r19: enumerated
	{
		if ie.Lpwus_LP_SS_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_LP_SS_r19, lPWUSSupportedBandInfoR19LpwusLPSSR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. minimumTimeGap-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.MinimumTimeGap_r19, lPWUSSupportedBandInfoR19MinimumTimeGapR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LPWUS_SupportedBandInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lPWUSSupportedBandInfoR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandIndicator-r19: ref
	{
		if err := ie.SupportedBandIndicator_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. lpwus-OOK-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(lPWUSSupportedBandInfoR19LpwusOOKR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_OOK_r19 = &idx
		}
	}

	// 4. lpwus-OFDM-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(lPWUSSupportedBandInfoR19LpwusOFDMR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_OFDM_r19 = &idx
		}
	}

	// 5. lpwus-LP-SS-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(lPWUSSupportedBandInfoR19LpwusLPSSR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_LP_SS_r19 = &idx
		}
	}

	// 6. minimumTimeGap-r19: enumerated
	{
		v4, err := d.DecodeEnumerated(lPWUSSupportedBandInfoR19MinimumTimeGapR19Constraints)
		if err != nil {
			return err
		}
		ie.MinimumTimeGap_r19 = v4
	}

	return nil
}
