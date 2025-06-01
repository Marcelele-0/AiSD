package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// OperationStats przechowuje statystyki pojedynczej operacji
type OperationStats struct {
	Comparisons    int `json:"comparisons"`
	PointerUpdates int `json:"pointer_updates"`
	Height         int `json:"height"`
}

// TestResult przechowuje wyniki testu dla konkretnego n
type TestResult struct {
	N                    int              `json:"n"`
	TestNumber           int              `json:"test_number"`
	Scenario             string           `json:"scenario"`
	InsertOperations     []OperationStats `json:"insert_operations"`
	DeleteOperations     []OperationStats `json:"delete_operations"`
	InsertAvgComparisons float64          `json:"insert_avg_comparisons"`
	InsertMaxComparisons int              `json:"insert_max_comparisons"`
	InsertAvgPointers    float64          `json:"insert_avg_pointers"`
	InsertMaxPointers    int              `json:"insert_max_pointers"`
	InsertAvgHeight      float64          `json:"insert_avg_height"`
	InsertMaxHeight      int              `json:"insert_max_height"`
	DeleteAvgComparisons float64          `json:"delete_avg_comparisons"`
	DeleteMaxComparisons int              `json:"delete_max_comparisons"`
	DeleteAvgPointers    float64          `json:"delete_avg_pointers"`
	DeleteMaxPointers    int              `json:"delete_max_pointers"`
	DeleteAvgHeight      float64          `json:"delete_avg_height"`
	DeleteMaxHeight      int              `json:"delete_max_height"`
	TotalTime            float64          `json:"total_time_ms"`
}

// AllResults przechowuje wszystkie wyniki testów
type AllResults struct {
	OrderedResults []TestResult `json:"ordered_results"`
	RandomResults  []TestResult `json:"random_results"`
	Summary        Summary      `json:"summary"`
}

// Summary przechowuje podsumowanie wyników
type Summary struct {
	OrderedScenario ScenarioSummary `json:"ordered_scenario"`
	RandomScenario  ScenarioSummary `json:"random_scenario"`
}

// ScenarioSummary przechowuje średnie dla scenariusza
type ScenarioSummary struct {
	AvgResults []NSummary `json:"avg_results"`
}

// NSummary przechowuje średnie dla konkretnego n
type NSummary struct {
	N                    int     `json:"n"`
	AvgInsertComparisons float64 `json:"avg_insert_comparisons"`
	AvgInsertPointers    float64 `json:"avg_insert_pointers"`
	AvgInsertHeight      float64 `json:"avg_insert_height"`
	AvgDeleteComparisons float64 `json:"avg_delete_comparisons"`
	AvgDeletePointers    float64 `json:"avg_delete_pointers"`
	AvgDeleteHeight      float64 `json:"avg_delete_height"`
	AvgTotalTime         float64 `json:"avg_total_time_ms"`
}

// AveragedResult przechowuje uśrednione wyniki dla konkretnego n i scenariusza
type AveragedResult struct {
	N                    int     `json:"n"`
	Scenario             string  `json:"scenario"`
	AvgInsertComparisons float64 `json:"avg_insert_comparisons"`
	AvgInsertPointers    float64 `json:"avg_insert_pointers"`
	AvgInsertHeight      float64 `json:"avg_insert_height"`
	AvgDeleteComparisons float64 `json:"avg_delete_comparisons"`
	AvgDeletePointers    float64 `json:"avg_delete_pointers"`
	AvgDeleteHeight      float64 `json:"avg_delete_height"`
	AvgTotalTime         float64 `json:"avg_total_time_ms"`
}

// AveragedResults przechowuje skrócone wyniki
type AveragedResults struct {
	AveragedResults []AveragedResult `json:"averaged_results"`
}

// TestTask reprezentuje zadanie testowe do wykonania w worker pool
type TestTask struct {
	N       int
	TestNum int
	Ordered bool
	Seed    int64
}

// ProgressInfo zawiera informacje o postępie
type ProgressInfo struct {
	Completed int
	Total     int
	Current   TestTask
}

// runSingleTestThreadSafe wykonuje pojedynczy test w sposób bezpieczny dla współbieżności
func runSingleTestThreadSafe(task TestTask) TestResult {
	// Każdy goroutine ma własny generator liczb losowych
	localRand := rand.New(rand.NewSource(task.Seed))

	tree := NewSplayTree()
	result := TestResult{
		N:          task.N,
		TestNumber: task.TestNum,
	}

	if task.Ordered {
		result.Scenario = "ordered"
	} else {
		result.Scenario = "random"
	}

	startTime := time.Now()

	// Przygotuj dane
	var keys []int
	if task.Ordered {
		keys = make([]int, task.N)
		for i := 0; i < task.N; i++ {
			keys[i] = i + 1
		}
	} else {
		keys = make([]int, task.N)
		for i := 0; i < task.N; i++ {
			keys[i] = i + 1
		}
		// Shuffle używając lokalnego generatora
		for i := len(keys) - 1; i > 0; i-- {
			j := localRand.Intn(i + 1)
			keys[i], keys[j] = keys[j], keys[i]
		}
	}
	// Testuj operacje INSERT
	result.InsertOperations = make([]OperationStats, task.N)
	insertComparisonsSum := 0
	insertPointersSum := 0
	insertHeightSum := 0
	insertMaxComparisons := 0
	insertMaxPointers := 0
	insertMaxHeight := 0

	for i, key := range keys {
		tree.resetStats()
		tree.insert(key)
		height := tree.calculateHeight()

		opStats := OperationStats{
			Comparisons:    tree.comparisons,
			PointerUpdates: tree.pointerUpdates,
			Height:         height,
		}

		result.InsertOperations[i] = opStats
		insertComparisonsSum += tree.comparisons
		insertPointersSum += tree.pointerUpdates
		insertHeightSum += height

		if tree.comparisons > insertMaxComparisons {
			insertMaxComparisons = tree.comparisons
		}
		if tree.pointerUpdates > insertMaxPointers {
			insertMaxPointers = tree.pointerUpdates
		}
		if height > insertMaxHeight {
			insertMaxHeight = height
		}
	}

	// Oblicz średnie dla INSERT
	result.InsertAvgComparisons = float64(insertComparisonsSum) / float64(task.N)
	result.InsertMaxComparisons = insertMaxComparisons
	result.InsertAvgPointers = float64(insertPointersSum) / float64(task.N)
	result.InsertMaxPointers = insertMaxPointers
	result.InsertAvgHeight = float64(insertHeightSum) / float64(task.N)
	result.InsertMaxHeight = insertMaxHeight

	// Przygotuj klucze do usuwania (w losowej kolejności)
	deleteKeys := make([]int, len(keys))
	copy(deleteKeys, keys)
	for i := len(deleteKeys) - 1; i > 0; i-- {
		j := localRand.Intn(i + 1)
		deleteKeys[i], deleteKeys[j] = deleteKeys[j], deleteKeys[i]
	}
	// Testuj operacje DELETE
	result.DeleteOperations = make([]OperationStats, task.N)
	deleteComparisonsSum := 0
	deletePointersSum := 0
	deleteHeightSum := 0
	deleteMaxComparisons := 0
	deleteMaxPointers := 0
	deleteMaxHeight := 0

	for i, key := range deleteKeys {
		tree.resetStats()
		tree.deleteNode(key)
		height := tree.calculateHeight()

		opStats := OperationStats{
			Comparisons:    tree.comparisons,
			PointerUpdates: tree.pointerUpdates,
			Height:         height,
		}

		result.DeleteOperations[i] = opStats
		deleteComparisonsSum += tree.comparisons
		deletePointersSum += tree.pointerUpdates
		deleteHeightSum += height

		if tree.comparisons > deleteMaxComparisons {
			deleteMaxComparisons = tree.comparisons
		}
		if tree.pointerUpdates > deleteMaxPointers {
			deleteMaxPointers = tree.pointerUpdates
		}
		if height > deleteMaxHeight {
			deleteMaxHeight = height
		}
	}

	// Oblicz średnie dla DELETE
	result.DeleteAvgComparisons = float64(deleteComparisonsSum) / float64(task.N)
	result.DeleteMaxComparisons = deleteMaxComparisons
	result.DeleteAvgPointers = float64(deletePointersSum) / float64(task.N)
	result.DeleteMaxPointers = deleteMaxPointers
	result.DeleteAvgHeight = float64(deleteHeightSum) / float64(task.N)
	result.DeleteMaxHeight = deleteMaxHeight

	result.TotalTime = float64(time.Since(startTime).Nanoseconds()) / 1e6 // milisekundy

	return result
}

// worker funkcja dla worker pool
func worker(tasks <-chan TestTask, results chan<- TestResult, progress chan<- ProgressInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		result := runSingleTestThreadSafe(task)
		results <- result

		// Wyślij informacje o postępie
		select {
		case progress <- ProgressInfo{Current: task}:
		default:
		}
	}
}

// runBenchmark wykonuje pełny benchmark
func runBenchmark() AllResults {
	fmt.Println("🌳 Splay Tree Performance Benchmark (Wielowątkowy)")
	fmt.Println("==============================================")

	// Konfiguracja
	nValues := []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000}
	testsPerN := 20
	numWorkers := 4

	// if runtime.NumCPU() < numWorkers {
	// 	numWorkers = runtime.NumCPU()
	// }

	fmt.Printf("🚀 Rozpoczynam wielowątkowe testy wydajności Splay Tree...\n")
	fmt.Printf("⚙️ Liczba wątków roboczych: %d (z %d dostępnych rdzeni)\n", numWorkers, runtime.NumCPU())
	fmt.Printf("📏 Wartości n: %v\n", nValues)
	fmt.Printf("🔄 Liczba testów dla każdego n: %d\n\n", testsPerN)

	totalTests := len(nValues) * testsPerN * 2 // 2 scenariusze
	fmt.Printf("📈 Łączna liczba testów do wykonania: %d\n\n", totalTests)

	// Przygotuj kanały
	tasks := make(chan TestTask, totalTests)
	results := make(chan TestResult, totalTests)
	progress := make(chan ProgressInfo, 100)

	// Uruchom worker pool
	var wg sync.WaitGroup
	fmt.Println("📤 Wysyłanie zadań do workerów...")

	// Uruchom workery
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, results, progress, &wg)
	}

	// Generuj zadania
	go func() {
		defer close(tasks)

		for _, n := range nValues {
			for testNum := 1; testNum <= testsPerN; testNum++ {
				// Test uporządkowany
				tasks <- TestTask{
					N:       n,
					TestNum: testNum,
					Ordered: true,
					Seed:    time.Now().UnixNano() + int64(n*testNum),
				}

				// Test losowy
				tasks <- TestTask{
					N:       n,
					TestNum: testNum,
					Ordered: false,
					Seed:    time.Now().UnixNano() + int64(n*testNum) + 1000000,
				}
			}
		}
	}()

	// Zbieraj wyniki z pokazaniem postępu
	go func() {
		defer close(results)
		wg.Wait()
	}()

	fmt.Println("📥 Zbieranie wyników...")

	var orderedResults []TestResult
	var randomResults []TestResult
	completed := 0

	// Pokazuj postęp podczas zbierania wyników
	progressTicker := time.NewTicker(2 * time.Second)
	defer progressTicker.Stop()

	for {
		select {
		case result, ok := <-results:
			if !ok {
				// Kanał zamknięty, wszystkie wyniki zebrane
				goto ProcessResults
			}

			if result.Scenario == "ordered" {
				orderedResults = append(orderedResults, result)
			} else {
				randomResults = append(randomResults, result)
			}
			completed++

		case <-progressTicker.C:
			fmt.Printf("⏳ Postęp: %d/%d testów zakończonych (%.1f%%)\n",
				completed, totalTests, float64(completed)/float64(totalTests)*100)
		}
	}

ProcessResults:
	fmt.Printf("✅ Wszystkie testy zakończone! (%d wyników)\n\n", completed)

	// Oblicz podsumowania
	summary := calculateSummary(orderedResults, randomResults, nValues)

	return AllResults{
		OrderedResults: orderedResults,
		RandomResults:  randomResults,
		Summary:        summary,
	}
}

// calculateSummary oblicza podsumowanie wyników
func calculateSummary(orderedResults, randomResults []TestResult, nValues []int) Summary {
	orderedSummary := calculateScenarioSummary(orderedResults, nValues)
	randomSummary := calculateScenarioSummary(randomResults, nValues)

	return Summary{
		OrderedScenario: orderedSummary,
		RandomScenario:  randomSummary,
	}
}

// calculateScenarioSummary oblicza podsumowanie dla scenariusza
func calculateScenarioSummary(results []TestResult, nValues []int) ScenarioSummary {
	var avgResults []NSummary

	for _, n := range nValues {
		var nResults []TestResult
		for _, result := range results {
			if result.N == n {
				nResults = append(nResults, result)
			}
		}

		if len(nResults) == 0 {
			continue
		}
		// Oblicz średnie
		var sumInsertComparisons, sumInsertPointers, sumInsertHeight float64
		var sumDeleteComparisons, sumDeletePointers, sumDeleteHeight float64
		var sumTotalTime float64

		for _, result := range nResults {
			sumInsertComparisons += result.InsertAvgComparisons
			sumInsertPointers += result.InsertAvgPointers
			sumInsertHeight += result.InsertAvgHeight
			sumDeleteComparisons += result.DeleteAvgComparisons
			sumDeletePointers += result.DeleteAvgPointers
			sumDeleteHeight += result.DeleteAvgHeight
			sumTotalTime += result.TotalTime
		}

		count := float64(len(nResults))
		avgResults = append(avgResults, NSummary{
			N:                    n,
			AvgInsertComparisons: sumInsertComparisons / count,
			AvgInsertPointers:    sumInsertPointers / count,
			AvgInsertHeight:      sumInsertHeight / count,
			AvgDeleteComparisons: sumDeleteComparisons / count,
			AvgDeletePointers:    sumDeletePointers / count,
			AvgDeleteHeight:      sumDeleteHeight / count,
			AvgTotalTime:         sumTotalTime / count,
		})
	}

	return ScenarioSummary{AvgResults: avgResults}
}

// generateAveragedResults generuje skrócone wyniki
func generateAveragedResults(allResults AllResults) AveragedResults {
	var averaged []AveragedResult
	// Dodaj wyniki dla scenariusza uporządkowanego
	for _, nSummary := range allResults.Summary.OrderedScenario.AvgResults {
		averaged = append(averaged, AveragedResult{
			N:                    nSummary.N,
			Scenario:             "ordered",
			AvgInsertComparisons: nSummary.AvgInsertComparisons,
			AvgInsertPointers:    nSummary.AvgInsertPointers,
			AvgInsertHeight:      nSummary.AvgInsertHeight,
			AvgDeleteComparisons: nSummary.AvgDeleteComparisons,
			AvgDeletePointers:    nSummary.AvgDeletePointers,
			AvgDeleteHeight:      nSummary.AvgDeleteHeight,
			AvgTotalTime:         nSummary.AvgTotalTime,
		})
	}

	// Dodaj wyniki dla scenariusza losowego
	for _, nSummary := range allResults.Summary.RandomScenario.AvgResults {
		averaged = append(averaged, AveragedResult{
			N:                    nSummary.N,
			Scenario:             "random",
			AvgInsertComparisons: nSummary.AvgInsertComparisons,
			AvgInsertPointers:    nSummary.AvgInsertPointers,
			AvgInsertHeight:      nSummary.AvgInsertHeight,
			AvgDeleteComparisons: nSummary.AvgDeleteComparisons,
			AvgDeletePointers:    nSummary.AvgDeletePointers,
			AvgDeleteHeight:      nSummary.AvgDeleteHeight,
			AvgTotalTime:         nSummary.AvgTotalTime,
		})
	}

	return AveragedResults{AveragedResults: averaged}
}

// runBenchmarkMain główna funkcja benchmark
func runBenchmarkMain() {
	startTime := time.Now()

	// Wykonaj benchmark
	allResults := runBenchmark()

	// Zapisz pełne wyniki
	fullData, err := json.MarshalIndent(allResults, "", "  ")
	if err != nil {
		fmt.Printf("❌ Błąd serializacji JSON: %v\n", err)
		return
	}

	fullFilename := "splay_benchmark_full.json"
	err = os.WriteFile(fullFilename, fullData, 0644)
	if err != nil {
		fmt.Printf("❌ Błąd zapisu pliku pełnego: %v\n", err)
		return
	}

	// Generuj i zapisz skrócone wyniki
	averagedResults := generateAveragedResults(allResults)
	shortData, err := json.MarshalIndent(averagedResults, "", "  ")
	if err != nil {
		fmt.Printf("❌ Błąd serializacji skróconych wyników: %v\n", err)
		return
	}

	shortFilename := "splay_benchmark_short.json"
	err = os.WriteFile(shortFilename, shortData, 0644)
	if err != nil {
		fmt.Printf("❌ Błąd zapisu skróconego pliku: %v\n", err)
		return
	}

	totalTime := time.Since(startTime)

	// Wyświetl podsumowanie
	fmt.Println("📊 PODSUMOWANIE WYNIKÓW")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("📁 Pełne wyniki zapisane w: %s\n", fullFilename)
	fmt.Printf("📁 Skrócone wyniki zapisane w: %s\n", shortFilename)
	fmt.Printf("⏱️ Łączny czas benchmark: %v\n", totalTime)
	fmt.Printf("🔢 Łączna liczba testów: %d\n", len(allResults.OrderedResults)+len(allResults.RandomResults))

	// Pokazuj przykładowe wyniki
	if len(allResults.Summary.OrderedScenario.AvgResults) > 0 {
		fmt.Println("\n📈 Przykładowe wyniki (scenariusz uporządkowany):")
		for i, result := range allResults.Summary.OrderedScenario.AvgResults {
			if i >= 3 { // Pokaż tylko pierwsze 3
				break
			}
			fmt.Printf("  n=%d: Avg Insert Comparisons=%.1f, Avg Height=%.1f, Avg Pointers=%.1f\n",
				result.N, result.AvgInsertComparisons, result.AvgInsertHeight, result.AvgInsertPointers)
		}
	}

	fmt.Println("\n✅ Benchmark zakończony pomyślnie!")
}
