// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersSidelinkCommon-r16 (line 25074).

var mACParametersSidelinkCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lcp-RestrictionSidelink-r16", Optional: true},
		{Name: "multipleConfiguredGrantsSidelink-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	MAC_ParametersSidelinkCommon_r16_Lcp_RestrictionSidelink_r16_Supported = 0
)

var mACParametersSidelinkCommonR16LcpRestrictionSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersSidelinkCommon_r16_Lcp_RestrictionSidelink_r16_Supported},
}

const (
	MAC_ParametersSidelinkCommon_r16_MultipleConfiguredGrantsSidelink_r16_Supported = 0
)

var mACParametersSidelinkCommonR16MultipleConfiguredGrantsSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersSidelinkCommon_r16_MultipleConfiguredGrantsSidelink_r16_Supported},
}

const (
	MAC_ParametersSidelinkCommon_r16_Ext_Drx_OnSidelink_r17_Supported = 0
)

var mACParametersSidelinkCommonR16ExtDrxOnSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersSidelinkCommon_r16_Ext_Drx_OnSidelink_r17_Supported},
}

const (
	MAC_ParametersSidelinkCommon_r16_Ext_Sl_LBT_FailureDectectionRecovery_r18_Supported = 0
)

var mACParametersSidelinkCommonR16ExtSlLBTFailureDectectionRecoveryR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersSidelinkCommon_r16_Ext_Sl_LBT_FailureDectectionRecovery_r18_Supported},
}

type MAC_ParametersSidelinkCommon_r16 struct {
	Lcp_RestrictionSidelink_r16          *int64
	MultipleConfiguredGrantsSidelink_r16 *int64
	Drx_OnSidelink_r17                   *int64
	Sl_LBT_FailureDectectionRecovery_r18 *int64
}

func (ie *MAC_ParametersSidelinkCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersSidelinkCommonR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Drx_OnSidelink_r17 != nil
	hasExtGroup1 := ie.Sl_LBT_FailureDectectionRecovery_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lcp_RestrictionSidelink_r16 != nil, ie.MultipleConfiguredGrantsSidelink_r16 != nil}); err != nil {
		return err
	}

	// 3. lcp-RestrictionSidelink-r16: enumerated
	{
		if ie.Lcp_RestrictionSidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Lcp_RestrictionSidelink_r16, mACParametersSidelinkCommonR16LcpRestrictionSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. multipleConfiguredGrantsSidelink-r16: enumerated
	{
		if ie.MultipleConfiguredGrantsSidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MultipleConfiguredGrantsSidelink_r16, mACParametersSidelinkCommonR16MultipleConfiguredGrantsSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "drx-OnSidelink-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Drx_OnSidelink_r17 != nil}); err != nil {
				return err
			}

			if ie.Drx_OnSidelink_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Drx_OnSidelink_r17, mACParametersSidelinkCommonR16ExtDrxOnSidelinkR17Constraints); err != nil {
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
					{Name: "sl-LBT-FailureDectectionRecovery-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_LBT_FailureDectectionRecovery_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_LBT_FailureDectectionRecovery_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_LBT_FailureDectectionRecovery_r18, mACParametersSidelinkCommonR16ExtSlLBTFailureDectectionRecoveryR18Constraints); err != nil {
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

func (ie *MAC_ParametersSidelinkCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersSidelinkCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. lcp-RestrictionSidelink-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersSidelinkCommonR16LcpRestrictionSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Lcp_RestrictionSidelink_r16 = &idx
		}
	}

	// 4. multipleConfiguredGrantsSidelink-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mACParametersSidelinkCommonR16MultipleConfiguredGrantsSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.MultipleConfiguredGrantsSidelink_r16 = &idx
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
				{Name: "drx-OnSidelink-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersSidelinkCommonR16ExtDrxOnSidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.Drx_OnSidelink_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-LBT-FailureDectectionRecovery-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersSidelinkCommonR16ExtSlLBTFailureDectectionRecoveryR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LBT_FailureDectectionRecovery_r18 = &v
		}
	}

	return nil
}
