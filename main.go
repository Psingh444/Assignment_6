// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// // Function to identify the browser from the User-Agent string
// func getBrowserName(userAgent string) string {
// 	// Map of browser signatures to browser names
// 	browserSignatures := map[string]string{
// 		"Edg":     "Microsoft Edge",
// 		"OPR":     "Opera",
// 		"Opera":   "Opera",
// 		"Firefox": "Mozilla Firefox",
// 		"Chrome":  "Google Chrome",
// 		"Safari":  "Apple Safari",
// 		"MSIE":    "Internet Explorer",
// 		"Trident": "Internet Explorer",
// 	}

// 	// Check for the presence of each browser signature in the User-Agent string
// 	for signature, name := range browserSignatures {
// 		if strings.Contains(userAgent, signature) {
// 			return name
// 		}
// 	}

// 	// If no known browser is found, return "Unknown Browser"
// 	return "Unknown Browser"
// }

// // HTTP handler function to process incoming requests
// func handler(w http.ResponseWriter, r *http.Request) {
// 	userAgent := r.Header.Get("User-Agent")
// 	browserName := getBrowserName(userAgent)
// 	log.Printf("User-Agent: %s", userAgent) // Log the User-Agent string
// 	response := fmt.Sprintf(`<html><body style="text-align:center; padding:50px;">
//                               <h1>Browser Name: %s</h1>
//                               </body></html>`, browserName)
// 	fmt.Fprint(w, response)
// }

// // Main function to start the HTTP server
// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Starting server on port 80...")
// 	if err := http.ListenAndServe(":80", nil); err != nil {
// 		log.Fatalf("Error starting server: %s\n", err)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Function to identify the browser from the User-Agent string
func getBrowserName(userAgent string) string {
	// Check for known browser signatures with more specific rules
	if strings.Contains(userAgent, "Edg/") {
		return "Microsoft Edge"
	} else if strings.Contains(userAgent, "OPR/") || strings.Contains(userAgent, "Opera") {
		return "Opera"
	} else if strings.Contains(userAgent, "Firefox/") {
		return "Mozilla Firefox"
	} else if strings.Contains(userAgent, "Chrome/") && strings.Contains(userAgent, "Safari/") && !strings.Contains(userAgent, "Edg/") {
		return "Google Chrome"
	} else if strings.Contains(userAgent, "Safari/") && !strings.Contains(userAgent, "Chrome/") {
		return "Apple Safari"
	} else if strings.Contains(userAgent, "MSIE ") || strings.Contains(userAgent, "Trident/") {
		return "Internet Explorer"
	}

	// If no known browser is found, return "Unknown Browser"
	return "Unknown Browser"
}

// Function to get background color based on browser name
func getBackgroundColor(browserName string) string {
	browserColors := map[string]string{
		"Microsoft Edge":    "#0078D7", // Blue
		"Opera":             "#FF1B2D", // Red
		"Mozilla Firefox":   "#FF7139", // Orange
		"Google Chrome":     "#FFC0CB", // Pink
		"Apple Safari":      "#1E90FF", // Dodger Blue
		"Internet Explorer": "#1E90FF", // Dodger Blue (Same as Safari for simplicity)
		"Unknown Browser":   "#FFFFFF", // White
	}

	if color, exists := browserColors[browserName]; exists {
		return color
	}
	return "#FFFFFF" // Default to white if browser not found
}

// HTTP handler function to process incoming requests
func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	browserName := getBrowserName(userAgent)
	backgroundColor := getBackgroundColor(browserName)
	log.Printf("User-Agent: %s", userAgent) // Log the User-Agent string

	response := fmt.Sprintf(`
		<html>
			<body style="text-align:center; padding:50px; background-color:%s;">
				<h1>Browser Name: %s</h1>
			</body>
		</html>`, backgroundColor, browserName)
	fmt.Fprint(w, response)
}

// Main function to start the HTTP server
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
