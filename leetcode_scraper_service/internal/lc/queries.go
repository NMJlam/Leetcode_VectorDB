package lc 

/*
--------------------------------------------
Slugs: names of the leetcode questions 
--------------------------------------------
*/
var SlugQuery = `
	query problemsetQuestionList($limit: Int, $skip: Int) {
		problemsetQuestionList: problemsetQuestionListV2(
			categorySlug: ""
			limit: $limit
			skip: $skip
		) {
			questions: questions {
				acRate
				difficulty
				title
				titleSlug
			}
		}
	}
`

type SlugResponse struct {
		ProblemsetQuestionList struct {
			Questions []struct {
				Difficulty string  `json:"difficulty"`
				Title      string  `json:"title"`
				TitleSlug  string  `json:"titleSlug"`
				AcRate     float64 `json:"acRate"`
			} `json:"questions"`
		} `json:"problemsetQuestionList"`
	}
