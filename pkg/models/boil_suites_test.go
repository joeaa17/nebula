// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Crawls", testCrawls)
	t.Run("Latencies", testLatencies)
	t.Run("PeerProperties", testPeerProperties)
	t.Run("Peers", testPeers)
	t.Run("Sessions", testSessions)
}

func TestDelete(t *testing.T) {
	t.Run("Crawls", testCrawlsDelete)
	t.Run("Latencies", testLatenciesDelete)
	t.Run("PeerProperties", testPeerPropertiesDelete)
	t.Run("Peers", testPeersDelete)
	t.Run("Sessions", testSessionsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Crawls", testCrawlsQueryDeleteAll)
	t.Run("Latencies", testLatenciesQueryDeleteAll)
	t.Run("PeerProperties", testPeerPropertiesQueryDeleteAll)
	t.Run("Peers", testPeersQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Crawls", testCrawlsSliceDeleteAll)
	t.Run("Latencies", testLatenciesSliceDeleteAll)
	t.Run("PeerProperties", testPeerPropertiesSliceDeleteAll)
	t.Run("Peers", testPeersSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Crawls", testCrawlsExists)
	t.Run("Latencies", testLatenciesExists)
	t.Run("PeerProperties", testPeerPropertiesExists)
	t.Run("Peers", testPeersExists)
	t.Run("Sessions", testSessionsExists)
}

func TestFind(t *testing.T) {
	t.Run("Crawls", testCrawlsFind)
	t.Run("Latencies", testLatenciesFind)
	t.Run("PeerProperties", testPeerPropertiesFind)
	t.Run("Peers", testPeersFind)
	t.Run("Sessions", testSessionsFind)
}

func TestBind(t *testing.T) {
	t.Run("Crawls", testCrawlsBind)
	t.Run("Latencies", testLatenciesBind)
	t.Run("PeerProperties", testPeerPropertiesBind)
	t.Run("Peers", testPeersBind)
	t.Run("Sessions", testSessionsBind)
}

func TestOne(t *testing.T) {
	t.Run("Crawls", testCrawlsOne)
	t.Run("Latencies", testLatenciesOne)
	t.Run("PeerProperties", testPeerPropertiesOne)
	t.Run("Peers", testPeersOne)
	t.Run("Sessions", testSessionsOne)
}

func TestAll(t *testing.T) {
	t.Run("Crawls", testCrawlsAll)
	t.Run("Latencies", testLatenciesAll)
	t.Run("PeerProperties", testPeerPropertiesAll)
	t.Run("Peers", testPeersAll)
	t.Run("Sessions", testSessionsAll)
}

func TestCount(t *testing.T) {
	t.Run("Crawls", testCrawlsCount)
	t.Run("Latencies", testLatenciesCount)
	t.Run("PeerProperties", testPeerPropertiesCount)
	t.Run("Peers", testPeersCount)
	t.Run("Sessions", testSessionsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Crawls", testCrawlsHooks)
	t.Run("Latencies", testLatenciesHooks)
	t.Run("PeerProperties", testPeerPropertiesHooks)
	t.Run("Peers", testPeersHooks)
	t.Run("Sessions", testSessionsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Crawls", testCrawlsInsert)
	t.Run("Crawls", testCrawlsInsertWhitelist)
	t.Run("Latencies", testLatenciesInsert)
	t.Run("Latencies", testLatenciesInsertWhitelist)
	t.Run("PeerProperties", testPeerPropertiesInsert)
	t.Run("PeerProperties", testPeerPropertiesInsertWhitelist)
	t.Run("Peers", testPeersInsert)
	t.Run("Peers", testPeersInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("LatencyToPeerUsingPeer", testLatencyToOnePeerUsingPeer)
	t.Run("PeerPropertyToCrawlUsingCrawl", testPeerPropertyToOneCrawlUsingCrawl)
	t.Run("SessionToPeerUsingPeer", testSessionToOnePeerUsingPeer)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CrawlToPeerProperties", testCrawlToManyPeerProperties)
	t.Run("PeerToLatencies", testPeerToManyLatencies)
	t.Run("PeerToSessions", testPeerToManySessions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("LatencyToPeerUsingLatencies", testLatencyToOneSetOpPeerUsingPeer)
	t.Run("PeerPropertyToCrawlUsingPeerProperties", testPeerPropertyToOneSetOpCrawlUsingCrawl)
	t.Run("SessionToPeerUsingSessions", testSessionToOneSetOpPeerUsingPeer)
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
	t.Run("CrawlToPeerProperties", testCrawlToManyAddOpPeerProperties)
	t.Run("PeerToLatencies", testPeerToManyAddOpLatencies)
	t.Run("PeerToSessions", testPeerToManyAddOpSessions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Crawls", testCrawlsReload)
	t.Run("Latencies", testLatenciesReload)
	t.Run("PeerProperties", testPeerPropertiesReload)
	t.Run("Peers", testPeersReload)
	t.Run("Sessions", testSessionsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Crawls", testCrawlsReloadAll)
	t.Run("Latencies", testLatenciesReloadAll)
	t.Run("PeerProperties", testPeerPropertiesReloadAll)
	t.Run("Peers", testPeersReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Crawls", testCrawlsSelect)
	t.Run("Latencies", testLatenciesSelect)
	t.Run("PeerProperties", testPeerPropertiesSelect)
	t.Run("Peers", testPeersSelect)
	t.Run("Sessions", testSessionsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Crawls", testCrawlsUpdate)
	t.Run("Latencies", testLatenciesUpdate)
	t.Run("PeerProperties", testPeerPropertiesUpdate)
	t.Run("Peers", testPeersUpdate)
	t.Run("Sessions", testSessionsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Crawls", testCrawlsSliceUpdateAll)
	t.Run("Latencies", testLatenciesSliceUpdateAll)
	t.Run("PeerProperties", testPeerPropertiesSliceUpdateAll)
	t.Run("Peers", testPeersSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
}
