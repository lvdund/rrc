package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1720 struct {
	Rtt_BasedPDC_CSI_RS_ForTracking_r17 *FeatureSetDownlink_v1720_rtt_BasedPDC_CSI_RS_ForTracking_r17 `optional`
	Rtt_BasedPDC_PRS_r17                *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17                `optional`
	Sps_Multicast_r17                   *FeatureSetDownlink_v1720_sps_Multicast_r17                   `optional`
}

func (ie *FeatureSetDownlink_v1720) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17 != nil, ie.Rtt_BasedPDC_PRS_r17 != nil, ie.Sps_Multicast_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17 != nil {
		if err = ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rtt_BasedPDC_CSI_RS_ForTracking_r17", err)
		}
	}
	if ie.Rtt_BasedPDC_PRS_r17 != nil {
		if err = ie.Rtt_BasedPDC_PRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rtt_BasedPDC_PRS_r17", err)
		}
	}
	if ie.Sps_Multicast_r17 != nil {
		if err = ie.Sps_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sps_Multicast_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720) Decode(r *aper.AperReader) error {
	var err error
	var Rtt_BasedPDC_CSI_RS_ForTracking_r17Present bool
	if Rtt_BasedPDC_CSI_RS_ForTracking_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rtt_BasedPDC_PRS_r17Present bool
	if Rtt_BasedPDC_PRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sps_Multicast_r17Present bool
	if Sps_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rtt_BasedPDC_CSI_RS_ForTracking_r17Present {
		ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_CSI_RS_ForTracking_r17)
		if err = ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rtt_BasedPDC_CSI_RS_ForTracking_r17", err)
		}
	}
	if Rtt_BasedPDC_PRS_r17Present {
		ie.Rtt_BasedPDC_PRS_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17)
		if err = ie.Rtt_BasedPDC_PRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rtt_BasedPDC_PRS_r17", err)
		}
	}
	if Sps_Multicast_r17Present {
		ie.Sps_Multicast_r17 = new(FeatureSetDownlink_v1720_sps_Multicast_r17)
		if err = ie.Sps_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sps_Multicast_r17", err)
		}
	}
	return nil
}
