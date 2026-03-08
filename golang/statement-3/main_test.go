package main

import "testing"

func TestAddAndExists(t *testing.T) {
	store := NewTrie()
	store.Add("hello")
	if !store.Exists("hello") {
		t.Errorf("word 'hello' should exist after adding")
	}
}

func TestWordNotFound(t *testing.T) {
	dict := NewTrie()
	if dict.Exists("world") {
		t.Errorf("word 'world' should not exist")
	}
}

func TestDelete(t *testing.T) {
	t1 := NewTrie()
	t1.Add("golang")
	t1.Delete("golang")
	if t1.Exists("golang") {
		t.Errorf("word 'golang' should have been removed")
	}
}

func TestEmptyInput(t *testing.T) {
	store := NewTrie()
	store.Add("")
	if store.Exists("") {
		t.Errorf("empty string should not be stored")
	}
}
func TestDeleteNonExisting(t *testing.T) {
	dict := NewTrie()
	dict.Delete("random")
	if dict.Exists("random") {
		t.Errorf("random should not exist in trie")
	}
}
