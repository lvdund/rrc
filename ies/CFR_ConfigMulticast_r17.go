package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CFR_ConfigMulticast_r17 struct {
	LocationAndBandwidthMulticast_r17    *int64                                `lb:0,ub:37949,optional`
	Pdcch_ConfigMulticast_r17            *PDCCH_Config                         `optional`
	Pdsch_ConfigMulticast_r17            *PDSCH_Config                         `optional`
	Sps_ConfigMulticastToAddModList_r17  *SPS_ConfigMulticastToAddModList_r17  `optional`
	Sps_ConfigMulticastToReleaseList_r17 *SPS_ConfigMulticastToReleaseList_r17 `optional`
}

func (ie *CFR_ConfigMulticast_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LocationAndBandwidthMulticast_r17 != nil, ie.Pdcch_ConfigMulticast_r17 != nil, ie.Pdsch_ConfigMulticast_r17 != nil, ie.Sps_ConfigMulticastToAddModList_r17 != nil, ie.Sps_ConfigMulticastToReleaseList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.LocationAndBandwidthMulticast_r17 != nil {
		if err = w.WriteInteger(*ie.LocationAndBandwidthMulticast_r17, &aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			return utils.WrapError("Encode LocationAndBandwidthMulticast_r17", err)
		}
	}
	if ie.Pdcch_ConfigMulticast_r17 != nil {
		if err = ie.Pdcch_ConfigMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_ConfigMulticast_r17", err)
		}
	}
	if ie.Pdsch_ConfigMulticast_r17 != nil {
		if err = ie.Pdsch_ConfigMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ConfigMulticast_r17", err)
		}
	}
	if ie.Sps_ConfigMulticastToAddModList_r17 != nil {
		if err = ie.Sps_ConfigMulticastToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sps_ConfigMulticastToAddModList_r17", err)
		}
	}
	if ie.Sps_ConfigMulticastToReleaseList_r17 != nil {
		if err = ie.Sps_ConfigMulticastToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sps_ConfigMulticastToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *CFR_ConfigMulticast_r17) Decode(r *aper.AperReader) error {
	var err error
	var LocationAndBandwidthMulticast_r17Present bool
	if LocationAndBandwidthMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_ConfigMulticast_r17Present bool
	if Pdcch_ConfigMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ConfigMulticast_r17Present bool
	if Pdsch_ConfigMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sps_ConfigMulticastToAddModList_r17Present bool
	if Sps_ConfigMulticastToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sps_ConfigMulticastToReleaseList_r17Present bool
	if Sps_ConfigMulticastToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if LocationAndBandwidthMulticast_r17Present {
		var tmp_int_LocationAndBandwidthMulticast_r17 int64
		if tmp_int_LocationAndBandwidthMulticast_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			return utils.WrapError("Decode LocationAndBandwidthMulticast_r17", err)
		}
		ie.LocationAndBandwidthMulticast_r17 = &tmp_int_LocationAndBandwidthMulticast_r17
	}
	if Pdcch_ConfigMulticast_r17Present {
		ie.Pdcch_ConfigMulticast_r17 = new(PDCCH_Config)
		if err = ie.Pdcch_ConfigMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_ConfigMulticast_r17", err)
		}
	}
	if Pdsch_ConfigMulticast_r17Present {
		ie.Pdsch_ConfigMulticast_r17 = new(PDSCH_Config)
		if err = ie.Pdsch_ConfigMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ConfigMulticast_r17", err)
		}
	}
	if Sps_ConfigMulticastToAddModList_r17Present {
		ie.Sps_ConfigMulticastToAddModList_r17 = new(SPS_ConfigMulticastToAddModList_r17)
		if err = ie.Sps_ConfigMulticastToAddModList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sps_ConfigMulticastToAddModList_r17", err)
		}
	}
	if Sps_ConfigMulticastToReleaseList_r17Present {
		ie.Sps_ConfigMulticastToReleaseList_r17 = new(SPS_ConfigMulticastToReleaseList_r17)
		if err = ie.Sps_ConfigMulticastToReleaseList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sps_ConfigMulticastToReleaseList_r17", err)
		}
	}
	return nil
}
