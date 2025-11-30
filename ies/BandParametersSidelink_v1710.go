package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelink_v1710 struct {
	Tx_IUC_Scheme1_Mode2Sidelink_r17 *BandParametersSidelink_v1710_tx_IUC_Scheme1_Mode2Sidelink_r17 `optional`
	Tx_IUC_Scheme2_Mode2Sidelink_r17 *BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17 `optional`
}

func (ie *BandParametersSidelink_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil, ie.Tx_IUC_Scheme2_Mode2Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
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

func (ie *BandParametersSidelink_v1710) Decode(r *aper.AperReader) error {
	var err error
	var Tx_IUC_Scheme1_Mode2Sidelink_r17Present bool
	if Tx_IUC_Scheme1_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tx_IUC_Scheme2_Mode2Sidelink_r17Present bool
	if Tx_IUC_Scheme2_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Tx_IUC_Scheme1_Mode2Sidelink_r17Present {
		ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 = new(BandParametersSidelink_v1710_tx_IUC_Scheme1_Mode2Sidelink_r17)
		if err = ie.Tx_IUC_Scheme1_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if Tx_IUC_Scheme2_Mode2Sidelink_r17Present {
		ie.Tx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17)
		if err = ie.Tx_IUC_Scheme2_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}
