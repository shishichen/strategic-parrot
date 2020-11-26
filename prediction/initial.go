package prediction

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/shishichen/strategic-parrot/base"
)

const initialOutcomeFile = "data/initial_outcomes"

// GetInitialOutcomes returns, given a hole and empty board, the probability that the hole will win, tie, and lose at
// the end of the game, assumming random subsequent cards. i.e. the starting hand probabilities, which are always the
// same.
// TODO: add number of opponents
func GetInitialOutcomes(hole []base.Card) (float64, float64, float64, error) {
	if len(hole) != 2 {
		return 0, 0, 0, fmt.Errorf("initial outcomes can only be predicted for holes with 2 cards")
	}

	outcomes, err := readInitialOutcomes()
	if err != nil {
		return 0, 0, 0, err
	}

	key, _ := base.GetKey(hole)
	outcome, ok := outcomes[key]
	if !ok {
		return 0, 0, 0, fmt.Errorf("initial outcome for hole %v not found", hole)
	}

	return outcome.win, outcome.tie, outcome.lose, nil
}

// PrecomputeInitialOutcomes precomputes the initial outcomes and stores them in a file for GetInitialOutcomes to later use.
func PrecomputeInitialOutcomes() error {
	deck := base.NewDeck()
	boards := base.GetCombinations(deck.GetCards(), 5)

	log.Printf("num boards %v", len(boards))

	n := runtime.NumCPU()
	var wg sync.WaitGroup
	accumulations := make([]map[base.Key]accumulation, n)
	for id := 0; id < n; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			accumulations[id] = make(map[base.Key]accumulation)

			lower := id * len(boards) / n
			upper := (id + 1) * len(boards) / n
			log.Printf("thread %v starting work on boards %v to %v", id, lower, upper)
			for i := lower; i < upper; i++ {
				deck := base.NewDeck()
				deck.Remove(boards[i])
				levels := evaluate(deck.GetCards(), boards[i])

				// For every hole...
				for _, holes := range levels {
					for _, hole := range holes.getHoles() {
						key, _ := base.GetKey(hole)
						// Count coexisting holes by reiterating through the levels
						for _, level := range levels {
							size := countCoexisting(hole, level.getHoles())
							a := accumulations[id][key]
							a.total += size
							if level.getScore() > holes.getScore() {
								a.better += size
							} else if level.getScore() == holes.getScore() {
								a.same += size
							} else {
								a.worse += size
							}
							accumulations[id][key] = a
						}
					}
				}

				if ((i+1)-lower)%10000 == 0 {
					log.Printf("thread %v completed %v out of %v boards", id, (i+1)-lower, upper-lower)
				}
			}
		}(id)
	}
	wg.Wait()

	accumulationsTotal := make(map[base.Key]accumulation)
	for id := 0; id < n; id++ {
		for k, v := range accumulations[id] {
			a := accumulationsTotal[k]
			a.better += v.better
			a.same += v.same
			a.worse += v.worse
			a.total += v.total
			accumulationsTotal[k] = a
		}
	}
	outcomes := make(map[base.Key]outcome)
	for key, accumulation := range accumulationsTotal {
		outcomes[key] = outcome{
			float64(accumulation.worse) / float64(accumulation.total),
			float64(accumulation.same) / float64(accumulation.total),
			float64(accumulation.better) / float64(accumulation.total)}
	}

	return writeInitialOutcomes(outcomes)
}

type outcome struct {
	win  float64
	tie  float64
	lose float64
}

type accumulation struct {
	better int64
	same   int64
	worse  int64
	total  int64
}

func readInitialOutcomes() (map[base.Key]outcome, error) {
	file, err := os.Open(initialOutcomeFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	outcomes := make(map[base.Key]outcome)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 4 {
			return nil, fmt.Errorf("unable to parse line in outcome file: %v", line)
		}
		key, err := strconv.ParseUint(parts[0], 0, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse key in outcome file: %v", parts[0])
		}
		win, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse win in outcome file: %v", parts[1])
		}
		tie, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse tie in outcome file: %v", parts[2])
		}
		lose, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse lose in outcome file: %v", parts[3])
		}
		outcomes[base.Key(key)] = outcome{win, tie, lose}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return outcomes, nil
}

// count how many hands among hands can coexist with the hand, i.e. do not have overlapping cards
func countCoexisting(hand []base.Card, hands [][]base.Card) int64 {
	result := int64(0)
loop:
	for _, h := range hands {
		for _, x := range h {
			for _, y := range hand {
				if x == y {
					continue loop
				}
			}
		}
		result++
	}
	return result
}

func writeInitialOutcomes(outcomes map[base.Key]outcome) error {
	var b bytes.Buffer
	for key, outcome := range outcomes {
		b.WriteString(fmt.Sprintf("%#x %v %v %v\n", key, outcome.win, outcome.tie, outcome.lose))
	}

	file, err := os.Create(initialOutcomeFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b.Bytes())
	if err != nil {
		return err
	}
	return file.Sync()
}
