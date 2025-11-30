package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MasterKeyUpdate struct {
	KeySetChangeIndicator bool                 `madatory`
	NextHopChainingCount  NextHopChainingCount `madatory`
	Nas_Container         *[]byte              `optional`
}

func (ie *MasterKeyUpdate) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Nas_Container != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBoolean(ie.KeySetChangeIndicator); err != nil {
		return utils.WrapError("WriteBoolean KeySetChangeIndicator", err)
	}
	if err = ie.NextHopChainingCount.Encode(w); err != nil {
		return utils.WrapError("Encode NextHopChainingCount", err)
	}
	if ie.Nas_Container != nil {
		if err = w.WriteOctetString(*ie.Nas_Container, nil, false); err != nil {
			return utils.WrapError("Encode Nas_Container", err)
		}
	}
	return nil
}

func (ie *MasterKeyUpdate) Decode(r *aper.AperReader) error {
	var err error
	var Nas_ContainerPresent bool
	if Nas_ContainerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bool_KeySetChangeIndicator bool
	if tmp_bool_KeySetChangeIndicator, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean KeySetChangeIndicator", err)
	}
	ie.KeySetChangeIndicator = tmp_bool_KeySetChangeIndicator
	if err = ie.NextHopChainingCount.Decode(r); err != nil {
		return utils.WrapError("Decode NextHopChainingCount", err)
	}
	if Nas_ContainerPresent {
		var tmp_os_Nas_Container []byte
		if tmp_os_Nas_Container, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Nas_Container", err)
		}
		ie.Nas_Container = &tmp_os_Nas_Container
	}
	return nil
}
