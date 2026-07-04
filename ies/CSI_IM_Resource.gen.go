// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-IM-Resource (line 6940).

var cSIIMResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-IM-ResourceId"},
		{Name: "csi-IM-ResourceElementPattern", Optional: true},
		{Name: "freqBand", Optional: true},
		{Name: "periodicityAndOffset", Optional: true},
	},
}

var cSI_IM_ResourceCsiIMResourceElementPatternConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pattern0"},
		{Name: "pattern1"},
	},
}

const (
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0 = 0
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1 = 1
)

const (
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S0  = 0
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S2  = 1
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S4  = 2
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S6  = 3
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S8  = 4
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S10 = 5
)

var cSIIMResourceCsiIMResourceElementPatternPattern0SubcarrierLocationP0Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S0, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S2, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S4, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S6, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S8, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0_SubcarrierLocation_P0_S10},
}

const (
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1_SubcarrierLocation_P1_S0 = 0
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1_SubcarrierLocation_P1_S4 = 1
	CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1_SubcarrierLocation_P1_S8 = 2
)

var cSIIMResourceCsiIMResourceElementPatternPattern1SubcarrierLocationP1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1_SubcarrierLocation_P1_S0, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1_SubcarrierLocation_P1_S4, CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1_SubcarrierLocation_P1_S8},
}

type CSI_IM_Resource struct {
	Csi_IM_ResourceId             CSI_IM_ResourceId
	Csi_IM_ResourceElementPattern *struct {
		Choice   int
		Pattern0 *struct {
			SubcarrierLocation_P0 int64
			SymbolLocation_P0     int64
		}
		Pattern1 *struct {
			SubcarrierLocation_P1 int64
			SymbolLocation_P1     int64
		}
	}
	FreqBand             *CSI_FrequencyOccupation
	PeriodicityAndOffset *CSI_ResourcePeriodicityAndOffset
}

func (ie *CSI_IM_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIIMResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_IM_ResourceElementPattern != nil, ie.FreqBand != nil, ie.PeriodicityAndOffset != nil}); err != nil {
		return err
	}

	// 3. csi-IM-ResourceId: ref
	{
		if err := ie.Csi_IM_ResourceId.Encode(e); err != nil {
			return err
		}
	}

	// 4. csi-IM-ResourceElementPattern: choice
	{
		if ie.Csi_IM_ResourceElementPattern != nil {
			choiceEnc := e.NewChoiceEncoder(cSI_IM_ResourceCsiIMResourceElementPatternConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Csi_IM_ResourceElementPattern).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Csi_IM_ResourceElementPattern).Choice {
			case CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0:
				c := (*(*ie.Csi_IM_ResourceElementPattern).Pattern0)
				if err := e.EncodeEnumerated(c.SubcarrierLocation_P0, cSIIMResourceCsiIMResourceElementPatternPattern0SubcarrierLocationP0Constraints); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.SymbolLocation_P0, per.Constrained(0, 12)); err != nil {
					return err
				}
			case CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1:
				c := (*(*ie.Csi_IM_ResourceElementPattern).Pattern1)
				if err := e.EncodeEnumerated(c.SubcarrierLocation_P1, cSIIMResourceCsiIMResourceElementPatternPattern1SubcarrierLocationP1Constraints); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.SymbolLocation_P1, per.Constrained(0, 13)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Csi_IM_ResourceElementPattern).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. freqBand: ref
	{
		if ie.FreqBand != nil {
			if err := ie.FreqBand.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. periodicityAndOffset: ref
	{
		if ie.PeriodicityAndOffset != nil {
			if err := ie.PeriodicityAndOffset.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_IM_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIIMResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-IM-ResourceId: ref
	{
		if err := ie.Csi_IM_ResourceId.Decode(d); err != nil {
			return err
		}
	}

	// 4. csi-IM-ResourceElementPattern: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Csi_IM_ResourceElementPattern = &struct {
				Choice   int
				Pattern0 *struct {
					SubcarrierLocation_P0 int64
					SymbolLocation_P0     int64
				}
				Pattern1 *struct {
					SubcarrierLocation_P1 int64
					SymbolLocation_P1     int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(cSI_IM_ResourceCsiIMResourceElementPatternConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Csi_IM_ResourceElementPattern).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern0:
				(*ie.Csi_IM_ResourceElementPattern).Pattern0 = &struct {
					SubcarrierLocation_P0 int64
					SymbolLocation_P0     int64
				}{}
				c := (*(*ie.Csi_IM_ResourceElementPattern).Pattern0)
				{
					v, err := d.DecodeEnumerated(cSIIMResourceCsiIMResourceElementPatternPattern0SubcarrierLocationP0Constraints)
					if err != nil {
						return err
					}
					c.SubcarrierLocation_P0 = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(0, 12))
					if err != nil {
						return err
					}
					c.SymbolLocation_P0 = v
				}
			case CSI_IM_Resource_Csi_IM_ResourceElementPattern_Pattern1:
				(*ie.Csi_IM_ResourceElementPattern).Pattern1 = &struct {
					SubcarrierLocation_P1 int64
					SymbolLocation_P1     int64
				}{}
				c := (*(*ie.Csi_IM_ResourceElementPattern).Pattern1)
				{
					v, err := d.DecodeEnumerated(cSIIMResourceCsiIMResourceElementPatternPattern1SubcarrierLocationP1Constraints)
					if err != nil {
						return err
					}
					c.SubcarrierLocation_P1 = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(0, 13))
					if err != nil {
						return err
					}
					c.SymbolLocation_P1 = v
				}
			}
		}
	}

	// 5. freqBand: ref
	{
		if seq.IsComponentPresent(2) {
			ie.FreqBand = new(CSI_FrequencyOccupation)
			if err := ie.FreqBand.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. periodicityAndOffset: ref
	{
		if seq.IsComponentPresent(3) {
			ie.PeriodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
			if err := ie.PeriodicityAndOffset.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
