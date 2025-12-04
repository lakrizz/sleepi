package shared

// Weekday mirrors time.Weekday but with a clear domain mapping (1–7).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func ParseWeekdays(input []int) []Weekday {
	res := make([]Weekday, len(input))
	for i, v := range input {
		res[i] = Weekday(v)
	}

	return res
}
