// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DMRS-BundlingPUSCH-Config-r17 (line 7751).

var dMRSBundlingPUSCHConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-DMRS-Bundling-r17", Optional: true},
		{Name: "pusch-TimeDomainWindowLength-r17", Optional: true},
		{Name: "pusch-WindowRestart-r17", Optional: true},
		{Name: "pusch-FrequencyHoppingInterval-r17", Optional: true},
	},
}

const (
	DMRS_BundlingPUSCH_Config_r17_Pusch_DMRS_Bundling_r17_Enabled = 0
)

var dMRSBundlingPUSCHConfigR17PuschDMRSBundlingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_BundlingPUSCH_Config_r17_Pusch_DMRS_Bundling_r17_Enabled},
}

var dMRSBundlingPUSCHConfigR17PuschTimeDomainWindowLengthR17Constraints = per.Constrained(2, 32)

const (
	DMRS_BundlingPUSCH_Config_r17_Pusch_WindowRestart_r17_Enabled = 0
)

var dMRSBundlingPUSCHConfigR17PuschWindowRestartR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_BundlingPUSCH_Config_r17_Pusch_WindowRestart_r17_Enabled},
}

const (
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S2  = 0
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S4  = 1
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S5  = 2
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S6  = 3
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S8  = 4
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S10 = 5
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S12 = 6
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S14 = 7
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S16 = 8
	DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S20 = 9
)

var dMRSBundlingPUSCHConfigR17PuschFrequencyHoppingIntervalR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S2, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S4, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S5, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S6, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S8, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S10, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S12, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S14, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S16, DMRS_BundlingPUSCH_Config_r17_Pusch_FrequencyHoppingInterval_r17_S20},
}

type DMRS_BundlingPUSCH_Config_r17 struct {
	Pusch_DMRS_Bundling_r17            *int64
	Pusch_TimeDomainWindowLength_r17   *int64
	Pusch_WindowRestart_r17            *int64
	Pusch_FrequencyHoppingInterval_r17 *int64
}

func (ie *DMRS_BundlingPUSCH_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dMRSBundlingPUSCHConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pusch_DMRS_Bundling_r17 != nil, ie.Pusch_TimeDomainWindowLength_r17 != nil, ie.Pusch_WindowRestart_r17 != nil, ie.Pusch_FrequencyHoppingInterval_r17 != nil}); err != nil {
		return err
	}

	// 3. pusch-DMRS-Bundling-r17: enumerated
	{
		if ie.Pusch_DMRS_Bundling_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_DMRS_Bundling_r17, dMRSBundlingPUSCHConfigR17PuschDMRSBundlingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. pusch-TimeDomainWindowLength-r17: integer
	{
		if ie.Pusch_TimeDomainWindowLength_r17 != nil {
			if err := e.EncodeInteger(*ie.Pusch_TimeDomainWindowLength_r17, dMRSBundlingPUSCHConfigR17PuschTimeDomainWindowLengthR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. pusch-WindowRestart-r17: enumerated
	{
		if ie.Pusch_WindowRestart_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_WindowRestart_r17, dMRSBundlingPUSCHConfigR17PuschWindowRestartR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pusch-FrequencyHoppingInterval-r17: enumerated
	{
		if ie.Pusch_FrequencyHoppingInterval_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_FrequencyHoppingInterval_r17, dMRSBundlingPUSCHConfigR17PuschFrequencyHoppingIntervalR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DMRS_BundlingPUSCH_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dMRSBundlingPUSCHConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pusch-DMRS-Bundling-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(dMRSBundlingPUSCHConfigR17PuschDMRSBundlingR17Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_DMRS_Bundling_r17 = &idx
		}
	}

	// 4. pusch-TimeDomainWindowLength-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(dMRSBundlingPUSCHConfigR17PuschTimeDomainWindowLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_TimeDomainWindowLength_r17 = &v
		}
	}

	// 5. pusch-WindowRestart-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dMRSBundlingPUSCHConfigR17PuschWindowRestartR17Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_WindowRestart_r17 = &idx
		}
	}

	// 6. pusch-FrequencyHoppingInterval-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dMRSBundlingPUSCHConfigR17PuschFrequencyHoppingIntervalR17Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_FrequencyHoppingInterval_r17 = &idx
		}
	}

	return nil
}
