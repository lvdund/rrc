package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRAT_Container struct {
	Rat_Type                   RAT_Type `madatory`
	Ue_CapabilityRAT_Container []byte   `madatory`
}

func (ie *UE_CapabilityRAT_Container) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Rat_Type.Encode(w); err != nil {
		return utils.WrapError("Encode Rat_Type", err)
	}
	if err = w.WriteOctetString(ie.Ue_CapabilityRAT_Container, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString Ue_CapabilityRAT_Container", err)
	}
	return nil
}

func (ie *UE_CapabilityRAT_Container) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Rat_Type.Decode(r); err != nil {
		return utils.WrapError("Decode Rat_Type", err)
	}
	var tmp_os_Ue_CapabilityRAT_Container []byte
	if tmp_os_Ue_CapabilityRAT_Container, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString Ue_CapabilityRAT_Container", err)
	}
	ie.Ue_CapabilityRAT_Container = tmp_os_Ue_CapabilityRAT_Container
	return nil
}
