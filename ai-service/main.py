from fastapi import FastAPI
from psycopg import connect

app = FastAPI()


def ping():
    try:
        with connect("host=db dbname=algtutor user=postgres password=postgres") as conn:
            with conn.cursor() as cur:
                cur.execute("SELECT 1")
                print(cur.fetchone())
    except Exception as e:
        print(e)


@app.get("/api/v1/hello")
async def hello():
    return {"hello": "world"}


@app.get("/api/v1/pingdb")
async def ping_db():
    ping()
    return {"msg": "done"}
