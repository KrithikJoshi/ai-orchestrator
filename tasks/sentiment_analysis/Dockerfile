FROM python:3.9-slim

WORKDIR /app
COPY sentiment.py .
COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt
RUN python -m textblob.download_corpora

CMD ["python", "sentiment.py"]
