package processing

// write this as a struct
// add mutexes to lock write and read not an issue bc no concurrency
var userReviewTable map[string]map[string]interface{}

func WriteUserReview(userId string, reviewValues map[string]interface{}) {
	if userReviewTable == nil {
		userReviewTable = make(map[string]map[string]interface{})
	}

	userReviewTable[userId] = reviewValues
}

func UserReviewExists(userId string) bool {
	_, exists := userReviewTable[userId]
	return exists
}
