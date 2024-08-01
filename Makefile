.PHONY: bench check-regression

# Run benchmarks and save the results to benchmark.txt
bench:
	go test -bench=. -benchmem > benchmark.txt

# Check if benchmark results indicate a performance regression
check-regression:
	git fetch origin main || true

	if [ -f benchmark.txt ]; then \
		new_benchmark=$$(grep -E "go-benchmark-pipeline" benchmark.txt | awk '{print $$3}' | sed 's/ns\/op//g' | sed 's/s//g'); \
		\
		old_benchmark=$$(git show origin/main:benchmark.txt 2>/dev/null | grep -E "go-benchmark-pipeline" | awk '{print $$3}' | sed 's/ns\/op//g' | sed 's/s//g'); \
		\
		echo "New Benchmark: $$new_benchmark"; \
		echo "Old Benchmark: $$old_benchmark"; \
		\
		new_benchmark_num=$$(echo "$$new_benchmark" | awk '{print $1}'); \
		old_benchmark_num=$$(echo "$$old_benchmark" | awk '{print $1}'); \
		\
		if [ -z "$$new_benchmark_num" ] || [ -z "$$old_benchmark_num" ]; then \
			echo "Benchmark values could not be parsed. Skipping performance regression check."; \
			exit 1; \
		fi; \
		\
		echo "New Benchmark (numeric): $$new_benchmark_num"; \
		echo "Old Benchmark (numeric): $$old_benchmark_num"; \
		\
		threshold=$$(echo "$$old_benchmark_num * 1.2" | bc -l); \
		regression=$$(echo "$$new_benchmark_num > $$threshold" | bc -l); \
		if [  $$regression -eq 1 ]; then \
			echo "Performance regression detected!"; \
			echo "New Benchmark: $$new_benchmark seconds"; \
			echo "Old Benchmark: $$old_benchmark seconds"; \
			exit 1; \
		else \
			echo "No performance regression detected."; \
		fi; \
	else \
		echo "Benchmark file not found. Skipping performance regression check."; \
		exit 1; \
	fi
