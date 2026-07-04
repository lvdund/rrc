// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RACH-ConfigGeneric (line 12993).

var rACHConfigGenericConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "prach-ConfigurationIndex"},
		{Name: "msg1-FDM"},
		{Name: "msg1-FrequencyStart"},
		{Name: "zeroCorrelationZoneConfig"},
		{Name: "preambleReceivedTargetPower"},
		{Name: "preambleTransMax"},
		{Name: "powerRampingStep"},
		{Name: "ra-ResponseWindow"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var rACHConfigGenericPrachConfigurationIndexConstraints = per.Constrained(0, 255)

const (
	RACH_ConfigGeneric_Msg1_FDM_One   = 0
	RACH_ConfigGeneric_Msg1_FDM_Two   = 1
	RACH_ConfigGeneric_Msg1_FDM_Four  = 2
	RACH_ConfigGeneric_Msg1_FDM_Eight = 3
)

var rACHConfigGenericMsg1FDMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_Msg1_FDM_One, RACH_ConfigGeneric_Msg1_FDM_Two, RACH_ConfigGeneric_Msg1_FDM_Four, RACH_ConfigGeneric_Msg1_FDM_Eight},
}

var rACHConfigGenericMsg1FrequencyStartConstraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var rACHConfigGenericZeroCorrelationZoneConfigConstraints = per.Constrained(0, 15)

var rACHConfigGenericPreambleReceivedTargetPowerConstraints = per.Constrained(-202, -60)

const (
	RACH_ConfigGeneric_PreambleTransMax_N3   = 0
	RACH_ConfigGeneric_PreambleTransMax_N4   = 1
	RACH_ConfigGeneric_PreambleTransMax_N5   = 2
	RACH_ConfigGeneric_PreambleTransMax_N6   = 3
	RACH_ConfigGeneric_PreambleTransMax_N7   = 4
	RACH_ConfigGeneric_PreambleTransMax_N8   = 5
	RACH_ConfigGeneric_PreambleTransMax_N10  = 6
	RACH_ConfigGeneric_PreambleTransMax_N20  = 7
	RACH_ConfigGeneric_PreambleTransMax_N50  = 8
	RACH_ConfigGeneric_PreambleTransMax_N100 = 9
	RACH_ConfigGeneric_PreambleTransMax_N200 = 10
)

var rACHConfigGenericPreambleTransMaxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_PreambleTransMax_N3, RACH_ConfigGeneric_PreambleTransMax_N4, RACH_ConfigGeneric_PreambleTransMax_N5, RACH_ConfigGeneric_PreambleTransMax_N6, RACH_ConfigGeneric_PreambleTransMax_N7, RACH_ConfigGeneric_PreambleTransMax_N8, RACH_ConfigGeneric_PreambleTransMax_N10, RACH_ConfigGeneric_PreambleTransMax_N20, RACH_ConfigGeneric_PreambleTransMax_N50, RACH_ConfigGeneric_PreambleTransMax_N100, RACH_ConfigGeneric_PreambleTransMax_N200},
}

const (
	RACH_ConfigGeneric_PowerRampingStep_DB0 = 0
	RACH_ConfigGeneric_PowerRampingStep_DB2 = 1
	RACH_ConfigGeneric_PowerRampingStep_DB4 = 2
	RACH_ConfigGeneric_PowerRampingStep_DB6 = 3
)

var rACHConfigGenericPowerRampingStepConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_PowerRampingStep_DB0, RACH_ConfigGeneric_PowerRampingStep_DB2, RACH_ConfigGeneric_PowerRampingStep_DB4, RACH_ConfigGeneric_PowerRampingStep_DB6},
}

const (
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl1  = 0
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl2  = 1
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl4  = 2
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl8  = 3
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl10 = 4
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl20 = 5
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl40 = 6
	RACH_ConfigGeneric_Ra_ResponseWindow_Sl80 = 7
)

var rACHConfigGenericRaResponseWindowConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_Ra_ResponseWindow_Sl1, RACH_ConfigGeneric_Ra_ResponseWindow_Sl2, RACH_ConfigGeneric_Ra_ResponseWindow_Sl4, RACH_ConfigGeneric_Ra_ResponseWindow_Sl8, RACH_ConfigGeneric_Ra_ResponseWindow_Sl10, RACH_ConfigGeneric_Ra_ResponseWindow_Sl20, RACH_ConfigGeneric_Ra_ResponseWindow_Sl40, RACH_ConfigGeneric_Ra_ResponseWindow_Sl80},
}

const (
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf1  = 0
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf2  = 1
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf4  = 2
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf8  = 3
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf16 = 4
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf32 = 5
	RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf64 = 6
)

var rACHConfigGenericExtPrachConfigurationPeriodScalingIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf1, RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf2, RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf4, RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf8, RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf16, RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf32, RACH_ConfigGeneric_Ext_Prach_ConfigurationPeriodScaling_IAB_r16_Scf64},
}

var rACHConfigGenericPrachConfigurationFrameOffsetIABR16Constraints = per.Constrained(0, 63)

var rACHConfigGenericPrachConfigurationSOffsetIABR16Constraints = per.Constrained(0, 39)

const (
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1610_Sl60  = 0
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1610_Sl160 = 1
)

var rACHConfigGenericExtRaResponseWindowV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1610_Sl60, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1610_Sl160},
}

var rACHConfigGenericPrachConfigurationIndexV1610Constraints = per.Constrained(256, 262)

const (
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl240  = 0
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl320  = 1
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl640  = 2
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl960  = 3
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl1280 = 4
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl1920 = 5
	RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl2560 = 6
)

var rACHConfigGenericExtRaResponseWindowV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl240, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl320, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl640, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl960, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl1280, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl1920, RACH_ConfigGeneric_Ext_Ra_ResponseWindow_v1700_Sl2560},
}

var rACHConfigGenericSbfdRACHSingleConfigPreambleReceivedTargetPowerR19Constraints = per.Constrained(-202, -60)

type RACH_ConfigGeneric struct {
	Prach_ConfigurationIndex                               int64
	Msg1_FDM                                               int64
	Msg1_FrequencyStart                                    int64
	ZeroCorrelationZoneConfig                              int64
	PreambleReceivedTargetPower                            int64
	PreambleTransMax                                       int64
	PowerRampingStep                                       int64
	Ra_ResponseWindow                                      int64
	Prach_ConfigurationPeriodScaling_IAB_r16               *int64
	Prach_ConfigurationFrameOffset_IAB_r16                 *int64
	Prach_ConfigurationSOffset_IAB_r16                     *int64
	Ra_ResponseWindow_v1610                                *int64
	Prach_ConfigurationIndex_v1610                         *int64
	Ra_ResponseWindow_v1700                                *int64
	Sbfd_RACH_SingleConfig_PreambleReceivedTargetPower_r19 *int64
}

func (ie *RACH_ConfigGeneric) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHConfigGenericConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil || ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil || ie.Prach_ConfigurationSOffset_IAB_r16 != nil || ie.Ra_ResponseWindow_v1610 != nil || ie.Prach_ConfigurationIndex_v1610 != nil
	hasExtGroup1 := ie.Ra_ResponseWindow_v1700 != nil
	hasExtGroup2 := ie.Sbfd_RACH_SingleConfig_PreambleReceivedTargetPower_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. prach-ConfigurationIndex: integer
	{
		if err := e.EncodeInteger(ie.Prach_ConfigurationIndex, rACHConfigGenericPrachConfigurationIndexConstraints); err != nil {
			return err
		}
	}

	// 3. msg1-FDM: enumerated
	{
		if err := e.EncodeEnumerated(ie.Msg1_FDM, rACHConfigGenericMsg1FDMConstraints); err != nil {
			return err
		}
	}

	// 4. msg1-FrequencyStart: integer
	{
		if err := e.EncodeInteger(ie.Msg1_FrequencyStart, rACHConfigGenericMsg1FrequencyStartConstraints); err != nil {
			return err
		}
	}

	// 5. zeroCorrelationZoneConfig: integer
	{
		if err := e.EncodeInteger(ie.ZeroCorrelationZoneConfig, rACHConfigGenericZeroCorrelationZoneConfigConstraints); err != nil {
			return err
		}
	}

	// 6. preambleReceivedTargetPower: integer
	{
		if err := e.EncodeInteger(ie.PreambleReceivedTargetPower, rACHConfigGenericPreambleReceivedTargetPowerConstraints); err != nil {
			return err
		}
	}

	// 7. preambleTransMax: enumerated
	{
		if err := e.EncodeEnumerated(ie.PreambleTransMax, rACHConfigGenericPreambleTransMaxConstraints); err != nil {
			return err
		}
	}

	// 8. powerRampingStep: enumerated
	{
		if err := e.EncodeEnumerated(ie.PowerRampingStep, rACHConfigGenericPowerRampingStepConstraints); err != nil {
			return err
		}
	}

	// 9. ra-ResponseWindow: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ra_ResponseWindow, rACHConfigGenericRaResponseWindowConstraints); err != nil {
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
					{Name: "prach-ConfigurationPeriodScaling-IAB-r16", Optional: true},
					{Name: "prach-ConfigurationFrameOffset-IAB-r16", Optional: true},
					{Name: "prach-ConfigurationSOffset-IAB-r16", Optional: true},
					{Name: "ra-ResponseWindow-v1610", Optional: true},
					{Name: "prach-ConfigurationIndex-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil, ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil, ie.Prach_ConfigurationSOffset_IAB_r16 != nil, ie.Ra_ResponseWindow_v1610 != nil, ie.Prach_ConfigurationIndex_v1610 != nil}); err != nil {
				return err
			}

			if ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Prach_ConfigurationPeriodScaling_IAB_r16, rACHConfigGenericExtPrachConfigurationPeriodScalingIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil {
				if err := ex.EncodeInteger(*ie.Prach_ConfigurationFrameOffset_IAB_r16, rACHConfigGenericPrachConfigurationFrameOffsetIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Prach_ConfigurationSOffset_IAB_r16 != nil {
				if err := ex.EncodeInteger(*ie.Prach_ConfigurationSOffset_IAB_r16, rACHConfigGenericPrachConfigurationSOffsetIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ra_ResponseWindow_v1610 != nil {
				if err := ex.EncodeEnumerated(*ie.Ra_ResponseWindow_v1610, rACHConfigGenericExtRaResponseWindowV1610Constraints); err != nil {
					return err
				}
			}

			if ie.Prach_ConfigurationIndex_v1610 != nil {
				if err := ex.EncodeInteger(*ie.Prach_ConfigurationIndex_v1610, rACHConfigGenericPrachConfigurationIndexV1610Constraints); err != nil {
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
					{Name: "ra-ResponseWindow-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_ResponseWindow_v1700 != nil}); err != nil {
				return err
			}

			if ie.Ra_ResponseWindow_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.Ra_ResponseWindow_v1700, rACHConfigGenericExtRaResponseWindowV1700Constraints); err != nil {
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
					{Name: "sbfd-RACH-SingleConfig-preambleReceivedTargetPower-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_RACH_SingleConfig_PreambleReceivedTargetPower_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_RACH_SingleConfig_PreambleReceivedTargetPower_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sbfd_RACH_SingleConfig_PreambleReceivedTargetPower_r19, rACHConfigGenericSbfdRACHSingleConfigPreambleReceivedTargetPowerR19Constraints); err != nil {
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

func (ie *RACH_ConfigGeneric) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHConfigGenericConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. prach-ConfigurationIndex: integer
	{
		v0, err := d.DecodeInteger(rACHConfigGenericPrachConfigurationIndexConstraints)
		if err != nil {
			return err
		}
		ie.Prach_ConfigurationIndex = v0
	}

	// 3. msg1-FDM: enumerated
	{
		v1, err := d.DecodeEnumerated(rACHConfigGenericMsg1FDMConstraints)
		if err != nil {
			return err
		}
		ie.Msg1_FDM = v1
	}

	// 4. msg1-FrequencyStart: integer
	{
		v2, err := d.DecodeInteger(rACHConfigGenericMsg1FrequencyStartConstraints)
		if err != nil {
			return err
		}
		ie.Msg1_FrequencyStart = v2
	}

	// 5. zeroCorrelationZoneConfig: integer
	{
		v3, err := d.DecodeInteger(rACHConfigGenericZeroCorrelationZoneConfigConstraints)
		if err != nil {
			return err
		}
		ie.ZeroCorrelationZoneConfig = v3
	}

	// 6. preambleReceivedTargetPower: integer
	{
		v4, err := d.DecodeInteger(rACHConfigGenericPreambleReceivedTargetPowerConstraints)
		if err != nil {
			return err
		}
		ie.PreambleReceivedTargetPower = v4
	}

	// 7. preambleTransMax: enumerated
	{
		v5, err := d.DecodeEnumerated(rACHConfigGenericPreambleTransMaxConstraints)
		if err != nil {
			return err
		}
		ie.PreambleTransMax = v5
	}

	// 8. powerRampingStep: enumerated
	{
		v6, err := d.DecodeEnumerated(rACHConfigGenericPowerRampingStepConstraints)
		if err != nil {
			return err
		}
		ie.PowerRampingStep = v6
	}

	// 9. ra-ResponseWindow: enumerated
	{
		v7, err := d.DecodeEnumerated(rACHConfigGenericRaResponseWindowConstraints)
		if err != nil {
			return err
		}
		ie.Ra_ResponseWindow = v7
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
				{Name: "prach-ConfigurationPeriodScaling-IAB-r16", Optional: true},
				{Name: "prach-ConfigurationFrameOffset-IAB-r16", Optional: true},
				{Name: "prach-ConfigurationSOffset-IAB-r16", Optional: true},
				{Name: "ra-ResponseWindow-v1610", Optional: true},
				{Name: "prach-ConfigurationIndex-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rACHConfigGenericExtPrachConfigurationPeriodScalingIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Prach_ConfigurationPeriodScaling_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(rACHConfigGenericPrachConfigurationFrameOffsetIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Prach_ConfigurationFrameOffset_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(rACHConfigGenericPrachConfigurationSOffsetIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Prach_ConfigurationSOffset_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(rACHConfigGenericExtRaResponseWindowV1610Constraints)
			if err != nil {
				return err
			}
			ie.Ra_ResponseWindow_v1610 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(rACHConfigGenericPrachConfigurationIndexV1610Constraints)
			if err != nil {
				return err
			}
			ie.Prach_ConfigurationIndex_v1610 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-ResponseWindow-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rACHConfigGenericExtRaResponseWindowV1700Constraints)
			if err != nil {
				return err
			}
			ie.Ra_ResponseWindow_v1700 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sbfd-RACH-SingleConfig-preambleReceivedTargetPower-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(rACHConfigGenericSbfdRACHSingleConfigPreambleReceivedTargetPowerR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_RACH_SingleConfig_PreambleReceivedTargetPower_r19 = &v
		}
	}

	return nil
}
