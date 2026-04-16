package v419

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"

	"github.com/potlounge/multiversion/protocols/v419/types"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// NewReader creates a new initialised Reader with an underlying protocol.Reader to write to.
func NewReader(r *protocol.Reader) *Reader {
	return &Reader{r}
}

type Reader struct {
	*protocol.Reader
}

// SoundPos reads an mgl32.Vec3 that serves as a position for a sound.
func (r *Reader) SoundPos(x *mgl32.Vec3) {
	var b protocol.BlockPos
	r.UBlockPos(&b)
	*x = mgl32.Vec3{float32(b[0]) / 8, float32(b[1]) / 8, float32(b[2]) / 8}
}

// panicf panics with the format and values passed and assigns the error created to the Reader.
func (r *Reader) panicf(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}

func (r *Reader) LimitUint32(value uint32, max uint32) {
	if max == math.MaxUint32 {
		// Account for 0-1 overflowing into max.
		max = 0
	}
	if value > max {
		r.panicf("uint32 %v exceeds maximum of %v", value, max)
	}
}

func (r *Reader) LimitInt32(value int32, min, max int32) {
	if value < min {
		r.panicf("int32 %v exceeds minimum of %v", value, min)
	} else if value > max {
		r.panicf("int32 %v exceeds maximum of %v", value, max)
	}
}

func (r *Reader) Reads() bool {
	return true
}

func (r *Reader) LimitsEnabled() bool {
	return r.Reader.LimitsEnabled()
}

func (r *Reader) ItemInstance(i *protocol.ItemInstance) {
	r.Item(&i.Stack)
}

func (r *Reader) Item(x *protocol.ItemStack) {
	x.NBTData = make(map[string]any)
	r.Varint32(&x.NetworkID)
	if x.NetworkID == 0 {
		// The item was air, so there is no more data we should read for the item instance. After all, air
		// items aren't really anything.
		x.MetadataValue, x.Count, x.CanBePlacedOn, x.CanBreak = 0, 0, nil, nil
		return
	}

	var auxValue int32
	r.Varint32(&auxValue)
	x.MetadataValue = uint32(auxValue >> 8)
	x.Count = uint16(auxValue & 0xff)

	var length int16
	r.Int16(&length)

	if length == -1 {
		var version uint8
		r.Uint8(&version)

		switch version {
		case 1:
			r.NBT(&x.NBTData, nbt.NetworkLittleEndian)
		default:
			r.UnknownEnumOption(version, "item user data version")
			return
		}
	} else if length > 0 {
		r.NBT(&x.NBTData, nbt.LittleEndian)
	}

	var count int32

	r.Varint32(&count)
	x.CanBePlacedOn = make([]string, count)
	for i := int32(0); i < count; i++ {
		r.String(&x.CanBePlacedOn[i])
	}

	r.Varint32(&count)
	x.CanBreak = make([]string, count)
	for i := int32(0); i < count; i++ {
		r.String(&x.CanBreak[i])
	}

	if x.NetworkID == r.ShieldID() {
		var blockingTick int64
		r.Int64(&blockingTick)
	}
}

func (r *Reader) PlayerInventoryAction(x *protocol.UseItemTransactionData) {
	r.Varint32(&x.LegacyRequestID)
	if x.LegacyRequestID < -1 && (x.LegacyRequestID&1) == 0 {
		protocol.Slice(r, &x.LegacySetItemSlots)
	}
	protocol.Slice(r, &x.Actions)
	r.Varuint32(&x.ActionType)
	r.BlockPos(&x.BlockPosition)
	r.Varint32(&x.BlockFace)
	r.Varint32(&x.HotBarSlot)
	r.ItemInstance(&x.HeldItem)
	r.Vec3(&x.Position)
	r.Vec3(&x.ClickedPosition)
	r.Varuint32(&x.BlockRuntimeID)
}

func (r *Reader) GameRule(x *protocol.GameRule) {
	r.String(&x.Name)
	var t uint32
	r.Varuint32(&t)

	switch t {
	case 1:
		var v bool
		r.Bool(&v)
		x.Value = v
	case 2:
		var v uint32
		r.Varuint32(&v)
		x.Value = v
	case 3:
		var v float32
		r.Float32(&v)
		x.Value = v
	default:
		r.UnknownEnumOption(t, "game rule type")
	}
}

func (r *Reader) StackRequestAction(x *protocol.StackRequestAction) {
	var id uint8
	r.Uint8(&id)
	if !lookupStackRequestAction(id, x) {
		r.UnknownEnumOption(id, "stack request action type")
		return
	}
	(*x).Marshal(r)
}

func lookupStackRequestAction(id uint8, x *protocol.StackRequestAction) bool {
	switch id {
	case protocol.StackRequestActionTake:
		*x = &types.TakeStackRequestAction{TakeStackRequestAction: protocol.TakeStackRequestAction{}}
	case protocol.StackRequestActionPlace:
		*x = &types.PlaceStackRequestAction{PlaceStackRequestAction: protocol.PlaceStackRequestAction{}}
	case protocol.StackRequestActionSwap:
		*x = &types.SwapStackRequestAction{SwapStackRequestAction: protocol.SwapStackRequestAction{}}
	case protocol.StackRequestActionDrop:
		*x = &types.DropStackRequestAction{DropStackRequestAction: protocol.DropStackRequestAction{}}
	case protocol.StackRequestActionDestroy:
		*x = &types.DestroyStackRequestAction{DestroyStackRequestAction: protocol.DestroyStackRequestAction{}}
	case protocol.StackRequestActionConsume:
		*x = &types.ConsumeStackRequestAction{DestroyStackRequestAction: protocol.DestroyStackRequestAction{}}
	case protocol.StackRequestActionCreate:
		*x = &protocol.CreateStackRequestAction{}
	case protocol.StackRequestActionPlaceInContainer:
		*x = &types.PlaceInContainerStackRequestAction{PlaceInContainerStackRequestAction: protocol.PlaceInContainerStackRequestAction{}}
	case protocol.StackRequestActionTakeOutContainer:
		*x = &types.TakeOutContainerStackRequestAction{TakeOutContainerStackRequestAction: protocol.TakeOutContainerStackRequestAction{}}
	case protocol.StackRequestActionLabTableCombine:
		*x = &protocol.LabTableCombineStackRequestAction{}
	case protocol.StackRequestActionBeaconPayment:
		*x = &protocol.BeaconPaymentStackRequestAction{}
	case protocol.StackRequestActionMineBlock:
		*x = &protocol.MineBlockStackRequestAction{}
	case protocol.StackRequestActionCraftRecipe:
		*x = &protocol.CraftRecipeStackRequestAction{}
	case protocol.StackRequestActionCraftRecipeAuto:
		*x = &types.AutoCraftRecipeStackRequestAction{AutoCraftRecipeStackRequestAction: protocol.AutoCraftRecipeStackRequestAction{}}
	case protocol.StackRequestActionCraftCreative:
		*x = &protocol.CraftCreativeStackRequestAction{}
	case protocol.StackRequestActionCraftRecipeOptional:
		*x = &protocol.CraftRecipeOptionalStackRequestAction{}
	case protocol.StackRequestActionCraftGrindstone:
		*x = &protocol.CraftGrindstoneRecipeStackRequestAction{}
	case protocol.StackRequestActionCraftLoom:
		*x = &protocol.CraftLoomRecipeStackRequestAction{}
	case protocol.StackRequestActionCraftNonImplementedDeprecated:
		*x = &protocol.CraftNonImplementedStackRequestAction{}
	case protocol.StackRequestActionCraftResultsDeprecated:
		*x = &protocol.CraftResultsDeprecatedStackRequestAction{}
	default:
		return false
	}
	return true
}

func (r *Reader) TransactionDataType(x *protocol.InventoryTransactionData) {
	var transactionType uint32
	r.Varuint32(&transactionType)
	if !lookupTransactionData(transactionType, x) {
		r.UnknownEnumOption(transactionType, "inventory transaction data type")
	}
}

func lookupTransactionData(id uint32, x *protocol.InventoryTransactionData) bool {
	switch id {
	case protocol.InventoryTransactionTypeNormal:
		*x = &protocol.NormalTransactionData{}
	case protocol.InventoryTransactionTypeMismatch:
		*x = &protocol.MismatchTransactionData{}
	case protocol.InventoryTransactionTypeUseItem:
		*x = &types.UseItemTransactionData{}
	case protocol.InventoryTransactionTypeUseItemOnEntity:
		*x = &protocol.UseItemOnEntityTransactionData{}
	case protocol.InventoryTransactionTypeReleaseItem:
		*x = &protocol.ReleaseItemTransactionData{}
	default:
		return false
	}
	return true
}

// NewWriter creates a new initialised Writer with an underlying protocol.Writer to write to.
func NewWriter(w *protocol.Writer) *Writer {
	return &Writer{w}
}

type Writer struct {
	*protocol.Writer
}

// SoundPos writes an mgl32.Vec3 that serves as a position for a sound.
func (w *Writer) SoundPos(x *mgl32.Vec3) {
	b := protocol.BlockPos{int32((*x)[0] * 8), int32((*x)[1] * 8), int32((*x)[2] * 8)}
	w.UBlockPos(&b)
}

// ItemInstance writes an ItemInstance i to the underlying buffer.
func (w *Writer) ItemInstance(i *protocol.ItemInstance) {
	w.Item(&i.Stack)
}

// Item writes an ItemStack x to the underlying buffer.
func (w *Writer) Item(x *protocol.ItemStack) {
	w.Varint32(&x.NetworkID)
	if x.NetworkID == 0 {
		// The item was air, so there's no more data to follow. Return immediately.
		return
	}

	aux := int32(x.MetadataValue<<8) | int32(x.Count)
	w.Varint32(&aux)

	var length int16
	if len(x.NBTData) != 0 {
		length = int16(-1)
		version := uint8(1)

		w.Int16(&length)
		w.Uint8(&version)
		w.NBT(&x.NBTData, nbt.NetworkLittleEndian)
	} else {
		w.Int16(&length)
	}

	placeOnLen := int32(len(x.CanBePlacedOn))
	canBreak := int32(len(x.CanBreak))

	w.Varint32(&placeOnLen)
	for _, block := range x.CanBePlacedOn {
		w.String(&block)
	}
	w.Varint32(&canBreak)
	for _, block := range x.CanBreak {
		w.String(&block)
	}

	if x.NetworkID == w.ShieldID() {
		var blockingTick int64
		w.Varint64(&blockingTick)
	}
}

func (w *Writer) GameRule(x *protocol.GameRule) {
	w.String(&x.Name)

	switch v := x.Value.(type) {
	case bool:
		id := uint32(1)
		w.Varuint32(&id)
		w.Bool(&v)
	case uint32:
		id := uint32(2)
		w.Varuint32(&id)
		w.Varuint32(&v)
	case float32:
		id := uint32(3)
		w.Varuint32(&id)
		w.Float32(&v)
	default:
		w.UnknownEnumOption(fmt.Sprintf("%T", v), "game rule type")
	}
}

func (w *Writer) ItemDescriptorCount(i *protocol.ItemDescriptorCount) {
	var id byte
	switch descriptor := i.Descriptor.(type) {
	case *protocol.InvalidItemDescriptor:
		id = protocol.ItemDescriptorInvalid
	case *protocol.DefaultItemDescriptor:
		id = protocol.ItemDescriptorDefault
	case *protocol.MoLangItemDescriptor:
		id = protocol.ItemDescriptorMoLang
	case *protocol.ItemTagItemDescriptor:
		id = protocol.ItemDescriptorItemTag
	case *protocol.DeferredItemDescriptor:
		id = protocol.ItemDescriptorDeferred
	case *protocol.ComplexAliasItemDescriptor:
		id = protocol.ItemDescriptorComplexAlias
	case *types.DefaultItemDescriptor:
		descriptor.Marshal(w)
		if descriptor.NetworkID != 0 {
			w.Varint32(&i.Count)
		}
		return
	default:
		w.UnknownEnumOption(fmt.Sprintf("%T", i.Descriptor), "item descriptor type")
		return
	}
	w.Uint8(&id)

	i.Descriptor.Marshal(w)
	w.Varint32(&i.Count)
}

func (w *Writer) Recipe(x *protocol.Recipe) {
	var recipeType int32
	if !lookupRecipeType(*x, &recipeType) {
		w.UnknownEnumOption(fmt.Sprintf("%T", *x), "crafting recipe type")
	}
	w.Varint32(&recipeType)
	(*x).Marshal(w)
}

func lookupRecipeType(x protocol.Recipe, recipeType *int32) bool {
	switch x.(type) {
	case *protocol.ShapelessRecipe:
		*recipeType = protocol.RecipeShapeless
	case *protocol.ShapedRecipe:
		*recipeType = protocol.RecipeShaped
	case *protocol.FurnaceRecipe:
		*recipeType = protocol.RecipeFurnace
	case *protocol.FurnaceDataRecipe:
		*recipeType = protocol.RecipeFurnaceData
	case *protocol.MultiRecipe:
		*recipeType = protocol.RecipeMulti
	case *protocol.ShulkerBoxRecipe:
		*recipeType = protocol.RecipeShulkerBox
	case *protocol.ShapelessChemistryRecipe:
		*recipeType = protocol.RecipeShapelessChemistry
	case *protocol.ShapedChemistryRecipe:
		*recipeType = protocol.RecipeShapedChemistry
	case *protocol.SmithingTransformRecipe:
		*recipeType = protocol.RecipeSmithingTransform
	case *protocol.SmithingTrimRecipe:
		*recipeType = protocol.RecipeSmithingTrim
	default:
		return false
	}
	return true
}
