package protocols

import (
	v419 "github.com/potlounge/multiversion/protocols/v419"
	v486 "github.com/potlounge/multiversion/protocols/v486"
	"github.com/sandertv/gophertunnel/minecraft"
)

func All() []minecraft.Protocol {
	return []minecraft.Protocol{
		v419.New(),
		v486.New(),
	}
}
