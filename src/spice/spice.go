package spice

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	maxAmoutnPerSupplementaryCurrency = 1000000000000000000
)

const (
	Currency = iota
	SuplementaryCurrency
)

var (
	ErrValueOverflow      = errors.New("value overflow")
	ErrNoSufficientFounds = errors.New("no sufficient founds to process transaction")
)

// Melange is an asset that is digitally transferable between two wallets.
type Melange struct {
	Currency              uint64 `yaml:"currency"`
	SupplementaryCurrency uint64 `yaml:"supplementary_currency"`
}

// New creates new spice Melange from given currency and supplementary currency values.
func New(currency, supplementaryCurrency uint64) Melange {
	if supplementaryCurrency >= maxAmoutnPerSupplementaryCurrency {
		currency += 1
		supplementaryCurrency -= maxAmoutnPerSupplementaryCurrency
	}
	return Melange{
		Currency:              currency,
		SupplementaryCurrency: supplementaryCurrency,
	}
}

// Supply supplies spice of the given amount from the source to the entity.
func (m *Melange) Supply(amount Melange) error {
	mCp := m.clone()
	for _, unit := range []byte{Currency, SuplementaryCurrency} {
		switch unit {
		case Currency:
			if math.MaxUint64-amount.Currency < m.Currency {
				return ErrValueOverflow
			}
			m.Currency += amount.Currency
		case SuplementaryCurrency:
			if maxAmoutnPerSupplementaryCurrency-amount.SupplementaryCurrency < m.SupplementaryCurrency {
				if m.Currency == math.MaxUint64 {
					m.copyFrom(mCp)
					return ErrValueOverflow
				}
			}
			m.SupplementaryCurrency += amount.SupplementaryCurrency

			if m.SupplementaryCurrency >= maxAmoutnPerSupplementaryCurrency {
				m.Currency += 1
				m.SupplementaryCurrency -= maxAmoutnPerSupplementaryCurrency
			}
		}
	}
	return nil
}

// Drain drains amount from the function pointer receiver to the sink.
func (m *Melange) Drain(amount Melange, sink *Melange) error {
	return Transfer(amount, m, sink)
}

// Empty verifies if is spice empty.
func (m *Melange) Empty() bool {
	return m.Currency == 0 && m.SupplementaryCurrency == 0
}

// Transfer transfers given amount from one Melange asset to the other if possible or returns error otherwise.
func Transfer(amount Melange, from, to *Melange) error {
	toCp := to.clone()
	fromCp := from.clone()
	for _, unit := range []byte{Currency, SuplementaryCurrency} {
		switch unit {
		case Currency:
			if amount.Currency > from.Currency {
				return ErrNoSufficientFounds
			}
			if math.MaxUint64-amount.Currency < to.Currency {
				return ErrValueOverflow
			}
			to.Currency += amount.Currency
			from.Currency -= amount.Currency
		case SuplementaryCurrency:
			if maxAmoutnPerSupplementaryCurrency-amount.SupplementaryCurrency < to.SupplementaryCurrency {
				if to.Currency == math.MaxUint64 {
					to.copyFrom(toCp)
					from.copyFrom(fromCp)
					return ErrValueOverflow
				}
			}
			if amount.SupplementaryCurrency > from.SupplementaryCurrency {
				if from.Currency == 0 {
					to.copyFrom(toCp)
					from.copyFrom(fromCp)
					return ErrNoSufficientFounds
				}
				from.Currency -= 1
				from.SupplementaryCurrency = from.SupplementaryCurrency + maxAmoutnPerSupplementaryCurrency - amount.SupplementaryCurrency
				to.SupplementaryCurrency += amount.SupplementaryCurrency

				if to.SupplementaryCurrency >= maxAmoutnPerSupplementaryCurrency {
					to.Currency += 1
					to.SupplementaryCurrency -= maxAmoutnPerSupplementaryCurrency
				}
				continue
			}
			from.SupplementaryCurrency -= amount.SupplementaryCurrency
			to.SupplementaryCurrency += amount.SupplementaryCurrency

			if to.SupplementaryCurrency >= maxAmoutnPerSupplementaryCurrency {
				to.Currency += 1
				to.SupplementaryCurrency -= maxAmoutnPerSupplementaryCurrency
			}
		}
	}
	return nil
}

// String returns string representation of spice Melange.
func (m Melange) String() string {
	suplementary := fmt.Sprintf("%v", m.SupplementaryCurrency)
	zeros := 18 - len(suplementary)
	if zeros < 0 {
		suplementary = "0"
	}
	suplementary = strings.Trim(suplementary, "0")

	var buf strings.Builder
	if len(suplementary) != 0 {
		for i := 0; i < zeros; i++ {
			buf.WriteString("0")
		}
	}
	buf.WriteString(suplementary)
	supp := buf.String()
	if len(supp) == 0 {
		supp = "0"
	}
	curr := fmt.Sprintf("%v", m.Currency)
	return fmt.Sprintf("%s.%s", curr, supp)
}

func (m Melange) clone() Melange {
	return Melange{
		Currency:              m.Currency,
		SupplementaryCurrency: m.SupplementaryCurrency,
	}
}

func (m *Melange) copyFrom(c Melange) {
	m.Currency = c.Currency
	m.SupplementaryCurrency = c.SupplementaryCurrency
}