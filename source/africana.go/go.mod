module main.go

go 1.22.2

replace bcolors => ./src/core/bcolors/

replace agreements => ./src/agreements/

replace subprocess => ./src/core/subprocess/

replace banners => ./src/core/banners/

replace menus => ./src/core/menus/

replace utils => ./src/core/utils/

replace internals => ./src/internals/

replace butchers => ./src/butchers/

replace wireless => ./src/wireless/

replace crackers => ./src/crackers/

replace phishers => ./src/phishers/

replace webattackers => ./src/webattackers/

require (
	agreements v0.0.0-00010101000000-000000000000 // indirect
	banners v0.0.0-00010101000000-000000000000 // indirect
	bcolors v0.0.0-00010101000000-000000000000 // indirect
	butchers v0.0.0-00010101000000-000000000000 // indirect
	crackers v0.0.0-00010101000000-000000000000 // indirect
	internals v0.0.0-00010101000000-000000000000 // indirect
	menus v0.0.0-00010101000000-000000000000 // indirect
	phishers v0.0.0-00010101000000-000000000000 // indirect
	subprocess v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
	webattackers v0.0.0-00010101000000-000000000000 // indirect
	wireless v0.0.0-00010101000000-000000000000 // indirect
)
