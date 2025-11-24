package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC struct {
	SupportedSubcarrierSpacingUL    SubcarrierSpacing                      `madatory`
	SupportedBandwidthUL            SupportedBandwidth                     `madatory`
	ChannelBW_90mhz                 *FeatureSetUplinkPerCC_channelBW_90mhz `optional`
	Mimo_CB_PUSCH                   *FeatureSetUplinkPerCC_mimo_CB_PUSCH   `lb:1,ub:2,optional`
	MaxNumberMIMO_LayersNonCB_PUSCH *MIMO_LayersUL                         `optional`
	SupportedModulationOrderUL      *ModulationOrder                       `optional`
}

func (ie *FeatureSetUplinkPerCC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ChannelBW_90mhz != nil, ie.Mimo_CB_PUSCH != nil, ie.MaxNumberMIMO_LayersNonCB_PUSCH != nil, ie.SupportedModulationOrderUL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SupportedSubcarrierSpacingUL.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedSubcarrierSpacingUL", err)
	}
	if err = ie.SupportedBandwidthUL.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedBandwidthUL", err)
	}
	if ie.ChannelBW_90mhz != nil {
		if err = ie.ChannelBW_90mhz.Encode(w); err != nil {
			return utils.WrapError("Encode ChannelBW_90mhz", err)
		}
	}
	if ie.Mimo_CB_PUSCH != nil {
		if err = ie.Mimo_CB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Mimo_CB_PUSCH", err)
		}
	}
	if ie.MaxNumberMIMO_LayersNonCB_PUSCH != nil {
		if err = ie.MaxNumberMIMO_LayersNonCB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberMIMO_LayersNonCB_PUSCH", err)
		}
	}
	if ie.SupportedModulationOrderUL != nil {
		if err = ie.SupportedModulationOrderUL.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedModulationOrderUL", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC) Decode(r *uper.UperReader) error {
	var err error
	var ChannelBW_90mhzPresent bool
	if ChannelBW_90mhzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mimo_CB_PUSCHPresent bool
	if Mimo_CB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberMIMO_LayersNonCB_PUSCHPresent bool
	if MaxNumberMIMO_LayersNonCB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedModulationOrderULPresent bool
	if SupportedModulationOrderULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SupportedSubcarrierSpacingUL.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedSubcarrierSpacingUL", err)
	}
	if err = ie.SupportedBandwidthUL.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedBandwidthUL", err)
	}
	if ChannelBW_90mhzPresent {
		ie.ChannelBW_90mhz = new(FeatureSetUplinkPerCC_channelBW_90mhz)
		if err = ie.ChannelBW_90mhz.Decode(r); err != nil {
			return utils.WrapError("Decode ChannelBW_90mhz", err)
		}
	}
	if Mimo_CB_PUSCHPresent {
		ie.Mimo_CB_PUSCH = new(FeatureSetUplinkPerCC_mimo_CB_PUSCH)
		if err = ie.Mimo_CB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Mimo_CB_PUSCH", err)
		}
	}
	if MaxNumberMIMO_LayersNonCB_PUSCHPresent {
		ie.MaxNumberMIMO_LayersNonCB_PUSCH = new(MIMO_LayersUL)
		if err = ie.MaxNumberMIMO_LayersNonCB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberMIMO_LayersNonCB_PUSCH", err)
		}
	}
	if SupportedModulationOrderULPresent {
		ie.SupportedModulationOrderUL = new(ModulationOrder)
		if err = ie.SupportedModulationOrderUL.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedModulationOrderUL", err)
		}
	}
	return nil
}
