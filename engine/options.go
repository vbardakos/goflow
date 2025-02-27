package engine

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const addrSep = "::"
const uidSep = "/"

type ActorOption func(*ActorConfig)

type ActorConfig struct {
	addr string
	kind string
	uid  *uint64
}

func defaultConfig() *ActorConfig {
	tmp := new(big.Int).Lsh(big.NewInt(1), 64)
	tmp, _ = rand.Int(rand.Reader, tmp)
	uid := tmp.Uint64()
	return &ActorConfig{uid: &uid}
}

func NewActorConfig(kind string, opts ...ActorOption) *ActorConfig {
	cfg := defaultConfig()
	cfg.kind = kind

	for _, fn := range opts {
		fn(cfg)
	}
	return cfg
}

func fromAddr(addr string) ActorOption {
	return func(cfg *ActorConfig) {
		cfg.addr = addr
	}
}

func withUid(uid uint64) ActorOption {
	return func(cfg *ActorConfig) {
		cfg.uid = &uid
	}
}

func withAtomicUid(cfg *ActorConfig) {
	cfg.uid = nil
}

func (c *ActorConfig) ActorID() *ActorID {
	return &ActorID{
		Address: c.Address(),
		Uid:     c.uid,
	}
}

func (c *ActorConfig) Address() string {
	return c.addr + addrSep + c.kind
}

func (c *ActorConfig) ParentAddress() string {
	return c.addr
}

func (c *ActorConfig) Kind() string {
	return c.kind
}

func (c *ActorConfig) IsAtomic() bool {
	return c.uid == nil
}

func (c *ActorConfig) UID() uint64 {
	return *c.uid
}

func (c *ActorConfig) Instance() string {
	if c.IsAtomic() {
		return c.Address()
	}
	return c.Address() + uidSep + fmt.Sprintf("%d", *c.uid)
}
