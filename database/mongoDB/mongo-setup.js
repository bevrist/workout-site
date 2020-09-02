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
  "week": [
    {
      "Day": [            // Week one
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
    },
    {
      "Day": [
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
        }
      ]
    },
    {
      "Day": [
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
        }
      ]
    }
  ]
});