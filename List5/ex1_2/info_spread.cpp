/**
 * @file info_spread.cpp
 * @brief Implementation of optimal information spreading algorithm on MST trees
 * @details This program implements an algorithm that determines the optimal order
 *          for vertices to inform their children in a tree to minimize the number
 *          of rounds needed for information to reach all vertices.
 * @author Auto-documented
 * @date June 2025
 */

#include "mst_lib.h"  // Include MST functionality
#include <climits>

using Clock = chrono::high_resolution_clock;
using ms = chrono::duration<double, milli>;

/**
 * @brief Converts MST edges to adjacency list representation
 * @param mst_edges Vector of MST edges
 * @param n Number of vertices
 * @return Adjacency list representation of the tree
 */
vector<vector<int>> mst_to_adjacency_list(const vector<Edge>& mst_edges, int n) {
    vector<vector<int>> adj(n);
    for (const auto& edge : mst_edges) {
        adj[edge.u].push_back(edge.v);
        adj[edge.v].push_back(edge.u);
    }
    return adj;
}

/**
 * @brief Creates a rooted tree from undirected tree
 * @param adj Adjacency list of undirected tree
 * @param root Root vertex
 * @return Adjacency list where adj[v] contains children of v
 */
vector<vector<int>> create_rooted_tree(const vector<vector<int>>& adj, int root) {
    int n = adj.size();
    vector<vector<int>> children(n);
    vector<bool> visited(n, false);
    
    function<void(int)> dfs = [&](int v) {
        visited[v] = true;
        for (int u : adj[v]) {
            if (!visited[u]) {
                children[v].push_back(u);
                dfs(u);
            }
        }
    };
    
    dfs(root);
    return children;
}

/**
 * @brief Computes optimal spreading order for each vertex
 * @param children Adjacency list of rooted tree (children only)
 * @param v Current vertex
 * @return Number of rounds needed for subtree rooted at v
 * @details The algorithm works by:
 *          1. Computing rounds needed for each child's subtree
 *          2. Sorting children by rounds needed (descending)
 *          3. Informing children that need more rounds first
 *          4. Total rounds = max over all children of (position + child_rounds)
 */
int compute_optimal_spreading(const vector<vector<int>>& children, int v) {
    if (children[v].empty()) {
        return 1; // Leaf node needs 1 round (it's already informed)
    }
    
    // Calculate rounds needed for each child's subtree
    vector<pair<int, int>> child_rounds; // (rounds_needed, child_id)
    for (int child : children[v]) {
        int rounds = compute_optimal_spreading(children, child);
        child_rounds.push_back({rounds, child});
    }
    
    // Sort children by rounds needed (descending order)
    // We should inform children that need more rounds first
    sort(child_rounds.begin(), child_rounds.end(), greater<pair<int, int>>());
    
    int max_rounds = 0;
    for (int i = 0; i < child_rounds.size(); i++) {
        // Child at position i will be informed in round (i+1)
        // Total rounds for this child = round_when_informed + rounds_needed_for_subtree
        int total_rounds = (i + 1) + child_rounds[i].first;
        max_rounds = max(max_rounds, total_rounds);
    }
    
    return max_rounds;
}

/**
 * @brief Calculates the number of rounds needed for information spreading
 * @param mst_edges Edges of the MST
 * @param n Number of vertices
 * @param root Root vertex from which information starts
 * @return Number of rounds needed to inform all vertices
 */
int calculate_spreading_rounds(const vector<Edge>& mst_edges, int n, int root) {
    if (n == 1) return 1; // Single vertex case
    
    // Convert to adjacency list
    auto adj = mst_to_adjacency_list(mst_edges, n);
    
    // Create rooted tree
    auto children = create_rooted_tree(adj, root);
    
    // Compute optimal spreading
    return compute_optimal_spreading(children, root);
}

/**
 * @struct SpreadingStats
 * @brief Statistics for information spreading analysis
 */
struct SpreadingStats {
    double avg_rounds;
    int min_rounds;
    int max_rounds;
    int n; // Number of vertices
};

/**
 * @brief Performs average case analysis for a given graph size
 * @param n Number of vertices
 * @param rep Number of repetitions
 * @param seed Random seed
 * @return SpreadingStats with average, min, and max rounds
 */
SpreadingStats analyze_spreading_for_size(int n, int rep, unsigned int seed) {
    mt19937 rng(seed);
    
    vector<int> all_rounds;
    
    for (int r = 0; r < rep; r++) {
        // Generate random complete graph
        auto G = generate_graph(n, rng);
        
        // Generate MST using Kruskal's algorithm
        auto mst_result = kruskal(G);
        
        // Test with random root vertex
        uniform_int_distribution<int> root_dist(0, n - 1);
        int root = root_dist(rng);
        
        // Calculate rounds needed
        int rounds = calculate_spreading_rounds(mst_result.edges, n, root);
        all_rounds.push_back(rounds);
    }
    
    // Calculate statistics
    double avg_rounds = 0.0;
    for (int rounds : all_rounds) {
        avg_rounds += rounds;
    }
    avg_rounds /= rep;
    
    int min_rounds = *min_element(all_rounds.begin(), all_rounds.end());
    int max_rounds = *max_element(all_rounds.begin(), all_rounds.end());
    
    return {avg_rounds, min_rounds, max_rounds, n};
}

/**
 * @brief Main function - Analysis of information spreading on MST trees
 */
int main() {
    // Configuration parameters
    const int nMin = 100;           // Start with small graphs
    const int nMax = 5000;          // Maximum graph size to test
    const int step = 100;           // Step size for graph size increment
    const int rep = 50;            // Number of repetitions per graph size
    
    cout << "Starting information spreading analysis on MST trees...\n";
    cout << "Graph sizes: " << nMin << " to " << nMax << " (step=" << step << ", repetitions=" << rep << ")\n\n";
    
    // Open output file for results
    ofstream out("info_spread_results.csv");
    out << "n,avg_rounds,min_rounds,max_rounds\n";
    
    auto total_start = Clock::now();
    
    for (int n = nMin; n <= nMax; n += step) {
        cout << "Testing n = " << n << "... " << flush;
        
        auto start = Clock::now();
        auto stats = analyze_spreading_for_size(n, rep, 42 + n);
        auto end = Clock::now();
        
        double test_time = ms(end - start).count();
        
        // Output results
        out << n << "," << fixed << setprecision(2) << stats.avg_rounds << "," 
            << stats.min_rounds << "," << stats.max_rounds << "\n";
        
        cout << "avg=" << fixed << setprecision(2) << stats.avg_rounds 
             << ", min=" << stats.min_rounds 
             << ", max=" << stats.max_rounds 
             << " (took " << fixed << setprecision(1) << test_time << "ms)\n";
    }
    
    auto total_end = Clock::now();
    double total_time = ms(total_end - total_start).count() / 1000.0;
    
    cout << "\nAnalysis completed in " << fixed << setprecision(1) << total_time << " seconds\n";
    cout << "Results saved to info_spread_results.csv\n";
    
    out.close();
    return 0;
}
