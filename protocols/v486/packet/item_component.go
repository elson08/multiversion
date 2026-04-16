package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// ItemComponentEntry is sent in the ItemComponent item table. It holds a name and all of the components and
// properties associated to the item.
type ItemComponentEntry struct {
	// Name is the name of the item, which is a name like 'minecraft:stick'.
	Name string
	// Data is a map containing the components and properties of the item.
	Data map[string]any
}

// Marshal encodes/decodes an ItemComponentEntry.
func (x *ItemComponentEntry) Marshal(r protocol.IO) {
	r.String(&x.Name)
	r.NBT(&x.Data, nbt.NetworkLittleEndian)
}

// ItemComponent is sent by the server to attach client-side components to a custom item.
type ItemComponent struct {
	// Items holds a list of all custom items with their respective components set.
	Items []ItemComponentEntry
}

// ID ...
func (*ItemComponent) ID() uint32 {
	return packet.IDItemRegistry
}

func (pk *ItemComponent) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Items)
}
