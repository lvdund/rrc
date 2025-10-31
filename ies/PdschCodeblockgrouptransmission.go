package ies

import "rrc/utils"

// PDSCH-CodeBlockGroupTransmission ::= SEQUENCE
// Extensible
type PdschCodeblockgrouptransmission struct {
	Maxcodeblockgroupspertransportblock PdschCodeblockgrouptransmissionMaxcodeblockgroupspertransportblock
	Codeblockgroupflushindicator        utils.BOOLEAN
}
