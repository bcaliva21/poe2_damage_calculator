package main

import (
	"fmt"
	"math"
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
	fmt.Println("\nCurrent damage scaling values:", damageScaling)

	damageScaling.increasedDamageTotal = collectIncreasedDamageValues()
	fmt.Println("\nCurrent damage scaling values:", damageScaling)

	damageScaling.moreDamageTotal = collectMoreDamageValues()
	fmt.Println("\nCurrent damage scaling values:", damageScaling)

	damageScaling.addedDamageTotal = collectAddedDamageValues()
	fmt.Println("\nCurrent damage scaling values:", damageScaling)

	damageScaling.hitRate = collectHitRate()
	fmt.Println("\nCurrent damage scaling values:", damageScaling)

	calculateDamagePerSecond(&damageScaling)
	fmt.Println("\nCurrent damage scaling values:", damageScaling)
}

type DamageScaling struct {
	baseDamage                int
	addedDamageTotal          int
	gainedDamageTotal         int
	moreDamageTotal           float32
	increasedDamageTotal      float32
	hitRate                   float32
	criticalStrikeDamageTotal int
	criticalStrikeChanceTotal int
	damagePerSecond           float32
}

func convertToMultiplier(percentage int) float32 {
	return float32(percentage)/100 + 1
}

func calculateDamagePerSecond(damageScaling *DamageScaling) {
	totalBaseDamage := float32(damageScaling.baseDamage + damageScaling.addedDamageTotal)
	damageScaling.damagePerSecond = totalBaseDamage
	damageScaling.damagePerSecond *= damageScaling.increasedDamageTotal
	damageScaling.damagePerSecond *= damageScaling.moreDamageTotal
	damageScaling.damagePerSecond *= damageScaling.hitRate
}

func collectHitRate() float32 {
	fmt.Println("Enter your attacks/casts per second")

	var userInput string
	fmt.Print("Attack/cast per second: ")
	fmt.Scanln(&userInput)

	hitRateValue, err := strconv.ParseFloat(userInput, 32)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid integer value.")
	}
	return float32(hitRateValue)
}

func collectAddedDamageValues() int {
	fmt.Println("Enter values for added damage")
	var added int

	for {
		// Prompt the user to enter low and high values
		fmt.Println("\nEnter a low value and a high value, or type 'done' to finish:\n")

		var lowInput, highInput string
		fmt.Print("Low value: ")
		fmt.Scanln(&lowInput)

		// Check if the user wants to finish
		if lowInput == "done" {
			break
		}

		fmt.Print("High value: ")
		fmt.Scanln(&highInput) // Convert user input to integer

		if highInput == "done" {
			break
		}
		lowValue, errLow := strconv.Atoi(lowInput)
		highValue, errHigh := strconv.Atoi(highInput)

		if errLow != nil || errHigh != nil {
			fmt.Println("Invalid input, please enter a valid integer value.")
			continue
		}

		rawAvg := float64((lowValue + highValue) / 2)
		avgValue := int(math.Floor(rawAvg))

		added += avgValue
	}

	return added
}

func collectMoreDamageValues() float32 {
	fmt.Println("Enter more damage values: ")

	multiplier := float32(1)
	for {
		fmt.Println("\nEnter a value (percentage as integer, or type 'done' to finish):\n")
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

func collectIncreasedDamageValues() float32 {
	fmt.Println("Enter increased damage values: ")
	var sum float32
	for {
		fmt.Println("\nEnter a value (percentage as integer, or type 'done' to finish):\n")
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
		"{\n  baseDamage: %d,\n  addedDamageTotal: %d,\n  gainedDamageTotal: %d,\n  moreDamageTotal: %.2f,\n  increasedDamageTotal: %.2f,\n  hitRate: %.2f,\n  criticalStrikeDamageTotal: %d,\n  criticalStrikeChanceTotal: %d,\n  damagePerSecond: %.2f\n}",
		d.baseDamage,
		d.addedDamageTotal,
		d.gainedDamageTotal,
		d.moreDamageTotal,
		d.increasedDamageTotal,
		d.hitRate,
		d.criticalStrikeDamageTotal,
		d.criticalStrikeChanceTotal,
		d.damagePerSecond,
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
