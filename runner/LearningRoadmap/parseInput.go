package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(inputFilePath string) (map[string]*Skill, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := make(map[string]*Skill)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")
		if len(parts) == 2 {
			childName := strings.TrimSpace(parts[1])
			parentName := strings.TrimSpace(parts[0])
			childSkill, exists := graph[childName]
			if !exists {
				childSkill = &Skill{Name: childName}
				graph[childName] = childSkill
			}
			parentSkill, exists := graph[parentName]
			if !exists {
				parentSkill = &Skill{Name: parentName}
				graph[parentName] = parentSkill
			}
			childSkill.Prerequisites = append(childSkill.Prerequisites, parentSkill)
		} else if len(parts) == 1 && strings.Contains(parts[0], "=") {
			progressParts := strings.Split(parts[0], "=")
			name := strings.TrimSpace(progressParts[0])
			progress := strings.TrimSpace(progressParts[1])
			progressValue := 0.0
			fmt.Sscanf(progress, "%f", &progressValue)
			skill, exists := graph[name]
			if exists {
				skill.Progress = progressValue
			} else {
				graph[name] = &Skill{Name: name, Progress: progressValue}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}
