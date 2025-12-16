package dan

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Decode parses DAN text into a map[string]interface{}
// It accepts both string and []byte inputs
func Decode(text interface{}) (map[string]interface{}, error) {
	var input string
	
	switch v := text.(type) {
	case string:
		input = v
	case []byte:
		input = string(v)
	default:
		return nil, fmt.Errorf("expected string or []byte, got %T", text)
	}
	
	// Handle empty input - return empty map
	input = strings.TrimSpace(input)
	if input == "" {
		return make(map[string]interface{}), nil
	}
	
	lines := strings.Split(input, "\n")
	stack := []map[string]interface{}{make(map[string]interface{})}
	var currentTable *tableContext
	
	tableRe := regexp.MustCompile(`^(\w+):\s*table\(([^)]+)\)\s*\[$`)
	kvRe := regexp.MustCompile(`^(\w+):\s*(.+)$`)
	
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		
		// Remove comments and trim
		line = removeComments(line)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		top := stack[len(stack)-1]
		
		// Block start
		if strings.HasSuffix(line, "{") {
			key := strings.TrimSpace(line[:len(line)-1])
			newObj := make(map[string]interface{})
			top[key] = newObj
			stack = append(stack, newObj)
			continue
		}
		
		// Block end
		if line == "}" {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		
		// Table start
		if tableMatch := tableRe.FindStringSubmatch(line); tableMatch != nil {
			key := tableMatch[1]
			columnsStr := tableMatch[2]
			columns := strings.Split(columnsStr, ",")
			for i := range columns {
				columns[i] = strings.TrimSpace(columns[i])
			}
			table := []map[string]interface{}{}
			top[key] = table
			currentTable = &tableContext{
				table:   &table,
				columns: columns,
				key:     key,
				parent:  top,
			}
			continue
		}
		
		// Table end
		if line == "]" {
			if currentTable != nil {
				// Final update of the table in the map
				currentTable.parent[currentTable.key] = *currentTable.table
			}
			currentTable = nil
			continue
		}
		
		// Table row
		if currentTable != nil {
			row := make(map[string]interface{})
			parts := strings.Split(line, ",")
			for i, part := range parts {
				if i < len(currentTable.columns) {
					val := strings.TrimSpace(part)
					row[currentTable.columns[i]] = parseValue(val)
				}
			}
			*currentTable.table = append(*currentTable.table, row)
			// Update the map entry to reflect the new length
			currentTable.parent[currentTable.key] = *currentTable.table
			continue
		}
		
		// Key-value
		if kvMatch := kvRe.FindStringSubmatch(line); kvMatch != nil {
			key := kvMatch[1]
			val := kvMatch[2]
			top[key] = parseValue(strings.TrimSpace(val))
		}
	}
	
	return stack[0], nil
}

// Encode converts a map[string]interface{} to DAN format
func Encode(obj map[string]interface{}) string {
	return encodeValue(obj, 0)
}

func encodeValue(val interface{}, indent int) string {
	var lines []string
	pad := strings.Repeat("  ", indent)
	
	switch v := val.(type) {
	case map[string]interface{}:
		for key, val := range v {
			// Use type switch to handle different array types
			switch arrVal := val.(type) {
			case []map[string]interface{}:
				// Table ([]map[string]interface{})
				if len(arrVal) > 0 {
					first := arrVal[0]
					columns := make([]string, 0, len(first))
					for col := range first {
						columns = append(columns, col)
					}
					lines = append(lines, fmt.Sprintf("%s%s: table(%s) [", pad, key, strings.Join(columns, ", ")))
					for _, row := range arrVal {
						rowValues := make([]string, len(columns))
						for i, col := range columns {
							rowValues[i] = serializeValue(row[col])
						}
						lines = append(lines, pad+"  "+strings.Join(rowValues, ", "))
					}
					lines = append(lines, pad+"]")
				} else {
					// Empty table
					lines = append(lines, fmt.Sprintf("%s%s: table() [", pad, key))
					lines = append(lines, pad+"]")
				}
			case []interface{}:
				// Array of interface{} - could be table or regular array
				if len(arrVal) > 0 {
					if first, ok := arrVal[0].(map[string]interface{}); ok {
						// Table (from []interface{} containing maps)
						columns := make([]string, 0, len(first))
						for col := range first {
							columns = append(columns, col)
						}
						lines = append(lines, fmt.Sprintf("%s%s: table(%s) [", pad, key, strings.Join(columns, ", ")))
						for _, row := range arrVal {
							if rowMap, ok := row.(map[string]interface{}); ok {
								rowValues := make([]string, len(columns))
								for i, col := range columns {
									rowValues[i] = serializeValue(rowMap[col])
								}
								lines = append(lines, pad+"  "+strings.Join(rowValues, ", "))
							}
						}
						lines = append(lines, pad+"]")
					} else {
						// Regular array
						lines = append(lines, fmt.Sprintf("%s%s: %s", pad, key, serializeValue(val)))
					}
				} else {
					// Empty array
					lines = append(lines, fmt.Sprintf("%s%s: %s", pad, key, serializeValue(val)))
				}
			case map[string]interface{}:
				// Nested object
				lines = append(lines, fmt.Sprintf("%s%s {", pad, key))
				nestedLines := encodeValue(arrVal, indent+1)
				if nestedLines != "" {
					lines = append(lines, nestedLines)
				}
				lines = append(lines, pad+"}")
			default:
				// Regular value (string, number, bool, etc.)
				lines = append(lines, fmt.Sprintf("%s%s: %s", pad, key, serializeValue(val)))
			}
		}
	default:
		return ""
	}
	
	return strings.Join(lines, "\n")
}

type tableContext struct {
	table   *[]map[string]interface{}
	columns []string
	key     string
	parent  map[string]interface{}
}

func removeComments(line string) string {
	commentIndex1 := strings.Index(line, "#")
	commentIndex2 := strings.Index(line, "//")
	
	cutIndex := -1
	if commentIndex1 >= 0 && commentIndex2 >= 0 {
		if commentIndex1 < commentIndex2 {
			cutIndex = commentIndex1
		} else {
			cutIndex = commentIndex2
		}
	} else if commentIndex1 >= 0 {
		cutIndex = commentIndex1
	} else if commentIndex2 >= 0 {
		cutIndex = commentIndex2
	}
	
	if cutIndex >= 0 {
		return line[:cutIndex]
	}
	return line
}

func parseValue(val string) interface{} {
	val = strings.TrimSpace(val)
	
	// Boolean
	if val == "true" {
		return true
	}
	if val == "false" {
		return false
	}
	
	// String (quoted)
	if len(val) >= 2 && val[0] == '"' && val[len(val)-1] == '"' {
		return val[1 : len(val)-1]
	}
	
	// Number
	if num, err := strconv.ParseFloat(val, 64); err == nil {
		return num
	}
	
	// Array
	if len(val) >= 2 && val[0] == '[' && val[len(val)-1] == ']' {
		content := strings.TrimSpace(val[1 : len(val)-1])
		if content == "" {
			return []interface{}{}
		}
		parts := strings.Split(content, ",")
		result := make([]interface{}, len(parts))
		for i, part := range parts {
			result[i] = parseValue(strings.TrimSpace(part))
		}
		return result
	}
	
	// Default: return as string
	return val
}

func serializeValue(val interface{}) string {
	switch v := val.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return fmt.Sprintf(`"%s"`, v)
	case float64:
		// Check if it's an integer
		if v == float64(int64(v)) {
			return fmt.Sprintf("%.0f", v)
		}
		return fmt.Sprintf("%g", v)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case []interface{}:
		parts := make([]string, len(v))
		for i, item := range v {
			parts[i] = serializeValue(item)
		}
		return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
	default:
		return fmt.Sprintf("%v", v)
	}
}
