package domain

import (
	"errors"
	"testing"
)

func newTestTask(t *testing.T) Task {
	t.Helper()
	task, err := NewTask(1, "leader1", "member1", "Write tests", 200, 25, 100)
	if err != nil {
		t.Fatalf("NewTask() error = %v", err)
	}
	return task
}

func TestTaskHappyPathPaysRewardOnce(t *testing.T) {
	task := newTestTask(t)
	if err := task.Start("member1"); err != nil {
		t.Fatal(err)
	}
	if err := task.Submit("member1"); err != nil {
		t.Fatal(err)
	}
	reward, err := task.Review("leader1", true, "Looks good")
	if err != nil {
		t.Fatal(err)
	}
	if reward != 25 || task.Status != StatusCompleted || !task.RewardPaid {
		t.Fatalf("unexpected approved task state: task=%+v reward=%d", task, reward)
	}
	if _, err := task.Review("leader1", true, ""); !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("second review error = %v, want invalid transition", err)
	}
}

func TestOnlyAssigneeCanStart(t *testing.T) {
	task := newTestTask(t)
	if err := task.Start("attacker"); !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("Start() error = %v, want unauthorized", err)
	}
}

func TestRejectRequiresFeedbackAndReturnsToWork(t *testing.T) {
	task := newTestTask(t)
	if err := task.Start("member1"); err != nil {
		t.Fatal(err)
	}
	if err := task.Submit("member1"); err != nil {
		t.Fatal(err)
	}
	if _, err := task.Review("leader1", false, "  "); !errors.Is(err, ErrFeedbackRequired) {
		t.Fatalf("Review() error = %v, want feedback required", err)
	}
	if task.Status != StatusInReview {
		t.Fatalf("rejected invalid review should not change state, got %s", task.Status)
	}
	if reward, err := task.Review("leader1", false, "Add a test"); err != nil || reward != 0 {
		t.Fatalf("Review() = (%d, %v), want (0, nil)", reward, err)
	}
	if task.Status != StatusInProgress || task.RewardPaid || task.ReviewNote != "Add a test" {
		t.Fatalf("unexpected rejected task state: %+v", task)
	}
}

func TestLeaderCannotSubmitForMember(t *testing.T) {
	task := newTestTask(t)
	if err := task.Submit("leader1"); !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("Submit() error = %v, want unauthorized", err)
	}
}

func TestInvalidTaskIsRejected(t *testing.T) {
	if _, err := NewTask(0, "leader1", "member1", "Task", 200, 25, 100); !errors.Is(err, ErrInvalidTask) {
		t.Fatalf("NewTask() error = %v, want invalid task", err)
	}
	if _, err := NewTask(2, "leader1", "member1", "Task", 100, 25, 100); !errors.Is(err, ErrInvalidTask) {
		t.Fatalf("expired deadline error = %v, want invalid task", err)
	}
}
