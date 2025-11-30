package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceConfigExt_v1700 struct {
	PeriodicityAndOffset_r17 *SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17 `lb:0,ub:1279,optional`
}

func (ie *SchedulingRequestResourceConfigExt_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PeriodicityAndOffset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PeriodicityAndOffset_r17 != nil {
		if err = ie.PeriodicityAndOffset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndOffset_r17", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1700) Decode(r *aper.AperReader) error {
	var err error
	var PeriodicityAndOffset_r17Present bool
	if PeriodicityAndOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PeriodicityAndOffset_r17Present {
		ie.PeriodicityAndOffset_r17 = new(SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17)
		if err = ie.PeriodicityAndOffset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndOffset_r17", err)
		}
	}
	return nil
}
