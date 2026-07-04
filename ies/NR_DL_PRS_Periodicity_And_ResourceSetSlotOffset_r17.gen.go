// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-DL-PRS-Periodicity-and-ResourceSetSlotOffset-r17 (line 10618).

var nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scs15-r17"},
		{Name: "scs30-r17"},
		{Name: "scs60-r17"},
		{Name: "scs120-r17"},
	},
}

const (
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17  = 0
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17  = 1
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17  = 2
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17 = 3
)

var nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs15R17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n4-r17"},
		{Name: "n5-r17"},
		{Name: "n8-r17"},
		{Name: "n10-r17"},
		{Name: "n16-r17"},
		{Name: "n20-r17"},
		{Name: "n32-r17"},
		{Name: "n40-r17"},
		{Name: "n64-r17"},
		{Name: "n80-r17"},
		{Name: "n160-r17"},
		{Name: "n320-r17"},
		{Name: "n640-r17"},
		{Name: "n1280-r17"},
		{Name: "n2560-r17"},
		{Name: "n5120-r17"},
		{Name: "n10240-r17"},
	},
}

const (
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N4_r17     = 0
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N5_r17     = 1
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N8_r17     = 2
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N10_r17    = 3
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N16_r17    = 4
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N20_r17    = 5
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N32_r17    = 6
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N40_r17    = 7
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N64_r17    = 8
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N80_r17    = 9
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N160_r17   = 10
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N320_r17   = 11
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N640_r17   = 12
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N1280_r17  = 13
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N2560_r17  = 14
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N5120_r17  = 15
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N10240_r17 = 16
)

var nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs30R17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n8-r17"},
		{Name: "n10-r17"},
		{Name: "n16-r17"},
		{Name: "n20-r17"},
		{Name: "n32-r17"},
		{Name: "n40-r17"},
		{Name: "n64-r17"},
		{Name: "n80-r17"},
		{Name: "n128-r17"},
		{Name: "n160-r17"},
		{Name: "n320-r17"},
		{Name: "n640-r17"},
		{Name: "n1280-r17"},
		{Name: "n2560-r17"},
		{Name: "n5120-r17"},
		{Name: "n10240-r17"},
		{Name: "n20480-r17"},
	},
}

const (
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N8_r17     = 0
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N10_r17    = 1
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N16_r17    = 2
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N20_r17    = 3
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N32_r17    = 4
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N40_r17    = 5
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N64_r17    = 6
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N80_r17    = 7
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N128_r17   = 8
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N160_r17   = 9
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N320_r17   = 10
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N640_r17   = 11
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N1280_r17  = 12
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N2560_r17  = 13
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N5120_r17  = 14
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N10240_r17 = 15
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N20480_r17 = 16
)

var nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs60R17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n16-r17"},
		{Name: "n20-r17"},
		{Name: "n32-r17"},
		{Name: "n40-r17"},
		{Name: "n64-r17"},
		{Name: "n80-r17"},
		{Name: "n128-r17"},
		{Name: "n160-r17"},
		{Name: "n256-r17"},
		{Name: "n320-r17"},
		{Name: "n640-r17"},
		{Name: "n1280-r17"},
		{Name: "n2560-r17"},
		{Name: "n5120-r17"},
		{Name: "n10240-r17"},
		{Name: "n20480-r17"},
		{Name: "n40960-r17"},
	},
}

const (
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N16_r17    = 0
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N20_r17    = 1
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N32_r17    = 2
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N40_r17    = 3
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N64_r17    = 4
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N80_r17    = 5
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N128_r17   = 6
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N160_r17   = 7
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N256_r17   = 8
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N320_r17   = 9
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N640_r17   = 10
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N1280_r17  = 11
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N2560_r17  = 12
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N5120_r17  = 13
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N10240_r17 = 14
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N20480_r17 = 15
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N40960_r17 = 16
)

var nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs120R17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n32-r17"},
		{Name: "n40-r17"},
		{Name: "n64-r17"},
		{Name: "n80-r17"},
		{Name: "n128-r17"},
		{Name: "n160-r17"},
		{Name: "n256-r17"},
		{Name: "n320-r17"},
		{Name: "n512-r17"},
		{Name: "n640-r17"},
		{Name: "n1280-r17"},
		{Name: "n2560-r17"},
		{Name: "n5120-r17"},
		{Name: "n10240-r17"},
		{Name: "n20480-r17"},
		{Name: "n40960-r17"},
		{Name: "n81920-r17"},
	},
}

const (
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N32_r17    = 0
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N40_r17    = 1
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N64_r17    = 2
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N80_r17    = 3
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N128_r17   = 4
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N160_r17   = 5
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N256_r17   = 6
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N320_r17   = 7
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N512_r17   = 8
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N640_r17   = 9
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N1280_r17  = 10
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N2560_r17  = 11
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N5120_r17  = 12
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N10240_r17 = 13
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N20480_r17 = 14
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N40960_r17 = 15
	NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N81920_r17 = 16
)

type NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17 struct {
	Choice    int
	Scs15_r17 *struct {
		Choice     int
		N4_r17     *int64
		N5_r17     *int64
		N8_r17     *int64
		N10_r17    *int64
		N16_r17    *int64
		N20_r17    *int64
		N32_r17    *int64
		N40_r17    *int64
		N64_r17    *int64
		N80_r17    *int64
		N160_r17   *int64
		N320_r17   *int64
		N640_r17   *int64
		N1280_r17  *int64
		N2560_r17  *int64
		N5120_r17  *int64
		N10240_r17 *int64
	}
	Scs30_r17 *struct {
		Choice     int
		N8_r17     *int64
		N10_r17    *int64
		N16_r17    *int64
		N20_r17    *int64
		N32_r17    *int64
		N40_r17    *int64
		N64_r17    *int64
		N80_r17    *int64
		N128_r17   *int64
		N160_r17   *int64
		N320_r17   *int64
		N640_r17   *int64
		N1280_r17  *int64
		N2560_r17  *int64
		N5120_r17  *int64
		N10240_r17 *int64
		N20480_r17 *int64
	}
	Scs60_r17 *struct {
		Choice     int
		N16_r17    *int64
		N20_r17    *int64
		N32_r17    *int64
		N40_r17    *int64
		N64_r17    *int64
		N80_r17    *int64
		N128_r17   *int64
		N160_r17   *int64
		N256_r17   *int64
		N320_r17   *int64
		N640_r17   *int64
		N1280_r17  *int64
		N2560_r17  *int64
		N5120_r17  *int64
		N10240_r17 *int64
		N20480_r17 *int64
		N40960_r17 *int64
	}
	Scs120_r17 *struct {
		Choice     int
		N32_r17    *int64
		N40_r17    *int64
		N64_r17    *int64
		N80_r17    *int64
		N128_r17   *int64
		N160_r17   *int64
		N256_r17   *int64
		N320_r17   *int64
		N512_r17   *int64
		N640_r17   *int64
		N1280_r17  *int64
		N2560_r17  *int64
		N5120_r17  *int64
		N10240_r17 *int64
		N20480_r17 *int64
		N40960_r17 *int64
		N81920_r17 *int64
	}
}

func (ie *NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17:
		choiceEnc := e.NewChoiceEncoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs15R17Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs15_r17).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs15_r17).Choice {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N4_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N4_r17), per.Constrained(0, 3)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N5_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N5_r17), per.Constrained(0, 4)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N8_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N8_r17), per.Constrained(0, 7)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N10_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N10_r17), per.Constrained(0, 9)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N16_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N16_r17), per.Constrained(0, 15)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N20_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N20_r17), per.Constrained(0, 19)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N32_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N32_r17), per.Constrained(0, 31)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N40_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N40_r17), per.Constrained(0, 39)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N64_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N64_r17), per.Constrained(0, 63)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N80_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N80_r17), per.Constrained(0, 79)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N160_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N160_r17), per.Constrained(0, 159)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N320_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N320_r17), per.Constrained(0, 319)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N640_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N640_r17), per.Constrained(0, 639)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N1280_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N1280_r17), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N2560_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N2560_r17), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N5120_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N5120_r17), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N10240_r17:
			if err := e.EncodeInteger((*(*ie.Scs15_r17).N10240_r17), per.Constrained(0, 10239)); err != nil {
				return err
			}
		}
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17:
		choiceEnc := e.NewChoiceEncoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs30R17Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs30_r17).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs30_r17).Choice {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N8_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N8_r17), per.Constrained(0, 7)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N10_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N10_r17), per.Constrained(0, 9)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N16_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N16_r17), per.Constrained(0, 15)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N20_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N20_r17), per.Constrained(0, 19)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N32_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N32_r17), per.Constrained(0, 31)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N40_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N40_r17), per.Constrained(0, 39)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N64_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N64_r17), per.Constrained(0, 63)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N80_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N80_r17), per.Constrained(0, 79)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N128_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N128_r17), per.Constrained(0, 127)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N160_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N160_r17), per.Constrained(0, 159)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N320_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N320_r17), per.Constrained(0, 319)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N640_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N640_r17), per.Constrained(0, 639)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N1280_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N1280_r17), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N2560_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N2560_r17), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N5120_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N5120_r17), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N10240_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N10240_r17), per.Constrained(0, 10239)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N20480_r17:
			if err := e.EncodeInteger((*(*ie.Scs30_r17).N20480_r17), per.Constrained(0, 20479)); err != nil {
				return err
			}
		}
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17:
		choiceEnc := e.NewChoiceEncoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs60R17Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs60_r17).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs60_r17).Choice {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N16_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N16_r17), per.Constrained(0, 15)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N20_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N20_r17), per.Constrained(0, 19)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N32_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N32_r17), per.Constrained(0, 31)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N40_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N40_r17), per.Constrained(0, 39)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N64_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N64_r17), per.Constrained(0, 63)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N80_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N80_r17), per.Constrained(0, 79)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N128_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N128_r17), per.Constrained(0, 127)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N160_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N160_r17), per.Constrained(0, 159)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N256_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N256_r17), per.Constrained(0, 255)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N320_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N320_r17), per.Constrained(0, 319)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N640_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N640_r17), per.Constrained(0, 639)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N1280_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N1280_r17), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N2560_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N2560_r17), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N5120_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N5120_r17), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N10240_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N10240_r17), per.Constrained(0, 10239)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N20480_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N20480_r17), per.Constrained(0, 20479)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N40960_r17:
			if err := e.EncodeInteger((*(*ie.Scs60_r17).N40960_r17), per.Constrained(0, 40959)); err != nil {
				return err
			}
		}
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17:
		choiceEnc := e.NewChoiceEncoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs120R17Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs120_r17).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs120_r17).Choice {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N32_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N32_r17), per.Constrained(0, 31)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N40_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N40_r17), per.Constrained(0, 39)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N64_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N64_r17), per.Constrained(0, 63)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N80_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N80_r17), per.Constrained(0, 79)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N128_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N128_r17), per.Constrained(0, 127)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N160_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N160_r17), per.Constrained(0, 159)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N256_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N256_r17), per.Constrained(0, 255)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N320_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N320_r17), per.Constrained(0, 319)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N512_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N512_r17), per.Constrained(0, 511)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N640_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N640_r17), per.Constrained(0, 639)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N1280_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N1280_r17), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N2560_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N2560_r17), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N5120_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N5120_r17), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N10240_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N10240_r17), per.Constrained(0, 10239)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N20480_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N20480_r17), per.Constrained(0, 20479)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N40960_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N40960_r17), per.Constrained(0, 40959)); err != nil {
				return err
			}
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N81920_r17:
			if err := e.EncodeInteger((*(*ie.Scs120_r17).N81920_r17), per.Constrained(0, 81919)); err != nil {
				return err
			}
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "NR-DL-PRS-Periodicity-and-ResourceSetSlotOffset-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "NR-DL-PRS-Periodicity-and-ResourceSetSlotOffset-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17:
		ie.Scs15_r17 = &struct {
			Choice     int
			N4_r17     *int64
			N5_r17     *int64
			N8_r17     *int64
			N10_r17    *int64
			N16_r17    *int64
			N20_r17    *int64
			N32_r17    *int64
			N40_r17    *int64
			N64_r17    *int64
			N80_r17    *int64
			N160_r17   *int64
			N320_r17   *int64
			N640_r17   *int64
			N1280_r17  *int64
			N2560_r17  *int64
			N5120_r17  *int64
			N10240_r17 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs15R17Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs15_r17).Choice = int(index)
		switch index {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N4_r17:
			(*ie.Scs15_r17).N4_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N4_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N5_r17:
			(*ie.Scs15_r17).N5_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N5_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N8_r17:
			(*ie.Scs15_r17).N8_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N8_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N10_r17:
			(*ie.Scs15_r17).N10_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N10_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N16_r17:
			(*ie.Scs15_r17).N16_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N16_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N20_r17:
			(*ie.Scs15_r17).N20_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N20_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N32_r17:
			(*ie.Scs15_r17).N32_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N32_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N40_r17:
			(*ie.Scs15_r17).N40_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N40_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N64_r17:
			(*ie.Scs15_r17).N64_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N64_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N80_r17:
			(*ie.Scs15_r17).N80_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N80_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N160_r17:
			(*ie.Scs15_r17).N160_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N160_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N320_r17:
			(*ie.Scs15_r17).N320_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N320_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N640_r17:
			(*ie.Scs15_r17).N640_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N640_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N1280_r17:
			(*ie.Scs15_r17).N1280_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N1280_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N2560_r17:
			(*ie.Scs15_r17).N2560_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N2560_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N5120_r17:
			(*ie.Scs15_r17).N5120_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N5120_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs15_r17_N10240_r17:
			(*ie.Scs15_r17).N10240_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs15_r17).N10240_r17) = v
		}
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17:
		ie.Scs30_r17 = &struct {
			Choice     int
			N8_r17     *int64
			N10_r17    *int64
			N16_r17    *int64
			N20_r17    *int64
			N32_r17    *int64
			N40_r17    *int64
			N64_r17    *int64
			N80_r17    *int64
			N128_r17   *int64
			N160_r17   *int64
			N320_r17   *int64
			N640_r17   *int64
			N1280_r17  *int64
			N2560_r17  *int64
			N5120_r17  *int64
			N10240_r17 *int64
			N20480_r17 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs30R17Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs30_r17).Choice = int(index)
		switch index {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N8_r17:
			(*ie.Scs30_r17).N8_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N8_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N10_r17:
			(*ie.Scs30_r17).N10_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N10_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N16_r17:
			(*ie.Scs30_r17).N16_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N16_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N20_r17:
			(*ie.Scs30_r17).N20_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N20_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N32_r17:
			(*ie.Scs30_r17).N32_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N32_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N40_r17:
			(*ie.Scs30_r17).N40_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N40_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N64_r17:
			(*ie.Scs30_r17).N64_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N64_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N80_r17:
			(*ie.Scs30_r17).N80_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N80_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N128_r17:
			(*ie.Scs30_r17).N128_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N128_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N160_r17:
			(*ie.Scs30_r17).N160_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N160_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N320_r17:
			(*ie.Scs30_r17).N320_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N320_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N640_r17:
			(*ie.Scs30_r17).N640_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N640_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N1280_r17:
			(*ie.Scs30_r17).N1280_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N1280_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N2560_r17:
			(*ie.Scs30_r17).N2560_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N2560_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N5120_r17:
			(*ie.Scs30_r17).N5120_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N5120_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N10240_r17:
			(*ie.Scs30_r17).N10240_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N10240_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs30_r17_N20480_r17:
			(*ie.Scs30_r17).N20480_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 20479))
			if err != nil {
				return err
			}
			(*(*ie.Scs30_r17).N20480_r17) = v
		}
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17:
		ie.Scs60_r17 = &struct {
			Choice     int
			N16_r17    *int64
			N20_r17    *int64
			N32_r17    *int64
			N40_r17    *int64
			N64_r17    *int64
			N80_r17    *int64
			N128_r17   *int64
			N160_r17   *int64
			N256_r17   *int64
			N320_r17   *int64
			N640_r17   *int64
			N1280_r17  *int64
			N2560_r17  *int64
			N5120_r17  *int64
			N10240_r17 *int64
			N20480_r17 *int64
			N40960_r17 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs60R17Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs60_r17).Choice = int(index)
		switch index {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N16_r17:
			(*ie.Scs60_r17).N16_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N16_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N20_r17:
			(*ie.Scs60_r17).N20_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N20_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N32_r17:
			(*ie.Scs60_r17).N32_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N32_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N40_r17:
			(*ie.Scs60_r17).N40_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N40_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N64_r17:
			(*ie.Scs60_r17).N64_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N64_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N80_r17:
			(*ie.Scs60_r17).N80_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N80_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N128_r17:
			(*ie.Scs60_r17).N128_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N128_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N160_r17:
			(*ie.Scs60_r17).N160_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N160_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N256_r17:
			(*ie.Scs60_r17).N256_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N256_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N320_r17:
			(*ie.Scs60_r17).N320_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N320_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N640_r17:
			(*ie.Scs60_r17).N640_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N640_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N1280_r17:
			(*ie.Scs60_r17).N1280_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N1280_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N2560_r17:
			(*ie.Scs60_r17).N2560_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N2560_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N5120_r17:
			(*ie.Scs60_r17).N5120_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N5120_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N10240_r17:
			(*ie.Scs60_r17).N10240_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N10240_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N20480_r17:
			(*ie.Scs60_r17).N20480_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 20479))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N20480_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs60_r17_N40960_r17:
			(*ie.Scs60_r17).N40960_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 40959))
			if err != nil {
				return err
			}
			(*(*ie.Scs60_r17).N40960_r17) = v
		}
	case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17:
		ie.Scs120_r17 = &struct {
			Choice     int
			N32_r17    *int64
			N40_r17    *int64
			N64_r17    *int64
			N80_r17    *int64
			N128_r17   *int64
			N160_r17   *int64
			N256_r17   *int64
			N320_r17   *int64
			N512_r17   *int64
			N640_r17   *int64
			N1280_r17  *int64
			N2560_r17  *int64
			N5120_r17  *int64
			N10240_r17 *int64
			N20480_r17 *int64
			N40960_r17 *int64
			N81920_r17 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(nRDLPRSPeriodicityAndResourceSetSlotOffsetR17Scs120R17Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs120_r17).Choice = int(index)
		switch index {
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N32_r17:
			(*ie.Scs120_r17).N32_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N32_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N40_r17:
			(*ie.Scs120_r17).N40_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N40_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N64_r17:
			(*ie.Scs120_r17).N64_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N64_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N80_r17:
			(*ie.Scs120_r17).N80_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N80_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N128_r17:
			(*ie.Scs120_r17).N128_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N128_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N160_r17:
			(*ie.Scs120_r17).N160_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N160_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N256_r17:
			(*ie.Scs120_r17).N256_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N256_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N320_r17:
			(*ie.Scs120_r17).N320_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N320_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N512_r17:
			(*ie.Scs120_r17).N512_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 511))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N512_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N640_r17:
			(*ie.Scs120_r17).N640_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N640_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N1280_r17:
			(*ie.Scs120_r17).N1280_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N1280_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N2560_r17:
			(*ie.Scs120_r17).N2560_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N2560_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N5120_r17:
			(*ie.Scs120_r17).N5120_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N5120_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N10240_r17:
			(*ie.Scs120_r17).N10240_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N10240_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N20480_r17:
			(*ie.Scs120_r17).N20480_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 20479))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N20480_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N40960_r17:
			(*ie.Scs120_r17).N40960_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 40959))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N40960_r17) = v
		case NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17_Scs120_r17_N81920_r17:
			(*ie.Scs120_r17).N81920_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 81919))
			if err != nil {
				return err
			}
			(*(*ie.Scs120_r17).N81920_r17) = v
		}
	default:
		return &per.DecodeError{TypeName: "NR-DL-PRS-Periodicity-and-ResourceSetSlotOffset-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
