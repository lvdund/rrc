// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookConfig (line 6045).

var codebookConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookType"},
	},
}

var codebookConfigCodebookTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1"},
		{Name: "type2"},
	},
}

const (
	CodebookConfig_CodebookType_Type1 = 0
	CodebookConfig_CodebookType_Type2 = 1
)

var codebookConfigCodebookTypeType1SubTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "typeI-SinglePanel"},
		{Name: "typeI-MultiPanel"},
	},
}

const (
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel = 0
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel  = 1
)

var codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two"},
		{Name: "moreThanTwo"},
	},
}

const (
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_Two         = 0
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo = 1
)

var codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "n1-n2"},
		{Name: "typeI-SinglePanel-codebookSubsetRestriction-i2", Optional: true},
	},
}

var codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoN1N2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-one-TypeI-SinglePanel-Restriction"},
		{Name: "two-two-TypeI-SinglePanel-Restriction"},
		{Name: "four-one-TypeI-SinglePanel-Restriction"},
		{Name: "three-two-TypeI-SinglePanel-Restriction"},
		{Name: "six-one-TypeI-SinglePanel-Restriction"},
		{Name: "four-two-TypeI-SinglePanel-Restriction"},
		{Name: "eight-one-TypeI-SinglePanel-Restriction"},
		{Name: "four-three-TypeI-SinglePanel-Restriction"},
		{Name: "six-two-TypeI-SinglePanel-Restriction"},
		{Name: "twelve-one-TypeI-SinglePanel-Restriction"},
		{Name: "four-four-TypeI-SinglePanel-Restriction"},
		{Name: "eight-two-TypeI-SinglePanel-Restriction"},
		{Name: "sixteen-one-TypeI-SinglePanel-Restriction"},
	},
}

const (
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction     = 0
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction     = 1
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction    = 2
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction   = 3
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction     = 4
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction    = 5
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction   = 6
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction  = 7
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction     = 8
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction  = 9
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction   = 10
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction   = 11
	CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction = 12
)

var codebookConfigCodebookTypeType1SubTypeTypeIMultiPanelNgN1N2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-two-one-TypeI-MultiPanel-Restriction"},
		{Name: "two-four-one-TypeI-MultiPanel-Restriction"},
		{Name: "four-two-one-TypeI-MultiPanel-Restriction"},
		{Name: "two-two-two-TypeI-MultiPanel-Restriction"},
		{Name: "two-eight-one-TypeI-MultiPanel-Restriction"},
		{Name: "four-four-one-TypeI-MultiPanel-Restriction"},
		{Name: "two-four-two-TypeI-MultiPanel-Restriction"},
		{Name: "four-two-two-TypeI-MultiPanel-Restriction"},
	},
}

const (
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Two_One_TypeI_MultiPanel_Restriction   = 0
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Four_One_TypeI_MultiPanel_Restriction  = 1
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Two_One_TypeI_MultiPanel_Restriction  = 2
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Two_Two_TypeI_MultiPanel_Restriction   = 3
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Eight_One_TypeI_MultiPanel_Restriction = 4
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Four_One_TypeI_MultiPanel_Restriction = 5
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Four_Two_TypeI_MultiPanel_Restriction  = 6
	CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Two_Two_TypeI_MultiPanel_Restriction  = 7
)

var codebookConfigCodebookTypeType2SubTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "typeII"},
		{Name: "typeII-PortSelection"},
	},
}

const (
	CodebookConfig_CodebookType_Type2_SubType_TypeII               = 0
	CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection = 1
)

var codebookConfigCodebookTypeType2SubTypeTypeIIN1N2CodebookSubsetRestrictionConstraints = per.ChoiceConstraints{
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
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Two_One     = 0
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Two_Two     = 1
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_One    = 2
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Three_Two   = 3
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Six_One     = 4
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Two    = 5
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Eight_One   = 6
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Three  = 7
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Six_Two     = 8
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Twelve_One  = 9
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Four   = 10
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Eight_Two   = 11
	CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Sixteen_One = 12
)

var codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "portSelectionSamplingSize", Optional: true},
		{Name: "typeII-PortSelectionRI-Restriction"},
	},
}

const (
	CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N1 = 0
	CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N2 = 1
	CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N3 = 2
	CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N4 = 3
)

var codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionPortSelectionSamplingSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N1, CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N2, CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N3, CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection_PortSelectionSamplingSize_N4},
}

const (
	CodebookConfig_CodebookType_Type2_PhaseAlphabetSize_N4 = 0
	CodebookConfig_CodebookType_Type2_PhaseAlphabetSize_N8 = 1
)

var codebookConfigCodebookTypeType2PhaseAlphabetSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_CodebookType_Type2_PhaseAlphabetSize_N4, CodebookConfig_CodebookType_Type2_PhaseAlphabetSize_N8},
}

const (
	CodebookConfig_CodebookType_Type2_NumberOfBeams_Two   = 0
	CodebookConfig_CodebookType_Type2_NumberOfBeams_Three = 1
	CodebookConfig_CodebookType_Type2_NumberOfBeams_Four  = 2
)

var codebookConfigCodebookTypeType2NumberOfBeamsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_CodebookType_Type2_NumberOfBeams_Two, CodebookConfig_CodebookType_Type2_NumberOfBeams_Three, CodebookConfig_CodebookType_Type2_NumberOfBeams_Four},
}

type CodebookConfig struct {
	CodebookType struct {
		Choice int
		Type1  *struct {
			SubType struct {
				Choice            int
				TypeI_SinglePanel *struct {
					NrOfAntennaPorts struct {
						Choice      int
						Two         *struct{ TwoTX_CodebookSubsetRestriction per.BitString }
						MoreThanTwo *struct {
							N1_N2 struct {
								Choice                                    int
								Two_One_TypeI_SinglePanel_Restriction     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction *per.BitString
							}
							TypeI_SinglePanel_CodebookSubsetRestriction_I2 *per.BitString
						}
					}
					TypeI_SinglePanel_Ri_Restriction per.BitString
				}
				TypeI_MultiPanel *struct {
					Ng_N1_N2 struct {
						Choice                                     int
						Two_Two_One_TypeI_MultiPanel_Restriction   *per.BitString
						Two_Four_One_TypeI_MultiPanel_Restriction  *per.BitString
						Four_Two_One_TypeI_MultiPanel_Restriction  *per.BitString
						Two_Two_Two_TypeI_MultiPanel_Restriction   *per.BitString
						Two_Eight_One_TypeI_MultiPanel_Restriction *per.BitString
						Four_Four_One_TypeI_MultiPanel_Restriction *per.BitString
						Two_Four_Two_TypeI_MultiPanel_Restriction  *per.BitString
						Four_Two_Two_TypeI_MultiPanel_Restriction  *per.BitString
					}
					Ri_Restriction per.BitString
				}
			}
			CodebookMode int64
		}
		Type2 *struct {
			SubType struct {
				Choice int
				TypeII *struct {
					N1_N2_CodebookSubsetRestriction struct {
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
					TypeII_RI_Restriction per.BitString
				}
				TypeII_PortSelection *struct {
					PortSelectionSamplingSize          *int64
					TypeII_PortSelectionRI_Restriction per.BitString
				}
			}
			PhaseAlphabetSize int64
			SubbandAmplitude  bool
			NumberOfBeams     int64
		}
	}
}

func (ie *CodebookConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookConfigConstraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CodebookType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CodebookType.Choice {
		case CodebookConfig_CodebookType_Type1:
			c := (*ie.CodebookType.Type1)
			{
				choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeType1SubTypeConstraints)
				if err := choiceEnc.EncodeChoice(int64(c.SubType.Choice), false, nil); err != nil {
					return err
				}
				switch c.SubType.Choice {
				case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel:
					c := (*c.SubType.TypeI_SinglePanel)
					{
						choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsConstraints)
						if err := choiceEnc.EncodeChoice(int64(c.NrOfAntennaPorts.Choice), false, nil); err != nil {
							return err
						}
						switch c.NrOfAntennaPorts.Choice {
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_Two:
							c := (*c.NrOfAntennaPorts.Two)
							if err := e.EncodeBitString(c.TwoTX_CodebookSubsetRestriction, per.FixedSize(6)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo:
							c := (*c.NrOfAntennaPorts.MoreThanTwo)
							codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoSeq := e.NewSequenceEncoder(codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoConstraints)
							if err := codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoSeq.EncodePreamble([]bool{c.TypeI_SinglePanel_CodebookSubsetRestriction_I2 != nil}); err != nil {
								return err
							}
							{
								choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoN1N2Constraints)
								if err := choiceEnc.EncodeChoice(int64(c.N1_N2.Choice), false, nil); err != nil {
									return err
								}
								switch c.N1_N2.Choice {
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Two_One_TypeI_SinglePanel_Restriction), per.FixedSize(8)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction), per.FixedSize(64)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Four_One_TypeI_SinglePanel_Restriction), per.FixedSize(16)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction), per.FixedSize(96)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Six_One_TypeI_SinglePanel_Restriction), per.FixedSize(24)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction), per.FixedSize(128)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction), per.FixedSize(32)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction), per.FixedSize(192)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction), per.FixedSize(192)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction), per.FixedSize(48)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction), per.FixedSize(256)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction), per.FixedSize(256)); err != nil {
										return err
									}
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction:
									if err := e.EncodeBitString((*c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction), per.FixedSize(64)); err != nil {
										return err
									}
								}
							}
							if c.TypeI_SinglePanel_CodebookSubsetRestriction_I2 != nil {
								if err := e.EncodeBitString((*c.TypeI_SinglePanel_CodebookSubsetRestriction_I2), per.FixedSize(16)); err != nil {
									return err
								}
							}
						}
					}
					if err := e.EncodeBitString(c.TypeI_SinglePanel_Ri_Restriction, per.FixedSize(8)); err != nil {
						return err
					}
				case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel:
					c := (*c.SubType.TypeI_MultiPanel)
					{
						choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeType1SubTypeTypeIMultiPanelNgN1N2Constraints)
						if err := choiceEnc.EncodeChoice(int64(c.Ng_N1_N2.Choice), false, nil); err != nil {
							return err
						}
						switch c.Ng_N1_N2.Choice {
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Two_One_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Two_Two_One_TypeI_MultiPanel_Restriction), per.FixedSize(8)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Four_One_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Two_Four_One_TypeI_MultiPanel_Restriction), per.FixedSize(16)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Two_One_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Four_Two_One_TypeI_MultiPanel_Restriction), per.FixedSize(8)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Two_Two_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Two_Two_Two_TypeI_MultiPanel_Restriction), per.FixedSize(64)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Eight_One_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Two_Eight_One_TypeI_MultiPanel_Restriction), per.FixedSize(32)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Four_One_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Four_Four_One_TypeI_MultiPanel_Restriction), per.FixedSize(16)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Four_Two_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Two_Four_Two_TypeI_MultiPanel_Restriction), per.FixedSize(128)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Two_Two_TypeI_MultiPanel_Restriction:
							if err := e.EncodeBitString((*c.Ng_N1_N2.Four_Two_Two_TypeI_MultiPanel_Restriction), per.FixedSize(64)); err != nil {
								return err
							}
						}
					}
					if err := e.EncodeBitString(c.Ri_Restriction, per.FixedSize(4)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.CodebookMode, per.Constrained(1, 2)); err != nil {
				return err
			}
		case CodebookConfig_CodebookType_Type2:
			c := (*ie.CodebookType.Type2)
			{
				choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeType2SubTypeConstraints)
				if err := choiceEnc.EncodeChoice(int64(c.SubType.Choice), false, nil); err != nil {
					return err
				}
				switch c.SubType.Choice {
				case CodebookConfig_CodebookType_Type2_SubType_TypeII:
					c := (*c.SubType.TypeII)
					{
						choiceEnc := e.NewChoiceEncoder(codebookConfigCodebookTypeType2SubTypeTypeIIN1N2CodebookSubsetRestrictionConstraints)
						if err := choiceEnc.EncodeChoice(int64(c.N1_N2_CodebookSubsetRestriction.Choice), false, nil); err != nil {
							return err
						}
						switch c.N1_N2_CodebookSubsetRestriction.Choice {
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Two_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Two_One), per.FixedSize(16)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Two_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Two_Two), per.FixedSize(43)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Four_One), per.FixedSize(32)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Three_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Three_Two), per.FixedSize(59)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Six_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Six_One), per.FixedSize(48)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Four_Two), per.FixedSize(75)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Eight_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Eight_One), per.FixedSize(64)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Three:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Four_Three), per.FixedSize(107)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Six_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Six_Two), per.FixedSize(107)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Twelve_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Twelve_One), per.FixedSize(96)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Four:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Four_Four), per.FixedSize(139)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Eight_Two:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Eight_Two), per.FixedSize(139)); err != nil {
								return err
							}
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Sixteen_One:
							if err := e.EncodeBitString((*c.N1_N2_CodebookSubsetRestriction.Sixteen_One), per.FixedSize(128)); err != nil {
								return err
							}
						}
					}
					if err := e.EncodeBitString(c.TypeII_RI_Restriction, per.FixedSize(2)); err != nil {
						return err
					}
				case CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection:
					c := (*c.SubType.TypeII_PortSelection)
					codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionSeq := e.NewSequenceEncoder(codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionConstraints)
					if err := codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionSeq.EncodePreamble([]bool{c.PortSelectionSamplingSize != nil}); err != nil {
						return err
					}
					if c.PortSelectionSamplingSize != nil {
						if err := e.EncodeEnumerated((*c.PortSelectionSamplingSize), codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionPortSelectionSamplingSizeConstraints); err != nil {
							return err
						}
					}
					if err := e.EncodeBitString(c.TypeII_PortSelectionRI_Restriction, per.FixedSize(2)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.PhaseAlphabetSize, codebookConfigCodebookTypeType2PhaseAlphabetSizeConstraints); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.SubbandAmplitude); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.NumberOfBeams, codebookConfigCodebookTypeType2NumberOfBeamsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CodebookType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CodebookConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookConfigConstraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CodebookType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CodebookConfig_CodebookType_Type1:
			ie.CodebookType.Type1 = &struct {
				SubType struct {
					Choice            int
					TypeI_SinglePanel *struct {
						NrOfAntennaPorts struct {
							Choice      int
							Two         *struct{ TwoTX_CodebookSubsetRestriction per.BitString }
							MoreThanTwo *struct {
								N1_N2 struct {
									Choice                                    int
									Two_One_TypeI_SinglePanel_Restriction     *per.BitString
									Two_Two_TypeI_SinglePanel_Restriction     *per.BitString
									Four_One_TypeI_SinglePanel_Restriction    *per.BitString
									Three_Two_TypeI_SinglePanel_Restriction   *per.BitString
									Six_One_TypeI_SinglePanel_Restriction     *per.BitString
									Four_Two_TypeI_SinglePanel_Restriction    *per.BitString
									Eight_One_TypeI_SinglePanel_Restriction   *per.BitString
									Four_Three_TypeI_SinglePanel_Restriction  *per.BitString
									Six_Two_TypeI_SinglePanel_Restriction     *per.BitString
									Twelve_One_TypeI_SinglePanel_Restriction  *per.BitString
									Four_Four_TypeI_SinglePanel_Restriction   *per.BitString
									Eight_Two_TypeI_SinglePanel_Restriction   *per.BitString
									Sixteen_One_TypeI_SinglePanel_Restriction *per.BitString
								}
								TypeI_SinglePanel_CodebookSubsetRestriction_I2 *per.BitString
							}
						}
						TypeI_SinglePanel_Ri_Restriction per.BitString
					}
					TypeI_MultiPanel *struct {
						Ng_N1_N2 struct {
							Choice                                     int
							Two_Two_One_TypeI_MultiPanel_Restriction   *per.BitString
							Two_Four_One_TypeI_MultiPanel_Restriction  *per.BitString
							Four_Two_One_TypeI_MultiPanel_Restriction  *per.BitString
							Two_Two_Two_TypeI_MultiPanel_Restriction   *per.BitString
							Two_Eight_One_TypeI_MultiPanel_Restriction *per.BitString
							Four_Four_One_TypeI_MultiPanel_Restriction *per.BitString
							Two_Four_Two_TypeI_MultiPanel_Restriction  *per.BitString
							Four_Two_Two_TypeI_MultiPanel_Restriction  *per.BitString
						}
						Ri_Restriction per.BitString
					}
				}
				CodebookMode int64
			}{}
			c := (*ie.CodebookType.Type1)
			{
				choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeType1SubTypeConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.SubType.Choice = int(index)
				switch index {
				case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel:
					c.SubType.TypeI_SinglePanel = &struct {
						NrOfAntennaPorts struct {
							Choice      int
							Two         *struct{ TwoTX_CodebookSubsetRestriction per.BitString }
							MoreThanTwo *struct {
								N1_N2 struct {
									Choice                                    int
									Two_One_TypeI_SinglePanel_Restriction     *per.BitString
									Two_Two_TypeI_SinglePanel_Restriction     *per.BitString
									Four_One_TypeI_SinglePanel_Restriction    *per.BitString
									Three_Two_TypeI_SinglePanel_Restriction   *per.BitString
									Six_One_TypeI_SinglePanel_Restriction     *per.BitString
									Four_Two_TypeI_SinglePanel_Restriction    *per.BitString
									Eight_One_TypeI_SinglePanel_Restriction   *per.BitString
									Four_Three_TypeI_SinglePanel_Restriction  *per.BitString
									Six_Two_TypeI_SinglePanel_Restriction     *per.BitString
									Twelve_One_TypeI_SinglePanel_Restriction  *per.BitString
									Four_Four_TypeI_SinglePanel_Restriction   *per.BitString
									Eight_Two_TypeI_SinglePanel_Restriction   *per.BitString
									Sixteen_One_TypeI_SinglePanel_Restriction *per.BitString
								}
								TypeI_SinglePanel_CodebookSubsetRestriction_I2 *per.BitString
							}
						}
						TypeI_SinglePanel_Ri_Restriction per.BitString
					}{}
					c := (*c.SubType.TypeI_SinglePanel)
					{
						choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsConstraints)
						index, _, _, err := choiceDec.DecodeChoice()
						if err != nil {
							return err
						}
						c.NrOfAntennaPorts.Choice = int(index)
						switch index {
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_Two:
							c.NrOfAntennaPorts.Two = &struct{ TwoTX_CodebookSubsetRestriction per.BitString }{}
							c := (*c.NrOfAntennaPorts.Two)
							{
								v, err := d.DecodeBitString(per.FixedSize(6))
								if err != nil {
									return err
								}
								c.TwoTX_CodebookSubsetRestriction = v
							}
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo:
							c.NrOfAntennaPorts.MoreThanTwo = &struct {
								N1_N2 struct {
									Choice                                    int
									Two_One_TypeI_SinglePanel_Restriction     *per.BitString
									Two_Two_TypeI_SinglePanel_Restriction     *per.BitString
									Four_One_TypeI_SinglePanel_Restriction    *per.BitString
									Three_Two_TypeI_SinglePanel_Restriction   *per.BitString
									Six_One_TypeI_SinglePanel_Restriction     *per.BitString
									Four_Two_TypeI_SinglePanel_Restriction    *per.BitString
									Eight_One_TypeI_SinglePanel_Restriction   *per.BitString
									Four_Three_TypeI_SinglePanel_Restriction  *per.BitString
									Six_Two_TypeI_SinglePanel_Restriction     *per.BitString
									Twelve_One_TypeI_SinglePanel_Restriction  *per.BitString
									Four_Four_TypeI_SinglePanel_Restriction   *per.BitString
									Eight_Two_TypeI_SinglePanel_Restriction   *per.BitString
									Sixteen_One_TypeI_SinglePanel_Restriction *per.BitString
								}
								TypeI_SinglePanel_CodebookSubsetRestriction_I2 *per.BitString
							}{}
							c := (*c.NrOfAntennaPorts.MoreThanTwo)
							codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoSeq := d.NewSequenceDecoder(codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoConstraints)
							if err := codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoSeq.DecodePreamble(); err != nil {
								return err
							}
							{
								choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoN1N2Constraints)
								index, _, _, err := choiceDec.DecodeChoice()
								if err != nil {
									return err
								}
								c.N1_N2.Choice = int(index)
								switch index {
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction:
									c.N1_N2.Two_One_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(8))
									if err != nil {
										return err
									}
									(*c.N1_N2.Two_One_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction:
									c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(64))
									if err != nil {
										return err
									}
									(*c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction:
									c.N1_N2.Four_One_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(16))
									if err != nil {
										return err
									}
									(*c.N1_N2.Four_One_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction:
									c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(96))
									if err != nil {
										return err
									}
									(*c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction:
									c.N1_N2.Six_One_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(24))
									if err != nil {
										return err
									}
									(*c.N1_N2.Six_One_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction:
									c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(128))
									if err != nil {
										return err
									}
									(*c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction:
									c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(32))
									if err != nil {
										return err
									}
									(*c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction:
									c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(192))
									if err != nil {
										return err
									}
									(*c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction:
									c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(192))
									if err != nil {
										return err
									}
									(*c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction:
									c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(48))
									if err != nil {
										return err
									}
									(*c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction:
									c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(256))
									if err != nil {
										return err
									}
									(*c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction:
									c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(256))
									if err != nil {
										return err
									}
									(*c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction) = v
								case CodebookConfig_CodebookType_Type1_SubType_TypeI_SinglePanel_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction:
									c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction = new(per.BitString)
									v, err := d.DecodeBitString(per.FixedSize(64))
									if err != nil {
										return err
									}
									(*c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction) = v
								}
							}
							if codebookConfigCodebookTypeType1SubTypeTypeISinglePanelNrOfAntennaPortsMoreThanTwoSeq.IsComponentPresent(1) {
								c.TypeI_SinglePanel_CodebookSubsetRestriction_I2 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(16))
								if err != nil {
									return err
								}
								(*c.TypeI_SinglePanel_CodebookSubsetRestriction_I2) = v
							}
						}
					}
					{
						v, err := d.DecodeBitString(per.FixedSize(8))
						if err != nil {
							return err
						}
						c.TypeI_SinglePanel_Ri_Restriction = v
					}
				case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel:
					c.SubType.TypeI_MultiPanel = &struct {
						Ng_N1_N2 struct {
							Choice                                     int
							Two_Two_One_TypeI_MultiPanel_Restriction   *per.BitString
							Two_Four_One_TypeI_MultiPanel_Restriction  *per.BitString
							Four_Two_One_TypeI_MultiPanel_Restriction  *per.BitString
							Two_Two_Two_TypeI_MultiPanel_Restriction   *per.BitString
							Two_Eight_One_TypeI_MultiPanel_Restriction *per.BitString
							Four_Four_One_TypeI_MultiPanel_Restriction *per.BitString
							Two_Four_Two_TypeI_MultiPanel_Restriction  *per.BitString
							Four_Two_Two_TypeI_MultiPanel_Restriction  *per.BitString
						}
						Ri_Restriction per.BitString
					}{}
					c := (*c.SubType.TypeI_MultiPanel)
					{
						choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeType1SubTypeTypeIMultiPanelNgN1N2Constraints)
						index, _, _, err := choiceDec.DecodeChoice()
						if err != nil {
							return err
						}
						c.Ng_N1_N2.Choice = int(index)
						switch index {
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Two_One_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Two_Two_One_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(8))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Two_Two_One_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Four_One_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Two_Four_One_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(16))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Two_Four_One_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Two_One_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Four_Two_One_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(8))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Four_Two_One_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Two_Two_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Two_Two_Two_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(64))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Two_Two_Two_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Eight_One_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Two_Eight_One_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(32))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Two_Eight_One_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Four_One_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Four_Four_One_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(16))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Four_Four_One_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Two_Four_Two_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Two_Four_Two_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(128))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Two_Four_Two_TypeI_MultiPanel_Restriction) = v
						case CodebookConfig_CodebookType_Type1_SubType_TypeI_MultiPanel_Ng_N1_N2_Four_Two_Two_TypeI_MultiPanel_Restriction:
							c.Ng_N1_N2.Four_Two_Two_TypeI_MultiPanel_Restriction = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(64))
							if err != nil {
								return err
							}
							(*c.Ng_N1_N2.Four_Two_Two_TypeI_MultiPanel_Restriction) = v
						}
					}
					{
						v, err := d.DecodeBitString(per.FixedSize(4))
						if err != nil {
							return err
						}
						c.Ri_Restriction = v
					}
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.CodebookMode = v
			}
		case CodebookConfig_CodebookType_Type2:
			ie.CodebookType.Type2 = &struct {
				SubType struct {
					Choice int
					TypeII *struct {
						N1_N2_CodebookSubsetRestriction struct {
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
						TypeII_RI_Restriction per.BitString
					}
					TypeII_PortSelection *struct {
						PortSelectionSamplingSize          *int64
						TypeII_PortSelectionRI_Restriction per.BitString
					}
				}
				PhaseAlphabetSize int64
				SubbandAmplitude  bool
				NumberOfBeams     int64
			}{}
			c := (*ie.CodebookType.Type2)
			{
				choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeType2SubTypeConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.SubType.Choice = int(index)
				switch index {
				case CodebookConfig_CodebookType_Type2_SubType_TypeII:
					c.SubType.TypeII = &struct {
						N1_N2_CodebookSubsetRestriction struct {
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
						TypeII_RI_Restriction per.BitString
					}{}
					c := (*c.SubType.TypeII)
					{
						choiceDec := d.NewChoiceDecoder(codebookConfigCodebookTypeType2SubTypeTypeIIN1N2CodebookSubsetRestrictionConstraints)
						index, _, _, err := choiceDec.DecodeChoice()
						if err != nil {
							return err
						}
						c.N1_N2_CodebookSubsetRestriction.Choice = int(index)
						switch index {
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Two_One:
							c.N1_N2_CodebookSubsetRestriction.Two_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(16))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Two_One) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Two_Two:
							c.N1_N2_CodebookSubsetRestriction.Two_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(43))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Two_Two) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_One:
							c.N1_N2_CodebookSubsetRestriction.Four_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(32))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Four_One) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Three_Two:
							c.N1_N2_CodebookSubsetRestriction.Three_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(59))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Three_Two) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Six_One:
							c.N1_N2_CodebookSubsetRestriction.Six_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(48))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Six_One) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Two:
							c.N1_N2_CodebookSubsetRestriction.Four_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(75))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Four_Two) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Eight_One:
							c.N1_N2_CodebookSubsetRestriction.Eight_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(64))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Eight_One) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Three:
							c.N1_N2_CodebookSubsetRestriction.Four_Three = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(107))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Four_Three) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Six_Two:
							c.N1_N2_CodebookSubsetRestriction.Six_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(107))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Six_Two) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Twelve_One:
							c.N1_N2_CodebookSubsetRestriction.Twelve_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(96))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Twelve_One) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Four_Four:
							c.N1_N2_CodebookSubsetRestriction.Four_Four = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(139))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Four_Four) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Eight_Two:
							c.N1_N2_CodebookSubsetRestriction.Eight_Two = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(139))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Eight_Two) = v
						case CodebookConfig_CodebookType_Type2_SubType_TypeII_N1_N2_CodebookSubsetRestriction_Sixteen_One:
							c.N1_N2_CodebookSubsetRestriction.Sixteen_One = new(per.BitString)
							v, err := d.DecodeBitString(per.FixedSize(128))
							if err != nil {
								return err
							}
							(*c.N1_N2_CodebookSubsetRestriction.Sixteen_One) = v
						}
					}
					{
						v, err := d.DecodeBitString(per.FixedSize(2))
						if err != nil {
							return err
						}
						c.TypeII_RI_Restriction = v
					}
				case CodebookConfig_CodebookType_Type2_SubType_TypeII_PortSelection:
					c.SubType.TypeII_PortSelection = &struct {
						PortSelectionSamplingSize          *int64
						TypeII_PortSelectionRI_Restriction per.BitString
					}{}
					c := (*c.SubType.TypeII_PortSelection)
					codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionSeq := d.NewSequenceDecoder(codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionConstraints)
					if err := codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionSeq.DecodePreamble(); err != nil {
						return err
					}
					if codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionSeq.IsComponentPresent(0) {
						c.PortSelectionSamplingSize = new(int64)
						v, err := d.DecodeEnumerated(codebookConfigCodebookTypeType2SubTypeTypeIIPortSelectionPortSelectionSamplingSizeConstraints)
						if err != nil {
							return err
						}
						(*c.PortSelectionSamplingSize) = v
					}
					{
						v, err := d.DecodeBitString(per.FixedSize(2))
						if err != nil {
							return err
						}
						c.TypeII_PortSelectionRI_Restriction = v
					}
				}
			}
			{
				v, err := d.DecodeEnumerated(codebookConfigCodebookTypeType2PhaseAlphabetSizeConstraints)
				if err != nil {
					return err
				}
				c.PhaseAlphabetSize = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.SubbandAmplitude = v
			}
			{
				v, err := d.DecodeEnumerated(codebookConfigCodebookTypeType2NumberOfBeamsConstraints)
				if err != nil {
					return err
				}
				c.NumberOfBeams = v
			}
		}
	}

	return nil
}
