package types

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// ItemStackResponse is a response to an individual ItemStackRequest.
type ItemStackResponse struct {
	// Status specifies if the request with the RequestID below was successful. If this is the case, the
	// ContainerInfo below will have information on what slots ended up changing. If not, the container info
	// will be empty.
	// A non-0 status means an error occurred and will result in the action being reverted.
	Status uint8
	// RequestID is the unique ID of the request that this response is in reaction to. If rejected, the client
	// will undo the actions from the request with this ID.
	RequestID int32
	// ContainerInfo holds information on the containers that had their contents changed as a result of the
	// request.
	ContainerInfo []StackResponseContainerInfo
}

// Marshal encodes/decodes an ItemStackResponse.
func (x *ItemStackResponse) Marshal(r protocol.IO) {
	r.Uint8(&x.Status)
	r.Varint32(&x.RequestID)
	if x.Status == protocol.ItemStackResponseStatusOK {
		protocol.Slice(r, &x.ContainerInfo)
	}
}

// StackResponseContainerInfo holds information on what slots in a container have what item stack in them.
type StackResponseContainerInfo struct {
	// ContainerID is the container ID of the container that the slots that follow are in. For the main
	// inventory, this value seems to be 0x1b. For the cursor, this value seems to be 0x3a. For the crafting
	// grid, this value seems to be 0x0d.
	ContainerID byte
	// SlotInfo holds information on what item stack should be present in specific slots in the container.
	SlotInfo []StackResponseSlotInfo
}

// Marshal encodes/decodes a StackResponseContainerInfo.
func (x *StackResponseContainerInfo) Marshal(r protocol.IO) {
	r.Uint8(&x.ContainerID)
	protocol.Slice(r, &x.SlotInfo)
}

// StackResponseSlotInfo holds information on what item stack should be present in a specific slot.
type StackResponseSlotInfo struct {
	protocol.StackResponseSlotInfo
}

// Marshal encodes/decodes a StackResponseSlotInfo.
func (x *StackResponseSlotInfo) Marshal(r protocol.IO) {
	r.Uint8(&x.Slot)
	r.Uint8(&x.HotbarSlot)
	r.Uint8(&x.Count)
	r.Varint32(&x.StackNetworkID)
	if x.Slot != x.HotbarSlot {
		r.InvalidValue(x.HotbarSlot, "hotbar slot", "hot bar slot must be equal to normal slot")
	}
}
