// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasGapConfig (line 9146).

var measGapConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gapFR2", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var measGapConfigGapFR2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasGapConfig_GapFR2_Release = 0
	MeasGapConfig_GapFR2_Setup   = 1
)

var measGapConfigExtGapFR1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasGapConfig_Ext_GapFR1_Release = 0
	MeasGapConfig_Ext_GapFR1_Setup   = 1
)

var measGapConfigExtGapUEConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasGapConfig_Ext_GapUE_Release = 0
	MeasGapConfig_Ext_GapUE_Setup   = 1
)

var measGapConfigExtGapToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofGapId_r17)

var measGapConfigExtGapToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofGapId_r17)

type MeasGapConfig struct {
	GapFR2 *struct {
		Choice  int
		Release *struct{}
		Setup   *GapConfig
	}
	GapFR1 *struct {
		Choice  int
		Release *struct{}
		Setup   *GapConfig
	}
	GapUE *struct {
		Choice  int
		Release *struct{}
		Setup   *GapConfig
	}
	GapToAddModList_r17                  []GapConfig_r17
	GapToReleaseList_r17                 []MeasGapId_r17
	PosMeasGapPreConfigToAddModList_r17  *PosMeasGapPreConfigToAddModList_r17
	PosMeasGapPreConfigToReleaseList_r17 *PosMeasGapPreConfigToReleaseList_r17
}

func (ie *MeasGapConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measGapConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.GapFR1 != nil || ie.GapUE != nil
	hasExtGroup1 := ie.GapToAddModList_r17 != nil || ie.GapToReleaseList_r17 != nil || ie.PosMeasGapPreConfigToAddModList_r17 != nil || ie.PosMeasGapPreConfigToReleaseList_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GapFR2 != nil}); err != nil {
		return err
	}

	// 3. gapFR2: choice
	{
		if ie.GapFR2 != nil {
			choiceEnc := e.NewChoiceEncoder(measGapConfigGapFR2Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.GapFR2).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.GapFR2).Choice {
			case MeasGapConfig_GapFR2_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MeasGapConfig_GapFR2_Setup:
				if err := (*ie.GapFR2).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.GapFR2).Choice), Constraint: "index out of range"}
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
					{Name: "gapFR1", Optional: true},
					{Name: "gapUE", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GapFR1 != nil, ie.GapUE != nil}); err != nil {
				return err
			}

			if ie.GapFR1 != nil {
				choiceEnc := ex.NewChoiceEncoder(measGapConfigExtGapFR1Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.GapFR1).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.GapFR1).Choice {
				case MeasGapConfig_Ext_GapFR1_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasGapConfig_Ext_GapFR1_Setup:
					if err := (*ie.GapFR1).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.GapUE != nil {
				choiceEnc := ex.NewChoiceEncoder(measGapConfigExtGapUEConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.GapUE).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.GapUE).Choice {
				case MeasGapConfig_Ext_GapUE_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasGapConfig_Ext_GapUE_Setup:
					if err := (*ie.GapUE).Setup.Encode(ex); err != nil {
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
					{Name: "gapToAddModList-r17", Optional: true},
					{Name: "gapToReleaseList-r17", Optional: true},
					{Name: "posMeasGapPreConfigToAddModList-r17", Optional: true},
					{Name: "posMeasGapPreConfigToReleaseList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GapToAddModList_r17 != nil, ie.GapToReleaseList_r17 != nil, ie.PosMeasGapPreConfigToAddModList_r17 != nil, ie.PosMeasGapPreConfigToReleaseList_r17 != nil}); err != nil {
				return err
			}

			if ie.GapToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(measGapConfigExtGapToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.GapToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.GapToAddModList_r17 {
					if err := ie.GapToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.GapToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(measGapConfigExtGapToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.GapToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.GapToReleaseList_r17 {
					if err := ie.GapToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PosMeasGapPreConfigToAddModList_r17 != nil {
				if err := ie.PosMeasGapPreConfigToAddModList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosMeasGapPreConfigToReleaseList_r17 != nil {
				if err := ie.PosMeasGapPreConfigToReleaseList_r17.Encode(ex); err != nil {
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

func (ie *MeasGapConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measGapConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. gapFR2: choice
	{
		if seq.IsComponentPresent(0) {
			ie.GapFR2 = &struct {
				Choice  int
				Release *struct{}
				Setup   *GapConfig
			}{}
			choiceDec := d.NewChoiceDecoder(measGapConfigGapFR2Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapFR2).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MeasGapConfig_GapFR2_Release:
				(*ie.GapFR2).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MeasGapConfig_GapFR2_Setup:
				(*ie.GapFR2).Setup = new(GapConfig)
				if err := (*ie.GapFR2).Setup.Decode(d); err != nil {
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
				{Name: "gapFR1", Optional: true},
				{Name: "gapUE", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.GapFR1 = &struct {
				Choice  int
				Release *struct{}
				Setup   *GapConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(measGapConfigExtGapFR1Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapFR1).Choice = int(index)
			switch index {
			case MeasGapConfig_Ext_GapFR1_Release:
				(*ie.GapFR1).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasGapConfig_Ext_GapFR1_Setup:
				(*ie.GapFR1).Setup = new(GapConfig)
				if err := (*ie.GapFR1).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.GapUE = &struct {
				Choice  int
				Release *struct{}
				Setup   *GapConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(measGapConfigExtGapUEConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapUE).Choice = int(index)
			switch index {
			case MeasGapConfig_Ext_GapUE_Release:
				(*ie.GapUE).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasGapConfig_Ext_GapUE_Setup:
				(*ie.GapUE).Setup = new(GapConfig)
				if err := (*ie.GapUE).Setup.Decode(dx); err != nil {
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
				{Name: "gapToAddModList-r17", Optional: true},
				{Name: "gapToReleaseList-r17", Optional: true},
				{Name: "posMeasGapPreConfigToAddModList-r17", Optional: true},
				{Name: "posMeasGapPreConfigToReleaseList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(measGapConfigExtGapToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.GapToAddModList_r17 = make([]GapConfig_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.GapToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(measGapConfigExtGapToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.GapToReleaseList_r17 = make([]MeasGapId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.GapToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.PosMeasGapPreConfigToAddModList_r17 = new(PosMeasGapPreConfigToAddModList_r17)
			if err := ie.PosMeasGapPreConfigToAddModList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.PosMeasGapPreConfigToReleaseList_r17 = new(PosMeasGapPreConfigToReleaseList_r17)
			if err := ie.PosMeasGapPreConfigToReleaseList_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
