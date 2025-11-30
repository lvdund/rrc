package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1530_IEs struct {
	MasterCellGroup                    *[]byte                                  `optional`
	FullConfig                         *RRCReconfiguration_v1530_IEs_fullConfig `optional`
	DedicatedNAS_MessageList           []DedicatedNAS_Message                   `lb:1,ub:maxDRB,optional`
	MasterKeyUpdate                    *MasterKeyUpdate                         `optional`
	DedicatedSIB1_Delivery             *[]byte                                  `optional`
	DedicatedSystemInformationDelivery *[]byte                                  `optional`
	OtherConfig                        *OtherConfig                             `optional`
	NonCriticalExtension               *RRCReconfiguration_v1540_IEs            `optional`
}

func (ie *RRCReconfiguration_v1530_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MasterCellGroup != nil, ie.FullConfig != nil, len(ie.DedicatedNAS_MessageList) > 0, ie.MasterKeyUpdate != nil, ie.DedicatedSIB1_Delivery != nil, ie.DedicatedSystemInformationDelivery != nil, ie.OtherConfig != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MasterCellGroup != nil {
		if err = w.WriteOctetString(*ie.MasterCellGroup, nil, false); err != nil {
			return utils.WrapError("Encode MasterCellGroup", err)
		}
	}
	if ie.FullConfig != nil {
		if err = ie.FullConfig.Encode(w); err != nil {
			return utils.WrapError("Encode FullConfig", err)
		}
	}
	if len(ie.DedicatedNAS_MessageList) > 0 {
		tmp_DedicatedNAS_MessageList := utils.NewSequence[*DedicatedNAS_Message]([]*DedicatedNAS_Message{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
		for _, i := range ie.DedicatedNAS_MessageList {
			tmp_DedicatedNAS_MessageList.Value = append(tmp_DedicatedNAS_MessageList.Value, &i)
		}
		if err = tmp_DedicatedNAS_MessageList.Encode(w); err != nil {
			return utils.WrapError("Encode DedicatedNAS_MessageList", err)
		}
	}
	if ie.MasterKeyUpdate != nil {
		if err = ie.MasterKeyUpdate.Encode(w); err != nil {
			return utils.WrapError("Encode MasterKeyUpdate", err)
		}
	}
	if ie.DedicatedSIB1_Delivery != nil {
		if err = w.WriteOctetString(*ie.DedicatedSIB1_Delivery, nil, false); err != nil {
			return utils.WrapError("Encode DedicatedSIB1_Delivery", err)
		}
	}
	if ie.DedicatedSystemInformationDelivery != nil {
		if err = w.WriteOctetString(*ie.DedicatedSystemInformationDelivery, nil, false); err != nil {
			return utils.WrapError("Encode DedicatedSystemInformationDelivery", err)
		}
	}
	if ie.OtherConfig != nil {
		if err = ie.OtherConfig.Encode(w); err != nil {
			return utils.WrapError("Encode OtherConfig", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1530_IEs) Decode(r *aper.AperReader) error {
	var err error
	var MasterCellGroupPresent bool
	if MasterCellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FullConfigPresent bool
	if FullConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DedicatedNAS_MessageListPresent bool
	if DedicatedNAS_MessageListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MasterKeyUpdatePresent bool
	if MasterKeyUpdatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DedicatedSIB1_DeliveryPresent bool
	if DedicatedSIB1_DeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DedicatedSystemInformationDeliveryPresent bool
	if DedicatedSystemInformationDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OtherConfigPresent bool
	if OtherConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MasterCellGroupPresent {
		var tmp_os_MasterCellGroup []byte
		if tmp_os_MasterCellGroup, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode MasterCellGroup", err)
		}
		ie.MasterCellGroup = &tmp_os_MasterCellGroup
	}
	if FullConfigPresent {
		ie.FullConfig = new(RRCReconfiguration_v1530_IEs_fullConfig)
		if err = ie.FullConfig.Decode(r); err != nil {
			return utils.WrapError("Decode FullConfig", err)
		}
	}
	if DedicatedNAS_MessageListPresent {
		tmp_DedicatedNAS_MessageList := utils.NewSequence[*DedicatedNAS_Message]([]*DedicatedNAS_Message{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
		fn_DedicatedNAS_MessageList := func() *DedicatedNAS_Message {
			return new(DedicatedNAS_Message)
		}
		if err = tmp_DedicatedNAS_MessageList.Decode(r, fn_DedicatedNAS_MessageList); err != nil {
			return utils.WrapError("Decode DedicatedNAS_MessageList", err)
		}
		ie.DedicatedNAS_MessageList = []DedicatedNAS_Message{}
		for _, i := range tmp_DedicatedNAS_MessageList.Value {
			ie.DedicatedNAS_MessageList = append(ie.DedicatedNAS_MessageList, *i)
		}
	}
	if MasterKeyUpdatePresent {
		ie.MasterKeyUpdate = new(MasterKeyUpdate)
		if err = ie.MasterKeyUpdate.Decode(r); err != nil {
			return utils.WrapError("Decode MasterKeyUpdate", err)
		}
	}
	if DedicatedSIB1_DeliveryPresent {
		var tmp_os_DedicatedSIB1_Delivery []byte
		if tmp_os_DedicatedSIB1_Delivery, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode DedicatedSIB1_Delivery", err)
		}
		ie.DedicatedSIB1_Delivery = &tmp_os_DedicatedSIB1_Delivery
	}
	if DedicatedSystemInformationDeliveryPresent {
		var tmp_os_DedicatedSystemInformationDelivery []byte
		if tmp_os_DedicatedSystemInformationDelivery, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode DedicatedSystemInformationDelivery", err)
		}
		ie.DedicatedSystemInformationDelivery = &tmp_os_DedicatedSystemInformationDelivery
	}
	if OtherConfigPresent {
		ie.OtherConfig = new(OtherConfig)
		if err = ie.OtherConfig.Decode(r); err != nil {
			return utils.WrapError("Decode OtherConfig", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfiguration_v1540_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
