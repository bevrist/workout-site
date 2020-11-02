db = db.getSiblingDB("workoutsite")

db.clients.insert({
  "uid": "testUID",
  "first_name": "Anthony",
  "last_name": "Hanna",
  "weight": 215,
  "waistcirc": 35.5,
  "heightinches": 75,
  "leanbodymass": 15,
  "age": 20,
  "gender": "male",
  "start_date": "2020-08-15",
  "recommendation": [{
    "HighDayProtein": 10,
    "HighDayCarb": 11,
    "HighDayFat": 12,
    "HighDayCalories": 13,
    "NormalDayProtein": 14,
    "NormalDayCarb": 15,
    "NormalDayFat": 16,
    "NormalDayCalories": 17,
    "LowDayProtein": 18,
    "LowDayCarb": 19,
    "LowDayFat": 20,
    "LowDayCalories": 21,
    "HIITCurrentCardioSession": 22,
    "HIITChangeCardioSession": 23,
    "HIITCurrentCardioIntervals": 24,
    "HIITChangeCardioIntervals": 25,
    "ModifiedDate": "2020-09-13",
  }],
  "week": [
    {
      "day": [            // Week one
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 20,
          "carbs": 20,
          "protein": 20,
          "total_calories": 32,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "no"
        },
        {
          "fat": 30,
          "carbs": 30,
          "protein": 30,
          "total_calories": 33,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 40,
          "carbs": 40,
          "protein": 40,
          "total_calories": 34,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 100,
          "carbs": 100,
          "protein": 100,
          "total_calories": 300,
          "day_calorie": "normal",
          "weight": 321,
          "cardio": "missed",
          "weight_training": "no"
        }
      ]
    },
    {
      "Day": [
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "no"
        }
      ]
    },
    {
      "Day": [
        {
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        }
      ]
    }
  ]
});
