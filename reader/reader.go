package reader

import (
	"bufio"
	"io"
	"routely/data"
	"strconv"
)

// ReadData read test data from io stream
func ReadData(mapFile io.Reader, queryFile io.Reader, isBonus bool) ([]*data.Intersection, []*data.Road, []*data.Query) {
	scannerd := bufio.NewScanner(mapFile)
	scannerd.Split(bufio.ScanWords)

	scannerq := bufio.NewScanner(queryFile)
	scannerq.Split(bufio.ScanWords)

	return readIntersections(scannerd), readRoads(scannerd, isBonus), readQueries(scannerq)
}

func readIntersections(scanner *bufio.Scanner) []*data.Intersection {
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	ret := make([]*data.Intersection, n)
	for i := 0; i < n; i++ {
		ret[i] = &data.Intersection{}

		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		ret[i].ID = id

		scanner.Scan()
		ret[i].X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		ret[i].Y, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
	}
	return ret
}

func readRoads(scanner *bufio.Scanner, isBonus bool) []*data.Road {
	scanner.Scan()
	m, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	if isBonus {
		scanner.Scan()
		speedConut, _ := strconv.Atoi(scanner.Text())
		data.SpeedCount = speedConut

		scanner.Scan()
		speedInterval, _ := strconv.Atoi(scanner.Text())
		data.SpeedInterval = speedInterval
	} else {
		data.SpeedCount = 1
		data.SpeedInterval = 1e9
	}

	ret := make([]*data.Road, m)
	for i := 0; i < m; i++ {
		ret[i] = &data.Road{}
		ret[i].Speeds = make([]float64, data.SpeedCount)

		scanner.Scan()
		from, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		ret[i].From = from
		scanner.Scan()
		to, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		ret[i].To = to

		scanner.Scan()
		ret[i].Length, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		for j := 0; j < data.SpeedCount; j++ {
			scanner.Scan()
			ret[i].Speeds[j], err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				panic(err)
			}
		}
	}
	return ret
}

func readQueries(scanner *bufio.Scanner) []*data.Query {
	scanner.Scan()
	q, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	ret := make([]*data.Query, q)
	for i := 0; i < q; i++ {
		ret[i] = &data.Query{}

		scanner.Scan()
		ret[i].From.X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		ret[i].From.Y, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		ret[i].To.X, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		ret[i].To.Y, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		ret[i].WalkingRadius, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		ret[i].WalkingRadius /= 1000.0
	}
	return ret
}
