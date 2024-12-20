package replinternal

import (
	"fmt"
	"strings"
	"github.com/som-pat/poke_dex/internal/config"

)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.ConfigState, ...string) (string,[]string,error)
}

func get_command() map[string] cliCommand{
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays available commands",
			callback:    call_help,
		},

		"map":{
			name:		 "map",
			description: "Display next 20 loactions",
			callback:	 call_map,
		},
		
		"mapb":{
			name:		 "mapb",
			description: "Display previous 20 locations",
			callback:    call_mapb,
		},
		
		"explore":{
			name:		 "explore 'location_area' ",
			description: "Display Pokemons in chosen location",
			callback:    call_explore,
		},
		"scout":{
			name:		 "scout",
			description: "Scour the region to find Pokemons/Items",
			callback:    call_scout,
		},

		"battle":{
			name:		 "battle{Pokemon name}",
			description: "battle Wild Pokemons",
			callback:    call_battle,
		},

		"catch":{
			name:		 "catch{Pokemon name}",
			description: "Catch Pokemons",
			callback:    call_catch,
		},

		"forage":{
			name:		 "forage{Item name}",
			description: "Collect Item",
			callback:    call_forage,
		},

		"inventory":{
			name:		 "inventory",
			description: "View caught Pokemons/ held Items",
			callback:    call_inventory,
		},

		"inspect":{
			name:		 "inspect{Pokemon/Item name}",
			description: "Inspect caught Pokemons",
			callback:    call_pokeInspect,
		},

	}
}

func ReplInput(cfg_state *config.ConfigState, input string) (string,[]string){
	new_input := input_clean(input)
	
	// Empty commands
	if len(new_input) == 0{
		return "No command entered. Type 'help' for a list of commands.",nil
	}
	com := new_input[0]
	args := []string{}
	if len(new_input)>1{
		args = new_input[1:]
	}

	avail_com := get_command()
	
	// Check if valid command
	route_com,ok  := avail_com[com]
	if !ok{
		return "Unknown command. Type 'help' for a list of available commands.",nil
	}

	res, lis, err :=route_com.callback(cfg_state, args...)
	if err != nil {
		return fmt.Sprintf("Error: %v",err),nil
	}

	return res,lis
		
}

func input_clean(input string) ([]string) {
	lower :=  strings.ToLower(input)
	words :=  strings.Fields(lower)
	return words
}