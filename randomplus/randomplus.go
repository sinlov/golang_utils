package randomplus

import (
	"time"
	"math/rand"
)

func Positive(size int) (int, error) {
	if size > 13 {
		return 0, RandomSizeError{
			13, "13",
		}
	}
	time_int := time.Now().UnixNano()
	r := rand.New(rand.NewSource(time_int))
	max_num := 1
	for i := 1; i < size + 1; i++ {
		max_num = max_num * 10
	}
	max_num--
	min_num := 1
	for i := 1; i < size; i++ {
		min_num = min_num * 10
	}
	min_num--
	NewRandom:
	res_int := r.Intn(max_num)
	if res_int < min_num {
		goto NewRandom
	}
	return res_int, nil
}

func PositiveNegative(size int) (int, error) {
	if size > 13 {
		return 0, RandomSizeError{
			13, "13",
		}
	}
	time_int := time.Now().UnixNano()
	r := rand.New(rand.NewSource(time_int))
	max_num := 1
	for i := 1; i < size + 1; i++ {
		max_num = max_num * 10
	}
	max_num--
	min_num := 1
	for i := 1; i < size; i++ {
		min_num = min_num * 10
	}
	min_num--
	NewRandom:
	res_int := r.Intn(max_num)
	if res_int < min_num {
		goto NewRandom
	}
	if (time_int % 2) == 0 {
		res_int = res_int * -1
	}
	return res_int, nil
}