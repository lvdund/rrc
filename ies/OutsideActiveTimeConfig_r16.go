package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OutsideActiveTimeConfig_r16 struct {
	FirstOutsideActiveTimeBWP_Id_r16   *BWP_Id              `optional`
	DormancyGroupOutsideActiveTime_r16 *DormancyGroupID_r16 `optional`
}

func (ie *OutsideActiveTimeConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FirstOutsideActiveTimeBWP_Id_r16 != nil, ie.DormancyGroupOutsideActiveTime_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FirstOutsideActiveTimeBWP_Id_r16 != nil {
		if err = ie.FirstOutsideActiveTimeBWP_Id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FirstOutsideActiveTimeBWP_Id_r16", err)
		}
	}
	if ie.DormancyGroupOutsideActiveTime_r16 != nil {
		if err = ie.DormancyGroupOutsideActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DormancyGroupOutsideActiveTime_r16", err)
		}
	}
	return nil
}

func (ie *OutsideActiveTimeConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var FirstOutsideActiveTimeBWP_Id_r16Present bool
	if FirstOutsideActiveTimeBWP_Id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DormancyGroupOutsideActiveTime_r16Present bool
	if DormancyGroupOutsideActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FirstOutsideActiveTimeBWP_Id_r16Present {
		ie.FirstOutsideActiveTimeBWP_Id_r16 = new(BWP_Id)
		if err = ie.FirstOutsideActiveTimeBWP_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FirstOutsideActiveTimeBWP_Id_r16", err)
		}
	}
	if DormancyGroupOutsideActiveTime_r16Present {
		ie.DormancyGroupOutsideActiveTime_r16 = new(DormancyGroupID_r16)
		if err = ie.DormancyGroupOutsideActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DormancyGroupOutsideActiveTime_r16", err)
		}
	}
	return nil
}
