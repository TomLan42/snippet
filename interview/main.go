package main

import "fmt"

type Log struct {
	user string
	page string
}

func main() {
	// log := []*Log{
	// 	{user: "1", page: "A"},
	// 	{user: "2", page: "B"},
	// 	{user: "1", page: "B"},
	// 	{user: "1", page: "D"},
	// 	{user: "2", page: "A"},
	// 	{user: "3", page: "B"},
	// 	{user: "3", page: "D"},
	// 	{user: "1", page: "C"},
	// 	{user: "3", page: "C"},
	// 	{user: "1", page: "C"},
	// 	{user: "2", page: "C"},
	// 	{user: "3", page: "B"},
	// 	{user: "1", page: "A"},
	// 	{user: "3", page: "C"},
	// }

	// log1 := []*Log{
	// 	{user: "1", page: "A"},
	// 	{user: "2", page: "B"},
	// 	{user: "1", page: "B"},
	// 	{user: "1", page: "D"},
	// 	{user: "2", page: "A"},
	// 	{user: "3", page: "B"},
	// 	{user: "3", page: "D"},
	// 	{user: "1", page: "C"},
	// 	{user: "3", page: "C"},
	// 	{user: "1", page: "C"},
	// 	{user: "2", page: "C"},
	// 	{user: "3", page: "B"},
	// 	{user: "1", page: "A"},
	// 	{user: "3", page: "C"},
	// 	{user: "4", page: "B"},
	// 	{user: "4", page: "A"},
	// 	{user: "4", page: "C"},
	// }

	log2 := []*Log{
		{user: "1", page: "A"},
		{user: "2", page: "B"},
		{user: "1", page: "B"},
		{user: "1", page: "D"},
		{user: "2", page: "A"},
		{user: "2", page: "A"},
	}

	fmt.Println(solution(log2))

}

// 1:[A B C] => "ABC"++
// 1:[B C D] => "BCD"++

func solution(records []*Log) []string {

	userMap := make(map[string][]string)
	countMap := make(map[string]int)

	for _, record := range records {
		_, ok := userMap[record.user]
		if !ok {
			userMap[record.user] = make([]string, 0)
		}
		userMap[record.user] = append(userMap[record.user], record.page)
		if len(userMap[record.user]) >= 3 {
			countMap[join(userMap[record.user])]++
			userMap[record.user] = userMap[record.user][1:]
		}
	}

	maxCount := 0
	totalCount := 0
	for _, v := range countMap {
		if v > maxCount {
			maxCount = v
		}
		totalCount += v
	}

	reverseMap := make(map[int][]string)

	for k, v := range countMap {
		if reverseMap[v] == nil {
			reverseMap[v] = make([]string, 0)
		}
		reverseMap[v] = append(reverseMap[v], k)
	}

	if len(reverseMap[maxCount]) == len(countMap) {
		return []string{}
	}

	return reverseMap[maxCount]
}

func join(inputs []string) string {
	output := ""

	for _, input := range inputs {
		output += input
	}

	return output
}


O(n)