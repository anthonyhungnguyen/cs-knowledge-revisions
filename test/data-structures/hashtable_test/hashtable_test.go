package hashtable_test

import (
	hashtable "hash_table"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HashTableSuite struct {
	suite.Suite
	data hashtable.HashTable
}

func (suite *HashTableSuite) SetupSuite() {
	suite.data = *hashtable.NewHashTable()
}

func (suite *HashTableSuite) TestAddToEmptyHashTable() {
	hashTable := suite.data
	hashTable.Add(1, 1)
	res, _ := suite.data.Get(1)
	hashTable.Print()
	assert.Equal(suite.T(), 1, res, "Expected 1 but got %d", res)
}

func (suite *HashTableSuite) TestAddCollisionSameKey() {
	hashTable := suite.data
	hashTable.Add(1, 1)
	hashTable.Add(1, 2)
	res, _ := hashTable.Get(1)
	hashTable.Print()
	assert.Equal(suite.T(), 2, res, "Expected 2 but got %d", res)
}

func (suite *HashTableSuite) TestAddCollisionDifferentKey() {
	hashTable := suite.data
	hashTable.Add(1, 1)
	hashTable.Add(11, 2)
	res_1, _ := hashTable.Get(1)
	res_2, _ := hashTable.Get(11)
	// hashtable.Print()
	assert.Equal(suite.T(), 1, res_1, "Expected 1 but got %d", res_1)
	assert.Equal(suite.T(), 2, res_2, "Expected 2 but got %d", res_2)
}

func (suite *HashTableSuite) TestExistsForKeyThatExists() {
	hashtable := suite.data
	hashtable.Add(1, 1)
	isExisted := hashtable.Exists(1)
	assert.True(suite.T(), isExisted, "Expected key 1 existed but %b", isExisted)
}

func (suite *HashTableSuite) TestExistsForKeyNotExists() {
	hashtable := suite.data
	hashtable.Add(1, 1)
	isExisted := hashtable.Exists(2)
	assert.False(suite.T(), isExisted, "Expected key 2 not existed but %b", isExisted)
}

func (suite *HashTableSuite) TestRemoveExistingKey() {
	hashTable := suite.data
	hashTable.Add(1, 1)
	existed := hashTable.Exists(1)
	assert.True(suite.T(), existed)
	hashTable.Remove(1)
	existed = hashTable.Exists(1)
	assert.False(suite.T(), existed, "Expected key 1 removed successfully but got error")
}

func (suite *HashTableSuite) TestRemoveNonExistingKey() {
	hashTable := suite.data
	err := hashTable.Remove(1)
	assert.Error(suite.T(), err)
}

func (suite *HashTableSuite) TestAddAndRemoveMultipleKeys() {
	hashTable := suite.data
	hashTable.Add(1, 1)
	hashTable.Add(2, 2)
	hashTable.Add(3, 3)
	err := hashTable.Remove(2)
	assert.NoError(suite.T(), err, "Expected key 2 removed successfully but got error")
	err = hashTable.Remove(1)
	assert.NoError(suite.T(), err, "Expected key 1 removed successfully but got error")
	err = hashTable.Remove(3)
	assert.NoError(suite.T(), err, "Expected key 3 removed successfully but got error")
}

func TestHashTable(t *testing.T) {
	suite.Run(t, new(HashTableSuite))
}
