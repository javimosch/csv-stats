package main
import ("encoding/csv";"fmt";"math";"os";"strconv";"strings")
func main() {
	r := csv.NewReader(os.Stdin)
	rows, err := r.ReadAll()
	if err != nil { fmt.Fprintln(os.Stderr,"CSV error:",err); os.Exit(1) }
	if len(rows) < 2 { fmt.Println("[]"); return }
	headers := rows[0]
	cols := len(headers)
	data := rows[1:]
	
	fmt.Println("[")
	for i := 0; i < cols; i++ {
		if i > 0 { fmt.Println(",") }
		count := len(data)
		var nums []float64
		strings := make(map[string]int)
		nulls := 0
		for _, row := range data {
			if i >= len(row) { nulls++; continue }
			v := strings.TrimSpace(row[i])
			if v == "" { nulls++; continue }
			if n, err := strconv.ParseFloat(v, 64); err == nil { nums = append(nums, n) }
			strings[v]++
		}
		fmt.Printf(`{"column":"%s","total":%d,"nulls":%d,"non_null":%d,"unique_strings":%d`, headers[i], count, nulls, count-nulls, len(strings))
		if len(nums) > 0 {
			min, max, sum := nums[0], nums[0], 0.0
			for _, n := range nums { if n < min { min = n }; if n > max { max = n }; sum += n }
			avg := sum / float64(len(nums))
			// median
			sortNums(nums)
			median := nums[len(nums)/2]
			fmt.Printf(`,"numeric":true,"min":%v,"max":%v,"avg":%.2f,"median":%v,"count":%d`, min, max, avg, median, len(nums))
		} else {
			fmt.Printf(`,"numeric":false`)
		}
		fmt.Printf(`}`)
	}
	fmt.Println("\n]")
}
func sortNums(n []float64) {
	for i := 0; i < len(n); i++ {
		for j := i+1; j < len(n); j++ {
			if n[j] < n[i] { n[i], n[j] = n[j], n[i] }
		}
	}
}
