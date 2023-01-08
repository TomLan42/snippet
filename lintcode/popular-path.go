package lintcode

// example of input (1B, 2B, 3C, 4A, 1A, 1C)
// output the most frequent visit path (of length 3)
// for this case, it is BAC
func PopularPath(records []string) string {

	userRecord := make(map[byte][]byte)
	pathCount := make(map[string]int)

	for _, record := range records {
		uid := record[0]
		page := record[1]

		// fmt.Println(string(uid), string(page))

		_, ok := userRecord[uid]
		if !ok {
			userRecord[uid] = make([]byte, 0)
		}

		userRecord[uid] = append(userRecord[uid], page)

		if len(userRecord[uid]) == 3 {
			pathCount[string(userRecord[uid])]++
			userRecord[uid] = userRecord[uid][1:]
		}
	}

	max := 0
	result := ""
	for k, v := range pathCount {
		if v > max {
			max = v
			result = k
		}
	}

	return result
}
