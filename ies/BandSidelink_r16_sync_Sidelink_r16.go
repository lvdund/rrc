package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sync_Sidelink_r16 struct {
	GNB_Sync_r16                              *BandSidelink_r16_sync_Sidelink_r16_gNB_Sync_r16                              `optional`
	GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 *BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 `optional`
	GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16    *BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16    `optional`
}

func (ie *BandSidelink_r16_sync_Sidelink_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.GNB_Sync_r16 != nil, ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 != nil, ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.GNB_Sync_r16 != nil {
		if err = ie.GNB_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_Sync_r16", err)
		}
	}
	if ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 != nil {
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16", err)
		}
	}
	if ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 != nil {
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sync_Sidelink_r16) Decode(r *aper.AperReader) error {
	var err error
	var GNB_Sync_r16Present bool
	if GNB_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16Present bool
	if GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16Present bool
	if GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if GNB_Sync_r16Present {
		ie.GNB_Sync_r16 = new(BandSidelink_r16_sync_Sidelink_r16_gNB_Sync_r16)
		if err = ie.GNB_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_Sync_r16", err)
		}
	}
	if GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16Present {
		ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 = new(BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16)
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16", err)
		}
	}
	if GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16Present {
		ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 = new(BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16)
		if err = ie.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16", err)
		}
	}
	return nil
}
