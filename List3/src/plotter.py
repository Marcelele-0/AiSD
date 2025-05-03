import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import os

# Wczytaj dane
file_path = "results/select_comparison.csv"
df = pd.read_csv(file_path)

# Przekształć dane — jeden wiersz na jedną metrykę (comparisons lub swaps)
df_melted = df.melt(
    id_vars=["n", "k", "algorithm"],
    value_vars=["avg_comparisons", "avg_swaps"],
    var_name="metric_type",
    value_name="value"
)

# Uproszczenie nazw
df_melted["metric_type"] = df_melted["metric_type"].map({
    "avg_comparisons": "Comparisons",
    "avg_swaps": "Swaps"
})

# Unikalne wartości k
k_values = df["k"].unique()

# Tworzenie wykresów
os.makedirs("results/plots", exist_ok=True)

for k_val in k_values:
    subset = df_melted[df_melted["k"] == k_val]
    g = sns.relplot(
        data=subset,
        x="n", y="value", hue="algorithm", col="metric_type", kind="line",
        facet_kws={'sharey': False, 'sharex': True}
    )
    g.fig.suptitle(f"Select Comparison Metrics (k = {k_val})", y=1.05)
    g.set_axis_labels("n", "Average Count")
    g.set_titles(col_template="{col_name}")
    plt.tight_layout()
    plt.savefig(f"results/plots/select_comparison_k{k_val}.png")
    plt.close()

print("Plots saved to 'results/plots/' folder.")
