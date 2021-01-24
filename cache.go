package main

type Cache struct {
	courseTimes map[string]map[string]CourseTimes
}

type CourseTimes struct {
	Name string

}

type TeeTime struct {
	Price int
	Time string
}