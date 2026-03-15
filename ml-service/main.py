from fastapi import FastAPI
from psycopg import connect

app = FastAPI()


@app.get("/api/v1/users")
async def get_users():
    users = {"users": []}
    try:
        with connect("host=db dbname=algtutor user=postgres password=postgres") as conn:
            with conn.cursor() as cur:
                cur.execute("SELECT * from app_user")
                for user in cur.fetchall():
                    users["users"].append({"first_name": user[1], "last_name": user[2]})
    except Exception as e:
        print(e)
    return users
