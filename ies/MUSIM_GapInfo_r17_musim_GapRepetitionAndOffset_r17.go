package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_nothing uint64 = iota
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms20_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms40_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms80_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms160_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms320_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms640_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms1280_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms2560_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms5120_r17
)

type MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17 struct {
	Choice     uint64
	Ms20_r17   int64 `lb:0,ub:19,madatory`
	Ms40_r17   int64 `lb:0,ub:39,madatory`
	Ms80_r17   int64 `lb:0,ub:79,madatory`
	Ms160_r17  int64 `lb:0,ub:159,madatory`
	Ms320_r17  int64 `lb:0,ub:319,madatory`
	Ms640_r17  int64 `lb:0,ub:639,madatory`
	Ms1280_r17 int64 `lb:0,ub:1279,madatory`
	Ms2560_r17 int64 `lb:0,ub:2559,madatory`
	Ms5120_r17 int64 `lb:0,ub:5119,madatory`
}

func (ie *MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms20_r17:
		if err = w.WriteInteger(int64(ie.Ms20_r17), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Ms20_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms40_r17:
		if err = w.WriteInteger(int64(ie.Ms40_r17), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Ms40_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms80_r17:
		if err = w.WriteInteger(int64(ie.Ms80_r17), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Ms80_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms160_r17:
		if err = w.WriteInteger(int64(ie.Ms160_r17), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Ms160_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms320_r17:
		if err = w.WriteInteger(int64(ie.Ms320_r17), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Ms320_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms640_r17:
		if err = w.WriteInteger(int64(ie.Ms640_r17), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Ms640_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms1280_r17:
		if err = w.WriteInteger(int64(ie.Ms1280_r17), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode Ms1280_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms2560_r17:
		if err = w.WriteInteger(int64(ie.Ms2560_r17), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode Ms2560_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms5120_r17:
		if err = w.WriteInteger(int64(ie.Ms5120_r17), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode Ms5120_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms20_r17:
		var tmp_int_Ms20_r17 int64
		if tmp_int_Ms20_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Ms20_r17", err)
		}
		ie.Ms20_r17 = tmp_int_Ms20_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms40_r17:
		var tmp_int_Ms40_r17 int64
		if tmp_int_Ms40_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Ms40_r17", err)
		}
		ie.Ms40_r17 = tmp_int_Ms40_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms80_r17:
		var tmp_int_Ms80_r17 int64
		if tmp_int_Ms80_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Ms80_r17", err)
		}
		ie.Ms80_r17 = tmp_int_Ms80_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms160_r17:
		var tmp_int_Ms160_r17 int64
		if tmp_int_Ms160_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Ms160_r17", err)
		}
		ie.Ms160_r17 = tmp_int_Ms160_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms320_r17:
		var tmp_int_Ms320_r17 int64
		if tmp_int_Ms320_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Ms320_r17", err)
		}
		ie.Ms320_r17 = tmp_int_Ms320_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms640_r17:
		var tmp_int_Ms640_r17 int64
		if tmp_int_Ms640_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Ms640_r17", err)
		}
		ie.Ms640_r17 = tmp_int_Ms640_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms1280_r17:
		var tmp_int_Ms1280_r17 int64
		if tmp_int_Ms1280_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode Ms1280_r17", err)
		}
		ie.Ms1280_r17 = tmp_int_Ms1280_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms2560_r17:
		var tmp_int_Ms2560_r17 int64
		if tmp_int_Ms2560_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode Ms2560_r17", err)
		}
		ie.Ms2560_r17 = tmp_int_Ms2560_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_Ms5120_r17:
		var tmp_int_Ms5120_r17 int64
		if tmp_int_Ms5120_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode Ms5120_r17", err)
		}
		ie.Ms5120_r17 = tmp_int_Ms5120_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
