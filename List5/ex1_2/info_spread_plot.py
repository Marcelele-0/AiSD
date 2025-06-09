#!/usr/bin/env python3
"""
Plot information spreading analysis results.
Visualizes average, minimum, and maximum rounds needed for information spreading
on MST trees as a function of tree size.
"""

import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
import sys
import os

def plot_info_spread_results(csv_file='info_spread_results.csv', output_file='info_spread_plot.png'):
    """
    Plot information spreading analysis results.
    
    Args:
        csv_file: Path to CSV file with results
        output_file: Output file for the plot
    """
    
    # Check if results file exists
    if not os.path.exists(csv_file):
        print(f"Error: Results file '{csv_file}' not found!")
        print("Please run the info_spread program first to generate results.")
        return False
    
    try:
        # Read the data
        df = pd.read_csv(csv_file)
        print(f"Loaded {len(df)} data points from {csv_file}")
        
        # Create the plot
        plt.figure(figsize=(14, 10))
        
        # Create subplot layout
        gs = plt.GridSpec(2, 2, height_ratios=[2, 1], width_ratios=[1, 1])
        
        # Main plot
        ax1 = plt.subplot(gs[0, :])
        
        # Plot average, min, and max rounds
        ax1.plot(df['n'], df['avg_rounds'], 'b-', linewidth=2.5, marker='o', markersize=5, 
                label='Average rounds', alpha=0.8)
        ax1.plot(df['n'], df['min_rounds'], 'g--', linewidth=2, marker='^', markersize=4, 
                label='Minimum rounds', alpha=0.8)
        ax1.plot(df['n'], df['max_rounds'], 'r--', linewidth=2, marker='v', markersize=4, 
                label='Maximum rounds', alpha=0.8)
        
        # Fill area between min and max
        ax1.fill_between(df['n'], df['min_rounds'], df['max_rounds'], 
                        alpha=0.2, color='gray', label='Min-Max range')
        
        # Customize main plot
        ax1.set_xlabel('Number of vertices (n)', fontsize=12)
        ax1.set_ylabel('Number of rounds', fontsize=12)
        ax1.set_title('Information Spreading Analysis on MST Trees\n'
                     'Optimal rounds needed to inform all vertices from random root', 
                     fontsize=14, fontweight='bold')
        ax1.grid(True, alpha=0.3)
        ax1.legend(fontsize=11)
        
        # Growth rate analysis (subplot 1)
        ax2 = plt.subplot(gs[1, 0])
        if len(df) > 1:
            growth_rates = np.diff(df['avg_rounds']) / np.diff(df['n'])
            ax2.plot(df['n'][1:], growth_rates, 'purple', marker='o', markersize=3)
            ax2.set_xlabel('Number of vertices (n)', fontsize=10)
            ax2.set_ylabel('Growth rate\n(rounds/vertex)', fontsize=10)
            ax2.set_title('Average Growth Rate', fontsize=11)
            ax2.grid(True, alpha=0.3)
        
        # Logarithmic analysis (subplot 2)
        ax3 = plt.subplot(gs[1, 1])
        log_n = np.log2(df['n'])
        ax3.scatter(log_n, df['avg_rounds'], alpha=0.6, color='orange', s=20)
        
        # Fit logarithmic trend
        coeffs = np.polyfit(log_n, df['avg_rounds'], 1)
        trend_line = coeffs[0] * log_n + coeffs[1]
        ax3.plot(log_n, trend_line, 'red', linestyle='--', alpha=0.8, linewidth=2)
        
        ax3.set_xlabel('log₂(n)', fontsize=10)
        ax3.set_ylabel('Average rounds', fontsize=10)
        ax3.set_title('Logarithmic Relationship', fontsize=11)
        ax3.grid(True, alpha=0.3)
        
        # Add correlation coefficient
        correlation = np.corrcoef(log_n, df['avg_rounds'])[0, 1]
        ax3.text(0.05, 0.95, f'R² = {correlation**2:.3f}', transform=ax3.transAxes, 
                fontsize=10, verticalalignment='top',
                bbox=dict(boxstyle='round', facecolor='white', alpha=0.8))
        
        # Calculate and display statistics
        avg_growth_rate = (df['avg_rounds'].iloc[-1] - df['avg_rounds'].iloc[0]) / (df['n'].iloc[-1] - df['n'].iloc[0])
        max_avg_rounds = df['avg_rounds'].max()
        min_avg_rounds = df['avg_rounds'].min()
        
        # Add text box with statistics on main plot
        stats_text = f'Statistics:\n'
        stats_text += f'Avg growth rate: {avg_growth_rate:.4f} rounds/vertex\n'
        stats_text += f'Max avg rounds: {max_avg_rounds:.1f}\n'
        stats_text += f'Min avg rounds: {min_avg_rounds:.1f}\n'
        stats_text += f'Log₂(n) correlation: {correlation:.3f}'
        
        ax1.text(0.02, 0.98, stats_text, transform=ax1.transAxes, 
                fontsize=10, verticalalignment='top',
                bbox=dict(boxstyle='round', facecolor='white', alpha=0.9))
        
        # Tight layout and save
        plt.tight_layout()
        plt.savefig(output_file, dpi=300, bbox_inches='tight')
        print(f"Plot saved as {output_file}")
        
        # Display insights
        print("\nAnalysis Insights:")
        print(f"- Average rounds grow at rate of {avg_growth_rate:.4f} rounds per additional vertex")
        print(f"- Range varies from {min_avg_rounds:.1f} to {max_avg_rounds:.1f} average rounds")
        print(f"- Correlation with log₂(n): {correlation:.3f}")
        
        if correlation > 0.9:
            print("  → Strong logarithmic growth pattern detected!")
            print("  → This suggests the spreading rounds grow as O(log n)")
        elif correlation > 0.7:
            print("  → Moderate logarithmic growth pattern")
        else:
            print("  → Growth pattern may not be purely logarithmic")
        
        # Theoretical analysis
        theoretical_log_n = np.log2(df['n'])
        print(f"- Theoretical log₂({df['n'].iloc[-1]}) = {theoretical_log_n.iloc[-1]:.2f}")
        print(f"- Actual avg rounds for n={df['n'].iloc[-1]}: {df['avg_rounds'].iloc[-1]:.2f}")
        
        plt.show()
        return True
        
    except Exception as e:
        print(f"Error processing results: {e}")
        return False

def main():
    """Main function to run the plotting script."""
    
    # Check for command line arguments
    csv_file = 'info_spread_results.csv'
    output_file = 'info_spread_plot.png'
    
    if len(sys.argv) > 1:
        csv_file = sys.argv[1]
    if len(sys.argv) > 2:
        output_file = sys.argv[2]
    
    print("Information Spreading Analysis Plotter")
    print("=" * 50)
    
    success = plot_info_spread_results(csv_file, output_file)
    
    if not success:
        print("\nTo generate results, compile and run:")
        print("g++ -std=c++17 -O2 info_spread.cpp -o info_spread")
        print("./info_spread")
        sys.exit(1)

if __name__ == "__main__":
    main()
