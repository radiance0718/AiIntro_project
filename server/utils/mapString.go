package utils

import "hash/fnv"

func MapStringToNumber(input string) int {
	hash := fnv.New32a()
	hash.Write([]byte(input))
	hashValue := hash.Sum32()

	return int(hashValue % 6)
}
