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
        tavarat := []struct { //couldn't get time.Date in an understandable form
                nimi string //-rm, that the function reading the slice
                pvm time.Time //would've understood before the time limit.
        }{ //I need further research in software development on the time var
                {"kalja", t5},
                {"maito", t4},
		{"kananmuna", t3},
		{"juusto", t2},
		{"tomaatti", t1},
		{"kala", t0},

        }


sort.Slice(tavarat, func(i, j int) bool {
	return tavarat[i].pvm.Before(tavarat[j].pvm) })
//sorts time.Time values from smallest to biggest
        for _, tavara := range tavarat { //prints all structs in the slice
                fmt.Println(tavara) //couldn't figure out a way to print them
//				in my desired format starting with day of the mo
	} // this was intended to be a last consumption day list of the products
} //	in your fridge. (Viimeinen käyttöpäivä)


