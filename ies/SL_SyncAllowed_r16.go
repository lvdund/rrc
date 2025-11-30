package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_SyncAllowed_r16 struct {
	Gnss_Sync_r16   *SL_SyncAllowed_r16_gnss_Sync_r16   `optional`
	GnbEnb_Sync_r16 *SL_SyncAllowed_r16_gnbEnb_Sync_r16 `optional`
	Ue_Sync_r16     *SL_SyncAllowed_r16_ue_Sync_r16     `optional`
}

func (ie *SL_SyncAllowed_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Gnss_Sync_r16 != nil, ie.GnbEnb_Sync_r16 != nil, ie.Ue_Sync_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Gnss_Sync_r16 != nil {
		if err = ie.Gnss_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnss_Sync_r16", err)
		}
	}
	if ie.GnbEnb_Sync_r16 != nil {
		if err = ie.GnbEnb_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode GnbEnb_Sync_r16", err)
		}
	}
	if ie.Ue_Sync_r16 != nil {
		if err = ie.Ue_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_Sync_r16", err)
		}
	}
	return nil
}

func (ie *SL_SyncAllowed_r16) Decode(r *aper.AperReader) error {
	var err error
	var Gnss_Sync_r16Present bool
	if Gnss_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GnbEnb_Sync_r16Present bool
	if GnbEnb_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_Sync_r16Present bool
	if Ue_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Gnss_Sync_r16Present {
		ie.Gnss_Sync_r16 = new(SL_SyncAllowed_r16_gnss_Sync_r16)
		if err = ie.Gnss_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnss_Sync_r16", err)
		}
	}
	if GnbEnb_Sync_r16Present {
		ie.GnbEnb_Sync_r16 = new(SL_SyncAllowed_r16_gnbEnb_Sync_r16)
		if err = ie.GnbEnb_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode GnbEnb_Sync_r16", err)
		}
	}
	if Ue_Sync_r16Present {
		ie.Ue_Sync_r16 = new(SL_SyncAllowed_r16_ue_Sync_r16)
		if err = ie.Ue_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_Sync_r16", err)
		}
	}
	return nil
}
