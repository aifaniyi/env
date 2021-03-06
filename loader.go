// Package loader provides primitives for fetching environement variables
// (or return a default value) as one of the common go primitive types i.e int, string
package env

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// LoadInt : load int64 environement variable if it is set or
// return a specified option value if not found
func LoadInt(key string, option int64) int64 {
	if envVar := loadTrimmed(key); envVar != "" {
		val, err := strconv.ParseInt(envVar, 10, 64)
		if err != nil {
			log.Printf("The provided environement value %s for %s is not an int, setting %s = %d",
				envVar, key, key, option)
			return option
		}
		log.Printf("Successfully loaded environment variable %s = %s\n", key, envVar)
		return val
	}
	log.Printf("Unable to load %s, setting %s = %d.\n", key, key, option)
	return option
}

// LoadFloat : load int environement variable if it is set or
// return a specified option value if not found
func LoadFloat(key string, option float64) float64 {
	if envVar := loadTrimmed(key); envVar != "" {
		val, err := strconv.ParseFloat(envVar, 64)
		if err != nil {
			log.Printf("The provided environement value %s for %s is not a float, setting %s = %f",
				envVar, key, key, option)
			return option
		}
		log.Printf("Successfully loaded environment variable %s = %f\n", key, val)
		return val
	}
	log.Printf("Unable to load %s, setting %s = %f.\n", key, key, option)
	return option
}

// LoadBool : load boolean environement variable if it is set or
// return a specified option value if not found
func LoadBool(key string, option bool) bool {
	if envVar := loadTrimmed(key); envVar != "" {
		val := false
		if envVar == "true" {
			val = true
		} else if envVar == "false" {
			val = false
		} else {
			log.Printf("The provided environment value %s for %s is not a boolean, setting %s = %t",
				envVar, key, key, option)
			return option
		}
		log.Printf("Successfully loaded environment variable %s = %t\n", key, val)
		return val
	}
	log.Printf("Unable to load %s, setting %s = %t.\n", key, key, option)
	return option
}

// LoadString : load string environement variable if it is set or
// return a specified option value if not found
func LoadString(key string, option string) string {
	if envVar := loadTrimmed(key); envVar != "" {
		log.Printf("Successfully loaded environment variable %s = %s\n", key, envVar)
		return envVar
	}
	log.Printf("Unable to load %s, setting %s = %s.\n", key, key, option)
	return option
}

// LoadArray : load array environement variable if set, where
// values are separated by a specified separator, or
// return default value if not set
func LoadArray(key, separator string, option []string) []string {
	val := strings.Split(loadTrimmed(key), separator)
	entries := make([]string, 0)
	for _, entry := range val {
		if entry != "" {
			entries = append(entries, strings.TrimSpace(entry))
		}
	}
	if len(entries) > 0 { // valid if at least one entry exists
		log.Printf("Successfully loaded environment variable %s = %v\n", key, entries)
		return entries
	}
	log.Printf("Unable to load %s, setting %s = %s.\n", key, key, option)
	return option
}

func loadTrimmed(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
