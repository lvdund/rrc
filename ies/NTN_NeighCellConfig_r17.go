package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NTN_NeighCellConfig_r17 struct {
	Ntn_Config_r17  *NTN_Config_r17 `optional`
	CarrierFreq_r17 *ARFCN_ValueNR  `optional`
	PhysCellId_r17  *PhysCellId     `optional`
}

func (ie *NTN_NeighCellConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ntn_Config_r17 != nil, ie.CarrierFreq_r17 != nil, ie.PhysCellId_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ntn_Config_r17 != nil {
		if err = ie.Ntn_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_Config_r17", err)
		}
	}
	if ie.CarrierFreq_r17 != nil {
		if err = ie.CarrierFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CarrierFreq_r17", err)
		}
	}
	if ie.PhysCellId_r17 != nil {
		if err = ie.PhysCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PhysCellId_r17", err)
		}
	}
	return nil
}

func (ie *NTN_NeighCellConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Ntn_Config_r17Present bool
	if Ntn_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CarrierFreq_r17Present bool
	if CarrierFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PhysCellId_r17Present bool
	if PhysCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ntn_Config_r17Present {
		ie.Ntn_Config_r17 = new(NTN_Config_r17)
		if err = ie.Ntn_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_Config_r17", err)
		}
	}
	if CarrierFreq_r17Present {
		ie.CarrierFreq_r17 = new(ARFCN_ValueNR)
		if err = ie.CarrierFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CarrierFreq_r17", err)
		}
	}
	if PhysCellId_r17Present {
		ie.PhysCellId_r17 = new(PhysCellId)
		if err = ie.PhysCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PhysCellId_r17", err)
		}
	}
	return nil
}
