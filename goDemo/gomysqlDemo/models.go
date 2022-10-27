package main

import "time"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type Project struct {
	Base          string     `json:"base"`
	Bd            string     `json:"bd"`
	Sa            string     `json:"sa"`
	Manager       string     `json:"manager"`
	Province      string     `json:"province"`
	City          string     `json:"city"`
	Company       string     `json:"company"`
	Spec          string     `json:"spec"`
	Workload      float64    `json:"workload"`
	Starting_date *time.Time `json:"starting_date"`
	Ending_date   *time.Time `json:"ending_date"`
	Status        string     `json:"status"`
	Payment       float64    `json:"payment"`
	Remarks       string     `json:"remarks"`
}
