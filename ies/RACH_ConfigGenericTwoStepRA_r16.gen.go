// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RACH-ConfigGenericTwoStepRA-r16 (line 13021).

var rACHConfigGenericTwoStepRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "msgA-PRACH-ConfigurationIndex-r16", Optional: true},
		{Name: "msgA-RO-FDM-r16", Optional: true},
		{Name: "msgA-RO-FrequencyStart-r16", Optional: true},
		{Name: "msgA-ZeroCorrelationZoneConfig-r16", Optional: true},
		{Name: "msgA-PreamblePowerRampingStep-r16", Optional: true},
		{Name: "msgA-PreambleReceivedTargetPower-r16", Optional: true},
		{Name: "msgB-ResponseWindow-r16", Optional: true},
		{Name: "preambleTransMax-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var rACHConfigGenericTwoStepRAR16MsgAPRACHConfigurationIndexR16Constraints = per.Constrained(0, 262)

const (
	RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_One   = 0
	RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_Two   = 1
	RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_Four  = 2
	RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_Eight = 3
)

var rACHConfigGenericTwoStepRAR16MsgAROFDMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_One, RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_Two, RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_Four, RACH_ConfigGenericTwoStepRA_r16_MsgA_RO_FDM_r16_Eight},
}

var rACHConfigGenericTwoStepRAR16MsgAROFrequencyStartR16Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var rACHConfigGenericTwoStepRAR16MsgAZeroCorrelationZoneConfigR16Constraints = per.Constrained(0, 15)

const (
	RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB0 = 0
	RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB2 = 1
	RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB4 = 2
	RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB6 = 3
)

var rACHConfigGenericTwoStepRAR16MsgAPreamblePowerRampingStepR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB0, RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB2, RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB4, RACH_ConfigGenericTwoStepRA_r16_MsgA_PreamblePowerRampingStep_r16_DB6},
}

var rACHConfigGenericTwoStepRAR16MsgAPreambleReceivedTargetPowerR16Constraints = per.Constrained(-202, -60)

const (
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl1   = 0
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl2   = 1
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl4   = 2
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl8   = 3
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl10  = 4
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl20  = 5
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl40  = 6
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl80  = 7
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl160 = 8
	RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl320 = 9
)

var rACHConfigGenericTwoStepRAR16MsgBResponseWindowR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl1, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl2, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl4, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl8, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl10, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl20, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl40, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl80, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl160, RACH_ConfigGenericTwoStepRA_r16_MsgB_ResponseWindow_r16_Sl320},
}

const (
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N3   = 0
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N4   = 1
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N5   = 2
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N6   = 3
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N7   = 4
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N8   = 5
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N10  = 6
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N20  = 7
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N50  = 8
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N100 = 9
	RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N200 = 10
)

var rACHConfigGenericTwoStepRAR16PreambleTransMaxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N3, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N4, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N5, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N6, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N7, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N8, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N10, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N20, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N50, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N100, RACH_ConfigGenericTwoStepRA_r16_PreambleTransMax_r16_N200},
}

const (
	RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl240  = 0
	RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl640  = 1
	RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl960  = 2
	RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl1280 = 3
	RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl1920 = 4
	RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl2560 = 5
)

var rACHConfigGenericTwoStepRAR16ExtMsgBResponseWindowV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl240, RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl640, RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl960, RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl1280, RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl1920, RACH_ConfigGenericTwoStepRA_r16_Ext_MsgB_ResponseWindow_v1700_Sl2560},
}

type RACH_ConfigGenericTwoStepRA_r16 struct {
	MsgA_PRACH_ConfigurationIndex_r16    *int64
	MsgA_RO_FDM_r16                      *int64
	MsgA_RO_FrequencyStart_r16           *int64
	MsgA_ZeroCorrelationZoneConfig_r16   *int64
	MsgA_PreamblePowerRampingStep_r16    *int64
	MsgA_PreambleReceivedTargetPower_r16 *int64
	MsgB_ResponseWindow_r16              *int64
	PreambleTransMax_r16                 *int64
	MsgB_ResponseWindow_v1700            *int64
}

func (ie *RACH_ConfigGenericTwoStepRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHConfigGenericTwoStepRAR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MsgB_ResponseWindow_v1700 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MsgA_PRACH_ConfigurationIndex_r16 != nil, ie.MsgA_RO_FDM_r16 != nil, ie.MsgA_RO_FrequencyStart_r16 != nil, ie.MsgA_ZeroCorrelationZoneConfig_r16 != nil, ie.MsgA_PreamblePowerRampingStep_r16 != nil, ie.MsgA_PreambleReceivedTargetPower_r16 != nil, ie.MsgB_ResponseWindow_r16 != nil, ie.PreambleTransMax_r16 != nil}); err != nil {
		return err
	}

	// 3. msgA-PRACH-ConfigurationIndex-r16: integer
	{
		if ie.MsgA_PRACH_ConfigurationIndex_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_PRACH_ConfigurationIndex_r16, rACHConfigGenericTwoStepRAR16MsgAPRACHConfigurationIndexR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. msgA-RO-FDM-r16: enumerated
	{
		if ie.MsgA_RO_FDM_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_RO_FDM_r16, rACHConfigGenericTwoStepRAR16MsgAROFDMR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. msgA-RO-FrequencyStart-r16: integer
	{
		if ie.MsgA_RO_FrequencyStart_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_RO_FrequencyStart_r16, rACHConfigGenericTwoStepRAR16MsgAROFrequencyStartR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. msgA-ZeroCorrelationZoneConfig-r16: integer
	{
		if ie.MsgA_ZeroCorrelationZoneConfig_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_ZeroCorrelationZoneConfig_r16, rACHConfigGenericTwoStepRAR16MsgAZeroCorrelationZoneConfigR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. msgA-PreamblePowerRampingStep-r16: enumerated
	{
		if ie.MsgA_PreamblePowerRampingStep_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_PreamblePowerRampingStep_r16, rACHConfigGenericTwoStepRAR16MsgAPreamblePowerRampingStepR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. msgA-PreambleReceivedTargetPower-r16: integer
	{
		if ie.MsgA_PreambleReceivedTargetPower_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_PreambleReceivedTargetPower_r16, rACHConfigGenericTwoStepRAR16MsgAPreambleReceivedTargetPowerR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. msgB-ResponseWindow-r16: enumerated
	{
		if ie.MsgB_ResponseWindow_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgB_ResponseWindow_r16, rACHConfigGenericTwoStepRAR16MsgBResponseWindowR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. preambleTransMax-r16: enumerated
	{
		if ie.PreambleTransMax_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PreambleTransMax_r16, rACHConfigGenericTwoStepRAR16PreambleTransMaxR16Constraints); err != nil {
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
					{Name: "msgB-ResponseWindow-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MsgB_ResponseWindow_v1700 != nil}); err != nil {
				return err
			}

			if ie.MsgB_ResponseWindow_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.MsgB_ResponseWindow_v1700, rACHConfigGenericTwoStepRAR16ExtMsgBResponseWindowV1700Constraints); err != nil {
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

func (ie *RACH_ConfigGenericTwoStepRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHConfigGenericTwoStepRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. msgA-PRACH-ConfigurationIndex-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(rACHConfigGenericTwoStepRAR16MsgAPRACHConfigurationIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PRACH_ConfigurationIndex_r16 = &v
		}
	}

	// 4. msgA-RO-FDM-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rACHConfigGenericTwoStepRAR16MsgAROFDMR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RO_FDM_r16 = &idx
		}
	}

	// 5. msgA-RO-FrequencyStart-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(rACHConfigGenericTwoStepRAR16MsgAROFrequencyStartR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RO_FrequencyStart_r16 = &v
		}
	}

	// 6. msgA-ZeroCorrelationZoneConfig-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(rACHConfigGenericTwoStepRAR16MsgAZeroCorrelationZoneConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_ZeroCorrelationZoneConfig_r16 = &v
		}
	}

	// 7. msgA-PreamblePowerRampingStep-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(rACHConfigGenericTwoStepRAR16MsgAPreamblePowerRampingStepR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PreamblePowerRampingStep_r16 = &idx
		}
	}

	// 8. msgA-PreambleReceivedTargetPower-r16: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(rACHConfigGenericTwoStepRAR16MsgAPreambleReceivedTargetPowerR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PreambleReceivedTargetPower_r16 = &v
		}
	}

	// 9. msgB-ResponseWindow-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(rACHConfigGenericTwoStepRAR16MsgBResponseWindowR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgB_ResponseWindow_r16 = &idx
		}
	}

	// 10. preambleTransMax-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(rACHConfigGenericTwoStepRAR16PreambleTransMaxR16Constraints)
			if err != nil {
				return err
			}
			ie.PreambleTransMax_r16 = &idx
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
				{Name: "msgB-ResponseWindow-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rACHConfigGenericTwoStepRAR16ExtMsgBResponseWindowV1700Constraints)
			if err != nil {
				return err
			}
			ie.MsgB_ResponseWindow_v1700 = &v
		}
	}

	return nil
}
