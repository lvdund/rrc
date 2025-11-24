package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1630_IEs struct {
	Uac_BarringInfo_v1630 *SIB1_v1630_IEs_uac_BarringInfo_v1630 `lb:2,ub:maxPLMN,optional`
	NonCriticalExtension  *SIB1_v1700_IEs                       `optional`
}

func (ie *SIB1_v1630_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Uac_BarringInfo_v1630 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Uac_BarringInfo_v1630 != nil {
		if err = ie.Uac_BarringInfo_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_BarringInfo_v1630", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB1_v1630_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Uac_BarringInfo_v1630Present bool
	if Uac_BarringInfo_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Uac_BarringInfo_v1630Present {
		ie.Uac_BarringInfo_v1630 = new(SIB1_v1630_IEs_uac_BarringInfo_v1630)
		if err = ie.Uac_BarringInfo_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_BarringInfo_v1630", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(SIB1_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
