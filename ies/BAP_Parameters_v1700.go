package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BAP_Parameters_v1700 struct {
	BapHeaderRewriting_Rerouting_r17 *BAP_Parameters_v1700_bapHeaderRewriting_Rerouting_r17 `optional`
	BapHeaderRewriting_Routing_r17   *BAP_Parameters_v1700_bapHeaderRewriting_Routing_r17   `optional`
}

func (ie *BAP_Parameters_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BapHeaderRewriting_Rerouting_r17 != nil, ie.BapHeaderRewriting_Routing_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BapHeaderRewriting_Rerouting_r17 != nil {
		if err = ie.BapHeaderRewriting_Rerouting_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BapHeaderRewriting_Rerouting_r17", err)
		}
	}
	if ie.BapHeaderRewriting_Routing_r17 != nil {
		if err = ie.BapHeaderRewriting_Routing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BapHeaderRewriting_Routing_r17", err)
		}
	}
	return nil
}

func (ie *BAP_Parameters_v1700) Decode(r *aper.AperReader) error {
	var err error
	var BapHeaderRewriting_Rerouting_r17Present bool
	if BapHeaderRewriting_Rerouting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BapHeaderRewriting_Routing_r17Present bool
	if BapHeaderRewriting_Routing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BapHeaderRewriting_Rerouting_r17Present {
		ie.BapHeaderRewriting_Rerouting_r17 = new(BAP_Parameters_v1700_bapHeaderRewriting_Rerouting_r17)
		if err = ie.BapHeaderRewriting_Rerouting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BapHeaderRewriting_Rerouting_r17", err)
		}
	}
	if BapHeaderRewriting_Routing_r17Present {
		ie.BapHeaderRewriting_Routing_r17 = new(BAP_Parameters_v1700_bapHeaderRewriting_Routing_r17)
		if err = ie.BapHeaderRewriting_Routing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BapHeaderRewriting_Routing_r17", err)
		}
	}
	return nil
}
