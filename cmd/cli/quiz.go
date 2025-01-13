package cli

import (
	"context"
	"fmt"
	"github.com/manifoldco/promptui"
	openapi "go-quiz/pkg/client"
	"log"
)

type QuizContext struct {
	ApiClient         *openapi.APIClient
	Context           context.Context
	SelectedQuizId    int32
	QuestionAnswerMap map[int32]int32
}

func NewQuizContext(openApiClient *openapi.APIClient, ctx context.Context) *QuizContext {
	return &QuizContext{
		ApiClient:         openApiClient,
		Context:           ctx,
		SelectedQuizId:    0,
		QuestionAnswerMap: make(map[int32]int32),
	}
}

func (q *QuizContext) AuthenticateUser(username string, password string) error {
	payload := openapi.ApiCreateTokenPayload{
		Email:    "demo@quiz.com",
		Password: "password",
	}
	request := q.ApiClient.AuthAPI.AuthTokenPost(q.Context).Payload(payload)

	response, _, err := q.ApiClient.AuthAPI.AuthTokenPostExecute(request)
	if err != nil {
		return err
	}

	log.Printf("Demo user authenticated with token %v", *response.Data.Token)

	apiKeys := map[string]openapi.APIKey{
		"BearerAuth": {Key: *response.Data.Token, Prefix: "Bearer"},
	}

	authenticatedContext := context.WithValue(q.Context, openapi.ContextAPIKeys, apiKeys)
	q.Context = authenticatedContext
	return nil
}

func (q *QuizContext) SelectQuiz() error {
	request := q.ApiClient.QuizzesAPI.QuizzesGet(q.Context)
	response, _, err := q.ApiClient.QuizzesAPI.QuizzesGetExecute(request)

	if err != nil {
		return err
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Title }}",
		Inactive: "  {{ .Title }}",
		Selected: "Starting quiz {{ .Title }}",
		Details: `
-------- Quiz --------
{{ "Title:" | faint }}	{{ .Title }}
{{ "Description:" | faint }}	{{ .Description }}
`,
	}

	prompt := promptui.Select{
		Label:     "Which quiz do you want to take",
		Items:     response.Data,
		Templates: templates,
		Size:      4,
	}

	index, _, err := prompt.Run()
	quiz := response.Data[index]
	q.SelectedQuizId = *quiz.Id

	if err != nil {
		return err
	}

	return nil
}

func (q *QuizContext) AnswerQuestions() error {
	request := q.ApiClient.QuizzesAPI.QuizzesQuizIdGet(q.Context, q.SelectedQuizId)
	response, _, err := q.ApiClient.QuizzesAPI.QuizzesQuizIdGetExecute(request)

	if err != nil {
		return err
	}

	for index, question := range response.Data.Questions {
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "> {{ . }}",
			Inactive: "  {{ . }}",
			Selected: fmt.Sprintf("Question #%d answer saved. moving to next question", index+1),
		}

		prompt := promptui.Select{
			Label:     fmt.Sprintf("Question #%d: %s", index+1, *question.Value),
			Items:     response.Data.Questions[index].Answers,
			Templates: templates,
			Size:      4,
		}

		answerIndex, _, err := prompt.Run()

		if err != nil {
			return err
		}

		q.QuestionAnswerMap[*question.Id] = int32(answerIndex + 1)
	}

	return nil
}

func (q *QuizContext) SubmitAnswers() error {
	var answers []openapi.ApiSubmitQuizAnswersPayloadAnswersInner

	for questionId, answerIndex := range q.QuestionAnswerMap {
		answers = append(answers, openapi.ApiSubmitQuizAnswersPayloadAnswersInner{
			QuestionId:  questionId,
			AnswerIndex: answerIndex,
		})
	}

	payload := openapi.ApiSubmitQuizAnswersPayload{
		Answers: answers,
	}

	request := q.ApiClient.QuizzesAPI.QuizzesQuizIdSubmitPost(q.Context, q.SelectedQuizId).Payload(payload)
	_, err := q.ApiClient.QuizzesAPI.QuizzesQuizIdSubmitPostExecute(request)

	if err != nil {
		return fmt.Errorf("failed to submit quiz answers %v", err)
	}

	return nil
}

func (q *QuizContext) DisplayResults() error {
	request := q.ApiClient.QuizzesAPI.QuizzesQuizIdResultsGet(q.Context, q.SelectedQuizId)
	response, _, err := q.ApiClient.QuizzesAPI.QuizzesQuizIdResultsGetExecute(request)

	if err != nil {
		return err
	}

	fmt.Printf("Congratulations, you completed the quiz!\nFrom the %d questions you answered %d correctly\nThat places you in the top %d percent of our users!", *response.Data.QuestionCount, *response.Data.UserScore, *response.Data.UserPercentile)

	return nil
}
