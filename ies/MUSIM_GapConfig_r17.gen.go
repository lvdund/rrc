// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapConfig-r17 (line 10193).

var mUSIMGapConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-GapToReleaseList-r17", Optional: true},
		{Name: "musim-GapToAddModList-r17", Optional: true},
		{Name: "musim-AperiodicGap-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var mUSIMGapConfigR17MusimGapToReleaseListR17Constraints = per.SizeRange(1, 3)

var mUSIMGapConfigR17MusimGapToAddModListR17Constraints = per.SizeRange(1, 3)

var mUSIMGapConfigR17ExtMusimGapToAddModListExtV1820Constraints = per.SizeRange(1, 3)

const (
	MUSIM_GapConfig_r17_Ext_Musim_GapKeep_r18_True = 0
)

var mUSIMGapConfigR17ExtMusimGapKeepR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_GapConfig_r17_Ext_Musim_GapKeep_r18_True},
}

type MUSIM_GapConfig_r17 struct {
	Musim_GapToReleaseList_r17     []MUSIM_GapId_r17
	Musim_GapToAddModList_r17      []MUSIM_Gap_r17
	Musim_AperiodicGap_r17         *MUSIM_GapInfo_r17
	Musim_GapToAddModListExt_v1820 []MUSIM_GapExt_v1820
	Musim_GapKeep_r18              *int64
}

func (ie *MUSIM_GapConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMGapConfigR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Musim_GapToAddModListExt_v1820 != nil || ie.Musim_GapKeep_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_GapToReleaseList_r17 != nil, ie.Musim_GapToAddModList_r17 != nil, ie.Musim_AperiodicGap_r17 != nil}); err != nil {
		return err
	}

	// 3. musim-GapToReleaseList-r17: sequence-of
	{
		if ie.Musim_GapToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mUSIMGapConfigR17MusimGapToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Musim_GapToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Musim_GapToReleaseList_r17 {
				if err := ie.Musim_GapToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. musim-GapToAddModList-r17: sequence-of
	{
		if ie.Musim_GapToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mUSIMGapConfigR17MusimGapToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Musim_GapToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Musim_GapToAddModList_r17 {
				if err := ie.Musim_GapToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. musim-AperiodicGap-r17: ref
	{
		if ie.Musim_AperiodicGap_r17 != nil {
			if err := ie.Musim_AperiodicGap_r17.Encode(e); err != nil {
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
					{Name: "musim-GapToAddModListExt-v1820", Optional: true},
					{Name: "musim-GapKeep-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Musim_GapToAddModListExt_v1820 != nil, ie.Musim_GapKeep_r18 != nil}); err != nil {
				return err
			}

			if ie.Musim_GapToAddModListExt_v1820 != nil {
				seqOf := ex.NewSequenceOfEncoder(mUSIMGapConfigR17ExtMusimGapToAddModListExtV1820Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Musim_GapToAddModListExt_v1820))); err != nil {
					return err
				}
				for i := range ie.Musim_GapToAddModListExt_v1820 {
					if err := ie.Musim_GapToAddModListExt_v1820[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Musim_GapKeep_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Musim_GapKeep_r18, mUSIMGapConfigR17ExtMusimGapKeepR18Constraints); err != nil {
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

func (ie *MUSIM_GapConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMGapConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. musim-GapToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(mUSIMGapConfigR17MusimGapToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Musim_GapToReleaseList_r17 = make([]MUSIM_GapId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Musim_GapToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. musim-GapToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(mUSIMGapConfigR17MusimGapToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Musim_GapToAddModList_r17 = make([]MUSIM_Gap_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Musim_GapToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. musim-AperiodicGap-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Musim_AperiodicGap_r17 = new(MUSIM_GapInfo_r17)
			if err := ie.Musim_AperiodicGap_r17.Decode(d); err != nil {
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
				{Name: "musim-GapToAddModListExt-v1820", Optional: true},
				{Name: "musim-GapKeep-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(mUSIMGapConfigR17ExtMusimGapToAddModListExtV1820Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Musim_GapToAddModListExt_v1820 = make([]MUSIM_GapExt_v1820, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Musim_GapToAddModListExt_v1820[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mUSIMGapConfigR17ExtMusimGapKeepR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_GapKeep_r18 = &v
		}
	}

	return nil
}
