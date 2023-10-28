package adapter

import "encoding/xml"

type Stock struct {
}

func (s *Stock) Data() ([]byte, error) {
	recs := []Record{}
	recs = append(recs, Record{"007", 700, 10000, 10000})
	recs = append(recs, Record{"alibaba", 300, 10000, 10000})

	return xml.Marshal(recs)
}

type Record struct {
	Name    string
	Price   float64
	BuyNum  int64
	SellNum int64
}
