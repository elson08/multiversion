package types

import (
	"math"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Command holds the data that a command requires to be shown to a player client-side. The command is shown in
// the /help command and auto-completed using this data.
type Command struct {
	// Name is the name of the command. The command may be executed using this name, and will be shown in the
	// /help list with it. It currently seems that the client crashes if the Name contains uppercase letters.
	Name string
	// Description is the description of the command. It is shown in the /help list and when starting to write
	// a command.
	Description string
	// Flags is a combination of flags not currently known. Leaving the Flags field empty appears to work.
	Flags uint16
	// PermissionLevel is the command permission level that the player required to execute this command. The
	// field no longer seems to serve a purpose, as the client does not handle the execution of commands
	// anymore: The permissions should be checked server-side.
	PermissionLevel byte
	// AliasesOffset is the offset to a CommandEnum that holds the values that
	// should be used as aliases for this command.
	AliasesOffset uint32
	// Overloads is a list of command overloads that specify the ways in which a command may be executed. The
	// overloads may be completely different.
	Overloads []CommandOverload
}

func (c *Command) Marshal(r protocol.IO) {
	r.String(&c.Name)
	r.String(&c.Description)
	r.Uint16(&c.Flags)
	r.Uint8(&c.PermissionLevel)
	r.Uint32(&c.AliasesOffset)
	protocol.Slice(r, &c.Overloads)
}

// CommandOverload represents an overload of a command. This overload can be compared to function overloading
// in languages such as java. It represents a single usage of the command. A command may have multiple
// different overloads, which are handled differently.
type CommandOverload struct {
	// Parameters is a list of command parameters that are part of the overload. These parameters specify the
	// usage of the command when this overload is applied.
	Parameters []protocol.CommandParameter
}

func (c *CommandOverload) Marshal(r protocol.IO) {
	protocol.Slice(r, &c.Parameters)
}

const (
	CommandOriginPlayer = iota
	CommandOriginBlock
	CommandOriginMinecartBlock
	CommandOriginDevConsole
	CommandOriginTest
	CommandOriginAutomationPlayer
	CommandOriginClientAutomation
	CommandOriginDedicatedServer
	CommandOriginEntity
	CommandOriginVirtual
	CommandOriginGameArgument
	CommandOriginEntityServer
)

// CommandOrigin holds data that identifies the origin of the requesting of a command. It holds several
// fields that may be used to get specific information.
// When sent in a CommandRequest packet, the same CommandOrigin should be sent in a CommandOutput packet.
type CommandOrigin struct {
	// Origin is one of the values above that specifies the origin of the command. The origin may change,
	// depending on what part of the client actually called the command. The command may be issued by a
	// websocket server, for example.
	Origin uint32
	// UUID is the UUID of the command called. This UUID is a bit odd as it is not specified by the server. It
	// is not clear what exactly this UUID is meant to identify, but it is unique for each command called.
	UUID uuid.UUID
	// RequestID is an ID that identifies the request of the client. The server should send a CommandOrigin
	// with the same request ID to ensure it can be matched with the request by the caller of the command.
	// This is especially important for websocket servers and it seems that this field is only non-empty for
	// these websocket servers.
	RequestID string
	// PlayerUniqueID is an ID that identifies the player, the same as the one found in the AdventureSettings
	// packet. Filling it out with 0 seems to work.
	// PlayerUniqueID is only written if Origin is CommandOriginDevConsole or CommandOriginTest.
	PlayerUniqueID int64
}

// CommandOriginData reads/writes a CommandOrigin x using IO r.
func CommandOriginData(r protocol.IO, x *CommandOrigin) {
	r.Varuint32(&x.Origin)
	r.UUID(&x.UUID)
	r.String(&x.RequestID)
	if x.Origin == CommandOriginDevConsole || x.Origin == CommandOriginTest {
		r.Varint64(&x.PlayerUniqueID)
	}
}

// CommandEnum represents an enum in a command usage. The enum typically has a type and a set of options that
// are valid. A value that is not one of the options results in a failure during execution.
type CommandEnum struct {
	// Type is the type of the command enum. The type will show up in the command usage as the type of the
	// argument if it has a certain amount of arguments, or when Options is set to true in the
	// command holding the enum.
	Type string
	// ValueIndices holds a list of indices that point to the EnumValues slice in the
	// AvailableCommandsPacket. These represent the options of the enum.
	ValueIndices []uint
}

// CommandEnumContext holds context required for encoding command enums.
type CommandEnumContext struct {
	EnumValues []string
}

// Marshal encodes/decodes a CommandEnum.
func (ctx CommandEnumContext) Marshal(r protocol.IO, x *CommandEnum) {
	r.String(&x.Type)
	protocol.FuncIOSlice(r, &x.ValueIndices, ctx.enumOption)
}

// enumOption writes/reads a command enum option as a byte/uint16/uint32,
// depending on the amount of enum values.
func (ctx CommandEnumContext) enumOption(r protocol.IO, opt *uint) {
	n := len(ctx.EnumValues)
	switch {
	case n <= math.MaxUint8:
		val := byte(*opt)
		r.Uint8(&val)
		*opt = uint(val)
	case n <= math.MaxUint16:
		val := uint16(*opt)
		r.Uint16(&val)
		*opt = uint(val)
	default:
		val := uint32(*opt)
		r.Uint32(&val)
		*opt = uint(val)
	}
}
