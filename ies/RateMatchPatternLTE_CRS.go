package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPatternLTE_CRS struct {
	CarrierFreqDL            int64                                      `lb:0,ub:16383,madatory`
	CarrierBandwidthDL       RateMatchPatternLTE_CRS_carrierBandwidthDL `madatory`
	Mbsfn_SubframeConfigList *EUTRA_MBSFN_SubframeConfigList            `optional`
	NrofCRS_Ports            RateMatchPatternLTE_CRS_nrofCRS_Ports      `madatory`
	V_Shift                  RateMatchPatternLTE_CRS_v_Shift            `madatory`
}

func (ie *RateMatchPatternLTE_CRS) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mbsfn_SubframeConfigList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.CarrierFreqDL, &aper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return utils.WrapError("WriteInteger CarrierFreqDL", err)
	}
	if err = ie.CarrierBandwidthDL.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierBandwidthDL", err)
	}
	if ie.Mbsfn_SubframeConfigList != nil {
		if err = ie.Mbsfn_SubframeConfigList.Encode(w); err != nil {
			return utils.WrapError("Encode Mbsfn_SubframeConfigList", err)
		}
	}
	if err = ie.NrofCRS_Ports.Encode(w); err != nil {
		return utils.WrapError("Encode NrofCRS_Ports", err)
	}
	if err = ie.V_Shift.Encode(w); err != nil {
		return utils.WrapError("Encode V_Shift", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS) Decode(r *aper.AperReader) error {
	var err error
	var Mbsfn_SubframeConfigListPresent bool
	if Mbsfn_SubframeConfigListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_CarrierFreqDL int64
	if tmp_int_CarrierFreqDL, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return utils.WrapError("ReadInteger CarrierFreqDL", err)
	}
	ie.CarrierFreqDL = tmp_int_CarrierFreqDL
	if err = ie.CarrierBandwidthDL.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierBandwidthDL", err)
	}
	if Mbsfn_SubframeConfigListPresent {
		ie.Mbsfn_SubframeConfigList = new(EUTRA_MBSFN_SubframeConfigList)
		if err = ie.Mbsfn_SubframeConfigList.Decode(r); err != nil {
			return utils.WrapError("Decode Mbsfn_SubframeConfigList", err)
		}
	}
	if err = ie.NrofCRS_Ports.Decode(r); err != nil {
		return utils.WrapError("Decode NrofCRS_Ports", err)
	}
	if err = ie.V_Shift.Decode(r); err != nil {
		return utils.WrapError("Decode V_Shift", err)
	}
	return nil
}
