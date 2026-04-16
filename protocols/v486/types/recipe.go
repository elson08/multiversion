package types

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Recipe represents a recipe that may be sent in a CraftingData packet to let the client know what recipes
// are available server-side.
type Recipe interface {
	// Marshal encodes the recipe data to its binary representation into buf.
	Marshal(w *protocol.Writer)
	// Unmarshal decodes a serialised recipe from Reader r into the recipe instance.
	Unmarshal(r *protocol.Reader)
}

// ShapelessRecipe is a recipe that has no particular shape. Its functionality is shared with the
// RecipeShulkerBox and RecipeShapelessChemistry types.
type ShapelessRecipe struct {
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	RecipeID string
	// Input is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output.
	Input []RecipeIngredientItem
	// Output is a list of items that are created as a result of crafting the recipe.
	Output []protocol.ItemStack
	// UUID is a UUID identifying the recipe. This can actually be set to an empty UUID if the CraftingEvent
	// packet is not used.
	UUID uuid.UUID
	// Block is the block name that is required to craft the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'crafting_table' as an example.
	// The available blocks are:
	// - crafting_table
	// - cartography_table
	// - stonecutter
	// - furnace
	// - blast_furnace
	// - smoker
	// - campfire
	Block string
	// Priority ...
	Priority int32
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
}

// Marshal ...
func (recipe *ShapelessRecipe) Marshal(w *protocol.Writer) {
	marshalShapeless(w, recipe)
}

// Unmarshal ...
func (recipe *ShapelessRecipe) Unmarshal(r *protocol.Reader) {
	marshalShapeless(r, recipe)
}

// ShapedRecipe is a recipe that has a specific shape that must be used to craft the output of the recipe.
// Trying to craft the item in any other shape will not work. The ShapedRecipe is of the same structure as the
// ShapedChemistryRecipe.
type ShapedRecipe struct {
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	RecipeID string
	// Width is the width of the recipe's shape.
	Width int32
	// Height is the height of the recipe's shape.
	Height int32
	// Input is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output. The amount of input items must be exactly equal to Width * Height.
	Input []RecipeIngredientItem
	// Output is a list of items that are created as a result of crafting the recipe.
	Output []protocol.ItemStack
	// UUID is a UUID identifying the recipe. This can actually be set to an empty UUID if the CraftingEvent
	// packet is not used.
	UUID uuid.UUID
	// Block is the block name that is required to craft the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'crafting_table' as an example.
	Block string
	// Priority ...
	Priority int32
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	RecipeNetworkID uint32
}

// Marshal ...
func (recipe *ShapedRecipe) Marshal(w *protocol.Writer) {
	marshalShaped(w, recipe)
}

// Unmarshal ...
func (recipe *ShapedRecipe) Unmarshal(r *protocol.Reader) {
	marshalShaped(r, recipe)
}

// marshalShaped ...
func marshalShaped(r protocol.IO, recipe *ShapedRecipe) {
	r.String(&recipe.RecipeID)
	r.Varint32(&recipe.Width)
	r.Varint32(&recipe.Height)
	protocol.FuncIOSliceOfLen(r, uint32(recipe.Width*recipe.Height), &recipe.Input, protocol.Single[RecipeIngredientItem])
	protocol.FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShapeless ...
func marshalShapeless(r protocol.IO, recipe *ShapelessRecipe) {
	r.String(&recipe.RecipeID)
	protocol.Slice(r, &recipe.Input)
	protocol.FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	r.Varuint32(&recipe.RecipeNetworkID)
}
