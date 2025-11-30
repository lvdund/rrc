package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB10_r16 struct {
	Hrnn_List_r16            *HRNN_List_r16 `optional`
	LateNonCriticalExtension *[]byte        `optional`
}

func (ie *SIB10_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Hrnn_List_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Hrnn_List_r16 != nil {
		if err = ie.Hrnn_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Hrnn_List_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB10_r16) Decode(r *aper.AperReader) error {
	var err error
	var Hrnn_List_r16Present bool
	if Hrnn_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Hrnn_List_r16Present {
		ie.Hrnn_List_r16 = new(HRNN_List_r16)
		if err = ie.Hrnn_List_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Hrnn_List_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
