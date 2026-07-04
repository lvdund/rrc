// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NCR-MsPeriodicityAndSlotOffset-r18 (line 10352).

var nCRMsPeriodicityAndSlotOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms1"},
		{Name: "ms2"},
		{Name: "ms4"},
		{Name: "ms5"},
		{Name: "ms8"},
		{Name: "ms10"},
		{Name: "ms16"},
		{Name: "ms20"},
		{Name: "ms32"},
		{Name: "ms40"},
		{Name: "ms64"},
		{Name: "ms80"},
		{Name: "ms128"},
		{Name: "ms160"},
		{Name: "ms256"},
		{Name: "ms320"},
		{Name: "ms512"},
		{Name: "ms640"},
		{Name: "ms1024"},
		{Name: "ms1280"},
		{Name: "ms2560"},
		{Name: "ms5120"},
		{Name: "ms10240"},
	},
}

const (
	NCR_MsPeriodicityAndSlotOffset_r18_Ms1     = 0
	NCR_MsPeriodicityAndSlotOffset_r18_Ms2     = 1
	NCR_MsPeriodicityAndSlotOffset_r18_Ms4     = 2
	NCR_MsPeriodicityAndSlotOffset_r18_Ms5     = 3
	NCR_MsPeriodicityAndSlotOffset_r18_Ms8     = 4
	NCR_MsPeriodicityAndSlotOffset_r18_Ms10    = 5
	NCR_MsPeriodicityAndSlotOffset_r18_Ms16    = 6
	NCR_MsPeriodicityAndSlotOffset_r18_Ms20    = 7
	NCR_MsPeriodicityAndSlotOffset_r18_Ms32    = 8
	NCR_MsPeriodicityAndSlotOffset_r18_Ms40    = 9
	NCR_MsPeriodicityAndSlotOffset_r18_Ms64    = 10
	NCR_MsPeriodicityAndSlotOffset_r18_Ms80    = 11
	NCR_MsPeriodicityAndSlotOffset_r18_Ms128   = 12
	NCR_MsPeriodicityAndSlotOffset_r18_Ms160   = 13
	NCR_MsPeriodicityAndSlotOffset_r18_Ms256   = 14
	NCR_MsPeriodicityAndSlotOffset_r18_Ms320   = 15
	NCR_MsPeriodicityAndSlotOffset_r18_Ms512   = 16
	NCR_MsPeriodicityAndSlotOffset_r18_Ms640   = 17
	NCR_MsPeriodicityAndSlotOffset_r18_Ms1024  = 18
	NCR_MsPeriodicityAndSlotOffset_r18_Ms1280  = 19
	NCR_MsPeriodicityAndSlotOffset_r18_Ms2560  = 20
	NCR_MsPeriodicityAndSlotOffset_r18_Ms5120  = 21
	NCR_MsPeriodicityAndSlotOffset_r18_Ms10240 = 22
)

type NCR_MsPeriodicityAndSlotOffset_r18 struct {
	Choice  int
	Ms1     *int64
	Ms2     *int64
	Ms4     *int64
	Ms5     *int64
	Ms8     *int64
	Ms10    *int64
	Ms16    *int64
	Ms20    *int64
	Ms32    *int64
	Ms40    *int64
	Ms64    *int64
	Ms80    *int64
	Ms128   *int64
	Ms160   *int64
	Ms256   *int64
	Ms320   *int64
	Ms512   *int64
	Ms640   *int64
	Ms1024  *int64
	Ms1280  *int64
	Ms2560  *int64
	Ms5120  *int64
	Ms10240 *int64
}

func (ie *NCR_MsPeriodicityAndSlotOffset_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(nCRMsPeriodicityAndSlotOffsetR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms1:
		if err := e.EncodeInteger((*ie.Ms1), per.Constrained(0, 15)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms2:
		if err := e.EncodeInteger((*ie.Ms2), per.Constrained(0, 31)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms4:
		if err := e.EncodeInteger((*ie.Ms4), per.Constrained(0, 63)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms5:
		if err := e.EncodeInteger((*ie.Ms5), per.Constrained(0, 79)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms8:
		if err := e.EncodeInteger((*ie.Ms8), per.Constrained(0, 127)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms10:
		if err := e.EncodeInteger((*ie.Ms10), per.Constrained(0, 159)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms16:
		if err := e.EncodeInteger((*ie.Ms16), per.Constrained(0, 255)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms20:
		if err := e.EncodeInteger((*ie.Ms20), per.Constrained(0, 319)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms32:
		if err := e.EncodeInteger((*ie.Ms32), per.Constrained(0, 511)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms40:
		if err := e.EncodeInteger((*ie.Ms40), per.Constrained(0, 639)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms64:
		if err := e.EncodeInteger((*ie.Ms64), per.Constrained(0, 1023)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms80:
		if err := e.EncodeInteger((*ie.Ms80), per.Constrained(0, 1279)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms128:
		if err := e.EncodeInteger((*ie.Ms128), per.Constrained(0, 2047)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms160:
		if err := e.EncodeInteger((*ie.Ms160), per.Constrained(0, 2559)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms256:
		if err := e.EncodeInteger((*ie.Ms256), per.Constrained(0, 4095)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms320:
		if err := e.EncodeInteger((*ie.Ms320), per.Constrained(0, 5119)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms512:
		if err := e.EncodeInteger((*ie.Ms512), per.Constrained(0, 8191)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms640:
		if err := e.EncodeInteger((*ie.Ms640), per.Constrained(0, 10239)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms1024:
		if err := e.EncodeInteger((*ie.Ms1024), per.Constrained(0, 16383)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms1280:
		if err := e.EncodeInteger((*ie.Ms1280), per.Constrained(0, 20479)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms2560:
		if err := e.EncodeInteger((*ie.Ms2560), per.Constrained(0, 40959)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms5120:
		if err := e.EncodeInteger((*ie.Ms5120), per.Constrained(0, 81919)); err != nil {
			return err
		}
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms10240:
		if err := e.EncodeInteger((*ie.Ms10240), per.Constrained(0, 163839)); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "NCR-MsPeriodicityAndSlotOffset-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *NCR_MsPeriodicityAndSlotOffset_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(nCRMsPeriodicityAndSlotOffsetR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "NCR-MsPeriodicityAndSlotOffset-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms1:
		ie.Ms1 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return err
		}
		(*ie.Ms1) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms2:
		ie.Ms2 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie.Ms2) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms4:
		ie.Ms4 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie.Ms4) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms5:
		ie.Ms5 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 79))
		if err != nil {
			return err
		}
		(*ie.Ms5) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms8:
		ie.Ms8 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 127))
		if err != nil {
			return err
		}
		(*ie.Ms8) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms10:
		ie.Ms10 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 159))
		if err != nil {
			return err
		}
		(*ie.Ms10) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms16:
		ie.Ms16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return err
		}
		(*ie.Ms16) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms20:
		ie.Ms20 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 319))
		if err != nil {
			return err
		}
		(*ie.Ms20) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms32:
		ie.Ms32 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 511))
		if err != nil {
			return err
		}
		(*ie.Ms32) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms40:
		ie.Ms40 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 639))
		if err != nil {
			return err
		}
		(*ie.Ms40) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms64:
		ie.Ms64 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1023))
		if err != nil {
			return err
		}
		(*ie.Ms64) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms80:
		ie.Ms80 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1279))
		if err != nil {
			return err
		}
		(*ie.Ms80) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms128:
		ie.Ms128 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 2047))
		if err != nil {
			return err
		}
		(*ie.Ms128) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms160:
		ie.Ms160 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 2559))
		if err != nil {
			return err
		}
		(*ie.Ms160) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms256:
		ie.Ms256 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 4095))
		if err != nil {
			return err
		}
		(*ie.Ms256) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms320:
		ie.Ms320 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 5119))
		if err != nil {
			return err
		}
		(*ie.Ms320) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms512:
		ie.Ms512 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 8191))
		if err != nil {
			return err
		}
		(*ie.Ms512) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms640:
		ie.Ms640 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 10239))
		if err != nil {
			return err
		}
		(*ie.Ms640) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms1024:
		ie.Ms1024 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 16383))
		if err != nil {
			return err
		}
		(*ie.Ms1024) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms1280:
		ie.Ms1280 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 20479))
		if err != nil {
			return err
		}
		(*ie.Ms1280) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms2560:
		ie.Ms2560 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 40959))
		if err != nil {
			return err
		}
		(*ie.Ms2560) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms5120:
		ie.Ms5120 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 81919))
		if err != nil {
			return err
		}
		(*ie.Ms5120) = v
	case NCR_MsPeriodicityAndSlotOffset_r18_Ms10240:
		ie.Ms10240 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 163839))
		if err != nil {
			return err
		}
		(*ie.Ms10240) = v
	default:
		return &per.DecodeError{TypeName: "NCR-MsPeriodicityAndSlotOffset-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
