package main

import "gitlab.com/jacob-ernst/mets/pkg/models"

var activityRecords = []models.Activity{
	{Name: "power mower", Effort: 4.5},
	{Name: "running, 4 mph", Description: "Running 15 min/mile", Effort: 6},
	{Name: "sitting tasks, light effort", Description: "Examples are office work, chemistry lab work, computer work, light assembly repair, watch repair, reading, desk work", Effort: 1.5},
}

var roleRecords = []models.Role{
	{Name: "user", Description: "Default user role"},
	{Name: "admin", Description: "For privileged admin access"},
}
