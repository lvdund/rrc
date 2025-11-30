package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DormantBWP_Config_r16 struct {
	DormantBWP_Id_r16           *BWP_Id                      `optional`
	WithinActiveTimeConfig_r16  *WithinActiveTimeConfig_r16  `optional,setuprelease`
	OutsideActiveTimeConfig_r16 *OutsideActiveTimeConfig_r16 `optional,setuprelease`
}

func (ie *DormantBWP_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DormantBWP_Id_r16 != nil, ie.WithinActiveTimeConfig_r16 != nil, ie.OutsideActiveTimeConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DormantBWP_Id_r16 != nil {
		if err = ie.DormantBWP_Id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DormantBWP_Id_r16", err)
		}
	}
	if ie.WithinActiveTimeConfig_r16 != nil {
		tmp_WithinActiveTimeConfig_r16 := utils.SetupRelease[*WithinActiveTimeConfig_r16]{
			Setup: ie.WithinActiveTimeConfig_r16,
		}
		if err = tmp_WithinActiveTimeConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode WithinActiveTimeConfig_r16", err)
		}
	}
	if ie.OutsideActiveTimeConfig_r16 != nil {
		tmp_OutsideActiveTimeConfig_r16 := utils.SetupRelease[*OutsideActiveTimeConfig_r16]{
			Setup: ie.OutsideActiveTimeConfig_r16,
		}
		if err = tmp_OutsideActiveTimeConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OutsideActiveTimeConfig_r16", err)
		}
	}
	return nil
}

func (ie *DormantBWP_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var DormantBWP_Id_r16Present bool
	if DormantBWP_Id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var WithinActiveTimeConfig_r16Present bool
	if WithinActiveTimeConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OutsideActiveTimeConfig_r16Present bool
	if OutsideActiveTimeConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DormantBWP_Id_r16Present {
		ie.DormantBWP_Id_r16 = new(BWP_Id)
		if err = ie.DormantBWP_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DormantBWP_Id_r16", err)
		}
	}
	if WithinActiveTimeConfig_r16Present {
		tmp_WithinActiveTimeConfig_r16 := utils.SetupRelease[*WithinActiveTimeConfig_r16]{}
		if err = tmp_WithinActiveTimeConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode WithinActiveTimeConfig_r16", err)
		}
		ie.WithinActiveTimeConfig_r16 = tmp_WithinActiveTimeConfig_r16.Setup
	}
	if OutsideActiveTimeConfig_r16Present {
		tmp_OutsideActiveTimeConfig_r16 := utils.SetupRelease[*OutsideActiveTimeConfig_r16]{}
		if err = tmp_OutsideActiveTimeConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OutsideActiveTimeConfig_r16", err)
		}
		ie.OutsideActiveTimeConfig_r16 = tmp_OutsideActiveTimeConfig_r16.Setup
	}
	return nil
}
