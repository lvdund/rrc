package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1640 struct {
	Csi_RSRP_AndRSRQ_MeasWithSSB_r16    *SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithSSB_r16    `optional`
	Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 *SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 `optional`
	Csi_SINR_Meas_r16                   *SharedSpectrumChAccessParamsPerBand_v1640_csi_SINR_Meas_r16                   `optional`
	Ssb_AndCSI_RS_RLM_r16               *SharedSpectrumChAccessParamsPerBand_v1640_ssb_AndCSI_RS_RLM_r16               `optional`
	Csi_RS_CFRA_ForHO_r16               *SharedSpectrumChAccessParamsPerBand_v1640_csi_RS_CFRA_ForHO_r16               `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1640) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16 != nil, ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 != nil, ie.Csi_SINR_Meas_r16 != nil, ie.Ssb_AndCSI_RS_RLM_r16 != nil, ie.Csi_RS_CFRA_ForHO_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16 != nil {
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RSRP_AndRSRQ_MeasWithSSB_r16", err)
		}
	}
	if ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 != nil {
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16", err)
		}
	}
	if ie.Csi_SINR_Meas_r16 != nil {
		if err = ie.Csi_SINR_Meas_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_SINR_Meas_r16", err)
		}
	}
	if ie.Ssb_AndCSI_RS_RLM_r16 != nil {
		if err = ie.Ssb_AndCSI_RS_RLM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_AndCSI_RS_RLM_r16", err)
		}
	}
	if ie.Csi_RS_CFRA_ForHO_r16 != nil {
		if err = ie.Csi_RS_CFRA_ForHO_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_CFRA_ForHO_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1640) Decode(r *aper.AperReader) error {
	var err error
	var Csi_RSRP_AndRSRQ_MeasWithSSB_r16Present bool
	if Csi_RSRP_AndRSRQ_MeasWithSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16Present bool
	if Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_SINR_Meas_r16Present bool
	if Csi_SINR_Meas_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_AndCSI_RS_RLM_r16Present bool
	if Ssb_AndCSI_RS_RLM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_CFRA_ForHO_r16Present bool
	if Csi_RS_CFRA_ForHO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Csi_RSRP_AndRSRQ_MeasWithSSB_r16Present {
		ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithSSB_r16)
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RSRP_AndRSRQ_MeasWithSSB_r16", err)
		}
	}
	if Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16Present {
		ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithoutSSB_r16)
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RSRP_AndRSRQ_MeasWithoutSSB_r16", err)
		}
	}
	if Csi_SINR_Meas_r16Present {
		ie.Csi_SINR_Meas_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_SINR_Meas_r16)
		if err = ie.Csi_SINR_Meas_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_SINR_Meas_r16", err)
		}
	}
	if Ssb_AndCSI_RS_RLM_r16Present {
		ie.Ssb_AndCSI_RS_RLM_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_ssb_AndCSI_RS_RLM_r16)
		if err = ie.Ssb_AndCSI_RS_RLM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_AndCSI_RS_RLM_r16", err)
		}
	}
	if Csi_RS_CFRA_ForHO_r16Present {
		ie.Csi_RS_CFRA_ForHO_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_RS_CFRA_ForHO_r16)
		if err = ie.Csi_RS_CFRA_ForHO_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_CFRA_ForHO_r16", err)
		}
	}
	return nil
}
