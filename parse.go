package svgpath

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

var ErrEmptyPath = errors.New("empty path")
var ErrNoPathElements = errors.New("no path elements")

var segmentRegExp = regexp.MustCompilePOSIX(`([astvzqmhlcASTVZQMHLC])([^astvzqmhlcASTVZQMHLC]*)`)
var numberRegExp = regexp.MustCompile(`-?[0-9]*\.?[0-9]+(?:e[-+]?\d+)?`)

type Segment struct {
	Command rune
	Args    []float64
}

var length = make(map[rune]int)

func init() {
	length['a'] = 7
	length['c'] = 6
	length['h'] = 1
	length['l'] = 2
	length['m'] = 2
	length['q'] = 4
	length['s'] = 4
	length['t'] = 2
	length['v'] = 1
	length['z'] = 0
}

func Parse(path string) ([]Segment, error) {
	if path == "" {
		return nil, ErrEmptyPath
	}

	segments := segmentRegExp.FindAllString(path, -1)
	if segments == nil || len(segments) == 0 {
		return nil, ErrNoPathElements
	}

	r, err := reduceE(segments, func(acc []Segment, segment string) ([]Segment, error) {
		command := rune(segment[0])
		t := unicode.ToLower(command)

		args, err2 := parseValues(segment[1:])
		if err2 != nil {
			return acc, err2
		}

		if t == 'm' || len(args) > 2 {
			acc = append(acc, Segment{Command: t, Args: args})
			t = 'l'
			if command == 'M' {
				command = 'L'
			} else if command == 'm' {
				command = 'l'
			}
		}

		for len(args) > 0 {
			if len(args) == length[t] {
				acc = append(acc, Segment{Command: command, Args: args})
				break
			}

			if len(args) < length[t] {
				return acc, fmt.Errorf("malformed path data: '%c' must have %d elements and has %d: '%s'", command, length[t], len(args), segment)
			}

			acc = append(acc, Segment{Command: command, Args: args[:length[t]]})
		}

		return acc, nil
	}, make([]Segment, 0, len(segments)))

	if err != nil {
		return nil, err
	}

	return r, nil
}

func parseValues(args string) ([]float64, error) {
	numbers := numberRegExp.FindAllString(args, -1)
	if numbers == nil || len(numbers) == 0 {
		return nil, fmt.Errorf("invalid command, no numbers found in '%s'", args)
	}

	return reduceE(numbers, func(acc []float64, number string) ([]float64, error) {
		val, err := strconv.ParseFloat(number, 64)
		if err != nil {
			return acc, err
		}

		return append(acc, val), nil
	}, make([]float64, 0, len(numbers)))
}
