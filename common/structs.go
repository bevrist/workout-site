package structs

// Auth holds authentication info for communicating with the AUTH service
type Auth struct {
	IsValid bool
	UID     string
}

//UserInfo holds user information for communicating with the BACKEND service
type UserInfo struct {
	FirstName    string
	LastName     string
	Weight       int
	WaistCirc    int
	HeightInches int
	LeanBodyMass int
	Age          int
	Gender       string
}

//Client holds user information in the same format as the Mongo database
type Client struct {
	UID string
	FirstName string
	LastName string
	Weight int
	WaistCirc int
	HeightInches int
	LeanBodyMass int
	Age int
	Gender string
	Week []Week
}

type Week struct {
	Day []Day
}

type Day struct {
	Fat int
	Carbs int
	Protein	int
	TotalCalories int
	DayCalorie int
	Weight int
	Cardio string
	WeightTraining string
}
