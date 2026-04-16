package types

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// CraftRecipeStackRequestAction is sent by the client the moment it begins crafting an item. This is the
// first action sent, before the Consume and Create item stack request actions.
// This action is also sent when an item is enchanted. Enchanting should be treated mostly the same way as
// crafting, where the old item is consumed.
type CraftRecipeStackRequestAction struct {
	// RecipeNetworkID is the network ID of the recipe that is about to be crafted. This network ID matches
	// one of the recipes sent in the CraftingData packet, where each of the recipes have a RecipeNetworkID as
	// of 1.16.
	RecipeNetworkID uint32
}

// Marshal ...
func (a *CraftRecipeStackRequestAction) Marshal(r protocol.IO) {
	r.Varuint32(&a.RecipeNetworkID)
}

// CraftCreativeStackRequestAction is sent by the client when it takes an item out fo the creative inventory.
// The item is thus not really crafted, but instantly created.
type CraftCreativeStackRequestAction struct {
	// CreativeItemNetworkID is the network ID of the creative item that is being created. This is one of the
	// creative item network IDs sent in the CreativeContent packet.
	CreativeItemNetworkID uint32
}

// Marshal ...
func (a *CraftCreativeStackRequestAction) Marshal(r protocol.IO) {
	r.Varuint32(&a.CreativeItemNetworkID)
}

// CraftRecipeOptionalStackRequestAction is sent when using an anvil. When this action is sent, the
// FilterStrings field in the respective stack request is non-empty and contains the name of the item created
// using the anvil or cartography table.
type CraftRecipeOptionalStackRequestAction struct {
	// RecipeNetworkID is the network ID of the multi-recipe that is about to be crafted. This network ID matches
	// one of the multi-recipes sent in the CraftingData packet, where each of the recipes have a RecipeNetworkID as
	// of 1.16.
	RecipeNetworkID uint32
	// FilterStringIndex is the index of a filter string sent in a ItemStackRequest.
	FilterStringIndex int32
}

// Marshal ...
func (c *CraftRecipeOptionalStackRequestAction) Marshal(r protocol.IO) {
	r.Varuint32(&c.RecipeNetworkID)
	r.Int32(&c.FilterStringIndex)
}

// CraftGrindstoneRecipeStackRequestAction is sent when a grindstone recipe is crafted. It contains the RecipeNetworkID
// to identify the recipe crafted, and the cost for crafting the recipe.
type CraftGrindstoneRecipeStackRequestAction struct {
	// RecipeNetworkID is the network ID of the recipe that is about to be crafted. This network ID matches
	// one of the recipes sent in the CraftingData packet, where each of the recipes have a RecipeNetworkID as
	// of 1.16.
	RecipeNetworkID uint32
	// Cost is the cost of the recipe that was crafted.
	Cost int32
}

// Marshal ...
func (c *CraftGrindstoneRecipeStackRequestAction) Marshal(r protocol.IO) {
	r.Varuint32(&c.RecipeNetworkID)
	r.Varint32(&c.Cost)
}

// StackRequestSlotInfo holds information on a specific slot client-side.
type StackRequestSlotInfo struct {
	// ContainerID is the ID of the container that the slot was in.
	ContainerID byte
	// Slot is the index of the slot within the container with the ContainerID above.
	Slot byte
	// StackNetworkID is the unique stack ID that the client assumes to be present in this slot. The server
	// must check if these IDs match. If they do not match, servers should reject the stack request that the
	// action holding this info was in.
	StackNetworkID int32
}
