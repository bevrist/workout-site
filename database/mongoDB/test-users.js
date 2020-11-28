db = db.getSiblingDB("workoutsite")

//test2
db.clients.insert({ //reasonable looking user
  "uid": "test2",
  "first_name": "Brett",
  "last_name": "Evrist",
  "weight": 100,
  "waistcirc": 31,
  "heightinches": 75,
  "leanbodymass": 15,
  "age": 20,
  "gender": "male",
  "start_date": "2020-10-01",
  "recommendation": [{
    "HighDayProtein": 101,
    "HighDayCarb": 111,
    "HighDayFat": 121,
    "HighDayCalories": 131,
    "NormalDayProtein": 141,
    "NormalDayCarb": 151,
    "NormalDayFat": 161,
    "NormalDayCalories": 171,
    "LowDayProtein": 181,
    "LowDayCarb": 191,
    "LowDayFat": 201,
    "LowDayCalories": 211,
    "ModifiedDate": "2020-09-15",
  },
  {},
  {},
  {
    "HighDayProtein": 22,
    "HighDayCarb": 12,
    "HighDayFat": 12,
    "HighDayCalories": 12,
    "NormalDayProtein": 12,
    "NormalDayCarb": 12,
    "NormalDayFat": 12,
    "NormalDayCalories": 12,
    "LowDayProtein": 12,
    "LowDayCarb": 12,
    "LowDayFat": 22,
    "LowDayCalories": 22,
    "HIITCurrentCardioSession": 22,
    "HIITCurrentCardioIntervals": 22,
    "ModifiedDate": "2020-09-16",
    },
    {
      "HighDayProtein": 23,
      "HighDayCarb": 13,
      "HighDayFat": 13,
      "HighDayCalories": 13,
      "NormalDayProtein": 13,
      "NormalDayCarb": 13,
      "NormalDayFat": 13,
      "NormalDayCalories": 13,
      "LowDayProtein": 13,
      "LowDayCarb": 13,
      "LowDayFat": 23,
      "LowDayCalories": 23,
      "HIITCurrentCardioSession": 231,
      "HIITCurrentCardioIntervals": 232,
      "ModifiedDate": "2020-09-16",
    }],
  "week": [
    {
      "day": [
        {
          "waistcirc": 20,
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
        {},
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "waistcirc": 21,
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "waistcirc": 22,
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "normal",
          "weight": 123,
          "cardio": "missed",
          "weight_training": "yes"
        },
        {},
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "waistcirc": 23,
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "waistcirc": 25,
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        }
      ]
    },
    {
      "Day": [
        {},
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {},
        {
          "fat": 10,
          "carbs": 10,
          "protein": 10,
          "total_calories": 30,
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        }
      ]
    }
  ]
});

//test3
db.clients.insert({ //user who completed all sessions
  "uid": "test3",
  "first_name": "Maxed",
  "last_name": "Out",
  "weight": 99,
  "waistcirc": 30,
  "heightinches": 35,
  "leanbodymass": 30,
  "age": 30,
  "gender": "female",
  "start_date": "2020-10-02",
  "recommendation": [{
    "HighDayProtein": 111,
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
    "HIITCurrentCardioIntervals": 24,
    "ModifiedDate": "2020-09-13",
  },
  {},
  {},
  {
    "HighDayProtein": 22,
    "HighDayCarb": 12,
    "HighDayFat": 12,
    "HighDayCalories": 12,
    "NormalDayProtein": 12,
    "NormalDayCarb": 12,
    "NormalDayFat": 12,
    "NormalDayCalories": 12,
    "LowDayProtein": 12,
    "LowDayCarb": 12,
    "LowDayFat": 22,
    "LowDayCalories": 22,
    "HIITCurrentCardioSession": 22,
    "HIITCurrentCardioIntervals": 22,
    "ModifiedDate": "2020-09-16",
    },
    {
      "HighDayProtein": 23,
      "HighDayCarb": 13,
      "HighDayFat": 13,
      "HighDayCalories": 13,
      "NormalDayProtein": 13,
      "NormalDayCarb": 13,
      "NormalDayFat": 13,
      "NormalDayCalories": 13,
      "LowDayProtein": 13,
      "LowDayCarb": 13,
      "LowDayFat": 23,
      "LowDayCalories": 23,
      "HIITCurrentCardioSession": 23,
      "HIITCurrentCardioIntervals": 23,
      "ModifiedDate": "2020-09-16",
    }],
  "week": [
    {
      "day": [
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "waistcirc": 10,
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "waistcirc": 20,
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "waistcirc": 30,
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        }
      ]
    },
    {
      "Day": [
        {
          "waistcirc": 40,
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
        }
      ]
    },
    {
      "day": [
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
        }
      ]
    },
    {
      "Day": [
        {
          "waistcirc": 50,
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
        }
      ]
    },
    {
      "day": [
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
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
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
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
          "fat": 110,
          "carbs": 110,
          "protein": 110,
          "total_calories": 310,
          "day_calorie": "high",
          "weight": 123,
          "cardio": "hit",
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
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
          "day_calorie": "low",
          "weight": 123,
          "cardio": "hit",
          "weight_training": "no"
        },
        {
          "fat": 11,
          "carbs": 11,
          "protein": 11,
          "total_calories": 31,
          "day_calorie": "normal",
          "weight": 222,
          "cardio": "missed",
          "weight_training": "yes"
        }
      ]
    },

  ]
});
