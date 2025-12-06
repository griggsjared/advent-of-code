#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Get the script directory (project root)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

print_usage() {
    echo -e "${BLUE}Advent of Code Go Runner${NC}"
    echo ""
    echo "Usage:"
    echo "  ./aoc.sh create [YEAR] [DAY]"
    echo "  ./aoc.sh [test|run] [YEAR] [DAY]"
    echo "  ./aoc.sh [test|run] [YEAR] all"
    echo "  ./aoc.sh [test|run] all"
    echo ""
    echo "Examples:"
    echo "  ./aoc.sh create 2024 1         # Create files for day 1 of 2024"
    echo "  ./aoc.sh test 2024 1           # Test day 1 of 2024"
    echo "  ./aoc.sh run 2024 1            # Run day 1 of 2024"
    echo "  ./aoc.sh test 2024 all         # Test all days of 2024"
    echo "  ./aoc.sh run all               # Run all available days"
    echo "  ./aoc.sh test                  # Interactive mode"
    echo ""
}

# Function to create a new day from starter template
create_day() {
    local year=$1
    local day=$2
    local day_padded=$(printf "%02d" $day)
    local day_dir="$SCRIPT_DIR/$year/$day_padded"
    
    # Validate year
    if ! [[ "$year" =~ ^[0-9]{4}$ ]]; then
        echo -e "${RED}Error: Year must be a 4-digit number${NC}"
        return 1
    fi
    
    # Validate day
    if ! [[ "$day" =~ ^[0-9]+$ ]] || [ "$day" -lt 1 ] || [ "$day" -gt 25 ]; then
        echo -e "${RED}Error: Day must be a number between 1 and 25${NC}"
        return 1
    fi
    
    # Check if directory already exists
    if [ -d "$day_dir" ]; then
        echo -e "${YELLOW}Directory $year/$day_padded already exists!${NC}"
        read -p "Overwrite? (y/N): " confirm
        if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
            echo -e "${CYAN}Cancelled${NC}"
            return 1
        fi
    fi
    
    # Create directory
    mkdir -p "$day_dir"
    
    # Copy and modify main.go
    echo -e "${CYAN}Creating $year/$day_padded/main.go...${NC}"
    sed -e "s/day  = 9999/day  = $day/" \
        -e "s/year = 9999/year = $year/" \
        "$SCRIPT_DIR/starter/main.go" > "$day_dir/main.go"
    
    # Copy and modify main_test.go
    echo -e "${CYAN}Creating $year/$day_padded/main_test.go...${NC}"
    sed "s|github.com/griggsjared/advent-of-code/go/year/day|github.com/griggsjared/advent-of-code/go/$year/$day_padded|" \
        "$SCRIPT_DIR/starter/main_test.go" > "$day_dir/main_test.go"
    
    echo -e "${GREEN}Successfully created files for $year day $day!${NC}"
    echo -e "${YELLOW}Don't forget to create $year/$day_padded/input.txt with your puzzle input${NC}"
    return 0
}

# Function to check if required files exist
check_files() {
    local year=$1
    local day=$2
    local mode=$3
    
    local day_padded=$(printf "%02d" $day)
    local day_dir="$SCRIPT_DIR/$year/$day_padded"
    
    if [ ! -d "$day_dir" ]; then
        echo -e "${RED}Error: Directory $year/$day_padded not found!${NC}"
        return 1
    fi
    
    if [ ! -f "$day_dir/main.go" ]; then
        echo -e "${RED}Error: $year/$day_padded/main.go not found!${NC}"
        return 1
    fi
    
    if [ "$mode" = "test" ]; then
        if [ ! -f "$day_dir/main_test.go" ]; then
            echo -e "${RED}Error: $year/$day_padded/main_test.go not found!${NC}"
            return 1
        fi
    fi
    
    if [ "$mode" = "run" ]; then
        if [ ! -f "$day_dir/input.txt" ]; then
            echo -e "${RED}Error: $year/$day_padded/input.txt not found!${NC}"
            echo -e "${YELLOW}Please create an input file at $year/$day_padded/input.txt${NC}"
            return 1
        fi
    fi
    
    return 0
}

# Function to run tests for a specific day
run_test() {
    local year=$1
    local day=$2
    local day_padded=$(printf "%02d" $day)
    
    if ! check_files "$year" "$day" "test"; then
        return 1
    fi
    
    echo -e "${CYAN}Testing $year Day $day...${NC}"
    cd "$SCRIPT_DIR/$year/$day_padded" && go test -v
    local result=$?
    cd "$SCRIPT_DIR"
    return $result
}

# Function to run solution for a specific day
run_solution() {
    local year=$1
    local day=$2
    local day_padded=$(printf "%02d" $day)
    
    if ! check_files "$year" "$day" "run"; then
        return 1
    fi
    
    echo -e "${CYAN}Running $year Day $day...${NC}"
    cd "$SCRIPT_DIR" && go run "$year/$day_padded/main.go"
    local result=$?
    return $result
}

# Function to get all available years
get_available_years() {
    for year_dir in "$SCRIPT_DIR"/[0-9][0-9][0-9][0-9]; do
        if [ -d "$year_dir" ]; then
            basename "$year_dir"
        fi
    done | sort
}

# Function to get all available days for a year
get_available_days() {
    local year=$1
    for day_dir in "$SCRIPT_DIR/$year"/[0-9][0-9]; do
        if [ -d "$day_dir" ]; then
            # Remove leading zero for output
            local day=$(basename "$day_dir")
            echo $((10#$day))
        fi
    done | sort -n
}

# Function to run all days for a year
run_all_days_for_year() {
    local mode=$1
    local year=$2
    
    local days=$(get_available_days "$year")
    if [ -z "$days" ]; then
        echo -e "${YELLOW}No days found for year $year${NC}"
        return 1
    fi
    
    local total=0
    local passed=0
    local failed=0
    
    for day in $days; do
        ((total++))
        echo ""
        if [ "$mode" = "test" ]; then
            if run_test "$year" "$day"; then
                ((passed++))
            else
                ((failed++))
            fi
        else
            if run_solution "$year" "$day"; then
                ((passed++))
            else
                ((failed++))
            fi
        fi
    done
    
    echo ""
    echo -e "${BLUE}========================================${NC}"
    echo -e "${BLUE}Summary for $year:${NC}"
    echo -e "${GREEN}Passed: $passed${NC}"
    if [ $failed -gt 0 ]; then
        echo -e "${RED}Failed: $failed${NC}"
    fi
    echo -e "${BLUE}Total: $total${NC}"
    echo -e "${BLUE}========================================${NC}"
}

# Function to run all available years and days
run_all() {
    local mode=$1
    
    local years=$(get_available_years)
    if [ -z "$years" ]; then
        echo -e "${YELLOW}No years found${NC}"
        return 1
    fi
    
    for year in $years; do
        echo -e "${BLUE}=======================================${NC}"
        echo -e "${BLUE}Year $year${NC}"
        echo -e "${BLUE}=======================================${NC}"
        run_all_days_for_year "$mode" "$year"
        echo ""
    done
}

# Interactive mode
interactive_mode() {
    local mode=$1
    
    # Get available years
    local years=$(get_available_years)
    if [ -z "$years" ]; then
        echo -e "${RED}No years found in the project!${NC}"
        return 1
    fi
    
    echo -e "${BLUE}Available years:${NC}"
    echo "$years" | while read -r y; do
        echo "  - $y"
    done
    echo "  - all"
    echo ""
    
    read -p "Enter year (or 'all'): " year
    
    if [ "$year" = "all" ]; then
        run_all "$mode"
        return
    fi
    
    # Validate year
    if ! echo "$years" | grep -q "^$year$"; then
        echo -e "${RED}Year $year not found!${NC}"
        return 1
    fi
    
    # Get available days for the year
    local days=$(get_available_days "$year")
    if [ -z "$days" ]; then
        echo -e "${YELLOW}No days found for year $year${NC}"
        return 1
    fi
    
    echo -e "${BLUE}Available days for $year:${NC}"
    echo "$days" | while read -r d; do
        echo "  - $d"
    done
    echo "  - all"
    echo ""
    
    read -p "Enter day (or 'all'): " day
    
    if [ "$day" = "all" ]; then
        run_all_days_for_year "$mode" "$year"
        return
    fi
    
    # Validate day
    if ! echo "$days" | grep -q "^$day$"; then
        echo -e "${RED}Day $day not found for year $year!${NC}"
        return 1
    fi
    
    if [ "$mode" = "test" ]; then
        run_test "$year" "$day"
    else
        run_solution "$year" "$day"
    fi
}

# Main script logic
main() {
    if [ $# -eq 0 ]; then
        print_usage
        echo -e "${YELLOW}No arguments provided. Starting interactive mode...${NC}"
        echo ""
        read -p "Choose mode (create/test/run): " mode
        
        if [ "$mode" != "create" ] && [ "$mode" != "test" ] && [ "$mode" != "run" ]; then
            echo -e "${RED}Invalid mode. Must be 'create', 'test', or 'run'${NC}"
            exit 1
        fi
        
        if [ "$mode" = "create" ]; then
            read -p "Enter year: " year
            read -p "Enter day: " day
            create_day "$year" "$day"
            exit $?
        fi
        
        interactive_mode "$mode"
        exit $?
    fi
    
    local mode=$1
    
    # Handle create command
    if [ "$mode" = "create" ]; then
        if [ $# -ne 3 ]; then
            echo -e "${RED}Error: create requires YEAR and DAY arguments${NC}"
            print_usage
            exit 1
        fi
        create_day "$2" "$3"
        exit $?
    fi
    
    if [ "$mode" != "test" ] && [ "$mode" != "run" ]; then
        echo -e "${RED}Invalid mode. Must be 'create', 'test', or 'run'${NC}"
        print_usage
        exit 1
    fi
    
    # No year/day specified - interactive mode
    if [ $# -eq 1 ]; then
        interactive_mode "$mode"
        exit $?
    fi
    
    local year=$2
    
    # Run all years
    if [ "$year" = "all" ]; then
        run_all "$mode"
        exit $?
    fi
    
    # No day specified - interactive mode for day selection
    if [ $# -eq 2 ]; then
        # Check if year exists
        if [ ! -d "$SCRIPT_DIR/$year" ]; then
            echo -e "${RED}Year $year not found!${NC}"
            exit 1
        fi
        
        echo -e "${BLUE}Available days for $year:${NC}"
        get_available_days "$year" | while read -r d; do
            echo "  - $d"
        done
        echo "  - all"
        echo ""
        
        read -p "Enter day (or 'all'): " day
        
        if [ "$day" = "all" ]; then
            run_all_days_for_year "$mode" "$year"
            exit $?
        fi
    else
        local day=$3
        
        # Run all days for a year
        if [ "$day" = "all" ]; then
            run_all_days_for_year "$mode" "$year"
            exit $?
        fi
    fi
    
    # Run specific day
    if [ "$mode" = "test" ]; then
        run_test "$year" "$day"
    else
        run_solution "$year" "$day"
    fi
}

main "$@"
