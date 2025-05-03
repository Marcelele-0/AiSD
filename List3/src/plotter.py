import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Wczytaj dane
file_path = "results/select_comparison_results.csv"
df = pd.read_csv(file_path)

# Przekształć dane, by nadawały się do wykresów
df_melted = df.melt(id_vars=["n", "k"], 
                    value_vars=["my_select_comparisons", "randomized_select_comparisons", 
                                 "my_select_swaps", "randomized_select_swaps"], 
                    var_name="metric", value_name="value")

# Dodaj kolumny algorytm i rodzaj metryki (porównania/przestawienia)
df_melted["algorithm"] = df_melted["metric"].apply(lambda x: "MySelect" if "my_select" in x else "RandomizedSelect")
df_melted["type"] = df_melted["metric"].apply(lambda x: "Comparisons" if "comparisons" in x else "Swaps")

# Wykresy: dla każdego k osobny rysunek
k_values = df["k"].unique()

for k_val in k_values:
    subset = df_melted[df_melted["k"] == k_val]
    g = sns.relplot(data=subset, x="n", y="value", hue="algorithm", col="type", kind="line",
                    facet_kws={'sharey': False, 'sharex': True})
    g.fig.suptitle(f"Select Comparison (k = {k_val})", y=1.05)
    g.set_axis_labels("n", "Average Count")
    g.set_titles(col_template="{col_name}")
    plt.tight_layout()
    plt.savefig(f"results/select_comparison_k{k_val}.png")
    plt.close()

print("Plots saved to 'results/' folder.")
