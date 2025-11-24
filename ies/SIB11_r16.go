package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB11_r16 struct {
	MeasIdleConfigSIB_r16    *MeasIdleConfigSIB_r16 `optional`
	LateNonCriticalExtension *[]byte                `optional`
}

func (ie *SIB11_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasIdleConfigSIB_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasIdleConfigSIB_r16 != nil {
		if err = ie.MeasIdleConfigSIB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdleConfigSIB_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB11_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasIdleConfigSIB_r16Present bool
	if MeasIdleConfigSIB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasIdleConfigSIB_r16Present {
		ie.MeasIdleConfigSIB_r16 = new(MeasIdleConfigSIB_r16)
		if err = ie.MeasIdleConfigSIB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdleConfigSIB_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
