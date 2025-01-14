import "testing"

// TestAddRoute tests the AddRoute function
func TestAddRoute(t *testing.T) {
    // Initialize test data
    initialRoutes := []Route{
        {StartPoint: "A", EndPoint: "B", Duration: 15},
        {StartPoint: "B", EndPoint: "C", Duration: 20},
    }

    // Expected result after adding a new route
    expectedRoutes := []Route{
        {StartPoint: "A", EndPoint: "B", Duration: 15},
        {StartPoint: "B", EndPoint: "C", Duration: 20},
        {StartPoint: "C", EndPoint: "D", Duration: 10},
    }

    // Add a new route
    updatedRoutes := AddRoute(initialRoutes, "C", "D", 10)

    // Verify length
    if len(updatedRoutes) != len(expectedRoutes) {
        t.Fatalf("Expected %d routes, got %d", len(expectedRoutes), len(updatedRoutes))
    }

    // Verify each route
    for i, route := range updatedRoutes {
        if route != expectedRoutes[i] {
            t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, expectedRoutes[i], route)
        }
    }
}


func TestCalculateRouteDuration(t *testing.T) {
    // Test data
    routes := []Route{
        {StartPoint: "A", EndPoint: "B", Duration: 15},
        {StartPoint: "B", EndPoint: "C", Duration: 20},
        {StartPoint: "C", EndPoint: "D", Duration: 10},
    }

    // Define test cases
    testCases := []struct {
        startPoint string
        endPoint   string
        expected   int
        testName   string
    }{
        {"A", "B", 15, "Single route from A to B"},
        {"A", "C", 35, "Multiple routes from A to C"},
        {"A", "D", 45, "Full route from A to D"},
        {"B", "D", 30, "Route starting from B to D"},
        {"C", "A", 0, "No route from C to A (invalid)"},
    }

    // Run each test case
    for _, tc := range testCases {
        t.Run(tc.testName, func(t *testing.T) {
            result := CalculateRouteDuration(routes, tc.startPoint, tc.endPoint)
            if result != tc.expected {
                t.Errorf("Expected %d, got %d for route from %s to %s", tc.expected, result, tc.startPoint, tc.endPoint)
            }
        })
    }
}
