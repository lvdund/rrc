// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSCCH-Config-r16 (line 28027).

var sLPSCCHConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TimeResourcePSCCH-r16", Optional: true},
		{Name: "sl-FreqResourcePSCCH-r16", Optional: true},
		{Name: "sl-DMRS-ScrambleID-r16", Optional: true},
		{Name: "sl-NumReservedBits-r16", Optional: true},
	},
}

const (
	SL_PSCCH_Config_r16_Sl_TimeResourcePSCCH_r16_N2 = 0
	SL_PSCCH_Config_r16_Sl_TimeResourcePSCCH_r16_N3 = 1
)

var sLPSCCHConfigR16SlTimeResourcePSCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSCCH_Config_r16_Sl_TimeResourcePSCCH_r16_N2, SL_PSCCH_Config_r16_Sl_TimeResourcePSCCH_r16_N3},
}

const (
	SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N10 = 0
	SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N12 = 1
	SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N15 = 2
	SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N20 = 3
	SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N25 = 4
)

var sLPSCCHConfigR16SlFreqResourcePSCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N10, SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N12, SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N15, SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N20, SL_PSCCH_Config_r16_Sl_FreqResourcePSCCH_r16_N25},
}

var sLPSCCHConfigR16SlDMRSScrambleIDR16Constraints = per.Constrained(0, 65535)

var sLPSCCHConfigR16SlNumReservedBitsR16Constraints = per.Constrained(2, 4)

type SL_PSCCH_Config_r16 struct {
	Sl_TimeResourcePSCCH_r16 *int64
	Sl_FreqResourcePSCCH_r16 *int64
	Sl_DMRS_ScrambleID_r16   *int64
	Sl_NumReservedBits_r16   *int64
}

func (ie *SL_PSCCH_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSCCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TimeResourcePSCCH_r16 != nil, ie.Sl_FreqResourcePSCCH_r16 != nil, ie.Sl_DMRS_ScrambleID_r16 != nil, ie.Sl_NumReservedBits_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-TimeResourcePSCCH-r16: enumerated
	{
		if ie.Sl_TimeResourcePSCCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TimeResourcePSCCH_r16, sLPSCCHConfigR16SlTimeResourcePSCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-FreqResourcePSCCH-r16: enumerated
	{
		if ie.Sl_FreqResourcePSCCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_FreqResourcePSCCH_r16, sLPSCCHConfigR16SlFreqResourcePSCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-DMRS-ScrambleID-r16: integer
	{
		if ie.Sl_DMRS_ScrambleID_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_DMRS_ScrambleID_r16, sLPSCCHConfigR16SlDMRSScrambleIDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-NumReservedBits-r16: integer
	{
		if ie.Sl_NumReservedBits_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumReservedBits_r16, sLPSCCHConfigR16SlNumReservedBitsR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PSCCH_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSCCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-TimeResourcePSCCH-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLPSCCHConfigR16SlTimeResourcePSCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeResourcePSCCH_r16 = &idx
		}
	}

	// 4. sl-FreqResourcePSCCH-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLPSCCHConfigR16SlFreqResourcePSCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_FreqResourcePSCCH_r16 = &idx
		}
	}

	// 5. sl-DMRS-ScrambleID-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLPSCCHConfigR16SlDMRSScrambleIDR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DMRS_ScrambleID_r16 = &v
		}
	}

	// 6. sl-NumReservedBits-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLPSCCHConfigR16SlNumReservedBitsR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumReservedBits_r16 = &v
		}
	}

	return nil
}
