// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-ReportSubConfig-v1900 (line 7346).

var cSIReportSubConfigV1900Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSubConfigParams-v1900"},
	},
}

var cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "portSubsetIndicator-v1900", Optional: true},
	},
}

var cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900PortSubsetIndicatorV1900Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "p48"},
		{Name: "p64"},
		{Name: "p128"},
	},
}

const (
	CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P48  = 0
	CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P64  = 1
	CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P128 = 2
)

type CSI_ReportSubConfig_v1900 struct {
	ReportSubConfigParams_v1900 struct {
		A1_Parameters_v1900 struct {
			PortSubsetIndicator_v1900 *struct {
				Choice int
				P48    *per.BitString
				P64    *per.BitString
				P128   *per.BitString
			}
		}
	}
}

func (ie *CSI_ReportSubConfig_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportSubConfigV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. reportSubConfigParams-v1900: sequence
	{
		{
			c := &ie.ReportSubConfigParams_v1900
			{
				c := &c.A1_Parameters_v1900
				cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Seq := e.NewSequenceEncoder(cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Constraints)
				if err := cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Seq.EncodePreamble([]bool{c.PortSubsetIndicator_v1900 != nil}); err != nil {
					return err
				}
				if c.PortSubsetIndicator_v1900 != nil {
					choiceEnc := e.NewChoiceEncoder(cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900PortSubsetIndicatorV1900Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.PortSubsetIndicator_v1900).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.PortSubsetIndicator_v1900).Choice {
					case CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P48:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_v1900).P48), per.FixedSize(48)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P64:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_v1900).P64), per.FixedSize(64)); err != nil {
							return err
						}
					case CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P128:
						if err := e.EncodeBitString((*(*c.PortSubsetIndicator_v1900).P128), per.FixedSize(128)); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (ie *CSI_ReportSubConfig_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportSubConfigV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. reportSubConfigParams-v1900: sequence
	{
		{
			c := &ie.ReportSubConfigParams_v1900
			{
				c := &c.A1_Parameters_v1900
				cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Seq := d.NewSequenceDecoder(cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Constraints)
				if err := cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Seq.DecodePreamble(); err != nil {
					return err
				}
				if cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900Seq.IsComponentPresent(0) {
					c.PortSubsetIndicator_v1900 = &struct {
						Choice int
						P48    *per.BitString
						P64    *per.BitString
						P128   *per.BitString
					}{}
					choiceDec := d.NewChoiceDecoder(cSIReportSubConfigV1900ReportSubConfigParamsV1900A1ParametersV1900PortSubsetIndicatorV1900Constraints)
					index, _, _, err := choiceDec.DecodeChoice()
					if err != nil {
						return err
					}
					(*c.PortSubsetIndicator_v1900).Choice = int(index)
					switch index {
					case CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P48:
						(*c.PortSubsetIndicator_v1900).P48 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(48))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_v1900).P48) = v
					case CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P64:
						(*c.PortSubsetIndicator_v1900).P64 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(64))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_v1900).P64) = v
					case CSI_ReportSubConfig_v1900_ReportSubConfigParams_v1900_A1_Parameters_v1900_PortSubsetIndicator_v1900_P128:
						(*c.PortSubsetIndicator_v1900).P128 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(128))
						if err != nil {
							return err
						}
						(*(*c.PortSubsetIndicator_v1900).P128) = v
					}
				}
			}
		}
	}

	return nil
}
