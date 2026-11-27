package main

import (
	"slices"
	"strings"
	"testing"
)

func TestCobraGh(t *testing.T) {
	root := importTestdata(t, "gh", "pr", "list")
	if root.Format != FormatCobra {
		t.Fatalf("expected format %q, got: %q", FormatCobra, root.Format)
	}
	exp := []string{"CORE COMMANDS", "GITHUB ACTIONS COMMANDS", "ALIAS COMMANDS", "ADDITIONAL COMMANDS"}
	if !slices.Equal(root.CommandSections, exp) {
		t.Errorf("expected command sections %q, got: %q", exp, root.CommandSections)
	}
	if root.Command("help") != nil || root.Command("completion") != nil {
		t.Errorf("expected no help and no completion command")
	}
	// co is a command that only __complete lists
	if c := root.Command("co"); c == nil || c.Usage != `Alias for "pr checkout"` {
		t.Errorf("expected the command co from __complete, got: %v", c)
	}
	pr := root.Command("pr")
	if pr == nil || pr.Stub {
		t.Fatalf("expected gh pr to be read")
	}
	if c := root.Command("issue"); c == nil || !c.Stub {
		t.Errorf("expected gh issue to be a stub")
	}
	if exp := []string{"GENERAL COMMANDS", "TARGETED COMMANDS"}; !slices.Equal(pr.CommandSections, exp) {
		t.Errorf("expected command sections %q, got: %q", exp, pr.CommandSections)
	}
	list := pr.Command("list")
	if list == nil || list.Stub {
		t.Fatalf("expected gh pr list to be read")
	}
	if exp := []string{"ls"}; !slices.Equal(list.Aliases, exp) {
		t.Errorf("expected aliases %q, got: %q", exp, list.Aliases)
	}
	if list.Spec != "[flags]" {
		t.Errorf("expected spec [flags], got: %q", list.Spec)
	}
	checkFlags(t, list, []testFlag{
		{"app", "", "string", "", "", -1, nil},
		{"draft", "d", "bool", "", "", -1, nil},
		{"jq", "q", "string", "expression", "", -1, nil},
		{"label", "l", "slice", "", "", -1, nil},
		{"limit", "L", "int", "", "30", -1, nil},
		{"state", "s", "string", "", "open", -1, nil},
		{"repo", "R", "string", "[HOST/]OWNER/REPO", "", 0, nil},
	})
	if g := list.Flag("state"); g != nil {
		if exp := []string{"open", "closed", "merged", "all"}; !slices.Equal(g.Values, exp) || !g.Exact {
			t.Errorf("expected exact values %q, got: %q %t", exp, g.Values, g.Exact)
		}
		if exp := "Filter by state: {open|closed|merged|all}"; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
	if g := list.Flag("json"); g != nil && (g.Exact || !slices.Contains(g.Values, "additions")) {
		t.Errorf("expected values that are not exact, got: %q %t", g.Values, g.Exact)
	}
	if g := list.Flag("repo"); g != nil && !g.Inherited {
		t.Errorf("expected --repo to be inherited")
	}
	if !strings.HasPrefix(list.Footer, "JSON FIELDS\n  additions,") {
		t.Errorf("expected footer, got: %q", list.Footer)
	}
}

func TestCobraHelm(t *testing.T) {
	root := importTestdata(t, "helm", "list")
	checkFlags(t, root, []testFlag{
		{"burst-limit", "", "int", "", "100", -1, nil},
		{"color", "", "string", "", "auto", -1, nil},
		{"content-cache", "", "string", "", "$HOME/.cache/helm/content", -1, nil},
		{"kube-as-group", "", "array", "", "", -1, nil},
		{"qps", "", "float32", "", "", -1, nil},
	})
	list := root.Command("list")
	if list == nil || list.Stub {
		t.Fatalf("expected helm list to be read")
	}
	if exp := []string{"ls"}; !slices.Equal(list.Aliases, exp) {
		t.Errorf("expected aliases %q, got: %q", exp, list.Aliases)
	}
	checkFlags(t, list, []testFlag{
		{"all-namespaces", "A", "bool", "", "", -1, nil},
		{"max", "m", "int", "", "256", -1, nil},
		{"output", "o", "string", "format", "table", -1, nil},
	})
	if g := list.Flag("output"); g != nil {
		if exp := []string{"json", "table", "yaml"}; !slices.Equal(g.Values, exp) || !g.Exact {
			t.Errorf("expected exact values %q, got: %q %t", exp, g.Values, g.Exact)
		}
	}
	if g := list.Flag("burst-limit"); g == nil || !g.Inherited {
		t.Errorf("expected --burst-limit to be inherited")
	}
}

func TestCobraKubectl(t *testing.T) {
	root := importTestdata(t, "kubectl", "get")
	// the help of kubectl lists no flags, so they come from __complete
	checkFlags(t, root, []testFlag{
		{"namespace", "n", "string", "", "", -1, nil},
		{"disable-compression", "", "bool", "", "", -1, nil},
	})
	if g := root.Flag("help"); g == nil || g.Hook != "help" {
		t.Errorf("expected a help hook")
	}
	if !slices.Contains(root.CommandSections, "Basic Commands (Beginner)") {
		t.Errorf("expected command sections, got: %q", root.CommandSections)
	}
	get := root.Command("get")
	if get == nil || get.Stub {
		t.Fatalf("expected kubectl get to be read")
	}
	checkFlags(t, get, []testFlag{
		{"all-namespaces", "A", "bool", "", "", -1, nil},
		{"allow-missing-template-keys", "", "bool", "", "true", -1, nil},
		{"chunk-size", "", "int", "", "500", -1, nil},
		{"filename", "f", "slice", "", "", -1, nil},
		{"output", "o", "string", "", "", -1, nil},
	})
	if g := get.Flag("filename"); g != nil && (g.Path != "file" || !slices.Contains(g.Exts, "yaml")) {
		t.Errorf("expected a file with extensions, got: %q %q", g.Path, g.Exts)
	}
	if g := get.Flag("all-namespaces"); g != nil && !strings.HasPrefix(g.Usage, "If present, list the requested object(s)") {
		t.Errorf("expected usage, got: %q", g.Usage)
	}
}

func TestSplitSections(t *testing.T) {
	help := "Banner line.\n\nUsage:\n  app [flags]\n\nSee:\n<https://example.com>\n\nFLAGS\n  -a, --all   all\n\nUse \"app --help\".\n"
	pre, sections, post := splitSections(splitLines(help))
	if exp := []string{"Banner line.", ""}; !slices.Equal(pre, exp) {
		t.Errorf("expected pre %q, got: %q", exp, pre)
	}
	var titles []string
	for _, sect := range sections {
		titles = append(titles, sect.title)
	}
	// See: is not a title, because the line after it is not indented
	if exp := []string{"Usage", "FLAGS"}; !slices.Equal(titles, exp) {
		t.Errorf("expected titles %q, got: %q", exp, titles)
	}
	if exp := []string{`Use "app --help".`}; !slices.Equal(post, exp) {
		t.Errorf("expected post %q, got: %q", exp, post)
	}
}
