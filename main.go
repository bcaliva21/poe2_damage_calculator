package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to the CLI Question App!")

	var damageScaling DamageScaling

	// Ask the user a question
	question := "What is your base damage?"
	damageScaling.baseDamage = collectInt(question)

	// Display the answer
	fmt.Println("Current damage scaling values:", damageScaling)

	increasedDamageValue := collectIncreasedDamageValues("Increased damage values")
	fmt.Printf("The final damage value is: %.2f\n", increasedDamageValue)

	moreDamageValue := collectMoreDamageValues("More damage values")
	fmt.Printf("%.2f", moreDamageValue)
}

type DamageScaling struct {
	baseDamage                int
	addedDamageTotal          int
	gainedDamageTotal         int
	moreDamageTotal           float32
	increasedDamageTotal      float32
	hitRate                   int
	criticalStrikeDamageTotal int
	criticalStrikeChanceTotal int
}

func convertToMultiplier(percentage int) float32 {
	return float32(percentage)/100 + 1
}

func collectMoreDamageValues(category string) float32 {
	fmt.Println("Enter values for ", category)

	multiplier := float32(1)
	for {
		fmt.Println("Enter a value (percentage as integer, or type 'done' to finish):")
		var userInput string
		fmt.Scanln(&userInput)

		// Check if user wants to finish
		if userInput == "done" {
			break
		}

		// Convert user input to integer
		value, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer value.")
			continue
		}

		// Add value to sum (convert percentage to a multiplier)
		subValue := convertToMultiplier(value)
		fmt.Printf("%.2f", subValue)
		multiplier *= subValue
	}

	// Return the sum (add base of 100)
	return multiplier
}

func collectIncreasedDamageValues(category string) float32 {
	// Ask user for input type (individual or total)
	// Normalize input type (to handle casing)
	// Handle individual values input
	fmt.Println("Enter values for ", category)
	var sum float32
	for {
		fmt.Println("Enter a value (percentage as integer, or type 'done' to finish):")
		var userInput string
		fmt.Scanln(&userInput)

		// Check if user wants to finish
		if userInput == "done" {
			break
		}

		// Convert user input to integer
		value, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer value.")
			continue
		}

		// Add value to sum (convert percentage to a multiplier)
		quotient := float32(value) / float32(100)
		fmt.Printf("%.2f", quotient)
		sum += quotient
	}

	return sum + float32(1)
}

func (d DamageScaling) String() string {
	return fmt.Sprintf(
		"{\n  baseDamage: %d,\n  addedDamageTotal: %d,\n  gainedDamageTotal: %d,\n  moreDamageTotal: %.2f,\n  increasedDamageTotal: %d,\n  hitRate: %d,\n  criticalStrikeDamageTotal: %d,\n  criticalStrikeChanceTotal: %d\n}",
		d.baseDamage,
		d.addedDamageTotal,
		d.gainedDamageTotal,
		d.moreDamageTotal,
		d.increasedDamageTotal,
		d.hitRate,
		d.criticalStrikeDamageTotal,
		d.criticalStrikeChanceTotal,
	)
}

// Helper function to format slices
func formatSlice(slice []float32) string {
	var str []string
	for _, val := range slice {
		str = append(str, fmt.Sprintf("%.2f", val))
	}
	return fmt.Sprintf("[%s]", strings.Join(str, ", "))
}

func collectInt(question string) int {
	fmt.Printf("%s\n", question)

	// Read the input from the user
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Return the user's input, trimming extra spaces
	return input
}

// askQuestion displays a question and waits for the user input
func askQuestion(question string) string {
	fmt.Printf("%s\n", question)

	// Read the input from the user
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Return the user's input, trimming extra spaces
	return strings.TrimSpace(input)
}
