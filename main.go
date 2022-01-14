package main

import (
	"fmt"
	"time"
)

func printYear() {
	fmt.Println("\t\t\tJan\t\tFeb\t\tMar\t\tApr\t\tMay\t\tJun\t\tJul\t\tAug\t\tSep\t\tOct\t\tNov\t\tDec")
}

func printLine() {
	fmt.Println("------------------------------------------------------" +
		"-------------------------------------------------------------------" +
		"-------------------------------------------------------------------" +
		"---------------")
}

/*
A Budget should have a set period
- for instance every year
*/

func add(x [12]float32, y [12]float32) [12]float32 {
	var t [12]float32

	for i := 0; i < 12; i++ {
		t[i] = x[i] + y[i]
	}
	return t
}

func subtract(x [12]float32, y [12]float32) [12]float32 {
	var t [12]float32

	for i := 0; i < 12; i++ {
		t[i] = x[i] - y[i]
	}
	return t
}

type Event struct {
	StartDate time.Time
	EndDate time.Time
	Value float32
}

type Item struct {
	Name string
	Events []Event
}

type Group struct {
	Name string
	Items []Item
}

type Budget struct {
	Name string
	Incomes []Group
	Expenses []Group
}

func (i *Item) AppendNewEvent(startTime string, endTime string, value float32) *Event {
	event := &Event{
		StartDate: parseTime(startTime),
		EndDate: parseTime(endTime),
		Value: value,
	}
	i.Events = append(i.Events, *event)
	return event
}

func (i *Item) MonthlyTotal(year int) [12]float32 {
	var a [12]float32

	for _, ev := range i.Events {
		evi := yearArray(year, ev)
		a = add(a, evi)
	}
	return a
}

func (i *Item) PrintMonthlyTotal(year int) {
	a := i.MonthlyTotal(year)
	fmt.Printf("- %-11.11s : %11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\n", i.Name, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
}

func (g *Group) MonthlyTotal(year int) [12]float32 {
	var t [12]float32

	for _, i := range g.Items {
		a := i.MonthlyTotal(year)
		t = add(t, a)
	}
	return t
}

func (g *Group) PrintMonthlyTotal(year int) {
	a := g.MonthlyTotal(year)
	fmt.Printf("%-13.13s : %11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\n", g.Name, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
}

func (g *Group) PrintMonthlyTotalItems(year int) {
	g.PrintMonthlyTotal(year)
	for _, i := range g.Items {
		i.PrintMonthlyTotal(year)
	}
}

func (b *Budget) PrintMonthlyTotal(year int) {
	printYear()
	printLine()
	b.PrintMonthlyBalance(year)
	b.PrintMonthlyRemainder(year)
	printLine()
	b.PrintMonthlyIncome(year)
	printLine()
	for _, ig := range b.Incomes {
		ig.PrintMonthlyTotalItems(year)
	}
	printLine()
	b.PrintMonthlyExpense(year)
	printLine()
	for _, eg := range b.Expenses {
		eg.PrintMonthlyTotalItems(year)
	}
	printLine()
}

func (b *Budget) MonthlyIncome(year int) [12]float32 {
	var ic [12]float32

	for _, ig := range b.Incomes {
		a := ig.MonthlyTotal(year)
		ic = add(ic, a)
	}

	return ic
}
func (b *Budget) PrintMonthlyIncome(year int) {
	a := b.MonthlyIncome(year)
	fmt.Printf("%-13.13s : %11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\n", "Income", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
}


func (b *Budget) MonthlyExpense(year int) [12]float32 {
	var ex [12]float32

	for _, eg := range b.Expenses {
		a := eg.MonthlyTotal(year)
		ex = add(ex, a)
	}

	return ex
}
func (b *Budget) PrintMonthlyExpense(year int) {
	a := b.MonthlyExpense(year)
	fmt.Printf("%-13.13s : %11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\n", "Expense", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
}

func (b *Budget) MonthlyRemainder(year int) [12]float32 {
	var rem [12]float32
	ic := b.MonthlyIncome(year)
	ex := b.MonthlyExpense(year)

	rem = subtract(ic, ex)

	return rem
}
func (b *Budget) PrintMonthlyRemainder(year int) {
	a := b.MonthlyRemainder(year)
	fmt.Printf("%-13.13s : %11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\n", "Remainder", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
}

func (b *Budget) MonthlyBalance(year int) [12]float32 {
	var bal [12]float32

	rem := b.MonthlyRemainder(year)
	bal[0] = rem[0]
	for i := 1; i < 12; i++ {
		bal[i] = bal[i-1] + rem[i]
	}

	return bal
}
func (b *Budget) PrintMonthlyBalance(year int) {
	a := b.MonthlyBalance(year)
	fmt.Printf("%-13.13s : %11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\t%11.2f\n", "Balance", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
}


func main() {
	medicalAid := Item{Name: "Medical Aid"}
	medicalAid.AppendNewEvent("2021-01-10", "2021-12-10", 1990)
	medicalExpenses := Group{
		Name: "Medical Expenses",
		Items: []Item{medicalAid},
	}

	tuna := Item{Name: "Tuna"}
	tuna.AppendNewEvent("2021-01-10", "2021-01-10", 24)
	milk := Item{Name: "Milk"}
	milk.AppendNewEvent("2021-01-10", "2021-12-10", 105)
	foodExpenses := Group{
		Name: "Food Expenses",
		Items: []Item{tuna, milk},
	}

	salary := Item{Name: "Salary"}
	salary.AppendNewEvent("2021-01-25", "2021-12-25", 34000)

	bonus := Item{Name: "Bonus"}
	bonus.AppendNewEvent("2021-12-25", "2021-12-25", 24000)

	incomeGroup := Group{
		Name: "Salary",
		Items: []Item{salary, bonus},
	}

	b := Budget{
		Name: "Personal Budget",
		Incomes: []Group{incomeGroup},
		Expenses: []Group{medicalExpenses, foodExpenses},
	}

	var year = 2021
	b.PrintMonthlyTotal(year)
}

// parseTime takes a YYYY-MM-DD string and converts it to a time.Time type.
func parseTime(t string) time.Time {
	const shortForm string = "2006-01-02"  // YYYY-MM-DD

	ti, err := time.Parse(shortForm, t)
	if err != nil {
		fmt.Println(err)
	}
	return ti
}

// parseStringTime takes a year, month, day int and converts it to a
// YYYY-MM-DD string.
func parseStringTime(y int, m time.Month, d int) string {
	var ms, ds string

	ms = fmt.Sprintf("%d", m)
	if m < 10 {
		ms = fmt.Sprintf("0%d", m)
	}

	ds = fmt.Sprintf("%d", d)
	if d < 10 {
		ds = fmt.Sprintf("0%d", d)
	}

	return fmt.Sprintf("%d-%s-%s", y, ms, ds)
}

// yearArray converts the event from a time period it and array of monthly amounts.
// TODO: improve to incorporate specific dates, such as specific date (day) transactions or closest (last friday of month)
func yearArray(year int, e Event) [12]float32 {
	// declare variable
	var m [12]float32

	// do computation
	// loop through months
	yi, mi, _ := e.StartDate.Date()
	yj, mj, _ := e.EndDate.Date()
	//start := parseTime(fmt.Sprintf(startYear, ))
	if (yi <= year) && (yj >= year) {
		for i := 1; i <= 12; i++ {
			tmi := time.Month(i)
			//fmt.Println(mi, tmi, mj, mi <= tmi && tmi <= mj)
			if mi <= tmi && tmi <= mj {
				m[i-1] = e.Value
			}
		}
	}

	return m
}