package internal

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/NinjaCrusader/pokedexcli/internal/pokecache"
)

type Maps struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Areas struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// First Checks if the api information is cached
// If it's not, calls the pokiapi and provides 20 area locations and adds that call to the cache
func GetMapHelper(URL string, cache *pokecache.Cache) (Maps, error) {
	var mapData Maps

	cacheCheck, ok := cache.Get(URL)
	if ok {
		json.Unmarshal(cacheCheck, &mapData)
		return mapData, nil
	}

	if len(URL) == 0 {
		URL = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(URL)
	if err != nil {
		return mapData, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	cache.Add(URL, body)

	if err := json.Unmarshal(body, &mapData); err != nil {
		return mapData, err
	}

	if res.StatusCode > 299 {
		return mapData, err
	}
	if err != nil {
		return mapData, err
	}
	return mapData, nil
}

func GetAreaInformationHelper(area string, cache *pokecache.Cache) (Areas, error) {
	var areaData Areas

	URL := "https://pokeapi.co/api/v2/location-area/" + area

	cacheCheck, ok := cache.Get(URL)
	if ok {
		json.Unmarshal(cacheCheck, &areaData)
		return areaData, nil
	}

	res, err := http.Get(URL)
	if err != nil {
		return areaData, nil
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	cache.Add(URL, body)

	if err := json.Unmarshal(body, &areaData); err != nil {
		return areaData, err
	}

	if res.StatusCode > 299 {
		return areaData, err
	}

	if err != nil {
		return areaData, err
	}

	return areaData, nil
}
