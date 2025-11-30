package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NTN_Parameters_r17 struct {
	InactiveStateNTN_r17               *NTN_Parameters_r17_inactiveStateNTN_r17 `optional`
	Ra_SDT_NTN_r17                     *NTN_Parameters_r17_ra_SDT_NTN_r17       `optional`
	Srb_SDT_NTN_r17                    *NTN_Parameters_r17_srb_SDT_NTN_r17      `optional`
	MeasAndMobParametersNTN_r17        *MeasAndMobParameters                    `optional`
	Mac_ParametersNTN_r17              *MAC_Parameters                          `optional`
	Phy_ParametersNTN_r17              *Phy_Parameters                          `optional`
	Fdd_Add_UE_NR_CapabilitiesNTN_r17  *UE_NR_CapabilityAddXDD_Mode             `optional`
	Fr1_Add_UE_NR_CapabilitiesNTN_r17  *UE_NR_CapabilityAddFRX_Mode             `optional`
	Ue_BasedPerfMeas_ParametersNTN_r17 *UE_BasedPerfMeas_Parameters_r16         `optional`
	Son_ParametersNTN_r17              *SON_Parameters_r16                      `optional`
}

func (ie *NTN_Parameters_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.InactiveStateNTN_r17 != nil, ie.Ra_SDT_NTN_r17 != nil, ie.Srb_SDT_NTN_r17 != nil, ie.MeasAndMobParametersNTN_r17 != nil, ie.Mac_ParametersNTN_r17 != nil, ie.Phy_ParametersNTN_r17 != nil, ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17 != nil, ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17 != nil, ie.Ue_BasedPerfMeas_ParametersNTN_r17 != nil, ie.Son_ParametersNTN_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InactiveStateNTN_r17 != nil {
		if err = ie.InactiveStateNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InactiveStateNTN_r17", err)
		}
	}
	if ie.Ra_SDT_NTN_r17 != nil {
		if err = ie.Ra_SDT_NTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_SDT_NTN_r17", err)
		}
	}
	if ie.Srb_SDT_NTN_r17 != nil {
		if err = ie.Srb_SDT_NTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srb_SDT_NTN_r17", err)
		}
	}
	if ie.MeasAndMobParametersNTN_r17 != nil {
		if err = ie.MeasAndMobParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersNTN_r17", err)
		}
	}
	if ie.Mac_ParametersNTN_r17 != nil {
		if err = ie.Mac_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersNTN_r17", err)
		}
	}
	if ie.Phy_ParametersNTN_r17 != nil {
		if err = ie.Phy_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersNTN_r17", err)
		}
	}
	if ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17 != nil {
		if err = ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17 != nil {
		if err = ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if ie.Ue_BasedPerfMeas_ParametersNTN_r17 != nil {
		if err = ie.Ue_BasedPerfMeas_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_BasedPerfMeas_ParametersNTN_r17", err)
		}
	}
	if ie.Son_ParametersNTN_r17 != nil {
		if err = ie.Son_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Son_ParametersNTN_r17", err)
		}
	}
	return nil
}

func (ie *NTN_Parameters_r17) Decode(r *aper.AperReader) error {
	var err error
	var InactiveStateNTN_r17Present bool
	if InactiveStateNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_SDT_NTN_r17Present bool
	if Ra_SDT_NTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srb_SDT_NTN_r17Present bool
	if Srb_SDT_NTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersNTN_r17Present bool
	if MeasAndMobParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_ParametersNTN_r17Present bool
	if Mac_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersNTN_r17Present bool
	if Phy_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdd_Add_UE_NR_CapabilitiesNTN_r17Present bool
	if Fdd_Add_UE_NR_CapabilitiesNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_Add_UE_NR_CapabilitiesNTN_r17Present bool
	if Fr1_Add_UE_NR_CapabilitiesNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_BasedPerfMeas_ParametersNTN_r17Present bool
	if Ue_BasedPerfMeas_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Son_ParametersNTN_r17Present bool
	if Son_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InactiveStateNTN_r17Present {
		ie.InactiveStateNTN_r17 = new(NTN_Parameters_r17_inactiveStateNTN_r17)
		if err = ie.InactiveStateNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InactiveStateNTN_r17", err)
		}
	}
	if Ra_SDT_NTN_r17Present {
		ie.Ra_SDT_NTN_r17 = new(NTN_Parameters_r17_ra_SDT_NTN_r17)
		if err = ie.Ra_SDT_NTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_SDT_NTN_r17", err)
		}
	}
	if Srb_SDT_NTN_r17Present {
		ie.Srb_SDT_NTN_r17 = new(NTN_Parameters_r17_srb_SDT_NTN_r17)
		if err = ie.Srb_SDT_NTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srb_SDT_NTN_r17", err)
		}
	}
	if MeasAndMobParametersNTN_r17Present {
		ie.MeasAndMobParametersNTN_r17 = new(MeasAndMobParameters)
		if err = ie.MeasAndMobParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersNTN_r17", err)
		}
	}
	if Mac_ParametersNTN_r17Present {
		ie.Mac_ParametersNTN_r17 = new(MAC_Parameters)
		if err = ie.Mac_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersNTN_r17", err)
		}
	}
	if Phy_ParametersNTN_r17Present {
		ie.Phy_ParametersNTN_r17 = new(Phy_Parameters)
		if err = ie.Phy_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersNTN_r17", err)
		}
	}
	if Fdd_Add_UE_NR_CapabilitiesNTN_r17Present {
		ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17 = new(UE_NR_CapabilityAddXDD_Mode)
		if err = ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if Fr1_Add_UE_NR_CapabilitiesNTN_r17Present {
		ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17 = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if Ue_BasedPerfMeas_ParametersNTN_r17Present {
		ie.Ue_BasedPerfMeas_ParametersNTN_r17 = new(UE_BasedPerfMeas_Parameters_r16)
		if err = ie.Ue_BasedPerfMeas_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_BasedPerfMeas_ParametersNTN_r17", err)
		}
	}
	if Son_ParametersNTN_r17Present {
		ie.Son_ParametersNTN_r17 = new(SON_Parameters_r16)
		if err = ie.Son_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Son_ParametersNTN_r17", err)
		}
	}
	return nil
}
