import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv('../clean_train.csv')

bar_chart_data = df[["Overall Qual","Overall Cond", "SalePrice","Year Built","Year Remod/Add", "Yr Sold"]]