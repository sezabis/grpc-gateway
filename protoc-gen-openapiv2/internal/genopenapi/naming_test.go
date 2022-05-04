package genopenapi

import "testing"

func TestNaming(t *testing.T) {
	type expectedNames struct {
		fqn, legacy, simple, noSep string
	}
	messageNameToExpected := map[string]expectedNames{
		".A":       {"A", "A", "A", "A"},
		".a.B.C":   {"a.B.C", "aBC", "B.C", "aBC"},
		".a.D.C":   {"a.D.C", "aDC", "D.C", "aDC"},
		".a.E.F":   {"a.E.F", "aEF", "a.E.F", "aEF"},
		".b.E.F":   {"b.E.F", "bEF", "b.E.F", "bEF"},
		".c.G.H.D": {"c.G.H.D", "HD", "D", "cGHD"},
		".a.c.G.H": {"a.c.G.H", "acGH", "a.c.G.H", "acGH"},
		".b.c.G.H": {"b.c.G.H", "bcGH", "b.c.G.H", "bcGH"},
	}

	allMessageNames := make([]string, 0, len(messageNameToExpected))
	for msgName := range messageNameToExpected {
		allMessageNames = append(allMessageNames, msgName)
	}

	t.Run("fqn", func(t *testing.T) {
		uniqueNames := resolveNamesFQN(allMessageNames)
		for _, msgName := range allMessageNames {
			expected := messageNameToExpected[msgName].fqn
			actual := uniqueNames[msgName]
			if expected != actual {
				t.Errorf("fqn unique name %q does not match expected name %q, source %q", actual, expected, msgName)
			}
		}
	})
	t.Run("legacy", func(t *testing.T) {
		uniqueNames := resolveNamesLegacy(allMessageNames)
		for _, msgName := range allMessageNames {
			expected := messageNameToExpected[msgName].legacy
			actual := uniqueNames[msgName]
			if expected != actual {
				t.Errorf("legacy unique name %q does not match expected name %q, source %q", actual, expected, msgName)
			}
		}
	})
	t.Run("simple", func(t *testing.T) {
		uniqueNames := resolveNamesSimple(allMessageNames)
		for _, msgName := range allMessageNames {
			expected := messageNameToExpected[msgName].simple
			actual := uniqueNames[msgName]
			if expected != actual {
				t.Errorf("simple unique name %q does not match expected name %q, source %q", actual, expected, msgName)
			}
		}
	})
	t.Run("full", func(t *testing.T) {
		uniqueNames := resolveNamesNoSeparator(allMessageNames)
		for _, msgName := range allMessageNames {
			expected := messageNameToExpected[msgName].noSep
			actual := uniqueNames[msgName]
			if expected != actual {
				t.Errorf("no-sep unique name %q does not match expected name %q, source %q", actual, expected, msgName)
			}
		}
	})
}
