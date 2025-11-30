package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_srs_PortReport_r17 struct {
	CapVal1_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal1_r17 `optional`
	CapVal2_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal2_r17 `optional`
	CapVal3_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal3_r17 `optional`
	CapVal4_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal4_r17 `optional`
}

func (ie *MIMO_ParametersPerBand_srs_PortReport_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CapVal1_r17 != nil, ie.CapVal2_r17 != nil, ie.CapVal3_r17 != nil, ie.CapVal4_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CapVal1_r17 != nil {
		if err = ie.CapVal1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CapVal1_r17", err)
		}
	}
	if ie.CapVal2_r17 != nil {
		if err = ie.CapVal2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CapVal2_r17", err)
		}
	}
	if ie.CapVal3_r17 != nil {
		if err = ie.CapVal3_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CapVal3_r17", err)
		}
	}
	if ie.CapVal4_r17 != nil {
		if err = ie.CapVal4_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CapVal4_r17", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_srs_PortReport_r17) Decode(r *aper.AperReader) error {
	var err error
	var CapVal1_r17Present bool
	if CapVal1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CapVal2_r17Present bool
	if CapVal2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CapVal3_r17Present bool
	if CapVal3_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CapVal4_r17Present bool
	if CapVal4_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CapVal1_r17Present {
		ie.CapVal1_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal1_r17)
		if err = ie.CapVal1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CapVal1_r17", err)
		}
	}
	if CapVal2_r17Present {
		ie.CapVal2_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal2_r17)
		if err = ie.CapVal2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CapVal2_r17", err)
		}
	}
	if CapVal3_r17Present {
		ie.CapVal3_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal3_r17)
		if err = ie.CapVal3_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CapVal3_r17", err)
		}
	}
	if CapVal4_r17Present {
		ie.CapVal4_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal4_r17)
		if err = ie.CapVal4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CapVal4_r17", err)
		}
	}
	return nil
}
