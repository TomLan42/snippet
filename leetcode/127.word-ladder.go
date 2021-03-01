/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {

	// what a interesting problem
	// makes me think of dp problems of eidt distance

	// don't have to use all the words from the dict

	// we can do bfs, first found is the shortest

	visited := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, beginWord)

	result := 1

	for _, word := range wordList {
		visited[word] = false
	}

	if _, found := visited[endWord]; !found {
		return 0
	}

	for len(queue) > 0 {
		// search neighbor from the dict that has not been visited
		l := len(queue)
		result++
		for i := 0; i < l; i++ {
			for _, word := range wordList {
				if isPath(queue[0], word) && !visited[word] {
					//enqueue
					queue = append(queue, word)
					//mark visited
					visited[word] = true
					if word == endWord {
						return result
					}
				}
			}
			// dequeue
			queue = queue[1:]
		}
	}

	return 0
}

func isPath(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
	}
	if count == 1 {
		return true
	}
	return false
}

// @lc code=end

