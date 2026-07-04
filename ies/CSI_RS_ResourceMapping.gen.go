// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-RS-ResourceMapping (line 7573).

var cSIRSResourceMappingConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyDomainAllocation"},
		{Name: "nrofPorts"},
		{Name: "firstOFDMSymbolInTimeDomain"},
		{Name: "firstOFDMSymbolInTimeDomain2", Optional: true},
		{Name: "cdm-Type"},
		{Name: "density"},
		{Name: "freqBand"},
	},
}

var cSI_RS_ResourceMappingFrequencyDomainAllocationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "row1"},
		{Name: "row2"},
		{Name: "row4"},
		{Name: "other"},
	},
}

const (
	CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row1  = 0
	CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row2  = 1
	CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row4  = 2
	CSI_RS_ResourceMapping_FrequencyDomainAllocation_Other = 3
)

const (
	CSI_RS_ResourceMapping_NrofPorts_P1  = 0
	CSI_RS_ResourceMapping_NrofPorts_P2  = 1
	CSI_RS_ResourceMapping_NrofPorts_P4  = 2
	CSI_RS_ResourceMapping_NrofPorts_P8  = 3
	CSI_RS_ResourceMapping_NrofPorts_P12 = 4
	CSI_RS_ResourceMapping_NrofPorts_P16 = 5
	CSI_RS_ResourceMapping_NrofPorts_P24 = 6
	CSI_RS_ResourceMapping_NrofPorts_P32 = 7
)

var cSIRSResourceMappingNrofPortsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_RS_ResourceMapping_NrofPorts_P1, CSI_RS_ResourceMapping_NrofPorts_P2, CSI_RS_ResourceMapping_NrofPorts_P4, CSI_RS_ResourceMapping_NrofPorts_P8, CSI_RS_ResourceMapping_NrofPorts_P12, CSI_RS_ResourceMapping_NrofPorts_P16, CSI_RS_ResourceMapping_NrofPorts_P24, CSI_RS_ResourceMapping_NrofPorts_P32},
}

var cSIRSResourceMappingFirstOFDMSymbolInTimeDomainConstraints = per.Constrained(0, 13)

var cSIRSResourceMappingFirstOFDMSymbolInTimeDomain2Constraints = per.Constrained(2, 12)

const (
	CSI_RS_ResourceMapping_Cdm_Type_NoCDM        = 0
	CSI_RS_ResourceMapping_Cdm_Type_Fd_CDM2      = 1
	CSI_RS_ResourceMapping_Cdm_Type_Cdm4_FD2_TD2 = 2
	CSI_RS_ResourceMapping_Cdm_Type_Cdm8_FD2_TD4 = 3
)

var cSIRSResourceMappingCdmTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_RS_ResourceMapping_Cdm_Type_NoCDM, CSI_RS_ResourceMapping_Cdm_Type_Fd_CDM2, CSI_RS_ResourceMapping_Cdm_Type_Cdm4_FD2_TD2, CSI_RS_ResourceMapping_Cdm_Type_Cdm8_FD2_TD4},
}

var cSI_RS_ResourceMappingDensityConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dot5"},
		{Name: "one"},
		{Name: "three"},
		{Name: "spare"},
	},
}

const (
	CSI_RS_ResourceMapping_Density_Dot5  = 0
	CSI_RS_ResourceMapping_Density_One   = 1
	CSI_RS_ResourceMapping_Density_Three = 2
	CSI_RS_ResourceMapping_Density_Spare = 3
)

const (
	CSI_RS_ResourceMapping_Density_Dot5_EvenPRBs = 0
	CSI_RS_ResourceMapping_Density_Dot5_OddPRBs  = 1
)

var cSIRSResourceMappingDensityDot5Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_RS_ResourceMapping_Density_Dot5_EvenPRBs, CSI_RS_ResourceMapping_Density_Dot5_OddPRBs},
}

type CSI_RS_ResourceMapping struct {
	FrequencyDomainAllocation struct {
		Choice int
		Row1   *per.BitString
		Row2   *per.BitString
		Row4   *per.BitString
		Other  *per.BitString
	}
	NrofPorts                    int64
	FirstOFDMSymbolInTimeDomain  int64
	FirstOFDMSymbolInTimeDomain2 *int64
	Cdm_Type                     int64
	Density                      struct {
		Choice int
		Dot5   *int64
		One    *struct{}
		Three  *struct{}
		Spare  *struct{}
	}
	FreqBand CSI_FrequencyOccupation
}

func (ie *CSI_RS_ResourceMapping) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSResourceMappingConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FirstOFDMSymbolInTimeDomain2 != nil}); err != nil {
		return err
	}

	// 3. frequencyDomainAllocation: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_RS_ResourceMappingFrequencyDomainAllocationConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.FrequencyDomainAllocation.Choice), false, nil); err != nil {
			return err
		}
		switch ie.FrequencyDomainAllocation.Choice {
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row1:
			if err := e.EncodeBitString((*ie.FrequencyDomainAllocation.Row1), per.FixedSize(4)); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row2:
			if err := e.EncodeBitString((*ie.FrequencyDomainAllocation.Row2), per.FixedSize(12)); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row4:
			if err := e.EncodeBitString((*ie.FrequencyDomainAllocation.Row4), per.FixedSize(3)); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Other:
			if err := e.EncodeBitString((*ie.FrequencyDomainAllocation.Other), per.FixedSize(6)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.FrequencyDomainAllocation.Choice), Constraint: "index out of range"}
		}
	}

	// 4. nrofPorts: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofPorts, cSIRSResourceMappingNrofPortsConstraints); err != nil {
			return err
		}
	}

	// 5. firstOFDMSymbolInTimeDomain: integer
	{
		if err := e.EncodeInteger(ie.FirstOFDMSymbolInTimeDomain, cSIRSResourceMappingFirstOFDMSymbolInTimeDomainConstraints); err != nil {
			return err
		}
	}

	// 6. firstOFDMSymbolInTimeDomain2: integer
	{
		if ie.FirstOFDMSymbolInTimeDomain2 != nil {
			if err := e.EncodeInteger(*ie.FirstOFDMSymbolInTimeDomain2, cSIRSResourceMappingFirstOFDMSymbolInTimeDomain2Constraints); err != nil {
				return err
			}
		}
	}

	// 7. cdm-Type: enumerated
	{
		if err := e.EncodeEnumerated(ie.Cdm_Type, cSIRSResourceMappingCdmTypeConstraints); err != nil {
			return err
		}
	}

	// 8. density: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_RS_ResourceMappingDensityConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Density.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Density.Choice {
		case CSI_RS_ResourceMapping_Density_Dot5:
			if err := e.EncodeEnumerated((*ie.Density.Dot5), cSIRSResourceMappingDensityDot5Constraints); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_Density_One:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_Density_Three:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_Density_Spare:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Density.Choice), Constraint: "index out of range"}
		}
	}

	// 9. freqBand: ref
	{
		if err := ie.FreqBand.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_RS_ResourceMapping) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSResourceMappingConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. frequencyDomainAllocation: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_RS_ResourceMappingFrequencyDomainAllocationConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.FrequencyDomainAllocation.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row1:
			ie.FrequencyDomainAllocation.Row1 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(4))
			if err != nil {
				return err
			}
			(*ie.FrequencyDomainAllocation.Row1) = v
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row2:
			ie.FrequencyDomainAllocation.Row2 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(12))
			if err != nil {
				return err
			}
			(*ie.FrequencyDomainAllocation.Row2) = v
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Row4:
			ie.FrequencyDomainAllocation.Row4 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(3))
			if err != nil {
				return err
			}
			(*ie.FrequencyDomainAllocation.Row4) = v
		case CSI_RS_ResourceMapping_FrequencyDomainAllocation_Other:
			ie.FrequencyDomainAllocation.Other = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(6))
			if err != nil {
				return err
			}
			(*ie.FrequencyDomainAllocation.Other) = v
		}
	}

	// 4. nrofPorts: enumerated
	{
		v1, err := d.DecodeEnumerated(cSIRSResourceMappingNrofPortsConstraints)
		if err != nil {
			return err
		}
		ie.NrofPorts = v1
	}

	// 5. firstOFDMSymbolInTimeDomain: integer
	{
		v2, err := d.DecodeInteger(cSIRSResourceMappingFirstOFDMSymbolInTimeDomainConstraints)
		if err != nil {
			return err
		}
		ie.FirstOFDMSymbolInTimeDomain = v2
	}

	// 6. firstOFDMSymbolInTimeDomain2: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(cSIRSResourceMappingFirstOFDMSymbolInTimeDomain2Constraints)
			if err != nil {
				return err
			}
			ie.FirstOFDMSymbolInTimeDomain2 = &v
		}
	}

	// 7. cdm-Type: enumerated
	{
		v4, err := d.DecodeEnumerated(cSIRSResourceMappingCdmTypeConstraints)
		if err != nil {
			return err
		}
		ie.Cdm_Type = v4
	}

	// 8. density: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_RS_ResourceMappingDensityConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Density.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_RS_ResourceMapping_Density_Dot5:
			ie.Density.Dot5 = new(int64)
			v, err := d.DecodeEnumerated(cSIRSResourceMappingDensityDot5Constraints)
			if err != nil {
				return err
			}
			(*ie.Density.Dot5) = v
		case CSI_RS_ResourceMapping_Density_One:
			ie.Density.One = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_Density_Three:
			ie.Density.Three = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_RS_ResourceMapping_Density_Spare:
			ie.Density.Spare = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	}

	// 9. freqBand: ref
	{
		if err := ie.FreqBand.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
