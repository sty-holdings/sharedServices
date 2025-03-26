package sharedServices

import (
	"time"

	"gorm.io/gorm"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

// NOTE: Table go files do not follow the Constant Type Vars file format.

type AnalyzedQuestionResults struct {
	AnalysisID                          string    `gorm:"type:VARCHAR(50);primaryKey;default:gen_random_uuid()"`
	CreateTimestamp                     time.Time `gorm:"type:TIMESTAMP WITH TIME ZONE;default:now()"`
	ElapseTimeSeconds                   float64   `gorm:"type:DOUBLE PRECISION"`
	AverageFlag                         bool      `gorm:"type:BOOLEAN"`
	ComparisonFlag                      bool      `gorm:"type:BOOLEAN"`
	CompoundFlag                        bool      `gorm:"type:BOOLEAN"`
	CountFlag                           bool      `gorm:"type:BOOLEAN"`
	DetailFlag                          bool      `gorm:"type:BOOLEAN"`
	ForecastFlag                        bool      `gorm:"type:BOOLEAN"`
	MaximumFlag                         bool      `gorm:"type:BOOLEAN"`
	MinimumFlag                         bool      `gorm:"type:BOOLEAN"`
	PercentageFlag                      bool      `gorm:"type:BOOLEAN"`
	ReportFlag                          bool      `gorm:"type:BOOLEAN"`
	SubtotalFlag                        bool      `gorm:"type:BOOLEAN"`
	SummaryFlag                         bool      `gorm:"type:BOOLEAN"`
	ToDateFlag                          bool      `gorm:"type:BOOLEAN"`
	TotalFlag                           bool      `gorm:"type:BOOLEAN"`
	TransactionFlag                     bool      `gorm:"type:BOOLEAN"`
	TrendFlag                           bool      `gorm:"type:BOOLEAN"`
	AIPrompt                            string    `gorm:"type:VARCHAR(1024)"`
	Category                            string    `gorm:"type:VARCHAR(255)"`
	CountBySubject                      string    `gorm:"type:VARCHAR(255)"`
	RelativeTimeWord                    string    `gorm:"type:VARCHAR(255)"`
	SentenceSubject                     string    `gorm:"type:VARCHAR(255)"`
	SentenceSubjectAdverb               string    `gorm:"type:VARCHAR(255)"`
	UserPrompt                          string    `gorm:"type:VARCHAR(1024)"`
	CategorySentenceCandidateTokenCount int       `gorm:"type:INTEGER"`
	CategorySentencePromptTokenCount    int       `gorm:"type:INTEGER"`
	CategorySentenceTotalTokenCount     int       `gorm:"type:INTEGER"`
	SpecialWordsCandidateTokenCount     int       `gorm:"type:INTEGER"`
	SpecialWordsPromptTokenCount        int       `gorm:"type:INTEGER"`
	SpecialWordsTotalTokenCount         int       `gorm:"type:INTEGER"`
	TimePeriodValuesCandidateTokenCount int       `gorm:"type:INTEGER"`
	TimePeriodValuesPromptTokenCount    int       `gorm:"type:INTEGER"`
	TimePeriodValuesTotalTokenCount     int       `gorm:"type:INTEGER"`
	ValuesYear                          string    `gorm:"type:VARCHAR(255)"`
	ValuesQuarter                       string    `gorm:"type:VARCHAR(255)"`
	ValuesMonth                         string    `gorm:"type:VARCHAR(255)"`
	ValuesDay                           string    `gorm:"type:VARCHAR(255)"`
	ValuesWeek                          string    `gorm:"type:VARCHAR(255)"`
	SundayDate                          string    `gorm:"type:VARCHAR(255)"`
	BatchName                           string    `gorm:"type:VARCHAR(255)"`
}

func (AnalyzedQuestionResults) TableName() string {
	return "dka.analyzed_question_results"
}

// InsertAnalyzedQuestions - populates a batch (max: 100) of rows and inserts them into the answers database
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (psqlServicePtr *PSQLService) InsertAnalyzedQuestions(analyzedQuestionBatch []ctv.AnalyzedQuestion, batchName string, elapseTime float64, fileTag string) (
	errorInfo errs.ErrorInfo,
) {

	var (
		batchRows        []AnalyzedQuestionResults
		analyzedQuestion ctv.AnalyzedQuestion
		result           *gorm.DB
	)

	for _, analyzedQuestion = range analyzedQuestionBatch {
		batchRows = append(
			batchRows, AnalyzedQuestionResults{
				AnalysisID:        analyzedQuestion.AnalysisId,
				CreateTimestamp:   time.Now().UTC(),
				ElapseTimeSeconds: elapseTime,
				AverageFlag:       analyzedQuestion.SpecialWords.AverageFlag,
				ComparisonFlag:    analyzedQuestion.SpecialWords.ComparisonFlag,
				CompoundFlag:      analyzedQuestion.SpecialWords.CompoundFlag,
				CountFlag:         analyzedQuestion.SpecialWords.CountFlag,
				DetailFlag:        analyzedQuestion.SpecialWords.DetailFlag,
				ForecastFlag:      analyzedQuestion.SpecialWords.ForecastFlag,
				MaximumFlag:       analyzedQuestion.SpecialWords.MaximumFlag,
				MinimumFlag:       analyzedQuestion.SpecialWords.MinimumFlag,
				PercentageFlag:    analyzedQuestion.SpecialWords.PercentageFlag,
				ReportFlag:        analyzedQuestion.SpecialWords.ReportFlag,
				SubtotalFlag:      analyzedQuestion.SpecialWords.SubTotalFlag,
				SummaryFlag:       analyzedQuestion.SpecialWords.SummaryFlag,
				ToDateFlag:        analyzedQuestion.TimePeriodValues.ToDate,
				TotalFlag:         analyzedQuestion.SpecialWords.TotalFlag,
				TransactionFlag:   analyzedQuestion.SpecialWords.TransactionFlag,
				TrendFlag:         analyzedQuestion.SpecialWords.TrendFlag,
				AIPrompt:          analyzedQuestion.CategorySentence.Prompt,
				UserPrompt:        analyzedQuestion.UserQuestion,
				BatchName:         batchName,
			},
		)

		if result = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].CreateInBatches(batchRows, 100); result.Error != nil {
			errorInfo = errs.NewErrorInfo(result.Error, errs.BuildLabelValueMessage(ctv.LBL_EXTENSION_HAL, ctv.LBL_PSQL_BATCH, batchName, ctv.TXT_FAILED))
		}
	}

	return
}
