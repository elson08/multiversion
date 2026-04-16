package types

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// ItemEntry is an item sent in the StartGame item table. It holds a name and a legacy ID, which is used to
// point back to that name.
type ItemEntry struct {
	// Name if the name of the item, which is a name like 'minecraft:stick'.
	Name string
	// RuntimeID is the ID that is used to identify the item over network. After sending all items in the
	// StartGame packet, items will then be identified using these numerical IDs.
	RuntimeID int16
	// ComponentBased specifies if the item was created using components, meaning the item is a custom item.
	ComponentBased bool
}

// Marshal encodes/decodes an ItemEntry.
func (x *ItemEntry) Marshal(r protocol.IO) {
	r.String(&x.Name)
	r.Int16(&x.RuntimeID)
	r.Bool(&x.ComponentBased)
}

// RecipeIngredientItem represents an item that may be used as a recipe ingredient.
type RecipeIngredientItem struct {
	// NetworkID is the numerical network ID of the item. This is sometimes a positive ID, and sometimes a
	// negative ID, depending on what item it concerns.
	NetworkID int32
	// MetadataValue is the metadata value of the item. For some items, this is the damage value, whereas for
	// other items it is simply an identifier of a variant of the item.
	MetadataValue int32
	// Count is the count of items that the recipe ingredient is required to have.
	Count int32
}

// Marshal encodes/decodes a RecipeIngredientItem.
func (x *RecipeIngredientItem) Marshal(r protocol.IO) {
	r.Varint32(&x.NetworkID)
	if x.NetworkID == 0 {
		return
	}
	r.Varint32(&x.MetadataValue)
	r.Varint32(&x.Count)
}
