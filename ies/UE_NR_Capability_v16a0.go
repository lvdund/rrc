package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v16a0 struct {
	Phy_Parameters_v16a0 *Phy_Parameters_v16a0 `optional`
	Rf_Parameters_v16a0  *RF_Parameters_v16a0  `optional`
	NonCriticalExtension interface{}           `optional`
}

func (ie *UE_NR_Capability_v16a0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Phy_Parameters_v16a0 != nil, ie.Rf_Parameters_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Phy_Parameters_v16a0 != nil {
		if err = ie.Phy_Parameters_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_Parameters_v16a0", err)
		}
	}
	if ie.Rf_Parameters_v16a0 != nil {
		if err = ie.Rf_Parameters_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode Rf_Parameters_v16a0", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v16a0) Decode(r *aper.AperReader) error {
	var err error
	var Phy_Parameters_v16a0Present bool
	if Phy_Parameters_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rf_Parameters_v16a0Present bool
	if Rf_Parameters_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Phy_Parameters_v16a0Present {
		ie.Phy_Parameters_v16a0 = new(Phy_Parameters_v16a0)
		if err = ie.Phy_Parameters_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_Parameters_v16a0", err)
		}
	}
	if Rf_Parameters_v16a0Present {
		ie.Rf_Parameters_v16a0 = new(RF_Parameters_v16a0)
		if err = ie.Rf_Parameters_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode Rf_Parameters_v16a0", err)
		}
	}
	return nil
}
