from sys import exit
from os import listdir
from psycopg import connect


def execute_sqls_in_directory(directory: str):
    try:
        with connect("host=db dbname=algtutor user=postgres password=postgres") as conn:
            with conn.cursor() as cur:
                for sql_file in listdir(directory):
                    sql_file_path = directory + sql_file
                    with open(sql_file_path, "r") as opened_sql_file:
                        sql_text = opened_sql_file.read()
                        sql_commands = sql_text.split(";")
                        for command in sql_commands:
                            cur.execute(command)
    except Exception as e:
        print(e)
        raise e


def main():
    try:
        print("dropping tables")
        execute_sqls_in_directory("./drop_tables/")
        print("creating tables")
        execute_sqls_in_directory("./create_tables/")
        print("seeding tables")
        execute_sqls_in_directory("./seed_tables/")
    except Exception:
        return 1
    return 0


if __name__ == "__main__":
    exit(main())
