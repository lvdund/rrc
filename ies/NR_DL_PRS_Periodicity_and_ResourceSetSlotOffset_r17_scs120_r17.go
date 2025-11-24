package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_nothing uint64 = iota
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N32_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N40_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N64_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N80_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N128_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N160_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N256_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N320_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N512_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N640_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N1280_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N2560_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N5120_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N10240_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N20480_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N40960_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N81920_r17
)

type NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17 struct {
	Choice     uint64
	N32_r17    int64 `lb:0,ub:31,madatory`
	N40_r17    int64 `lb:0,ub:39,madatory`
	N64_r17    int64 `lb:0,ub:63,madatory`
	N80_r17    int64 `lb:0,ub:79,madatory`
	N128_r17   int64 `lb:0,ub:127,madatory`
	N160_r17   int64 `lb:0,ub:159,madatory`
	N256_r17   int64 `lb:0,ub:255,madatory`
	N320_r17   int64 `lb:0,ub:319,madatory`
	N512_r17   int64 `lb:0,ub:511,madatory`
	N640_r17   int64 `lb:0,ub:639,madatory`
	N1280_r17  int64 `lb:0,ub:1279,madatory`
	N2560_r17  int64 `lb:0,ub:2559,madatory`
	N5120_r17  int64 `lb:0,ub:5119,madatory`
	N10240_r17 int64 `lb:0,ub:10239,madatory`
	N20480_r17 int64 `lb:0,ub:20479,madatory`
	N40960_r17 int64 `lb:0,ub:40959,madatory`
	N81920_r17 int64 `lb:0,ub:81919,madatory`
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N32_r17:
		if err = w.WriteInteger(int64(ie.N32_r17), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode N32_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N40_r17:
		if err = w.WriteInteger(int64(ie.N40_r17), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode N40_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N64_r17:
		if err = w.WriteInteger(int64(ie.N64_r17), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode N64_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N80_r17:
		if err = w.WriteInteger(int64(ie.N80_r17), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode N80_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N128_r17:
		if err = w.WriteInteger(int64(ie.N128_r17), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode N128_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N160_r17:
		if err = w.WriteInteger(int64(ie.N160_r17), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode N160_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N256_r17:
		if err = w.WriteInteger(int64(ie.N256_r17), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode N256_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N320_r17:
		if err = w.WriteInteger(int64(ie.N320_r17), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode N320_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N512_r17:
		if err = w.WriteInteger(int64(ie.N512_r17), &uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			err = utils.WrapError("Encode N512_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N640_r17:
		if err = w.WriteInteger(int64(ie.N640_r17), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode N640_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N1280_r17:
		if err = w.WriteInteger(int64(ie.N1280_r17), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode N1280_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N2560_r17:
		if err = w.WriteInteger(int64(ie.N2560_r17), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode N2560_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N5120_r17:
		if err = w.WriteInteger(int64(ie.N5120_r17), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode N5120_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N10240_r17:
		if err = w.WriteInteger(int64(ie.N10240_r17), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode N10240_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N20480_r17:
		if err = w.WriteInteger(int64(ie.N20480_r17), &uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode N20480_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N40960_r17:
		if err = w.WriteInteger(int64(ie.N40960_r17), &uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			err = utils.WrapError("Encode N40960_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N81920_r17:
		if err = w.WriteInteger(int64(ie.N81920_r17), &uper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			err = utils.WrapError("Encode N81920_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N32_r17:
		var tmp_int_N32_r17 int64
		if tmp_int_N32_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode N32_r17", err)
		}
		ie.N32_r17 = tmp_int_N32_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N40_r17:
		var tmp_int_N40_r17 int64
		if tmp_int_N40_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode N40_r17", err)
		}
		ie.N40_r17 = tmp_int_N40_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N64_r17:
		var tmp_int_N64_r17 int64
		if tmp_int_N64_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode N64_r17", err)
		}
		ie.N64_r17 = tmp_int_N64_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N80_r17:
		var tmp_int_N80_r17 int64
		if tmp_int_N80_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode N80_r17", err)
		}
		ie.N80_r17 = tmp_int_N80_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N128_r17:
		var tmp_int_N128_r17 int64
		if tmp_int_N128_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode N128_r17", err)
		}
		ie.N128_r17 = tmp_int_N128_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N160_r17:
		var tmp_int_N160_r17 int64
		if tmp_int_N160_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode N160_r17", err)
		}
		ie.N160_r17 = tmp_int_N160_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N256_r17:
		var tmp_int_N256_r17 int64
		if tmp_int_N256_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode N256_r17", err)
		}
		ie.N256_r17 = tmp_int_N256_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N320_r17:
		var tmp_int_N320_r17 int64
		if tmp_int_N320_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode N320_r17", err)
		}
		ie.N320_r17 = tmp_int_N320_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N512_r17:
		var tmp_int_N512_r17 int64
		if tmp_int_N512_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			return utils.WrapError("Decode N512_r17", err)
		}
		ie.N512_r17 = tmp_int_N512_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N640_r17:
		var tmp_int_N640_r17 int64
		if tmp_int_N640_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode N640_r17", err)
		}
		ie.N640_r17 = tmp_int_N640_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N1280_r17:
		var tmp_int_N1280_r17 int64
		if tmp_int_N1280_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode N1280_r17", err)
		}
		ie.N1280_r17 = tmp_int_N1280_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N2560_r17:
		var tmp_int_N2560_r17 int64
		if tmp_int_N2560_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode N2560_r17", err)
		}
		ie.N2560_r17 = tmp_int_N2560_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N5120_r17:
		var tmp_int_N5120_r17 int64
		if tmp_int_N5120_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode N5120_r17", err)
		}
		ie.N5120_r17 = tmp_int_N5120_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N10240_r17:
		var tmp_int_N10240_r17 int64
		if tmp_int_N10240_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode N10240_r17", err)
		}
		ie.N10240_r17 = tmp_int_N10240_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N20480_r17:
		var tmp_int_N20480_r17 int64
		if tmp_int_N20480_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode N20480_r17", err)
		}
		ie.N20480_r17 = tmp_int_N20480_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N40960_r17:
		var tmp_int_N40960_r17 int64
		if tmp_int_N40960_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode N40960_r17", err)
		}
		ie.N40960_r17 = tmp_int_N40960_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17_Choice_N81920_r17:
		var tmp_int_N81920_r17 int64
		if tmp_int_N81920_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			return utils.WrapError("Decode N81920_r17", err)
		}
		ie.N81920_r17 = tmp_int_N81920_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
