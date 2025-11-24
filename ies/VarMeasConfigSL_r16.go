package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasConfigSL_r16 struct {
	Sl_MeasIdList_r16       *SL_MeasIdList_r16       `optional`
	Sl_MeasObjectList_r16   *SL_MeasObjectList_r16   `optional`
	Sl_reportConfigList_r16 *SL_ReportConfigList_r16 `optional`
	Sl_QuantityConfig_r16   *SL_QuantityConfig_r16   `optional`
}

func (ie *VarMeasConfigSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_MeasIdList_r16 != nil, ie.Sl_MeasObjectList_r16 != nil, ie.Sl_reportConfigList_r16 != nil, ie.Sl_QuantityConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_MeasIdList_r16 != nil {
		if err = ie.Sl_MeasIdList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasIdList_r16", err)
		}
	}
	if ie.Sl_MeasObjectList_r16 != nil {
		if err = ie.Sl_MeasObjectList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasObjectList_r16", err)
		}
	}
	if ie.Sl_reportConfigList_r16 != nil {
		if err = ie.Sl_reportConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_reportConfigList_r16", err)
		}
	}
	if ie.Sl_QuantityConfig_r16 != nil {
		if err = ie.Sl_QuantityConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_QuantityConfig_r16", err)
		}
	}
	return nil
}

func (ie *VarMeasConfigSL_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_MeasIdList_r16Present bool
	if Sl_MeasIdList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasObjectList_r16Present bool
	if Sl_MeasObjectList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_reportConfigList_r16Present bool
	if Sl_reportConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_QuantityConfig_r16Present bool
	if Sl_QuantityConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_MeasIdList_r16Present {
		ie.Sl_MeasIdList_r16 = new(SL_MeasIdList_r16)
		if err = ie.Sl_MeasIdList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasIdList_r16", err)
		}
	}
	if Sl_MeasObjectList_r16Present {
		ie.Sl_MeasObjectList_r16 = new(SL_MeasObjectList_r16)
		if err = ie.Sl_MeasObjectList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasObjectList_r16", err)
		}
	}
	if Sl_reportConfigList_r16Present {
		ie.Sl_reportConfigList_r16 = new(SL_ReportConfigList_r16)
		if err = ie.Sl_reportConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_reportConfigList_r16", err)
		}
	}
	if Sl_QuantityConfig_r16Present {
		ie.Sl_QuantityConfig_r16 = new(SL_QuantityConfig_r16)
		if err = ie.Sl_QuantityConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_QuantityConfig_r16", err)
		}
	}
	return nil
}
