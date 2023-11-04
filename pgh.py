import argparse

def main():
    parser = argparse.ArgumentParser()
    parser.description = "Simple utility for interactive with PostgreSQL"
    parser.add_argument("-h", "--host", required=True)
    parser.add_argument("-u", "--username", required=True)
    parser.add_argument("-p", "--password", required=True)
    parser.add_argument("-q", "--raw-sql")
    args = parser.parse_args()

if __name__ == '__main__':
    main()
