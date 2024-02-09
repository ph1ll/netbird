package integrated_validator

import (
	"github.com/netbirdio/netbird/management/server/account"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedValidator interface {
	PreparePeer(peer *nbpeer.Peer, extraSettings *account.ExtraSettings, groups []string) *nbpeer.Peer
	ValidatePeer(peer *nbpeer.Peer, groups []string) (bool, error)
}
