import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv('../clean_train.csv')

bar_chart_data = df[["Overall Qual","Overall Cond", "SalePrice","Year Built","Year Remod/Add", "Yr Sold"]]

temp = bar_chart_data.groupby(['Year Built', 'Overall Qual'])[['SalePrice']].sum()

temp.reset_index(inplace=True)

stacked_bar_chart = temp.pivot(index='Year Built', 
           columns='Overall Qual', 
           values='SalePrice').plot(kind='bar', 
                                    figsize=(20,10),
                                    stacked=True)

fig = stacked_bar_chart.get_figure()
fig.savefig("./temp.png") 
