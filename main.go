package main

import (
    "fmt"
    "math"
    "net/http"
    "strconv"
    "strings"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/submit", handleFormSubmit)

    fmt.Println("Starting server at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
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

func handleFormSubmit(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var damageScaling DamageScaling

        baseDamage, _ := strconv.Atoi(r.FormValue("baseDamage"))
        addedDamage := r.FormValue("addedDamage")
        moreDamage := r.FormValue("moreDamage")
        increasedDamage := r.FormValue("increasedDamage")
        hitRate, _ := strconv.ParseFloat(r.FormValue("hitRate"), 32)

        damageScaling.baseDamage = baseDamage
        damageScaling.addedDamageTotal = collectAddedDamageValuesFromString(addedDamage)
        damageScaling.moreDamageTotal = collectMoreDamageValuesFromString(moreDamage)
        damageScaling.increasedDamageTotal = collectIncreasedDamageValuesFromString(increasedDamage)
        damageScaling.hitRate = float32(hitRate)

        calculateDamagePerSecond(&damageScaling)

        fmt.Fprintf(w, "Damage Per Second: %.2f", damageScaling.damagePerSecond)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func collectAddedDamageValuesFromString(input string) int {
    if input == "" {
        return 0
    }
    var added int
    values := strings.Split(input, ",")
    for i := 0; i < len(values); i += 2 {
        lowValue, _ := strconv.Atoi(strings.TrimSpace(values[i]))
        highValue, _ := strconv.Atoi(strings.TrimSpace(values[i+1]))
        rawAvg := float64((lowValue + highValue) / 2)
        avgValue := int(math.Floor(rawAvg))
        added += avgValue
    }
    return added
}

func collectMoreDamageValuesFromString(input string) float32 {
    multiplier := float32(1)
    if input == "" {
        return multiplier
    }
    values := strings.Split(input, ",")
    for _, valueStr := range values {
        value, _ := strconv.Atoi(strings.TrimSpace(valueStr))
        subValue := convertToMultiplier(value)
        multiplier *= subValue
    }
    return multiplier
}

func collectIncreasedDamageValuesFromString(input string) float32 {
    var sum float32
    if input == "" {
        return float32(1)
    }
    values := strings.Split(input, ",")
    for _, valueStr := range values {
        value, _ := strconv.Atoi(strings.TrimSpace(valueStr))
        quotient := float32(value) / float32(100)
        sum += quotient
    }
    return sum + float32(1)
}
