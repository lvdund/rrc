package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasConfig struct {
	MeasIdList       *MeasIdToAddModList            `optional`
	MeasObjectList   *MeasObjectToAddModList        `optional`
	ReportConfigList *ReportConfigToAddModList      `optional`
	QuantityConfig   *QuantityConfig                `optional`
	S_MeasureConfig  *VarMeasConfig_s_MeasureConfig `optional`
}

func (ie *VarMeasConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasIdList != nil, ie.MeasObjectList != nil, ie.ReportConfigList != nil, ie.QuantityConfig != nil, ie.S_MeasureConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasIdList != nil {
		if err = ie.MeasIdList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdList", err)
		}
	}
	if ie.MeasObjectList != nil {
		if err = ie.MeasObjectList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasObjectList", err)
		}
	}
	if ie.ReportConfigList != nil {
		if err = ie.ReportConfigList.Encode(w); err != nil {
			return utils.WrapError("Encode ReportConfigList", err)
		}
	}
	if ie.QuantityConfig != nil {
		if err = ie.QuantityConfig.Encode(w); err != nil {
			return utils.WrapError("Encode QuantityConfig", err)
		}
	}
	if ie.S_MeasureConfig != nil {
		if err = ie.S_MeasureConfig.Encode(w); err != nil {
			return utils.WrapError("Encode S_MeasureConfig", err)
		}
	}
	return nil
}

func (ie *VarMeasConfig) Decode(r *aper.AperReader) error {
	var err error
	var MeasIdListPresent bool
	if MeasIdListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasObjectListPresent bool
	if MeasObjectListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportConfigListPresent bool
	if ReportConfigListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var QuantityConfigPresent bool
	if QuantityConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var S_MeasureConfigPresent bool
	if S_MeasureConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasIdListPresent {
		ie.MeasIdList = new(MeasIdToAddModList)
		if err = ie.MeasIdList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdList", err)
		}
	}
	if MeasObjectListPresent {
		ie.MeasObjectList = new(MeasObjectToAddModList)
		if err = ie.MeasObjectList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectList", err)
		}
	}
	if ReportConfigListPresent {
		ie.ReportConfigList = new(ReportConfigToAddModList)
		if err = ie.ReportConfigList.Decode(r); err != nil {
			return utils.WrapError("Decode ReportConfigList", err)
		}
	}
	if QuantityConfigPresent {
		ie.QuantityConfig = new(QuantityConfig)
		if err = ie.QuantityConfig.Decode(r); err != nil {
			return utils.WrapError("Decode QuantityConfig", err)
		}
	}
	if S_MeasureConfigPresent {
		ie.S_MeasureConfig = new(VarMeasConfig_s_MeasureConfig)
		if err = ie.S_MeasureConfig.Decode(r); err != nil {
			return utils.WrapError("Decode S_MeasureConfig", err)
		}
	}
	return nil
}
