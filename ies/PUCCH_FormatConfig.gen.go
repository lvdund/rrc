// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-FormatConfig (line 12047).

var pUCCHFormatConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interslotFrequencyHopping", Optional: true},
		{Name: "additionalDMRS", Optional: true},
		{Name: "maxCodeRate", Optional: true},
		{Name: "nrofSlots", Optional: true},
		{Name: "pi2BPSK", Optional: true},
		{Name: "simultaneousHARQ-ACK-CSI", Optional: true},
	},
}

const (
	PUCCH_FormatConfig_InterslotFrequencyHopping_Enabled = 0
)

var pUCCHFormatConfigInterslotFrequencyHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_FormatConfig_InterslotFrequencyHopping_Enabled},
}

const (
	PUCCH_FormatConfig_AdditionalDMRS_True = 0
)

var pUCCHFormatConfigAdditionalDMRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_FormatConfig_AdditionalDMRS_True},
}

const (
	PUCCH_FormatConfig_NrofSlots_N2 = 0
	PUCCH_FormatConfig_NrofSlots_N4 = 1
	PUCCH_FormatConfig_NrofSlots_N8 = 2
)

var pUCCHFormatConfigNrofSlotsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_FormatConfig_NrofSlots_N2, PUCCH_FormatConfig_NrofSlots_N4, PUCCH_FormatConfig_NrofSlots_N8},
}

const (
	PUCCH_FormatConfig_Pi2BPSK_Enabled = 0
)

var pUCCHFormatConfigPi2BPSKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_FormatConfig_Pi2BPSK_Enabled},
}

const (
	PUCCH_FormatConfig_SimultaneousHARQ_ACK_CSI_True = 0
)

var pUCCHFormatConfigSimultaneousHARQACKCSIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_FormatConfig_SimultaneousHARQ_ACK_CSI_True},
}

type PUCCH_FormatConfig struct {
	InterslotFrequencyHopping *int64
	AdditionalDMRS            *int64
	MaxCodeRate               *PUCCH_MaxCodeRate
	NrofSlots                 *int64
	Pi2BPSK                   *int64
	SimultaneousHARQ_ACK_CSI  *int64
}

func (ie *PUCCH_FormatConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormatConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterslotFrequencyHopping != nil, ie.AdditionalDMRS != nil, ie.MaxCodeRate != nil, ie.NrofSlots != nil, ie.Pi2BPSK != nil, ie.SimultaneousHARQ_ACK_CSI != nil}); err != nil {
		return err
	}

	// 2. interslotFrequencyHopping: enumerated
	{
		if ie.InterslotFrequencyHopping != nil {
			if err := e.EncodeEnumerated(*ie.InterslotFrequencyHopping, pUCCHFormatConfigInterslotFrequencyHoppingConstraints); err != nil {
				return err
			}
		}
	}

	// 3. additionalDMRS: enumerated
	{
		if ie.AdditionalDMRS != nil {
			if err := e.EncodeEnumerated(*ie.AdditionalDMRS, pUCCHFormatConfigAdditionalDMRSConstraints); err != nil {
				return err
			}
		}
	}

	// 4. maxCodeRate: ref
	{
		if ie.MaxCodeRate != nil {
			if err := ie.MaxCodeRate.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. nrofSlots: enumerated
	{
		if ie.NrofSlots != nil {
			if err := e.EncodeEnumerated(*ie.NrofSlots, pUCCHFormatConfigNrofSlotsConstraints); err != nil {
				return err
			}
		}
	}

	// 6. pi2BPSK: enumerated
	{
		if ie.Pi2BPSK != nil {
			if err := e.EncodeEnumerated(*ie.Pi2BPSK, pUCCHFormatConfigPi2BPSKConstraints); err != nil {
				return err
			}
		}
	}

	// 7. simultaneousHARQ-ACK-CSI: enumerated
	{
		if ie.SimultaneousHARQ_ACK_CSI != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousHARQ_ACK_CSI, pUCCHFormatConfigSimultaneousHARQACKCSIConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_FormatConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormatConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. interslotFrequencyHopping: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUCCHFormatConfigInterslotFrequencyHoppingConstraints)
			if err != nil {
				return err
			}
			ie.InterslotFrequencyHopping = &idx
		}
	}

	// 3. additionalDMRS: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pUCCHFormatConfigAdditionalDMRSConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalDMRS = &idx
		}
	}

	// 4. maxCodeRate: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MaxCodeRate = new(PUCCH_MaxCodeRate)
			if err := ie.MaxCodeRate.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nrofSlots: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(pUCCHFormatConfigNrofSlotsConstraints)
			if err != nil {
				return err
			}
			ie.NrofSlots = &idx
		}
	}

	// 6. pi2BPSK: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(pUCCHFormatConfigPi2BPSKConstraints)
			if err != nil {
				return err
			}
			ie.Pi2BPSK = &idx
		}
	}

	// 7. simultaneousHARQ-ACK-CSI: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(pUCCHFormatConfigSimultaneousHARQACKCSIConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousHARQ_ACK_CSI = &idx
		}
	}

	return nil
}
