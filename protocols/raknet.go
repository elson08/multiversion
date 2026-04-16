package protocols

import (
	"context"
	"log/slog"
	"net"
	"time"

	"github.com/sandertv/go-raknet"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// MultiRakNet is an implementation of a RakNet v9/10 Network.
type MultiRakNet struct {
	l *slog.Logger
}

// DialContext ...
func (r MultiRakNet) DialContext(ctx context.Context, address string) (net.Conn, error) {
	return raknet.Dialer{
		ErrorLog: slog.Default(),
	}.DialContext(ctx, address)
}

// PingContext ...
func (r MultiRakNet) PingContext(ctx context.Context, address string) (response []byte, err error) {
	return raknet.Dialer{
		ErrorLog: slog.Default(),
	}.PingContext(ctx, address)
}

// Listen ...
func (r MultiRakNet) Listen(address string) (minecraft.NetworkListener, error) {
	return raknet.ListenConfig{
		ErrorLog:         slog.Default(),
		ProtocolVersions: []byte{10},
		BlockDuration:    time.Second * 10,
	}.Listen(address)
}

// Compression ...
func (MultiRakNet) Compression(conn net.Conn) packet.Compression {
	return packet.FlateCompression
}

// init registers the MultiRakNet network. It overrides the existing minecraft.RakNet network.
func init() {
	minecraft.RegisterNetwork("raknet", func(l *slog.Logger) minecraft.Network {
		return MultiRakNet{l: l}
	})
}
