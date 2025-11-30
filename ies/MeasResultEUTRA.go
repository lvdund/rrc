package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultEUTRA struct {
	Eutra_PhysCellId PhysCellId               `madatory`
	MeasResult       MeasQuantityResultsEUTRA `madatory`
	Cgi_Info         *CGI_InfoEUTRA           `optional`
}

func (ie *MeasResultEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Cgi_Info != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Eutra_PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode Eutra_PhysCellId", err)
	}
	if err = ie.MeasResult.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResult", err)
	}
	if ie.Cgi_Info != nil {
		if err = ie.Cgi_Info.Encode(w); err != nil {
			return utils.WrapError("Encode Cgi_Info", err)
		}
	}
	return nil
}

func (ie *MeasResultEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var Cgi_InfoPresent bool
	if Cgi_InfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Eutra_PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode Eutra_PhysCellId", err)
	}
	if err = ie.MeasResult.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResult", err)
	}
	if Cgi_InfoPresent {
		ie.Cgi_Info = new(CGI_InfoEUTRA)
		if err = ie.Cgi_Info.Decode(r); err != nil {
			return utils.WrapError("Decode Cgi_Info", err)
		}
	}
	return nil
}
