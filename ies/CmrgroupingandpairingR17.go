package ies

import "rrc/utils"

// CMRGroupingAndPairing-r17 ::= SEQUENCE
type CmrgroupingandpairingR17 struct {
	Nrofresourcesgroup1R17 utils.INTEGER `lb:0,ub:7`
	Pair1ofnzpCsiRsR17     *NzpCsiRsPairingR17
	Pair2ofnzpCsiRsR17     *NzpCsiRsPairingR17
}
