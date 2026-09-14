package ox

import "testing"

func TestUserStateDir(t *testing.T) {
	dir, err := UserStateDir()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	t.Logf("dir: %s", dir)
}
