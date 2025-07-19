package lc 

/*
--------------------------------------------
Slugs: names of the leetcode questions 
--------------------------------------------
*/
const SlugQuery = `
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
/*
--------------------------------------------
Question Descriptions: Problem Descriptions  
--------------------------------------------
*/
const DescriptionQuery = `query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    content 
    difficulty
    similarQuestions
    codeSnippets {
        lang 
        langSlug
        code
    }
    solution {
        canSeeDetail
        paidOnly
        hasVideoSolution
        paidOnlyVideo
    }
  }
}
`

type DescriptionResponse struct {
    Question struct {
        QuestionID       string `json:"questionId"`
        Content          string `json:"content"`
        Difficulty       string `json:"difficulty"`
        SimilarQuestions string `json:"similarQuestions"`
        CodeSnippets     []struct {
            Lang     string `json:"lang"`
            LangSlug string `json:"langSlug"`
            Code     string `json:"code"`
        } `json:"codeSnippets"`
        Solution struct {
            CanSeeDetail     bool `json:"canSeeDetail"`
            PaidOnly         bool `json:"paidOnly"`
            HasVideoSolution bool `json:"hasVideoSolution"`
            PaidOnlyVideo    bool `json:"paidOnlyVideo"`
        } `json:"solution"`
    } `json:"question"`
}

