FROM python:3.9-slim

WORKDIR /app
COPY clean.py .
COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "clean.py"]
