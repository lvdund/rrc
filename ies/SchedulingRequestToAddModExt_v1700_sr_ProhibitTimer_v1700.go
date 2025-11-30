package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms192  aper.Enumerated = 0
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms256  aper.Enumerated = 1
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms320  aper.Enumerated = 2
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms384  aper.Enumerated = 3
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms448  aper.Enumerated = 4
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms512  aper.Enumerated = 5
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms576  aper.Enumerated = 6
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms640  aper.Enumerated = 7
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms1082 aper.Enumerated = 8
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare7 aper.Enumerated = 9
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare6 aper.Enumerated = 10
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare5 aper.Enumerated = 11
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare4 aper.Enumerated = 12
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare3 aper.Enumerated = 13
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare2 aper.Enumerated = 14
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare1 aper.Enumerated = 15
)

type SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
