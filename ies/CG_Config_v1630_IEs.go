package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1630_IEs struct {
	SelectedToffset_r16  *T_Offset_r16        `optional`
	NonCriticalExtension *CG_Config_v1640_IEs `optional`
}

func (ie *CG_Config_v1630_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SelectedToffset_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SelectedToffset_r16 != nil {
		if err = ie.SelectedToffset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SelectedToffset_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1630_IEs) Decode(r *aper.AperReader) error {
	var err error
	var SelectedToffset_r16Present bool
	if SelectedToffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SelectedToffset_r16Present {
		ie.SelectedToffset_r16 = new(T_Offset_r16)
		if err = ie.SelectedToffset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SelectedToffset_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1640_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
