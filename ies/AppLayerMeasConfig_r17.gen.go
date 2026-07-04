// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AppLayerMeasConfig-r17 (line 26010).

var appLayerMeasConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measConfigAppLayerToAddModList-r17", Optional: true},
		{Name: "measConfigAppLayerToReleaseList-r17", Optional: true},
		{Name: "rrc-SegAllowedSRB4-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var appLayerMeasConfigR17MeasConfigAppLayerToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofAppLayerMeas_r17)

var appLayerMeasConfigR17MeasConfigAppLayerToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofAppLayerMeas_r17)

const (
	AppLayerMeasConfig_r17_Rrc_SegAllowedSRB4_r17_Enabled = 0
)

var appLayerMeasConfigR17RrcSegAllowedSRB4R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasConfig_r17_Rrc_SegAllowedSRB4_r17_Enabled},
}

const (
	AppLayerMeasConfig_r17_Ext_Rrc_SegAllowedSRB5_r18_Enabled = 0
)

var appLayerMeasConfigR17ExtRrcSegAllowedSRB5R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasConfig_r17_Ext_Rrc_SegAllowedSRB5_r18_Enabled},
}

const (
	AppLayerMeasConfig_r17_Ext_IdleInactiveReportAllowed_r18_Enabled = 0
)

var appLayerMeasConfigR17ExtIdleInactiveReportAllowedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AppLayerMeasConfig_r17_Ext_IdleInactiveReportAllowed_r18_Enabled},
}

type AppLayerMeasConfig_r17 struct {
	MeasConfigAppLayerToAddModList_r17  []MeasConfigAppLayer_r17
	MeasConfigAppLayerToReleaseList_r17 []MeasConfigAppLayerId_r17
	Rrc_SegAllowedSRB4_r17              *int64
	Rrc_SegAllowedSRB5_r18              *int64
	IdleInactiveReportAllowed_r18       *int64
}

func (ie *AppLayerMeasConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(appLayerMeasConfigR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Rrc_SegAllowedSRB5_r18 != nil || ie.IdleInactiveReportAllowed_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasConfigAppLayerToAddModList_r17 != nil, ie.MeasConfigAppLayerToReleaseList_r17 != nil, ie.Rrc_SegAllowedSRB4_r17 != nil}); err != nil {
		return err
	}

	// 3. measConfigAppLayerToAddModList-r17: sequence-of
	{
		if ie.MeasConfigAppLayerToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(appLayerMeasConfigR17MeasConfigAppLayerToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.MeasConfigAppLayerToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.MeasConfigAppLayerToAddModList_r17 {
				if err := ie.MeasConfigAppLayerToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. measConfigAppLayerToReleaseList-r17: sequence-of
	{
		if ie.MeasConfigAppLayerToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(appLayerMeasConfigR17MeasConfigAppLayerToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.MeasConfigAppLayerToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.MeasConfigAppLayerToReleaseList_r17 {
				if err := ie.MeasConfigAppLayerToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. rrc-SegAllowedSRB4-r17: enumerated
	{
		if ie.Rrc_SegAllowedSRB4_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Rrc_SegAllowedSRB4_r17, appLayerMeasConfigR17RrcSegAllowedSRB4R17Constraints); err != nil {
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
					{Name: "rrc-SegAllowedSRB5-r18", Optional: true},
					{Name: "idleInactiveReportAllowed-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rrc_SegAllowedSRB5_r18 != nil, ie.IdleInactiveReportAllowed_r18 != nil}); err != nil {
				return err
			}

			if ie.Rrc_SegAllowedSRB5_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Rrc_SegAllowedSRB5_r18, appLayerMeasConfigR17ExtRrcSegAllowedSRB5R18Constraints); err != nil {
					return err
				}
			}

			if ie.IdleInactiveReportAllowed_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.IdleInactiveReportAllowed_r18, appLayerMeasConfigR17ExtIdleInactiveReportAllowedR18Constraints); err != nil {
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

func (ie *AppLayerMeasConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(appLayerMeasConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measConfigAppLayerToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(appLayerMeasConfigR17MeasConfigAppLayerToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasConfigAppLayerToAddModList_r17 = make([]MeasConfigAppLayer_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasConfigAppLayerToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. measConfigAppLayerToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(appLayerMeasConfigR17MeasConfigAppLayerToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasConfigAppLayerToReleaseList_r17 = make([]MeasConfigAppLayerId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasConfigAppLayerToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. rrc-SegAllowedSRB4-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(appLayerMeasConfigR17RrcSegAllowedSRB4R17Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_SegAllowedSRB4_r17 = &idx
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
				{Name: "rrc-SegAllowedSRB5-r18", Optional: true},
				{Name: "idleInactiveReportAllowed-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(appLayerMeasConfigR17ExtRrcSegAllowedSRB5R18Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_SegAllowedSRB5_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(appLayerMeasConfigR17ExtIdleInactiveReportAllowedR18Constraints)
			if err != nil {
				return err
			}
			ie.IdleInactiveReportAllowed_r18 = &v
		}
	}

	return nil
}
