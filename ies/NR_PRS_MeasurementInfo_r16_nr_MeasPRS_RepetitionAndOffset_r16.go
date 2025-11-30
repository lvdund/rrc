package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_nothing uint64 = iota
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms20_r16
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms40_r16
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms80_r16
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms160_r16
)

type NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16 struct {
	Choice    uint64
	Ms20_r16  int64 `lb:0,ub:19,madatory`
	Ms40_r16  int64 `lb:0,ub:39,madatory`
	Ms80_r16  int64 `lb:0,ub:79,madatory`
	Ms160_r16 int64 `lb:0,ub:159,madatory`
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms20_r16:
		if err = w.WriteInteger(int64(ie.Ms20_r16), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Ms20_r16", err)
		}
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms40_r16:
		if err = w.WriteInteger(int64(ie.Ms40_r16), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Ms40_r16", err)
		}
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms80_r16:
		if err = w.WriteInteger(int64(ie.Ms80_r16), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Ms80_r16", err)
		}
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms160_r16:
		if err = w.WriteInteger(int64(ie.Ms160_r16), &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Ms160_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms20_r16:
		var tmp_int_Ms20_r16 int64
		if tmp_int_Ms20_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Ms20_r16", err)
		}
		ie.Ms20_r16 = tmp_int_Ms20_r16
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms40_r16:
		var tmp_int_Ms40_r16 int64
		if tmp_int_Ms40_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Ms40_r16", err)
		}
		ie.Ms40_r16 = tmp_int_Ms40_r16
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms80_r16:
		var tmp_int_Ms80_r16 int64
		if tmp_int_Ms80_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Ms80_r16", err)
		}
		ie.Ms80_r16 = tmp_int_Ms80_r16
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_Ms160_r16:
		var tmp_int_Ms160_r16 int64
		if tmp_int_Ms160_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Ms160_r16", err)
		}
		ie.Ms160_r16 = tmp_int_Ms160_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
