package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"dan-cli/dan"
)

func main() {
	decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)
	encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)
	prettyCmd := flag.NewFlagSet("pretty", flag.ExitOnError)
	
	decodeJSON := decodeCmd.Bool("json", false, "Output as JSON")
	encodeJSON := encodeCmd.Bool("json", false, "Input as JSON")
	prettyJSON := prettyCmd.Bool("json", false, "Output as JSON")
	
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	
	command := os.Args[1]
	
	switch command {
	case "decode":
		decodeCmd.Parse(os.Args[2:])
		filePath := ""
		if len(decodeCmd.Args()) > 0 {
			filePath = decodeCmd.Args()[0]
		}
		if err := runDecode(*decodeJSON, filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "encode":
		encodeCmd.Parse(os.Args[2:])
		filePath := ""
		if len(encodeCmd.Args()) > 0 {
			filePath = encodeCmd.Args()[0]
		}
		if err := runEncode(*encodeJSON, filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "pretty":
		prettyCmd.Parse(os.Args[2:])
		filePath := ""
		if len(prettyCmd.Args()) > 0 {
			filePath = prettyCmd.Args()[0]
		}
		if err := runPretty(*prettyJSON, filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("DAN (Data Advanced Notation) CLI Tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  dan decode [flags]              Decode DAN to JSON")
	fmt.Println("  dan encode [flags]              Encode JSON to DAN")
	fmt.Println("  dan pretty [flags]              Pretty-print DAN")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -json                          Use JSON format for input/output")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  dan decode file.dan            Decode DAN file to JSON")
	fmt.Println("  dan decode < file.dan          Decode DAN from stdin")
	fmt.Println("  dan encode -json file.json     Encode JSON file to DAN")
	fmt.Println("  dan pretty file.dan            Pretty-print DAN file")
	fmt.Println("  echo 'name: \"test\"' | dan decode")
}

func runDecode(outputJSON bool, filePath string) error {
	input, err := readInput(filePath)
	if err != nil {
		return err
	}
	
	result, err := dan.Decode(input)
	if err != nil {
		return err
	}
	
	if outputJSON {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(result)
	}
	
	// Output as DAN (pretty-printed)
	output := dan.Encode(result)
	fmt.Println(output)
	return nil
}

func runEncode(inputJSON bool, filePath string) error {
	input, err := readInput(filePath)
	if err != nil {
		return err
	}
	
	var obj map[string]interface{}
	
	if inputJSON {
		if err := json.Unmarshal([]byte(input), &obj); err != nil {
			return fmt.Errorf("invalid JSON: %w", err)
		}
	} else {
		// Assume input is DAN
		var err error
		obj, err = dan.Decode(input)
		if err != nil {
			return err
		}
	}
	
	output := dan.Encode(obj)
	fmt.Println(output)
	return nil
}

func runPretty(outputJSON bool, filePath string) error {
	input, err := readInput(filePath)
	if err != nil {
		return err
	}
	
	result, err := dan.Decode(input)
	if err != nil {
		return err
	}
	
	if outputJSON {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(result)
	}
	
	// Pretty-print as DAN
	output := dan.Encode(result)
	fmt.Println(output)
	return nil
}

func readInput(filePath string) (string, error) {
	// If file path is provided, read from file
	if filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
		}
		return string(data), nil
	}
	
	// Check if there's data in stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped in
		var input strings.Builder
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input.WriteString(scanner.Text())
			input.WriteString("\n")
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return input.String(), nil
	}
	
	// Read from stdin interactively
	var input strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter DAN content (Ctrl+D or Ctrl+Z to finish):")
	for scanner.Scan() {
		input.WriteString(scanner.Text())
		input.WriteString("\n")
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return "", err
	}
	return input.String(), nil
}
