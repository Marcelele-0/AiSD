/**
 * @file mst.cpp
 * @brief Implementation and performance comparison of Minimum Spanning Tree algorithms
 * @details This program implements both Prim's and Kruskal's algorithms for finding
 *          minimum spanning trees and compares their performance on randomly generated graphs.
 * @author Auto-documented
 * @date June 2025
 */

#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>
#include <random>
#include <chrono>
#include <fstream>
#include <iomanip>
#include <thread>
#include <future>
#include <mutex>

using namespace std;
using Clock = chrono::high_resolution_clock;
using ms = chrono::duration<double, milli>;

/**
 * @struct Edge
 * @brief Represents an edge in a graph with two vertices and a weight
 * @details Used primarily in Kruskal's algorithm for storing and sorting edges by weight
 */
struct Edge {
    int u;      ///< First vertex of the edge
    int v;      ///< Second vertex of the edge
    double w;   ///< Weight of the edge
    
    /**
     * @brief Comparison operator for sorting edges by weight
     * @param other The other edge to compare with
     * @return true if this edge has smaller weight than other
     */
    bool operator<(const Edge &other) const {
        return w < other.w;
    }
};

/**
 * @struct DSU
 * @brief Disjoint Set Union (Union-Find) data structure for Kruskal's algorithm
 * @details Implements path compression and union by rank optimizations
 *          for efficient cycle detection in MST construction
 */
struct DSU {
    vector<int> parent; ///< Parent array for union-find structure
    vector<int> rank;   ///< Rank array for union by rank optimization
    
    /**
     * @brief Constructor that initializes DSU for n elements
     * @param n Number of elements (vertices) in the disjoint set
     */
    DSU(int n) : parent(n), rank(n, 0) {
        for (int i = 0; i < n; ++i) parent[i] = i;
    }
    
    /**
     * @brief Find operation with path compression
     * @param x Element to find the root of
     * @return Root representative of the set containing x
     */
    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    
    /**
     * @brief Union operation with union by rank
     * @param x First element to unite
     * @param y Second element to unite
     * @return true if elements were in different sets and were united, false if already connected
     */
    bool unite(int x, int y) {
        int xr = find(x), yr = find(y);
        if (xr == yr) return false;
        if (rank[xr] < rank[yr]) parent[xr] = yr;
        else {
            parent[yr] = xr;
            if (rank[xr] == rank[yr]) rank[xr]++;
        }
        return true;
    }
};

/**
 * @brief Generates a complete graph with random edge weights
 * @param n Number of vertices in the graph
 * @param rng Reference to random number generator for reproducible results
 * @return 2D adjacency matrix representing the complete graph with random weights [0,1)
 * @details Creates a symmetric adjacency matrix where G[i][j] = G[j][i] represents
 *          the weight of edge between vertices i and j. Self-loops have weight 0.
 * @complexity Time: O(n²), Space: O(n²)
 */
vector<vector<double>> generate_graph(int n, mt19937 &rng) {
    uniform_real_distribution<double> dist(0.0, 1.0);
    vector<vector<double>> G(n, vector<double>(n));
    for (int i = 0; i < n; ++i)
        for (int j = i + 1; j < n; ++j)
            G[i][j] = G[j][i] = dist(rng);
    return G;
}

/**
 * @struct MSTResult
 * @brief Structure to hold both cost and edges of MST
 */
struct MSTResult {
    double cost;                ///< Total cost of the MST
    vector<Edge> edges;         ///< Edges that form the MST
};

/**
 * @brief Prim's algorithm for finding Minimum Spanning Tree
 * @param G Const reference to adjacency matrix representing the graph
 * @return MSTResult containing cost and edges of the minimum spanning tree
 * @details Implements Prim's algorithm using a priority queue (min-heap).
 *          Starts from vertex 0 and greedily adds the minimum weight edge
 *          that connects a vertex in MST to a vertex outside MST.
 * @complexity Time: O(V² log V) with adjacency matrix, Space: O(V)
 * @note Uses structured bindings (C++17 feature)
 */
MSTResult prim(const vector<vector<double>> &G) {
    int n = static_cast<int>(G.size());
    vector<bool> inMST(n, false);           // Track vertices included in MST
    vector<double> key(n, 1e9);             // Minimum weight to connect each vertex
    vector<int> parent(n, -1);              // Track parent to reconstruct MST
    priority_queue<pair<double, int>, vector<pair<double, int>>, greater<>> pq;
    
    key[0] = 0;
    pq.emplace(0.0, 0);
    double cost = 0.0;
    vector<Edge> mst_edges;

    while (!pq.empty()) {
        auto [w, u] = pq.top(); 
        pq.pop();
        
        if (inMST[u]) continue;
        
        inMST[u] = true;
        cost += w;
        
        // Add edge to MST (except for root)
        if (parent[u] != -1) {
            mst_edges.push_back({parent[u], u, G[parent[u]][u]});
        }
        
        // Update keys of adjacent vertices
        for (int v = 0; v < n; ++v) {
            if (!inMST[v] && G[u][v] < key[v]) {
                key[v] = G[u][v];
                parent[v] = u;
                pq.emplace(key[v], v);
            }
        }
    }
    return {cost, mst_edges};
}

/**
 * @brief Kruskal's algorithm for finding Minimum Spanning Tree
 * @param G Const reference to adjacency matrix representing the graph
 * @return MSTResult containing cost and edges of the minimum spanning tree
 * @details Implements Kruskal's algorithm using edge sorting and Union-Find.
 *          Sorts all edges by weight and greedily adds edges that don't create cycles.
 * @complexity Time: O(E log E) where E = V(V-1)/2 for complete graph, Space: O(E + V)
 * @note More efficient for sparse graphs, but this implementation uses complete graphs
 */
MSTResult kruskal(const vector<vector<double>> &G) {
    int n = static_cast<int>(G.size());
    vector<Edge> edges;
    
    // Convert adjacency matrix to edge list
    for (int i = 0; i < n; ++i)
        for (int j = i + 1; j < n; ++j)
            edges.push_back({i, j, G[i][j]});

    sort(edges.begin(), edges.end());  // Sort edges by weight
    DSU dsu(n);
    double cost = 0.0;
    vector<Edge> mst_edges;
    
    // Process edges in order of increasing weight
    for (const auto &e : edges) {
        if (dsu.unite(e.u, e.v)) {  // If vertices are in different components
            cost += e.w;
            mst_edges.push_back(e);
        }
    }
    return {cost, mst_edges};
}

/**
 * @struct TestResult
 * @brief Structure to hold test results for a specific graph size
 */
struct TestResult {
    int n;                    ///< Graph size (number of vertices)
    double prim_avg_time;     ///< Average time for Prim's algorithm
    double kruskal_avg_time;  ///< Average time for Kruskal's algorithm
};

/**
 * @brief Test MST algorithms for a specific graph size
 * @param n Number of vertices in the graph
 * @param rep Number of repetitions for averaging
 * @param seed Random seed for this test
 * @return TestResult containing average execution times
 */
TestResult test_graph_size(int n, int rep, unsigned int seed) {
    mt19937 rng(seed);
    double prim_total = 0.0;
    double kruskal_total = 0.0;
    
    for (int r = 0; r < rep; ++r) {
        auto G = generate_graph(n, rng);
        
        // Time Prim's algorithm
        auto start = Clock::now();
        prim(G).cost;  // Only need cost for timing
        auto end = Clock::now();
        prim_total += ms(end - start).count();
        
        // Time Kruskal's algorithm
        start = Clock::now();
        kruskal(G).cost;  // Only need cost for timing
        end = Clock::now();
        kruskal_total += ms(end - start).count();
    }
    
    return {n, prim_total / rep, kruskal_total / rep};
}

/**
 * @brief Main function - Performance comparison of MST algorithms
 * @return 0 on successful execution
 * @details Compares Prim's and Kruskal's algorithms on complete graphs of varying sizes.
 *          Generates performance data and outputs to CSV file for analysis.
 *          Tests graph sizes from 100 to 1000 vertices in steps of 100.
 *          Each test is repeated 5 times and averaged for statistical reliability.
 * @note Results are saved to "ex1/results.csv" in the current directory
 */
int main() {
    // Configuration parameters
    const int nMin = 1000;          ///< Minimum graph size to test
    const int nMax = 20000;         ///< Maximum graph size to test
    const int step = 1000;           ///< Step size for graph size increment
    const int rep = 20;             ///< Number of repetitions per graph size
    
    // Determine number of threads (use 4 cores less than available)
    const unsigned int available_cores = thread::hardware_concurrency();
    const unsigned int num_threads = max(1u, available_cores > 10 ? available_cores - 10 : 1);
    cout << "Available cores: " << available_cores << ", using " << num_threads << " threads for parallel execution\n";
    
    // Prepare test cases
    vector<int> test_sizes;
    for (int n = nMin; n <= nMax; n += step) {
        test_sizes.push_back(n);
    }
    
    // Open output file for results
    ofstream out("results.csv");
    out << "n,prim_time,kruskal_time\n";
    
    // Execute tests in parallel batches
    vector<TestResult> results;
    results.reserve(test_sizes.size());
    
    cout << "Starting performance tests...\n";
    cout << "Testing graph sizes: " << nMin << " to " << nMax << " (step=" << step << ", repetitions=" << rep << ")\n";
    cout << "Total test cases: " << test_sizes.size() << "\n\n";
    auto total_start = Clock::now();
    
    // Process tests in batches to avoid creating too many threads
    for (size_t i = 0; i < test_sizes.size(); i += num_threads) {
        vector<future<TestResult>> futures;
        
        // Launch threads for current batch
        for (size_t j = i; j < min(i + num_threads, test_sizes.size()); ++j) {
            int n = test_sizes[j];
            unsigned int seed = 42 + j;  // Different seed for each test
            
            futures.push_back(async(launch::async, test_graph_size, n, rep, seed));
        }
        
        // Collect results from current batch
        for (auto& future : futures) {
            results.push_back(future.get());
        }
        
        // Progress update with percentage and current test sizes
        int current_batch = (i / num_threads + 1);
        int total_batches = ((test_sizes.size() + num_threads - 1) / num_threads);
        double progress_percent = (double)current_batch / total_batches * 100.0;
        
        cout << "Progress: " << fixed << setprecision(1) << progress_percent << "% "
             << "(batch " << current_batch << "/" << total_batches << ") - ";
        
        // Show which graph sizes were tested in this batch
        cout << "tested n = ";
        for (size_t j = i; j < min(i + num_threads, test_sizes.size()); ++j) {
            if (j > i) cout << ", ";
            cout << test_sizes[j];
        }
        cout << "\n";
    }
    
    // Sort results by graph size (in case they completed out of order)
    sort(results.begin(), results.end(), [](const TestResult& a, const TestResult& b) {
        return a.n < b.n;
    });
    
    // Write results to file
    for (const auto& result : results) {
        out << result.n << ','
            << fixed << setprecision(3) << result.prim_avg_time << ','
            << fixed << setprecision(3) << result.kruskal_avg_time << '\n';
        cout << "n=" << result.n 
             << " | Prim: " << fixed << setprecision(1) << result.prim_avg_time << "ms"
             << " | Kruskal: " << fixed << setprecision(1) << result.kruskal_avg_time << "ms\n";
    }
    
    auto total_end = Clock::now();
    double total_time = ms(total_end - total_start).count() / 1000.0; // seconds
    
    cout << "\nAll tests completed in " << fixed << setprecision(1) << total_time << " seconds\n";
    out.close();
    return 0;
}
