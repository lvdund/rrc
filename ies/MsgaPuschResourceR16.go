package ies

import "rrc/utils"

// MsgA-PUSCH-Resource-r16 ::= SEQUENCE
// Extensible
type MsgaPuschResourceR16 struct {
	MsgaMcsR16                        utils.INTEGER `lb:0,ub:15`
	NrofslotsmsgaPuschR16             utils.INTEGER `lb:0,ub:4`
	NrofmsgaPoPerslotR16              MsgaPuschResourceR16NrofmsgaPoPerslotR16
	MsgaPuschTimedomainoffsetR16      utils.INTEGER  `lb:0,ub:32`
	MsgaPuschTimedomainallocationR16  *utils.INTEGER `lb:0,ub:maxNrofULAllocations`
	StartsymbolandlengthmsgaPoR16     *utils.INTEGER `lb:0,ub:127`
	MappingtypemsgaPuschR16           *MsgaPuschResourceR16MappingtypemsgaPuschR16
	GuardperiodmsgaPuschR16           *utils.INTEGER `lb:0,ub:3`
	GuardbandmsgaPuschR16             utils.INTEGER  `lb:0,ub:1`
	FrequencystartmsgaPuschR16        utils.INTEGER  `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	NrofprbsPermsgaPoR16              utils.INTEGER  `lb:0,ub:32`
	NrofmsgaPoFdmR16                  MsgaPuschResourceR16NrofmsgaPoFdmR16
	MsgaIntraslotfrequencyhoppingR16  *MsgaPuschResourceR16MsgaIntraslotfrequencyhoppingR16
	MsgaHoppingbitsR16                *utils.BITSTRING `lb:2,ub:2`
	MsgaDmrsConfigR16                 MsgaDmrsConfigR16
	NrofdmrsSequencesR16              utils.INTEGER `lb:0,ub:2`
	MsgaAlphaR16                      *MsgaPuschResourceR16MsgaAlphaR16
	InterlaceindexfirstpoMsgaPuschR16 *utils.INTEGER `lb:0,ub:10`
	NrofinterlacespermsgaPoR16        *utils.INTEGER `lb:0,ub:10`
}
