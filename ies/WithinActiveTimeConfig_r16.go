package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type WithinActiveTimeConfig_r16 struct {
	FirstWithinActiveTimeBWP_Id_r16   *BWP_Id              `optional`
	DormancyGroupWithinActiveTime_r16 *DormancyGroupID_r16 `optional`
}

func (ie *WithinActiveTimeConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FirstWithinActiveTimeBWP_Id_r16 != nil, ie.DormancyGroupWithinActiveTime_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FirstWithinActiveTimeBWP_Id_r16 != nil {
		if err = ie.FirstWithinActiveTimeBWP_Id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FirstWithinActiveTimeBWP_Id_r16", err)
		}
	}
	if ie.DormancyGroupWithinActiveTime_r16 != nil {
		if err = ie.DormancyGroupWithinActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DormancyGroupWithinActiveTime_r16", err)
		}
	}
	return nil
}

func (ie *WithinActiveTimeConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var FirstWithinActiveTimeBWP_Id_r16Present bool
	if FirstWithinActiveTimeBWP_Id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DormancyGroupWithinActiveTime_r16Present bool
	if DormancyGroupWithinActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FirstWithinActiveTimeBWP_Id_r16Present {
		ie.FirstWithinActiveTimeBWP_Id_r16 = new(BWP_Id)
		if err = ie.FirstWithinActiveTimeBWP_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FirstWithinActiveTimeBWP_Id_r16", err)
		}
	}
	if DormancyGroupWithinActiveTime_r16Present {
		ie.DormancyGroupWithinActiveTime_r16 = new(DormancyGroupID_r16)
		if err = ie.DormancyGroupWithinActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DormancyGroupWithinActiveTime_r16", err)
		}
	}
	return nil
}
