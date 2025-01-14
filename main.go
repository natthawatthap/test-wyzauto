package main

import "fmt"

type Route struct {
    StartPoint string
    EndPoint   string
    Duration   int
}

func AddRoute(routes []Route, startPoint, endPoint string, duration int) []Route {
    newRoute := Route{StartPoint: startPoint, EndPoint: endPoint, Duration: duration}
    return append(routes, newRoute)
}

func CalculateRouteDuration(routes []Route, startPoint, endPoint string) int {
    totalDuration := 0
    currentPoint := startPoint
    visited := make(map[string]bool)

    for currentPoint != "" {
        found := false
        for _, route := range routes {
            if route.StartPoint == currentPoint && !visited[currentPoint] {
                totalDuration += route.Duration
                visited[currentPoint] = true
                currentPoint = route.EndPoint
                found = true
                if currentPoint == endPoint {
                    return totalDuration
                }
                break
            }
        }
        if !found {
            return 0 
        }
    }

    return 0 
}

func main() {
    // Create an array of Route
    routes := []Route{
        {StartPoint: "A", EndPoint: "B", Duration: 15},
        {StartPoint: "B", EndPoint: "C", Duration: 20},
        {StartPoint: "C", EndPoint: "D", Duration: 10},
    }

    // Add a new route
    routes = AddRoute(routes, "D", "E", 25)

    // Print each route
    for _, route := range routes {
        fmt.Printf("Route: %s -> %s, Duration: %d minutes\n", route.StartPoint, route.EndPoint, route.Duration)
    }

    // Calculate the total duration from A to C
    totalDuration := CalculateRouteDuration(routes, "A", "C")
    fmt.Printf("Total duration from A to C: %d minutes\n", totalDuration)
}
