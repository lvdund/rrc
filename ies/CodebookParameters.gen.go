// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParameters (line 18489).

var codebookParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type1"},
		{Name: "type2", Optional: true},
		{Name: "type2-PortSelection", Optional: true},
	},
}

var codebookParametersType1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "singlePanel"},
		{Name: "multiPanel", Optional: true},
	},
}

var codebookParametersType1SinglePanelSupportedCSIRSResourceListConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

const (
	CodebookParameters_Type1_SinglePanel_Modes_Mode1         = 0
	CodebookParameters_Type1_SinglePanel_Modes_Mode1andMode2 = 1
)

var codebookParametersType1SinglePanelModesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameters_Type1_SinglePanel_Modes_Mode1, CodebookParameters_Type1_SinglePanel_Modes_Mode1andMode2},
}

var codebookParametersType1MultiPanelSupportedCSIRSResourceListConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

const (
	CodebookParameters_Type1_MultiPanel_Modes_Mode1 = 0
	CodebookParameters_Type1_MultiPanel_Modes_Mode2 = 1
	CodebookParameters_Type1_MultiPanel_Modes_Both  = 2
)

var codebookParametersType1MultiPanelModesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameters_Type1_MultiPanel_Modes_Mode1, CodebookParameters_Type1_MultiPanel_Modes_Mode2, CodebookParameters_Type1_MultiPanel_Modes_Both},
}

const (
	CodebookParameters_Type1_MultiPanel_NrofPanels_N2 = 0
	CodebookParameters_Type1_MultiPanel_NrofPanels_N4 = 1
)

var codebookParametersType1MultiPanelNrofPanelsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameters_Type1_MultiPanel_NrofPanels_N2, CodebookParameters_Type1_MultiPanel_NrofPanels_N4},
}

var codebookParametersType2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedCSI-RS-ResourceList"},
		{Name: "parameterLx"},
		{Name: "amplitudeScalingType"},
		{Name: "amplitudeSubsetRestriction", Optional: true},
	},
}

var codebookParametersType2SupportedCSIRSResourceListConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

const (
	CodebookParameters_Type2_AmplitudeScalingType_Wideband           = 0
	CodebookParameters_Type2_AmplitudeScalingType_WidebandAndSubband = 1
)

var codebookParametersType2AmplitudeScalingTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameters_Type2_AmplitudeScalingType_Wideband, CodebookParameters_Type2_AmplitudeScalingType_WidebandAndSubband},
}

const (
	CodebookParameters_Type2_AmplitudeSubsetRestriction_Supported = 0
)

var codebookParametersType2AmplitudeSubsetRestrictionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameters_Type2_AmplitudeSubsetRestriction_Supported},
}

var codebookParametersType2PortSelectionSupportedCSIRSResourceListConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

const (
	CodebookParameters_Type2_PortSelection_AmplitudeScalingType_Wideband           = 0
	CodebookParameters_Type2_PortSelection_AmplitudeScalingType_WidebandAndSubband = 1
)

var codebookParametersType2PortSelectionAmplitudeScalingTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameters_Type2_PortSelection_AmplitudeScalingType_Wideband, CodebookParameters_Type2_PortSelection_AmplitudeScalingType_WidebandAndSubband},
}

type CodebookParameters struct {
	Type1 struct {
		SinglePanel struct {
			SupportedCSI_RS_ResourceList   []SupportedCSI_RS_Resource
			Modes                          int64
			MaxNumberCSI_RS_PerResourceSet int64
		}
		MultiPanel *struct {
			SupportedCSI_RS_ResourceList   []SupportedCSI_RS_Resource
			Modes                          int64
			NrofPanels                     int64
			MaxNumberCSI_RS_PerResourceSet int64
		}
	}
	Type2 *struct {
		SupportedCSI_RS_ResourceList []SupportedCSI_RS_Resource
		ParameterLx                  int64
		AmplitudeScalingType         int64
		AmplitudeSubsetRestriction   *int64
	}
	Type2_PortSelection *struct {
		SupportedCSI_RS_ResourceList []SupportedCSI_RS_Resource
		ParameterLx                  int64
		AmplitudeScalingType         int64
	}
}

func (ie *CodebookParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Type2 != nil, ie.Type2_PortSelection != nil}); err != nil {
		return err
	}

	// 2. type1: sequence
	{
		{
			c := &ie.Type1
			codebookParametersType1Seq := e.NewSequenceEncoder(codebookParametersType1Constraints)
			if err := codebookParametersType1Seq.EncodePreamble([]bool{c.MultiPanel != nil}); err != nil {
				return err
			}
			{
				c := &c.SinglePanel
				{
					seqOf := e.NewSequenceOfEncoder(codebookParametersType1SinglePanelSupportedCSIRSResourceListConstraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList))); err != nil {
						return err
					}
					for i := range c.SupportedCSI_RS_ResourceList {
						if err := c.SupportedCSI_RS_ResourceList[i].Encode(e); err != nil {
							return err
						}
					}
				}
				if err := e.EncodeEnumerated(c.Modes, codebookParametersType1SinglePanelModesConstraints); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.MaxNumberCSI_RS_PerResourceSet, per.Constrained(1, 8)); err != nil {
					return err
				}
			}
			if c.MultiPanel != nil {
				c := (*c.MultiPanel)
				{
					seqOf := e.NewSequenceOfEncoder(codebookParametersType1MultiPanelSupportedCSIRSResourceListConstraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList))); err != nil {
						return err
					}
					for i := range c.SupportedCSI_RS_ResourceList {
						if err := c.SupportedCSI_RS_ResourceList[i].Encode(e); err != nil {
							return err
						}
					}
				}
				if err := e.EncodeEnumerated(c.Modes, codebookParametersType1MultiPanelModesConstraints); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.NrofPanels, codebookParametersType1MultiPanelNrofPanelsConstraints); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.MaxNumberCSI_RS_PerResourceSet, per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	// 3. type2: sequence
	{
		if ie.Type2 != nil {
			c := ie.Type2
			codebookParametersType2Seq := e.NewSequenceEncoder(codebookParametersType2Constraints)
			if err := codebookParametersType2Seq.EncodePreamble([]bool{c.AmplitudeSubsetRestriction != nil}); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType2SupportedCSIRSResourceListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceList {
					if err := c.SupportedCSI_RS_ResourceList[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.ParameterLx, per.Constrained(2, 4)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.AmplitudeScalingType, codebookParametersType2AmplitudeScalingTypeConstraints); err != nil {
				return err
			}
			if c.AmplitudeSubsetRestriction != nil {
				if err := e.EncodeEnumerated((*c.AmplitudeSubsetRestriction), codebookParametersType2AmplitudeSubsetRestrictionConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. type2-PortSelection: sequence
	{
		if ie.Type2_PortSelection != nil {
			c := ie.Type2_PortSelection
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType2PortSelectionSupportedCSIRSResourceListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceList {
					if err := c.SupportedCSI_RS_ResourceList[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.ParameterLx, per.Constrained(2, 4)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.AmplitudeScalingType, codebookParametersType2PortSelectionAmplitudeScalingTypeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. type1: sequence
	{
		{
			c := &ie.Type1
			codebookParametersType1Seq := d.NewSequenceDecoder(codebookParametersType1Constraints)
			if err := codebookParametersType1Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.SinglePanel
				{
					seqOf := d.NewSequenceOfDecoder(codebookParametersType1SinglePanelSupportedCSIRSResourceListConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList = make([]SupportedCSI_RS_Resource, n)
					for i := int64(0); i < n; i++ {
						if err := c.SupportedCSI_RS_ResourceList[i].Decode(d); err != nil {
							return err
						}
					}
				}
				{
					v, err := d.DecodeEnumerated(codebookParametersType1SinglePanelModesConstraints)
					if err != nil {
						return err
					}
					c.Modes = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 8))
					if err != nil {
						return err
					}
					c.MaxNumberCSI_RS_PerResourceSet = v
				}
			}
			if codebookParametersType1Seq.IsComponentPresent(1) {
				c.MultiPanel = &struct {
					SupportedCSI_RS_ResourceList   []SupportedCSI_RS_Resource
					Modes                          int64
					NrofPanels                     int64
					MaxNumberCSI_RS_PerResourceSet int64
				}{}
				c := (*c.MultiPanel)
				{
					seqOf := d.NewSequenceOfDecoder(codebookParametersType1MultiPanelSupportedCSIRSResourceListConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList = make([]SupportedCSI_RS_Resource, n)
					for i := int64(0); i < n; i++ {
						if err := c.SupportedCSI_RS_ResourceList[i].Decode(d); err != nil {
							return err
						}
					}
				}
				{
					v, err := d.DecodeEnumerated(codebookParametersType1MultiPanelModesConstraints)
					if err != nil {
						return err
					}
					c.Modes = v
				}
				{
					v, err := d.DecodeEnumerated(codebookParametersType1MultiPanelNrofPanelsConstraints)
					if err != nil {
						return err
					}
					c.NrofPanels = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 8))
					if err != nil {
						return err
					}
					c.MaxNumberCSI_RS_PerResourceSet = v
				}
			}
		}
	}

	// 3. type2: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Type2 = &struct {
				SupportedCSI_RS_ResourceList []SupportedCSI_RS_Resource
				ParameterLx                  int64
				AmplitudeScalingType         int64
				AmplitudeSubsetRestriction   *int64
			}{}
			c := ie.Type2
			codebookParametersType2Seq := d.NewSequenceDecoder(codebookParametersType2Constraints)
			if err := codebookParametersType2Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType2SupportedCSIRSResourceListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList = make([]SupportedCSI_RS_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ResourceList[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.ParameterLx = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersType2AmplitudeScalingTypeConstraints)
				if err != nil {
					return err
				}
				c.AmplitudeScalingType = v
			}
			if codebookParametersType2Seq.IsComponentPresent(3) {
				c.AmplitudeSubsetRestriction = new(int64)
				v, err := d.DecodeEnumerated(codebookParametersType2AmplitudeSubsetRestrictionConstraints)
				if err != nil {
					return err
				}
				(*c.AmplitudeSubsetRestriction) = v
			}
		}
	}

	// 4. type2-PortSelection: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.Type2_PortSelection = &struct {
				SupportedCSI_RS_ResourceList []SupportedCSI_RS_Resource
				ParameterLx                  int64
				AmplitudeScalingType         int64
			}{}
			c := ie.Type2_PortSelection
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType2PortSelectionSupportedCSIRSResourceListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList = make([]SupportedCSI_RS_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ResourceList[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.ParameterLx = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersType2PortSelectionAmplitudeScalingTypeConstraints)
				if err != nil {
					return err
				}
				c.AmplitudeScalingType = v
			}
		}
	}

	return nil
}
