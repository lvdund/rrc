package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceConfigExt_v1610 struct {
	Phy_PriorityIndex_r16 *SchedulingRequestResourceConfigExt_v1610_phy_PriorityIndex_r16 `optional`
}

func (ie *SchedulingRequestResourceConfigExt_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Phy_PriorityIndex_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Phy_PriorityIndex_r16 != nil {
		if err = ie.Phy_PriorityIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_PriorityIndex_r16", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1610) Decode(r *aper.AperReader) error {
	var err error
	var Phy_PriorityIndex_r16Present bool
	if Phy_PriorityIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Phy_PriorityIndex_r16Present {
		ie.Phy_PriorityIndex_r16 = new(SchedulingRequestResourceConfigExt_v1610_phy_PriorityIndex_r16)
		if err = ie.Phy_PriorityIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_PriorityIndex_r16", err)
		}
	}
	return nil
}
