from sys import exit
from os import listdir
from psycopg import connect


def create_tables():
    base_dir = "./tables/"
    try:
        with connect("host=db dbname=algtutor user=postgres password=postgres") as conn:
            with conn.cursor() as cur:
                for sql_file in listdir(base_dir):
                    sql_file_path = base_dir + sql_file
                    with open(sql_file_path, "r") as opened_sql_file:
                        sql_text = opened_sql_file.read()
                        sql_commands = sql_text.split(";")
                        for command in sql_commands:
                            cur.execute(command)
    except Exception as e:
        print(e)
        raise e


def seed_tables():
    base_dir = "./seed/"
    try:
        with connect("host=db dbname=algtutor user=postgres password=postgres") as conn:
            with conn.cursor() as cur:
                for sql_file in listdir(base_dir):
                    sql_file_path = base_dir + sql_file
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
        print("creating tables")
        create_tables()
        # print("seeding tables")
        # seed_tables()
    except Exception:
        return 1
    return 0


if __name__ == "__main__":
    exit(main())
