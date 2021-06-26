// Package language brings up a bridge among client, backend and database
// in this package all the ties between language and its interpretions should be handled
package language

// ID language id
type ID int32

const (
	// Farsi fa-IR , fa
	Farsi ID = 1 + iota
	// English en-US, en
	English
)

// Symbol signifies language symbolic representations, for instance fa-IR, en-US etc.
type Symbol string

const (
	// SymFarsi farsi symbol
	SymFarsi Symbol = "fa-IR"
	// SymEnglish english symbol
	SymEnglish Symbol = "en-US"
)

var (
	// IDToSym which converts languages
	IDToSym = map[int32]string{
		int32(English): string(SymEnglish),
		int32(Farsi):   string(SymFarsi),
	}
	// SymToID sym to id
	SymToID = map[Symbol]ID{
		SymEnglish: English,
		SymFarsi:   Farsi,
	}
)

// ToString convert symbol to string
func (sym Symbol) ToString() string {
	return string(sym)
}

// ToInt convert ID to int
func (id ID) ToInt() int {
	return int(id)
}

// ToID convert symbol to ID
func (sym Symbol) ToID() ID {
	// switch cases are faster than maps
	switch sym {
	case SymFarsi:
		return Farsi
	case SymEnglish:
		return English
	default:
		return Farsi
	}
}

// ToSymbol convert ID to Symbol
func (id ID) ToSymbol() Symbol {
	switch id {
	case Farsi:
		return SymFarsi
	case English:
		return SymEnglish
	default:
		return SymFarsi
	}
}
