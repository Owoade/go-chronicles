package main

import (
	"errors"
	"fmt"
)

func main() {

	jamb_score := make(map[string]int)

	jamb_score["english"] = 68

	jamb_score["physics"] = 75

	jamb_score["chemistry"] = 76

	jamb_score["biology"] = 77

	o_levels := map[string]string{
		"maths":   "A1",
		"f.maths": "A1",
		"english": "c6",
		"chem":    "b3",
		"physics": "b3",
		"bio":     "b3",
	}

	elem, ok := o_levels["mats"]

	delete(o_levels, "maths")

	if _, ok = o_levels["maths"]; !ok {
		fmt.Println(("Mathematics is compulsory"))
	}

	fmt.Println(ok, elem)

	delete(jamb_score, "chemistry")

	fmt.Println(jamb_score)

	user_map, err := map_name_to_phone([]string{"mark", "bayero", "collins"}, []int{900, 230})

	fmt.Println(user_map, err)

}

func map_name_to_phone(names []string, phones []int) (map[string]int, error) {

	name_phone_map := make(map[string]int)

	if len(names) != len(phones) {
		return nil, errors.New("length of phones and names must be equal")
	}

	for i, name := range names {
		name_phone_map[name] = phones[i]
	}

	return name_phone_map, nil

}
