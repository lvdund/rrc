// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasGapSharingConfig (line 9218).

var measGapSharingConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gapSharingFR2", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var measGapSharingConfigGapSharingFR2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasGapSharingConfig_GapSharingFR2_Release = 0
	MeasGapSharingConfig_GapSharingFR2_Setup   = 1
)

var measGapSharingConfigExtGapSharingFR1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasGapSharingConfig_Ext_GapSharingFR1_Release = 0
	MeasGapSharingConfig_Ext_GapSharingFR1_Setup   = 1
)

var measGapSharingConfigExtGapSharingUEConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasGapSharingConfig_Ext_GapSharingUE_Release = 0
	MeasGapSharingConfig_Ext_GapSharingUE_Setup   = 1
)

type MeasGapSharingConfig struct {
	GapSharingFR2 *struct {
		Choice  int
		Release *struct{}
		Setup   *MeasGapSharingScheme
	}
	GapSharingFR1 *struct {
		Choice  int
		Release *struct{}
		Setup   *MeasGapSharingScheme
	}
	GapSharingUE *struct {
		Choice  int
		Release *struct{}
		Setup   *MeasGapSharingScheme
	}
}

func (ie *MeasGapSharingConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measGapSharingConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.GapSharingFR1 != nil || ie.GapSharingUE != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GapSharingFR2 != nil}); err != nil {
		return err
	}

	// 3. gapSharingFR2: choice
	{
		if ie.GapSharingFR2 != nil {
			choiceEnc := e.NewChoiceEncoder(measGapSharingConfigGapSharingFR2Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.GapSharingFR2).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.GapSharingFR2).Choice {
			case MeasGapSharingConfig_GapSharingFR2_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MeasGapSharingConfig_GapSharingFR2_Setup:
				if err := (*ie.GapSharingFR2).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.GapSharingFR2).Choice), Constraint: "index out of range"}
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
					{Name: "gapSharingFR1", Optional: true},
					{Name: "gapSharingUE", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GapSharingFR1 != nil, ie.GapSharingUE != nil}); err != nil {
				return err
			}

			if ie.GapSharingFR1 != nil {
				choiceEnc := ex.NewChoiceEncoder(measGapSharingConfigExtGapSharingFR1Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.GapSharingFR1).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.GapSharingFR1).Choice {
				case MeasGapSharingConfig_Ext_GapSharingFR1_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasGapSharingConfig_Ext_GapSharingFR1_Setup:
					if err := (*ie.GapSharingFR1).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.GapSharingUE != nil {
				choiceEnc := ex.NewChoiceEncoder(measGapSharingConfigExtGapSharingUEConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.GapSharingUE).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.GapSharingUE).Choice {
				case MeasGapSharingConfig_Ext_GapSharingUE_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasGapSharingConfig_Ext_GapSharingUE_Setup:
					if err := (*ie.GapSharingUE).Setup.Encode(ex); err != nil {
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

func (ie *MeasGapSharingConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measGapSharingConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. gapSharingFR2: choice
	{
		if seq.IsComponentPresent(0) {
			ie.GapSharingFR2 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MeasGapSharingScheme
			}{}
			choiceDec := d.NewChoiceDecoder(measGapSharingConfigGapSharingFR2Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapSharingFR2).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MeasGapSharingConfig_GapSharingFR2_Release:
				(*ie.GapSharingFR2).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MeasGapSharingConfig_GapSharingFR2_Setup:
				(*ie.GapSharingFR2).Setup = new(MeasGapSharingScheme)
				if err := (*ie.GapSharingFR2).Setup.Decode(d); err != nil {
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
				{Name: "gapSharingFR1", Optional: true},
				{Name: "gapSharingUE", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.GapSharingFR1 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MeasGapSharingScheme
			}{}
			choiceDec := dx.NewChoiceDecoder(measGapSharingConfigExtGapSharingFR1Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapSharingFR1).Choice = int(index)
			switch index {
			case MeasGapSharingConfig_Ext_GapSharingFR1_Release:
				(*ie.GapSharingFR1).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasGapSharingConfig_Ext_GapSharingFR1_Setup:
				(*ie.GapSharingFR1).Setup = new(MeasGapSharingScheme)
				if err := (*ie.GapSharingFR1).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.GapSharingUE = &struct {
				Choice  int
				Release *struct{}
				Setup   *MeasGapSharingScheme
			}{}
			choiceDec := dx.NewChoiceDecoder(measGapSharingConfigExtGapSharingUEConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.GapSharingUE).Choice = int(index)
			switch index {
			case MeasGapSharingConfig_Ext_GapSharingUE_Release:
				(*ie.GapSharingUE).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasGapSharingConfig_Ext_GapSharingUE_Setup:
				(*ie.GapSharingUE).Setup = new(MeasGapSharingScheme)
				if err := (*ie.GapSharingUE).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
