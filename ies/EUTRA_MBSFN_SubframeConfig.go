package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_MBSFN_SubframeConfig struct {
	RadioframeAllocationPeriod EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod `madatory`
	RadioframeAllocationOffset int64                                                 `lb:0,ub:7,madatory`
	SubframeAllocation1        EUTRA_MBSFN_SubframeConfig_subframeAllocation1        `lb:6,ub:6,madatory`
	SubframeAllocation2        *EUTRA_MBSFN_SubframeConfig_subframeAllocation2       `lb:2,ub:2,optional`
}

func (ie *EUTRA_MBSFN_SubframeConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SubframeAllocation2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.RadioframeAllocationPeriod.Encode(w); err != nil {
		return utils.WrapError("Encode RadioframeAllocationPeriod", err)
	}
	if err = w.WriteInteger(ie.RadioframeAllocationOffset, &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger RadioframeAllocationOffset", err)
	}
	if err = ie.SubframeAllocation1.Encode(w); err != nil {
		return utils.WrapError("Encode SubframeAllocation1", err)
	}
	if ie.SubframeAllocation2 != nil {
		if err = ie.SubframeAllocation2.Encode(w); err != nil {
			return utils.WrapError("Encode SubframeAllocation2", err)
		}
	}
	return nil
}

func (ie *EUTRA_MBSFN_SubframeConfig) Decode(r *aper.AperReader) error {
	var err error
	var SubframeAllocation2Present bool
	if SubframeAllocation2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.RadioframeAllocationPeriod.Decode(r); err != nil {
		return utils.WrapError("Decode RadioframeAllocationPeriod", err)
	}
	var tmp_int_RadioframeAllocationOffset int64
	if tmp_int_RadioframeAllocationOffset, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger RadioframeAllocationOffset", err)
	}
	ie.RadioframeAllocationOffset = tmp_int_RadioframeAllocationOffset
	if err = ie.SubframeAllocation1.Decode(r); err != nil {
		return utils.WrapError("Decode SubframeAllocation1", err)
	}
	if SubframeAllocation2Present {
		ie.SubframeAllocation2 = new(EUTRA_MBSFN_SubframeConfig_subframeAllocation2)
		if err = ie.SubframeAllocation2.Decode(r); err != nil {
			return utils.WrapError("Decode SubframeAllocation2", err)
		}
	}
	return nil
}
