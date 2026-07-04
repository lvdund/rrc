// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSets (line 19997).

var featureSetsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSetsDownlink", Optional: true},
		{Name: "featureSetsDownlinkPerCC", Optional: true},
		{Name: "featureSetsUplink", Optional: true},
		{Name: "featureSetsUplinkPerCC", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
		{Name: "extension-addition-7", Optional: true},
		{Name: "extension-addition-8", Optional: true},
		{Name: "extension-addition-9", Optional: true},
		{Name: "extension-addition-10", Optional: true},
		{Name: "extension-addition-11", Optional: true},
		{Name: "extension-addition-12", Optional: true},
		{Name: "extension-addition-13", Optional: true},
		{Name: "extension-addition-14", Optional: true},
	},
}

var featureSetsFeatureSetsDownlinkConstraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsFeatureSetsDownlinkPerCCConstraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsFeatureSetsUplinkConstraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsFeatureSetsUplinkPerCCConstraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV1540Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsUplinkV1540Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1540Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV15a0Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkV1610Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsUplinkV1610Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetDownlinkPerCCV1620Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkV1630Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsUplinkV1640Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkV1700Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1700Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkV1710Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1700Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV1720Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1720Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkV1720Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkV1730Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1730Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1780Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1780Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV1800Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1800Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkV1800Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1800Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV1830Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1840Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1840Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkV1850Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1850Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV1860Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsDownlinkPerCCV1900Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsUplinkPerCCV1900Constraints = per.SizeRange(1, common.MaxPerCC_FeatureSets)

var featureSetsExtFeatureSetsDownlinkV1900Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

var featureSetsExtFeatureSetsUplinkV1900Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

type FeatureSets struct {
	FeatureSetsDownlink            []FeatureSetDownlink
	FeatureSetsDownlinkPerCC       []FeatureSetDownlinkPerCC
	FeatureSetsUplink              []FeatureSetUplink
	FeatureSetsUplinkPerCC         []FeatureSetUplinkPerCC
	FeatureSetsDownlink_v1540      []FeatureSetDownlink_v1540
	FeatureSetsUplink_v1540        []FeatureSetUplink_v1540
	FeatureSetsUplinkPerCC_v1540   []FeatureSetUplinkPerCC_v1540
	FeatureSetsDownlink_V15a0      []FeatureSetDownlink_V15a0
	FeatureSetsDownlink_v1610      []FeatureSetDownlink_v1610
	FeatureSetsUplink_v1610        []FeatureSetUplink_v1610
	FeatureSetDownlinkPerCC_v1620  []FeatureSetDownlinkPerCC_v1620
	FeatureSetsUplink_v1630        []FeatureSetUplink_v1630
	FeatureSetsUplink_v1640        []FeatureSetUplink_v1640
	FeatureSetsDownlink_v1700      []FeatureSetDownlink_v1700
	FeatureSetsDownlinkPerCC_v1700 []FeatureSetDownlinkPerCC_v1700
	FeatureSetsUplink_v1710        []FeatureSetUplink_v1710
	FeatureSetsUplinkPerCC_v1700   []FeatureSetUplinkPerCC_v1700
	FeatureSetsDownlink_v1720      []FeatureSetDownlink_v1720
	FeatureSetsDownlinkPerCC_v1720 []FeatureSetDownlinkPerCC_v1720
	FeatureSetsUplink_v1720        []FeatureSetUplink_v1720
	FeatureSetsDownlink_v1730      []FeatureSetDownlink_v1730
	FeatureSetsDownlinkPerCC_v1730 []FeatureSetDownlinkPerCC_v1730
	FeatureSetsDownlinkPerCC_v1780 []FeatureSetDownlinkPerCC_v1780
	FeatureSetsUplinkPerCC_v1780   []FeatureSetUplinkPerCC_v1780
	FeatureSetsDownlink_v1800      []FeatureSetDownlink_v1800
	FeatureSetsDownlinkPerCC_v1800 []FeatureSetDownlinkPerCC_v1800
	FeatureSetsUplink_v1800        []FeatureSetUplink_v1800
	FeatureSetsUplinkPerCC_v1800   []FeatureSetUplinkPerCC_v1800
	FeatureSetsDownlink_v1830      []FeatureSetDownlink_v1830
	FeatureSetsDownlinkPerCC_v1840 []FeatureSetDownlinkPerCC_v1840
	FeatureSetsUplinkPerCC_v1840   []FeatureSetUplinkPerCC_v1840
	FeatureSetsUplink_v1850        []FeatureSetUplink_v1850
	FeatureSetsUplinkPerCC_v1850   []FeatureSetUplinkPerCC_v1850
	FeatureSetsDownlink_v1860      []FeatureSetDownlink_v1860
	FeatureSetsDownlinkPerCC_v1900 []FeatureSetDownlinkPerCC_v1900
	FeatureSetsUplinkPerCC_v1900   []FeatureSetUplinkPerCC_v1900
	FeatureSetsDownlink_v1900      []FeatureSetDownlink_v1900
	FeatureSetsUplink_v1900        []FeatureSetUplink_v1900
}

func (ie *FeatureSets) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetsConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.FeatureSetsDownlink_v1540 != nil || ie.FeatureSetsUplink_v1540 != nil || ie.FeatureSetsUplinkPerCC_v1540 != nil
	hasExtGroup1 := ie.FeatureSetsDownlink_V15a0 != nil
	hasExtGroup2 := ie.FeatureSetsDownlink_v1610 != nil || ie.FeatureSetsUplink_v1610 != nil || ie.FeatureSetDownlinkPerCC_v1620 != nil
	hasExtGroup3 := ie.FeatureSetsUplink_v1630 != nil
	hasExtGroup4 := ie.FeatureSetsUplink_v1640 != nil
	hasExtGroup5 := ie.FeatureSetsDownlink_v1700 != nil || ie.FeatureSetsDownlinkPerCC_v1700 != nil || ie.FeatureSetsUplink_v1710 != nil || ie.FeatureSetsUplinkPerCC_v1700 != nil
	hasExtGroup6 := ie.FeatureSetsDownlink_v1720 != nil || ie.FeatureSetsDownlinkPerCC_v1720 != nil || ie.FeatureSetsUplink_v1720 != nil
	hasExtGroup7 := ie.FeatureSetsDownlink_v1730 != nil || ie.FeatureSetsDownlinkPerCC_v1730 != nil
	hasExtGroup8 := ie.FeatureSetsDownlinkPerCC_v1780 != nil || ie.FeatureSetsUplinkPerCC_v1780 != nil
	hasExtGroup9 := ie.FeatureSetsDownlink_v1800 != nil || ie.FeatureSetsDownlinkPerCC_v1800 != nil || ie.FeatureSetsUplink_v1800 != nil || ie.FeatureSetsUplinkPerCC_v1800 != nil
	hasExtGroup10 := ie.FeatureSetsDownlink_v1830 != nil
	hasExtGroup11 := ie.FeatureSetsDownlinkPerCC_v1840 != nil || ie.FeatureSetsUplinkPerCC_v1840 != nil
	hasExtGroup12 := ie.FeatureSetsUplink_v1850 != nil || ie.FeatureSetsUplinkPerCC_v1850 != nil
	hasExtGroup13 := ie.FeatureSetsDownlink_v1860 != nil
	hasExtGroup14 := ie.FeatureSetsDownlinkPerCC_v1900 != nil || ie.FeatureSetsUplinkPerCC_v1900 != nil || ie.FeatureSetsDownlink_v1900 != nil || ie.FeatureSetsUplink_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12 || hasExtGroup13 || hasExtGroup14

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSetsDownlink != nil, ie.FeatureSetsDownlinkPerCC != nil, ie.FeatureSetsUplink != nil, ie.FeatureSetsUplinkPerCC != nil}); err != nil {
		return err
	}

	// 3. featureSetsDownlink: sequence-of
	{
		if ie.FeatureSetsDownlink != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsFeatureSetsDownlinkConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsDownlink {
				if err := ie.FeatureSetsDownlink[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. featureSetsDownlinkPerCC: sequence-of
	{
		if ie.FeatureSetsDownlinkPerCC != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsFeatureSetsDownlinkPerCCConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsDownlinkPerCC {
				if err := ie.FeatureSetsDownlinkPerCC[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. featureSetsUplink: sequence-of
	{
		if ie.FeatureSetsUplink != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsFeatureSetsUplinkConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsUplink {
				if err := ie.FeatureSetsUplink[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. featureSetsUplinkPerCC: sequence-of
	{
		if ie.FeatureSetsUplinkPerCC != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsFeatureSetsUplinkPerCCConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsUplinkPerCC {
				if err := ie.FeatureSetsUplinkPerCC[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10, hasExtGroup11, hasExtGroup12, hasExtGroup13, hasExtGroup14}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1540", Optional: true},
					{Name: "featureSetsUplink-v1540", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1540", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1540 != nil, ie.FeatureSetsUplink_v1540 != nil, ie.FeatureSetsUplinkPerCC_v1540 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1540 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1540Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1540))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1540 {
					if err := ie.FeatureSetsDownlink_v1540[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplink_v1540 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1540Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1540))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1540 {
					if err := ie.FeatureSetsUplink_v1540[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1540 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1540Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1540))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1540 {
					if err := ie.FeatureSetsUplinkPerCC_v1540[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "featureSetsDownlink-v15a0", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_V15a0 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_V15a0 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV15a0Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_V15a0))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_V15a0 {
					if err := ie.FeatureSetsDownlink_V15a0[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "featureSetsDownlink-v1610", Optional: true},
					{Name: "featureSetsUplink-v1610", Optional: true},
					{Name: "featureSetDownlinkPerCC-v1620", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1610 != nil, ie.FeatureSetsUplink_v1610 != nil, ie.FeatureSetDownlinkPerCC_v1620 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1610))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1610 {
					if err := ie.FeatureSetsDownlink_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplink_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1610))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1610 {
					if err := ie.FeatureSetsUplink_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetDownlinkPerCC_v1620 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetDownlinkPerCCV1620Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetDownlinkPerCC_v1620))); err != nil {
					return err
				}
				for i := range ie.FeatureSetDownlinkPerCC_v1620 {
					if err := ie.FeatureSetDownlinkPerCC_v1620[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "featureSetsUplink-v1630", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsUplink_v1630 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsUplink_v1630 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1630Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1630))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1630 {
					if err := ie.FeatureSetsUplink_v1630[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsUplink-v1640", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsUplink_v1640 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsUplink_v1640 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1640Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1640))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1640 {
					if err := ie.FeatureSetsUplink_v1640[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1700", Optional: true},
					{Name: "featureSetsDownlinkPerCC-v1700", Optional: true},
					{Name: "featureSetsUplink-v1710", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1700 != nil, ie.FeatureSetsDownlinkPerCC_v1700 != nil, ie.FeatureSetsUplink_v1710 != nil, ie.FeatureSetsUplinkPerCC_v1700 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1700 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1700Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1700))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1700 {
					if err := ie.FeatureSetsDownlink_v1700[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsDownlinkPerCC_v1700 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1700Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1700))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1700 {
					if err := ie.FeatureSetsDownlinkPerCC_v1700[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplink_v1710 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1710Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1710))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1710 {
					if err := ie.FeatureSetsUplink_v1710[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1700 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1700Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1700))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1700 {
					if err := ie.FeatureSetsUplinkPerCC_v1700[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1720", Optional: true},
					{Name: "featureSetsDownlinkPerCC-v1720", Optional: true},
					{Name: "featureSetsUplink-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1720 != nil, ie.FeatureSetsDownlinkPerCC_v1720 != nil, ie.FeatureSetsUplink_v1720 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1720 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1720Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1720))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1720 {
					if err := ie.FeatureSetsDownlink_v1720[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsDownlinkPerCC_v1720 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1720Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1720))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1720 {
					if err := ie.FeatureSetsDownlinkPerCC_v1720[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplink_v1720 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1720Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1720))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1720 {
					if err := ie.FeatureSetsUplink_v1720[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 7:
		if hasExtGroup7 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1730", Optional: true},
					{Name: "featureSetsDownlinkPerCC-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1730 != nil, ie.FeatureSetsDownlinkPerCC_v1730 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1730 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1730Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1730))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1730 {
					if err := ie.FeatureSetsDownlink_v1730[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsDownlinkPerCC_v1730 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1730Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1730))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1730 {
					if err := ie.FeatureSetsDownlinkPerCC_v1730[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 8:
		if hasExtGroup8 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlinkPerCC-v1780", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1780", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlinkPerCC_v1780 != nil, ie.FeatureSetsUplinkPerCC_v1780 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlinkPerCC_v1780 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1780Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1780))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1780 {
					if err := ie.FeatureSetsDownlinkPerCC_v1780[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1780 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1780Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1780))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1780 {
					if err := ie.FeatureSetsUplinkPerCC_v1780[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 9:
		if hasExtGroup9 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1800", Optional: true},
					{Name: "featureSetsDownlinkPerCC-v1800", Optional: true},
					{Name: "featureSetsUplink-v1800", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1800 != nil, ie.FeatureSetsDownlinkPerCC_v1800 != nil, ie.FeatureSetsUplink_v1800 != nil, ie.FeatureSetsUplinkPerCC_v1800 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1800))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1800 {
					if err := ie.FeatureSetsDownlink_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsDownlinkPerCC_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1800))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1800 {
					if err := ie.FeatureSetsDownlinkPerCC_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplink_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1800))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1800 {
					if err := ie.FeatureSetsUplink_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1800))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1800 {
					if err := ie.FeatureSetsUplinkPerCC_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 10:
		if hasExtGroup10 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1830", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1830 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1830 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1830Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1830))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1830 {
					if err := ie.FeatureSetsDownlink_v1830[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 11:
		if hasExtGroup11 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlinkPerCC-v1840", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1840", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlinkPerCC_v1840 != nil, ie.FeatureSetsUplinkPerCC_v1840 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlinkPerCC_v1840 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1840Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1840))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1840 {
					if err := ie.FeatureSetsDownlinkPerCC_v1840[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1840 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1840Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1840))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1840 {
					if err := ie.FeatureSetsUplinkPerCC_v1840[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 12:
		if hasExtGroup12 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsUplink-v1850", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1850", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsUplink_v1850 != nil, ie.FeatureSetsUplinkPerCC_v1850 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsUplink_v1850 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1850Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1850))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1850 {
					if err := ie.FeatureSetsUplink_v1850[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1850 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1850Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1850))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1850 {
					if err := ie.FeatureSetsUplinkPerCC_v1850[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 13:
		if hasExtGroup13 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlink-v1860", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlink_v1860 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlink_v1860 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1860Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1860))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1860 {
					if err := ie.FeatureSetsDownlink_v1860[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 14:
		if hasExtGroup14 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featureSetsDownlinkPerCC-v1900", Optional: true},
					{Name: "featureSetsUplinkPerCC-v1900", Optional: true},
					{Name: "featureSetsDownlink-v1900", Optional: true},
					{Name: "featureSetsUplink-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeatureSetsDownlinkPerCC_v1900 != nil, ie.FeatureSetsUplinkPerCC_v1900 != nil, ie.FeatureSetsDownlink_v1900 != nil, ie.FeatureSetsUplink_v1900 != nil}); err != nil {
				return err
			}

			if ie.FeatureSetsDownlinkPerCC_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkPerCCV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlinkPerCC_v1900))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlinkPerCC_v1900 {
					if err := ie.FeatureSetsDownlinkPerCC_v1900[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplinkPerCC_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkPerCCV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplinkPerCC_v1900))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplinkPerCC_v1900 {
					if err := ie.FeatureSetsUplinkPerCC_v1900[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsDownlink_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsDownlinkV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_v1900))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsDownlink_v1900 {
					if err := ie.FeatureSetsDownlink_v1900[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FeatureSetsUplink_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(featureSetsExtFeatureSetsUplinkV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_v1900))); err != nil {
					return err
				}
				for i := range ie.FeatureSetsUplink_v1900 {
					if err := ie.FeatureSetsUplink_v1900[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *FeatureSets) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetsConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. featureSetsDownlink: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(featureSetsFeatureSetsDownlinkConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink = make([]FeatureSetDownlink, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. featureSetsDownlinkPerCC: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(featureSetsFeatureSetsDownlinkPerCCConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC = make([]FeatureSetDownlinkPerCC, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. featureSetsUplink: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(featureSetsFeatureSetsUplinkConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink = make([]FeatureSetUplink, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. featureSetsUplinkPerCC: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(featureSetsFeatureSetsUplinkPerCCConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC = make([]FeatureSetUplinkPerCC, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC[i].Decode(d); err != nil {
					return err
				}
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
				{Name: "featureSetsDownlink-v1540", Optional: true},
				{Name: "featureSetsUplink-v1540", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1540", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1540Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1540 = make([]FeatureSetDownlink_v1540, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1540[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1540Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1540 = make([]FeatureSetUplink_v1540, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1540[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1540Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1540 = make([]FeatureSetUplinkPerCC_v1540, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1540[i].Decode(dx); err != nil {
					return err
				}
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
				{Name: "featureSetsDownlink-v15a0", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV15a0Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_V15a0 = make([]FeatureSetDownlink_V15a0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_V15a0[i].Decode(dx); err != nil {
					return err
				}
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
				{Name: "featureSetsDownlink-v1610", Optional: true},
				{Name: "featureSetsUplink-v1610", Optional: true},
				{Name: "featureSetDownlinkPerCC-v1620", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1610 = make([]FeatureSetDownlink_v1610, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1610 = make([]FeatureSetUplink_v1610, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetDownlinkPerCCV1620Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetDownlinkPerCC_v1620 = make([]FeatureSetDownlinkPerCC_v1620, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetDownlinkPerCC_v1620[i].Decode(dx); err != nil {
					return err
				}
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
				{Name: "featureSetsUplink-v1630", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1630Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1630 = make([]FeatureSetUplink_v1630, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1630[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsUplink-v1640", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1640Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1640 = make([]FeatureSetUplink_v1640, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1640[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlink-v1700", Optional: true},
				{Name: "featureSetsDownlinkPerCC-v1700", Optional: true},
				{Name: "featureSetsUplink-v1710", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1700 = make([]FeatureSetDownlink_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1700[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1700 = make([]FeatureSetDownlinkPerCC_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1700[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1710Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1710 = make([]FeatureSetUplink_v1710, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1710[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1700 = make([]FeatureSetUplinkPerCC_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1700[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlink-v1720", Optional: true},
				{Name: "featureSetsDownlinkPerCC-v1720", Optional: true},
				{Name: "featureSetsUplink-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1720Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1720 = make([]FeatureSetDownlink_v1720, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1720[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1720Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1720 = make([]FeatureSetDownlinkPerCC_v1720, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1720[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1720Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1720 = make([]FeatureSetUplink_v1720, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1720[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlink-v1730", Optional: true},
				{Name: "featureSetsDownlinkPerCC-v1730", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1730Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1730 = make([]FeatureSetDownlink_v1730, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1730[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1730Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1730 = make([]FeatureSetDownlinkPerCC_v1730, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1730[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlinkPerCC-v1780", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1780", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1780Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1780 = make([]FeatureSetDownlinkPerCC_v1780, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1780[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1780Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1780 = make([]FeatureSetUplinkPerCC_v1780, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1780[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlink-v1800", Optional: true},
				{Name: "featureSetsDownlinkPerCC-v1800", Optional: true},
				{Name: "featureSetsUplink-v1800", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1800 = make([]FeatureSetDownlink_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1800 = make([]FeatureSetDownlinkPerCC_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1800 = make([]FeatureSetUplink_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1800 = make([]FeatureSetUplinkPerCC_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlink-v1830", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1830Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1830 = make([]FeatureSetDownlink_v1830, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1830[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlinkPerCC-v1840", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1840", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1840Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1840 = make([]FeatureSetDownlinkPerCC_v1840, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1840[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1840Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1840 = make([]FeatureSetUplinkPerCC_v1840, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1840[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsUplink-v1850", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1850", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1850Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1850 = make([]FeatureSetUplink_v1850, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1850[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1850Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1850 = make([]FeatureSetUplinkPerCC_v1850, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1850[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 13:
	if seq.IsExtensionPresent(13) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlink-v1860", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1860Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1860 = make([]FeatureSetDownlink_v1860, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1860[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 14:
	if seq.IsExtensionPresent(14) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featureSetsDownlinkPerCC-v1900", Optional: true},
				{Name: "featureSetsUplinkPerCC-v1900", Optional: true},
				{Name: "featureSetsDownlink-v1900", Optional: true},
				{Name: "featureSetsUplink-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkPerCCV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlinkPerCC_v1900 = make([]FeatureSetDownlinkPerCC_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlinkPerCC_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkPerCCV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplinkPerCC_v1900 = make([]FeatureSetUplinkPerCC_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplinkPerCC_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsDownlinkV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_v1900 = make([]FeatureSetDownlink_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(featureSetsExtFeatureSetsUplinkV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_v1900 = make([]FeatureSetUplink_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
