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

require mechs v0.0.0-00010101000000-000000000000

require (
	colorscheme v0.0.0-00010101000000-000000000000 // indirect
	core_powers v0.0.0-00010101000000-000000000000 // indirect
	core_stats v0.0.0-00010101000000-000000000000 // indirect
	dataload v0.0.0-00010101000000-000000000000 // indirect
	frames v0.0.0-00010101000000-000000000000 // indirect
	github.com/rs/cors v1.11.1 // indirect
	mounts v0.0.0-00010101000000-000000000000 // indirect
	traits v0.0.0-00010101000000-000000000000 // indirect
	weapons v0.0.0-00010101000000-000000000000 // indirect
)
