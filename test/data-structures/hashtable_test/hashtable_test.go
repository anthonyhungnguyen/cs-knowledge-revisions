package hashtable_test

import (
	hashtable "hash_table"

	"github.com/stretchr/testify/suite"
)

type HashTableSuite struct {
	suite.Suite
	data hashtable.HashTable
}
