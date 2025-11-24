package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBSBroadcastConfiguration_r17_IEs struct {
	Mbs_SessionInfoList_r17        *MBS_SessionInfoList_r17        `optional`
	Mbs_NeighbourCellList_r17      *MBS_NeighbourCellList_r17      `optional`
	Drx_ConfigPTM_List_r17         []DRX_ConfigPTM_r17             `lb:1,ub:maxNrofDRX_ConfigPTM_r17,optional`
	Pdsch_ConfigMTCH_r17           *PDSCH_ConfigBroadcast_r17      `optional`
	Mtch_SSB_MappingWindowList_r17 *MTCH_SSB_MappingWindowList_r17 `optional`
	LateNonCriticalExtension       *[]byte                         `optional`
	NonCriticalExtension           interface{}                     `optional`
}

func (ie *MBSBroadcastConfiguration_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mbs_SessionInfoList_r17 != nil, ie.Mbs_NeighbourCellList_r17 != nil, len(ie.Drx_ConfigPTM_List_r17) > 0, ie.Pdsch_ConfigMTCH_r17 != nil, ie.Mtch_SSB_MappingWindowList_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mbs_SessionInfoList_r17 != nil {
		if err = ie.Mbs_SessionInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_SessionInfoList_r17", err)
		}
	}
	if ie.Mbs_NeighbourCellList_r17 != nil {
		if err = ie.Mbs_NeighbourCellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_NeighbourCellList_r17", err)
		}
	}
	if len(ie.Drx_ConfigPTM_List_r17) > 0 {
		tmp_Drx_ConfigPTM_List_r17 := utils.NewSequence[*DRX_ConfigPTM_r17]([]*DRX_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDRX_ConfigPTM_r17}, false)
		for _, i := range ie.Drx_ConfigPTM_List_r17 {
			tmp_Drx_ConfigPTM_List_r17.Value = append(tmp_Drx_ConfigPTM_List_r17.Value, &i)
		}
		if err = tmp_Drx_ConfigPTM_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_ConfigPTM_List_r17", err)
		}
	}
	if ie.Pdsch_ConfigMTCH_r17 != nil {
		if err = ie.Pdsch_ConfigMTCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ConfigMTCH_r17", err)
		}
	}
	if ie.Mtch_SSB_MappingWindowList_r17 != nil {
		if err = ie.Mtch_SSB_MappingWindowList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mtch_SSB_MappingWindowList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MBSBroadcastConfiguration_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Mbs_SessionInfoList_r17Present bool
	if Mbs_SessionInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mbs_NeighbourCellList_r17Present bool
	if Mbs_NeighbourCellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_ConfigPTM_List_r17Present bool
	if Drx_ConfigPTM_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ConfigMTCH_r17Present bool
	if Pdsch_ConfigMTCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mtch_SSB_MappingWindowList_r17Present bool
	if Mtch_SSB_MappingWindowList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mbs_SessionInfoList_r17Present {
		ie.Mbs_SessionInfoList_r17 = new(MBS_SessionInfoList_r17)
		if err = ie.Mbs_SessionInfoList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_SessionInfoList_r17", err)
		}
	}
	if Mbs_NeighbourCellList_r17Present {
		ie.Mbs_NeighbourCellList_r17 = new(MBS_NeighbourCellList_r17)
		if err = ie.Mbs_NeighbourCellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_NeighbourCellList_r17", err)
		}
	}
	if Drx_ConfigPTM_List_r17Present {
		tmp_Drx_ConfigPTM_List_r17 := utils.NewSequence[*DRX_ConfigPTM_r17]([]*DRX_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDRX_ConfigPTM_r17}, false)
		fn_Drx_ConfigPTM_List_r17 := func() *DRX_ConfigPTM_r17 {
			return new(DRX_ConfigPTM_r17)
		}
		if err = tmp_Drx_ConfigPTM_List_r17.Decode(r, fn_Drx_ConfigPTM_List_r17); err != nil {
			return utils.WrapError("Decode Drx_ConfigPTM_List_r17", err)
		}
		ie.Drx_ConfigPTM_List_r17 = []DRX_ConfigPTM_r17{}
		for _, i := range tmp_Drx_ConfigPTM_List_r17.Value {
			ie.Drx_ConfigPTM_List_r17 = append(ie.Drx_ConfigPTM_List_r17, *i)
		}
	}
	if Pdsch_ConfigMTCH_r17Present {
		ie.Pdsch_ConfigMTCH_r17 = new(PDSCH_ConfigBroadcast_r17)
		if err = ie.Pdsch_ConfigMTCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ConfigMTCH_r17", err)
		}
	}
	if Mtch_SSB_MappingWindowList_r17Present {
		ie.Mtch_SSB_MappingWindowList_r17 = new(MTCH_SSB_MappingWindowList_r17)
		if err = ie.Mtch_SSB_MappingWindowList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mtch_SSB_MappingWindowList_r17", err)
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
