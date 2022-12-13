package main

import (
	"context"
	"fmt"
	"github.com/zhenghaoz/gorse/client"
)

func main() {
	// Create a client
	gorse := client.NewGorseClient("http://127.0.0.1:8088", "api_key")
	ctx := context.Background()
	// Insert feedback
	gorse.InsertFeedback(ctx, []client.Feedback{
		{FeedbackType: "read", UserId: "0-vortex", ItemId: "vuejs:vue", Timestamp: "2022-12-02"},
		{FeedbackType: "star", UserId: "0-vortex", ItemId: "vuejs:vue", Timestamp: "2022-12-02"},
		{FeedbackType: "read", UserId: "0-vortex", ItemId: "d3:d3", Timestamp: "2022-12-04"},
		{FeedbackType: "like", UserId: "0-vortex", ItemId: "d3:d3", Timestamp: "2022-12-04"},
		{FeedbackType: "read", UserId: "0-vortex", ItemId: "dogfalo:materialize", Timestamp: "2022-12-01"},
		{FeedbackType: "star", UserId: "0-vortex", ItemId: "dogfalo:materialize", Timestamp: "2022-12-01"},
		{FeedbackType: "read", UserId: "0-vortex", ItemId: "mozilla:pdf.js", Timestamp: "2022-12-01"},
		{FeedbackType: "like", UserId: "0-vortex", ItemId: "mozilla:pdf.js", Timestamp: "2022-12-01"},
		{FeedbackType: "read", UserId: "0-vortex", ItemId: "moment:moment", Timestamp: "2022-12-01"},
		{FeedbackType: "star", UserId: "0-vortex", ItemId: "moment:moment", Timestamp: "2022-12-01"},
	})

	// Get recommendation.
	//recommend, err := gorse.GetRecommend(ctx, "bob33", "", 10)
	//if err != nil {
	//	return
	//}
	//for i, v := range recommend {
	//	fmt.Println(i, ":", v)
	//}
	//user, err := gorse.GetUser(ctx, "bob")
	//if err != nil {
	//	return
	//}
	//for i, v := range user.Labels {
	//	fmt.Println(i, ":", v)
	//}
	userNeighbors, err := gorse.GetUserNeighbors(ctx, "0-vortex", 10)
	if err != nil {
		return
	}
	for _, v := range userNeighbors {
		fmt.Println(v.Id)
	}
}
