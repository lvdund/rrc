package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17 struct {
	MinBeamApplicationTime_r17 *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_minBeamApplicationTime_r17 `optional`
	MaxNumMAC_CE_PerCC         MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC          `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MinBeamApplicationTime_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MinBeamApplicationTime_r17 != nil {
		if err = ie.MinBeamApplicationTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MinBeamApplicationTime_r17", err)
		}
	}
	if err = ie.MaxNumMAC_CE_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumMAC_CE_PerCC", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17) Decode(r *aper.AperReader) error {
	var err error
	var MinBeamApplicationTime_r17Present bool
	if MinBeamApplicationTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MinBeamApplicationTime_r17Present {
		ie.MinBeamApplicationTime_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_minBeamApplicationTime_r17)
		if err = ie.MinBeamApplicationTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MinBeamApplicationTime_r17", err)
		}
	}
	if err = ie.MaxNumMAC_CE_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumMAC_CE_PerCC", err)
	}
	return nil
}
