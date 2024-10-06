package tasks

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const TimeFormat = "20060102"

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}

func (task *Task) Validate(checkWithId bool) error {

	var errs []string

	if checkWithId {
		if task.Id == "" {
			errs = append(errs, "the Id is blank")
		} else {
			taskId, err := strconv.Atoi(task.Id)
			if err != nil {
				errs = append(errs, "wrong format of Id")
			} else {
				if taskId <= 0 {
					errs = append(errs, "wrong Id: Id must be > 0")
				}
			}
		}
	}

	if task.Title == "" {
		errs = append(errs, "the Title is blank")
	}

	currentTime := time.Now().Truncate(24 * time.Hour)
	if task.Date == "" {
		task.Date = currentTime.Format(TimeFormat)
	}

	if task.Repeat == "" {
		checkedDate, err := time.Parse(TimeFormat, task.Date)
		if err != nil {
			errs = append(errs, "wrong format of the Date")
		} else {
			if checkedDate.Before(currentTime) {
				task.Date = currentTime.Format(TimeFormat)
			}
		}
	} else {
		nextDate, err := NextDate(currentTime, task.Date, task.Repeat)
		if err != nil {
			errs = append(errs, err.Error())
		} else if taskDate, _ := time.Parse(TimeFormat, task.Date); taskDate.Before(currentTime) {
			task.Date = nextDate
		}
	}

	if len(errs) > 0 {
		errDescription := strings.Join(errs, "; ")
		return fmt.Errorf("failed to validate task.Task: %s", errDescription)
	}

	return nil
}

func NextDate(now time.Time, sourceDate, repeat string) (string, error) {

	date, err := time.Parse(TimeFormat, sourceDate)
	if err != nil {
		return "", fmt.Errorf("failed to parse date: %w", err)
	}

	if strings.Contains(repeat, "d") {
		return nextDateWithD(now, date, repeat)
	} else if strings.Contains(repeat, "y") {
		return nextDateWithY(now, date, repeat)
	} else {
		return "", fmt.Errorf("repeat rule is undefined. The rule: \"%s\"", repeat)
	}

}

func errWrongRepeatFormat(ruleType, repeatRule string) error {
	return fmt.Errorf("wrong repeat rule format with \"%s\": rule - \"%s\"", ruleType, repeatRule)
}

func nextDateWithD(now, date time.Time, repeat string) (string, error) {

	if !validRepeatD(repeat) {
		return "", errWrongRepeatFormat("d", repeat)
	}

	days, _ := strconv.Atoi(repeat[2:])
	if date.Before(now) {
		for date.Before(now) {
			date = date.Add(time.Hour * time.Duration(days*24))
		}
	} else {
		date = date.Add(time.Hour * time.Duration(days*24))
	}

	return date.Format(TimeFormat), nil
}

func validRepeatD(rule string) bool {

	rule = strings.TrimSpace(rule)
	if valid, _ := regexp.MatchString(`^d\s\d{1,3}$`, rule); !valid {
		return false
	}

	days, err := strconv.Atoi(rule[2:])
	if err != nil || days > 400 {
		return false
	}

	return true
}

func nextDateWithY(now, date time.Time, repeat string) (string, error) {

	if !validRepeatY(repeat) {
		return "", errWrongRepeatFormat("y", repeat)
	}

	if date.Before(now) {
		for date.Before(now) {
			date = date.AddDate(1, 0, 0)
		}
	} else {
		date = date.AddDate(1, 0, 0)
	}

	return date.Format(TimeFormat), nil
}

func validRepeatY(rule string) bool {

	rule = strings.TrimSpace(rule)
	return rule == "y"
}
