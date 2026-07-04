// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookConfig-r18 (line 6230).

var codebookConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookType"},
	},
}

var codebookConfig_r18CodebookTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type2"},
	},
}

const (
	CodebookConfig_r18_CodebookType_Type2 = 0
)

var codebookConfigR18CodebookTypeType2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "typeII-CJT-r18"},
		{Name: "typeII-CJT-PortSelection-r18"},
		{Name: "typeII-Doppler-r18"},
		{Name: "typeII-DopplerPortSelection-r18"},
	},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18                  = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18    = 1
	CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18              = 2
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18 = 3
)

var codebookConfigR18CodebookTypeType2TypeIICJTR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "n1-n2-codebookSubsetRestrictionList-r18"},
		{Name: "paramCombination-CJT-r18"},
		{Name: "paramCombination-CJT-L-r18"},
		{Name: "restrictedCMR-Selection-r18"},
		{Name: "valueOfO3-r18", Optional: true},
		{Name: "numberOfPMI-SubbandsPerCQI-Subband-r18"},
		{Name: "typeII-RI-Restriction-r18"},
		{Name: "codebookMode-r18"},
	},
}

var codebookConfigR18CodebookTypeType2TypeIICJTR18ParamCombinationCJTLR18Constraints = per.SizeRange(1, 4)

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18_RestrictedCMR_Selection_r18_Enable = 0
)

var codebookConfigR18CodebookTypeType2TypeIICJTR18RestrictedCMRSelectionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18_RestrictedCMR_Selection_r18_Enable},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18_ValueOfO3_r18_N1 = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18_ValueOfO3_r18_N4 = 1
)

var codebookConfigR18CodebookTypeType2TypeIICJTR18ValueOfO3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18_ValueOfO3_r18_N1, CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18_ValueOfO3_r18_N4},
}

var codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "paramCombination-CJT-PS-r18"},
		{Name: "paramCombination-CJT-PS-alpha-r18"},
		{Name: "restrictedCMR-Selection-r18"},
		{Name: "valueOfO3-r18", Optional: true},
		{Name: "valueOfN-CJT-r18", Optional: true},
		{Name: "numberOfPMI-SubbandsPerCQI-Subband-r18"},
		{Name: "typeII-PortSelectionRI-Restriction-r18"},
		{Name: "codebookMode-r18"},
	},
}

var codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ParamCombinationCJTPSAlphaR18Constraints = per.SizeRange(1, 4)

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_RestrictedCMR_Selection_r18_Enable = 0
)

var codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18RestrictedCMRSelectionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_RestrictedCMR_Selection_r18_Enable},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfO3_r18_N1 = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfO3_r18_N4 = 1
)

var codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ValueOfO3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfO3_r18_N1, CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfO3_r18_N4},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfN_CJT_r18_N2 = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfN_CJT_r18_N4 = 1
)

var codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ValueOfNCJTR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfN_CJT_r18_N2, CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18_ValueOfN_CJT_r18_N4},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_M0 = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_N0 = 1
	CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_N1 = 2
	CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_N2 = 3
)

var codebookConfigR18CodebookTypeType2TypeIIDopplerR18PredictionDelayR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_M0, CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_N0, CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_N1, CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18_PredictionDelay_r18_N2},
}

var codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "paramCombinationDoppler-PS-r18"},
		{Name: "td-dd-config-r18"},
		{Name: "valueOfN-Doppler-r18", Optional: true},
		{Name: "numberOfPMI-SubbandsPerCQI-Subband-r18"},
		{Name: "predictionDelay-r18"},
		{Name: "typeII-PortSelectionRI-Restriction-r18"},
	},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_ValueOfN_Doppler_r18_N2 = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_ValueOfN_Doppler_r18_N4 = 1
)

var codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18ValueOfNDopplerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_ValueOfN_Doppler_r18_N2, CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_ValueOfN_Doppler_r18_N4},
}

const (
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_M0 = 0
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_N0 = 1
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_N1 = 2
	CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_N2 = 3
)

var codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18PredictionDelayR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_M0, CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_N0, CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_N1, CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18_PredictionDelay_r18_N2},
}

type CodebookConfig_r18 struct {
	CodebookType struct {
		Choice int
		Type2  *struct {
			Choice         int
			TypeII_CJT_r18 *struct {
				N1_N2_CodebookSubsetRestrictionList_r18 N1_N2_CBSR_List_r18
				ParamCombination_CJT_r18                int64
				ParamCombination_CJT_L_r18              []int64
				RestrictedCMR_Selection_r18             int64
				ValueOfO3_r18                           *int64
				NumberOfPMI_SubbandsPerCQI_Subband_r18  int64
				TypeII_RI_Restriction_r18               per.BitString
				CodebookMode_r18                        int64
			}
			TypeII_CJT_PortSelection_r18 *struct {
				ParamCombination_CJT_PS_r18            int64
				ParamCombination_CJT_PS_Alpha_r18      []int64
				RestrictedCMR_Selection_r18            int64
				ValueOfO3_r18                          *int64
				ValueOfN_CJT_r18                       *int64
				NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
				TypeII_PortSelectionRI_Restriction_r18 per.BitString
				CodebookMode_r18                       int64
			}
			TypeII_Doppler_r18 *struct {
				N1_N2_CodebookSubsetRestriction_r18    N1_N2_CBSR_r18
				ParamCombination_Doppler_r18           int64
				Td_Dd_Config_r18                       TD_DD_Config_r18
				NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
				PredictionDelay_r18                    int64
				TypeII_RI_Restriction_r18              per.BitString
			}
			TypeII_DopplerPortSelection_r18 *struct {
				ParamCombinationDoppler_PS_r18         int64
				Td_Dd_Config_r18                       TD_DD_Config_r18
				ValueOfN_Doppler_r18                   *int64
				NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
				PredictionDelay_r18                    int64
				TypeII_PortSelectionRI_Restriction_r18 per.BitString
			}
		}
	}
}

func (ie *CodebookConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookConfigR18Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceEnc := e.NewChoiceEncoder(codebookConfig_r18CodebookTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CodebookType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CodebookType.Choice {
		case CodebookConfig_r18_CodebookType_Type2:
			choiceEnc := e.NewChoiceEncoder(codebookConfigR18CodebookTypeType2Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CodebookType.Type2).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CodebookType.Type2).Choice {
			case CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18:
				c := (*(*ie.CodebookType.Type2).TypeII_CJT_r18)
				codebookConfigR18CodebookTypeType2TypeIICJTR18Seq := e.NewSequenceEncoder(codebookConfigR18CodebookTypeType2TypeIICJTR18Constraints)
				if err := codebookConfigR18CodebookTypeType2TypeIICJTR18Seq.EncodePreamble([]bool{c.ValueOfO3_r18 != nil}); err != nil {
					return err
				}
				if err := c.N1_N2_CodebookSubsetRestrictionList_r18.Encode(e); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.ParamCombination_CJT_r18, per.Constrained(1, 7)); err != nil {
					return err
				}
				{
					seqOf := e.NewSequenceOfEncoder(codebookConfigR18CodebookTypeType2TypeIICJTR18ParamCombinationCJTLR18Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ParamCombination_CJT_L_r18))); err != nil {
						return err
					}
					for i := range c.ParamCombination_CJT_L_r18 {
						if err := e.EncodeInteger(c.ParamCombination_CJT_L_r18[i], per.Constrained(1, 5)); err != nil {
							return err
						}
					}
				}
				if err := e.EncodeEnumerated(c.RestrictedCMR_Selection_r18, codebookConfigR18CodebookTypeType2TypeIICJTR18RestrictedCMRSelectionR18Constraints); err != nil {
					return err
				}
				if c.ValueOfO3_r18 != nil {
					if err := e.EncodeEnumerated((*c.ValueOfO3_r18), codebookConfigR18CodebookTypeType2TypeIICJTR18ValueOfO3R18Constraints); err != nil {
						return err
					}
				}
				if err := e.EncodeInteger(c.NumberOfPMI_SubbandsPerCQI_Subband_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := e.EncodeBitString(c.TypeII_RI_Restriction_r18, per.FixedSize(4)); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.CodebookMode_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
			case CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18:
				c := (*(*ie.CodebookType.Type2).TypeII_CJT_PortSelection_r18)
				codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Seq := e.NewSequenceEncoder(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Constraints)
				if err := codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Seq.EncodePreamble([]bool{c.ValueOfO3_r18 != nil, c.ValueOfN_CJT_r18 != nil}); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.ParamCombination_CJT_PS_r18, per.Constrained(1, 5)); err != nil {
					return err
				}
				{
					seqOf := e.NewSequenceOfEncoder(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ParamCombinationCJTPSAlphaR18Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ParamCombination_CJT_PS_Alpha_r18))); err != nil {
						return err
					}
					for i := range c.ParamCombination_CJT_PS_Alpha_r18 {
						if err := e.EncodeInteger(c.ParamCombination_CJT_PS_Alpha_r18[i], per.Constrained(1, 8)); err != nil {
							return err
						}
					}
				}
				if err := e.EncodeEnumerated(c.RestrictedCMR_Selection_r18, codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18RestrictedCMRSelectionR18Constraints); err != nil {
					return err
				}
				if c.ValueOfO3_r18 != nil {
					if err := e.EncodeEnumerated((*c.ValueOfO3_r18), codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ValueOfO3R18Constraints); err != nil {
						return err
					}
				}
				if c.ValueOfN_CJT_r18 != nil {
					if err := e.EncodeEnumerated((*c.ValueOfN_CJT_r18), codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ValueOfNCJTR18Constraints); err != nil {
						return err
					}
				}
				if err := e.EncodeInteger(c.NumberOfPMI_SubbandsPerCQI_Subband_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := e.EncodeBitString(c.TypeII_PortSelectionRI_Restriction_r18, per.FixedSize(4)); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.CodebookMode_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
			case CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18:
				c := (*(*ie.CodebookType.Type2).TypeII_Doppler_r18)
				if err := c.N1_N2_CodebookSubsetRestriction_r18.Encode(e); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.ParamCombination_Doppler_r18, per.Constrained(1, 9)); err != nil {
					return err
				}
				if err := c.Td_Dd_Config_r18.Encode(e); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.NumberOfPMI_SubbandsPerCQI_Subband_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.PredictionDelay_r18, codebookConfigR18CodebookTypeType2TypeIIDopplerR18PredictionDelayR18Constraints); err != nil {
					return err
				}
				if err := e.EncodeBitString(c.TypeII_RI_Restriction_r18, per.FixedSize(4)); err != nil {
					return err
				}
			case CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18:
				c := (*(*ie.CodebookType.Type2).TypeII_DopplerPortSelection_r18)
				codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Seq := e.NewSequenceEncoder(codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Constraints)
				if err := codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Seq.EncodePreamble([]bool{c.ValueOfN_Doppler_r18 != nil}); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.ParamCombinationDoppler_PS_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := c.Td_Dd_Config_r18.Encode(e); err != nil {
					return err
				}
				if c.ValueOfN_Doppler_r18 != nil {
					if err := e.EncodeEnumerated((*c.ValueOfN_Doppler_r18), codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18ValueOfNDopplerR18Constraints); err != nil {
						return err
					}
				}
				if err := e.EncodeInteger(c.NumberOfPMI_SubbandsPerCQI_Subband_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.PredictionDelay_r18, codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18PredictionDelayR18Constraints); err != nil {
					return err
				}
				if err := e.EncodeBitString(c.TypeII_PortSelectionRI_Restriction_r18, per.FixedSize(4)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CodebookType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CodebookConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookConfigR18Constraints)
	_ = seq

	// 1. codebookType: choice
	{
		choiceDec := d.NewChoiceDecoder(codebookConfig_r18CodebookTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CodebookType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CodebookConfig_r18_CodebookType_Type2:
			ie.CodebookType.Type2 = &struct {
				Choice         int
				TypeII_CJT_r18 *struct {
					N1_N2_CodebookSubsetRestrictionList_r18 N1_N2_CBSR_List_r18
					ParamCombination_CJT_r18                int64
					ParamCombination_CJT_L_r18              []int64
					RestrictedCMR_Selection_r18             int64
					ValueOfO3_r18                           *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r18  int64
					TypeII_RI_Restriction_r18               per.BitString
					CodebookMode_r18                        int64
				}
				TypeII_CJT_PortSelection_r18 *struct {
					ParamCombination_CJT_PS_r18            int64
					ParamCombination_CJT_PS_Alpha_r18      []int64
					RestrictedCMR_Selection_r18            int64
					ValueOfO3_r18                          *int64
					ValueOfN_CJT_r18                       *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
					TypeII_PortSelectionRI_Restriction_r18 per.BitString
					CodebookMode_r18                       int64
				}
				TypeII_Doppler_r18 *struct {
					N1_N2_CodebookSubsetRestriction_r18    N1_N2_CBSR_r18
					ParamCombination_Doppler_r18           int64
					Td_Dd_Config_r18                       TD_DD_Config_r18
					NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
					PredictionDelay_r18                    int64
					TypeII_RI_Restriction_r18              per.BitString
				}
				TypeII_DopplerPortSelection_r18 *struct {
					ParamCombinationDoppler_PS_r18         int64
					Td_Dd_Config_r18                       TD_DD_Config_r18
					ValueOfN_Doppler_r18                   *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
					PredictionDelay_r18                    int64
					TypeII_PortSelectionRI_Restriction_r18 per.BitString
				}
			}{}
			choiceDec := d.NewChoiceDecoder(codebookConfigR18CodebookTypeType2Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CodebookType.Type2).Choice = int(index)
			switch index {
			case CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_r18:
				(*ie.CodebookType.Type2).TypeII_CJT_r18 = &struct {
					N1_N2_CodebookSubsetRestrictionList_r18 N1_N2_CBSR_List_r18
					ParamCombination_CJT_r18                int64
					ParamCombination_CJT_L_r18              []int64
					RestrictedCMR_Selection_r18             int64
					ValueOfO3_r18                           *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r18  int64
					TypeII_RI_Restriction_r18               per.BitString
					CodebookMode_r18                        int64
				}{}
				c := (*(*ie.CodebookType.Type2).TypeII_CJT_r18)
				codebookConfigR18CodebookTypeType2TypeIICJTR18Seq := d.NewSequenceDecoder(codebookConfigR18CodebookTypeType2TypeIICJTR18Constraints)
				if err := codebookConfigR18CodebookTypeType2TypeIICJTR18Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.N1_N2_CodebookSubsetRestrictionList_r18.Decode(d); err != nil {
						return err
					}
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 7))
					if err != nil {
						return err
					}
					c.ParamCombination_CJT_r18 = v
				}
				{
					seqOf := d.NewSequenceOfDecoder(codebookConfigR18CodebookTypeType2TypeIICJTR18ParamCombinationCJTLR18Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.ParamCombination_CJT_L_r18 = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := d.DecodeInteger(per.Constrained(1, 5))
						if err != nil {
							return err
						}
						c.ParamCombination_CJT_L_r18[i] = v
					}
				}
				{
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIICJTR18RestrictedCMRSelectionR18Constraints)
					if err != nil {
						return err
					}
					c.RestrictedCMR_Selection_r18 = v
				}
				if codebookConfigR18CodebookTypeType2TypeIICJTR18Seq.IsComponentPresent(4) {
					c.ValueOfO3_r18 = new(int64)
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIICJTR18ValueOfO3R18Constraints)
					if err != nil {
						return err
					}
					(*c.ValueOfO3_r18) = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.NumberOfPMI_SubbandsPerCQI_Subband_r18 = v
				}
				{
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					c.TypeII_RI_Restriction_r18 = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.CodebookMode_r18 = v
				}
			case CodebookConfig_r18_CodebookType_Type2_TypeII_CJT_PortSelection_r18:
				(*ie.CodebookType.Type2).TypeII_CJT_PortSelection_r18 = &struct {
					ParamCombination_CJT_PS_r18            int64
					ParamCombination_CJT_PS_Alpha_r18      []int64
					RestrictedCMR_Selection_r18            int64
					ValueOfO3_r18                          *int64
					ValueOfN_CJT_r18                       *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
					TypeII_PortSelectionRI_Restriction_r18 per.BitString
					CodebookMode_r18                       int64
				}{}
				c := (*(*ie.CodebookType.Type2).TypeII_CJT_PortSelection_r18)
				codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Seq := d.NewSequenceDecoder(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Constraints)
				if err := codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 5))
					if err != nil {
						return err
					}
					c.ParamCombination_CJT_PS_r18 = v
				}
				{
					seqOf := d.NewSequenceOfDecoder(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ParamCombinationCJTPSAlphaR18Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.ParamCombination_CJT_PS_Alpha_r18 = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := d.DecodeInteger(per.Constrained(1, 8))
						if err != nil {
							return err
						}
						c.ParamCombination_CJT_PS_Alpha_r18[i] = v
					}
				}
				{
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18RestrictedCMRSelectionR18Constraints)
					if err != nil {
						return err
					}
					c.RestrictedCMR_Selection_r18 = v
				}
				if codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Seq.IsComponentPresent(3) {
					c.ValueOfO3_r18 = new(int64)
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ValueOfO3R18Constraints)
					if err != nil {
						return err
					}
					(*c.ValueOfO3_r18) = v
				}
				if codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18Seq.IsComponentPresent(4) {
					c.ValueOfN_CJT_r18 = new(int64)
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIICJTPortSelectionR18ValueOfNCJTR18Constraints)
					if err != nil {
						return err
					}
					(*c.ValueOfN_CJT_r18) = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.NumberOfPMI_SubbandsPerCQI_Subband_r18 = v
				}
				{
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					c.TypeII_PortSelectionRI_Restriction_r18 = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.CodebookMode_r18 = v
				}
			case CodebookConfig_r18_CodebookType_Type2_TypeII_Doppler_r18:
				(*ie.CodebookType.Type2).TypeII_Doppler_r18 = &struct {
					N1_N2_CodebookSubsetRestriction_r18    N1_N2_CBSR_r18
					ParamCombination_Doppler_r18           int64
					Td_Dd_Config_r18                       TD_DD_Config_r18
					NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
					PredictionDelay_r18                    int64
					TypeII_RI_Restriction_r18              per.BitString
				}{}
				c := (*(*ie.CodebookType.Type2).TypeII_Doppler_r18)
				{
					if err := c.N1_N2_CodebookSubsetRestriction_r18.Decode(d); err != nil {
						return err
					}
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 9))
					if err != nil {
						return err
					}
					c.ParamCombination_Doppler_r18 = v
				}
				{
					if err := c.Td_Dd_Config_r18.Decode(d); err != nil {
						return err
					}
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.NumberOfPMI_SubbandsPerCQI_Subband_r18 = v
				}
				{
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIIDopplerR18PredictionDelayR18Constraints)
					if err != nil {
						return err
					}
					c.PredictionDelay_r18 = v
				}
				{
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					c.TypeII_RI_Restriction_r18 = v
				}
			case CodebookConfig_r18_CodebookType_Type2_TypeII_DopplerPortSelection_r18:
				(*ie.CodebookType.Type2).TypeII_DopplerPortSelection_r18 = &struct {
					ParamCombinationDoppler_PS_r18         int64
					Td_Dd_Config_r18                       TD_DD_Config_r18
					ValueOfN_Doppler_r18                   *int64
					NumberOfPMI_SubbandsPerCQI_Subband_r18 int64
					PredictionDelay_r18                    int64
					TypeII_PortSelectionRI_Restriction_r18 per.BitString
				}{}
				c := (*(*ie.CodebookType.Type2).TypeII_DopplerPortSelection_r18)
				codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Seq := d.NewSequenceDecoder(codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Constraints)
				if err := codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 8))
					if err != nil {
						return err
					}
					c.ParamCombinationDoppler_PS_r18 = v
				}
				{
					if err := c.Td_Dd_Config_r18.Decode(d); err != nil {
						return err
					}
				}
				if codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18Seq.IsComponentPresent(2) {
					c.ValueOfN_Doppler_r18 = new(int64)
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18ValueOfNDopplerR18Constraints)
					if err != nil {
						return err
					}
					(*c.ValueOfN_Doppler_r18) = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.NumberOfPMI_SubbandsPerCQI_Subband_r18 = v
				}
				{
					v, err := d.DecodeEnumerated(codebookConfigR18CodebookTypeType2TypeIIDopplerPortSelectionR18PredictionDelayR18Constraints)
					if err != nil {
						return err
					}
					c.PredictionDelay_r18 = v
				}
				{
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					c.TypeII_PortSelectionRI_Restriction_r18 = v
				}
			}
		}
	}

	return nil
}
