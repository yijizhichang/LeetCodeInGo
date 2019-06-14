package problem

import (
	"time"
	"sort"
)

/*
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
示例:
Twitter twitter = new Twitter();

// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);

// 用户1关注了用户2.
twitter.follow(1, 2);

// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);

// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);

// 用户1取消关注了用户2.
twitter.unfollow(1, 2);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);
 */

type Twit struct {
	TwitId int
	Time   time.Time
}

type Twits []Twit

func (t Twits) Len() int {
	return len(t)
}

func (t Twits) Less(i, j int) bool {
	if t[i].Time.After(t[j].Time) {
		return true
	}
	return false
}

func (t Twits) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type Twitter struct {
	UserTwit   map[int]Twits
	UserFollow map[int][]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		UserTwit:   make(map[int]Twits),
		UserFollow: make(map[int][]int),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.UserTwit[userId] = append(this.UserTwit[userId], Twit{
		tweetId,
		time.Now(),
	})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	ret := make([]int, 0)
	temp := make(Twits, this.UserTwit[userId].Len())
	copy(temp, this.UserTwit[userId])
	for _, v := range this.UserFollow[userId] {
		temp = append(temp, this.UserTwit[v]...)
	}
	sort.Sort(temp)
	for i := 0; i < 10 && i < temp.Len(); i++ {
		ret = append(ret, temp[i].TwitId)
	}
	return ret
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	for _, v := range this.UserFollow[followerId] {
		if v == followeeId {
			return
		}
	}
	this.UserFollow[followerId] = append(this.UserFollow[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	for i, v := range this.UserFollow[followerId] {
		if v == followeeId {
			this.UserFollow[followerId] = append(this.UserFollow[followerId][:i], this.UserFollow[followerId][i+1:]...)
		}
	}
}
