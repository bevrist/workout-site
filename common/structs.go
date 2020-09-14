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
	HighDayProtein             int    `bson:"highdayprotein,omitempty" json:",omitempty"`
	HighDayCarb                int    `bson:"highdaycarb,omitempty" json:",omitempty"`
	HighDayFat                 int    `bson:"highdayfat,omitempty" json:",omitempty"`
	HighDayCalories            int    `bson:"highdaycalories,omitempty" json:",omitempty"`
	NormalDayProtein           int    `bson:"normaldayprotein,omitempty" json:",omitempty"`
	NormalDayCarb              int    `bson:"normaldaycarb,omitempty" json:",omitempty"`
	NormalDayFat               int    `bson:"normaldayfat,omitempty" json:",omitempty"`
	NormalDayCalories          int    `bson:"normaldaycalories,omitempty" json:",omitempty"`
	LowDayProtein              int    `bson:"lowdayprotein,omitempty" json:",omitempty"`
	LowDayCarb                 int    `bson:"lowdaycarb,omitempty" json:",omitempty"`
	LowDayFat                  int    `bson:"lowdayfat,omitempty" json:",omitempty"`
	LowDayCalories             int    `bson:"lowdaycalories,omitempty" json:",omitempty"`
	HIITCurrentCardioSession   int    `bson:"hiitcurrentcardiosession,omitempty" json:",omitempty"`
	HIITChangeCardioSession    int    `bson:"hiitchangecardiosession,omitempty" json:",omitempty"`
	HIITCurrentCardioIntervals int    `bson:"hiitcurrentcardiointervals,omitempty" json:",omitempty"`
	HIITChangeCarioIntervals   int    `bson:"hiitchangecariointervals,omitempty" json:",omitempty"`
	Week                       int    `bson:"week,omitempty" json:",omitempty"`
	ModifiedDate               string `bson:"modifieddate,omitempty" json:",omitempty"`
}
