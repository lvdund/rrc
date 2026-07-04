// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-PPW-PeriodicityAndStartSlot-r17 (line 7653).

var dLPPWPeriodicityAndStartSlotR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scs15"},
		{Name: "scs30"},
		{Name: "scs60"},
		{Name: "scs120"},
	},
}

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15  = 0
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30  = 1
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60  = 2
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120 = 3
)

var dLPPWPeriodicityAndStartSlotR17Scs15Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n4"},
		{Name: "n5"},
		{Name: "n8"},
		{Name: "n10"},
		{Name: "n16"},
		{Name: "n20"},
		{Name: "n32"},
		{Name: "n40"},
		{Name: "n64"},
		{Name: "n80"},
		{Name: "n160"},
		{Name: "n320"},
		{Name: "n640"},
		{Name: "n1280"},
		{Name: "n2560"},
		{Name: "n5120"},
		{Name: "n10240"},
	},
}

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N4     = 0
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N5     = 1
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N8     = 2
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N10    = 3
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N16    = 4
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N20    = 5
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N32    = 6
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N40    = 7
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N64    = 8
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N80    = 9
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N160   = 10
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N320   = 11
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N640   = 12
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N1280  = 13
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N2560  = 14
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N5120  = 15
	DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N10240 = 16
)

var dLPPWPeriodicityAndStartSlotR17Scs30Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n8"},
		{Name: "n10"},
		{Name: "n16"},
		{Name: "n20"},
		{Name: "n32"},
		{Name: "n40"},
		{Name: "n64"},
		{Name: "n80"},
		{Name: "n128"},
		{Name: "n160"},
		{Name: "n320"},
		{Name: "n640"},
		{Name: "n1280"},
		{Name: "n2560"},
		{Name: "n5120"},
		{Name: "n10240"},
		{Name: "n20480"},
	},
}

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N8     = 0
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N10    = 1
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N16    = 2
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N20    = 3
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N32    = 4
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N40    = 5
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N64    = 6
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N80    = 7
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N128   = 8
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N160   = 9
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N320   = 10
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N640   = 11
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N1280  = 12
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N2560  = 13
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N5120  = 14
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N10240 = 15
	DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N20480 = 16
)

var dLPPWPeriodicityAndStartSlotR17Scs60Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n16"},
		{Name: "n20"},
		{Name: "n32"},
		{Name: "n40"},
		{Name: "n64"},
		{Name: "n80"},
		{Name: "n128"},
		{Name: "n160"},
		{Name: "n256"},
		{Name: "n320"},
		{Name: "n640"},
		{Name: "n1280"},
		{Name: "n2560"},
		{Name: "n5120"},
		{Name: "n10240"},
		{Name: "n20480"},
		{Name: "n40960"},
	},
}

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N16    = 0
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N20    = 1
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N32    = 2
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N40    = 3
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N64    = 4
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N80    = 5
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N128   = 6
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N160   = 7
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N256   = 8
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N320   = 9
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N640   = 10
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N1280  = 11
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N2560  = 12
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N5120  = 13
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N10240 = 14
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N20480 = 15
	DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N40960 = 16
)

var dLPPWPeriodicityAndStartSlotR17Scs120Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n32"},
		{Name: "n40"},
		{Name: "n64"},
		{Name: "n80"},
		{Name: "n128"},
		{Name: "n160"},
		{Name: "n256"},
		{Name: "n320"},
		{Name: "n512"},
		{Name: "n640"},
		{Name: "n1280"},
		{Name: "n2560"},
		{Name: "n5120"},
		{Name: "n10240"},
		{Name: "n20480"},
		{Name: "n40960"},
		{Name: "n81920"},
	},
}

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N32    = 0
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N40    = 1
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N64    = 2
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N80    = 3
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N128   = 4
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N160   = 5
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N256   = 6
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N320   = 7
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N512   = 8
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N640   = 9
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N1280  = 10
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N2560  = 11
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N5120  = 12
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N10240 = 13
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N20480 = 14
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N40960 = 15
	DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N81920 = 16
)

type DL_PPW_PeriodicityAndStartSlot_r17 struct {
	Choice int
	Scs15  *struct {
		Choice int
		N4     *int64
		N5     *int64
		N8     *int64
		N10    *int64
		N16    *int64
		N20    *int64
		N32    *int64
		N40    *int64
		N64    *int64
		N80    *int64
		N160   *int64
		N320   *int64
		N640   *int64
		N1280  *int64
		N2560  *int64
		N5120  *int64
		N10240 *int64
	}
	Scs30 *struct {
		Choice int
		N8     *int64
		N10    *int64
		N16    *int64
		N20    *int64
		N32    *int64
		N40    *int64
		N64    *int64
		N80    *int64
		N128   *int64
		N160   *int64
		N320   *int64
		N640   *int64
		N1280  *int64
		N2560  *int64
		N5120  *int64
		N10240 *int64
		N20480 *int64
	}
	Scs60 *struct {
		Choice int
		N16    *int64
		N20    *int64
		N32    *int64
		N40    *int64
		N64    *int64
		N80    *int64
		N128   *int64
		N160   *int64
		N256   *int64
		N320   *int64
		N640   *int64
		N1280  *int64
		N2560  *int64
		N5120  *int64
		N10240 *int64
		N20480 *int64
		N40960 *int64
	}
	Scs120 *struct {
		Choice int
		N32    *int64
		N40    *int64
		N64    *int64
		N80    *int64
		N128   *int64
		N160   *int64
		N256   *int64
		N320   *int64
		N512   *int64
		N640   *int64
		N1280  *int64
		N2560  *int64
		N5120  *int64
		N10240 *int64
		N20480 *int64
		N40960 *int64
		N81920 *int64
	}
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(dLPPWPeriodicityAndStartSlotR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs15:
		choiceEnc := e.NewChoiceEncoder(dLPPWPeriodicityAndStartSlotR17Scs15Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs15).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs15).Choice {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N4:
			if err := e.EncodeInteger((*(*ie.Scs15).N4), per.Constrained(0, 3)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N5:
			if err := e.EncodeInteger((*(*ie.Scs15).N5), per.Constrained(0, 4)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N8:
			if err := e.EncodeInteger((*(*ie.Scs15).N8), per.Constrained(0, 7)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N10:
			if err := e.EncodeInteger((*(*ie.Scs15).N10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N16:
			if err := e.EncodeInteger((*(*ie.Scs15).N16), per.Constrained(0, 15)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N20:
			if err := e.EncodeInteger((*(*ie.Scs15).N20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N32:
			if err := e.EncodeInteger((*(*ie.Scs15).N32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N40:
			if err := e.EncodeInteger((*(*ie.Scs15).N40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N64:
			if err := e.EncodeInteger((*(*ie.Scs15).N64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N80:
			if err := e.EncodeInteger((*(*ie.Scs15).N80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N160:
			if err := e.EncodeInteger((*(*ie.Scs15).N160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N320:
			if err := e.EncodeInteger((*(*ie.Scs15).N320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N640:
			if err := e.EncodeInteger((*(*ie.Scs15).N640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N1280:
			if err := e.EncodeInteger((*(*ie.Scs15).N1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N2560:
			if err := e.EncodeInteger((*(*ie.Scs15).N2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N5120:
			if err := e.EncodeInteger((*(*ie.Scs15).N5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N10240:
			if err := e.EncodeInteger((*(*ie.Scs15).N10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs30:
		choiceEnc := e.NewChoiceEncoder(dLPPWPeriodicityAndStartSlotR17Scs30Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs30).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs30).Choice {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N8:
			if err := e.EncodeInteger((*(*ie.Scs30).N8), per.Constrained(0, 7)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N10:
			if err := e.EncodeInteger((*(*ie.Scs30).N10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N16:
			if err := e.EncodeInteger((*(*ie.Scs30).N16), per.Constrained(0, 15)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N20:
			if err := e.EncodeInteger((*(*ie.Scs30).N20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N32:
			if err := e.EncodeInteger((*(*ie.Scs30).N32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N40:
			if err := e.EncodeInteger((*(*ie.Scs30).N40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N64:
			if err := e.EncodeInteger((*(*ie.Scs30).N64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N80:
			if err := e.EncodeInteger((*(*ie.Scs30).N80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N128:
			if err := e.EncodeInteger((*(*ie.Scs30).N128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N160:
			if err := e.EncodeInteger((*(*ie.Scs30).N160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N320:
			if err := e.EncodeInteger((*(*ie.Scs30).N320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N640:
			if err := e.EncodeInteger((*(*ie.Scs30).N640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N1280:
			if err := e.EncodeInteger((*(*ie.Scs30).N1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N2560:
			if err := e.EncodeInteger((*(*ie.Scs30).N2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N5120:
			if err := e.EncodeInteger((*(*ie.Scs30).N5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N10240:
			if err := e.EncodeInteger((*(*ie.Scs30).N10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N20480:
			if err := e.EncodeInteger((*(*ie.Scs30).N20480), per.Constrained(0, 20479)); err != nil {
				return err
			}
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs60:
		choiceEnc := e.NewChoiceEncoder(dLPPWPeriodicityAndStartSlotR17Scs60Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs60).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs60).Choice {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N16:
			if err := e.EncodeInteger((*(*ie.Scs60).N16), per.Constrained(0, 15)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N20:
			if err := e.EncodeInteger((*(*ie.Scs60).N20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N32:
			if err := e.EncodeInteger((*(*ie.Scs60).N32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N40:
			if err := e.EncodeInteger((*(*ie.Scs60).N40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N64:
			if err := e.EncodeInteger((*(*ie.Scs60).N64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N80:
			if err := e.EncodeInteger((*(*ie.Scs60).N80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N128:
			if err := e.EncodeInteger((*(*ie.Scs60).N128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N160:
			if err := e.EncodeInteger((*(*ie.Scs60).N160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N256:
			if err := e.EncodeInteger((*(*ie.Scs60).N256), per.Constrained(0, 255)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N320:
			if err := e.EncodeInteger((*(*ie.Scs60).N320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N640:
			if err := e.EncodeInteger((*(*ie.Scs60).N640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N1280:
			if err := e.EncodeInteger((*(*ie.Scs60).N1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N2560:
			if err := e.EncodeInteger((*(*ie.Scs60).N2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N5120:
			if err := e.EncodeInteger((*(*ie.Scs60).N5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N10240:
			if err := e.EncodeInteger((*(*ie.Scs60).N10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N20480:
			if err := e.EncodeInteger((*(*ie.Scs60).N20480), per.Constrained(0, 20479)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N40960:
			if err := e.EncodeInteger((*(*ie.Scs60).N40960), per.Constrained(0, 40959)); err != nil {
				return err
			}
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs120:
		choiceEnc := e.NewChoiceEncoder(dLPPWPeriodicityAndStartSlotR17Scs120Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Scs120).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Scs120).Choice {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N32:
			if err := e.EncodeInteger((*(*ie.Scs120).N32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N40:
			if err := e.EncodeInteger((*(*ie.Scs120).N40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N64:
			if err := e.EncodeInteger((*(*ie.Scs120).N64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N80:
			if err := e.EncodeInteger((*(*ie.Scs120).N80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N128:
			if err := e.EncodeInteger((*(*ie.Scs120).N128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N160:
			if err := e.EncodeInteger((*(*ie.Scs120).N160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N256:
			if err := e.EncodeInteger((*(*ie.Scs120).N256), per.Constrained(0, 255)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N320:
			if err := e.EncodeInteger((*(*ie.Scs120).N320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N512:
			if err := e.EncodeInteger((*(*ie.Scs120).N512), per.Constrained(0, 511)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N640:
			if err := e.EncodeInteger((*(*ie.Scs120).N640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N1280:
			if err := e.EncodeInteger((*(*ie.Scs120).N1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N2560:
			if err := e.EncodeInteger((*(*ie.Scs120).N2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N5120:
			if err := e.EncodeInteger((*(*ie.Scs120).N5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N10240:
			if err := e.EncodeInteger((*(*ie.Scs120).N10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N20480:
			if err := e.EncodeInteger((*(*ie.Scs120).N20480), per.Constrained(0, 20479)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N40960:
			if err := e.EncodeInteger((*(*ie.Scs120).N40960), per.Constrained(0, 40959)); err != nil {
				return err
			}
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N81920:
			if err := e.EncodeInteger((*(*ie.Scs120).N81920), per.Constrained(0, 81919)); err != nil {
				return err
			}
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "DL-PPW-PeriodicityAndStartSlot-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(dLPPWPeriodicityAndStartSlotR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "DL-PPW-PeriodicityAndStartSlot-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs15:
		ie.Scs15 = &struct {
			Choice int
			N4     *int64
			N5     *int64
			N8     *int64
			N10    *int64
			N16    *int64
			N20    *int64
			N32    *int64
			N40    *int64
			N64    *int64
			N80    *int64
			N160   *int64
			N320   *int64
			N640   *int64
			N1280  *int64
			N2560  *int64
			N5120  *int64
			N10240 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(dLPPWPeriodicityAndStartSlotR17Scs15Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs15).Choice = int(index)
		switch index {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N4:
			(*ie.Scs15).N4 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N4) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N5:
			(*ie.Scs15).N5 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N5) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N8:
			(*ie.Scs15).N8 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N8) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N10:
			(*ie.Scs15).N10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N10) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N16:
			(*ie.Scs15).N16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N16) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N20:
			(*ie.Scs15).N20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N20) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N32:
			(*ie.Scs15).N32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N32) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N40:
			(*ie.Scs15).N40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N40) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N64:
			(*ie.Scs15).N64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N64) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N80:
			(*ie.Scs15).N80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N80) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N160:
			(*ie.Scs15).N160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N160) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N320:
			(*ie.Scs15).N320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N320) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N640:
			(*ie.Scs15).N640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N640) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N1280:
			(*ie.Scs15).N1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N1280) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N2560:
			(*ie.Scs15).N2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N2560) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N5120:
			(*ie.Scs15).N5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N5120) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs15_N10240:
			(*ie.Scs15).N10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs15).N10240) = v
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs30:
		ie.Scs30 = &struct {
			Choice int
			N8     *int64
			N10    *int64
			N16    *int64
			N20    *int64
			N32    *int64
			N40    *int64
			N64    *int64
			N80    *int64
			N128   *int64
			N160   *int64
			N320   *int64
			N640   *int64
			N1280  *int64
			N2560  *int64
			N5120  *int64
			N10240 *int64
			N20480 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(dLPPWPeriodicityAndStartSlotR17Scs30Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs30).Choice = int(index)
		switch index {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N8:
			(*ie.Scs30).N8 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N8) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N10:
			(*ie.Scs30).N10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N10) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N16:
			(*ie.Scs30).N16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N16) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N20:
			(*ie.Scs30).N20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N20) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N32:
			(*ie.Scs30).N32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N32) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N40:
			(*ie.Scs30).N40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N40) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N64:
			(*ie.Scs30).N64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N64) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N80:
			(*ie.Scs30).N80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N80) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N128:
			(*ie.Scs30).N128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N128) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N160:
			(*ie.Scs30).N160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N160) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N320:
			(*ie.Scs30).N320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N320) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N640:
			(*ie.Scs30).N640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N640) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N1280:
			(*ie.Scs30).N1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N1280) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N2560:
			(*ie.Scs30).N2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N2560) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N5120:
			(*ie.Scs30).N5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N5120) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N10240:
			(*ie.Scs30).N10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N10240) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs30_N20480:
			(*ie.Scs30).N20480 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 20479))
			if err != nil {
				return err
			}
			(*(*ie.Scs30).N20480) = v
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs60:
		ie.Scs60 = &struct {
			Choice int
			N16    *int64
			N20    *int64
			N32    *int64
			N40    *int64
			N64    *int64
			N80    *int64
			N128   *int64
			N160   *int64
			N256   *int64
			N320   *int64
			N640   *int64
			N1280  *int64
			N2560  *int64
			N5120  *int64
			N10240 *int64
			N20480 *int64
			N40960 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(dLPPWPeriodicityAndStartSlotR17Scs60Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs60).Choice = int(index)
		switch index {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N16:
			(*ie.Scs60).N16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N16) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N20:
			(*ie.Scs60).N20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N20) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N32:
			(*ie.Scs60).N32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N32) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N40:
			(*ie.Scs60).N40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N40) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N64:
			(*ie.Scs60).N64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N64) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N80:
			(*ie.Scs60).N80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N80) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N128:
			(*ie.Scs60).N128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N128) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N160:
			(*ie.Scs60).N160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N160) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N256:
			(*ie.Scs60).N256 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N256) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N320:
			(*ie.Scs60).N320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N320) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N640:
			(*ie.Scs60).N640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N640) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N1280:
			(*ie.Scs60).N1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N1280) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N2560:
			(*ie.Scs60).N2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N2560) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N5120:
			(*ie.Scs60).N5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N5120) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N10240:
			(*ie.Scs60).N10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N10240) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N20480:
			(*ie.Scs60).N20480 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 20479))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N20480) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs60_N40960:
			(*ie.Scs60).N40960 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 40959))
			if err != nil {
				return err
			}
			(*(*ie.Scs60).N40960) = v
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Scs120:
		ie.Scs120 = &struct {
			Choice int
			N32    *int64
			N40    *int64
			N64    *int64
			N80    *int64
			N128   *int64
			N160   *int64
			N256   *int64
			N320   *int64
			N512   *int64
			N640   *int64
			N1280  *int64
			N2560  *int64
			N5120  *int64
			N10240 *int64
			N20480 *int64
			N40960 *int64
			N81920 *int64
		}{}
		choiceDec := d.NewChoiceDecoder(dLPPWPeriodicityAndStartSlotR17Scs120Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Scs120).Choice = int(index)
		switch index {
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N32:
			(*ie.Scs120).N32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N32) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N40:
			(*ie.Scs120).N40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N40) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N64:
			(*ie.Scs120).N64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N64) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N80:
			(*ie.Scs120).N80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N80) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N128:
			(*ie.Scs120).N128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N128) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N160:
			(*ie.Scs120).N160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N160) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N256:
			(*ie.Scs120).N256 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N256) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N320:
			(*ie.Scs120).N320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N320) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N512:
			(*ie.Scs120).N512 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 511))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N512) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N640:
			(*ie.Scs120).N640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N640) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N1280:
			(*ie.Scs120).N1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N1280) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N2560:
			(*ie.Scs120).N2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N2560) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N5120:
			(*ie.Scs120).N5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N5120) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N10240:
			(*ie.Scs120).N10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N10240) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N20480:
			(*ie.Scs120).N20480 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 20479))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N20480) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N40960:
			(*ie.Scs120).N40960 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 40959))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N40960) = v
		case DL_PPW_PeriodicityAndStartSlot_r17_Scs120_N81920:
			(*ie.Scs120).N81920 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 81919))
			if err != nil {
				return err
			}
			(*(*ie.Scs120).N81920) = v
		}
	default:
		return &per.DecodeError{TypeName: "DL-PPW-PeriodicityAndStartSlot-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
