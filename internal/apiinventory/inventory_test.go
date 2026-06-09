package apiinventory

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateFromFixtureHeaders(t *testing.T) {
	dir := t.TempDir()
	writeHeader(t, dir, "libraw.h", `
extern "C" {
  DllDef const char *libraw_version(void);
  DllDef libraw_data_t *libraw_init(unsigned int flags);
  DllDef int libraw_open_file(libraw_data_t *, const char *);
}
`)
	writeHeader(t, dir, "libraw_const.h", `
#define LIBRAW_SAMPLE_MACRO 1
enum LibRaw_errors { LIBRAW_SUCCESS = 0 };
`)
	writeHeader(t, dir, "libraw_types.h", `
typedef struct { int raw_height; } libraw_image_sizes_t;
`)
	writeHeader(t, dir, "libraw_version.h", `
#define LIBRAW_MAJOR_VERSION 1
#define LIBRAW_MINOR_VERSION 2
#define LIBRAW_PATCH_VERSION 3
#define LIBRAW_VERSION_TAIL Release
#define LIBRAW_VERSION_STR "1.2.3-Release"
`)

	inv, err := Generate(Options{HeaderDir: dir})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if inv.Version != "1.2.3-Release" {
		t.Fatalf("Version = %q, want 1.2.3-Release", inv.Version)
	}

	want := map[string]bool{
		"function\tlibraw_version":      false,
		"function\tlibraw_init":         false,
		"function\tlibraw_open_file":    false,
		"enum\tLibRaw_errors":           false,
		"macro\tLIBRAW_SAMPLE_MACRO":    false,
		"struct\tlibraw_image_sizes_t":  false,
		"version\tLIBRAW_MAJOR_VERSION": false,
		"version\tLIBRAW_VERSION_STR":   false,
		"version\tLIBRAW_VERSION_TAIL":  false,
		"version\tLIBRAW_PATCH_VERSION": false,
		"version\tLIBRAW_MINOR_VERSION": false,
	}
	for _, sym := range inv.Symbols {
		key := string(sym.Kind) + "\t" + sym.Name
		if _, ok := want[key]; ok {
			want[key] = true
		}
	}
	for key, seen := range want {
		if !seen {
			t.Fatalf("missing symbol %s in inventory %#v", key, inv.Symbols)
		}
	}
}

func TestLoadCoverageAndMissingCoverage(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "coverage.tsv")
	if err := os.WriteFile(path, []byte("# kind\tname\tstatus\tnote\nfunction\tlibraw_version\twrapped\tversion helper\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	coverage, err := LoadCoverage(path)
	if err != nil {
		t.Fatalf("LoadCoverage() error = %v", err)
	}

	inv := Inventory{Symbols: []Symbol{
		{Kind: FunctionKind, Name: "libraw_version", Header: "libraw.h"},
		{Kind: FunctionKind, Name: "libraw_open_file", Header: "libraw.h"},
	}}
	missing := MissingCoverage(inv, coverage)
	if len(missing) != 1 || missing[0].Name != "libraw_open_file" {
		t.Fatalf("missing = %#v, want libraw_open_file", missing)
	}
}

func TestRenderMarkdownReportsUnmapped(t *testing.T) {
	inv := Inventory{
		HeaderDir: "/tmp/libraw",
		Version:   "1.2.3-Release",
		Symbols: []Symbol{
			{Kind: FunctionKind, Name: "libraw_version", Header: "libraw.h"},
		},
	}

	body, missing := RenderMarkdown(inv, map[string]CoverageEntry{})
	if len(missing) != 1 {
		t.Fatalf("missing len = %d, want 1", len(missing))
	}
	if !strings.Contains(string(body), "`unmapped`") {
		t.Fatalf("rendered markdown does not mention unmapped:\n%s", body)
	}
}

func TestRenderCoverageTSVPreservesExistingEntries(t *testing.T) {
	inv := Inventory{Symbols: []Symbol{
		{Kind: FunctionKind, Name: "libraw_version", Header: "libraw.h"},
		{Kind: FunctionKind, Name: "libraw_open_file", Header: "libraw.h"},
	}}
	coverage := map[string]CoverageEntry{
		"function\tlibraw_version": {
			Kind:   FunctionKind,
			Name:   "libraw_version",
			Status: "wrapped",
			Note:   "public version helper",
		},
	}

	body := string(RenderCoverageTSV(inv, coverage))
	if !strings.Contains(body, "function\tlibraw_version\twrapped\tpublic version helper") {
		t.Fatalf("existing coverage entry not preserved:\n%s", body)
	}
	if !strings.Contains(body, "function\tlibraw_open_file\tdeferred\ttracked for a future workflowr task") {
		t.Fatalf("missing default deferred entry:\n%s", body)
	}
}

func TestRenderCoverageReportSummarizesReleaseGate(t *testing.T) {
	inv := Inventory{
		HeaderDir: "/tmp/libraw",
		Version:   "1.2.3-Release",
		Symbols: []Symbol{
			{Kind: FunctionKind, Name: "libraw_version", Header: "libraw.h"},
			{Kind: FunctionKind, Name: "libraw_open_wfile", Header: "libraw.h"},
			{Kind: MacroKind, Name: "LIBRAW_WIN32_CALLS", Header: "libraw.h"},
		},
	}
	coverage := map[string]CoverageEntry{
		"function\tlibraw_version": {
			Kind:   FunctionKind,
			Name:   "libraw_version",
			Status: "wrapped",
			Note:   "public version helper",
		},
		"function\tlibraw_open_wfile": {
			Kind:   FunctionKind,
			Name:   "libraw_open_wfile",
			Status: "deferred",
			Note:   "tracked for Windows path support",
		},
		"macro\tLIBRAW_WIN32_CALLS": {
			Kind:   MacroKind,
			Name:   "LIBRAW_WIN32_CALLS",
			Status: "unsupported",
			Note:   "platform macro",
		},
	}

	body := string(RenderCoverageReport(inv, coverage))
	for _, want := range []string{
		"- Release gate: `pass`",
		"| `wrapped` | `1` |",
		"| `deferred` | `1` |",
		"| `unsupported` | `1` |",
		"`function` `libraw_open_wfile`",
		"`macro` `LIBRAW_WIN32_CALLS`",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("coverage report missing %q:\n%s", want, body)
		}
	}
}

func TestCoverageReportFailsGateForUnmappedSymbols(t *testing.T) {
	inv := Inventory{Symbols: []Symbol{
		{Kind: FunctionKind, Name: "libraw_new_symbol", Header: "libraw.h"},
	}}

	body := string(RenderCoverageReport(inv, map[string]CoverageEntry{}))
	if !strings.Contains(body, "- Release gate: `fail`") {
		t.Fatalf("coverage report did not fail gate for unmapped symbol:\n%s", body)
	}
}

func TestGenerateFromInstalledHeaders(t *testing.T) {
	dir, err := FindHeaderDir("")
	if err != nil {
		t.Skipf("installed LibRaw headers not available: %v", err)
	}

	inv, err := Generate(Options{HeaderDir: dir})
	if err != nil {
		t.Fatalf("Generate() with installed headers error = %v", err)
	}
	if len(inv.Symbols) == 0 {
		t.Fatal("installed header inventory is empty")
	}
	if inv.Version == "" {
		t.Fatal("installed header version is empty")
	}
}

func writeHeader(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}
