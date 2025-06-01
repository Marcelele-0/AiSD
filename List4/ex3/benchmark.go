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
	Completed   int
	Total       int
	CurrentN    int
	Scenario    string
	WorkerCount int
}

// createPermutationThreadSafe tworzy losową permutację z własnym generatorem (thread-safe)
func createPermutationThreadSafe(n int, seed int64) []int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i + 1
	}

	// Używamy lokalnego generatora z konkretnym seed dla powtarzalności
	localRand := rand.New(rand.NewSource(seed))
	localRand.Shuffle(len(perm), func(i, j int) {
		perm[i], perm[j] = perm[j], perm[i]
	})

	return perm
}

// runSingleTestThreadSafe wykonuje pojedynczy test w sposób thread-safe
func runSingleTestThreadSafe(task TestTask) TestResult {
	startTime := time.Now()

	scenario := "random"
	if task.Ordered {
		scenario = "ordered"
	}

	result := TestResult{
		N:                task.N,
		TestNumber:       task.TestNum,
		Scenario:         scenario,
		InsertOperations: make([]OperationStats, 0, task.N),
		DeleteOperations: make([]OperationStats, 0, task.N),
	}

	rb := NewRB_BST()

	// Przygotowanie sekwencji do wstawiania
	var insertSequence []int
	if task.Ordered {
		insertSequence = make([]int, task.N)
		for i := 0; i < task.N; i++ {
			insertSequence[i] = i + 1
		}
	} else {
		insertSequence = createPermutationThreadSafe(task.N, task.Seed)
	}

	// Wstawianie elementów
	for _, key := range insertSequence {
		prevComparisons := rb.comparisons
		prevPointers := rb.pointerUpdates

		rb.insert(key)
		height := rb.calculateHeight()

		opStats := OperationStats{
			Comparisons:    rb.comparisons - prevComparisons,
			PointerUpdates: rb.pointerUpdates - prevPointers,
			Height:         height,
		}

		result.InsertOperations = append(result.InsertOperations, opStats)
	}

	// Przygotowanie sekwencji do usuwania (zawsze losowa)
	deleteSequence := createPermutationThreadSafe(task.N, task.Seed+1000)

	// Reset statystyk dla operacji usuwania
	rb.resetStats()

	// Usuwanie elementów
	for _, key := range deleteSequence {
		prevComparisons := rb.comparisons
		prevPointers := rb.pointerUpdates

		rb.deleteNode(key)
		height := rb.calculateHeight()

		opStats := OperationStats{
			Comparisons:    rb.comparisons - prevComparisons,
			PointerUpdates: rb.pointerUpdates - prevPointers,
			Height:         height,
		}

		result.DeleteOperations = append(result.DeleteOperations, opStats)
	}

	// Obliczanie statystyk
	result.calculateStats()
	result.TotalTime = float64(time.Since(startTime).Nanoseconds()) / 1e6 // w milisekundach

	return result
}

// calculateStats oblicza średnie i maksymalne wartości
func (tr *TestResult) calculateStats() {
	// Statystyki wstawiania
	if len(tr.InsertOperations) > 0 {
		totalComp, totalPtr, totalHeight := 0, 0, 0
		maxComp, maxPtr, maxHeight := 0, 0, 0

		for _, op := range tr.InsertOperations {
			totalComp += op.Comparisons
			totalPtr += op.PointerUpdates
			totalHeight += op.Height

			if op.Comparisons > maxComp {
				maxComp = op.Comparisons
			}
			if op.PointerUpdates > maxPtr {
				maxPtr = op.PointerUpdates
			}
			if op.Height > maxHeight {
				maxHeight = op.Height
			}
		}

		count := float64(len(tr.InsertOperations))
		tr.InsertAvgComparisons = float64(totalComp) / count
		tr.InsertAvgPointers = float64(totalPtr) / count
		tr.InsertAvgHeight = float64(totalHeight) / count
		tr.InsertMaxComparisons = maxComp
		tr.InsertMaxPointers = maxPtr
		tr.InsertMaxHeight = maxHeight
	}

	// Statystyki usuwania
	if len(tr.DeleteOperations) > 0 {
		totalComp, totalPtr, totalHeight := 0, 0, 0
		maxComp, maxPtr, maxHeight := 0, 0, 0

		for _, op := range tr.DeleteOperations {
			totalComp += op.Comparisons
			totalPtr += op.PointerUpdates
			totalHeight += op.Height

			if op.Comparisons > maxComp {
				maxComp = op.Comparisons
			}
			if op.PointerUpdates > maxPtr {
				maxPtr = op.PointerUpdates
			}
			if op.Height > maxHeight {
				maxHeight = op.Height
			}
		}

		count := float64(len(tr.DeleteOperations))
		tr.DeleteAvgComparisons = float64(totalComp) / count
		tr.DeleteAvgPointers = float64(totalPtr) / count
		tr.DeleteAvgHeight = float64(totalHeight) / count
		tr.DeleteMaxComparisons = maxComp
		tr.DeleteMaxPointers = maxPtr
		tr.DeleteMaxHeight = maxHeight
	}
}

// runBenchmark wykonuje kompletny benchmark wielowątkowy
func runBenchmark() AllResults {
	results := AllResults{
		OrderedResults: make([]TestResult, 0),
		RandomResults:  make([]TestResult, 0),
	}

	// Konfiguracja wielowątkowości - ustaw tutaj liczbę wątków
	numWorkers := 4 // Można zmienić tę wartość (np. 4, 8, 16)
	if numWorkers > runtime.NumCPU() {
		numWorkers = runtime.NumCPU()
	}

	// Wartości n do testowania
	nValues := []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000}
	testsPerN := 20

	fmt.Println("🚀 Rozpoczynam wielowątkowe testy wydajności RB-BST...")
	fmt.Printf("⚙️ Liczba wątków roboczych: %d (z %d dostępnych rdzeni)\n", numWorkers, runtime.NumCPU())
	fmt.Printf("📏 Wartości n: %v\n", nValues)
	fmt.Printf("🔄 Liczba testów dla każdego n: %d\n", testsPerN)
	fmt.Println()

	totalTests := len(nValues) * testsPerN * 2 // 2 scenariusze
	fmt.Printf("📈 Łączna liczba testów do wykonania: %d\n", totalTests)
	fmt.Println()

	// Przygotowanie zadań
	tasks := make([]TestTask, 0, totalTests)
	taskIndex := 0

	for _, n := range nValues {
		for test := 1; test <= testsPerN; test++ {
			// Zadanie dla scenariusza uporządkowanego
			tasks = append(tasks, TestTask{
				N:       n,
				TestNum: test,
				Ordered: true,
				Seed:    int64(taskIndex * 12345), // Unikalny seed dla każdego zadania
			})
			taskIndex++

			// Zadanie dla scenariusza losowego
			tasks = append(tasks, TestTask{
				N:       n,
				TestNum: test,
				Ordered: false,
				Seed:    int64(taskIndex * 54321), // Unikalny seed dla każdego zadania
			})
			taskIndex++
		}
	}

	// Kanały do komunikacji
	taskCh := make(chan TestTask, 100)
	resultCh := make(chan TestResult, 100)
	progressCh := make(chan ProgressInfo, 10)

	var wg sync.WaitGroup

	// Uruchomienie workerów
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskCh {
				result := runSingleTestThreadSafe(task)
				resultCh <- result

				// Wysłanie informacji o postępie
				scenario := "random"
				if task.Ordered {
					scenario = "ordered"
				}
				progressCh <- ProgressInfo{
					CurrentN:    task.N,
					Scenario:    scenario,
					WorkerCount: numWorkers,
				}
			}
		}(i)
	}

	// Goroutine do śledzenia postępu
	go func() {
		completed := 0
		lastPrintTime := time.Now()
		printInterval := 2 * time.Second

		for progress := range progressCh {
			completed++

			// Wyświetlanie postępu co określony interwał czasowy
			if time.Since(lastPrintTime) >= printInterval || completed == totalTests {
				percentage := float64(completed) / float64(totalTests) * 100
				fmt.Printf("🔄 Postęp: %d/%d (%.1f%%) | Aktualnie: n=%d, scenariusz=%s | Wątki: %d\n",
					completed, totalTests, percentage, progress.CurrentN, progress.Scenario, progress.WorkerCount)
				lastPrintTime = time.Now()
			}

			if completed == totalTests {
				close(progressCh)
				return
			}
		}
	}()

	// Wysyłanie zadań do workerów
	fmt.Println("📤 Wysyłanie zadań do workerów...")
	go func() {
		for _, task := range tasks {
			taskCh <- task
		}
		close(taskCh)
	}()

	// Oczekiwanie na zakończenie wszystkich workerów
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Zbieranie wyników
	fmt.Println("📥 Zbieranie wyników...")
	for result := range resultCh {
		if result.Scenario == "ordered" {
			results.OrderedResults = append(results.OrderedResults, result)
		} else {
			results.RandomResults = append(results.RandomResults, result)
		}
	}

	fmt.Println("✅ Wszystkie testy wielowątkowe zakończone!")

	// Obliczanie podsumowań
	results.Summary = calculateSummary(results, nValues)

	return results
}

// calculateSummary oblicza średnie dla każdego n
func calculateSummary(results AllResults, nValues []int) Summary {
	summary := Summary{
		OrderedScenario: ScenarioSummary{AvgResults: make([]NSummary, 0)},
		RandomScenario:  ScenarioSummary{AvgResults: make([]NSummary, 0)},
	}

	for _, n := range nValues {
		// Średnie dla scenariusza uporządkowanego
		orderedResults := make([]TestResult, 0)
		for _, result := range results.OrderedResults {
			if result.N == n {
				orderedResults = append(orderedResults, result)
			}
		}

		if len(orderedResults) > 0 {
			summary.OrderedScenario.AvgResults = append(summary.OrderedScenario.AvgResults, calculateNSummary(n, orderedResults))
		}

		// Średnie dla scenariusza losowego
		randomResults := make([]TestResult, 0)
		for _, result := range results.RandomResults {
			if result.N == n {
				randomResults = append(randomResults, result)
			}
		}

		if len(randomResults) > 0 {
			summary.RandomScenario.AvgResults = append(summary.RandomScenario.AvgResults, calculateNSummary(n, randomResults))
		}
	}

	return summary
}

// calculateNSummary oblicza średnie dla konkretnego n
func calculateNSummary(n int, results []TestResult) NSummary {
	totalInsertComp, totalInsertPtr, totalInsertHeight := 0.0, 0.0, 0.0
	totalDeleteComp, totalDeletePtr, totalDeleteHeight := 0.0, 0.0, 0.0
	totalTime := 0.0

	for _, result := range results {
		totalInsertComp += result.InsertAvgComparisons
		totalInsertPtr += result.InsertAvgPointers
		totalInsertHeight += result.InsertAvgHeight
		totalDeleteComp += result.DeleteAvgComparisons
		totalDeletePtr += result.DeleteAvgPointers
		totalDeleteHeight += result.DeleteAvgHeight
		totalTime += result.TotalTime
	}

	count := float64(len(results))
	return NSummary{
		N:                    n,
		AvgInsertComparisons: totalInsertComp / count,
		AvgInsertPointers:    totalInsertPtr / count,
		AvgInsertHeight:      totalInsertHeight / count,
		AvgDeleteComparisons: totalDeleteComp / count,
		AvgDeletePointers:    totalDeletePtr / count,
		AvgDeleteHeight:      totalDeleteHeight / count,
		AvgTotalTime:         totalTime / count,
	}
}

// createAveragedResults tworzy skrócony format z uśrednionymi wynikami
func createAveragedResults(results AllResults, nValues []int) AveragedResults {
	averaged := AveragedResults{
		AveragedResults: make([]AveragedResult, 0),
	}

	for _, n := range nValues {
		// Średnie dla scenariusza uporządkowanego
		orderedResults := make([]TestResult, 0)
		for _, result := range results.OrderedResults {
			if result.N == n {
				orderedResults = append(orderedResults, result)
			}
		}

		if len(orderedResults) > 0 {
			avgResult := calculateAveragedResult(n, "ordered", orderedResults)
			averaged.AveragedResults = append(averaged.AveragedResults, avgResult)
		}

		// Średnie dla scenariusza losowego
		randomResults := make([]TestResult, 0)
		for _, result := range results.RandomResults {
			if result.N == n {
				randomResults = append(randomResults, result)
			}
		}

		if len(randomResults) > 0 {
			avgResult := calculateAveragedResult(n, "random", randomResults)
			averaged.AveragedResults = append(averaged.AveragedResults, avgResult)
		}
	}

	return averaged
}

// calculateAveragedResult oblicza średnie dla konkretnego n i scenariusza
func calculateAveragedResult(n int, scenario string, results []TestResult) AveragedResult {
	totalInsertComp, totalInsertPtr, totalInsertHeight := 0.0, 0.0, 0.0
	totalDeleteComp, totalDeletePtr, totalDeleteHeight := 0.0, 0.0, 0.0
	totalTime := 0.0

	for _, result := range results {
		totalInsertComp += result.InsertAvgComparisons
		totalInsertPtr += result.InsertAvgPointers
		totalInsertHeight += result.InsertAvgHeight
		totalDeleteComp += result.DeleteAvgComparisons
		totalDeletePtr += result.DeleteAvgPointers
		totalDeleteHeight += result.DeleteAvgHeight
		totalTime += result.TotalTime
	}

	count := float64(len(results))
	return AveragedResult{
		N:                    n,
		Scenario:             scenario,
		AvgInsertComparisons: totalInsertComp / count,
		AvgInsertPointers:    totalInsertPtr / count,
		AvgInsertHeight:      totalInsertHeight / count,
		AvgDeleteComparisons: totalDeleteComp / count,
		AvgDeletePointers:    totalDeletePtr / count,
		AvgDeleteHeight:      totalDeleteHeight / count,
		AvgTotalTime:         totalTime / count,
	}
}

// saveResultsToJSON zapisuje wyniki do pliku JSON
func saveResultsToJSON(results AllResults, filename string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// saveAveragedResultsToJSON zapisuje skrócone wyniki do pliku JSON
func saveAveragedResultsToJSON(results AveragedResults, filename string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// printSummary wyświetla podsumowanie wyników
func printSummary(summary Summary) {
	fmt.Printf("\n📊 PODSUMOWANIE WYDAJNOŚCI RB-BST\n")
	fmt.Printf("=================================\n")

	fmt.Printf("\n🔢 SCENARIUSZ UPORZĄDKOWANY (Wstawianie 1,2,3...n):\n")
	fmt.Printf("%-8s %-12s %-12s %-10s %-12s %-12s %-10s %-10s\n",
		"n", "Ins.Comp", "Ins.Ptr", "Ins.H", "Del.Comp", "Del.Ptr", "Del.H", "Czas[ms]")
	fmt.Printf("%s\n", strings.Repeat("-", 95))

	for _, avg := range summary.OrderedScenario.AvgResults {
		fmt.Printf("%-8d %-12.1f %-12.1f %-10.1f %-12.1f %-12.1f %-10.1f %-10.1f\n",
			avg.N, avg.AvgInsertComparisons, avg.AvgInsertPointers, avg.AvgInsertHeight,
			avg.AvgDeleteComparisons, avg.AvgDeletePointers, avg.AvgDeleteHeight, avg.AvgTotalTime)
	}

	fmt.Printf("\n🎲 SCENARIUSZ LOSOWY (Permutacja losowa):\n")
	fmt.Printf("%-8s %-12s %-12s %-10s %-12s %-12s %-10s %-10s\n",
		"n", "Ins.Comp", "Ins.Ptr", "Ins.H", "Del.Comp", "Del.Ptr", "Del.H", "Czas[ms]")
	fmt.Printf("%s\n", strings.Repeat("-", 95))

	for _, avg := range summary.RandomScenario.AvgResults {
		fmt.Printf("%-8d %-12.1f %-12.1f %-10.1f %-12.1f %-12.1f %-10.1f %-10.1f\n",
			avg.N, avg.AvgInsertComparisons, avg.AvgInsertPointers, avg.AvgInsertHeight,
			avg.AvgDeleteComparisons, avg.AvgDeletePointers, avg.AvgDeleteHeight, avg.AvgTotalTime)
	}
}

func runBenchmarkMain() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🌳 RB-BST Performance Benchmark (Wielowątkowy)")
	fmt.Println("==============================================")

	startTime := time.Now()
	results := runBenchmark()
	totalTime := time.Since(startTime)

	fmt.Printf("\n🎯 PODSUMOWANIE WYKONANIA\n")
	fmt.Printf("========================\n")
	fmt.Printf("✅ Testy zakończone w czasie: %v\n", totalTime)
	fmt.Printf("🔧 Użyto architektury wielowątkowej dla maksymalnej wydajności\n")
	fmt.Printf("📊 Przetestowano %d różnych wartości n\n", 10)
	fmt.Printf("🔄 Wykonano %d testów dla każdego scenariusza\n", 20)
	// Zapis do pliku JSON - pełny format
	fullFilename := "rbbst_benchmark_results_full.json"
	fmt.Printf("\n💾 Zapisywanie pełnych wyników do pliku %s...\n", fullFilename)

	err := saveResultsToJSON(results, fullFilename)
	if err != nil {
		fmt.Printf("❌ Błąd podczas zapisu pełnego pliku: %v\n", err)
		return
	}

	fmt.Printf("✅ Pełne wyniki zapisane pomyślnie\n")

	// Zapis do pliku JSON - skrócony format
	nValues := []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000}
	averagedResults := createAveragedResults(results, nValues)
	shortFilename := "rbbst_benchmark_results_short.json"
	fmt.Printf("💾 Zapisywanie skróconych wyników do pliku %s...\n", shortFilename)

	err = saveAveragedResultsToJSON(averagedResults, shortFilename)
	if err != nil {
		fmt.Printf("❌ Błąd podczas zapisu skróconego pliku: %v\n", err)
		return
	}

	fmt.Printf("✅ Skrócone wyniki zapisane pomyślnie\n")

	// Wyświetlanie podsumowania
	printSummary(results.Summary)
	fmt.Printf("\n📈 Łączna liczba wykonanych testów: %d\n", len(results.OrderedResults)+len(results.RandomResults))
	fmt.Printf("🚀 Testy wielowątkowe przyspieszyły wykonanie!\n")
	fmt.Printf("📁 Plik pełny: %s\n", fullFilename)
	fmt.Printf("📁 Plik skrócony: %s\n", shortFilename)
	fmt.Println("🎉 Wielowątkowy benchmark RB-BST zakończony pomyślnie!")
}
