package main

import (
	"fmt"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	channels := make([]chan interface{}, len(cmds)+1)
	for i := range channels {
		channels[i] = make(chan interface{})
	}

	wg := &sync.WaitGroup{}
	for i, command := range cmds {
		wg.Add(1)
		go func(cmd cmd, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			cmd(in, out)
		}(command, channels[i], channels[i+1])
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	seenUsers := make(map[uint64]bool)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for email := range in {
		emailStr := email.(string)
		wg.Add(1)

		go func(e string) {
			defer wg.Done()
			user := GetUser(e)

			mu.Lock()
			if !seenUsers[user.ID] {
				seenUsers[user.ID] = true
				mu.Unlock()
				out <- user
			} else {
				mu.Unlock()
			}
		}(emailStr)
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	batch := make([]User, 0, GetMessagesMaxUsersBatch)

	for user := range in {
		u := user.(User)
		batch = append(batch, u)

		if len(batch) >= GetMessagesMaxUsersBatch {
			wg.Add(1)
			currentBatch := make([]User, len(batch))
			copy(currentBatch, batch)
			batch = batch[:0]

			go func(users []User) {
				defer wg.Done()
				msgIDs, err := GetMessages(users...)
				if err != nil {
					return
				}
				for _, msgID := range msgIDs {
					out <- msgID
				}
			}(currentBatch)
		}
	}

	if len(batch) > 0 {
		wg.Add(1)
		go func(users []User) {
			defer wg.Done()
			msgIDs, err := GetMessages(users...)
			if err != nil {
				return
			}
			for _, msgID := range msgIDs {
				out <- msgID
			}
		}(batch)
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	semaphore := make(chan struct{}, HasSpamMaxAsyncRequests)
	wg := &sync.WaitGroup{}

	for msgID := range in {
		id := msgID.(MsgID)
		wg.Add(1)

		go func(m MsgID) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			hasSpam, err := HasSpam(m)
			if err != nil {
				return
			}

			out <- MsgData{
				ID:      m,
				HasSpam: hasSpam,
			}
		}(id)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	results := make([]MsgData, 0)

	for data := range in {
		msgData := data.(MsgData)
		results = append(results, msgData)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].HasSpam != results[j].HasSpam {
			return results[i].HasSpam
		}
		return results[i].ID < results[j].ID
	})

	for _, msg := range results {
		if msg.HasSpam {
			out <- fmt.Sprintf("true %d", msg.ID)
		} else {
			out <- fmt.Sprintf("false %d", msg.ID)
		}
	}
}
