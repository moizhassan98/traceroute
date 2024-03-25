package traceroute

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// Route represents information about a single hop in the trace route
type Route struct {
	Hop     int
	Address string
	Time1   string
	Time2   string
	Time3   string
}

// Trace represents the trace route result
type Trace struct {
	Destination     string
	Routes          []Route
	traceSuccessful bool
}

func GetHops() (Trace, error) {
	// Command to execute
	cmd := exec.Command("tracert", "google.com")
	var e error

	// Run the command and capture output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		e = err
		return Trace{}, err
	}

	// Convert output to string
	outputStr := string(output)
	fmt.Println("RAW Output: ")
	fmt.Println(outputStr)

	// Parse output
	var trace Trace = parseTrace(outputStr, false)

	return trace, e
}

func GetHopsJSON(trace Trace) ([]byte, error) {
	jsonData, err := json.Marshal(trace)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return []byte{}, err
	}
	return jsonData, nil
}

func parseTrace(output string, includeTimeouts bool) Trace {
	lines := strings.Split(output, "\n")
	destination := ""
	var routes []Route

	for _, line := range lines {
		if strings.Contains(line, "Tracing route to") {
			destination = strings.TrimSpace(strings.Split(line, "Tracing route to")[1])
		}
		if strings.Contains(line, "ms") {
			fields := strings.Fields(line)
			// hop := strings.Trim(fields[0], "[]")
			address := strings.Join(fields[7:], " ")
			times := fields[1:7]
			routes = append(routes, Route{
				Hop:     len(routes) + 1,
				Address: address,
				Time1:   times[0] + times[1],
				Time2:   times[2] + times[3],
				Time3:   times[4] + times[5],
			})
		} else if strings.Contains(line, "*") && includeTimeouts {
			fields := strings.Fields(line)
			address := strings.Join(fields[4:], " ")
			times := fields[1:4]
			routes = append(routes, Route{
				Hop:     len(routes) + 1,
				Address: address,
				Time1:   times[0],
				Time2:   times[1],
				Time3:   times[2],
			})
		}

	}

	traceSuccessful := false
	fmt.Println("Last Line: ", lines[len(lines)-1])
	if strings.Contains(lines[len(lines)-1], "ms") {
		fmt.Println("TRACE SUCCESSFUL")
		traceSuccessful = true
	}

	return Trace{
		Destination:     destination,
		Routes:          routes,
		traceSuccessful: traceSuccessful,
	}
}
