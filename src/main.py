from pandasArrow import pandasArrow

def main():
    """
    this function to be used mainly to run the applicaiton via command line and for testing
    """

    pa = pandasArrow('../examples/crx.csv', header=None)
    pa.head()
    print(pa.dtype())

    print("Numeric DF")
    print(pa.identify_numeric())
    
    print("DropNA")
    print(pa.dropna())    
  
    # Fill NaN values with the mean of each column
    print("Fill NaN values with the mean of each column")
    print(pa.fillna_with_mean_of_column())   

    print("Corr Matrix")
    print(pa.corr())   

if __name__ == "__main__":
    main()