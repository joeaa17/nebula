// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPeers(t *testing.T) {
	t.Parallel()

	query := Peers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPeersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Peers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PeerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Peer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PeerExists to return true, but got false.")
	}
}

func testPeersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	peerFound, err := FindPeer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if peerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPeersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Peers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPeersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Peers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPeersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	peerOne := &Peer{}
	peerTwo := &Peer{}
	if err = randomize.Struct(seed, peerOne, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}
	if err = randomize.Struct(seed, peerTwo, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Peers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPeersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	peerOne := &Peer{}
	peerTwo := &Peer{}
	if err = randomize.Struct(seed, peerOne, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}
	if err = randomize.Struct(seed, peerTwo, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func peerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func testPeersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Peer{}
	o := &Peer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, peerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Peer object: %s", err)
	}

	AddPeerHook(boil.BeforeInsertHook, peerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	peerBeforeInsertHooks = []PeerHook{}

	AddPeerHook(boil.AfterInsertHook, peerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	peerAfterInsertHooks = []PeerHook{}

	AddPeerHook(boil.AfterSelectHook, peerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	peerAfterSelectHooks = []PeerHook{}

	AddPeerHook(boil.BeforeUpdateHook, peerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	peerBeforeUpdateHooks = []PeerHook{}

	AddPeerHook(boil.AfterUpdateHook, peerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	peerAfterUpdateHooks = []PeerHook{}

	AddPeerHook(boil.BeforeDeleteHook, peerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	peerBeforeDeleteHooks = []PeerHook{}

	AddPeerHook(boil.AfterDeleteHook, peerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	peerAfterDeleteHooks = []PeerHook{}

	AddPeerHook(boil.BeforeUpsertHook, peerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	peerBeforeUpsertHooks = []PeerHook{}

	AddPeerHook(boil.AfterUpsertHook, peerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	peerAfterUpsertHooks = []PeerHook{}
}

func testPeersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(peerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeerToManyConnections(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c Connection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, connectionDBTypes, false, connectionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, connectionDBTypes, false, connectionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PeerID = a.ID
	c.PeerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Connections().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PeerID == b.PeerID {
			bFound = true
		}
		if v.PeerID == c.PeerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadConnections(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Connections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Connections = nil
	if err = a.L.LoadConnections(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Connections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyLatencies(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c Latency

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, latencyDBTypes, false, latencyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, latencyDBTypes, false, latencyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PeerID = a.ID
	c.PeerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Latencies().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PeerID == b.PeerID {
			bFound = true
		}
		if v.PeerID == c.PeerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadLatencies(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Latencies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Latencies = nil
	if err = a.L.LoadLatencies(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Latencies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyNeighbourNeighbours(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c Neighbour

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.NeighbourID = a.ID
	c.NeighbourID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.NeighbourNeighbours().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.NeighbourID == b.NeighbourID {
			bFound = true
		}
		if v.NeighbourID == c.NeighbourID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadNeighbourNeighbours(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NeighbourNeighbours); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.NeighbourNeighbours = nil
	if err = a.L.LoadNeighbourNeighbours(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NeighbourNeighbours); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyNeighbours(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c Neighbour

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PeerID = a.ID
	c.PeerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Neighbours().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PeerID == b.PeerID {
			bFound = true
		}
		if v.PeerID == c.PeerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadNeighbours(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Neighbours); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Neighbours = nil
	if err = a.L.LoadNeighbours(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Neighbours); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManySessions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c Session

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PeerID = a.ID
	c.PeerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Sessions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PeerID == b.PeerID {
			bFound = true
		}
		if v.PeerID == c.PeerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadSessions(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sessions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Sessions = nil
	if err = a.L.LoadSessions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sessions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyAddOpConnections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e Connection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Connection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, connectionDBTypes, false, strmangle.SetComplement(connectionPrimaryKeyColumns, connectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Connection{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddConnections(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PeerID {
			t.Error("foreign key was wrong value", a.ID, first.PeerID)
		}
		if a.ID != second.PeerID {
			t.Error("foreign key was wrong value", a.ID, second.PeerID)
		}

		if first.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Connections[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Connections[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Connections().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpLatencies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e Latency

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Latency{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, latencyDBTypes, false, strmangle.SetComplement(latencyPrimaryKeyColumns, latencyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Latency{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLatencies(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PeerID {
			t.Error("foreign key was wrong value", a.ID, first.PeerID)
		}
		if a.ID != second.PeerID {
			t.Error("foreign key was wrong value", a.ID, second.PeerID)
		}

		if first.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Latencies[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Latencies[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Latencies().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpNeighbourNeighbours(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e Neighbour

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Neighbour{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, neighbourDBTypes, false, strmangle.SetComplement(neighbourPrimaryKeyColumns, neighbourColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Neighbour{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNeighbourNeighbours(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.NeighbourID {
			t.Error("foreign key was wrong value", a.ID, first.NeighbourID)
		}
		if a.ID != second.NeighbourID {
			t.Error("foreign key was wrong value", a.ID, second.NeighbourID)
		}

		if first.R.Neighbour != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Neighbour != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.NeighbourNeighbours[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.NeighbourNeighbours[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.NeighbourNeighbours().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpNeighbours(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e Neighbour

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Neighbour{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, neighbourDBTypes, false, strmangle.SetComplement(neighbourPrimaryKeyColumns, neighbourColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Neighbour{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNeighbours(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PeerID {
			t.Error("foreign key was wrong value", a.ID, first.PeerID)
		}
		if a.ID != second.PeerID {
			t.Error("foreign key was wrong value", a.ID, second.PeerID)
		}

		if first.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Neighbours[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Neighbours[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Neighbours().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpSessions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e Session

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Session{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sessionDBTypes, false, strmangle.SetComplement(sessionPrimaryKeyColumns, sessionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Session{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSessions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PeerID {
			t.Error("foreign key was wrong value", a.ID, first.PeerID)
		}
		if a.ID != second.PeerID {
			t.Error("foreign key was wrong value", a.ID, second.PeerID)
		}

		if first.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Sessions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Sessions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Sessions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPeersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Peers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	peerDBTypes = map[string]string{`PeerID`: `character varying`, `MultiAddresses`: `ARRAYcharacter varying`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`, `OldMultiAddresses`: `ARRAYcharacter varying`, `AgentVersion`: `character varying`, `Protocol`: `ARRAYcharacter varying`, `ID`: `integer`}
	_           = bytes.MinRead
)

func testPeersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(peerAllColumns) == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peerDBTypes, true, peerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPeersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(peerAllColumns) == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peerDBTypes, true, peerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(peerAllColumns, peerPrimaryKeyColumns) {
		fields = peerAllColumns
	} else {
		fields = strmangle.SetComplement(
			peerAllColumns,
			peerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PeerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPeersUpsert(t *testing.T) {
	t.Parallel()

	if len(peerAllColumns) == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Peer{}
	if err = randomize.Struct(seed, &o, peerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Peer: %s", err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, peerDBTypes, false, peerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Peer: %s", err)
	}

	count, err = Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
