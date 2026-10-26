package domain

import (
	"errors"
	"fmt"
	"strings"
)

type Status string

const (
	StatusAssigned   Status = "assigned"
	StatusInProgress Status = "in_progress"
	StatusInReview   Status = "in_review"
	StatusCompleted  Status = "completed"
)

var (
	ErrInvalidTask       = errors.New("invalid task")
	ErrUnauthorized      = errors.New("actor is not allowed to perform this action")
	ErrInvalidTransition = errors.New("invalid task state transition")
	ErrFeedbackRequired  = errors.New("feedback is required when rejecting a task")
	ErrAlreadyRewarded   = errors.New("task reward has already been paid")
)

// Task is the small domain model used to test business rules before wiring them
// into the Cosmos SDK state machine. It does not persist state or transfer coins.
type Task struct {
	ID          uint64
	Leader      string
	Assignee    string
	Title       string
	Deadline    int64
	Reward      uint64
	Status      Status
	ReviewNote  string
	RewardPaid  bool
}

func NewTask(id uint64, leader, assignee, title string, deadline int64, reward uint64, now int64) (Task, error) {
	if id == 0 || strings.TrimSpace(leader) == "" || strings.TrimSpace(assignee) == "" ||
		strings.TrimSpace(title) == "" || deadline <= now || reward == 0 {
		return Task{}, ErrInvalidTask
	}
	if leader == assignee {
		return Task{}, fmt.Errorf("%w: leader and assignee must differ", ErrInvalidTask)
	}
	return Task{
		ID: id, Leader: leader, Assignee: assignee, Title: strings.TrimSpace(title),
		Deadline: deadline, Reward: reward, Status: StatusAssigned,
	}, nil
}

func (t *Task) Start(actor string) error {
	if actor != t.Assignee {
		return ErrUnauthorized
	}
	if t.Status != StatusAssigned {
		return ErrInvalidTransition
	}
	t.Status = StatusInProgress
	return nil
}

func (t *Task) Submit(actor string) error {
	if actor != t.Assignee {
		return ErrUnauthorized
	}
	if t.Status != StatusInProgress {
		return ErrInvalidTransition
	}
	t.Status = StatusInReview
	return nil
}

// Review returns the reward amount only on approval. A real chain implementation
// must transfer funded coins through the bank keeper in the same transaction;
// this model only checks the business rule and marks the logical reward as paid.
func (t *Task) Review(actor string, approve bool, feedback string) (uint64, error) {
	if actor != t.Leader {
		return 0, ErrUnauthorized
	}
	if t.Status != StatusInReview {
		return 0, ErrInvalidTransition
	}
	if t.RewardPaid {
		return 0, ErrAlreadyRewarded
	}
	if !approve && strings.TrimSpace(feedback) == "" {
		return 0, ErrFeedbackRequired
	}
	t.ReviewNote = strings.TrimSpace(feedback)
	if !approve {
		t.Status = StatusInProgress
		return 0, nil
	}
	t.Status = StatusCompleted
	t.RewardPaid = true
	return t.Reward, nil
}
