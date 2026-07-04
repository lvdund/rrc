// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSCCH-ConfigDedicatedSL-PRS-RP-r18 (line 27649).

var sLPSCCHConfigDedicatedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TimeResourcePSCCH-DedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-FreqResourcePSCCH-DedicatedSL-PRS-RP-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18_N2 = 0
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18_N3 = 1
)

var sLPSCCHConfigDedicatedSLPRSRPR18SlTimeResourcePSCCHDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18_N2, SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18_N3},
}

const (
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N10 = 0
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N12 = 1
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N15 = 2
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N20 = 3
	SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N25 = 4
)

var sLPSCCHConfigDedicatedSLPRSRPR18SlFreqResourcePSCCHDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N10, SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N12, SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N15, SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N20, SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18_Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18_N25},
}

var sLPSCCHConfigDedicatedSLPRSRPR18SlDMRSScrambleIDDedicatedSLPRSRPR18Constraints = per.Constrained(0, 65535)

type SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18 struct {
	Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18 *int64
	Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18 *int64
	Sl_DMRS_ScrambleID_DedicatedSL_PRS_RP_r18   *int64
}

func (ie *SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSCCHConfigDedicatedSLPRSRPR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_DMRS_ScrambleID_DedicatedSL_PRS_RP_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18 != nil, ie.Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-TimeResourcePSCCH-DedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18, sLPSCCHConfigDedicatedSLPRSRPR18SlTimeResourcePSCCHDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-FreqResourcePSCCH-DedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18, sLPSCCHConfigDedicatedSLPRSRPR18SlFreqResourcePSCCHDedicatedSLPRSRPR18Constraints); err != nil {
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
					{Name: "sl-DMRS-ScrambleID-DedicatedSL-PRS-RP-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DMRS_ScrambleID_DedicatedSL_PRS_RP_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_DMRS_ScrambleID_DedicatedSL_PRS_RP_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_DMRS_ScrambleID_DedicatedSL_PRS_RP_r18, sLPSCCHConfigDedicatedSLPRSRPR18SlDMRSScrambleIDDedicatedSLPRSRPR18Constraints); err != nil {
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

func (ie *SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSCCHConfigDedicatedSLPRSRPR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-TimeResourcePSCCH-DedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLPSCCHConfigDedicatedSLPRSRPR18SlTimeResourcePSCCHDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeResourcePSCCH_DedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 4. sl-FreqResourcePSCCH-DedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLPSCCHConfigDedicatedSLPRSRPR18SlFreqResourcePSCCHDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_FreqResourcePSCCH_DedicatedSL_PRS_RP_r18 = &idx
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
				{Name: "sl-DMRS-ScrambleID-DedicatedSL-PRS-RP-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sLPSCCHConfigDedicatedSLPRSRPR18SlDMRSScrambleIDDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DMRS_ScrambleID_DedicatedSL_PRS_RP_r18 = &v
		}
	}

	return nil
}
