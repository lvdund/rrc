// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-InitiatedResource-r19 (line 7407).

var uEInitiatedResourceR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-r19"},
		{Name: "resource-r19"},
		{Name: "ul-BWP-Id-r19"},
	},
}

var uE_InitiatedResource_r19PeriodicityAndOffsetR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sym2"},
		{Name: "sym6or7"},
		{Name: "sl1"},
		{Name: "sl2"},
		{Name: "sl4"},
		{Name: "sl5"},
		{Name: "sl8"},
		{Name: "sl10"},
		{Name: "sl16"},
		{Name: "sl20"},
		{Name: "sl40"},
		{Name: "sl80"},
		{Name: "sl160"},
		{Name: "sl320"},
		{Name: "sl640"},
	},
}

const (
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sym2    = 0
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sym6or7 = 1
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl1     = 2
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl2     = 3
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl4     = 4
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl5     = 5
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl8     = 6
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl10    = 7
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl16    = 8
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl20    = 9
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl40    = 10
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl80    = 11
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl160   = 12
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl320   = 13
	UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl640   = 14
)

type UE_InitiatedResource_r19 struct {
	PeriodicityAndOffset_r19 struct {
		Choice  int
		Sym2    *struct{}
		Sym6or7 *struct{}
		Sl1     *struct{}
		Sl2     *int64
		Sl4     *int64
		Sl5     *int64
		Sl8     *int64
		Sl10    *int64
		Sl16    *int64
		Sl20    *int64
		Sl40    *int64
		Sl80    *int64
		Sl160   *int64
		Sl320   *int64
		Sl640   *int64
	}
	Resource_r19  PUCCH_ResourceId
	Ul_BWP_Id_r19 BWP_Id
}

func (ie *UE_InitiatedResource_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInitiatedResourceR19Constraints)
	_ = seq

	// 1. periodicityAndOffset-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(uE_InitiatedResource_r19PeriodicityAndOffsetR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PeriodicityAndOffset_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PeriodicityAndOffset_r19.Choice {
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sym2:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sym6or7:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl2:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl2), per.Constrained(0, 1)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl4:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl4), per.Constrained(0, 3)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl5:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl5), per.Constrained(0, 4)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl8:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl8), per.Constrained(0, 7)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl10:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl16:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl16), per.Constrained(0, 15)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl20:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl40:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl80:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl160:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl320:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl640:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r19.Sl640), per.Constrained(0, 639)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PeriodicityAndOffset_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 2. resource-r19: ref
	{
		if err := ie.Resource_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. ul-BWP-Id-r19: ref
	{
		if err := ie.Ul_BWP_Id_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UE_InitiatedResource_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInitiatedResourceR19Constraints)
	_ = seq

	// 1. periodicityAndOffset-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(uE_InitiatedResource_r19PeriodicityAndOffsetR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PeriodicityAndOffset_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sym2:
			ie.PeriodicityAndOffset_r19.Sym2 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sym6or7:
			ie.PeriodicityAndOffset_r19.Sym6or7 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl1:
			ie.PeriodicityAndOffset_r19.Sl1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl2:
			ie.PeriodicityAndOffset_r19.Sl2 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl2) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl4:
			ie.PeriodicityAndOffset_r19.Sl4 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl4) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl5:
			ie.PeriodicityAndOffset_r19.Sl5 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl5) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl8:
			ie.PeriodicityAndOffset_r19.Sl8 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl8) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl10:
			ie.PeriodicityAndOffset_r19.Sl10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl10) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl16:
			ie.PeriodicityAndOffset_r19.Sl16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl16) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl20:
			ie.PeriodicityAndOffset_r19.Sl20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl20) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl40:
			ie.PeriodicityAndOffset_r19.Sl40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl40) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl80:
			ie.PeriodicityAndOffset_r19.Sl80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl80) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl160:
			ie.PeriodicityAndOffset_r19.Sl160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl160) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl320:
			ie.PeriodicityAndOffset_r19.Sl320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl320) = v
		case UE_InitiatedResource_r19_PeriodicityAndOffset_r19_Sl640:
			ie.PeriodicityAndOffset_r19.Sl640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r19.Sl640) = v
		}
	}

	// 2. resource-r19: ref
	{
		if err := ie.Resource_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. ul-BWP-Id-r19: ref
	{
		if err := ie.Ul_BWP_Id_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
