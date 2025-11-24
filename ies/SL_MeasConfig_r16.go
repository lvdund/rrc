package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasConfig_r16 struct {
	Sl_MeasObjectToRemoveList_r16   *SL_MeasObjectToRemoveList_r16   `optional`
	Sl_MeasObjectToAddModList_r16   *SL_MeasObjectList_r16           `optional`
	Sl_ReportConfigToRemoveList_r16 *SL_ReportConfigToRemoveList_r16 `optional`
	Sl_ReportConfigToAddModList_r16 *SL_ReportConfigList_r16         `optional`
	Sl_MeasIdToRemoveList_r16       *SL_MeasIdToRemoveList_r16       `optional`
	Sl_MeasIdToAddModList_r16       *SL_MeasIdList_r16               `optional`
	Sl_QuantityConfig_r16           *SL_QuantityConfig_r16           `optional`
}

func (ie *SL_MeasConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_MeasObjectToRemoveList_r16 != nil, ie.Sl_MeasObjectToAddModList_r16 != nil, ie.Sl_ReportConfigToRemoveList_r16 != nil, ie.Sl_ReportConfigToAddModList_r16 != nil, ie.Sl_MeasIdToRemoveList_r16 != nil, ie.Sl_MeasIdToAddModList_r16 != nil, ie.Sl_QuantityConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_MeasObjectToRemoveList_r16 != nil {
		if err = ie.Sl_MeasObjectToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasObjectToRemoveList_r16", err)
		}
	}
	if ie.Sl_MeasObjectToAddModList_r16 != nil {
		if err = ie.Sl_MeasObjectToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasObjectToAddModList_r16", err)
		}
	}
	if ie.Sl_ReportConfigToRemoveList_r16 != nil {
		if err = ie.Sl_ReportConfigToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ReportConfigToRemoveList_r16", err)
		}
	}
	if ie.Sl_ReportConfigToAddModList_r16 != nil {
		if err = ie.Sl_ReportConfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ReportConfigToAddModList_r16", err)
		}
	}
	if ie.Sl_MeasIdToRemoveList_r16 != nil {
		if err = ie.Sl_MeasIdToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasIdToRemoveList_r16", err)
		}
	}
	if ie.Sl_MeasIdToAddModList_r16 != nil {
		if err = ie.Sl_MeasIdToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasIdToAddModList_r16", err)
		}
	}
	if ie.Sl_QuantityConfig_r16 != nil {
		if err = ie.Sl_QuantityConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_QuantityConfig_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_MeasObjectToRemoveList_r16Present bool
	if Sl_MeasObjectToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasObjectToAddModList_r16Present bool
	if Sl_MeasObjectToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ReportConfigToRemoveList_r16Present bool
	if Sl_ReportConfigToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ReportConfigToAddModList_r16Present bool
	if Sl_ReportConfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasIdToRemoveList_r16Present bool
	if Sl_MeasIdToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasIdToAddModList_r16Present bool
	if Sl_MeasIdToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_QuantityConfig_r16Present bool
	if Sl_QuantityConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_MeasObjectToRemoveList_r16Present {
		ie.Sl_MeasObjectToRemoveList_r16 = new(SL_MeasObjectToRemoveList_r16)
		if err = ie.Sl_MeasObjectToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasObjectToRemoveList_r16", err)
		}
	}
	if Sl_MeasObjectToAddModList_r16Present {
		ie.Sl_MeasObjectToAddModList_r16 = new(SL_MeasObjectList_r16)
		if err = ie.Sl_MeasObjectToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasObjectToAddModList_r16", err)
		}
	}
	if Sl_ReportConfigToRemoveList_r16Present {
		ie.Sl_ReportConfigToRemoveList_r16 = new(SL_ReportConfigToRemoveList_r16)
		if err = ie.Sl_ReportConfigToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ReportConfigToRemoveList_r16", err)
		}
	}
	if Sl_ReportConfigToAddModList_r16Present {
		ie.Sl_ReportConfigToAddModList_r16 = new(SL_ReportConfigList_r16)
		if err = ie.Sl_ReportConfigToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ReportConfigToAddModList_r16", err)
		}
	}
	if Sl_MeasIdToRemoveList_r16Present {
		ie.Sl_MeasIdToRemoveList_r16 = new(SL_MeasIdToRemoveList_r16)
		if err = ie.Sl_MeasIdToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasIdToRemoveList_r16", err)
		}
	}
	if Sl_MeasIdToAddModList_r16Present {
		ie.Sl_MeasIdToAddModList_r16 = new(SL_MeasIdList_r16)
		if err = ie.Sl_MeasIdToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasIdToAddModList_r16", err)
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
