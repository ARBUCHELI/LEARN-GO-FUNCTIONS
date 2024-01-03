package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int;
  switch planet {
    case "Venus":
      fuel = 300000
      break
    case "Mercury":
      fuel = 500000
      break
    case "Mars":
      fuel = 700000
      break
    default:
      fuel = 0
  }
  fmt.Println("For going to", planet, "you need", fuel, "gallons of fuel!")
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
 var fuelRemaining, fuelCost int
 fuelRemaining = fuel
 fuelCost = calculateFuel(planet)

 if (fuelRemaining >= fuelCost) {
  greetPlanet(planet)
  fuelRemaining -= fuelCost
 } else {
  cantFly()
 }
 return fuelRemaining
}

func main() {
  // Test your functions!
  //fuelGauge(140)
  //calculateFuel("Venus")
  //greetPlanet("Venus")
 //cantFly()
  //fmt.Println(flyToPlanet("Mercury", 600000))
  // Create `planetChoice` and `fuel`
  var fuel int
  fuel = 1000000
  var planetChoice string
  planetChoice = "Venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
}