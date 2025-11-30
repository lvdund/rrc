package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_Uu_RLF          aper.Enumerated = 0
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_HO              aper.Enumerated = 1
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_CellReselection aper.Enumerated = 2
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_Uu_RRC_Failure  aper.Enumerated = 3
)

type NotificationMessageSidelink_r17_IEs_indicationType_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *NotificationMessageSidelink_r17_IEs_indicationType_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode NotificationMessageSidelink_r17_IEs_indicationType_r17", err)
	}
	return nil
}

func (ie *NotificationMessageSidelink_r17_IEs_indicationType_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode NotificationMessageSidelink_r17_IEs_indicationType_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
