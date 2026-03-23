module lancer

go 1.23.0

toolchain go1.24.10

replace mechs => ./mechs

replace frames => ./mechs/frames

replace traits => ./mechs/traits

replace weapons => ./mechs/weapons

replace mounts => ./mechs/mounts

replace colorscheme => ./mechs/colorscheme

replace core_stats => ./mechs/frames/core_stats

replace core_powers => ./mechs/frames/core_powers

replace dataload => ./data/dataload

require (
	dataload v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.11.1
	mechs v0.0.0-00010101000000-000000000000
)

require (
	colorscheme v0.0.0-00010101000000-000000000000 // indirect
	frames v0.0.0-00010101000000-000000000000 // indirect
	mounts v0.0.0-00010101000000-000000000000 // indirect
	weapons v0.0.0-00010101000000-000000000000 // indirect
)
