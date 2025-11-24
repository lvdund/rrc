package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventD1_r17 struct {
	DistanceThreshFromReference1_r17 int64                  `lb:0,ub:65525,madatory`
	DistanceThreshFromReference2_r17 int64                  `lb:0,ub:65525,madatory`
	ReferenceLocation1_r17           ReferenceLocation_r17  `madatory`
	ReferenceLocation2_r17           ReferenceLocation_r17  `madatory`
	HysteresisLocation_r17           HysteresisLocation_r17 `madatory`
	TimeToTrigger_r17                TimeToTrigger          `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventD1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.DistanceThreshFromReference1_r17, &uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("WriteInteger DistanceThreshFromReference1_r17", err)
	}
	if err = w.WriteInteger(ie.DistanceThreshFromReference2_r17, &uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("WriteInteger DistanceThreshFromReference2_r17", err)
	}
	if err = ie.ReferenceLocation1_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceLocation1_r17", err)
	}
	if err = ie.ReferenceLocation2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceLocation2_r17", err)
	}
	if err = ie.HysteresisLocation_r17.Encode(w); err != nil {
		return utils.WrapError("Encode HysteresisLocation_r17", err)
	}
	if err = ie.TimeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode TimeToTrigger_r17", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventD1_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_DistanceThreshFromReference1_r17 int64
	if tmp_int_DistanceThreshFromReference1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("ReadInteger DistanceThreshFromReference1_r17", err)
	}
	ie.DistanceThreshFromReference1_r17 = tmp_int_DistanceThreshFromReference1_r17
	var tmp_int_DistanceThreshFromReference2_r17 int64
	if tmp_int_DistanceThreshFromReference2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("ReadInteger DistanceThreshFromReference2_r17", err)
	}
	ie.DistanceThreshFromReference2_r17 = tmp_int_DistanceThreshFromReference2_r17
	if err = ie.ReferenceLocation1_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceLocation1_r17", err)
	}
	if err = ie.ReferenceLocation2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceLocation2_r17", err)
	}
	if err = ie.HysteresisLocation_r17.Decode(r); err != nil {
		return utils.WrapError("Decode HysteresisLocation_r17", err)
	}
	if err = ie.TimeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode TimeToTrigger_r17", err)
	}
	return nil
}
