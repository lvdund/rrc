package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_v1710_nr struct {
	Sl_TransmissionMode2_PartialSensing_r17 *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17 `optional`
	Rx_sidelinkPSFCH_r17                    *BandParametersSidelinkEUTRA_NR_v1710_nr_rx_sidelinkPSFCH_r17                    `optional`
	Tx_IUC_Scheme1_Mode2Sidelink_r17        *BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17        `optional`
	Tx_IUC_Scheme2_Mode2Sidelink_r17        *BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme2_Mode2Sidelink_r17        `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TransmissionMode2_PartialSensing_r17 != nil, ie.Rx_sidelinkPSFCH_r17 != nil, ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil, ie.Tx_IUC_Scheme2_Mode2Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TransmissionMode2_PartialSensing_r17 != nil {
		if err = ie.Sl_TransmissionMode2_PartialSensing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TransmissionMode2_PartialSensing_r17", err)
		}
	}
	if ie.Rx_sidelinkPSFCH_r17 != nil {
		if err = ie.Rx_sidelinkPSFCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rx_sidelinkPSFCH_r17", err)
		}
	}
	if ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil {
		if err = ie.Tx_IUC_Scheme1_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if ie.Tx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
		if err = ie.Tx_IUC_Scheme2_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr) Decode(r *uper.UperReader) error {
	var err error
	var Sl_TransmissionMode2_PartialSensing_r17Present bool
	if Sl_TransmissionMode2_PartialSensing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rx_sidelinkPSFCH_r17Present bool
	if Rx_sidelinkPSFCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tx_IUC_Scheme1_Mode2Sidelink_r17Present bool
	if Tx_IUC_Scheme1_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tx_IUC_Scheme2_Mode2Sidelink_r17Present bool
	if Tx_IUC_Scheme2_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TransmissionMode2_PartialSensing_r17Present {
		ie.Sl_TransmissionMode2_PartialSensing_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17)
		if err = ie.Sl_TransmissionMode2_PartialSensing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TransmissionMode2_PartialSensing_r17", err)
		}
	}
	if Rx_sidelinkPSFCH_r17Present {
		ie.Rx_sidelinkPSFCH_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_rx_sidelinkPSFCH_r17)
		if err = ie.Rx_sidelinkPSFCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rx_sidelinkPSFCH_r17", err)
		}
	}
	if Tx_IUC_Scheme1_Mode2Sidelink_r17Present {
		ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17)
		if err = ie.Tx_IUC_Scheme1_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if Tx_IUC_Scheme2_Mode2Sidelink_r17Present {
		ie.Tx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme2_Mode2Sidelink_r17)
		if err = ie.Tx_IUC_Scheme2_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}
