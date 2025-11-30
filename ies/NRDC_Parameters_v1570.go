package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v1570 struct {
	Sfn_SyncNRDC *NRDC_Parameters_v1570_sfn_SyncNRDC `optional`
}

func (ie *NRDC_Parameters_v1570) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sfn_SyncNRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sfn_SyncNRDC != nil {
		if err = ie.Sfn_SyncNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SyncNRDC", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters_v1570) Decode(r *aper.AperReader) error {
	var err error
	var Sfn_SyncNRDCPresent bool
	if Sfn_SyncNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sfn_SyncNRDCPresent {
		ie.Sfn_SyncNRDC = new(NRDC_Parameters_v1570_sfn_SyncNRDC)
		if err = ie.Sfn_SyncNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SyncNRDC", err)
		}
	}
	return nil
}
