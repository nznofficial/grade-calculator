package main

import "fmt"

func calculateAverage(scores []float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

func getLetterGrade(average float64) string {
	switch {
	case average >= 90:
		return "A"
	case average >= 80:
		return "B"
	case average >= 70:
		return "C"
	case average >= 60:
		return "D"
	default:
		return "F"
	}
}

func provideFeedback(letterGrade string, average float64) {
	fmt.Printf("Your average score is %.2f\n", average)
	fmt.Printf("Letter grade: %s\n", letterGrade)

	if letterGrade == "A" {
		fmt.Println("Excellent work! Keep up the outstanding performance!")
	} else if letterGrade == "B" || letterGrade == "C" {
		fmt.Println("Good job! With a little more effort, you can achieve even better results")
	} else {
		fmt.Println("There's room for improvement. Consider seeking additional help or study resources.")
	}
}

func main() {
	fmt.Println("=== Student Grade Calculator ===")

	//Sample test scores for demonstration
	studentScores := []float64{85.5, 92.0, 78.5, 88.0, 95.5}
	studentName := "Turbo Nitro"

	fmt.Printf("Calculating grades for: %s\n", studentName)
	fmt.Println("Test Scores:", studentScores)

	average := calculateAverage(studentScores)
	letterGrade := getLetterGrade(average)

	fmt.Println("\n--- Results ---")
	provideFeedback(letterGrade, average)

	//Demonstrate conditional logic with different score scenarios
	fmt.Println("\n--- Additional Examples ---")

	lowScores := []float64{55.0, 62.0, 58.5}
	lowAverage := calculateAverage(lowScores)
	lowLetterGrade := getLetterGrade(lowAverage)

	fmt.Printf("Low performer average: %.2f (Grade %s)\n", lowAverage, lowLetterGrade)

	highScores := []float64{96.0, 98.5, 94.0, 97.5}
	highAverage := calculateAverage(highScores)
	highLetterGrade := getLetterGrade(highAverage)

	fmt.Printf("High performer average: %.2f (Grade %s)\n", highAverage, highLetterGrade)
}
