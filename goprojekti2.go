package main
import "fmt"
import "time"
import "sort"



func main() {
	t0 := time.Date(2020,8,24,22,22,22,22,time.UTC)
	t1 := time.Date(2020,8,22,22,22,22,22,time.UTC)
	t2 := time.Date(2020,8,22,22,22,22,22,time.UTC)
	t3 := time.Date(2020,8,20,22,22,22,22,time.UTC)
	t4 := time.Date(2020,8,19,22,22,22,22,time.UTC)
	t5 := time.Date(2020,7,31,22,22,22,22,time.UTC)
        tavarat := []struct {
                nimi string
                pvm time.Time
        }{
                {"kalja", t5},
                {"maito", t4},
		{"kananmuna", t3},
		{"juusto", t2},
		{"tomaatti", t1},
		{"kala", t0},

        }


sort.Slice(tavarat, func(i, j int) bool {
	return tavarat[i].pvm.Before(tavarat[j].pvm) })

        for _, tavara := range tavarat {
                fmt.Println(tavara)

	}
}


