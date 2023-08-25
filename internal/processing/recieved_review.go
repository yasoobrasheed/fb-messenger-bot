package processing

import (
	"fmt"

	"github.com/cdipaolo/sentiment"
)

func HandleReceivedReview(reviewText string) string {
	goodExperienceResponseText := `Thank you so much for your review of your purchase! We hope you love it. 
								   Here's your 20% discount code: ` + generateCode()
	badExperienceResponseText := `We're so sorry that you had a bad experience and want to make it up to you. 
								  Please give us a call at (847) 847-1847, and we'll get it fixed for you. 
								  Here's your 20% discount code: ` + generateCode()

	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	analysis := model.SentimentAnalysis(reviewText, sentiment.English)

	if analysis.Score > 0 {
		return goodExperienceResponseText
	} else {
		return badExperienceResponseText
	}
}

func generateCode() string {
	// make this randomized & hash it to avoid collisions
	return "G0F45TB04T5M0J1T05"
}
