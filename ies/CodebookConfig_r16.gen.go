// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookConfig-r16 (line 6124).

var codebookConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookType"},
	},
}

var codebookConfig_r16CodebookTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type2"},
	},
}

const (
	CodebookConfig_r16_CodebookType_Type2 = 0
)

var codebookConfigR16CodebookTypeType2SubTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "typeII-r16"},
		{Name: "typeII-PortSelection-r16"},
	},
}

const (
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16               = 0
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16 = 1
)

var codebookConfigR16CodebookTypeType2SubTypeTypeIIR16N1N2CodebookSubsetRestrictionR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-one"},
		{Name: "two-two"},
		{Name: "four-one"},
		{Name: "three-two"},
		{Name: "six-one"},
		{Name: "four-two"},
		{Name: "eight-one"},
		{Name: "four-three"},
		{Name: "six-two"},
		{Name: "twelve-one"},
		{Name: "four-four"},
		{Name: "eight-two"},
		{Name: "sixteen-one"},
	},
}

const (
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Two_One     = 0
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Two_Two     = 1
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_One    = 2
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Three_Two   = 3
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Six_One     = 4
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Two    = 5
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Eight_One   = 6
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Three  = 7
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Six_Two     = 8
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Twelve_One  = 9
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Four   = 10
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Eight_Two   = 11
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Sixteen_One = 12
)

const (
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N1 = 0
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N2 = 1
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N3 = 2
	CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N4 = 3
)

var codebookConfigR16CodebookTypeType2SubTypeTypeIIPortSelectionR16PortSelectionSamplingSizeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N1, CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N2, CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N3, CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16_PortSelectionSamplingSize_r16_N4},
}

type CodebookConfig_r16 struct {
	CodebookType struct {
		Choice int
		Type2  *struct {
			SubType struct {
				Choice     int
				TypeII_r16 *struct {
					N1_N2_CodebookSubsetRestriction_r16 struct {
						Choice      int
						Two_One     *per.BitString
						Two_Two     *per.BitString
						Four_One    *per.BitString
						Three_Two   *per.BitString
						Six_One     *per.BitString
						Four_Two    *per.BitString
						Eight_One   *per.BitString
						Four_Three  *per.BitString
						Six_Two     *per.BitString
						Twelve_One  *per.BitString
						Four_Four   *per.BitString
						Eight_Two   *per.BitString
						Sixteen_One *per.BitString
					}
					TypeII_RI_Restriction_r16 per.BitString
				}
				TypeII_PortSelection_r16 *struct {
					PortSelectionSamplingSize_r16          int64
					TypeII_PortSelectionRI_Restriction_r16 per.BitString
				}
			}
			NumberOfPMI_SubbandsPerCQI_Subband_r16 int64
			ParamCombination_r16                   int64
		}
	}
}

func (ie *CodebookConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookConfigR16Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceEnc := e.NewChoiceEncoder(codebookConfig_r16CodebookTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CodebookType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CodebookType.Choice {
		case CodebookConfig_r16_CodebookType_Type2:
			c := (*ie.CodebookType.Type2)
			{
				choiceEnc := e.NewChoiceEncoder(codebookConfigR16CodebookTypeType2SubTypeConstraints)
				if err := choiceEnc.EncodeChoice(int64(c.SubType.Choice), false, nil); err != nil {
					return err
				}
				switch c.SubType.Choice {
				case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16:
					c := (*c.SubType.TypeII_r16)
					{
						choiceEnc := e.NewChoiceEncoder(codebookConfigR16CodebookTypeType2SubTypeTypeIIR16N1N2CodebookSubsetRestrictionR16Constraints)
						if err := choiceEnc.EncodeChoice(int64(c.N1_N2_CodebookSubsetRestriction_r16.Choice), false, nil); err != nil {
							return err
						}
						switch c.N1_N2_CodebookSubsetRestriction_r16.Choice {
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Two_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Two_One), per.FixedSize(16)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Two_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Two_Two), per.FixedSize(43)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Four_One), per.FixedSize(32)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Three_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Three_Two), per.FixedSize(59)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Six_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Six_One), per.FixedSize(48)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Four_Two), per.FixedSize(75)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Eight_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Eight_One), per.FixedSize(64)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Three:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Four_Three), per.FixedSize(107)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Six_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Six_Two), per.FixedSize(107)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Twelve_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Twelve_One), per.FixedSize(96)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Four:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Four_Four), per.FixedSize(139)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Eight_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Eight_Two), per.FixedSize(139)); err != nil {
								return err
							}
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Sixteen_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction_r16.Sixteen_One), per.FixedSize(128)); err != nil {
								return err
							}
						}
					}
					if err := e.EncodeBitString(c.TypeII_RI_Restriction_r16, per.FixedSize(4)); err != nil {
						return err
					}
				case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16:
					c := (*c.SubType.TypeII_PortSelection_r16)
					if err := e.EncodeEnumerated(c.PortSelectionSamplingSize_r16, codebookConfigR16CodebookTypeType2SubTypeTypeIIPortSelectionR16PortSelectionSamplingSizeR16Constraints); err != nil {
						return err
					}
					if err := e.EncodeBitString(c.TypeII_PortSelectionRI_Restriction_r16, per.FixedSize(4)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.NumberOfPMI_SubbandsPerCQI_Subband_r16, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ParamCombination_r16, per.Constrained(1, 8)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CodebookType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CodebookConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookConfigR16Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceDec := d.NewChoiceDecoder(codebookConfig_r16CodebookTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CodebookType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CodebookConfig_r16_CodebookType_Type2:
			ie.CodebookType.Type2 = &struct {
				SubType struct {
					Choice     int
					TypeII_r16 *struct {
						N1_N2_CodebookSubsetRestriction_r16 struct {
							Choice      int
							Two_One     *per.BitString
							Two_Two     *per.BitString
							Four_One    *per.BitString
							Three_Two   *per.BitString
							Six_One     *per.BitString
							Four_Two    *per.BitString
							Eight_One   *per.BitString
							Four_Three  *per.BitString
							Six_Two     *per.BitString
							Twelve_One  *per.BitString
							Four_Four   *per.BitString
							Eight_Two   *per.BitString
							Sixteen_One *per.BitString
						}
						TypeII_RI_Restriction_r16 per.BitString
					}
					TypeII_PortSelection_r16 *struct {
						PortSelectionSamplingSize_r16          int64
						TypeII_PortSelectionRI_Restriction_r16 per.BitString
					}
				}
				NumberOfPMI_SubbandsPerCQI_Subband_r16 int64
				ParamCombination_r16                   int64
			}{}
			c := (*ie.CodebookType.Type2)
			{
				choiceDec := d.NewChoiceDecoder(codebookConfigR16CodebookTypeType2SubTypeConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.SubType.Choice = int(index)
				switch index {
				case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16:
					c.SubType.TypeII_r16 = &struct {
						N1_N2_CodebookSubsetRestriction_r16 struct {
							Choice      int
							Two_One     *per.BitString
							Two_Two     *per.BitString
							Four_One    *per.BitString
							Three_Two   *per.BitString
							Six_One     *per.BitString
							Four_Two    *per.BitString
							Eight_One   *per.BitString
							Four_Three  *per.BitString
							Six_Two     *per.BitString
							Twelve_One  *per.BitString
							Four_Four   *per.BitString
							Eight_Two   *per.BitString
							Sixteen_One *per.BitString
						}
						TypeII_RI_Restriction_r16 per.BitString
					}{}
					c := (*c.SubType.TypeII_r16)
					{
						choiceDec := d.NewChoiceDecoder(codebookConfigR16CodebookTypeType2SubTypeTypeIIR16N1N2CodebookSubsetRestrictionR16Constraints)
						index, _, _, err := choiceDec.DecodeChoice()
						if err != nil {
							return err
						}
						c.N1_N2_CodebookSubsetRestriction_r16.Choice = int(index)
						switch index {
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Two_One:
							c.N1_N2_CodebookSubsetRestriction_r16.Two_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(16))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Two_One) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Two_Two:
							c.N1_N2_CodebookSubsetRestriction_r16.Two_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(43))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Two_Two) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_One:
							c.N1_N2_CodebookSubsetRestriction_r16.Four_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(32))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Four_One) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Three_Two:
							c.N1_N2_CodebookSubsetRestriction_r16.Three_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(59))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Three_Two) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Six_One:
							c.N1_N2_CodebookSubsetRestriction_r16.Six_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(48))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Six_One) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Two:
							c.N1_N2_CodebookSubsetRestriction_r16.Four_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(75))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Four_Two) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Eight_One:
							c.N1_N2_CodebookSubsetRestriction_r16.Eight_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(64))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Eight_One) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Three:
							c.N1_N2_CodebookSubsetRestriction_r16.Four_Three = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(107))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Four_Three) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Six_Two:
							c.N1_N2_CodebookSubsetRestriction_r16.Six_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(107))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Six_Two) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Twelve_One:
							c.N1_N2_CodebookSubsetRestriction_r16.Twelve_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(96))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Twelve_One) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Four_Four:
							c.N1_N2_CodebookSubsetRestriction_r16.Four_Four = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(139))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Four_Four) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Eight_Two:
							c.N1_N2_CodebookSubsetRestriction_r16.Eight_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(139))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Eight_Two) = v
						case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_r16_N1_N2_CodebookSubsetRestriction_r16_Sixteen_One:
							c.N1_N2_CodebookSubsetRestriction_r16.Sixteen_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(128))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction_r16.Sixteen_One) = v
						}
					}
					{
						v, err := d.DecodeBitString(per.FixedSize(4))
						if err != nil {
							return err
						}
						c.TypeII_RI_Restriction_r16 = v
					}
				case CodebookConfig_r16_CodebookType_Type2_SubType_TypeII_PortSelection_r16:
					c.SubType.TypeII_PortSelection_r16 = &struct {
						PortSelectionSamplingSize_r16          int64
						TypeII_PortSelectionRI_Restriction_r16 per.BitString
					}{}
					c := (*c.SubType.TypeII_PortSelection_r16)
					{
						v, err := d.DecodeEnumerated(codebookConfigR16CodebookTypeType2SubTypeTypeIIPortSelectionR16PortSelectionSamplingSizeR16Constraints)
						if err != nil {
							return err
						}
						c.PortSelectionSamplingSize_r16 = v
					}
					{
						v, err := d.DecodeBitString(per.FixedSize(4))
						if err != nil {
							return err
						}
						c.TypeII_PortSelectionRI_Restriction_r16 = v
					}
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.NumberOfPMI_SubbandsPerCQI_Subband_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.ParamCombination_r16 = v
			}
		}
	}

	return nil
}
