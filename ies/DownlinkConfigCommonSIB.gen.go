// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DownlinkConfigCommonSIB (line 7830).

var downlinkConfigCommonSIBConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoDL"},
		{Name: "initialDownlinkBWP"},
		{Name: "bcch-Config"},
		{Name: "pcch-Config"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

type DownlinkConfigCommonSIB struct {
	FrequencyInfoDL               FrequencyInfoDL_SIB
	InitialDownlinkBWP            BWP_DownlinkCommon
	Bcch_Config                   BCCH_Config
	Pcch_Config                   PCCH_Config
	Pei_Config_r17                *PEI_Config_r17
	InitialDownlinkBWP_RedCap_r17 *BWP_DownlinkCommon
	FrequencyInfoDL_v1800         *FrequencyInfoDL_SIB_v1800
	LowPowerConfig_r19            *LowPowerConfig_r19
	PagingAdaptPEI_Config_r19     *PEI_Config_r19
}

func (ie *DownlinkConfigCommonSIB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(downlinkConfigCommonSIBConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Pei_Config_r17 != nil || ie.InitialDownlinkBWP_RedCap_r17 != nil
	hasExtGroup1 := ie.FrequencyInfoDL_v1800 != nil
	hasExtGroup2 := ie.LowPowerConfig_r19 != nil || ie.PagingAdaptPEI_Config_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. frequencyInfoDL: ref
	{
		if err := ie.FrequencyInfoDL.Encode(e); err != nil {
			return err
		}
	}

	// 3. initialDownlinkBWP: ref
	{
		if err := ie.InitialDownlinkBWP.Encode(e); err != nil {
			return err
		}
	}

	// 4. bcch-Config: ref
	{
		if err := ie.Bcch_Config.Encode(e); err != nil {
			return err
		}
	}

	// 5. pcch-Config: ref
	{
		if err := ie.Pcch_Config.Encode(e); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "pei-Config-r17", Optional: true},
					{Name: "initialDownlinkBWP-RedCap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pei_Config_r17 != nil, ie.InitialDownlinkBWP_RedCap_r17 != nil}); err != nil {
				return err
			}

			if ie.Pei_Config_r17 != nil {
				if err := ie.Pei_Config_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.InitialDownlinkBWP_RedCap_r17 != nil {
				if err := ie.InitialDownlinkBWP_RedCap_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "frequencyInfoDL-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FrequencyInfoDL_v1800 != nil}); err != nil {
				return err
			}

			if ie.FrequencyInfoDL_v1800 != nil {
				if err := ie.FrequencyInfoDL_v1800.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "lowPowerConfig-r19", Optional: true},
					{Name: "pagingAdaptPEI-Config-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LowPowerConfig_r19 != nil, ie.PagingAdaptPEI_Config_r19 != nil}); err != nil {
				return err
			}

			if ie.LowPowerConfig_r19 != nil {
				if err := ie.LowPowerConfig_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PagingAdaptPEI_Config_r19 != nil {
				if err := ie.PagingAdaptPEI_Config_r19.Encode(ex); err != nil {
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

func (ie *DownlinkConfigCommonSIB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(downlinkConfigCommonSIBConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. frequencyInfoDL: ref
	{
		if err := ie.FrequencyInfoDL.Decode(d); err != nil {
			return err
		}
	}

	// 3. initialDownlinkBWP: ref
	{
		if err := ie.InitialDownlinkBWP.Decode(d); err != nil {
			return err
		}
	}

	// 4. bcch-Config: ref
	{
		if err := ie.Bcch_Config.Decode(d); err != nil {
			return err
		}
	}

	// 5. pcch-Config: ref
	{
		if err := ie.Pcch_Config.Decode(d); err != nil {
			return err
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
				{Name: "pei-Config-r17", Optional: true},
				{Name: "initialDownlinkBWP-RedCap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pei_Config_r17 = new(PEI_Config_r17)
			if err := ie.Pei_Config_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.InitialDownlinkBWP_RedCap_r17 = new(BWP_DownlinkCommon)
			if err := ie.InitialDownlinkBWP_RedCap_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "frequencyInfoDL-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FrequencyInfoDL_v1800 = new(FrequencyInfoDL_SIB_v1800)
			if err := ie.FrequencyInfoDL_v1800.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "lowPowerConfig-r19", Optional: true},
				{Name: "pagingAdaptPEI-Config-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.LowPowerConfig_r19 = new(LowPowerConfig_r19)
			if err := ie.LowPowerConfig_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.PagingAdaptPEI_Config_r19 = new(PEI_Config_r19)
			if err := ie.PagingAdaptPEI_Config_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
