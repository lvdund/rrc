package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17 struct {
	CellBarredRedCap1Rx_r17 RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17_cellBarredRedCap1Rx_r17 `madatory`
	CellBarredRedCap2Rx_r17 RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17_cellBarredRedCap2Rx_r17 `madatory`
}

func (ie *RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CellBarredRedCap1Rx_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CellBarredRedCap1Rx_r17", err)
	}
	if err = ie.CellBarredRedCap2Rx_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CellBarredRedCap2Rx_r17", err)
	}
	return nil
}

func (ie *RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CellBarredRedCap1Rx_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CellBarredRedCap1Rx_r17", err)
	}
	if err = ie.CellBarredRedCap2Rx_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CellBarredRedCap2Rx_r17", err)
	}
	return nil
}
