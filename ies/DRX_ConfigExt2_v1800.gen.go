// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-ConfigExt2-v1800 (line 8169).

var dRXConfigExt2V1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-NonIntegerLongCycleStartOffset-r18"},
		{Name: "shortDRX-r18", Optional: true},
		{Name: "drx-TimeReferenceSFN-r18", Optional: true},
	},
}

var dRX_ConfigExt2_v1800DrxNonIntegerLongCycleStartOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms1001over240"},
		{Name: "ms25over6"},
		{Name: "ms25over3"},
		{Name: "ms1001over120"},
		{Name: "ms100over9"},
		{Name: "ms25over2"},
		{Name: "ms40over3"},
		{Name: "ms125over9"},
		{Name: "ms50over3"},
		{Name: "ms1001over60"},
		{Name: "ms125over6"},
		{Name: "ms200over9"},
		{Name: "ms250over9"},
		{Name: "ms100over3"},
		{Name: "ms1001over30"},
		{Name: "ms75over2"},
		{Name: "ms125over3"},
		{Name: "ms1001over24"},
		{Name: "ms200over3"},
		{Name: "ms1001over15"},
		{Name: "ms250over3"},
		{Name: "ms1001over12"},
		{Name: "ms400over3"},
	},
}

const (
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over240 = 0
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over6     = 1
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over3     = 2
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over120 = 3
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms100over9    = 4
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over2     = 5
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms40over3     = 6
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over9    = 7
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms50over3     = 8
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over60  = 9
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over6    = 10
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms200over9    = 11
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms250over9    = 12
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms100over3    = 13
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over30  = 14
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms75over2     = 15
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over3    = 16
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over24  = 17
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms200over3    = 18
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over15  = 19
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms250over3    = 20
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over12  = 21
	DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms400over3    = 22
)

const (
	DRX_ConfigExt2_v1800_Drx_TimeReferenceSFN_r18_Sfn512 = 0
)

var dRXConfigExt2V1800DrxTimeReferenceSFNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigExt2_v1800_Drx_TimeReferenceSFN_r18_Sfn512},
}

const (
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over240 = 0
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms25over6     = 1
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms25over3     = 2
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over120 = 3
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms100over9    = 4
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms25over2     = 5
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms40over3     = 6
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms125over9    = 7
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms50over3     = 8
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over60  = 9
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms125over6    = 10
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms200over9    = 11
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms100over3    = 12
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over30  = 13
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms125over3    = 14
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over24  = 15
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms200over3    = 16
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare15       = 17
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare14       = 18
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare13       = 19
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare12       = 20
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare11       = 21
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare10       = 22
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare9        = 23
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare8        = 24
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare7        = 25
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare6        = 26
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare5        = 27
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare4        = 28
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare3        = 29
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare2        = 30
	DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare1        = 31
)

var dRXConfigExt2V1800ShortDRXR18DrxNonIntegerShortCycleR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over240, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms25over6, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms25over3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over120, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms100over9, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms25over2, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms40over3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms125over9, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms50over3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over60, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms125over6, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms200over9, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms100over3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over30, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms125over3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms1001over24, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Ms200over3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare15, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare14, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare13, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare12, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare11, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare10, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare9, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare8, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare7, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare6, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare5, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare4, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare3, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare2, DRX_ConfigExt2_v1800_ShortDRX_r18_Drx_NonIntegerShortCycle_r18_Spare1},
}

type DRX_ConfigExt2_v1800 struct {
	Drx_NonIntegerLongCycleStartOffset_r18 struct {
		Choice        int
		Ms1001over240 *int64
		Ms25over6     *int64
		Ms25over3     *int64
		Ms1001over120 *int64
		Ms100over9    *int64
		Ms25over2     *int64
		Ms40over3     *int64
		Ms125over9    *int64
		Ms50over3     *int64
		Ms1001over60  *int64
		Ms125over6    *int64
		Ms200over9    *int64
		Ms250over9    *int64
		Ms100over3    *int64
		Ms1001over30  *int64
		Ms75over2     *int64
		Ms125over3    *int64
		Ms1001over24  *int64
		Ms200over3    *int64
		Ms1001over15  *int64
		Ms250over3    *int64
		Ms1001over12  *int64
		Ms400over3    *int64
	}
	ShortDRX_r18 *struct {
		Drx_NonIntegerShortCycle_r18 int64
		Drx_ShortCycleTimer_r18      int64
	}
	Drx_TimeReferenceSFN_r18 *int64
}

func (ie *DRX_ConfigExt2_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXConfigExt2V1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ShortDRX_r18 != nil, ie.Drx_TimeReferenceSFN_r18 != nil}); err != nil {
		return err
	}

	// 2. drx-NonIntegerLongCycleStartOffset-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(dRX_ConfigExt2_v1800DrxNonIntegerLongCycleStartOffsetR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Drx_NonIntegerLongCycleStartOffset_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Drx_NonIntegerLongCycleStartOffset_r18.Choice {
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over240:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over240), per.Constrained(0, 3)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over6:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over6), per.Constrained(0, 3)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over3), per.Constrained(0, 7)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over120:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over120), per.Constrained(0, 7)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms100over9:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms100over9), per.Constrained(0, 10)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over2:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over2), per.Constrained(0, 11)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms40over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms40over3), per.Constrained(0, 12)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over9:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over9), per.Constrained(0, 12)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms50over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms50over3), per.Constrained(0, 15)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over60:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over60), per.Constrained(0, 15)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over6:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over6), per.Constrained(0, 19)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms200over9:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms200over9), per.Constrained(0, 21)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms250over9:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms250over9), per.Constrained(0, 26)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms100over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms100over3), per.Constrained(0, 32)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over30:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over30), per.Constrained(0, 32)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms75over2:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms75over2), per.Constrained(0, 36)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over3), per.Constrained(0, 40)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over24:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over24), per.Constrained(0, 40)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms200over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms200over3), per.Constrained(0, 65)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over15:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over15), per.Constrained(0, 65)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms250over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms250over3), per.Constrained(0, 82)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over12:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over12), per.Constrained(0, 82)); err != nil {
				return err
			}
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms400over3:
			if err := e.EncodeInteger((*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms400over3), per.Constrained(0, 132)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Drx_NonIntegerLongCycleStartOffset_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. shortDRX-r18: sequence
	{
		if ie.ShortDRX_r18 != nil {
			c := ie.ShortDRX_r18
			if err := e.EncodeEnumerated(c.Drx_NonIntegerShortCycle_r18, dRXConfigExt2V1800ShortDRXR18DrxNonIntegerShortCycleR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Drx_ShortCycleTimer_r18, per.Constrained(1, 16)); err != nil {
				return err
			}
		}
	}

	// 4. drx-TimeReferenceSFN-r18: enumerated
	{
		if ie.Drx_TimeReferenceSFN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Drx_TimeReferenceSFN_r18, dRXConfigExt2V1800DrxTimeReferenceSFNR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DRX_ConfigExt2_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXConfigExt2V1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. drx-NonIntegerLongCycleStartOffset-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(dRX_ConfigExt2_v1800DrxNonIntegerLongCycleStartOffsetR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Drx_NonIntegerLongCycleStartOffset_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over240:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over240) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over6:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over6 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over6) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over120:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over120) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms100over9:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms100over9 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms100over9) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms25over2:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over2 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 11))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms25over2) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms40over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms40over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 12))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms40over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over9:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over9 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 12))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over9) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms50over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms50over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms50over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over60:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over60 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over60) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over6:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over6 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over6) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms200over9:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms200over9 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 21))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms200over9) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms250over9:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms250over9 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 26))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms250over9) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms100over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms100over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 32))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms100over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over30:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over30 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 32))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over30) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms75over2:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms75over2 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 36))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms75over2) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms125over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 40))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms125over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over24:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over24 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 40))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over24) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms200over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms200over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 65))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms200over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over15:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over15 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 65))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over15) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms250over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms250over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 82))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms250over3) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms1001over12:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over12 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 82))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms1001over12) = v
		case DRX_ConfigExt2_v1800_Drx_NonIntegerLongCycleStartOffset_r18_Ms400over3:
			ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms400over3 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 132))
			if err != nil {
				return err
			}
			(*ie.Drx_NonIntegerLongCycleStartOffset_r18.Ms400over3) = v
		}
	}

	// 3. shortDRX-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.ShortDRX_r18 = &struct {
				Drx_NonIntegerShortCycle_r18 int64
				Drx_ShortCycleTimer_r18      int64
			}{}
			c := ie.ShortDRX_r18
			{
				v, err := d.DecodeEnumerated(dRXConfigExt2V1800ShortDRXR18DrxNonIntegerShortCycleR18Constraints)
				if err != nil {
					return err
				}
				c.Drx_NonIntegerShortCycle_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.Drx_ShortCycleTimer_r18 = v
			}
		}
	}

	// 4. drx-TimeReferenceSFN-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dRXConfigExt2V1800DrxTimeReferenceSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.Drx_TimeReferenceSFN_r18 = &idx
		}
	}

	return nil
}
