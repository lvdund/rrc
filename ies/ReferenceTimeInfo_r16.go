package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceTimeInfo_r16 struct {
	Time_r16         ReferenceTime_r16                       `madatory`
	Uncertainty_r16  *int64                                  `lb:0,ub:32767,optional`
	TimeInfoType_r16 *ReferenceTimeInfo_r16_timeInfoType_r16 `optional`
	ReferenceSFN_r16 *int64                                  `lb:0,ub:1023,optional`
}

func (ie *ReferenceTimeInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Uncertainty_r16 != nil, ie.TimeInfoType_r16 != nil, ie.ReferenceSFN_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Time_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Time_r16", err)
	}
	if ie.Uncertainty_r16 != nil {
		if err = w.WriteInteger(*ie.Uncertainty_r16, &aper.Constraint{Lb: 0, Ub: 32767}, false); err != nil {
			return utils.WrapError("Encode Uncertainty_r16", err)
		}
	}
	if ie.TimeInfoType_r16 != nil {
		if err = ie.TimeInfoType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TimeInfoType_r16", err)
		}
	}
	if ie.ReferenceSFN_r16 != nil {
		if err = w.WriteInteger(*ie.ReferenceSFN_r16, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode ReferenceSFN_r16", err)
		}
	}
	return nil
}

func (ie *ReferenceTimeInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var Uncertainty_r16Present bool
	if Uncertainty_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeInfoType_r16Present bool
	if TimeInfoType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReferenceSFN_r16Present bool
	if ReferenceSFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Time_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Time_r16", err)
	}
	if Uncertainty_r16Present {
		var tmp_int_Uncertainty_r16 int64
		if tmp_int_Uncertainty_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 32767}, false); err != nil {
			return utils.WrapError("Decode Uncertainty_r16", err)
		}
		ie.Uncertainty_r16 = &tmp_int_Uncertainty_r16
	}
	if TimeInfoType_r16Present {
		ie.TimeInfoType_r16 = new(ReferenceTimeInfo_r16_timeInfoType_r16)
		if err = ie.TimeInfoType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TimeInfoType_r16", err)
		}
	}
	if ReferenceSFN_r16Present {
		var tmp_int_ReferenceSFN_r16 int64
		if tmp_int_ReferenceSFN_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode ReferenceSFN_r16", err)
		}
		ie.ReferenceSFN_r16 = &tmp_int_ReferenceSFN_r16
	}
	return nil
}
