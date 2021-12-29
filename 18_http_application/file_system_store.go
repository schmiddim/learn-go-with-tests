package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {

			wins = player.Wins
			break
		}
	}
	return wins
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
