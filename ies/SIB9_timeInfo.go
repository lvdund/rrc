package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB9_timeInfo struct {
	TimeInfoUTC        int64           `lb:0,ub:549755813887,madatory`
	DayLightSavingTime *aper.BitString `lb:2,ub:2,optional`
	LeapSeconds        *int64          `lb:-127,ub:128,optional`
	LocalTimeOffset    *int64          `lb:-63,ub:64,optional`
}

func (ie *SIB9_timeInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DayLightSavingTime != nil, ie.LeapSeconds != nil, ie.LocalTimeOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.TimeInfoUTC, &aper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
		return utils.WrapError("WriteInteger TimeInfoUTC", err)
	}
	if ie.DayLightSavingTime != nil {
		if err = w.WriteBitString(ie.DayLightSavingTime.Bytes, uint(ie.DayLightSavingTime.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode DayLightSavingTime", err)
		}
	}
	if ie.LeapSeconds != nil {
		if err = w.WriteInteger(*ie.LeapSeconds, &aper.Constraint{Lb: -127, Ub: 128}, false); err != nil {
			return utils.WrapError("Encode LeapSeconds", err)
		}
	}
	if ie.LocalTimeOffset != nil {
		if err = w.WriteInteger(*ie.LocalTimeOffset, &aper.Constraint{Lb: -63, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode LocalTimeOffset", err)
		}
	}
	return nil
}

func (ie *SIB9_timeInfo) Decode(r *aper.AperReader) error {
	var err error
	var DayLightSavingTimePresent bool
	if DayLightSavingTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LeapSecondsPresent bool
	if LeapSecondsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LocalTimeOffsetPresent bool
	if LocalTimeOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_TimeInfoUTC int64
	if tmp_int_TimeInfoUTC, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
		return utils.WrapError("ReadInteger TimeInfoUTC", err)
	}
	ie.TimeInfoUTC = tmp_int_TimeInfoUTC
	if DayLightSavingTimePresent {
		var tmp_bs_DayLightSavingTime []byte
		var n_DayLightSavingTime uint
		if tmp_bs_DayLightSavingTime, n_DayLightSavingTime, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode DayLightSavingTime", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_DayLightSavingTime,
			NumBits: uint64(n_DayLightSavingTime),
		}
		ie.DayLightSavingTime = &tmp_bitstring
	}
	if LeapSecondsPresent {
		var tmp_int_LeapSeconds int64
		if tmp_int_LeapSeconds, err = r.ReadInteger(&aper.Constraint{Lb: -127, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode LeapSeconds", err)
		}
		ie.LeapSeconds = &tmp_int_LeapSeconds
	}
	if LocalTimeOffsetPresent {
		var tmp_int_LocalTimeOffset int64
		if tmp_int_LocalTimeOffset, err = r.ReadInteger(&aper.Constraint{Lb: -63, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode LocalTimeOffset", err)
		}
		ie.LocalTimeOffset = &tmp_int_LocalTimeOffset
	}
	return nil
}
