package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasConfigCommon_r16 struct {
	Sl_MeasObjectListCommon_r16   *SL_MeasObjectList_r16   `optional`
	Sl_ReportConfigListCommon_r16 *SL_ReportConfigList_r16 `optional`
	Sl_MeasIdListCommon_r16       *SL_MeasIdList_r16       `optional`
	Sl_QuantityConfigCommon_r16   *SL_QuantityConfig_r16   `optional`
}

func (ie *SL_MeasConfigCommon_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_MeasObjectListCommon_r16 != nil, ie.Sl_ReportConfigListCommon_r16 != nil, ie.Sl_MeasIdListCommon_r16 != nil, ie.Sl_QuantityConfigCommon_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_MeasObjectListCommon_r16 != nil {
		if err = ie.Sl_MeasObjectListCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasObjectListCommon_r16", err)
		}
	}
	if ie.Sl_ReportConfigListCommon_r16 != nil {
		if err = ie.Sl_ReportConfigListCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ReportConfigListCommon_r16", err)
		}
	}
	if ie.Sl_MeasIdListCommon_r16 != nil {
		if err = ie.Sl_MeasIdListCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasIdListCommon_r16", err)
		}
	}
	if ie.Sl_QuantityConfigCommon_r16 != nil {
		if err = ie.Sl_QuantityConfigCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_QuantityConfigCommon_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasConfigCommon_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_MeasObjectListCommon_r16Present bool
	if Sl_MeasObjectListCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ReportConfigListCommon_r16Present bool
	if Sl_ReportConfigListCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasIdListCommon_r16Present bool
	if Sl_MeasIdListCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_QuantityConfigCommon_r16Present bool
	if Sl_QuantityConfigCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_MeasObjectListCommon_r16Present {
		ie.Sl_MeasObjectListCommon_r16 = new(SL_MeasObjectList_r16)
		if err = ie.Sl_MeasObjectListCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasObjectListCommon_r16", err)
		}
	}
	if Sl_ReportConfigListCommon_r16Present {
		ie.Sl_ReportConfigListCommon_r16 = new(SL_ReportConfigList_r16)
		if err = ie.Sl_ReportConfigListCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ReportConfigListCommon_r16", err)
		}
	}
	if Sl_MeasIdListCommon_r16Present {
		ie.Sl_MeasIdListCommon_r16 = new(SL_MeasIdList_r16)
		if err = ie.Sl_MeasIdListCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasIdListCommon_r16", err)
		}
	}
	if Sl_QuantityConfigCommon_r16Present {
		ie.Sl_QuantityConfigCommon_r16 = new(SL_QuantityConfig_r16)
		if err = ie.Sl_QuantityConfigCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_QuantityConfigCommon_r16", err)
		}
	}
	return nil
}
