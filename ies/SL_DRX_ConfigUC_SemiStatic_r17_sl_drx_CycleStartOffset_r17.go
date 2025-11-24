package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_nothing uint64 = iota
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms10
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms20
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms32
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms40
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms60
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms64
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms70
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms80
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms128
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms160
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms256
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms320
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms512
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms640
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms1024
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms1280
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms2048
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms2560
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms5120
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms10240
)

type SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17 struct {
	Choice  uint64
	Ms10    int64 `lb:0,ub:9,madatory`
	Ms20    int64 `lb:0,ub:19,madatory`
	Ms32    int64 `lb:0,ub:31,madatory`
	Ms40    int64 `lb:0,ub:39,madatory`
	Ms60    int64 `lb:0,ub:59,madatory`
	Ms64    int64 `lb:0,ub:63,madatory`
	Ms70    int64 `lb:0,ub:69,madatory`
	Ms80    int64 `lb:0,ub:79,madatory`
	Ms128   int64 `lb:0,ub:127,madatory`
	Ms160   int64 `lb:0,ub:159,madatory`
	Ms256   int64 `lb:0,ub:255,madatory`
	Ms320   int64 `lb:0,ub:319,madatory`
	Ms512   int64 `lb:0,ub:511,madatory`
	Ms640   int64 `lb:0,ub:639,madatory`
	Ms1024  int64 `lb:0,ub:1023,madatory`
	Ms1280  int64 `lb:0,ub:1279,madatory`
	Ms2048  int64 `lb:0,ub:2047,madatory`
	Ms2560  int64 `lb:0,ub:2559,madatory`
	Ms5120  int64 `lb:0,ub:5119,madatory`
	Ms10240 int64 `lb:0,ub:10239,madatory`
}

func (ie *SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms10:
		if err = w.WriteInteger(int64(ie.Ms10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Ms10", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms20:
		if err = w.WriteInteger(int64(ie.Ms20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Ms20", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms32:
		if err = w.WriteInteger(int64(ie.Ms32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Ms32", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms40:
		if err = w.WriteInteger(int64(ie.Ms40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Ms40", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms60:
		if err = w.WriteInteger(int64(ie.Ms60), &uper.Constraint{Lb: 0, Ub: 59}, false); err != nil {
			err = utils.WrapError("Encode Ms60", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms64:
		if err = w.WriteInteger(int64(ie.Ms64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Ms64", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms70:
		if err = w.WriteInteger(int64(ie.Ms70), &uper.Constraint{Lb: 0, Ub: 69}, false); err != nil {
			err = utils.WrapError("Encode Ms70", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms80:
		if err = w.WriteInteger(int64(ie.Ms80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Ms80", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms128:
		if err = w.WriteInteger(int64(ie.Ms128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode Ms128", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms160:
		if err = w.WriteInteger(int64(ie.Ms160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Ms160", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms256:
		if err = w.WriteInteger(int64(ie.Ms256), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode Ms256", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms320:
		if err = w.WriteInteger(int64(ie.Ms320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Ms320", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms512:
		if err = w.WriteInteger(int64(ie.Ms512), &uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			err = utils.WrapError("Encode Ms512", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms640:
		if err = w.WriteInteger(int64(ie.Ms640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Ms640", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms1024:
		if err = w.WriteInteger(int64(ie.Ms1024), &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			err = utils.WrapError("Encode Ms1024", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms1280:
		if err = w.WriteInteger(int64(ie.Ms1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode Ms1280", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms2048:
		if err = w.WriteInteger(int64(ie.Ms2048), &uper.Constraint{Lb: 0, Ub: 2047}, false); err != nil {
			err = utils.WrapError("Encode Ms2048", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms2560:
		if err = w.WriteInteger(int64(ie.Ms2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode Ms2560", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms5120:
		if err = w.WriteInteger(int64(ie.Ms5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode Ms5120", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms10240:
		if err = w.WriteInteger(int64(ie.Ms10240), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode Ms10240", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms10:
		var tmp_int_Ms10 int64
		if tmp_int_Ms10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Ms10", err)
		}
		ie.Ms10 = tmp_int_Ms10
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms20:
		var tmp_int_Ms20 int64
		if tmp_int_Ms20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Ms20", err)
		}
		ie.Ms20 = tmp_int_Ms20
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms32:
		var tmp_int_Ms32 int64
		if tmp_int_Ms32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Ms32", err)
		}
		ie.Ms32 = tmp_int_Ms32
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms40:
		var tmp_int_Ms40 int64
		if tmp_int_Ms40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Ms40", err)
		}
		ie.Ms40 = tmp_int_Ms40
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms60:
		var tmp_int_Ms60 int64
		if tmp_int_Ms60, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 59}, false); err != nil {
			return utils.WrapError("Decode Ms60", err)
		}
		ie.Ms60 = tmp_int_Ms60
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms64:
		var tmp_int_Ms64 int64
		if tmp_int_Ms64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Ms64", err)
		}
		ie.Ms64 = tmp_int_Ms64
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms70:
		var tmp_int_Ms70 int64
		if tmp_int_Ms70, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 69}, false); err != nil {
			return utils.WrapError("Decode Ms70", err)
		}
		ie.Ms70 = tmp_int_Ms70
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms80:
		var tmp_int_Ms80 int64
		if tmp_int_Ms80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Ms80", err)
		}
		ie.Ms80 = tmp_int_Ms80
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms128:
		var tmp_int_Ms128 int64
		if tmp_int_Ms128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode Ms128", err)
		}
		ie.Ms128 = tmp_int_Ms128
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms160:
		var tmp_int_Ms160 int64
		if tmp_int_Ms160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Ms160", err)
		}
		ie.Ms160 = tmp_int_Ms160
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms256:
		var tmp_int_Ms256 int64
		if tmp_int_Ms256, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode Ms256", err)
		}
		ie.Ms256 = tmp_int_Ms256
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms320:
		var tmp_int_Ms320 int64
		if tmp_int_Ms320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Ms320", err)
		}
		ie.Ms320 = tmp_int_Ms320
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms512:
		var tmp_int_Ms512 int64
		if tmp_int_Ms512, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			return utils.WrapError("Decode Ms512", err)
		}
		ie.Ms512 = tmp_int_Ms512
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms640:
		var tmp_int_Ms640 int64
		if tmp_int_Ms640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Ms640", err)
		}
		ie.Ms640 = tmp_int_Ms640
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms1024:
		var tmp_int_Ms1024 int64
		if tmp_int_Ms1024, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode Ms1024", err)
		}
		ie.Ms1024 = tmp_int_Ms1024
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms1280:
		var tmp_int_Ms1280 int64
		if tmp_int_Ms1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode Ms1280", err)
		}
		ie.Ms1280 = tmp_int_Ms1280
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms2048:
		var tmp_int_Ms2048 int64
		if tmp_int_Ms2048, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2047}, false); err != nil {
			return utils.WrapError("Decode Ms2048", err)
		}
		ie.Ms2048 = tmp_int_Ms2048
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms2560:
		var tmp_int_Ms2560 int64
		if tmp_int_Ms2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode Ms2560", err)
		}
		ie.Ms2560 = tmp_int_Ms2560
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms5120:
		var tmp_int_Ms5120 int64
		if tmp_int_Ms5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode Ms5120", err)
		}
		ie.Ms5120 = tmp_int_Ms5120
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17_Choice_Ms10240:
		var tmp_int_Ms10240 int64
		if tmp_int_Ms10240, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode Ms10240", err)
		}
		ie.Ms10240 = tmp_int_Ms10240
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
