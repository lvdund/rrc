package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sync_Sidelink_v1710 struct {
	Sync_GNSS_r17                             *BandSidelink_r16_sync_Sidelink_v1710_sync_GNSS_r17                             `optional`
	GNB_Sync_r17                              *BandSidelink_r16_sync_Sidelink_v1710_gNB_Sync_r17                              `optional`
	GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 *BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 `optional`
	GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17    *BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17    `optional`
}

func (ie *BandSidelink_r16_sync_Sidelink_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sync_GNSS_r17 != nil, ie.GNB_Sync_r17 != nil, ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 != nil, ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sync_GNSS_r17 != nil {
		if err = ie.Sync_GNSS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sync_GNSS_r17", err)
		}
	}
	if ie.GNB_Sync_r17 != nil {
		if err = ie.GNB_Sync_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_Sync_r17", err)
		}
	}
	if ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 != nil {
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17", err)
		}
	}
	if ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 != nil {
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sync_Sidelink_v1710) Decode(r *uper.UperReader) error {
	var err error
	var Sync_GNSS_r17Present bool
	if Sync_GNSS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GNB_Sync_r17Present bool
	if GNB_Sync_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17Present bool
	if GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17Present bool
	if GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sync_GNSS_r17Present {
		ie.Sync_GNSS_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_sync_GNSS_r17)
		if err = ie.Sync_GNSS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sync_GNSS_r17", err)
		}
	}
	if GNB_Sync_r17Present {
		ie.GNB_Sync_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_gNB_Sync_r17)
		if err = ie.GNB_Sync_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_Sync_r17", err)
		}
	}
	if GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17Present {
		ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17)
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17", err)
		}
	}
	if GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17Present {
		ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17)
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17", err)
		}
	}
	return nil
}
