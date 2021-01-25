package app

type Cache struct {
	courseTimes map[string]map[string]CourseTimes
}

type Data struct {
	Date    string
	Courses []CourseTimes
}

type CourseTimes struct {
	Name  string
	Times []TeeTime
}

type TeeTime struct {
	Price int
	Time  string
}
