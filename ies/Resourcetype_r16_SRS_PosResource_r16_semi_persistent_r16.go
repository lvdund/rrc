package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16 struct {
	PeriodicityAndOffset_sp_r16     SRS_PeriodicityAndOffset_r16     `madatory`
	PeriodicityAndOffset_sp_Ext_r16 *SRS_PeriodicityAndOffsetExt_r16 `optional`
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PeriodicityAndOffset_sp_Ext_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PeriodicityAndOffset_sp_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodicityAndOffset_sp_r16", err)
	}
	if ie.PeriodicityAndOffset_sp_Ext_r16 != nil {
		if err = ie.PeriodicityAndOffset_sp_Ext_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndOffset_sp_Ext_r16", err)
		}
	}
	return nil
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16) Decode(r *uper.UperReader) error {
	var err error
	var PeriodicityAndOffset_sp_Ext_r16Present bool
	if PeriodicityAndOffset_sp_Ext_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PeriodicityAndOffset_sp_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodicityAndOffset_sp_r16", err)
	}
	if PeriodicityAndOffset_sp_Ext_r16Present {
		ie.PeriodicityAndOffset_sp_Ext_r16 = new(SRS_PeriodicityAndOffsetExt_r16)
		if err = ie.PeriodicityAndOffset_sp_Ext_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndOffset_sp_Ext_r16", err)
		}
	}
	return nil
}
