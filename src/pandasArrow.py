import pandas as pd
import pyarrow as pa

class pandasArrow():

    def __init__(self, data, header) -> None:
        self.df = pd.read_csv(data, dtype_backend="pyarrow", header=header)

    def head(self) -> pd.DataFrame:
        print(self.df.head())

    def dtype(self) -> object:
        return self.df.dtypes 

    def dropna(self) -> object:
        self.df = self.df.convert_dtypes()
        self.df = self.df.dropna(axis=0, how='all')
        self.df = self.df.dropna(axis=1, how='all')
        self.df = self.df.convert_dtypes(dtype_backend='pyarrow')
        return self.df

    def identify_numeric(self) -> object:
        self.df = self.df.apply(pd.to_numeric, errors='coerce')
        return self.df
    
    def fillna_with_mean_of_column(self) -> object:
        self.df = self.df.fillna(self.df.mean())
        return self.df
    
    def corr(self) -> object:
        return self.df.corr()