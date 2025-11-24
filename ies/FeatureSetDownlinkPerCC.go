package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC struct {
	SupportedSubcarrierSpacingDL SubcarrierSpacing                        `madatory`
	SupportedBandwidthDL         SupportedBandwidth                       `madatory`
	ChannelBW_90mhz              *FeatureSetDownlinkPerCC_channelBW_90mhz `optional`
	MaxNumberMIMO_LayersPDSCH    *MIMO_LayersDL                           `optional`
	SupportedModulationOrderDL   *ModulationOrder                         `optional`
}

func (ie *FeatureSetDownlinkPerCC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ChannelBW_90mhz != nil, ie.MaxNumberMIMO_LayersPDSCH != nil, ie.SupportedModulationOrderDL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SupportedSubcarrierSpacingDL.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedSubcarrierSpacingDL", err)
	}
	if err = ie.SupportedBandwidthDL.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedBandwidthDL", err)
	}
	if ie.ChannelBW_90mhz != nil {
		if err = ie.ChannelBW_90mhz.Encode(w); err != nil {
			return utils.WrapError("Encode ChannelBW_90mhz", err)
		}
	}
	if ie.MaxNumberMIMO_LayersPDSCH != nil {
		if err = ie.MaxNumberMIMO_LayersPDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberMIMO_LayersPDSCH", err)
		}
	}
	if ie.SupportedModulationOrderDL != nil {
		if err = ie.SupportedModulationOrderDL.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedModulationOrderDL", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC) Decode(r *uper.UperReader) error {
	var err error
	var ChannelBW_90mhzPresent bool
	if ChannelBW_90mhzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberMIMO_LayersPDSCHPresent bool
	if MaxNumberMIMO_LayersPDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedModulationOrderDLPresent bool
	if SupportedModulationOrderDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SupportedSubcarrierSpacingDL.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedSubcarrierSpacingDL", err)
	}
	if err = ie.SupportedBandwidthDL.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedBandwidthDL", err)
	}
	if ChannelBW_90mhzPresent {
		ie.ChannelBW_90mhz = new(FeatureSetDownlinkPerCC_channelBW_90mhz)
		if err = ie.ChannelBW_90mhz.Decode(r); err != nil {
			return utils.WrapError("Decode ChannelBW_90mhz", err)
		}
	}
	if MaxNumberMIMO_LayersPDSCHPresent {
		ie.MaxNumberMIMO_LayersPDSCH = new(MIMO_LayersDL)
		if err = ie.MaxNumberMIMO_LayersPDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberMIMO_LayersPDSCH", err)
		}
	}
	if SupportedModulationOrderDLPresent {
		ie.SupportedModulationOrderDL = new(ModulationOrder)
		if err = ie.SupportedModulationOrderDL.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedModulationOrderDL", err)
		}
	}
	return nil
}
