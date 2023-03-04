package main

import "fmt"

type Student struct {
	name  string
	score []int
}

func (s *Student) getAverageScore() float64 {
	sum := 0
	for _, score := range s.score {
		sum += score
	}
	return float64(sum) / float64(len(s.score))
}

func (s *Student) getMinScore() int {
	min := s.score[0]
	for _, score := range s.score {
		if score < min {
			min = score
		}
	}
	return min
}

func (s *Student) getMaxScore() int {
	max := s.score[0]
	for _, score := range s.score {
		if score > max {
			max = score
		}
	}
	return max
}

func main() {
	var students [5]Student
	for i := 0; i < 5; i++ {
		var name string
		var score []int
		fmt.Printf("Masukkan nama siswa ke-%d: ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Masukkan skor siswa ke-%d (pisahkan dengan spasi): ", i+1)
		for j := 0; j < 3; j++ {
			var s int
			fmt.Scan(&s)
			score = append(score, s)
		}
		students[i] = Student{name: name, score: score}
	}
	var sumScore float64
	for _, student := range students {
		sumScore += student.getAverageScore()
	}
	fmt.Printf("Skor rata-rata dari %d siswa: %f\n", len(students), sumScore/float64(len(students)))
	var minScore, maxScore int
	var minStudent, maxStudent string
	for _, student := range students {
		if student.getMinScore() < minScore || minScore == 0 {
			minScore = student.getMinScore()
			minStudent = student.name
		}
		if student.getMaxScore() > maxScore {
			maxScore = student.getMaxScore()
			maxStudent = student.name
		}
	}
	fmt.Printf("Siswa dengan skor minimum: %s (%d)\n", minStudent, minScore)
	fmt.Printf("Siswa dengan skor maksimum: %s (%d)\n", maxStudent, maxScore)
}
