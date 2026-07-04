// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersXDD-Diff (line 21055).

var mACParametersXDDDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "skipUplinkTxDynamic", Optional: true},
		{Name: "logicalChannelSR-DelayTimer", Optional: true},
		{Name: "longDRX-Cycle", Optional: true},
		{Name: "shortDRX-Cycle", Optional: true},
		{Name: "multipleSR-Configurations", Optional: true},
		{Name: "multipleConfiguredGrants", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	MAC_ParametersXDD_Diff_SkipUplinkTxDynamic_Supported = 0
)

var mACParametersXDDDiffSkipUplinkTxDynamicConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_SkipUplinkTxDynamic_Supported},
}

const (
	MAC_ParametersXDD_Diff_LogicalChannelSR_DelayTimer_Supported = 0
)

var mACParametersXDDDiffLogicalChannelSRDelayTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_LogicalChannelSR_DelayTimer_Supported},
}

const (
	MAC_ParametersXDD_Diff_LongDRX_Cycle_Supported = 0
)

var mACParametersXDDDiffLongDRXCycleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_LongDRX_Cycle_Supported},
}

const (
	MAC_ParametersXDD_Diff_ShortDRX_Cycle_Supported = 0
)

var mACParametersXDDDiffShortDRXCycleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_ShortDRX_Cycle_Supported},
}

const (
	MAC_ParametersXDD_Diff_MultipleSR_Configurations_Supported = 0
)

var mACParametersXDDDiffMultipleSRConfigurationsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_MultipleSR_Configurations_Supported},
}

const (
	MAC_ParametersXDD_Diff_MultipleConfiguredGrants_Supported = 0
)

var mACParametersXDDDiffMultipleConfiguredGrantsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_MultipleConfiguredGrants_Supported},
}

const (
	MAC_ParametersXDD_Diff_Ext_SecondaryDRX_Group_r16_Supported = 0
)

var mACParametersXDDDiffExtSecondaryDRXGroupR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_Ext_SecondaryDRX_Group_r16_Supported},
}

const (
	MAC_ParametersXDD_Diff_Ext_EnhancedSkipUplinkTxDynamic_r16_Supported = 0
)

var mACParametersXDDDiffExtEnhancedSkipUplinkTxDynamicR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_Ext_EnhancedSkipUplinkTxDynamic_r16_Supported},
}

const (
	MAC_ParametersXDD_Diff_Ext_EnhancedSkipUplinkTxConfigured_r16_Supported = 0
)

var mACParametersXDDDiffExtEnhancedSkipUplinkTxConfiguredR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_Ext_EnhancedSkipUplinkTxConfigured_r16_Supported},
}

const (
	MAC_ParametersXDD_Diff_Ext_Dummy1_Supported = 0
)

var mACParametersXDDDiffExtDummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_Ext_Dummy1_Supported},
}

const (
	MAC_ParametersXDD_Diff_Ext_Dummy2_Supported = 0
)

var mACParametersXDDDiffExtDummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersXDD_Diff_Ext_Dummy2_Supported},
}

type MAC_ParametersXDD_Diff struct {
	SkipUplinkTxDynamic                *int64
	LogicalChannelSR_DelayTimer        *int64
	LongDRX_Cycle                      *int64
	ShortDRX_Cycle                     *int64
	MultipleSR_Configurations          *int64
	MultipleConfiguredGrants           *int64
	SecondaryDRX_Group_r16             *int64
	EnhancedSkipUplinkTxDynamic_r16    *int64
	EnhancedSkipUplinkTxConfigured_r16 *int64
	Dummy1                             *int64
	Dummy2                             *int64
}

func (ie *MAC_ParametersXDD_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersXDDDiffConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SecondaryDRX_Group_r16 != nil
	hasExtGroup1 := ie.EnhancedSkipUplinkTxDynamic_r16 != nil || ie.EnhancedSkipUplinkTxConfigured_r16 != nil
	hasExtGroup2 := ie.Dummy1 != nil || ie.Dummy2 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SkipUplinkTxDynamic != nil, ie.LogicalChannelSR_DelayTimer != nil, ie.LongDRX_Cycle != nil, ie.ShortDRX_Cycle != nil, ie.MultipleSR_Configurations != nil, ie.MultipleConfiguredGrants != nil}); err != nil {
		return err
	}

	// 3. skipUplinkTxDynamic: enumerated
	{
		if ie.SkipUplinkTxDynamic != nil {
			if err := e.EncodeEnumerated(*ie.SkipUplinkTxDynamic, mACParametersXDDDiffSkipUplinkTxDynamicConstraints); err != nil {
				return err
			}
		}
	}

	// 4. logicalChannelSR-DelayTimer: enumerated
	{
		if ie.LogicalChannelSR_DelayTimer != nil {
			if err := e.EncodeEnumerated(*ie.LogicalChannelSR_DelayTimer, mACParametersXDDDiffLogicalChannelSRDelayTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 5. longDRX-Cycle: enumerated
	{
		if ie.LongDRX_Cycle != nil {
			if err := e.EncodeEnumerated(*ie.LongDRX_Cycle, mACParametersXDDDiffLongDRXCycleConstraints); err != nil {
				return err
			}
		}
	}

	// 6. shortDRX-Cycle: enumerated
	{
		if ie.ShortDRX_Cycle != nil {
			if err := e.EncodeEnumerated(*ie.ShortDRX_Cycle, mACParametersXDDDiffShortDRXCycleConstraints); err != nil {
				return err
			}
		}
	}

	// 7. multipleSR-Configurations: enumerated
	{
		if ie.MultipleSR_Configurations != nil {
			if err := e.EncodeEnumerated(*ie.MultipleSR_Configurations, mACParametersXDDDiffMultipleSRConfigurationsConstraints); err != nil {
				return err
			}
		}
	}

	// 8. multipleConfiguredGrants: enumerated
	{
		if ie.MultipleConfiguredGrants != nil {
			if err := e.EncodeEnumerated(*ie.MultipleConfiguredGrants, mACParametersXDDDiffMultipleConfiguredGrantsConstraints); err != nil {
				return err
			}
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
					{Name: "secondaryDRX-Group-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SecondaryDRX_Group_r16 != nil}); err != nil {
				return err
			}

			if ie.SecondaryDRX_Group_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondaryDRX_Group_r16, mACParametersXDDDiffExtSecondaryDRXGroupR16Constraints); err != nil {
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
					{Name: "enhancedSkipUplinkTxDynamic-r16", Optional: true},
					{Name: "enhancedSkipUplinkTxConfigured-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnhancedSkipUplinkTxDynamic_r16 != nil, ie.EnhancedSkipUplinkTxConfigured_r16 != nil}); err != nil {
				return err
			}

			if ie.EnhancedSkipUplinkTxDynamic_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedSkipUplinkTxDynamic_r16, mACParametersXDDDiffExtEnhancedSkipUplinkTxDynamicR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnhancedSkipUplinkTxConfigured_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedSkipUplinkTxConfigured_r16, mACParametersXDDDiffExtEnhancedSkipUplinkTxConfiguredR16Constraints); err != nil {
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
					{Name: "dummy1", Optional: true},
					{Name: "dummy2", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dummy1 != nil, ie.Dummy2 != nil}); err != nil {
				return err
			}

			if ie.Dummy1 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy1, mACParametersXDDDiffExtDummy1Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy2 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy2, mACParametersXDDDiffExtDummy2Constraints); err != nil {
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

func (ie *MAC_ParametersXDD_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. skipUplinkTxDynamic: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersXDDDiffSkipUplinkTxDynamicConstraints)
			if err != nil {
				return err
			}
			ie.SkipUplinkTxDynamic = &idx
		}
	}

	// 4. logicalChannelSR-DelayTimer: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mACParametersXDDDiffLogicalChannelSRDelayTimerConstraints)
			if err != nil {
				return err
			}
			ie.LogicalChannelSR_DelayTimer = &idx
		}
	}

	// 5. longDRX-Cycle: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mACParametersXDDDiffLongDRXCycleConstraints)
			if err != nil {
				return err
			}
			ie.LongDRX_Cycle = &idx
		}
	}

	// 6. shortDRX-Cycle: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mACParametersXDDDiffShortDRXCycleConstraints)
			if err != nil {
				return err
			}
			ie.ShortDRX_Cycle = &idx
		}
	}

	// 7. multipleSR-Configurations: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(mACParametersXDDDiffMultipleSRConfigurationsConstraints)
			if err != nil {
				return err
			}
			ie.MultipleSR_Configurations = &idx
		}
	}

	// 8. multipleConfiguredGrants: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(mACParametersXDDDiffMultipleConfiguredGrantsConstraints)
			if err != nil {
				return err
			}
			ie.MultipleConfiguredGrants = &idx
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
				{Name: "secondaryDRX-Group-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersXDDDiffExtSecondaryDRXGroupR16Constraints)
			if err != nil {
				return err
			}
			ie.SecondaryDRX_Group_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "enhancedSkipUplinkTxDynamic-r16", Optional: true},
				{Name: "enhancedSkipUplinkTxConfigured-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersXDDDiffExtEnhancedSkipUplinkTxDynamicR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedSkipUplinkTxDynamic_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersXDDDiffExtEnhancedSkipUplinkTxConfiguredR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedSkipUplinkTxConfigured_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dummy1", Optional: true},
				{Name: "dummy2", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersXDDDiffExtDummy1Constraints)
			if err != nil {
				return err
			}
			ie.Dummy1 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersXDDDiffExtDummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = &v
		}
	}

	return nil
}
