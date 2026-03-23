# syntax=docker/dockerfile:1

FROM python:3.12

WORKDIR /app

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY drop_tables ./drop_tables

COPY create_tables ./create_tables

COPY seed_tables ./seed_tables

COPY *.py ./

CMD ["python", "main.py"]