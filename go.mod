module main.go

go 1.23.6

replace credits => ./src/credits/

replace utils => ./src/core/utils/

replace menus => ./src/core/menus/

replace wireless => ./src/wireless/

replace crackers => ./src/crackers/

replace phishers => ./src/phishers/

replace handlers => ./src/handlers/

replace africana => ./src/core/main/

replace setups => ./src/core/setups/

replace internals => ./src/internals/

replace bcolors => ./src/core/bcolors/

replace banners => ./src/core/banners/

replace securities => ./src/securities/

replace agreements => ./src/agreements/

replace scriptures => ./src/scriptures/

replace webattackers => ./src/webattackers/

replace subprocess => ./src/core/subprocess/

require (
	africana v0.0.0-00010101000000-000000000000 // indirect
	agreements v0.0.0-00010101000000-000000000000 // indirect
	banners v0.0.0-00010101000000-000000000000 // indirect
	bcolors v0.0.0-00010101000000-000000000000 // indirect
	crackers v0.0.0-00010101000000-000000000000 // indirect
	credits v0.0.0-00010101000000-000000000000 // indirect
	handlers v0.0.0-00010101000000-000000000000 // indirect
	internals v0.0.0-00010101000000-000000000000 // indirect
	menus v0.0.0-00010101000000-000000000000 // indirect
	phishers v0.0.0-00010101000000-000000000000 // indirect
	scriptures v0.0.0-00010101000000-000000000000 // indirect
	securities v0.0.0-00010101000000-000000000000 // indirect
	setups v0.0.0-00010101000000-000000000000 // indirect
	subprocess v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
	webattackers v0.0.0-00010101000000-000000000000 // indirect
	wireless v0.0.0-00010101000000-000000000000 // indirect
)
