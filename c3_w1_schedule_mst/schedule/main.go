package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var o *string
var f *string

type job struct {
	w int
	l int
}

func (j *job) diff() float32 {
	return float32(j.w - j.l)
}

func (j *job) ratio() float32 {
	return float32(j.w) / float32(j.l)
}

func init() {
	o = flag.String("o", "ratio", "ordering")
	f = flag.String("f", "main", "data file")
	flag.Parse()
}

func main() {
	var less func(i, j job) bool
	_ = less
	switch *o {
	case "ratio":
		less = func(i, j job) bool { return i.ratio() > j.ratio() }
	case "diff":
		less = func(i, j job) bool {
			si := i.diff()
			sj := j.diff()
			if si > sj {
				return true
			}
			if si == sj && i.w > j.w {
				return true
			}
			return false
		}
	default:
		panic("unknown ordering mode " + *o)
	}
	var file string
	switch *f {
	case "main":
		file = "./jobs.txt"
	case "test":
		file = "./jobs_test.txt"
	case "test2":
		file = "./jobs_test2.txt"

	default:
		panic("unknown file" + *f)
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	if !s.Scan() {
		if s.Err() != nil {
			panic(err)
		}
		panic("empty file")
	}

	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	jobs := make([]job, n)
	i := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic("wrong parts number " + line)
		}
		w, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		l, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		jobs[i] = job{
			w: w,
			l: l,
		}
		i++
	}
	if i != n {
		fmt.Println("expected", n, "parsed:", i)
		panic("elements count not match expected number")
	}

	if s.Err() != nil {
		panic(s.Err())
	}

	quickSort(jobs, median, less)
	fmt.Println(score(jobs))
}

func score(jobs []job) int64 {
	c := 0
	score := int64(0)
	for _, job := range jobs {
		c += job.l
		score += int64(job.w * c)
	}
	return score
}
