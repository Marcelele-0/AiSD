#!/usr/bin/env python3
"""
🌳 Generator wykresów porównawczych dla drzew BST, RB-BST i Splay Tree
Automatycznie wczytuje wyniki JSON i generuje wykresy porównawcze
"""

import json
import matplotlib.pyplot as plt
import numpy as np
import os
import sys
from pathlib import Path

# Konfiguracja
TREE_CONFIGS = {
    'BST': {
        'folder': 'ex1',
        'name': 'Binary Search Tree',
        'color': '#2E8B57',  # Zielony (SeaGreen)
        'marker': 'o'
    },
    'RB-BST': {
        'folder': 'ex3', 
        'name': 'Red-Black Tree',
        'color': '#DC143C',  # Czerwony (Crimson) 
        'marker': 's'
    },
    'Splay Tree': {
        'folder': 'ex5',
        'name': 'Splay Tree', 
        'color': '#4169E1',  # Niebieski (RoyalBlue)
        'marker': '^'
    }
}

def load_averaged_results(folder_path):
    """Wczytuje skrócone wyniki z pliku JSON"""
    json_file = folder_path / "averaged_results.json"
    if not json_file.exists():
        print(f"⚠️  Brak pliku {json_file}")
        return None
    
    try:
        with open(json_file, 'r', encoding='utf-8') as f:
            data = json.load(f)
        return data.get('averaged_results', [])
    except Exception as e:
        print(f"❌ Błąd wczytywania {json_file}: {e}")
        return None

def prepare_data():
    """Przygotowuje dane ze wszystkich drzew"""
    all_data = {}
    
    for tree_type, config in TREE_CONFIGS.items():
        folder_path = Path(config['folder'])
        results = load_averaged_results(folder_path)
        
        if results:
            # Rozdziel na scenariusze
            ordered_data = [r for r in results if r['scenario'] == 'ordered']
            random_data = [r for r in results if r['scenario'] == 'random']
            
            all_data[tree_type] = {
                'ordered': sorted(ordered_data, key=lambda x: x['n']),
                'random': sorted(random_data, key=lambda x: x['n']),
                'config': config
            }
            print(f"✅ Wczytano dane dla {tree_type}")
        else:
            print(f"❌ Nie udało się wczytać danych dla {tree_type}")
    
    return all_data

def create_comparison_chart(all_data, metric, title, ylabel, filename):
    """Tworzy wykres porównawczy dla danej metryki"""
    fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(15, 6))
    
    scenarios = [
        ('ordered', 'Scenariusz uporządkowany', ax1),
        ('random', 'Scenariusz losowy', ax2)
    ]
    
    for scenario, scenario_title, ax in scenarios:
        for tree_type, data in all_data.items():
            if scenario not in data:
                continue
                
            scenario_data = data[scenario]
            config = data['config']
            
            if not scenario_data:
                continue
            
            n_values = [item['n'] for item in scenario_data]
            metric_values = [item[metric] for item in scenario_data]
            
            ax.plot(n_values, metric_values, 
                   color=config['color'], 
                   marker=config['marker'],
                   linewidth=2, 
                   markersize=6,
                   label=config['name'])
        
        ax.set_xlabel('Rozmiar danych (n)')
        ax.set_ylabel(ylabel)
        ax.set_title(scenario_title)
        ax.legend()
        ax.grid(True, alpha=0.3)
    
    plt.suptitle(title, fontsize=16, fontweight='bold')
    plt.tight_layout()
    plt.savefig(f'charts_{filename}.png', dpi=300, bbox_inches='tight')
    plt.show()
    print(f"📊 Wykres zapisany jako: charts_{filename}.png")

def create_height_comparison(all_data):
    """Specjalny wykres dla wysokości drzew (skala logarytmiczna)"""
    fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(15, 6))
    
    scenarios = [
        ('ordered', 'Scenariusz uporządkowany', ax1),
        ('random', 'Scenariusz losowy', ax2)
    ]
    
    for scenario, scenario_title, ax in scenarios:
        for tree_type, data in all_data.items():
            if scenario not in data:
                continue
                
            scenario_data = data[scenario]
            config = data['config']
            
            if not scenario_data:
                continue
            
            n_values = [item['n'] for item in scenario_data]
            insert_heights = [item['avg_insert_height'] for item in scenario_data]
            delete_heights = [item['avg_delete_height'] for item in scenario_data]
            
            # Średnia wysokość
            avg_heights = [(i + d) / 2 for i, d in zip(insert_heights, delete_heights)]
            
            ax.plot(n_values, avg_heights, 
                   color=config['color'], 
                   marker=config['marker'],
                   linewidth=2, 
                   markersize=6,
                   label=config['name'])
        
        # Dodaj teoretyczną linię log(n)
        if n_values:
            theoretical_heights = [np.log2(n) for n in n_values]
            ax.plot(n_values, theoretical_heights, 
                   'k--', alpha=0.5, label='Teoretyczne log₂(n)')
        
        ax.set_xlabel('Rozmiar danych (n)')
        ax.set_ylabel('Średnia wysokość drzewa')
        ax.set_title(scenario_title)
        ax.legend()
        ax.grid(True, alpha=0.3)
        ax.set_yscale('log')
    
    plt.suptitle('🌳 Porównanie wysokości drzew (skala logarytmiczna)', fontsize=16, fontweight='bold')
    plt.tight_layout()
    plt.savefig('charts_height_comparison.png', dpi=300, bbox_inches='tight')
    plt.show()
    print("📊 Wykres wysokości zapisany jako: charts_height_comparison.png")

def create_time_comparison(all_data):
    """Wykres porównawczy czasów wykonania"""
    fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(15, 6))
    
    scenarios = [
        ('ordered', 'Scenariusz uporządkowany', ax1),
        ('random', 'Scenariusz losowy', ax2)
    ]
    
    for scenario, scenario_title, ax in scenarios:
        for tree_type, data in all_data.items():
            if scenario not in data:
                continue
                
            scenario_data = data[scenario]
            config = data['config']
            
            if not scenario_data:
                continue
            
            n_values = [item['n'] for item in scenario_data]
            times = [item['avg_total_time_ms'] for item in scenario_data]
            
            ax.plot(n_values, times, 
                   color=config['color'], 
                   marker=config['marker'],
                   linewidth=2, 
                   markersize=6,
                   label=config['name'])
        
        ax.set_xlabel('Rozmiar danych (n)')
        ax.set_ylabel('Czas wykonania (ms)')
        ax.set_title(scenario_title)
        ax.legend()
        ax.grid(True, alpha=0.3)
    
    plt.suptitle('⏱️ Porównanie czasów wykonania', fontsize=16, fontweight='bold')
    plt.tight_layout()
    plt.savefig('charts_time_comparison.png', dpi=300, bbox_inches='tight')
    plt.show()
    print("📊 Wykres czasów zapisany jako: charts_time_comparison.png")

def generate_summary_table(all_data):
    """Generuje tabelę podsumowującą"""
    print("\n" + "="*80)
    print("📋 PODSUMOWANIE WYNIKÓW")
    print("="*80)
    
    for tree_type, data in all_data.items():
        if 'ordered' not in data or 'random' not in data:
            continue
            
        print(f"\n🌳 {data['config']['name']}:")
        
        for scenario in ['ordered', 'random']:
            scenario_data = data[scenario]
            if not scenario_data:
                continue
                
            scenario_name = "Uporządkowany" if scenario == 'ordered' else "Losowy"
            print(f"  📊 Scenariusz {scenario_name}:")
            
            # Znajdź najgorszy przypadek (największe n)
            worst_case = max(scenario_data, key=lambda x: x['n'])
            print(f"    n={worst_case['n']}: Avg Comparisons={worst_case['avg_insert_comparisons']:.1f}, "
                  f"Avg Height={worst_case['avg_insert_height']:.1f}")

def main():
    """Główna funkcja programu"""
    print("🌳 Generator wykresów porównawczych drzew")
    print("="*50)
    
    # Sprawdź czy matplotlib jest dostępne
    try:
        import matplotlib.pyplot as plt
        import numpy as np
    except ImportError:
        print("❌ Błąd: Brak wymaganych bibliotek!")
        print("💡 Zainstaluj: pip install matplotlib numpy")
        return
    
    # Wczytaj dane
    all_data = prepare_data()
    
    if not all_data:
        print("❌ Brak danych do wygenerowania wykresów!")
        print("💡 Uruchom najpierw benchmarki używając run_all.bat")
        return
    
    # Generuj wykresy
    print("\n📊 Generowanie wykresów...")
    
    # 1. Porównania liczby porównań
    create_comparison_chart(
        all_data, 
        'avg_insert_comparisons',
        '🔍 Porównanie średniej liczby porównań (INSERT)',
        'Średnia liczba porównań',
        'insert_comparisons'
    )
    
    # 2. Porównania aktualizacji wskaźników
    create_comparison_chart(
        all_data,
        'avg_insert_pointers', 
        '🔗 Porównanie średniej liczby aktualizacji wskaźników (INSERT)',
        'Średnia liczba aktualizacji wskaźników',
        'insert_pointers'
    )
    
    # 3. Specjalny wykres wysokości
    create_height_comparison(all_data)
    
    # 4. Porównanie czasów
    create_time_comparison(all_data)
    
    # 5. Tabela podsumowująca
    generate_summary_table(all_data)
    
    print("\n✅ Wszystkie wykresy zostały wygenerowane!")
    print("📁 Pliki PNG zostały zapisane w bieżącym folderze")

if __name__ == "__main__":
    main()
