package main
import ("encoding/csv";"fmt";"os")
func main() {
	r := csv.NewReader(os.Stdin)
	rows, err := r.ReadAll()
	if err != nil { fmt.Fprintln(os.Stderr,"CSV error:",err); os.Exit(1) }
	if len(rows) < 2 { fmt.Println("[]"); return }
	headers := rows[0]; cols := len(headers); data := rows[1:]
	fmt.Println("[")
	for i := 0; i < cols; i++ {
		if i > 0 { fmt.Println(",") }
		var nums []float64; strVals := make(map[string]int); nulls := 0
		for _, row := range data {
			if i >= len(row) { nulls++; continue }
			v := row[i]
			if v == "" { nulls++; continue }
			if n, err := fmt.Sscanf(v, "%f", new(float64)); err == nil && n == 1 {
				var f float64; fmt.Sscanf(v, "%f", &f); nums = append(nums, f)
			}
			strVals[v]++
		}
		fmt.Printf(`{"column":"%s","total":%d,"nulls":%d,"non_null":%d,"unique_strings":%d`, headers[i], len(data), nulls, len(data)-nulls, len(strVals))
		if len(nums) > 0 {
			min, max, sum := nums[0], nums[0], 0.0
			for _, n := range nums { if n < min { min = n }; if n > max { max = n }; sum += n }
			avg := sum / float64(len(nums))
			// median
			for a := 0; a < len(nums); a++ { for b := a+1; b < len(nums); b++ { if nums[b] < nums[a] { nums[a], nums[b] = nums[b], nums[a] } } }
			med := nums[len(nums)/2]
			fmt.Printf(`,"numeric":true,"min":%.2f,"max":%.2f,"avg":%.2f,"median":%.2f,"count":%d`, min, max, avg, med, len(nums))
		} else { fmt.Printf(`,"numeric":false`) }
		fmt.Printf(`}`)
	}
	fmt.Println("\n]")
}
