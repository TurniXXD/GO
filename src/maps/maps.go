package maps

import (
	"errors"
	"fmt"
	"log"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// Map provides key -> value object
	// Searching for something by a key in a map is faster then searching through slice
	userMap := make(map[string]user)

	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func addValueToNestedMap(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	// Check that the object in nested map exists
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country] = 420
}

type CountryCompositeKey struct {
	path, country string
}

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		// Boolean for element existence is not needed because count will just be 0
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

func getNameCounts(names []string) map[rune]map[string]int {
	// rune gives more assurity in type system
	counts := make(map[rune]map[string]int)

	// Continuously initialize new maps
	for _, name := range names {
		// Skip empty strings
		if name == "" {
			continue
		}
		firstChar := name[0]
		// Check if map with letter already exists
		_, ok := counts[rune(firstChar)]
		if !ok {
			// Initialize new map
			counts[rune(firstChar)] = make(map[string]int)
		}
		// We are sure that map exists now
		counts[rune(firstChar)][name]++
	}
	return counts
}

func userMap() {
	names := []string{"jacob", "john", "barry"}
	phoneNumbers := []int{735023812, 789654156, 598715879}

	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

	delete(users, "john")

	fmt.Println(users)

	// Nested map
	usersFields := make(map[string]map[string]int)

	addValueToNestedMap(usersFields, "jerry", "czechia")

	fmt.Println(usersFields)

	// Composite key map, most of the time this is the easier way of doing so
	userCountry := make(map[CountryCompositeKey]int)
	userCountry[CountryCompositeKey{"jerry", "slovakia"}] = 418

	fmt.Println(userCountry)

	userIds := []string{"a1", "a2", "a1", "a3"}
	fmt.Println(getCounts(userIds))

	fmt.Printf("Count users that start with same letter: %v", getNameCounts(names))
}

func Maps() {
	fmt.Println("\nMaps: ")

	userMap()
}
