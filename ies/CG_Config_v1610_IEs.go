package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1610_IEs struct {
	Drx_InfoSCG2         *DRX_Info2           `optional`
	NonCriticalExtension *CG_Config_v1620_IEs `optional`
}

func (ie *CG_Config_v1610_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Drx_InfoSCG2 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drx_InfoSCG2 != nil {
		if err = ie.Drx_InfoSCG2.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_InfoSCG2", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1610_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Drx_InfoSCG2Present bool
	if Drx_InfoSCG2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Drx_InfoSCG2Present {
		ie.Drx_InfoSCG2 = new(DRX_Info2)
		if err = ie.Drx_InfoSCG2.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_InfoSCG2", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1620_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
