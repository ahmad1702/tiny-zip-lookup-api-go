package seed

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type ZipInfoRecord struct {
	Zip        int
	City       string
	State      string
	StateAbbr  string
	County     string
	CountyCode int
	Latitude   float64
	Longitude  float64
}

func GetZipInfos() map[int]ZipInfoRecord {
	file, err := os.Open("./data/us_zips.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	zips := make(map[int]ZipInfoRecord)
	csvReader := csv.NewReader(file)
	i := 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		zip, err := strconv.Atoi(rec[0])
		if err != nil {
			log.Fatal(err)
		}
		county_code, err := strconv.Atoi(rec[5])
		if err != nil {
			log.Fatal(err)
		}

		latitude, err := strconv.ParseFloat(rec[6], 64)
		if err != nil {
			log.Fatal(err)
		}
		longitude, err := strconv.ParseFloat(rec[6], 64)
		if err != nil {
			log.Fatal(err)
		}

		zipInfo := ZipInfoRecord{
			Zip:        zip,
			City:       rec[1],
			State:      rec[2],
			StateAbbr:  rec[3],
			County:     rec[4],
			CountyCode: county_code,
			Latitude:   latitude,
			Longitude:  longitude,
		}

		zips[zip] = zipInfo
		i++
	}
	fmt.Printf("Read %d records\n", i)

	return zips
}
