/**
 * @file mst_lib.h
 * @brief Header file with MST functionality (without main function)
 */
#pragma once

#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>
#include <random>
#include <chrono>
#include <fstream>
#include <iomanip>
#include <functional>

using namespace std;

/**
 * @struct Edge
 * @brief Represents an edge in a graph with two vertices and a weight
 */
struct Edge {
    int u;      ///< First vertex of the edge
    int v;      ///< Second vertex of the edge
    double w;   ///< Weight of the edge
    
    bool operator<(const Edge &other) const {
        return w < other.w;
    }
};

/**
 * @struct DSU
 * @brief Disjoint Set Union (Union-Find) data structure for Kruskal's algorithm
 */
struct DSU {
    vector<int> parent;
    vector<int> rank;
    
    DSU(int n) : parent(n), rank(n, 0) {
        for (int i = 0; i < n; ++i) parent[i] = i;
    }
    
    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    
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
 * @struct MSTResult
 * @brief Structure to hold both cost and edges of MST
 */
struct MSTResult {
    double cost;
    vector<Edge> edges;
};

/**
 * @brief Generates a complete graph with random edge weights
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
 * @brief Prim's algorithm for finding MST
 */
MSTResult prim(const vector<vector<double>> &G) {
    int n = static_cast<int>(G.size());
    vector<bool> inMST(n, false);
    vector<double> key(n, 1e9);
    vector<int> parent(n, -1);
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
        
        if (parent[u] != -1) {
            mst_edges.push_back({parent[u], u, G[parent[u]][u]});
        }
        
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
 * @brief Kruskal's algorithm for finding MST
 */
MSTResult kruskal(const vector<vector<double>> &G) {
    int n = static_cast<int>(G.size());
    vector<Edge> edges;
    
    for (int i = 0; i < n; ++i)
        for (int j = i + 1; j < n; ++j)
            edges.push_back({i, j, G[i][j]});

    sort(edges.begin(), edges.end());
    DSU dsu(n);
    double cost = 0.0;
    vector<Edge> mst_edges;
    
    for (const auto &e : edges) {
        if (dsu.unite(e.u, e.v)) {
            cost += e.w;
            mst_edges.push_back(e);
        }
    }
    return {cost, mst_edges};
}
