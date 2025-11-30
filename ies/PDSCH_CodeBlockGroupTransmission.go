package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_CodeBlockGroupTransmission struct {
	MaxCodeBlockGroupsPerTransportBlock PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock `madatory`
	CodeBlockGroupFlushIndicator        bool                                                                 `madatory`
}

func (ie *PDSCH_CodeBlockGroupTransmission) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxCodeBlockGroupsPerTransportBlock.Encode(w); err != nil {
		return utils.WrapError("Encode MaxCodeBlockGroupsPerTransportBlock", err)
	}
	if err = w.WriteBoolean(ie.CodeBlockGroupFlushIndicator); err != nil {
		return utils.WrapError("WriteBoolean CodeBlockGroupFlushIndicator", err)
	}
	return nil
}

func (ie *PDSCH_CodeBlockGroupTransmission) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxCodeBlockGroupsPerTransportBlock.Decode(r); err != nil {
		return utils.WrapError("Decode MaxCodeBlockGroupsPerTransportBlock", err)
	}
	var tmp_bool_CodeBlockGroupFlushIndicator bool
	if tmp_bool_CodeBlockGroupFlushIndicator, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean CodeBlockGroupFlushIndicator", err)
	}
	ie.CodeBlockGroupFlushIndicator = tmp_bool_CodeBlockGroupFlushIndicator
	return nil
}
