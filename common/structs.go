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
	UID            string           `json:"UID,omitempty"`
	FirstName      string           `bson:"first_name,omitempty"`
	LastName       string           `bson:"last_name,omitempty"`
	Weight         float64          `bson:"weight,omitempty"`
	WaistCirc      float64          `bson:"waistcirc,omitempty"`
	HeightInches   int              `bson:"heightinches,omitempty"`
	LeanBodyMass   int              `bson:"leanbodymass,omitempty"`
	Age            int              `bson:"age,omitempty"`
	Gender         string           `bson:"gender,omitempty"`
	Week           []Week           `bson:"week,omitempty"`
	Recommendation []Recommendation `bson:"recommendation,omitempty"`
}
type Week struct {
	Day []Day `bson:"day,omitempty"`
}
type Day struct {
	Fat            int
	Carbs          int
	Protein        int
	TotalCalories  int    `bson:"total_calories"`
	DayCalorie     string `bson:"day_calorie"`
	Weight         float64
	Cardio         string
	WeightTraining string `bson:"weight_training"`
}
type Recommendation struct {
	HighDayProtein             int      `bson:"highdayprotein,omitempty"`
	HighDayCarb                int      `bson:"highdaycarb,omitempty"`
	HighDayFat                 int      `bson:"highdayfat,omitempty"`
	HighDayCalories            int      `bson:"highdaycalories,omitempty"`
	NormalDayProtein           int      `bson:"normaldayprotein,omitempty"`
	NormalDayCarb              int      `bson:"normaldaycarb,omitempty"`
	NormalDayFat               int      `bson:"normaldayfat,omitempty"`
	NormalDayCalories          int      `bson:"normaldaycalories,omitempty"`
	LowDayProtein              int      `bson:"lowdayprotein,omitempty"`
	LowDayCarb                 int      `bson:"lowdaycarb,omitempty"`
	LowDayFat                  int      `bson:"lowdayfat,omitempty"`
	LowDayCalories             int      `bson:"lowdaycalories,omitempty"`
	HIITCurrentCardioSession   int      `bson:"hiitcurrentcardiosession,omitempty"`
	HIITChangeCardioSession    int      `bson:"hiitchangecardiosession,omitempty"`
	HIITCurrentCardioIntervals int      `bson:"hiitcurrentcardiointervals,omitempty"`
	HIITChangeCarioIntervals   int      `bson:"hiitchangecariointervals,omitempty"`
	Week                       int      `bson:"week,omitempty"`
	ModifiedDate               string   `bson:"modifieddate,omitempty"`
	Nil                        []string `bson:"nil,omitempty"`
}
