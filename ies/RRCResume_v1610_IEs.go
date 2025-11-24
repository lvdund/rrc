package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResume_v1610_IEs struct {
	IdleModeMeasurementReq_r16  *RRCResume_v1610_IEs_idleModeMeasurementReq_r16  `optional`
	RestoreMCG_SCells_r16       *RRCResume_v1610_IEs_restoreMCG_SCells_r16       `optional`
	RestoreSCG_r16              *RRCResume_v1610_IEs_restoreSCG_r16              `optional`
	Mrdc_SecondaryCellGroup_r16 *RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16 `optional`
	NeedForGapsConfigNR_r16     *NeedForGapsConfigNR_r16                         `optional,setuprelease`
	NonCriticalExtension        *RRCResume_v1700_IEs                             `optional`
}

func (ie *RRCResume_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IdleModeMeasurementReq_r16 != nil, ie.RestoreMCG_SCells_r16 != nil, ie.RestoreSCG_r16 != nil, ie.Mrdc_SecondaryCellGroup_r16 != nil, ie.NeedForGapsConfigNR_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IdleModeMeasurementReq_r16 != nil {
		if err = ie.IdleModeMeasurementReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleModeMeasurementReq_r16", err)
		}
	}
	if ie.RestoreMCG_SCells_r16 != nil {
		if err = ie.RestoreMCG_SCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RestoreMCG_SCells_r16", err)
		}
	}
	if ie.RestoreSCG_r16 != nil {
		if err = ie.RestoreSCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RestoreSCG_r16", err)
		}
	}
	if ie.Mrdc_SecondaryCellGroup_r16 != nil {
		if err = ie.Mrdc_SecondaryCellGroup_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_SecondaryCellGroup_r16", err)
		}
	}
	if ie.NeedForGapsConfigNR_r16 != nil {
		tmp_NeedForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{
			Setup: ie.NeedForGapsConfigNR_r16,
		}
		if err = tmp_NeedForGapsConfigNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapsConfigNR_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResume_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var IdleModeMeasurementReq_r16Present bool
	if IdleModeMeasurementReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RestoreMCG_SCells_r16Present bool
	if RestoreMCG_SCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RestoreSCG_r16Present bool
	if RestoreSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_SecondaryCellGroup_r16Present bool
	if Mrdc_SecondaryCellGroup_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapsConfigNR_r16Present bool
	if NeedForGapsConfigNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IdleModeMeasurementReq_r16Present {
		ie.IdleModeMeasurementReq_r16 = new(RRCResume_v1610_IEs_idleModeMeasurementReq_r16)
		if err = ie.IdleModeMeasurementReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleModeMeasurementReq_r16", err)
		}
	}
	if RestoreMCG_SCells_r16Present {
		ie.RestoreMCG_SCells_r16 = new(RRCResume_v1610_IEs_restoreMCG_SCells_r16)
		if err = ie.RestoreMCG_SCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RestoreMCG_SCells_r16", err)
		}
	}
	if RestoreSCG_r16Present {
		ie.RestoreSCG_r16 = new(RRCResume_v1610_IEs_restoreSCG_r16)
		if err = ie.RestoreSCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RestoreSCG_r16", err)
		}
	}
	if Mrdc_SecondaryCellGroup_r16Present {
		ie.Mrdc_SecondaryCellGroup_r16 = new(RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16)
		if err = ie.Mrdc_SecondaryCellGroup_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_SecondaryCellGroup_r16", err)
		}
	}
	if NeedForGapsConfigNR_r16Present {
		tmp_NeedForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{}
		if err = tmp_NeedForGapsConfigNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapsConfigNR_r16", err)
		}
		ie.NeedForGapsConfigNR_r16 = tmp_NeedForGapsConfigNR_r16.Setup
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCResume_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
