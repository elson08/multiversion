module github.com/potlounge/multiversion

go 1.24.1

require (
	github.com/df-mc/dragonfly v0.10.9
	github.com/df-mc/worldupgrader v1.0.20
	github.com/go-gl/mathgl v1.2.0
	github.com/google/uuid v1.6.0
	github.com/rogpeppe/go-internal v1.14.1
	github.com/samber/lo v1.52.0
	github.com/sandertv/go-raknet v1.14.3-0.20250305181847-6af3e95113d6
	github.com/sandertv/gophertunnel v1.51.1
	github.com/segmentio/fasthash v1.0.3
	golang.org/x/exp v0.0.0-20251219203646-944ab1f22d93
	golang.org/x/image v0.34.0
)

require (
	github.com/brentp/intintmap v0.0.0-20251106190759-56907b1f8479 // indirect
	github.com/df-mc/goleveldb v1.1.9 // indirect
	github.com/df-mc/jsonc v1.0.5 // indirect
	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/klauspost/compress v1.18.2 // indirect
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/oauth2 v0.34.0 // indirect
	golang.org/x/text v0.32.0 // indirect
)

replace github.com/sandertv/go-raknet => ../tedac-raknet

replace github.com/sandertv/gophertunnel => ../tedac-gophertunnel

replace github.com/df-mc/dragonfly => ../dragonfly
