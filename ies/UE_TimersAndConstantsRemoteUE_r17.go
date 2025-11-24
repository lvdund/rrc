package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_TimersAndConstantsRemoteUE_r17 struct {
	T300_RemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17_t300_RemoteUE_r17 `optional`
	T301_RemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17_t301_RemoteUE_r17 `optional`
	T319_RemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17 `optional`
}

func (ie *UE_TimersAndConstantsRemoteUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.T300_RemoteUE_r17 != nil, ie.T301_RemoteUE_r17 != nil, ie.T319_RemoteUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.T300_RemoteUE_r17 != nil {
		if err = ie.T300_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T300_RemoteUE_r17", err)
		}
	}
	if ie.T301_RemoteUE_r17 != nil {
		if err = ie.T301_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T301_RemoteUE_r17", err)
		}
	}
	if ie.T319_RemoteUE_r17 != nil {
		if err = ie.T319_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode T319_RemoteUE_r17", err)
		}
	}
	return nil
}

func (ie *UE_TimersAndConstantsRemoteUE_r17) Decode(r *uper.UperReader) error {
	var err error
	var T300_RemoteUE_r17Present bool
	if T300_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T301_RemoteUE_r17Present bool
	if T301_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T319_RemoteUE_r17Present bool
	if T319_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if T300_RemoteUE_r17Present {
		ie.T300_RemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17_t300_RemoteUE_r17)
		if err = ie.T300_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T300_RemoteUE_r17", err)
		}
	}
	if T301_RemoteUE_r17Present {
		ie.T301_RemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17_t301_RemoteUE_r17)
		if err = ie.T301_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T301_RemoteUE_r17", err)
		}
	}
	if T319_RemoteUE_r17Present {
		ie.T319_RemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17)
		if err = ie.T319_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode T319_RemoteUE_r17", err)
		}
	}
	return nil
}
