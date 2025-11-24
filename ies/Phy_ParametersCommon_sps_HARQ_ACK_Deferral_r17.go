package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17 struct {
	Non_SharedSpectrumChAccess_r17 *Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17_non_SharedSpectrumChAccess_r17 `optional`
	SharedSpectrumChAccess_r17     *Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17_sharedSpectrumChAccess_r17     `optional`
}

func (ie *Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Non_SharedSpectrumChAccess_r17 != nil, ie.SharedSpectrumChAccess_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Non_SharedSpectrumChAccess_r17 != nil {
		if err = ie.Non_SharedSpectrumChAccess_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Non_SharedSpectrumChAccess_r17", err)
		}
	}
	if ie.SharedSpectrumChAccess_r17 != nil {
		if err = ie.SharedSpectrumChAccess_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SharedSpectrumChAccess_r17", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17) Decode(r *uper.UperReader) error {
	var err error
	var Non_SharedSpectrumChAccess_r17Present bool
	if Non_SharedSpectrumChAccess_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SharedSpectrumChAccess_r17Present bool
	if SharedSpectrumChAccess_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Non_SharedSpectrumChAccess_r17Present {
		ie.Non_SharedSpectrumChAccess_r17 = new(Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17_non_SharedSpectrumChAccess_r17)
		if err = ie.Non_SharedSpectrumChAccess_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Non_SharedSpectrumChAccess_r17", err)
		}
	}
	if SharedSpectrumChAccess_r17Present {
		ie.SharedSpectrumChAccess_r17 = new(Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17_sharedSpectrumChAccess_r17)
		if err = ie.SharedSpectrumChAccess_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SharedSpectrumChAccess_r17", err)
		}
	}
	return nil
}
