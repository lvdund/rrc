package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestToAddModExt_v1700 struct {
	Sr_ProhibitTimer_v1700 *SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700 `optional`
}

func (ie *SchedulingRequestToAddModExt_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sr_ProhibitTimer_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sr_ProhibitTimer_v1700 != nil {
		if err = ie.Sr_ProhibitTimer_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Sr_ProhibitTimer_v1700", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestToAddModExt_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Sr_ProhibitTimer_v1700Present bool
	if Sr_ProhibitTimer_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sr_ProhibitTimer_v1700Present {
		ie.Sr_ProhibitTimer_v1700 = new(SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700)
		if err = ie.Sr_ProhibitTimer_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sr_ProhibitTimer_v1700", err)
		}
	}
	return nil
}
