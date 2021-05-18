package greetings

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func canCross(stones []int) bool {
	var dict_stones = map[int]map[int]bool{}
	for _, stone := range stones {
		dict_stones[stone] = map[int]bool{}
	}
	dict_stones[stones[0]][1] = false
	for i, stone := range stones {
		for step, _ := range dict_stones[stone] {
			var exist bool
			var step_map map[int]bool
			step_map, exist = dict_stones[i+step]
			if exist {
				if step - 1 > 0 {
					step_map[step-1] = true
				}
				if step > 0 {
					step_map[step] = true
				}
				if step + 1 > 0 {
					step_map[step] = true
				}
			}
		}

	}
	var last_stone map[int]bool
	last_stone, _ = dict_stones[stones[len(stones)-1]]
	if len(last_stone) == 0 {
		return false
	}
	return true

}


func reverse(x int) int {
	var sign int
	if x >= 0 {
		sign = 1
	} else {
		sign = -1
	}
	x = int(math.Abs(float64(x)))
	var res = x % 10
	x = x / 10
	for x > 0 {
		res = res * 10 + x % 10
		x = x / 10
	}
	var downLimit = -1 * math.Pow(2, 31)
	var upLimit = math.Pow(2,31) - 1
	if x >= int(downLimit) && x <= int(upLimit) {
		return x
	}
	return sign * x
}

func xorOperation(n int, start int) int {
	res := 0
	var temp int
	for i := 0; i < n; i++ {
		temp = start + i * 2
		res = res ^ temp
	}
	return res
}

