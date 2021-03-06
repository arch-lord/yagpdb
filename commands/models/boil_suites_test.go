// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverrides)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverrides)
}

func TestDelete(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesDelete)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesQueryDeleteAll)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesSliceDeleteAll)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesExists)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesExists)
}

func TestFind(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesFind)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesFind)
}

func TestBind(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesBind)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesBind)
}

func TestOne(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesOne)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesOne)
}

func TestAll(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesAll)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesAll)
}

func TestCount(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesCount)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesCount)
}

func TestInsert(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesInsert)
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesInsertWhitelist)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesInsert)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CommandsCommandOverrideToCommandsChannelsOverrideUsingCommandsChannelsOverride", testCommandsCommandOverrideToOneCommandsChannelsOverrideUsingCommandsChannelsOverride)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CommandsChannelsOverrideToCommandsCommandOverrides", testCommandsChannelsOverrideToManyCommandsCommandOverrides)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CommandsCommandOverrideToCommandsChannelsOverrideUsingCommandsCommandOverrides", testCommandsCommandOverrideToOneSetOpCommandsChannelsOverrideUsingCommandsChannelsOverride)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CommandsChannelsOverrideToCommandsCommandOverrides", testCommandsChannelsOverrideToManyAddOpCommandsCommandOverrides)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesReload)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesReloadAll)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesSelect)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesUpdate)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("CommandsChannelsOverrides", testCommandsChannelsOverridesSliceUpdateAll)
	t.Run("CommandsCommandOverrides", testCommandsCommandOverridesSliceUpdateAll)
}
