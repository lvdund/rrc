package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v15g0 struct {
	Rf_Parameters_v15g0  *RF_Parameters_v15g0    `optional`
	NonCriticalExtension *UE_NR_Capability_v15j0 `optional`
}

func (ie *UE_NR_Capability_v15g0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rf_Parameters_v15g0 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rf_Parameters_v15g0 != nil {
		if err = ie.Rf_Parameters_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode Rf_Parameters_v15g0", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v15g0) Decode(r *aper.AperReader) error {
	var err error
	var Rf_Parameters_v15g0Present bool
	if Rf_Parameters_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Rf_Parameters_v15g0Present {
		ie.Rf_Parameters_v15g0 = new(RF_Parameters_v15g0)
		if err = ie.Rf_Parameters_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode Rf_Parameters_v15g0", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v15j0)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
