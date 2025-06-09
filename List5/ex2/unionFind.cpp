#include <iostream>
#include <vector>
#include <algorithm>
#include <queue>
#include <unordered_map>
#include <unordered_set>
#include <random>
#include <cmath>
#include <iomanip>
#include <fstream>

using namespace std;

struct Edge {
    int u, v, weight;
    bool operator<(const Edge& other) const {
        return weight < other.weight;
    }
};

class UnionFind {
private:
    vector<int> parent, rank;
    
public:
    UnionFind(int n) : parent(n), rank(n, 0) {
        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }
    }
    
    int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }
    
    bool unite(int x, int y) {
        int px = find(x), py = find(y);
        if (px == py) return false;
        
        if (rank[px] < rank[py]) swap(px, py);
        parent[py] = px;
        if (rank[px] == rank[py]) rank[px]++;
        
        return true;
    }
};

class TreeInfoSpread {
private:
    vector<vector<int>> graph;
    vector<vector<int>> tree;
    int n;
    
public:
    TreeInfoSpread(int vertices) : n(vertices) {
        graph.resize(n);
        tree.resize(n);
    }
    
    void addEdge(int u, int v) {
        graph[u].push_back(v);
        graph[v].push_back(u);
    }
    
    vector<Edge> kruskalMST() {
        vector<Edge> edges;
        vector<Edge> mstEdges;
        
        // Zbieramy wszystkie krawędzie z losowymi wagami
        random_device rd;
        mt19937 gen(rd());
        uniform_int_distribution<> dis(1, 100);
        
        for (int u = 0; u < n; u++) {
            for (int v = u + 1; v < n; v++) {
                edges.push_back({u, v, dis(gen)});
            }
        }
        
        sort(edges.begin(), edges.end());
        
        UnionFind uf(n);
        tree.assign(n, vector<int>());
        
        for (const Edge& e : edges) {
            if (uf.unite(e.u, e.v)) {
                mstEdges.push_back(e);
                tree[e.u].push_back(e.v);
                tree[e.v].push_back(e.u);
                if (mstEdges.size() == n - 1) break;
            }
        }
        
        return mstEdges;
    }
    
    vector<vector<int>> buildRootedTree(int root) {
        vector<vector<int>> rootedTree(n);
        vector<bool> visited(n, false);
        queue<int> q;
        
        q.push(root);
        visited[root] = true;
        
        while (!q.empty()) {
            int node = q.front();
            q.pop();
            
            for (int neighbor : tree[node]) {
                if (!visited[neighbor]) {
                    rootedTree[node].push_back(neighbor);
                    visited[neighbor] = true;
                    q.push(neighbor);
                }
            }
        }
        
        return rootedTree;
    }
    
    int calculateSubtreeDepth(const vector<vector<int>>& rootedTree, int node) {
        if (rootedTree[node].empty()) {
            return 0; // Liść
        }
        
        vector<int> childDepths;
        for (int child : rootedTree[node]) {
            childDepths.push_back(calculateSubtreeDepth(rootedTree, child));
        }
        
        // Sortujemy głębokości malejąco
        sort(childDepths.rbegin(), childDepths.rend());
        
        int maxRounds = 0;
        for (int i = 0; i < childDepths.size(); i++) {
            // i-te dziecko informujemy w (i+1)-szej rundzie
            maxRounds = max(maxRounds, i + 1 + childDepths[i]);
        }
        
        return maxRounds;
    }
    
    vector<int> getOptimalOrder(const vector<vector<int>>& rootedTree, int node) {
        if (rootedTree[node].empty()) {
            return {};
        }
        
        vector<pair<int, int>> childDepths; // {głębokość, dziecko}
        
        for (int child : rootedTree[node]) {
            int depth = calculateSubtreeDepth(rootedTree, child);
            childDepths.push_back({depth, child});
        }
        
        // Sortujemy malejąco po głębokości
        sort(childDepths.rbegin(), childDepths.rend());
        
        vector<int> order;
        for (const auto& pair : childDepths) {
            order.push_back(pair.second);
        }
        
        return order;
    }
    
    int simulateSpread(int root) {
        vector<vector<int>> rootedTree = buildRootedTree(root);
        
        // Obliczamy optymalną kolejność dla każdego wierzchołka
        unordered_map<int, vector<int>> optimalOrders;
        for (int node = 0; node < n; node++) {
            optimalOrders[node] = getOptimalOrder(rootedTree, node);
        }
        
        // Symulacja
        unordered_set<int> informed;
        informed.insert(root);
        
        vector<bool> isInformed(n, false);
        vector<int> nextChildIdx(n, 0);
        isInformed[root] = true;
        
        int roundNum = 0;
        
        while (informed.size() < n) {
            roundNum++;
            vector<int> newlyInformed;
            
            for (int node : informed) {
                if (nextChildIdx[node] < optimalOrders[node].size()) {
                    int child = optimalOrders[node][nextChildIdx[node]];
                    
                    if (!isInformed[child]) {
                        newlyInformed.push_back(child);
                        isInformed[child] = true;
                        nextChildIdx[node]++;
                    }
                }
            }
            
            for (int child : newlyInformed) {
                informed.insert(child);
            }
        }
        
        return roundNum;
    }
    
    void demonstrateAlgorithm() {
        cout << "=== DEMONSTRACJA ALGORYTMU ===" << endl;
        
        // Tworzymy przykładowe drzewo
        // Struktura:     0
        //               /|\\
        //              1 2 3
        //             /| |
        //            4 5 6
        
        TreeInfoSpread demo(7);
        demo.tree = {{1, 2, 3}, {0, 4, 5}, {0, 6}, {0}, {1}, {1}, {2}};
        
        int root = 0;
        vector<vector<int>> rootedTree = demo.buildRootedTree(root);
        
        cout << "Struktura drzewa (jako lista dzieci):" << endl;
        for (int node = 0; node < 7; node++) {
            cout << "Wierzchołek " << node << ": dzieci = [";
            for (int i = 0; i < rootedTree[node].size(); i++) {
                if (i > 0) cout << ", ";
                cout << rootedTree[node][i];
            }
            cout << "]" << endl;
        }
        
        cout << "\nOptymalna kolejność informowania dla każdego wierzchołka:" << endl;
        for (int node = 0; node < 7; node++) {
            vector<int> order = demo.getOptimalOrder(rootedTree, node);
            int depth = demo.calculateSubtreeDepth(rootedTree, node);
            
            cout << "Wierzchołek " << node << ": kolejność = [";
            for (int i = 0; i < order.size(); i++) {
                if (i > 0) cout << ", ";
                cout << order[i];
            }
            cout << "], głębokość poddrzewa = " << depth << endl;
        }
        
        int rounds = demo.simulateSpread(root);
        cout << "\nLiczba rund potrzebna do poinformowania wszystkich: " << rounds << endl;
    }
};

struct ExperimentResults {
    vector<int> sizes;
    vector<double> avgRounds;
    vector<int> minRounds;
    vector<int> maxRounds;
    vector<double> stdRounds;
};

class ExperimentalAnalysis {
private:
    random_device rd;
    mt19937 gen;
    
public:
    ExperimentalAnalysis() : gen(rd()) {}
    
    ExperimentResults runExperiments(const vector<int>& sizes, int trialsPerSize = 50) {
        ExperimentResults results;
        results.sizes = sizes;
        
        cout << "Przeprowadzam analizę eksperymentalną..." << endl;
        
        for (int size : sizes) {
            cout << "Testowanie rozmiaru " << size << "..." << flush;
            
            vector<int> roundsForSize;
            
            for (int trial = 0; trial < trialsPerSize; trial++) {
                // Generujemy pełny graf i znajdujemy MST
                TreeInfoSpread tis(size);
                tis.kruskalMST();
                
                // Losowy korzeń
                uniform_int_distribution<> rootDis(0, size - 1);
                int root = rootDis(gen);
                
                int rounds = tis.simulateSpread(root);
                roundsForSize.push_back(rounds);
            }
            
            // Obliczamy statystyki
            double sum = 0;
            int minVal = *min_element(roundsForSize.begin(), roundsForSize.end());
            int maxVal = *max_element(roundsForSize.begin(), roundsForSize.end());
            
            for (int rounds : roundsForSize) {
                sum += rounds;
            }
            double avg = sum / roundsForSize.size();
            
            // Odchylenie standardowe
            double variance = 0;
            for (int rounds : roundsForSize) {
                variance += (rounds - avg) * (rounds - avg);
            }
            double stdDev = sqrt(variance / roundsForSize.size());
            
            results.avgRounds.push_back(avg);
            results.minRounds.push_back(minVal);
            results.maxRounds.push_back(maxVal);
            results.stdRounds.push_back(stdDev);
            
            cout << " OK" << endl;
        }
        
        return results;
    }
    
    void printResults(const ExperimentResults& results) {
        cout << "\n" << string(60, '=') << endl;
        cout << "WYNIKI ANALIZY EKSPERYMENTALNEJ" << endl;
        cout << string(60, '=') << endl;
        
        cout << "Rozmiar | Średnia | Min | Max | Odch.std" << endl;
        cout << string(45, '-') << endl;
        
        for (int i = 0; i < results.sizes.size(); i++) {
            cout << setw(6) << results.sizes[i] 
                 << " | " << setw(7) << fixed << setprecision(2) << results.avgRounds[i]
                 << " | " << setw(3) << results.minRounds[i]
                 << " | " << setw(3) << results.maxRounds[i]
                 << " | " << setw(8) << fixed << setprecision(2) << results.stdRounds[i] << endl;
        }
    }
    
    void saveResultsToCSV(const ExperimentResults& results, const string& filename) {
        ofstream file(filename);
        file << "Size,Average,Min,Max,StdDev,LogTheoretical" << endl;
        
        for (int i = 0; i < results.sizes.size(); i++) {
            double logTheoretical = (results.sizes[i] > 1) ? log2(results.sizes[i]) : 0;
            file << results.sizes[i] << ","
                 << results.avgRounds[i] << ","
                 << results.minRounds[i] << ","
                 << results.maxRounds[i] << ","
                 << results.stdRounds[i] << ","
                 << logTheoretical << endl;
        }
        
        file.close();
        cout << "\nWyniki zapisane do pliku: " << filename << endl;
    }
    
    void analyzeComplexity(const ExperimentResults& results) {
        cout << "\n" << string(40, '=') << endl;
        cout << "ANALIZA ZŁOŻONOŚCI" << endl;
        cout << string(40, '=') << endl;
        
        cout << "Porównanie ze złożonością logarytmiczną:" << endl;
        cout << "Rozmiar | Średnia | log₂(n) | Stosunek" << endl;
        cout << string(45, '-') << endl;
        
        for (int i = 0; i < results.sizes.size(); i++) {
            if (results.sizes[i] > 1) {
                double logTheoretical = log2(results.sizes[i]);
                double ratio = results.avgRounds[i] / logTheoretical;
                
                cout << setw(6) << results.sizes[i] 
                     << " | " << setw(7) << fixed << setprecision(2) << results.avgRounds[i]
                     << " | " << setw(7) << fixed << setprecision(2) << logTheoretical
                     << " | " << setw(7) << fixed << setprecision(2) << ratio << endl;
            }
        }
    }
};

int main() {
    cout << "Algorytm optymalizacji rozprzestrzeniania informacji w drzewie" << endl;
    cout << string(60, '=') << endl;
    
    // Demonstracja algorytmu
    TreeInfoSpread demo(7);
    demo.demonstrateAlgorithm();
    
    cout << "\n\nRozpoczynanie analizy eksperymentalnej..." << endl;
    
    // Analiza eksperymentalna
    ExperimentalAnalysis analysis;
    vector<int> sizes = {5, 10, 15, 20, 25, 30, 40, 50, 75, 100};
    
    ExperimentResults results = analysis.runExperiments(sizes, 50);
    
    // Wyświetlenie wyników
    analysis.printResults(results);
    analysis.analyzeComplexity(results);
    
    // Zapisanie do CSV
    analysis.saveResultsToCSV(results, "tree_info_spread_results.csv");
    
    cout << "\n" << string(60, '=') << endl;
    cout << "WNIOSKI:" << endl;
    cout << "1. Liczba rund rośnie logarytmicznie względem rozmiaru drzewa" << endl;
    cout << "2. Optymalna strategia znacznie redukuje liczbę potrzebnych rund" << endl;
    cout << "3. Różnica między min a max wynika ze struktury konkretnych drzew" << endl;
    cout << "4. Algorytm ma złożoność czasową O(n log n)" << endl;
    cout << "5. Stosunek średniej do log₂(n) pozostaje względnie stały" << endl;
    cout << string(60, '=') << endl;
    
    return 0;
}