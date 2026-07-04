// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookConfig-r17 (line 6157).

var codebookConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookType"},
	},
}

var codebookConfig_r17CodebookTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1"},
		{Name: "type2"},
	},
}

const (
	CodebookConfig_r17_CodebookType_Type1 = 0
	CodebookConfig_r17_CodebookType_Type2 = 1
)

var codebookConfigR17CodebookTypeType1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "typeI-SinglePanel-Group1-r17", Optional: true},
		{Name: "typeI-SinglePanel-Group2-r17", Optional: true},
		{Name: "typeI-SinglePanel-ri-RestrictionSTRP-r17", Optional: true},
		{Name: "typeI-SinglePanel-ri-RestrictionSDM-r17", Optional: true},
	},
}

var codebookConfigR17CodebookTypeType1TypeISinglePanelGroup1R17NrOfAntennaPortsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two"},
		{Name: "moreThanTwo"},
	},
}

const (
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_Two         = 0
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo = 1
)

var codebookConfigR17CodebookTypeType1TypeISinglePanelGroup1R17NrOfAntennaPortsMoreThanTwoN1N2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-one-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "two-two-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "four-one-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "three-two-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "six-one-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "four-two-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "eight-one-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "four-three-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "six-two-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "twelve-one-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "four-four-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "eight-two-TypeI-SinglePanel-Restriction1-r17"},
		{Name: "sixteen-one-TypeI-SinglePanel-Restriction1-r17"},
	},
}

const (
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction1_r17     = 0
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction1_r17     = 1
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction1_r17    = 2
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction1_r17   = 3
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction1_r17     = 4
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction1_r17    = 5
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction1_r17   = 6
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction1_r17  = 7
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction1_r17     = 8
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction1_r17  = 9
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction1_r17   = 10
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction1_r17   = 11
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction1_r17 = 12
)

var codebookConfigR17CodebookTypeType1TypeISinglePanelGroup2R17NrOfAntennaPortsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two"},
		{Name: "moreThanTwo"},
	},
}

const (
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_Two         = 0
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo = 1
)

var codebookConfigR17CodebookTypeType1TypeISinglePanelGroup2R17NrOfAntennaPortsMoreThanTwoN1N2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-one-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "two-two-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "four-one-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "three-two-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "six-one-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "four-two-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "eight-one-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "four-three-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "six-two-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "twelve-one-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "four-four-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "eight-two-TypeI-SinglePanel-Restriction2-r17"},
		{Name: "sixteen-one-TypeI-SinglePanel-Restriction2-r17"},
	},
}

const (
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction2_r17     = 0
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction2_r17     = 1
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction2_r17    = 2
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction2_r17   = 3
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction2_r17     = 4
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction2_r17    = 5
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction2_r17   = 6
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction2_r17  = 7
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction2_r17     = 8
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction2_r17  = 9
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction2_r17   = 10
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction2_r17   = 11
	CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction2_r17 = 12
)

var codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "paramCombination-r17"},
		{Name: "valueOfN-r17", Optional: true},
		{Name: "numberOfPMI-SubbandsPerCQI-Subband-r17", Optional: true},
		{Name: "typeII-PortSelectionRI-Restriction-r17"},
	},
}

const (
	CodebookConfig_r17_CodebookType_Type2_TypeII_PortSelection_r17_ValueOfN_r17_N2 = 0
	CodebookConfig_r17_CodebookType_Type2_TypeII_PortSelection_r17_ValueOfN_r17_N4 = 1
)

var codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17ValueOfNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r17_CodebookType_Type2_TypeII_PortSelection_r17_ValueOfN_r17_N2, CodebookConfig_r17_CodebookType_Type2_TypeII_PortSelection_r17_ValueOfN_r17_N4},
}

type CodebookConfig_r17 struct {
	CodebookType struct {
		Choice int
		Type1  *struct {
			TypeI_SinglePanel_Group1_r17 *struct {
				NrOfAntennaPorts struct {
					Choice      int
					Two         *struct{ TwoTX_CodebookSubsetRestriction1_r17 per.BitString }
					MoreThanTwo *struct {
						N1_N2 struct {
							Choice                                         int
							Two_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
							Two_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
							Four_One_TypeI_SinglePanel_Restriction1_r17    *per.BitString
							Three_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
							Six_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
							Four_Two_TypeI_SinglePanel_Restriction1_r17    *per.BitString
							Eight_One_TypeI_SinglePanel_Restriction1_r17   *per.BitString
							Four_Three_TypeI_SinglePanel_Restriction1_r17  *per.BitString
							Six_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
							Twelve_One_TypeI_SinglePanel_Restriction1_r17  *per.BitString
							Four_Four_TypeI_SinglePanel_Restriction1_r17   *per.BitString
							Eight_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
							Sixteen_One_TypeI_SinglePanel_Restriction1_r17 *per.BitString
						}
					}
				}
			}
			TypeI_SinglePanel_Group2_r17 *struct {
				NrOfAntennaPorts struct {
					Choice      int
					Two         *struct{ TwoTX_CodebookSubsetRestriction2_r17 per.BitString }
					MoreThanTwo *struct {
						N1_N2 struct {
							Choice                                         int
							Two_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
							Two_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
							Four_One_TypeI_SinglePanel_Restriction2_r17    *per.BitString
							Three_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
							Six_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
							Four_Two_TypeI_SinglePanel_Restriction2_r17    *per.BitString
							Eight_One_TypeI_SinglePanel_Restriction2_r17   *per.BitString
							Four_Three_TypeI_SinglePanel_Restriction2_r17  *per.BitString
							Six_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
							Twelve_One_TypeI_SinglePanel_Restriction2_r17  *per.BitString
							Four_Four_TypeI_SinglePanel_Restriction2_r17   *per.BitString
							Eight_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
							Sixteen_One_TypeI_SinglePanel_Restriction2_r17 *per.BitString
						}
					}
				}
			}
			TypeI_SinglePanel_Ri_RestrictionSTRP_r17 *per.BitString
			TypeI_SinglePanel_Ri_RestrictionSDM_r17  *per.BitString
		}
		Type2 *struct {
			TypeII_PortSelection_r17 struct {
				ParamCombination_r17                   int64
				ValueOfN_r17                           *int64
				NumberOfPMI_SubbandsPerCQI_Subband_r17 *int64
				TypeII_PortSelectionRI_Restriction_r17 per.BitString
			}
		}
	}
}

func (ie *CodebookConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookConfigR17Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceEnc := e.NewChoiceEncoder(codebookConfig_r17CodebookTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CodebookType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CodebookType.Choice {
		case CodebookConfig_r17_CodebookType_Type1:
			c := (*ie.CodebookType.Type1)
			codebookConfigR17CodebookTypeType1Seq := e.NewSequenceEncoder(codebookConfigR17CodebookTypeType1Constraints)
			if err := codebookConfigR17CodebookTypeType1Seq.EncodePreamble([]bool{c.TypeI_SinglePanel_Group1_r17 != nil, c.TypeI_SinglePanel_Group2_r17 != nil, c.TypeI_SinglePanel_Ri_RestrictionSTRP_r17 != nil, c.TypeI_SinglePanel_Ri_RestrictionSDM_r17 != nil}); err != nil {
				return err
			}
			if c.TypeI_SinglePanel_Group1_r17 != nil {
				c := (*c.TypeI_SinglePanel_Group1_r17)
				{
					choiceEnc := e.NewChoiceEncoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup1R17NrOfAntennaPortsConstraints)
					if err := choiceEnc.EncodeChoice(int64(c.NrOfAntennaPorts.Choice), false, nil); err != nil {
						return err
					}
					switch c.NrOfAntennaPorts.Choice {
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_Two:
						c := (*c.NrOfAntennaPorts.Two)
						if err := e.EncodeBitString(c.TwoTX_CodebookSubsetRestriction1_r17, per.FixedSize(6)); err != nil {
							return err
						}
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo:
						c := (*c.NrOfAntennaPorts.MoreThanTwo)
						{
							choiceEnc := e.NewChoiceEncoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup1R17NrOfAntennaPortsMoreThanTwoN1N2Constraints)
							if err := choiceEnc.EncodeChoice(int64(c.N1_N2.Choice), false, nil); err != nil {
								return err
							}
							switch c.N1_N2.Choice {
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Two_One_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(8)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(64)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_One_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(16)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(96)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Six_One_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(24)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(128)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(32)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(192)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(192)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(48)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(256)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(256)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction1_r17:
								if err := e.EncodeBitString((*c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction1_r17), per.FixedSize(64)); err != nil {
									return err
								}
							}
						}
					}
				}
			}
			if c.TypeI_SinglePanel_Group2_r17 != nil {
				c := (*c.TypeI_SinglePanel_Group2_r17)
				{
					choiceEnc := e.NewChoiceEncoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup2R17NrOfAntennaPortsConstraints)
					if err := choiceEnc.EncodeChoice(int64(c.NrOfAntennaPorts.Choice), false, nil); err != nil {
						return err
					}
					switch c.NrOfAntennaPorts.Choice {
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_Two:
						c := (*c.NrOfAntennaPorts.Two)
						if err := e.EncodeBitString(c.TwoTX_CodebookSubsetRestriction2_r17, per.FixedSize(6)); err != nil {
							return err
						}
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo:
						c := (*c.NrOfAntennaPorts.MoreThanTwo)
						{
							choiceEnc := e.NewChoiceEncoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup2R17NrOfAntennaPortsMoreThanTwoN1N2Constraints)
							if err := choiceEnc.EncodeChoice(int64(c.N1_N2.Choice), false, nil); err != nil {
								return err
							}
							switch c.N1_N2.Choice {
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Two_One_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(8)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(64)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_One_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(16)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(96)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Six_One_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(24)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(128)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(32)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(192)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(192)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(48)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(256)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(256)); err != nil {
									return err
								}
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction2_r17:
								if err := e.EncodeBitString((*c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction2_r17), per.FixedSize(64)); err != nil {
									return err
								}
							}
						}
					}
				}
			}
			if c.TypeI_SinglePanel_Ri_RestrictionSTRP_r17 != nil {
				if err := e.EncodeBitString((*c.TypeI_SinglePanel_Ri_RestrictionSTRP_r17), per.FixedSize(8)); err != nil {
					return err
				}
			}
			if c.TypeI_SinglePanel_Ri_RestrictionSDM_r17 != nil {
				if err := e.EncodeBitString((*c.TypeI_SinglePanel_Ri_RestrictionSDM_r17), per.FixedSize(4)); err != nil {
					return err
				}
			}
		case CodebookConfig_r17_CodebookType_Type2:
			c := (*ie.CodebookType.Type2)
			{
				c := &c.TypeII_PortSelection_r17
				codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Seq := e.NewSequenceEncoder(codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Constraints)
				if err := codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Seq.EncodePreamble([]bool{c.ValueOfN_r17 != nil, c.NumberOfPMI_SubbandsPerCQI_Subband_r17 != nil}); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.ParamCombination_r17, per.Constrained(1, 8)); err != nil {
					return err
				}
				if c.ValueOfN_r17 != nil {
					if err := e.EncodeEnumerated((*c.ValueOfN_r17), codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17ValueOfNR17Constraints); err != nil {
						return err
					}
				}
				if c.NumberOfPMI_SubbandsPerCQI_Subband_r17 != nil {
					if err := e.EncodeInteger((*c.NumberOfPMI_SubbandsPerCQI_Subband_r17), per.Constrained(1, 2)); err != nil {
						return err
					}
				}
				if err := e.EncodeBitString(c.TypeII_PortSelectionRI_Restriction_r17, per.FixedSize(4)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CodebookType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CodebookConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookConfigR17Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceDec := d.NewChoiceDecoder(codebookConfig_r17CodebookTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CodebookType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CodebookConfig_r17_CodebookType_Type1:
			ie.CodebookType.Type1 = &struct {
				TypeI_SinglePanel_Group1_r17 *struct {
					NrOfAntennaPorts struct {
						Choice      int
						Two         *struct{ TwoTX_CodebookSubsetRestriction1_r17 per.BitString }
						MoreThanTwo *struct {
							N1_N2 struct {
								Choice                                         int
								Two_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction1_r17    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction1_r17    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction1_r17  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction1_r17  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction1_r17 *per.BitString
							}
						}
					}
				}
				TypeI_SinglePanel_Group2_r17 *struct {
					NrOfAntennaPorts struct {
						Choice      int
						Two         *struct{ TwoTX_CodebookSubsetRestriction2_r17 per.BitString }
						MoreThanTwo *struct {
							N1_N2 struct {
								Choice                                         int
								Two_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction2_r17    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction2_r17    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction2_r17  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction2_r17  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction2_r17 *per.BitString
							}
						}
					}
				}
				TypeI_SinglePanel_Ri_RestrictionSTRP_r17 *per.BitString
				TypeI_SinglePanel_Ri_RestrictionSDM_r17  *per.BitString
			}{}
			c := (*ie.CodebookType.Type1)
			codebookConfigR17CodebookTypeType1Seq := d.NewSequenceDecoder(codebookConfigR17CodebookTypeType1Constraints)
			if err := codebookConfigR17CodebookTypeType1Seq.DecodePreamble(); err != nil {
				return err
			}
			if codebookConfigR17CodebookTypeType1Seq.IsComponentPresent(0) {
				c.TypeI_SinglePanel_Group1_r17 = &struct {
					NrOfAntennaPorts struct {
						Choice      int
						Two         *struct{ TwoTX_CodebookSubsetRestriction1_r17 per.BitString }
						MoreThanTwo *struct {
							N1_N2 struct {
								Choice                                         int
								Two_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction1_r17    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction1_r17    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction1_r17  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction1_r17  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction1_r17 *per.BitString
							}
						}
					}
				}{}
				c := (*c.TypeI_SinglePanel_Group1_r17)
				{
					choiceDec := d.NewChoiceDecoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup1R17NrOfAntennaPortsConstraints)
					index, _, _, err := choiceDec.DecodeChoice()
					if err != nil {
						return err
					}
					c.NrOfAntennaPorts.Choice = int(index)
					switch index {
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_Two:
						c.NrOfAntennaPorts.Two = &struct{ TwoTX_CodebookSubsetRestriction1_r17 per.BitString }{}
						c := (*c.NrOfAntennaPorts.Two)
						{
							v, err := d.DecodeBitString(per.FixedSize(6))
							if err != nil {
								return err
							}
							c.TwoTX_CodebookSubsetRestriction1_r17 = v
						}
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo:
						c.NrOfAntennaPorts.MoreThanTwo = &struct {
							N1_N2 struct {
								Choice                                         int
								Two_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction1_r17    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction1_r17    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction1_r17  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction1_r17     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction1_r17  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction1_r17   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction1_r17 *per.BitString
							}
						}{}
						c := (*c.NrOfAntennaPorts.MoreThanTwo)
						{
							choiceDec := d.NewChoiceDecoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup1R17NrOfAntennaPortsMoreThanTwoN1N2Constraints)
							index, _, _, err := choiceDec.DecodeChoice()
							if err != nil {
								return err
							}
							c.N1_N2.Choice = int(index)
							switch index {
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Two_One_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(8))
								if err != nil {
									return err
								}
								(*c.N1_N2.Two_One_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(64))
								if err != nil {
									return err
								}
								(*c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Four_One_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(16))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_One_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(96))
								if err != nil {
									return err
								}
								(*c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Six_One_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(24))
								if err != nil {
									return err
								}
								(*c.N1_N2.Six_One_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(128))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(32))
								if err != nil {
									return err
								}
								(*c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(192))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(192))
								if err != nil {
									return err
								}
								(*c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(48))
								if err != nil {
									return err
								}
								(*c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(256))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(256))
								if err != nil {
									return err
								}
								(*c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction1_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group1_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction1_r17:
								c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction1_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(64))
								if err != nil {
									return err
								}
								(*c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction1_r17) = v
							}
						}
					}
				}
			}
			if codebookConfigR17CodebookTypeType1Seq.IsComponentPresent(1) {
				c.TypeI_SinglePanel_Group2_r17 = &struct {
					NrOfAntennaPorts struct {
						Choice      int
						Two         *struct{ TwoTX_CodebookSubsetRestriction2_r17 per.BitString }
						MoreThanTwo *struct {
							N1_N2 struct {
								Choice                                         int
								Two_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction2_r17    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction2_r17    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction2_r17  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction2_r17  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction2_r17 *per.BitString
							}
						}
					}
				}{}
				c := (*c.TypeI_SinglePanel_Group2_r17)
				{
					choiceDec := d.NewChoiceDecoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup2R17NrOfAntennaPortsConstraints)
					index, _, _, err := choiceDec.DecodeChoice()
					if err != nil {
						return err
					}
					c.NrOfAntennaPorts.Choice = int(index)
					switch index {
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_Two:
						c.NrOfAntennaPorts.Two = &struct{ TwoTX_CodebookSubsetRestriction2_r17 per.BitString }{}
						c := (*c.NrOfAntennaPorts.Two)
						{
							v, err := d.DecodeBitString(per.FixedSize(6))
							if err != nil {
								return err
							}
							c.TwoTX_CodebookSubsetRestriction2_r17 = v
						}
					case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo:
						c.NrOfAntennaPorts.MoreThanTwo = &struct {
							N1_N2 struct {
								Choice                                         int
								Two_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Two_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Four_One_TypeI_SinglePanel_Restriction2_r17    *per.BitString
								Three_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Six_One_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Four_Two_TypeI_SinglePanel_Restriction2_r17    *per.BitString
								Eight_One_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Four_Three_TypeI_SinglePanel_Restriction2_r17  *per.BitString
								Six_Two_TypeI_SinglePanel_Restriction2_r17     *per.BitString
								Twelve_One_TypeI_SinglePanel_Restriction2_r17  *per.BitString
								Four_Four_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Eight_Two_TypeI_SinglePanel_Restriction2_r17   *per.BitString
								Sixteen_One_TypeI_SinglePanel_Restriction2_r17 *per.BitString
							}
						}{}
						c := (*c.NrOfAntennaPorts.MoreThanTwo)
						{
							choiceDec := d.NewChoiceDecoder(codebookConfigR17CodebookTypeType1TypeISinglePanelGroup2R17NrOfAntennaPortsMoreThanTwoN1N2Constraints)
							index, _, _, err := choiceDec.DecodeChoice()
							if err != nil {
								return err
							}
							c.N1_N2.Choice = int(index)
							switch index {
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_One_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Two_One_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(8))
								if err != nil {
									return err
								}
								(*c.N1_N2.Two_One_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Two_Two_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(64))
								if err != nil {
									return err
								}
								(*c.N1_N2.Two_Two_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_One_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Four_One_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(16))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_One_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Three_Two_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(96))
								if err != nil {
									return err
								}
								(*c.N1_N2.Three_Two_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_One_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Six_One_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(24))
								if err != nil {
									return err
								}
								(*c.N1_N2.Six_One_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Two_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(128))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_Two_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_One_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(32))
								if err != nil {
									return err
								}
								(*c.N1_N2.Eight_One_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Three_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(192))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_Three_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Six_Two_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(192))
								if err != nil {
									return err
								}
								(*c.N1_N2.Six_Two_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Twelve_One_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(48))
								if err != nil {
									return err
								}
								(*c.N1_N2.Twelve_One_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Four_Four_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(256))
								if err != nil {
									return err
								}
								(*c.N1_N2.Four_Four_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Eight_Two_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(256))
								if err != nil {
									return err
								}
								(*c.N1_N2.Eight_Two_TypeI_SinglePanel_Restriction2_r17) = v
							case CodebookConfig_r17_CodebookType_Type1_TypeI_SinglePanel_Group2_r17_NrOfAntennaPorts_MoreThanTwo_N1_N2_Sixteen_One_TypeI_SinglePanel_Restriction2_r17:
								c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction2_r17 = new(per.BitString)
								v, err := d.DecodeBitString(per.FixedSize(64))
								if err != nil {
									return err
								}
								(*c.N1_N2.Sixteen_One_TypeI_SinglePanel_Restriction2_r17) = v
							}
						}
					}
				}
			}
			if codebookConfigR17CodebookTypeType1Seq.IsComponentPresent(2) {
				c.TypeI_SinglePanel_Ri_RestrictionSTRP_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*c.TypeI_SinglePanel_Ri_RestrictionSTRP_r17) = v
			}
			if codebookConfigR17CodebookTypeType1Seq.IsComponentPresent(3) {
				c.TypeI_SinglePanel_Ri_RestrictionSDM_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*c.TypeI_SinglePanel_Ri_RestrictionSDM_r17) = v
			}
		case CodebookConfig_r17_CodebookType_Type2:
			ie.CodebookType.Type2 = &struct {
				TypeII_PortSelection_r17 struct {
					ParamCombination_r17                   int64
					ValueOfN_r17                           *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r17 *int64
					TypeII_PortSelectionRI_Restriction_r17 per.BitString
				}
			}{}
			c := (*ie.CodebookType.Type2)
			{
				c := &c.TypeII_PortSelection_r17
				codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Seq := d.NewSequenceDecoder(codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Constraints)
				if err := codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 8))
					if err != nil {
						return err
					}
					c.ParamCombination_r17 = v
				}
				if codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Seq.IsComponentPresent(1) {
					c.ValueOfN_r17 = new(int64)
					v, err := d.DecodeEnumerated(codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17ValueOfNR17Constraints)
					if err != nil {
						return err
					}
					(*c.ValueOfN_r17) = v
				}
				if codebookConfigR17CodebookTypeType2TypeIIPortSelectionR17Seq.IsComponentPresent(2) {
					c.NumberOfPMI_SubbandsPerCQI_Subband_r17 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					(*c.NumberOfPMI_SubbandsPerCQI_Subband_r17) = v
				}
				{
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					c.TypeII_PortSelectionRI_Restriction_r17 = v
				}
			}
		}
	}

	return nil
}
