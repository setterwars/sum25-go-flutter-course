package message

import (
	"sync"
)

// Message represents a chat message
// Add more fields if needed

type Message struct {
	Sender    string
	Content   string
	Timestamp int64
}

// MessageStore stores chat messages
// Contains a slice of messages and a mutex for concurrency

type MessageStore struct {
	messages []Message
	mutex    sync.RWMutex
	// Add more fields if needed
}

// NewMessageStore creates a new MessageStore
func NewMessageStore() *MessageStore {
	// Initialize MessageStore fields
	return &MessageStore{
		messages: make([]Message, 0, 100),
	}
}

// AddMessage stores a new message
func (s *MessageStore) AddMessage(msg Message) error {
	// Add message to storage (concurrent safe)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.messages = append(s.messages, msg)
	return nil
}

// GetMessages retrieves messages (optionally by user)
func (s *MessageStore) GetMessages(user string) ([]Message, error) {
	// Retrieve messages (all or by user)
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if user == "" {
		// Return all messages
		result := make([]Message, len(s.messages))
		copy(result, s.messages)
		return result, nil
	}
	// Filter messages by sender
	var filtered []Message
	for _, msg := range s.messages {
		if msg.Sender == user {
			filtered = append(filtered, msg)
		}
	}
	return filtered, nil
}
