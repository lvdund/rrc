// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-BearerConfig (line 14010).

var rLCBearerConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "logicalChannelIdentity"},
		{Name: "servedRadioBearer", Optional: true},
		{Name: "reestablishRLC", Optional: true},
		{Name: "rlc-Config", Optional: true},
		{Name: "mac-LogicalChannelConfig", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var rLC_BearerConfigServedRadioBearerConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "srb-Identity"},
		{Name: "drb-Identity"},
	},
}

const (
	RLC_BearerConfig_ServedRadioBearer_Srb_Identity = 0
	RLC_BearerConfig_ServedRadioBearer_Drb_Identity = 1
)

const (
	RLC_BearerConfig_ReestablishRLC_True = 0
)

var rLCBearerConfigReestablishRLCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_BearerConfig_ReestablishRLC_True},
}

type RLC_BearerConfig struct {
	LogicalChannelIdentity LogicalChannelIdentity
	ServedRadioBearer      *struct {
		Choice       int
		Srb_Identity *SRB_Identity
		Drb_Identity *DRB_Identity
	}
	ReestablishRLC                *int64
	Rlc_Config                    *RLC_Config
	Mac_LogicalChannelConfig      *LogicalChannelConfig
	Rlc_Config_v1610              *RLC_Config_v1610
	Rlc_Config_v1700              *RLC_Config_v1700
	LogicalChannelIdentityExt_r17 *LogicalChannelIdentityExt_r17
	MulticastRLC_BearerConfig_r17 *MulticastRLC_BearerConfig_r17
	ServedRadioBearerSRB4_r17     *SRB_Identity_v1700
	ServedRadioBearerSRB5_r18     *SRB_Identity_v1800
	Rlc_Config_v1900              *RLC_Config_v1900
	ServedRadioBearerSRB6_r19     *SRB_Identity_v1900
}

func (ie *RLC_BearerConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCBearerConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Rlc_Config_v1610 != nil
	hasExtGroup1 := ie.Rlc_Config_v1700 != nil || ie.LogicalChannelIdentityExt_r17 != nil || ie.MulticastRLC_BearerConfig_r17 != nil || ie.ServedRadioBearerSRB4_r17 != nil
	hasExtGroup2 := ie.ServedRadioBearerSRB5_r18 != nil
	hasExtGroup3 := ie.Rlc_Config_v1900 != nil || ie.ServedRadioBearerSRB6_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServedRadioBearer != nil, ie.ReestablishRLC != nil, ie.Rlc_Config != nil, ie.Mac_LogicalChannelConfig != nil}); err != nil {
		return err
	}

	// 3. logicalChannelIdentity: ref
	{
		if err := ie.LogicalChannelIdentity.Encode(e); err != nil {
			return err
		}
	}

	// 4. servedRadioBearer: choice
	{
		if ie.ServedRadioBearer != nil {
			choiceEnc := e.NewChoiceEncoder(rLC_BearerConfigServedRadioBearerConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ServedRadioBearer).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ServedRadioBearer).Choice {
			case RLC_BearerConfig_ServedRadioBearer_Srb_Identity:
				if err := (*ie.ServedRadioBearer).Srb_Identity.Encode(e); err != nil {
					return err
				}
			case RLC_BearerConfig_ServedRadioBearer_Drb_Identity:
				if err := (*ie.ServedRadioBearer).Drb_Identity.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ServedRadioBearer).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. reestablishRLC: enumerated
	{
		if ie.ReestablishRLC != nil {
			if err := e.EncodeEnumerated(*ie.ReestablishRLC, rLCBearerConfigReestablishRLCConstraints); err != nil {
				return err
			}
		}
	}

	// 6. rlc-Config: ref
	{
		if ie.Rlc_Config != nil {
			if err := ie.Rlc_Config.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. mac-LogicalChannelConfig: ref
	{
		if ie.Mac_LogicalChannelConfig != nil {
			if err := ie.Mac_LogicalChannelConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rlc-Config-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rlc_Config_v1610 != nil}); err != nil {
				return err
			}

			if ie.Rlc_Config_v1610 != nil {
				if err := ie.Rlc_Config_v1610.Encode(ex); err != nil {
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
					{Name: "rlc-Config-v1700", Optional: true},
					{Name: "logicalChannelIdentityExt-r17", Optional: true},
					{Name: "multicastRLC-BearerConfig-r17", Optional: true},
					{Name: "servedRadioBearerSRB4-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rlc_Config_v1700 != nil, ie.LogicalChannelIdentityExt_r17 != nil, ie.MulticastRLC_BearerConfig_r17 != nil, ie.ServedRadioBearerSRB4_r17 != nil}); err != nil {
				return err
			}

			if ie.Rlc_Config_v1700 != nil {
				if err := ie.Rlc_Config_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.LogicalChannelIdentityExt_r17 != nil {
				if err := ie.LogicalChannelIdentityExt_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MulticastRLC_BearerConfig_r17 != nil {
				if err := ie.MulticastRLC_BearerConfig_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ServedRadioBearerSRB4_r17 != nil {
				if err := ie.ServedRadioBearerSRB4_r17.Encode(ex); err != nil {
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
					{Name: "servedRadioBearerSRB5-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ServedRadioBearerSRB5_r18 != nil}); err != nil {
				return err
			}

			if ie.ServedRadioBearerSRB5_r18 != nil {
				if err := ie.ServedRadioBearerSRB5_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rlc-Config-v1900", Optional: true},
					{Name: "servedRadioBearerSRB6-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rlc_Config_v1900 != nil, ie.ServedRadioBearerSRB6_r19 != nil}); err != nil {
				return err
			}

			if ie.Rlc_Config_v1900 != nil {
				if err := ie.Rlc_Config_v1900.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ServedRadioBearerSRB6_r19 != nil {
				if err := ie.ServedRadioBearerSRB6_r19.Encode(ex); err != nil {
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

func (ie *RLC_BearerConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCBearerConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. logicalChannelIdentity: ref
	{
		if err := ie.LogicalChannelIdentity.Decode(d); err != nil {
			return err
		}
	}

	// 4. servedRadioBearer: choice
	{
		if seq.IsComponentPresent(1) {
			ie.ServedRadioBearer = &struct {
				Choice       int
				Srb_Identity *SRB_Identity
				Drb_Identity *DRB_Identity
			}{}
			choiceDec := d.NewChoiceDecoder(rLC_BearerConfigServedRadioBearerConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ServedRadioBearer).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RLC_BearerConfig_ServedRadioBearer_Srb_Identity:
				(*ie.ServedRadioBearer).Srb_Identity = new(SRB_Identity)
				if err := (*ie.ServedRadioBearer).Srb_Identity.Decode(d); err != nil {
					return err
				}
			case RLC_BearerConfig_ServedRadioBearer_Drb_Identity:
				(*ie.ServedRadioBearer).Drb_Identity = new(DRB_Identity)
				if err := (*ie.ServedRadioBearer).Drb_Identity.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. reestablishRLC: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rLCBearerConfigReestablishRLCConstraints)
			if err != nil {
				return err
			}
			ie.ReestablishRLC = &idx
		}
	}

	// 6. rlc-Config: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Rlc_Config = new(RLC_Config)
			if err := ie.Rlc_Config.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. mac-LogicalChannelConfig: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mac_LogicalChannelConfig = new(LogicalChannelConfig)
			if err := ie.Mac_LogicalChannelConfig.Decode(d); err != nil {
				return err
			}
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
				{Name: "rlc-Config-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Rlc_Config_v1610 = new(RLC_Config_v1610)
			if err := ie.Rlc_Config_v1610.Decode(dx); err != nil {
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
				{Name: "rlc-Config-v1700", Optional: true},
				{Name: "logicalChannelIdentityExt-r17", Optional: true},
				{Name: "multicastRLC-BearerConfig-r17", Optional: true},
				{Name: "servedRadioBearerSRB4-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Rlc_Config_v1700 = new(RLC_Config_v1700)
			if err := ie.Rlc_Config_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.LogicalChannelIdentityExt_r17 = new(LogicalChannelIdentityExt_r17)
			if err := ie.LogicalChannelIdentityExt_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MulticastRLC_BearerConfig_r17 = new(MulticastRLC_BearerConfig_r17)
			if err := ie.MulticastRLC_BearerConfig_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.ServedRadioBearerSRB4_r17 = new(SRB_Identity_v1700)
			if err := ie.ServedRadioBearerSRB4_r17.Decode(dx); err != nil {
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
				{Name: "servedRadioBearerSRB5-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ServedRadioBearerSRB5_r18 = new(SRB_Identity_v1800)
			if err := ie.ServedRadioBearerSRB5_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "rlc-Config-v1900", Optional: true},
				{Name: "servedRadioBearerSRB6-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Rlc_Config_v1900 = new(RLC_Config_v1900)
			if err := ie.Rlc_Config_v1900.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ServedRadioBearerSRB6_r19 = new(SRB_Identity_v1900)
			if err := ie.ServedRadioBearerSRB6_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
