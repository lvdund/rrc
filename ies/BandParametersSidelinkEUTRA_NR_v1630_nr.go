package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_v1630_nr struct {
	Tx_Sidelink_r16               *BandParametersSidelinkEUTRA_NR_v1630_nr_tx_Sidelink_r16               `optional`
	Rx_Sidelink_r16               *BandParametersSidelinkEUTRA_NR_v1630_nr_rx_Sidelink_r16               `optional`
	Sl_CrossCarrierScheduling_r16 *BandParametersSidelinkEUTRA_NR_v1630_nr_sl_CrossCarrierScheduling_r16 `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630_nr) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Tx_Sidelink_r16 != nil, ie.Rx_Sidelink_r16 != nil, ie.Sl_CrossCarrierScheduling_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tx_Sidelink_r16 != nil {
		if err = ie.Tx_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_Sidelink_r16", err)
		}
	}
	if ie.Rx_Sidelink_r16 != nil {
		if err = ie.Rx_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rx_Sidelink_r16", err)
		}
	}
	if ie.Sl_CrossCarrierScheduling_r16 != nil {
		if err = ie.Sl_CrossCarrierScheduling_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CrossCarrierScheduling_r16", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630_nr) Decode(r *aper.AperReader) error {
	var err error
	var Tx_Sidelink_r16Present bool
	if Tx_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rx_Sidelink_r16Present bool
	if Rx_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CrossCarrierScheduling_r16Present bool
	if Sl_CrossCarrierScheduling_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Tx_Sidelink_r16Present {
		ie.Tx_Sidelink_r16 = new(BandParametersSidelinkEUTRA_NR_v1630_nr_tx_Sidelink_r16)
		if err = ie.Tx_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_Sidelink_r16", err)
		}
	}
	if Rx_Sidelink_r16Present {
		ie.Rx_Sidelink_r16 = new(BandParametersSidelinkEUTRA_NR_v1630_nr_rx_Sidelink_r16)
		if err = ie.Rx_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rx_Sidelink_r16", err)
		}
	}
	if Sl_CrossCarrierScheduling_r16Present {
		ie.Sl_CrossCarrierScheduling_r16 = new(BandParametersSidelinkEUTRA_NR_v1630_nr_sl_CrossCarrierScheduling_r16)
		if err = ie.Sl_CrossCarrierScheduling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CrossCarrierScheduling_r16", err)
		}
	}
	return nil
}
