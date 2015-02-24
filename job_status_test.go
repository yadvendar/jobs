// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

package jobs

import (
	"reflect"
	"testing"
)

func TestJobStatusCount(t *testing.T) {
	testingSetUp()
	defer testingTeardown()
	jobs, err := createAndSaveTestJobs(5)
	if err != nil {
		t.Errorf("Unexpected error: %s")
	}
	for _, status := range possibleStatuses {
		if status == StatusDestroyed {
			// Skip this one, since destroying a job means erasing all records from the database
			continue
		}
		for _, job := range jobs {
			job.setStatus(status)
		}
		count, err := status.Count()
		if err != nil {
			t.Errorf("Unexpected error in status.Count(): %s", err.Error())
		}
		if count != len(jobs) {
			t.Errorf("Expected %s.Count() to return %d after setting job statuses to %s, but got %d", status, len(jobs), status, count)
		}
	}
}

func TestJobStatusJobIds(t *testing.T) {
	testingSetUp()
	defer testingTeardown()
	jobs, err := createAndSaveTestJobs(5)
	if err != nil {
		t.Errorf("Unexpected error: %s")
	}
	jobIds := make([]string, len(jobs))
	for i, job := range jobs {
		jobIds[i] = job.id
	}
	for _, status := range possibleStatuses {
		if status == StatusDestroyed {
			// Skip this one, since destroying a job means erasing all records from the database
			continue
		}
		for _, job := range jobs {
			job.setStatus(status)
		}
		gotIds, err := status.JobIds()
		if err != nil {
			t.Errorf("Unexpected error in status.JobIds(): %s", err.Error())
		}
		if len(gotIds) != len(jobIds) {
			t.Errorf("%s.JobIds() was incorrect. Expected slice of length %d but got %d", len(jobIds), len(gotIds))
		}
		if !reflect.DeepEqual(jobIds, gotIds) {
			t.Errorf("%s.JobIds() was incorrect. Expected %v but got %v", status, jobIds, gotIds)
		}
	}
}
