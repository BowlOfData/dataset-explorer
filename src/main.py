from pandasArrow import pandasArrow

def main():
    """
    this function to be used mainly to run the applicaiton via command line and for testing
    """

    pa = pandasArrow('/Users/marco/code/dataset-explorer/examples/crx.csv')
    pa.head()
    print(pa.dtype())


if __name__ == "__main__":
    main()