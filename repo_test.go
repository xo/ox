//go:build !tinygo

package ox_test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// These tests read the repository and not the package. They hold the setup
// of D21, which every xo repository shares (dbmeta D110). The build
// constraint keeps them out of the TinyGo run, because they test files and
// not the compiler.

// TestSkillsAreCopies makes sure that each skill in skills-lock.json is an
// ordinary folder under .agents/skills and under .claude/skills, and that the
// two folders hold the same files.
//
// A symbolic link is refused because a Windows checkout writes one as a text
// file that holds the target path. Claude Code then finds a file where it
// expects a folder, and it loads no skill and reports nothing.
func TestSkillsAreCopies(t *testing.T) {
	t.Parallel()
	body, err := os.ReadFile("skills-lock.json")
	if err != nil {
		t.Fatal(err)
	}
	var lock struct {
		Skills map[string]json.RawMessage `json:"skills"`
	}
	if err := json.Unmarshal(body, &lock); err != nil {
		t.Fatalf("reading skills-lock.json: %v", err)
	}
	if len(lock.Skills) == 0 {
		t.Fatal("skills-lock.json names no skill")
	}
	roots := []string{filepath.Join(".agents", "skills"), filepath.Join(".claude", "skills")}
	for _, root := range roots {
		entries, err := os.ReadDir(root)
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range entries {
			if _, ok := lock.Skills[e.Name()]; !ok {
				t.Errorf("%s is not in skills-lock.json, so nobody can install it again. "+
					"Add it with the command in CONTRIBUTING.md", filepath.Join(root, e.Name()))
			}
		}
	}
	for _, name := range slices.Sorted(maps.Keys(lock.Skills)) {
		agents := skillFiles(t, filepath.Join(roots[0], name))
		claude := skillFiles(t, filepath.Join(roots[1], name))
		if agents == nil || claude == nil {
			continue
		}
		for _, path := range slices.Sorted(maps.Keys(agents)) {
			switch other, ok := claude[path]; {
			case !ok:
				t.Errorf("%s: %s is in %s and not in %s", name, path, roots[0], roots[1])
			case other != agents[path]:
				t.Errorf("%s: %s differs between %s and %s", name, path, roots[0], roots[1])
			}
		}
		for _, path := range slices.Sorted(maps.Keys(claude)) {
			if _, ok := agents[path]; !ok {
				t.Errorf("%s: %s is in %s and not in %s", name, path, roots[1], roots[0])
			}
		}
	}
}

// skillFiles returns the content of every file in one copy of a skill, keyed
// by its path inside the copy. It returns nil after it reports a copy that is
// missing or that is a symbolic link.
func skillFiles(t *testing.T, dir string) map[string]string {
	t.Helper()
	fi, err := os.Lstat(dir)
	switch {
	case err != nil:
		t.Errorf("%s is missing. Install the skill with the command in CONTRIBUTING.md", dir)
		return nil
	case !fi.IsDir():
		t.Errorf("%s is not a folder. A Windows checkout writes a symbolic link as a text file, "+
			"so install the skill with --copy. See D21", dir)
		return nil
	}
	out := map[string]string{}
	err = fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		switch {
		case err != nil:
			return err
		case d.Type()&fs.ModeSymlink != 0:
			t.Errorf("%s is a symbolic link. See D21", filepath.Join(dir, path))
		case d.Type().IsRegular():
			body, err := fs.ReadFile(os.DirFS(dir), path)
			if err != nil {
				return err
			}
			out[path] = string(body)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("reading %s: %v", dir, err)
	}
	return out
}

// TestClaudeImportsAgents makes sure that CLAUDE.md is an ordinary file that
// holds only @AGENTS.md. AGENTS.md holds the rules, because Codex and the
// other agents read it, and the import makes Claude Code read the same rules.
// A rule written in CLAUDE.md reaches Claude Code alone. A symbolic link is
// refused for the reason in TestSkillsAreCopies.
func TestClaudeImportsAgents(t *testing.T) {
	t.Parallel()
	fi, err := os.Lstat("CLAUDE.md")
	switch {
	case err != nil:
		t.Fatal(err)
	case !fi.Mode().IsRegular():
		t.Fatalf("CLAUDE.md is not an ordinary file, got mode %v. "+
			"Make it a file that holds @AGENTS.md. See D21", fi.Mode())
	}
	if got := strings.TrimSpace(read(t, "CLAUDE.md")); got != "@AGENTS.md" {
		t.Errorf("CLAUDE.md holds %q. It holds only @AGENTS.md, and the rules go in"+
			" AGENTS.md. See D21", got)
	}
	if _, err := os.Stat("AGENTS.md"); err != nil {
		t.Errorf("CLAUDE.md imports AGENTS.md, which is missing: %v", err)
	}
}

// TestTheRootHoldsFourDocuments makes sure that the repository root holds
// README.md, AGENTS.md, CLAUDE.md and CONTRIBUTING.md, and no other
// document. A document that appears in the root is one that nobody filed.
func TestTheRootHoldsFourDocuments(t *testing.T) {
	t.Parallel()
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	allowed := map[string]bool{
		"README.md": true, "AGENTS.md": true, "CLAUDE.md": true, "CONTRIBUTING.md": true,
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		if !allowed[e.Name()] {
			t.Errorf("%s is in the repository root. Only README, AGENTS, CLAUDE and CONTRIBUTING "+
				"belong there, and every other document goes in docs/. See D21", e.Name())
		}
	}
	for _, name := range slices.Sorted(maps.Keys(allowed)) {
		if _, err := os.Stat(name); err != nil {
			t.Errorf("expected %s in the repository root", name)
		}
	}
}

// decision is one decision in docs/PLAN.md.
type decision struct {
	num     int
	title   string
	status  string
	heading string
}

// statusPart is one part of the status at the end of a decision heading.
const statusPart = `(?:Decided|Proposed|Amends D\d+|Amended by D\d+|Supersedes D\d+|Superseded by D\d+)`

// decisionHeading matches a decision heading, such as:
//
//	### D3. The package is named ox. Decided.
//	### D9. A title. Decided. Amends D4.
var decisionHeading = regexp.MustCompile(`^### D(\d+)\. (.+?)\. (` + statusPart + `(?:\. ` + statusPart + `)*)\.$`)

// decisions reads every decision heading in docs/PLAN.md, in order.
func decisions(t *testing.T) []decision {
	t.Helper()
	var out []decision
	for line := range strings.Lines(read(t, filepath.Join("docs", "PLAN.md"))) {
		line = strings.TrimRight(line, "\n")
		if !strings.HasPrefix(line, "### D") {
			continue
		}
		m := decisionHeading.FindStringSubmatch(line)
		if m == nil {
			t.Errorf("docs/PLAN.md: %q is not a decision heading. Write it as "+
				"\"### D<n>. <title>. <status>.\", where the status is Decided, Proposed, "+
				"Amends D<n> or Amended by D<n>", line)
			continue
		}
		num, _ := strconv.Atoi(m[1])
		out = append(out, decision{
			num:     num,
			title:   m[2],
			status:  m[3],
			heading: strings.TrimPrefix(line, "### "),
		})
	}
	if len(out) == 0 {
		t.Fatal("docs/PLAN.md holds no decision heading, so the tests that read it guard nothing")
	}
	return out
}

// TestTheDecisionIndexIsComplete makes sure that each decision heading in
// docs/PLAN.md has a row in the index at the top of the file, and that the
// row has the same title and status as the heading and links to it. A reader
// finds a decision by its number in that table, so a missing row or a stale
// status hides it.
func TestTheDecisionIndexIsComplete(t *testing.T) {
	t.Parallel()
	row := regexp.MustCompile(`(?m)^\| \[D(\d+)\]\(#([^)]*)\) \| (.+) \| (.+) \|$`)
	type indexRow struct {
		anchor, title, status, line string
	}
	rows := make(map[int]indexRow)
	for _, m := range row.FindAllStringSubmatch(read(t, filepath.Join("docs", "PLAN.md")), -1) {
		num, _ := strconv.Atoi(m[1])
		if _, ok := rows[num]; ok {
			t.Errorf("docs/PLAN.md: the index has two rows for D%d", num)
		}
		rows[num] = indexRow{anchor: m[2], title: m[3], status: m[4], line: m[0]}
	}
	written := make(map[int]bool)
	for i, d := range decisions(t) {
		if d.num != i+1 {
			t.Errorf("docs/PLAN.md: the decision after D%d is D%d. Number each decision "+
				"one more than the one before it", i, d.num)
		}
		written[d.num] = true
		want := fmt.Sprintf("| [D%d](#%s) | %s | %s |", d.num, anchor(d.heading), d.title, d.status)
		switch got, ok := rows[d.num]; {
		case !ok:
			t.Errorf("D%d has no row in the index in docs/PLAN.md. Add:\n%s", d.num, want)
		case got.line != want:
			t.Errorf("D%d: the row in the index in docs/PLAN.md is\n%s\nand the heading says\n%s",
				d.num, got.line, want)
		}
	}
	for _, num := range slices.Sorted(maps.Keys(rows)) {
		if !written[num] {
			t.Errorf("docs/PLAN.md: the index has a row for D%d, and no heading holds it", num)
		}
	}
}

// anchor returns the link that GitHub gives a heading: the text in lower
// case, with each space changed to a hyphen, and with every character that is
// not a letter, a digit, a hyphen or an underscore removed.
func anchor(heading string) string {
	var sb strings.Builder
	for _, r := range strings.ToLower(heading) {
		switch {
		case r == ' ':
			_ = sb.WriteByte('-')
		case r == '-', r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			_, _ = sb.WriteRune(r)
		}
	}
	return sb.String()
}

// TestAnAmendmentPointsBothWays makes sure that when one decision amends or
// supersedes another, the status of each one names the other. A reader who
// finds the older decision must learn that it no longer holds as written.
func TestAnAmendmentPointsBothWays(t *testing.T) {
	t.Parallel()
	status := make(map[int]string)
	for _, d := range decisions(t) {
		status[d.num] = d.status
	}
	// back maps each word to the words that the other decision must use.
	back := map[string]string{
		"Amends":        "Amended by",
		"Amended by":    "Amends",
		"Supersedes":    "Superseded by",
		"Superseded by": "Supersedes",
	}
	naming := regexp.MustCompile(`(Amends|Amended by|Supersedes|Superseded by) D(\d+)`)
	for _, num := range slices.Sorted(maps.Keys(status)) {
		for _, m := range naming.FindAllStringSubmatch(status[num], -1) {
			other, _ := strconv.Atoi(m[2])
			otherStatus, ok := status[other]
			switch want := fmt.Sprintf("%s D%d", back[m[1]], num); {
			case !ok:
				t.Errorf("D%d says %q, and D%d is not a decision", num, m[0], other)
			case other == num:
				t.Errorf("D%d says %q, and a decision cannot name itself", num, m[0])
			case !strings.Contains(otherStatus, want):
				t.Errorf("D%d says %q, and the status of D%d is %q. Add %q to the heading of D%d "+
					"and to its row in the index, so that a reader of D%d learns of D%d",
					num, m[0], other, otherStatus, want, other, other, num)
			}
		}
	}
}

// read returns the content of the file at path.
func read(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
