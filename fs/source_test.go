package fs

import "testing"

func TestDefault(t *testing.T) {
	source, err := New("../testdata/devcontainer.gitignore")
	if err != nil {
		t.Fatal(err)
	}
	defer source.Close()
	if !source.HasNext() {
		t.Fatal("source should have next (1/3 line)")
	}
	line1, err := source.Next()
	if err != nil && line1 != ".cache" {
		t.Fatalf("expected .cache, got %s", line1)
	}
	line2, err := source.Next()
	if err != nil && line2 != ".dotnet" {
		t.Fatalf("expected .dotnet, got %s", line2)
	}
	if !source.HasNext() {
		t.Fatal("source should have next (3/3 line)")
	}
	line3, err := source.Next()
	if err != nil && line3 != ".vscode-server" {
		t.Fatalf("expected .vscode-server, got %s", line3)
	}
	if !source.HasNext() {
		t.Fatal("source should not have next (4/3 line)")
	}
}
