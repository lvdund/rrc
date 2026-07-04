// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DMRS-BundlingPUCCH-Config-r17 (line 7740).

var dMRSBundlingPUCCHConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-DMRS-Bundling-r17", Optional: true},
		{Name: "pucch-TimeDomainWindowLength-r17", Optional: true},
		{Name: "pucch-WindowRestart-r17", Optional: true},
		{Name: "pucch-FrequencyHoppingInterval-r17", Optional: true},
	},
}

const (
	DMRS_BundlingPUCCH_Config_r17_Pucch_DMRS_Bundling_r17_Enabled = 0
)

var dMRSBundlingPUCCHConfigR17PucchDMRSBundlingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_BundlingPUCCH_Config_r17_Pucch_DMRS_Bundling_r17_Enabled},
}

var dMRSBundlingPUCCHConfigR17PucchTimeDomainWindowLengthR17Constraints = per.Constrained(2, 8)

const (
	DMRS_BundlingPUCCH_Config_r17_Pucch_WindowRestart_r17_Enabled = 0
)

var dMRSBundlingPUCCHConfigR17PucchWindowRestartR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_BundlingPUCCH_Config_r17_Pucch_WindowRestart_r17_Enabled},
}

const (
	DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S2  = 0
	DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S4  = 1
	DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S5  = 2
	DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S10 = 3
)

var dMRSBundlingPUCCHConfigR17PucchFrequencyHoppingIntervalR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S2, DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S4, DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S5, DMRS_BundlingPUCCH_Config_r17_Pucch_FrequencyHoppingInterval_r17_S10},
}

type DMRS_BundlingPUCCH_Config_r17 struct {
	Pucch_DMRS_Bundling_r17            *int64
	Pucch_TimeDomainWindowLength_r17   *int64
	Pucch_WindowRestart_r17            *int64
	Pucch_FrequencyHoppingInterval_r17 *int64
}

func (ie *DMRS_BundlingPUCCH_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dMRSBundlingPUCCHConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pucch_DMRS_Bundling_r17 != nil, ie.Pucch_TimeDomainWindowLength_r17 != nil, ie.Pucch_WindowRestart_r17 != nil, ie.Pucch_FrequencyHoppingInterval_r17 != nil}); err != nil {
		return err
	}

	// 3. pucch-DMRS-Bundling-r17: enumerated
	{
		if ie.Pucch_DMRS_Bundling_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_DMRS_Bundling_r17, dMRSBundlingPUCCHConfigR17PucchDMRSBundlingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. pucch-TimeDomainWindowLength-r17: integer
	{
		if ie.Pucch_TimeDomainWindowLength_r17 != nil {
			if err := e.EncodeInteger(*ie.Pucch_TimeDomainWindowLength_r17, dMRSBundlingPUCCHConfigR17PucchTimeDomainWindowLengthR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. pucch-WindowRestart-r17: enumerated
	{
		if ie.Pucch_WindowRestart_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_WindowRestart_r17, dMRSBundlingPUCCHConfigR17PucchWindowRestartR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pucch-FrequencyHoppingInterval-r17: enumerated
	{
		if ie.Pucch_FrequencyHoppingInterval_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_FrequencyHoppingInterval_r17, dMRSBundlingPUCCHConfigR17PucchFrequencyHoppingIntervalR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DMRS_BundlingPUCCH_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dMRSBundlingPUCCHConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pucch-DMRS-Bundling-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(dMRSBundlingPUCCHConfigR17PucchDMRSBundlingR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_DMRS_Bundling_r17 = &idx
		}
	}

	// 4. pucch-TimeDomainWindowLength-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(dMRSBundlingPUCCHConfigR17PucchTimeDomainWindowLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_TimeDomainWindowLength_r17 = &v
		}
	}

	// 5. pucch-WindowRestart-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dMRSBundlingPUCCHConfigR17PucchWindowRestartR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_WindowRestart_r17 = &idx
		}
	}

	// 6. pucch-FrequencyHoppingInterval-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dMRSBundlingPUCCHConfigR17PucchFrequencyHoppingIntervalR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_FrequencyHoppingInterval_r17 = &idx
		}
	}

	return nil
}
