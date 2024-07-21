import pandas as pd
import pyarrow as pa

class pandasArrow():

    def __init__(self, data) -> None:
        self.df = pd.read_csv(data, dtype_backend="pyarrow")

    def head(self) -> pd.DataFrame:
        print(self.df.head())

    def dtype(self) -> object:
        return self.df.dtypes