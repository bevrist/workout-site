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
	UID          string
	FirstName    string  `bson:"first_name,omitempty"`
	LastName     string  `bson:"last_name,omitempty"`
	Weight       int     `bson:"weight,omitempty"`
	WaistCirc    float64 `bson:"waistcirc,omitempty"`
	HeightInches int     `bson:"heightinches,omitempty"`
	LeanBodyMass int     `bson:"leanbodymass,omitempty"`
	Age          int     `bson:"age,omitempty"`
	Gender       string  `bson:"gender,omitempty"`
	Week         []struct {
		Day []struct {
			Fat            int    `bson:"fat,omitempty"`
			Carbs          int    `bson:"carbs,omitempty"`
			Protein        int    `bson:"protein,omitempty"`
			TotalCalories  int    `bson:"total_calories,omitempty"`
			DayCalorie     string `bson:"day_calorie,omitempty"`
			Weight         int    `bson:"weight,omitempty"`
			Cardio         string `bson:"cardio,omitempty"`
			WeightTraining string `bson:"weight_training,omitempty"`
		}
	}
}
