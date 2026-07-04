// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-MTC3-r16 (line 15873).

var sSBMTC3R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-r16"},
		{Name: "duration-r16"},
		{Name: "pci-List-r16", Optional: true},
		{Name: "ssb-ToMeasure-r16", Optional: true},
	},
}

var sSB_MTC3_r16PeriodicityAndOffsetR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sf5-r16"},
		{Name: "sf10-r16"},
		{Name: "sf20-r16"},
		{Name: "sf40-r16"},
		{Name: "sf80-r16"},
		{Name: "sf160-r16"},
		{Name: "sf320-r16"},
		{Name: "sf640-r16"},
		{Name: "sf1280-r16"},
	},
}

const (
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf5_r16    = 0
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf10_r16   = 1
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf20_r16   = 2
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf40_r16   = 3
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf80_r16   = 4
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf160_r16  = 5
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf320_r16  = 6
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf640_r16  = 7
	SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf1280_r16 = 8
)

const (
	SSB_MTC3_r16_Duration_r16_Sf1 = 0
	SSB_MTC3_r16_Duration_r16_Sf2 = 1
	SSB_MTC3_r16_Duration_r16_Sf3 = 2
	SSB_MTC3_r16_Duration_r16_Sf4 = 3
	SSB_MTC3_r16_Duration_r16_Sf5 = 4
)

var sSBMTC3R16DurationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_MTC3_r16_Duration_r16_Sf1, SSB_MTC3_r16_Duration_r16_Sf2, SSB_MTC3_r16_Duration_r16_Sf3, SSB_MTC3_r16_Duration_r16_Sf4, SSB_MTC3_r16_Duration_r16_Sf5},
}

var sSBMTC3R16PciListR16Constraints = per.SizeRange(1, common.MaxNrofPCIsPerSMTC)

var sSB_MTC3_r16SsbToMeasureR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SSB_MTC3_r16_Ssb_ToMeasure_r16_Release = 0
	SSB_MTC3_r16_Ssb_ToMeasure_r16_Setup   = 1
)

type SSB_MTC3_r16 struct {
	PeriodicityAndOffset_r16 struct {
		Choice     int
		Sf5_r16    *int64
		Sf10_r16   *int64
		Sf20_r16   *int64
		Sf40_r16   *int64
		Sf80_r16   *int64
		Sf160_r16  *int64
		Sf320_r16  *int64
		Sf640_r16  *int64
		Sf1280_r16 *int64
	}
	Duration_r16      int64
	Pci_List_r16      []PhysCellId
	Ssb_ToMeasure_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SSB_ToMeasure
	}
}

func (ie *SSB_MTC3_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMTC3R16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pci_List_r16 != nil, ie.Ssb_ToMeasure_r16 != nil}); err != nil {
		return err
	}

	// 2. periodicityAndOffset-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sSB_MTC3_r16PeriodicityAndOffsetR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PeriodicityAndOffset_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PeriodicityAndOffset_r16.Choice {
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf5_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf5_r16), per.Constrained(0, 4)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf10_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf10_r16), per.Constrained(0, 9)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf20_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf20_r16), per.Constrained(0, 19)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf40_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf40_r16), per.Constrained(0, 39)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf80_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf80_r16), per.Constrained(0, 79)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf160_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf160_r16), per.Constrained(0, 159)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf320_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf320_r16), per.Constrained(0, 319)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf640_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf640_r16), per.Constrained(0, 639)); err != nil {
				return err
			}
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf1280_r16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r16.Sf1280_r16), per.Constrained(0, 1279)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PeriodicityAndOffset_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 3. duration-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Duration_r16, sSBMTC3R16DurationR16Constraints); err != nil {
			return err
		}
	}

	// 4. pci-List-r16: sequence-of
	{
		if ie.Pci_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sSBMTC3R16PciListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pci_List_r16))); err != nil {
				return err
			}
			for i := range ie.Pci_List_r16 {
				if err := ie.Pci_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. ssb-ToMeasure-r16: choice
	{
		if ie.Ssb_ToMeasure_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sSB_MTC3_r16SsbToMeasureR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_ToMeasure_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ssb_ToMeasure_r16).Choice {
			case SSB_MTC3_r16_Ssb_ToMeasure_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SSB_MTC3_r16_Ssb_ToMeasure_r16_Setup:
				if err := (*ie.Ssb_ToMeasure_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ssb_ToMeasure_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SSB_MTC3_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMTC3R16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. periodicityAndOffset-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sSB_MTC3_r16PeriodicityAndOffsetR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PeriodicityAndOffset_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf5_r16:
			ie.PeriodicityAndOffset_r16.Sf5_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf5_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf10_r16:
			ie.PeriodicityAndOffset_r16.Sf10_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf10_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf20_r16:
			ie.PeriodicityAndOffset_r16.Sf20_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf20_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf40_r16:
			ie.PeriodicityAndOffset_r16.Sf40_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf40_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf80_r16:
			ie.PeriodicityAndOffset_r16.Sf80_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf80_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf160_r16:
			ie.PeriodicityAndOffset_r16.Sf160_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf160_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf320_r16:
			ie.PeriodicityAndOffset_r16.Sf320_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf320_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf640_r16:
			ie.PeriodicityAndOffset_r16.Sf640_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf640_r16) = v
		case SSB_MTC3_r16_PeriodicityAndOffset_r16_Sf1280_r16:
			ie.PeriodicityAndOffset_r16.Sf1280_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r16.Sf1280_r16) = v
		}
	}

	// 3. duration-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sSBMTC3R16DurationR16Constraints)
		if err != nil {
			return err
		}
		ie.Duration_r16 = v1
	}

	// 4. pci-List-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sSBMTC3R16PciListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pci_List_r16 = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pci_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. ssb-ToMeasure-r16: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Ssb_ToMeasure_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SSB_ToMeasure
			}{}
			choiceDec := d.NewChoiceDecoder(sSB_MTC3_r16SsbToMeasureR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_ToMeasure_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SSB_MTC3_r16_Ssb_ToMeasure_r16_Release:
				(*ie.Ssb_ToMeasure_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SSB_MTC3_r16_Ssb_ToMeasure_r16_Setup:
				(*ie.Ssb_ToMeasure_r16).Setup = new(SSB_ToMeasure)
				if err := (*ie.Ssb_ToMeasure_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
