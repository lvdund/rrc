package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_nothing uint64 = iota
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl32
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl64
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl128
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl5120
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl10240
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl20480
)

type SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710 struct {
	Choice  uint64
	Sl32    int64 `lb:0,ub:31,madatory`
	Sl64    int64 `lb:0,ub:63,madatory`
	Sl128   int64 `lb:0,ub:127,madatory`
	Sl5120  int64 `lb:0,ub:5119,madatory`
	Sl10240 int64 `lb:0,ub:10239,madatory`
	Sl20480 int64 `lb:0,ub:20479,madatory`
}

func (ie *SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl32:
		if err = w.WriteInteger(int64(ie.Sl32), &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Sl32", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl64:
		if err = w.WriteInteger(int64(ie.Sl64), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Sl64", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl128:
		if err = w.WriteInteger(int64(ie.Sl128), &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode Sl128", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl5120:
		if err = w.WriteInteger(int64(ie.Sl5120), &aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode Sl5120", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl10240:
		if err = w.WriteInteger(int64(ie.Sl10240), &aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode Sl10240", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl20480:
		if err = w.WriteInteger(int64(ie.Sl20480), &aper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode Sl20480", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl32:
		var tmp_int_Sl32 int64
		if tmp_int_Sl32, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Sl32", err)
		}
		ie.Sl32 = tmp_int_Sl32
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl64:
		var tmp_int_Sl64 int64
		if tmp_int_Sl64, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Sl64", err)
		}
		ie.Sl64 = tmp_int_Sl64
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl128:
		var tmp_int_Sl128 int64
		if tmp_int_Sl128, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode Sl128", err)
		}
		ie.Sl128 = tmp_int_Sl128
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl5120:
		var tmp_int_Sl5120 int64
		if tmp_int_Sl5120, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode Sl5120", err)
		}
		ie.Sl5120 = tmp_int_Sl5120
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl10240:
		var tmp_int_Sl10240 int64
		if tmp_int_Sl10240, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode Sl10240", err)
		}
		ie.Sl10240 = tmp_int_Sl10240
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_Sl20480:
		var tmp_int_Sl20480 int64
		if tmp_int_Sl20480, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode Sl20480", err)
		}
		ie.Sl20480 = tmp_int_Sl20480
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
