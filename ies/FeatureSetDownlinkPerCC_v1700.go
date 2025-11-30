package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1700 struct {
	SupportedMinBandwidthDL_r17            *SupportedBandwidth_v1700                                             `optional`
	BroadcastSCell_r17                     *FeatureSetDownlinkPerCC_v1700_broadcastSCell_r17                     `optional`
	MaxNumberMIMO_LayersMulticastPDSCH_r17 *FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17 `optional`
	DynamicMulticastSCell_r17              *FeatureSetDownlinkPerCC_v1700_dynamicMulticastSCell_r17              `optional`
	SupportedBandwidthDL_v1710             *SupportedBandwidth_v1700                                             `optional`
	SupportedCRS_InterfMitigation_r17      *CRS_InterfMitigation_r17                                             `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedMinBandwidthDL_r17 != nil, ie.BroadcastSCell_r17 != nil, ie.MaxNumberMIMO_LayersMulticastPDSCH_r17 != nil, ie.DynamicMulticastSCell_r17 != nil, ie.SupportedBandwidthDL_v1710 != nil, ie.SupportedCRS_InterfMitigation_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedMinBandwidthDL_r17 != nil {
		if err = ie.SupportedMinBandwidthDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedMinBandwidthDL_r17", err)
		}
	}
	if ie.BroadcastSCell_r17 != nil {
		if err = ie.BroadcastSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BroadcastSCell_r17", err)
		}
	}
	if ie.MaxNumberMIMO_LayersMulticastPDSCH_r17 != nil {
		if err = ie.MaxNumberMIMO_LayersMulticastPDSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberMIMO_LayersMulticastPDSCH_r17", err)
		}
	}
	if ie.DynamicMulticastSCell_r17 != nil {
		if err = ie.DynamicMulticastSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicMulticastSCell_r17", err)
		}
	}
	if ie.SupportedBandwidthDL_v1710 != nil {
		if err = ie.SupportedBandwidthDL_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandwidthDL_v1710", err)
		}
	}
	if ie.SupportedCRS_InterfMitigation_r17 != nil {
		if err = ie.SupportedCRS_InterfMitigation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedCRS_InterfMitigation_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1700) Decode(r *aper.AperReader) error {
	var err error
	var SupportedMinBandwidthDL_r17Present bool
	if SupportedMinBandwidthDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BroadcastSCell_r17Present bool
	if BroadcastSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberMIMO_LayersMulticastPDSCH_r17Present bool
	if MaxNumberMIMO_LayersMulticastPDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicMulticastSCell_r17Present bool
	if DynamicMulticastSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandwidthDL_v1710Present bool
	if SupportedBandwidthDL_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedCRS_InterfMitigation_r17Present bool
	if SupportedCRS_InterfMitigation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedMinBandwidthDL_r17Present {
		ie.SupportedMinBandwidthDL_r17 = new(SupportedBandwidth_v1700)
		if err = ie.SupportedMinBandwidthDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedMinBandwidthDL_r17", err)
		}
	}
	if BroadcastSCell_r17Present {
		ie.BroadcastSCell_r17 = new(FeatureSetDownlinkPerCC_v1700_broadcastSCell_r17)
		if err = ie.BroadcastSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BroadcastSCell_r17", err)
		}
	}
	if MaxNumberMIMO_LayersMulticastPDSCH_r17Present {
		ie.MaxNumberMIMO_LayersMulticastPDSCH_r17 = new(FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17)
		if err = ie.MaxNumberMIMO_LayersMulticastPDSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberMIMO_LayersMulticastPDSCH_r17", err)
		}
	}
	if DynamicMulticastSCell_r17Present {
		ie.DynamicMulticastSCell_r17 = new(FeatureSetDownlinkPerCC_v1700_dynamicMulticastSCell_r17)
		if err = ie.DynamicMulticastSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicMulticastSCell_r17", err)
		}
	}
	if SupportedBandwidthDL_v1710Present {
		ie.SupportedBandwidthDL_v1710 = new(SupportedBandwidth_v1700)
		if err = ie.SupportedBandwidthDL_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandwidthDL_v1710", err)
		}
	}
	if SupportedCRS_InterfMitigation_r17Present {
		ie.SupportedCRS_InterfMitigation_r17 = new(CRS_InterfMitigation_r17)
		if err = ie.SupportedCRS_InterfMitigation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedCRS_InterfMitigation_r17", err)
		}
	}
	return nil
}
