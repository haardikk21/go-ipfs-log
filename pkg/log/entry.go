package log

import (
	"github.com/haardikk21/go-ipfs-log/pkg/identity"
	"github.com/ipfs/go-ipfs/core"
	"github.com/pkg/errors"
)

type Entry struct {
	IPFS     *core.IpfsNode
	Identity *identity.Identity
	LogID    string
	Payload  interface{}
	Next     []interface{}
	Clock    LamportClock
	Hash     string
	Version  int
}

func NewEntry(ipfs *core.IpfsNode, identity *identity.Identity, logID string, data interface{}, next []interface{}, clock LamportClock) (*Entry, error) {
	if ipfs == nil {
		return nil, errors.New("ipfs instance not defined")
	}

	if identity == nil {
		return nil, errors.New("identity not defined")
	}

	if logID == "" {
		return nil, errors.New("logId not defined")
	}

	if data == nil {
		return nil, errors.New("data not defined")
	}

	if next == nil {
		return nil, errors.New("next not defined")
	}

	var nexts []interface{}

	for _, n := range next {
		if n != nil {
			nexts = append(nexts, toEntry(n))
		}
	}

	entry := Entry{
		Hash:    "",
		LogID:   logID,
		Payload: data,
		Next:    nexts,
		Version: 2,
		Clock:   clock,
	}

	return &entry, nil

}

func toEntry(e interface{}) string {
	switch t := e.(type) {
	case Entry:
		return t.Hash
	case string:
		return t
	default:
		panic("invalid type")
	}

}
