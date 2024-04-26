module main.go

replace bcolors => ./bcolors

replace internals => ./internals

go 1.22.2

require bcolors v0.0.0-00010101000000-000000000000

require internals v0.0.0-00010101000000-000000000000 // indirect
