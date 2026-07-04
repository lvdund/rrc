// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookConfig-r19 (line 6273).

var codebookConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "codebookType"},
		{Name: "cri-TypeI-SinglePanel"},
		{Name: "cri-TypeII"},
	},
}

const (
	CodebookConfig_r19_CodebookType          = 0
	CodebookConfig_r19_Cri_TypeI_SinglePanel = 1
	CodebookConfig_r19_Cri_TypeII            = 2
)

var codebookConfigR19CodebookTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1"},
		{Name: "type2"},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type1 = 0
	CodebookConfig_r19_CodebookType_Type2 = 1
)

var codebookConfigR19CodebookTypeType1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "typeI-SinglePanel-r19"},
		{Name: "typeI-MultiPanel-r19"},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19 = 0
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19  = 1
)

var codebookConfigR19CodebookTypeType1TypeISinglePanelR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookMode-r19"},
		{Name: "typeI-SinglePanelRI-Restriction-r19"},
		{Name: "n1-n2-r19"},
		{Name: "typeI-CodebookSubsetRestriction-r19", Optional: true},
		{Name: "typeI-SoftScalingRank-r19", Optional: true},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_CodebookMode_r19_ModeA = 0
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_CodebookMode_r19_ModeB = 1
)

var codebookConfigR19CodebookTypeType1TypeISinglePanelR19CodebookModeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_CodebookMode_r19_ModeA, CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_CodebookMode_r19_ModeB},
}

const (
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Eight_Three  = 0
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Six_Four     = 1
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Sixteen_Two  = 2
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Eight_Four   = 3
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Sixteen_Four = 4
	CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Eight_Eight  = 5
)

var codebookConfigR19CodebookTypeType1TypeISinglePanelR19N1N2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Eight_Three, CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Six_Four, CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Sixteen_Two, CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Eight_Four, CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Sixteen_Four, CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19_N1_N2_r19_Eight_Eight},
}

var codebookConfigR19CodebookTypeType1TypeIMultiPanelR19NgN1N2CbsrR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-four-three-r19"},
		{Name: "two-six-two-r19"},
		{Name: "two-eight-two-r19"},
		{Name: "four-four-two-r19"},
		{Name: "two-four-four-r19"},
		{Name: "four-four-four-r19"},
		{Name: "four-eight-two-r19"},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Four_Three_r19 = 0
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Six_Two_r19    = 1
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Eight_Two_r19  = 2
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Four_Two_r19  = 3
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Four_Four_r19  = 4
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Four_Four_r19 = 5
	CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Eight_Two_r19 = 6
)

var codebookConfigR19CodebookTypeType2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "typeII"},
		{Name: "typeII-RI-Restriction-r19"},
		{Name: "numberOfPMI-SubbandsPerCQI-Subband-r19", Optional: true},
	},
}

var codebookConfigR19CodebookTypeType2TypeIIConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "etypeII-r19"},
		{Name: "typeII-FePortSelection-r19"},
		{Name: "typeII-Doppler-r19"},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19                = 0
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19 = 1
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19         = 2
)

var codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "paramCombination-r19"},
		{Name: "n1-n2-r19"},
		{Name: "typeII-CodebookSubsetRestriction-r19", Optional: true},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Eight_Three  = 0
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Six_Four     = 1
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Sixteen_Two  = 2
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Eight_Four   = 3
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Sixteen_Four = 4
	CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Eight_Eight  = 5
)

var codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19N1N2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Eight_Three, CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Six_Four, CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Sixteen_Two, CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Eight_Four, CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Sixteen_Four, CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19_N1_N2_r19_Eight_Eight},
}

var codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "paramCombination-r19"},
		{Name: "valueOfN-r19", Optional: true},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19_ValueOfN_r19_N2 = 0
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19_ValueOfN_r19_N4 = 1
)

var codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19ValueOfNR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19_ValueOfN_r19_N2, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19_ValueOfN_r19_N4},
}

var codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "paramCombination-Doppler-r19"},
		{Name: "td-dd-config-r19"},
		{Name: "predictionDelay-r19"},
		{Name: "n1-n2-r19"},
		{Name: "typeII-CodebookSubsetRestriction-r19", Optional: true},
	},
}

const (
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_M0 = 0
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_N0 = 1
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_N1 = 2
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_N2 = 3
)

var codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19PredictionDelayR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_M0, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_N0, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_N1, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_PredictionDelay_r19_N2},
}

const (
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Eight_Three  = 0
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Six_Four     = 1
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Sixteen_Two  = 2
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Eight_Four   = 3
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Sixteen_Four = 4
	CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Eight_Eight  = 5
)

var codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19N1N2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Eight_Three, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Six_Four, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Sixteen_Two, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Eight_Four, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Sixteen_Four, CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19_N1_N2_r19_Eight_Eight},
}

var codebookConfigR19CriTypeISinglePanelConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cri-TypeI-SinglePanelRI-Restriction-r19", Optional: true},
		{Name: "cri-TypeI-SinglePanelN1-N2-CBSR-r19", Optional: true},
	},
}

var codebookConfigR19CriTypeISinglePanelCriTypeISinglePanelRIRestrictionR19Constraints = per.SizeRange(1, 8)

var codebookConfigR19CriTypeIIConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cri-TypeII-RI-Restriction-r19", Optional: true},
		{Name: "cri-TypeII-N1-N2-CBSR-r19", Optional: true},
	},
}

var codebookConfigR19CriTypeIICriTypeIIRIRestrictionR19Constraints = per.SizeRange(1, 4)

type CodebookConfig_r19 struct {
	Choice       int
	CodebookType *struct {
		Choice int
		Type1  *struct {
			Choice                int
			TypeI_SinglePanel_r19 *struct {
				CodebookMode_r19                    int64
				TypeI_SinglePanelRI_Restriction_r19 per.BitString
				N1_N2_r19                           int64
				TypeI_CodebookSubsetRestriction_r19 *TypeI_X1_X2_CBSR_r19
				TypeI_SoftScalingRank_r19           *TypeI_X1_X2_SoftScalingRank_r19
			}
			TypeI_MultiPanel_r19 *struct {
				Ri_Restriction_r19 per.BitString
				Ng_N1_N2_Cbsr_r19  struct {
					Choice             int
					Two_Four_Three_r19 *per.BitString
					Two_Six_Two_r19    *per.BitString
					Two_Eight_Two_r19  *per.BitString
					Four_Four_Two_r19  *per.BitString
					Two_Four_Four_r19  *per.BitString
					Four_Four_Four_r19 *per.BitString
					Four_Eight_Two_r19 *per.BitString
				}
			}
		}
		Type2 *struct {
			TypeII struct {
				Choice      int
				EtypeII_r19 *struct {
					ParamCombination_r19                 int64
					N1_N2_r19                            int64
					TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
				}
				TypeII_FePortSelection_r19 *struct {
					ParamCombination_r19 int64
					ValueOfN_r19         *int64
				}
				TypeII_Doppler_r19 *struct {
					ParamCombination_Doppler_r19         int64
					Td_Dd_Config_r19                     TD_DD_Config_r18
					PredictionDelay_r19                  int64
					N1_N2_r19                            int64
					TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
				}
			}
			TypeII_RI_Restriction_r19              per.BitString
			NumberOfPMI_SubbandsPerCQI_Subband_r19 *int64
		}
	}
	Cri_TypeI_SinglePanel *struct {
		Cri_TypeI_SinglePanelRI_Restriction_r19 []per.BitString
		Cri_TypeI_SinglePanelN1_N2_CBSR_r19     *CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19
	}
	Cri_TypeII *struct {
		Cri_TypeII_RI_Restriction_r19 []per.BitString
		Cri_TypeII_N1_N2_CBSR_r19     *CRI_TypeII_N1_N2_CBSR_List_r19
	}
}

func (ie *CodebookConfig_r19) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(codebookConfigR19Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r19_CodebookType:
		choiceEnc := e.NewChoiceEncoder(codebookConfigR19CodebookTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.CodebookType).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.CodebookType).Choice {
		case CodebookConfig_r19_CodebookType_Type1:
			choiceEnc := e.NewChoiceEncoder(codebookConfigR19CodebookTypeType1Constraints)
			if err := choiceEnc.EncodeChoice(int64((*(*ie.CodebookType).Type1).Choice), false, nil); err != nil {
				return err
			}
			switch (*(*ie.CodebookType).Type1).Choice {
			case CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19:
				c := (*(*(*ie.CodebookType).Type1).TypeI_SinglePanel_r19)
				codebookConfigR19CodebookTypeType1TypeISinglePanelR19Seq := e.NewSequenceEncoder(codebookConfigR19CodebookTypeType1TypeISinglePanelR19Constraints)
				if err := codebookConfigR19CodebookTypeType1TypeISinglePanelR19Seq.EncodePreamble([]bool{c.TypeI_CodebookSubsetRestriction_r19 != nil, c.TypeI_SoftScalingRank_r19 != nil}); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.CodebookMode_r19, codebookConfigR19CodebookTypeType1TypeISinglePanelR19CodebookModeR19Constraints); err != nil {
					return err
				}
				if err := e.EncodeBitString(c.TypeI_SinglePanelRI_Restriction_r19, per.FixedSize(8)); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.N1_N2_r19, codebookConfigR19CodebookTypeType1TypeISinglePanelR19N1N2R19Constraints); err != nil {
					return err
				}
				if c.TypeI_CodebookSubsetRestriction_r19 != nil {
					if err := c.TypeI_CodebookSubsetRestriction_r19.Encode(e); err != nil {
						return err
					}
				}
				if c.TypeI_SoftScalingRank_r19 != nil {
					if err := c.TypeI_SoftScalingRank_r19.Encode(e); err != nil {
						return err
					}
				}
			case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19:
				c := (*(*(*ie.CodebookType).Type1).TypeI_MultiPanel_r19)
				if err := e.EncodeBitString(c.Ri_Restriction_r19, per.FixedSize(4)); err != nil {
					return err
				}
				{
					choiceEnc := e.NewChoiceEncoder(codebookConfigR19CodebookTypeType1TypeIMultiPanelR19NgN1N2CbsrR19Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.Ng_N1_N2_Cbsr_r19.Choice), false, nil); err != nil {
						return err
					}
					switch c.Ng_N1_N2_Cbsr_r19.Choice {
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Four_Three_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Two_Four_Three_r19), per.FixedSize(192)); err != nil {
							return err
						}
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Six_Two_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Two_Six_Two_r19), per.FixedSize(192)); err != nil {
							return err
						}
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Eight_Two_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Two_Eight_Two_r19), per.FixedSize(256)); err != nil {
							return err
						}
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Four_Two_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Four_Four_Two_r19), per.FixedSize(128)); err != nil {
							return err
						}
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Four_Four_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Two_Four_Four_r19), per.FixedSize(256)); err != nil {
							return err
						}
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Four_Four_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Four_Four_Four_r19), per.FixedSize(256)); err != nil {
							return err
						}
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Eight_Two_r19:
						if err := e.EncodeBitString((*c.Ng_N1_N2_Cbsr_r19.Four_Eight_Two_r19), per.FixedSize(256)); err != nil {
							return err
						}
					}
				}
			}
		case CodebookConfig_r19_CodebookType_Type2:
			c := (*(*ie.CodebookType).Type2)
			codebookConfigR19CodebookTypeType2Seq := e.NewSequenceEncoder(codebookConfigR19CodebookTypeType2Constraints)
			if err := codebookConfigR19CodebookTypeType2Seq.EncodePreamble([]bool{c.NumberOfPMI_SubbandsPerCQI_Subband_r19 != nil}); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(codebookConfigR19CodebookTypeType2TypeIIConstraints)
				if err := choiceEnc.EncodeChoice(int64(c.TypeII.Choice), false, nil); err != nil {
					return err
				}
				switch c.TypeII.Choice {
				case CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19:
					c := (*c.TypeII.EtypeII_r19)
					codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Seq := e.NewSequenceEncoder(codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Constraints)
					if err := codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Seq.EncodePreamble([]bool{c.TypeII_CodebookSubsetRestriction_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.ParamCombination_r19, per.Constrained(1, 8)); err != nil {
						return err
					}
					if err := e.EncodeEnumerated(c.N1_N2_r19, codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19N1N2R19Constraints); err != nil {
						return err
					}
					if c.TypeII_CodebookSubsetRestriction_r19 != nil {
						if err := c.TypeII_CodebookSubsetRestriction_r19.Encode(e); err != nil {
							return err
						}
					}
				case CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19:
					c := (*c.TypeII.TypeII_FePortSelection_r19)
					codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Seq := e.NewSequenceEncoder(codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Constraints)
					if err := codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Seq.EncodePreamble([]bool{c.ValueOfN_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.ParamCombination_r19, per.Constrained(1, 7)); err != nil {
						return err
					}
					if c.ValueOfN_r19 != nil {
						if err := e.EncodeEnumerated((*c.ValueOfN_r19), codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19ValueOfNR19Constraints); err != nil {
							return err
						}
					}
				case CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19:
					c := (*c.TypeII.TypeII_Doppler_r19)
					codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Seq := e.NewSequenceEncoder(codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Constraints)
					if err := codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Seq.EncodePreamble([]bool{c.TypeII_CodebookSubsetRestriction_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.ParamCombination_Doppler_r19, per.Constrained(1, 9)); err != nil {
						return err
					}
					if err := c.Td_Dd_Config_r19.Encode(e); err != nil {
						return err
					}
					if err := e.EncodeEnumerated(c.PredictionDelay_r19, codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19PredictionDelayR19Constraints); err != nil {
						return err
					}
					if err := e.EncodeEnumerated(c.N1_N2_r19, codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19N1N2R19Constraints); err != nil {
						return err
					}
					if c.TypeII_CodebookSubsetRestriction_r19 != nil {
						if err := c.TypeII_CodebookSubsetRestriction_r19.Encode(e); err != nil {
							return err
						}
					}
				}
			}
			if err := e.EncodeBitString(c.TypeII_RI_Restriction_r19, per.FixedSize(4)); err != nil {
				return err
			}
			if c.NumberOfPMI_SubbandsPerCQI_Subband_r19 != nil {
				if err := e.EncodeInteger((*c.NumberOfPMI_SubbandsPerCQI_Subband_r19), per.Constrained(1, 2)); err != nil {
					return err
				}
			}
		}
	case CodebookConfig_r19_Cri_TypeI_SinglePanel:
		c := (*ie.Cri_TypeI_SinglePanel)
		codebookConfigR19CriTypeISinglePanelSeq := e.NewSequenceEncoder(codebookConfigR19CriTypeISinglePanelConstraints)
		if err := codebookConfigR19CriTypeISinglePanelSeq.EncodePreamble([]bool{c.Cri_TypeI_SinglePanelRI_Restriction_r19 != nil, c.Cri_TypeI_SinglePanelN1_N2_CBSR_r19 != nil}); err != nil {
			return err
		}
		if c.Cri_TypeI_SinglePanelRI_Restriction_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookConfigR19CriTypeISinglePanelCriTypeISinglePanelRIRestrictionR19Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Cri_TypeI_SinglePanelRI_Restriction_r19))); err != nil {
				return err
			}
			for i := range c.Cri_TypeI_SinglePanelRI_Restriction_r19 {
				if err := e.EncodeBitString(c.Cri_TypeI_SinglePanelRI_Restriction_r19[i], per.FixedSize(8)); err != nil {
					return err
				}
			}
		}
		if c.Cri_TypeI_SinglePanelN1_N2_CBSR_r19 != nil {
			if err := c.Cri_TypeI_SinglePanelN1_N2_CBSR_r19.Encode(e); err != nil {
				return err
			}
		}
	case CodebookConfig_r19_Cri_TypeII:
		c := (*ie.Cri_TypeII)
		codebookConfigR19CriTypeIISeq := e.NewSequenceEncoder(codebookConfigR19CriTypeIIConstraints)
		if err := codebookConfigR19CriTypeIISeq.EncodePreamble([]bool{c.Cri_TypeII_RI_Restriction_r19 != nil, c.Cri_TypeII_N1_N2_CBSR_r19 != nil}); err != nil {
			return err
		}
		if c.Cri_TypeII_RI_Restriction_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookConfigR19CriTypeIICriTypeIIRIRestrictionR19Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Cri_TypeII_RI_Restriction_r19))); err != nil {
				return err
			}
			for i := range c.Cri_TypeII_RI_Restriction_r19 {
				if err := e.EncodeBitString(c.Cri_TypeII_RI_Restriction_r19[i], per.FixedSize(4)); err != nil {
					return err
				}
			}
		}
		if c.Cri_TypeII_N1_N2_CBSR_r19 != nil {
			if err := c.Cri_TypeII_N1_N2_CBSR_r19.Encode(e); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CodebookConfig-r19",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CodebookConfig_r19) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(codebookConfigR19Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CodebookConfig-r19", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CodebookConfig_r19_CodebookType:
		ie.CodebookType = &struct {
			Choice int
			Type1  *struct {
				Choice                int
				TypeI_SinglePanel_r19 *struct {
					CodebookMode_r19                    int64
					TypeI_SinglePanelRI_Restriction_r19 per.BitString
					N1_N2_r19                           int64
					TypeI_CodebookSubsetRestriction_r19 *TypeI_X1_X2_CBSR_r19
					TypeI_SoftScalingRank_r19           *TypeI_X1_X2_SoftScalingRank_r19
				}
				TypeI_MultiPanel_r19 *struct {
					Ri_Restriction_r19 per.BitString
					Ng_N1_N2_Cbsr_r19  struct {
						Choice             int
						Two_Four_Three_r19 *per.BitString
						Two_Six_Two_r19    *per.BitString
						Two_Eight_Two_r19  *per.BitString
						Four_Four_Two_r19  *per.BitString
						Two_Four_Four_r19  *per.BitString
						Four_Four_Four_r19 *per.BitString
						Four_Eight_Two_r19 *per.BitString
					}
				}
			}
			Type2 *struct {
				TypeII struct {
					Choice      int
					EtypeII_r19 *struct {
						ParamCombination_r19                 int64
						N1_N2_r19                            int64
						TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
					}
					TypeII_FePortSelection_r19 *struct {
						ParamCombination_r19 int64
						ValueOfN_r19         *int64
					}
					TypeII_Doppler_r19 *struct {
						ParamCombination_Doppler_r19         int64
						Td_Dd_Config_r19                     TD_DD_Config_r18
						PredictionDelay_r19                  int64
						N1_N2_r19                            int64
						TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
					}
				}
				TypeII_RI_Restriction_r19              per.BitString
				NumberOfPMI_SubbandsPerCQI_Subband_r19 *int64
			}
		}{}
		choiceDec := d.NewChoiceDecoder(codebookConfigR19CodebookTypeConstraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.CodebookType).Choice = int(index)
		switch index {
		case CodebookConfig_r19_CodebookType_Type1:
			(*ie.CodebookType).Type1 = &struct {
				Choice                int
				TypeI_SinglePanel_r19 *struct {
					CodebookMode_r19                    int64
					TypeI_SinglePanelRI_Restriction_r19 per.BitString
					N1_N2_r19                           int64
					TypeI_CodebookSubsetRestriction_r19 *TypeI_X1_X2_CBSR_r19
					TypeI_SoftScalingRank_r19           *TypeI_X1_X2_SoftScalingRank_r19
				}
				TypeI_MultiPanel_r19 *struct {
					Ri_Restriction_r19 per.BitString
					Ng_N1_N2_Cbsr_r19  struct {
						Choice             int
						Two_Four_Three_r19 *per.BitString
						Two_Six_Two_r19    *per.BitString
						Two_Eight_Two_r19  *per.BitString
						Four_Four_Two_r19  *per.BitString
						Two_Four_Four_r19  *per.BitString
						Four_Four_Four_r19 *per.BitString
						Four_Eight_Two_r19 *per.BitString
					}
				}
			}{}
			choiceDec := d.NewChoiceDecoder(codebookConfigR19CodebookTypeType1Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*(*ie.CodebookType).Type1).Choice = int(index)
			switch index {
			case CodebookConfig_r19_CodebookType_Type1_TypeI_SinglePanel_r19:
				(*(*ie.CodebookType).Type1).TypeI_SinglePanel_r19 = &struct {
					CodebookMode_r19                    int64
					TypeI_SinglePanelRI_Restriction_r19 per.BitString
					N1_N2_r19                           int64
					TypeI_CodebookSubsetRestriction_r19 *TypeI_X1_X2_CBSR_r19
					TypeI_SoftScalingRank_r19           *TypeI_X1_X2_SoftScalingRank_r19
				}{}
				c := (*(*(*ie.CodebookType).Type1).TypeI_SinglePanel_r19)
				codebookConfigR19CodebookTypeType1TypeISinglePanelR19Seq := d.NewSequenceDecoder(codebookConfigR19CodebookTypeType1TypeISinglePanelR19Constraints)
				if err := codebookConfigR19CodebookTypeType1TypeISinglePanelR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := d.DecodeEnumerated(codebookConfigR19CodebookTypeType1TypeISinglePanelR19CodebookModeR19Constraints)
					if err != nil {
						return err
					}
					c.CodebookMode_r19 = v
				}
				{
					v, err := d.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					c.TypeI_SinglePanelRI_Restriction_r19 = v
				}
				{
					v, err := d.DecodeEnumerated(codebookConfigR19CodebookTypeType1TypeISinglePanelR19N1N2R19Constraints)
					if err != nil {
						return err
					}
					c.N1_N2_r19 = v
				}
				if codebookConfigR19CodebookTypeType1TypeISinglePanelR19Seq.IsComponentPresent(3) {
					c.TypeI_CodebookSubsetRestriction_r19 = new(TypeI_X1_X2_CBSR_r19)
					if err := (*c.TypeI_CodebookSubsetRestriction_r19).Decode(d); err != nil {
						return err
					}
				}
				if codebookConfigR19CodebookTypeType1TypeISinglePanelR19Seq.IsComponentPresent(4) {
					c.TypeI_SoftScalingRank_r19 = new(TypeI_X1_X2_SoftScalingRank_r19)
					if err := (*c.TypeI_SoftScalingRank_r19).Decode(d); err != nil {
						return err
					}
				}
			case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19:
				(*(*ie.CodebookType).Type1).TypeI_MultiPanel_r19 = &struct {
					Ri_Restriction_r19 per.BitString
					Ng_N1_N2_Cbsr_r19  struct {
						Choice             int
						Two_Four_Three_r19 *per.BitString
						Two_Six_Two_r19    *per.BitString
						Two_Eight_Two_r19  *per.BitString
						Four_Four_Two_r19  *per.BitString
						Two_Four_Four_r19  *per.BitString
						Four_Four_Four_r19 *per.BitString
						Four_Eight_Two_r19 *per.BitString
					}
				}{}
				c := (*(*(*ie.CodebookType).Type1).TypeI_MultiPanel_r19)
				{
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					c.Ri_Restriction_r19 = v
				}
				{
					choiceDec := d.NewChoiceDecoder(codebookConfigR19CodebookTypeType1TypeIMultiPanelR19NgN1N2CbsrR19Constraints)
					index, _, _, err := choiceDec.DecodeChoice()
					if err != nil {
						return err
					}
					c.Ng_N1_N2_Cbsr_r19.Choice = int(index)
					switch index {
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Four_Three_r19:
						c.Ng_N1_N2_Cbsr_r19.Two_Four_Three_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(192))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Two_Four_Three_r19) = v
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Six_Two_r19:
						c.Ng_N1_N2_Cbsr_r19.Two_Six_Two_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(192))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Two_Six_Two_r19) = v
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Eight_Two_r19:
						c.Ng_N1_N2_Cbsr_r19.Two_Eight_Two_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(256))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Two_Eight_Two_r19) = v
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Four_Two_r19:
						c.Ng_N1_N2_Cbsr_r19.Four_Four_Two_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(128))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Four_Four_Two_r19) = v
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Two_Four_Four_r19:
						c.Ng_N1_N2_Cbsr_r19.Two_Four_Four_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(256))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Two_Four_Four_r19) = v
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Four_Four_r19:
						c.Ng_N1_N2_Cbsr_r19.Four_Four_Four_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(256))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Four_Four_Four_r19) = v
					case CodebookConfig_r19_CodebookType_Type1_TypeI_MultiPanel_r19_Ng_N1_N2_Cbsr_r19_Four_Eight_Two_r19:
						c.Ng_N1_N2_Cbsr_r19.Four_Eight_Two_r19 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(256))
						if err != nil {
							return err
						}
						(*c.Ng_N1_N2_Cbsr_r19.Four_Eight_Two_r19) = v
					}
				}
			}
		case CodebookConfig_r19_CodebookType_Type2:
			(*ie.CodebookType).Type2 = &struct {
				TypeII struct {
					Choice      int
					EtypeII_r19 *struct {
						ParamCombination_r19                 int64
						N1_N2_r19                            int64
						TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
					}
					TypeII_FePortSelection_r19 *struct {
						ParamCombination_r19 int64
						ValueOfN_r19         *int64
					}
					TypeII_Doppler_r19 *struct {
						ParamCombination_Doppler_r19         int64
						Td_Dd_Config_r19                     TD_DD_Config_r18
						PredictionDelay_r19                  int64
						N1_N2_r19                            int64
						TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
					}
				}
				TypeII_RI_Restriction_r19              per.BitString
				NumberOfPMI_SubbandsPerCQI_Subband_r19 *int64
			}{}
			c := (*(*ie.CodebookType).Type2)
			codebookConfigR19CodebookTypeType2Seq := d.NewSequenceDecoder(codebookConfigR19CodebookTypeType2Constraints)
			if err := codebookConfigR19CodebookTypeType2Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := d.NewChoiceDecoder(codebookConfigR19CodebookTypeType2TypeIIConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.TypeII.Choice = int(index)
				switch index {
				case CodebookConfig_r19_CodebookType_Type2_TypeII_EtypeII_r19:
					c.TypeII.EtypeII_r19 = &struct {
						ParamCombination_r19                 int64
						N1_N2_r19                            int64
						TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
					}{}
					c := (*c.TypeII.EtypeII_r19)
					codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Seq := d.NewSequenceDecoder(codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Constraints)
					if err := codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 8))
						if err != nil {
							return err
						}
						c.ParamCombination_r19 = v
					}
					{
						v, err := d.DecodeEnumerated(codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19N1N2R19Constraints)
						if err != nil {
							return err
						}
						c.N1_N2_r19 = v
					}
					if codebookConfigR19CodebookTypeType2TypeIIEtypeIIR19Seq.IsComponentPresent(2) {
						c.TypeII_CodebookSubsetRestriction_r19 = new(TypeII_X1_X2_CBSR_r19)
						if err := (*c.TypeII_CodebookSubsetRestriction_r19).Decode(d); err != nil {
							return err
						}
					}
				case CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_FePortSelection_r19:
					c.TypeII.TypeII_FePortSelection_r19 = &struct {
						ParamCombination_r19 int64
						ValueOfN_r19         *int64
					}{}
					c := (*c.TypeII.TypeII_FePortSelection_r19)
					codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Seq := d.NewSequenceDecoder(codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Constraints)
					if err := codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 7))
						if err != nil {
							return err
						}
						c.ParamCombination_r19 = v
					}
					if codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19Seq.IsComponentPresent(1) {
						c.ValueOfN_r19 = new(int64)
						v, err := d.DecodeEnumerated(codebookConfigR19CodebookTypeType2TypeIITypeIIFePortSelectionR19ValueOfNR19Constraints)
						if err != nil {
							return err
						}
						(*c.ValueOfN_r19) = v
					}
				case CodebookConfig_r19_CodebookType_Type2_TypeII_TypeII_Doppler_r19:
					c.TypeII.TypeII_Doppler_r19 = &struct {
						ParamCombination_Doppler_r19         int64
						Td_Dd_Config_r19                     TD_DD_Config_r18
						PredictionDelay_r19                  int64
						N1_N2_r19                            int64
						TypeII_CodebookSubsetRestriction_r19 *TypeII_X1_X2_CBSR_r19
					}{}
					c := (*c.TypeII.TypeII_Doppler_r19)
					codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Seq := d.NewSequenceDecoder(codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Constraints)
					if err := codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 9))
						if err != nil {
							return err
						}
						c.ParamCombination_Doppler_r19 = v
					}
					{
						if err := c.Td_Dd_Config_r19.Decode(d); err != nil {
							return err
						}
					}
					{
						v, err := d.DecodeEnumerated(codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19PredictionDelayR19Constraints)
						if err != nil {
							return err
						}
						c.PredictionDelay_r19 = v
					}
					{
						v, err := d.DecodeEnumerated(codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19N1N2R19Constraints)
						if err != nil {
							return err
						}
						c.N1_N2_r19 = v
					}
					if codebookConfigR19CodebookTypeType2TypeIITypeIIDopplerR19Seq.IsComponentPresent(4) {
						c.TypeII_CodebookSubsetRestriction_r19 = new(TypeII_X1_X2_CBSR_r19)
						if err := (*c.TypeII_CodebookSubsetRestriction_r19).Decode(d); err != nil {
							return err
						}
					}
				}
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				c.TypeII_RI_Restriction_r19 = v
			}
			if codebookConfigR19CodebookTypeType2Seq.IsComponentPresent(2) {
				c.NumberOfPMI_SubbandsPerCQI_Subband_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				(*c.NumberOfPMI_SubbandsPerCQI_Subband_r19) = v
			}
		}
	case CodebookConfig_r19_Cri_TypeI_SinglePanel:
		ie.Cri_TypeI_SinglePanel = &struct {
			Cri_TypeI_SinglePanelRI_Restriction_r19 []per.BitString
			Cri_TypeI_SinglePanelN1_N2_CBSR_r19     *CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19
		}{}
		c := (*ie.Cri_TypeI_SinglePanel)
		codebookConfigR19CriTypeISinglePanelSeq := d.NewSequenceDecoder(codebookConfigR19CriTypeISinglePanelConstraints)
		if err := codebookConfigR19CriTypeISinglePanelSeq.DecodePreamble(); err != nil {
			return err
		}
		if codebookConfigR19CriTypeISinglePanelSeq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(codebookConfigR19CriTypeISinglePanelCriTypeISinglePanelRIRestrictionR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Cri_TypeI_SinglePanelRI_Restriction_r19 = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				c.Cri_TypeI_SinglePanelRI_Restriction_r19[i] = v
			}
		}
		if codebookConfigR19CriTypeISinglePanelSeq.IsComponentPresent(1) {
			c.Cri_TypeI_SinglePanelN1_N2_CBSR_r19 = new(CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19)
			if err := (*c.Cri_TypeI_SinglePanelN1_N2_CBSR_r19).Decode(d); err != nil {
				return err
			}
		}
	case CodebookConfig_r19_Cri_TypeII:
		ie.Cri_TypeII = &struct {
			Cri_TypeII_RI_Restriction_r19 []per.BitString
			Cri_TypeII_N1_N2_CBSR_r19     *CRI_TypeII_N1_N2_CBSR_List_r19
		}{}
		c := (*ie.Cri_TypeII)
		codebookConfigR19CriTypeIISeq := d.NewSequenceDecoder(codebookConfigR19CriTypeIIConstraints)
		if err := codebookConfigR19CriTypeIISeq.DecodePreamble(); err != nil {
			return err
		}
		if codebookConfigR19CriTypeIISeq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(codebookConfigR19CriTypeIICriTypeIIRIRestrictionR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Cri_TypeII_RI_Restriction_r19 = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				c.Cri_TypeII_RI_Restriction_r19[i] = v
			}
		}
		if codebookConfigR19CriTypeIISeq.IsComponentPresent(1) {
			c.Cri_TypeII_N1_N2_CBSR_r19 = new(CRI_TypeII_N1_N2_CBSR_List_r19)
			if err := (*c.Cri_TypeII_N1_N2_CBSR_r19).Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "CodebookConfig-r19", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
