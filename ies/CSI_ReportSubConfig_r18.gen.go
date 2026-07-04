// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ReportSubConfig-r18 (line 7304).

var cSIReportSubConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSubConfigId-r18"},
		{Name: "reportSubConfigParams-r18", Optional: true},
		{Name: "powerOffset-r18", Optional: true},
	},
}

var cSI_ReportSubConfig_r18ReportSubConfigParamsR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "a1-parameters"},
		{Name: "a2-parameters"},
	},
}

const (
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters = 0
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A2_Parameters = 1
)

var cSIReportSubConfigR18PowerOffsetR18Constraints = per.Constrained(0, 23)

var cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookSubConfig-r18", Optional: true},
		{Name: "portSubsetIndicator-r18", Optional: true},
		{Name: "non-PMI-PortIndication-r18", Optional: true},
	},
}

var cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersPortSubsetIndicatorR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "p2"},
		{Name: "p4"},
		{Name: "p8"},
		{Name: "p12"},
		{Name: "p16"},
		{Name: "p24"},
		{Name: "p32"},
	},
}

const (
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P2  = 0
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P4  = 1
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P8  = 2
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P12 = 3
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P16 = 4
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P24 = 5
	CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P32 = 6
)

var cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersNonPMIPortIndicationR18Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourcesPerConfig)

var cSIReportSubConfigR18ReportSubConfigParamsR18A2ParametersNzpCSIRSResourceListR18Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourcesPerSet)

type CSI_ReportSubConfig_r18 struct {
	ReportSubConfigId_r18     CSI_ReportSubConfigId_r18
	ReportSubConfigParams_r18 *struct {
		Choice        int
		A1_Parameters *struct {
			CodebookSubConfig_r18   *CodebookConfig
			PortSubsetIndicator_r18 *struct {
				Choice int
				P2     *per.BitString
				P4     *per.BitString
				P8     *per.BitString
				P12    *per.BitString
				P16    *per.BitString
				P24    *per.BitString
				P32    *per.BitString
			}
			Non_PMI_PortIndication_r18 []PortIndexFor8Ranks
		}
		A2_Parameters *struct {
			Nzp_CSI_RS_ResourceList_r18 []NZP_CSI_RS_ResourceIndex_r18
		}
	}
	PowerOffset_r18 *int64
}

func (ie *CSI_ReportSubConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportSubConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReportSubConfigParams_r18 != nil, ie.PowerOffset_r18 != nil}); err != nil {
		return err
	}

	// 2. reportSubConfigId-r18: ref
	{
		if err := ie.ReportSubConfigId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. reportSubConfigParams-r18: choice
	{
		if ie.ReportSubConfigParams_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(cSI_ReportSubConfig_r18ReportSubConfigParamsR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ReportSubConfigParams_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ReportSubConfigParams_r18).Choice {
			case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters:
				c := (*(*ie.ReportSubConfigParams_r18).A1_Parameters)
				cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq := e.NewSequenceEncoder(cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersConstraints)
				if err := cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq.EncodePreamble([]bool{c.CodebookSubConfig_r18 != nil, c.PortSubsetIndicator_r18 != nil, c.Non_PMI_PortIndication_r18 != nil}); err != nil {
					return err
				}
				if c.CodebookSubConfig_r18 != nil {
					if err := c.CodebookSubConfig_r18.Encode(e); err != nil {
						return err
					}
				}
				if c.PortSubsetIndicator_r18 != nil {
					choiceEnc := e.NewChoiceEncoder(cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersPortSubsetIndicatorR18Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.PortSubsetIndicator_r18).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.PortSubsetIndicator_r18).Choice {
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P2:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P2), per.FixedSize(2)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P4:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P4), per.FixedSize(4)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P8:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P8), per.FixedSize(8)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P12:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P12), per.FixedSize(12)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P16:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P16), per.FixedSize(16)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P24:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P24), per.FixedSize(24)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P32:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_r18).P32), per.FixedSize(32)); err != nil {
							return err
						}
					}
				}
				if c.Non_PMI_PortIndication_r18 != nil {
					seqOf := e.NewSequenceOfEncoder(cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersNonPMIPortIndicationR18Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Non_PMI_PortIndication_r18))); err != nil {
						return err
					}
					for i := range c.Non_PMI_PortIndication_r18 {
						if err := c.Non_PMI_PortIndication_r18[i].Encode(e); err != nil {
							return err
						}
					}
				}
			case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A2_Parameters:
				c := (*(*ie.ReportSubConfigParams_r18).A2_Parameters)
				{
					seqOf := e.NewSequenceOfEncoder(cSIReportSubConfigR18ReportSubConfigParamsR18A2ParametersNzpCSIRSResourceListR18Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Nzp_CSI_RS_ResourceList_r18))); err != nil {
						return err
					}
					for i := range c.Nzp_CSI_RS_ResourceList_r18 {
						if err := c.Nzp_CSI_RS_ResourceList_r18[i].Encode(e); err != nil {
							return err
						}
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ReportSubConfigParams_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. powerOffset-r18: integer
	{
		if ie.PowerOffset_r18 != nil {
			if err := e.EncodeInteger(*ie.PowerOffset_r18, cSIReportSubConfigR18PowerOffsetR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_ReportSubConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportSubConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reportSubConfigId-r18: ref
	{
		if err := ie.ReportSubConfigId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. reportSubConfigParams-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.ReportSubConfigParams_r18 = &struct {
				Choice        int
				A1_Parameters *struct {
					CodebookSubConfig_r18   *CodebookConfig
					PortSubsetIndicator_r18 *struct {
						Choice int
						P2     *per.BitString
						P4     *per.BitString
						P8     *per.BitString
						P12    *per.BitString
						P16    *per.BitString
						P24    *per.BitString
						P32    *per.BitString
					}
					Non_PMI_PortIndication_r18 []PortIndexFor8Ranks
				}
				A2_Parameters *struct {
					Nzp_CSI_RS_ResourceList_r18 []NZP_CSI_RS_ResourceIndex_r18
				}
			}{}
			choiceDec := d.NewChoiceDecoder(cSI_ReportSubConfig_r18ReportSubConfigParamsR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReportSubConfigParams_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters:
				(*ie.ReportSubConfigParams_r18).A1_Parameters = &struct {
					CodebookSubConfig_r18   *CodebookConfig
					PortSubsetIndicator_r18 *struct {
						Choice int
						P2     *per.BitString
						P4     *per.BitString
						P8     *per.BitString
						P12    *per.BitString
						P16    *per.BitString
						P24    *per.BitString
						P32    *per.BitString
					}
					Non_PMI_PortIndication_r18 []PortIndexFor8Ranks
				}{}
				c := (*(*ie.ReportSubConfigParams_r18).A1_Parameters)
				cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq := d.NewSequenceDecoder(cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersConstraints)
				if err := cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq.DecodePreamble(); err != nil {
					return err
				}
				if cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq.IsComponentPresent(0) {
					c.CodebookSubConfig_r18 = new(CodebookConfig)
					if err := (*c.CodebookSubConfig_r18).Decode(d); err != nil {
						return err
					}
				}
				if cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq.IsComponentPresent(1) {
					c.PortSubsetIndicator_r18 = &struct {
						Choice int
						P2     *per.BitString
						P4     *per.BitString
						P8     *per.BitString
						P12    *per.BitString
						P16    *per.BitString
						P24    *per.BitString
						P32    *per.BitString
					}{}
					choiceDec := d.NewChoiceDecoder(cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersPortSubsetIndicatorR18Constraints)
					index, _, _, err := choiceDec.DecodeChoice()
					if err != nil {
						return err
					}
					(*c.PortSubsetIndicator_r18).Choice = int(index)
					switch index {
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P2:
						(*c.PortSubsetIndicator_r18).P2 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(2))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P2) = v
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P4:
						(*c.PortSubsetIndicator_r18).P4 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(4))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P4) = v
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P8:
						(*c.PortSubsetIndicator_r18).P8 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(8))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P8) = v
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P12:
						(*c.PortSubsetIndicator_r18).P12 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(12))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P12) = v
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P16:
						(*c.PortSubsetIndicator_r18).P16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P16) = v
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P24:
						(*c.PortSubsetIndicator_r18).P24 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(24))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P24) = v
					case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A1_Parameters_PortSubsetIndicator_r18_P32:
						(*c.PortSubsetIndicator_r18).P32 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(32))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_r18).P32) = v
					}
				}
				if cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersSeq.IsComponentPresent(2) {
					seqOf := d.NewSequenceOfDecoder(cSIReportSubConfigR18ReportSubConfigParamsR18A1ParametersNonPMIPortIndicationR18Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Non_PMI_PortIndication_r18 = make([]PortIndexFor8Ranks, n)
					for i := int64(0); i < n; i++ {
						if err := c.Non_PMI_PortIndication_r18[i].Decode(d); err != nil {
							return err
						}
					}
				}
			case CSI_ReportSubConfig_r18_ReportSubConfigParams_r18_A2_Parameters:
				(*ie.ReportSubConfigParams_r18).A2_Parameters = &struct {
					Nzp_CSI_RS_ResourceList_r18 []NZP_CSI_RS_ResourceIndex_r18
				}{}
				c := (*(*ie.ReportSubConfigParams_r18).A2_Parameters)
				{
					seqOf := d.NewSequenceOfDecoder(cSIReportSubConfigR18ReportSubConfigParamsR18A2ParametersNzpCSIRSResourceListR18Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Nzp_CSI_RS_ResourceList_r18 = make([]NZP_CSI_RS_ResourceIndex_r18, n)
					for i := int64(0); i < n; i++ {
						if err := c.Nzp_CSI_RS_ResourceList_r18[i].Decode(d); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	// 4. powerOffset-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cSIReportSubConfigR18PowerOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.PowerOffset_r18 = &v
		}
	}

	return nil
}
